// SPDX-License-Identifier: Zlib

package sdl2

import (
	"fmt"
	"log/slog"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/renderers/internal/overlay"
	"github.com/TotallyGamerJet/clay/renderers/internal/shapes"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	FontId uint32
	// Font is used for text without a font size, and for all text if Data is nil.
	Font *ttf.Font
	// Data is the font file, used to open the font at the font sizes of the text.
	// SDL reads from it while the font is in use, so it must not be modified.
	Data []byte
}

type sizedFontKey struct {
	data *byte
	size uint16
}

var sizedFonts = map[sizedFontKey]*ttf.Font{}

// sized returns the font at the given size.
func (f *Font) sized(size uint16) (*ttf.Font, error) {
	if size == 0 || len(f.Data) == 0 {
		return f.Font, nil
	}
	key := sizedFontKey{data: &f.Data[0], size: size}
	if font, ok := sizedFonts[key]; ok {
		return font, nil
	}
	rw, err := sdl.RWFromMem(f.Data)
	if err != nil {
		return nil, err
	}
	font, err := ttf.OpenFontRW(rw, 1, int(size))
	if err != nil {
		return nil, err
	}
	sizedFonts[key] = font
	return font, nil
}

func MeasureText(text string, config *clay.TextElementConfig, userData any) clay.Dimensions {
	fonts := *userData.(*[]Font)
	font, err := fonts[config.FontId].sized(config.FontSize)
	if err != nil {
		panic(fmt.Errorf("sdl2: failed to open font: %w", err))
	}
	width, height, err := font.SizeUTF8(text)
	if err != nil {
		panic(fmt.Errorf("sdl2: failed to measure text: %w", err))
	}
	width += int(config.LetterSpacing) * max(len([]rune(text))-1, 0)
	if config.LineHeight > 0 {
		height = int(config.LineHeight)
	}

	return clay.Dimensions{
		Width:  float32(width),
		Height: float32(height),
	}
}

// drawText draws one line of text, spacing the characters out by letterSpacing and
// centering them in the line box when the text sets a line height.
func drawText(renderer *sdl.Renderer, font *ttf.Font, str string, box clay.BoundingBox, letterSpacing int, color sdl.Color) error {
	draw := func(str string, x float32) (float32, error) {
		surface, err := font.RenderUTF8Blended(str, color)
		if err != nil {
			return 0, err
		}
		defer surface.Free()
		texture, err := renderer.CreateTextureFromSurface(surface)
		if err != nil {
			return 0, err
		}
		defer texture.Destroy()
		y := box.Y
		if float32(surface.H) < box.Height {
			y += (box.Height - float32(surface.H)) / 2
		}
		destination := sdl.Rect{X: int32(x), Y: int32(y), W: surface.W, H: surface.H}
		return float32(surface.W), renderer.Copy(texture, nil, &destination)
	}

	if letterSpacing == 0 {
		_, err := draw(str, box.X)
		return err
	}
	x := box.X
	for _, r := range str {
		width, err := draw(string(r), x)
		if err != nil {
			return err
		}
		x += width + float32(letterSpacing)
	}
	return nil
}

// feather is how far the edges of a shape fade out. SDL draws triangles without
// anti aliasing, so the shapes are drawn with a soft edge instead.
const feather = 0.5

func drawMesh(renderer *sdl.Renderer, mesh shapes.Mesh, c clay.Color) error {
	if len(mesh.Indices) == 0 {
		return nil
	}
	if err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		return err
	}
	vertices := make([]sdl.Vertex, len(mesh.Vertices))
	for i, v := range mesh.Vertices {
		vertices[i] = sdl.Vertex{
			Position: sdl.FPoint{X: v.X, Y: v.Y},
			Color:    sdl.Color{R: uint8(c.R), G: uint8(c.G), B: uint8(c.B), A: uint8(c.A * v.Alpha)},
		}
	}
	indices := make([]int32, len(mesh.Indices))
	for i, index := range mesh.Indices {
		indices[i] = int32(index)
	}
	return renderer.RenderGeometry(nil, vertices, indices)
}

func ClayRender(renderer *sdl.Renderer, renderCommands clay.RenderCommandArray, fonts []Font) error {
	var overlays overlay.Stack
	for _, renderCommand := range renderCommands {
		boundingBox := renderCommand.BoundingBox
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			color := overlays.Apply(config.BackgroundColor)
			if err := renderer.SetDrawColor(uint8(color.R), uint8(color.G), uint8(color.B), uint8(color.A)); err != nil {
				return err
			}
			rect := sdl.FRect{
				X: boundingBox.X,
				Y: boundingBox.Y,
				W: boundingBox.Width,
				H: boundingBox.Height,
			}
			if config.CornerRadius != (clay.CornerRadius{}) {
				if err := drawMesh(renderer, shapes.Fill(boundingBox, config.CornerRadius, feather), color); err != nil {
					return err
				}
			} else {
				if err := renderer.FillRectF(&rect); err != nil {
					return err
				}
			}
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			font, err := fonts[config.FontId].sized(config.FontSize)
			if err != nil {
				return err
			}
			textColor := overlays.Apply(config.TextColor)
			err = drawText(renderer, font, config.StringContents, boundingBox, int(config.LetterSpacing), sdl.Color{
				R: uint8(textColor.R),
				G: uint8(textColor.G),
				B: uint8(textColor.B),
				A: uint8(textColor.A),
			})
			if err != nil {
				return err
			}
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
			currentClippingRectangle := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if err := renderer.SetClipRect(&currentClippingRectangle); err != nil {
				return err
			}
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
			if err := renderer.SetClipRect(nil); err != nil {
				return err
			}
		case clay.RENDER_COMMAND_TYPE_OVERLAY_COLOR_START:
			overlays.Push(renderCommand.RenderData.OverlayColor.Color)
		case clay.RENDER_COMMAND_TYPE_OVERLAY_COLOR_END:
			overlays.Pop()
		case clay.RENDER_COMMAND_TYPE_IMAGE:
			config := &renderCommand.RenderData.Image
			texture, err := renderer.CreateTextureFromSurface(config.ImageData.(*sdl.Surface))
			if err != nil {
				return err
			}
			// The image is tinted by its background color, which is untinted when unset.
			if tint := config.BackgroundColor; tint != (clay.Color{}) {
				if err := texture.SetColorMod(uint8(tint.R), uint8(tint.G), uint8(tint.B)); err != nil {
					return err
				}
				if err := texture.SetAlphaMod(uint8(tint.A)); err != nil {
					return err
				}
			}
			destination := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if err := renderer.Copy(texture, nil, &destination); err != nil {
				return err
			}
			if err := texture.Destroy(); err != nil {
				return err
			}
			// Blend the image towards the overlays by drawing them over it.
			// This is exact for opaque images.
			for _, o := range overlays {
				if err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
					return err
				}
				if err := renderer.SetDrawColor(uint8(o.R), uint8(o.G), uint8(o.B), uint8(o.A)); err != nil {
					return err
				}
				if err := renderer.FillRect(&destination); err != nil {
					return err
				}
			}
		case clay.RENDER_COMMAND_TYPE_BORDER:
			config := &renderCommand.RenderData.Border
			config.Color = overlays.Apply(config.Color)
			if err := drawMesh(renderer, shapes.Border(boundingBox, config.CornerRadius, config.Width, feather), config.Color); err != nil {
				return err
			}
		case clay.RENDER_COMMAND_TYPE_NONE:
		case clay.RENDER_COMMAND_TYPE_CUSTOM:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}
	return nil
}
