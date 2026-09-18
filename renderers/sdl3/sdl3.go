// SPDX-License-Identifier: Zlib

package sdl3

import (
	"fmt"
	"log/slog"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/renderers/internal/overlay"
	"github.com/TotallyGamerJet/clay/renderers/internal/shapes"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

type RendererData struct {
	Renderer   *sdl.Renderer
	TextEngine *ttf.TextEngine
	Fonts      []*ttf.Font
}

type sizedFontKey struct {
	font *ttf.Font
	size uint16
}

var sizedFonts = map[sizedFontKey]*ttf.Font{}

// sizedFont returns a copy of font with the font size of the text.
// Copies are kept instead of resizing font, as resizing clears its glyph cache.
func sizedFont(font *ttf.Font, size uint16) (*ttf.Font, error) {
	if size == 0 {
		return font, nil
	}
	key := sizedFontKey{font: font, size: size}
	if sized, ok := sizedFonts[key]; ok {
		return sized, nil
	}
	sized, err := font.Copy()
	if err != nil {
		return nil, err
	}
	if err := sized.SetSize(float32(size)); err != nil {
		sized.Close()
		return nil, err
	}
	sizedFonts[key] = sized
	return sized, nil
}

func MeasureText(text string, config *clay.TextElementConfig, userData any) clay.Dimensions {
	fonts := *userData.(*[]*ttf.Font)
	font, err := sizedFont(fonts[config.FontId], config.FontSize)
	if err != nil {
		panic(fmt.Errorf("sdl3: failed to size font: %w", err))
	}

	width, height, err := font.StringSize(text)
	if err != nil {
		panic(fmt.Errorf("sdl3: failed to measure text: %w", err))
	}
	width += int32(config.LetterSpacing) * int32(max(len([]rune(text))-1, 0))
	if config.LineHeight > 0 {
		height = int32(config.LineHeight)
	}

	return clay.Dimensions{
		Width:  float32(width),
		Height: float32(height),
	}
}

// drawText draws one line of text, spacing the characters out by letterSpacing and
// centering them in the line box when the text sets a line height.
func drawText(engine *ttf.TextEngine, font *ttf.Font, str string, box clay.BoundingBox, letterSpacing float32, color clay.Color) error {
	draw := func(str string, x float32) (float32, error) {
		text, err := engine.CreateText(font, str)
		if err != nil {
			return 0, err
		}
		defer text.Destroy()
		if err := text.SetColor(uint8(color.R), uint8(color.G), uint8(color.B), uint8(color.A)); err != nil {
			return 0, err
		}
		width, height, err := font.StringSize(str)
		if err != nil {
			return 0, err
		}
		y := box.Y
		if float32(height) < box.Height {
			y += (box.Height - float32(height)) / 2
		}
		return float32(width), text.DrawRenderer(x, y)
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
		x += width + letterSpacing
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
			Color:    sdl.FColor{R: c.R / 255, G: c.G / 255, B: c.B / 255, A: c.A / 255 * v.Alpha},
		}
	}
	indices := make([]int32, len(mesh.Indices))
	for i, index := range mesh.Indices {
		indices[i] = int32(index)
	}
	return renderer.RenderGeometry(nil, vertices, indices)
}

func ClayRender(rendererData *RendererData, renderCommands clay.RenderCommandArray) error {
	renderer := rendererData.Renderer
	fonts := rendererData.Fonts
	textEngine := rendererData.TextEngine
	var overlays overlay.Stack
	for _, renderCommand := range renderCommands {
		boundingBox := renderCommand.BoundingBox
		rect := sdl.FRect{
			X: boundingBox.X,
			Y: boundingBox.Y,
			W: boundingBox.Width,
			H: boundingBox.Height,
		}
		switch renderCommand.CommandType {
		case clay.RenderCommandTypeRectangle:
			config := &renderCommand.RenderData.Rectangle
			config.BackgroundColor = overlays.Apply(config.BackgroundColor)
			renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
			renderer.SetDrawColor(
				uint8(config.BackgroundColor.R),
				uint8(config.BackgroundColor.G),
				uint8(config.BackgroundColor.B),
				uint8(config.BackgroundColor.A),
			)
			if config.CornerRadius != (clay.CornerRadius{}) {
				if err := drawMesh(renderer, shapes.Fill(boundingBox, config.CornerRadius, feather), config.BackgroundColor); err != nil {
					return err
				}
			} else {
				renderer.RenderFillRect(&rect)
			}
		case clay.RenderCommandTypeText:
			config := &renderCommand.RenderData.Text
			config.TextColor = overlays.Apply(config.TextColor)
			font, err := sizedFont(fonts[config.FontId], config.FontSize)
			if err != nil {
				return err
			}
			if err := drawText(textEngine, font, config.StringContents, boundingBox, float32(config.LetterSpacing), config.TextColor); err != nil {
				return err
			}
		case clay.RenderCommandTypeScissorStart:
			currentClippingRectangle := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if err := renderer.SetClipRect(&currentClippingRectangle); err != nil {
				return err
			}
		case clay.RenderCommandTypeScissorEnd:
			if err := renderer.SetClipRect(nil); err != nil {
				return err
			}
		case clay.RenderCommandTypeOverlayColorStart:
			overlays.Push(renderCommand.RenderData.OverlayColor.Color)
		case clay.RenderCommandTypeOverlayColorEnd:
			overlays.Pop()
		case clay.RenderCommandTypeImage:
			config := &renderCommand.RenderData.Image
			texture, err := renderer.CreateTextureFromSurface(config.ImageData.(*sdl.Surface))
			if err != nil {
				return err
			}
			destination := sdl.FRect{
				X: rect.X,
				Y: rect.Y,
				W: rect.W,
				H: rect.H,
			}
			// The image is tinted by its background color, which is untinted when unset.
			if tint := config.BackgroundColor; tint != (clay.Color{}) {
				if err := texture.SetColorModFloat(tint.R/255, tint.G/255, tint.B/255); err != nil {
					return err
				}
				if err := texture.SetAlphaModFloat(tint.A / 255); err != nil {
					return err
				}
			}
			if err := renderer.RenderTexture(texture, nil, &destination); err != nil {
				return err
			}
			texture.Destroy()
			// Blend the image towards the overlays by drawing them over it.
			// This is exact for opaque images.
			for _, o := range overlays {
				renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
				renderer.SetDrawColor(uint8(o.R), uint8(o.G), uint8(o.B), uint8(o.A))
				if err := renderer.RenderFillRect(&destination); err != nil {
					return err
				}
			}
		case clay.RenderCommandTypeBorder:
			config := &renderCommand.RenderData.Border
			config.Color = overlays.Apply(config.Color)
			if err := drawMesh(renderer, shapes.Border(boundingBox, config.CornerRadius, config.Width, feather), config.Color); err != nil {
				return err
			}
		case clay.RenderCommandTypeNone:
		case clay.RenderCommandTypeCustom:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}
	return nil
}
