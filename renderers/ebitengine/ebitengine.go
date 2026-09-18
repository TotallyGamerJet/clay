// SPDX-License-Identifier: Zlib

package ebitengine

import (
	"image"
	"image/color"
	"log/slog"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/renderers/internal/overlay"
	"github.com/TotallyGamerJet/clay/renderers/internal/shapes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	whiteImage      *ebiten.Image
	solidColorImage *ebiten.Image
)

func init() {
	// Creating a sub-image to avoid bleeding edges
	// https://github.com/hajimehoshi/ebiten/blob/1a4237213c92be1b9c16176887d992eb4183751b/vector/util.go#L26-L29
	img := ebiten.NewImage(3, 3)
	img.Fill(color.White)
	whiteImage = img.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	// Create a 1x1 solid color image for efficient rectangle drawing
	// This avoids the vector.DrawFilledRect bug on macOS/Retina
	solidColorImage = ebiten.NewImage(1, 1)
}

// RendererData is what the renderer draws with.
type RendererData struct {
	Fonts []text.Face

	// Fonts at the sizes text asks for, scaled for the display. They are kept here
	// rather than for the whole program, so that Close can let them go, along with the
	// fonts they were made from.
	sizedFaces map[sizedFaceKey]*text.GoTextFace
}

type sizedFaceKey struct {
	face *text.GoTextFace
	size float64
}

// Close lets go of the fonts made at other sizes.
func (r *RendererData) Close() {
	r.sizedFaces = nil
}

// sizedFace returns face with the font size of the text, scaled by scaleFactor.
// Only *text.GoTextFace can be resized; other faces and texts without a font size use face as is.
func (r *RendererData) sizedFace(face text.Face, fontSize uint16, scaleFactor float64) text.Face {
	f, ok := face.(*text.GoTextFace)
	if !ok || fontSize == 0 {
		return face
	}
	key := sizedFaceKey{face: f, size: float64(fontSize) * scaleFactor}
	sized, ok := r.sizedFaces[key]
	if !ok {
		c := *f
		c.Size = key.size
		sized = &c
		if r.sizedFaces == nil {
			r.sizedFaces = map[sizedFaceKey]*text.GoTextFace{}
		}
		r.sizedFaces[key] = sized
	}
	return sized
}

// MeasureText measures text for clay. userData must be the *RendererData the text is
// drawn with, which keeps the fonts at the sizes text asks for.
func MeasureText(txt string, config *clay.TextElementConfig, userData any) clay.Dimensions {
	r := userData.(*RendererData)

	scaleFactor := ebiten.Monitor().DeviceScaleFactor() // should we be passing the scaleFactor like we do in the renderer?
	font := r.sizedFace(r.Fonts[config.FontId], config.FontSize, scaleFactor)

	width, height := text.Measure(txt, font, font.Metrics().HLineGap)
	width += float64(config.LetterSpacing) * float64(max(len([]rune(txt))-1, 0)) * scaleFactor
	if config.LineHeight > 0 {
		height = float64(config.LineHeight) * scaleFactor
	}
	return clay.Dimensions{
		Width:  float32(width / scaleFactor),
		Height: float32(height / scaleFactor),
	}
}

// drawText draws one line of text, spacing the characters out by letterSpacing and
// centering them in the line box when the text sets a line height.
func drawText(screen *ebiten.Image, str string, face text.Face, box clay.BoundingBox, letterSpacing float32, colorScale ebiten.ColorScale) {
	metrics := face.Metrics()
	y := float64(box.Y)
	if height := metrics.HAscent + metrics.HDescent; float64(box.Height) > height {
		y += (float64(box.Height) - height) / 2
	}
	opts := &text.DrawOptions{}
	opts.ColorScale = colorScale
	if letterSpacing == 0 {
		opts.GeoM.Translate(float64(box.X), y)
		text.Draw(screen, str, face, opts)
		return
	}
	x := float64(box.X)
	for _, r := range str {
		opts.GeoM.Reset()
		opts.GeoM.Translate(x, y)
		text.Draw(screen, string(r), face, opts)
		x += text.Advance(string(r), face) + float64(letterSpacing)
	}
}

func ClayRender(screen *ebiten.Image, scaleFactor float32, renderCommands clay.RenderCommandArray, rendererData *RendererData) error {
	fullScreen := screen
	var overlays overlay.Stack
	for _, renderCommand := range renderCommands {
		boundingBox := renderCommand.BoundingBox
		boundingBox.X *= scaleFactor
		boundingBox.Y *= scaleFactor
		boundingBox.Width *= scaleFactor
		boundingBox.Height *= scaleFactor
		switch renderCommand.CommandType {
		case clay.RenderCommandTypeRectangle:
			config := &renderCommand.RenderData.Rectangle
			config.BackgroundColor = overlays.Apply(config.BackgroundColor)
			if config.CornerRadius != (clay.CornerRadius{}) {
				drawMesh(screen, shapes.Fill(boundingBox, scaleRadius(config.CornerRadius, scaleFactor), 0), config.BackgroundColor)
			} else {
				// Workaround for vector.DrawFilledRect bug on macOS/Retina displays
				rectColor := color.NRGBA{
					R: uint8(config.BackgroundColor.R),
					G: uint8(config.BackgroundColor.G),
					B: uint8(config.BackgroundColor.B),
					A: uint8(config.BackgroundColor.A),
				}
				solidColorImage.Fill(rectColor)
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Scale(float64(boundingBox.Width), float64(boundingBox.Height))
				opts.GeoM.Translate(float64(boundingBox.X), float64(boundingBox.Y))
				screen.DrawImage(solidColorImage, opts)
			}
		case clay.RenderCommandTypeText:
			config := &renderCommand.RenderData.Text
			config.TextColor = overlays.Apply(config.TextColor)
			font := rendererData.sizedFace(rendererData.Fonts[config.FontId], config.FontSize, float64(scaleFactor))

			var colorScale ebiten.ColorScale
			colorScale.Scale(
				config.TextColor.R/255,
				config.TextColor.G/255,
				config.TextColor.B/255,
				config.TextColor.A/255,
			)
			drawText(screen, config.StringContents, font, boundingBox, float32(config.LetterSpacing)*scaleFactor, colorScale)
		case clay.RenderCommandTypeScissorStart:
			screen = screen.SubImage(image.Rect(
				int(boundingBox.X), int(boundingBox.Y),
				int(boundingBox.X+boundingBox.Width),
				int(boundingBox.Y+boundingBox.Height),
			)).(*ebiten.Image)
		case clay.RenderCommandTypeScissorEnd:
			screen = fullScreen
		case clay.RenderCommandTypeOverlayColorStart:
			overlays.Push(renderCommand.RenderData.OverlayColor.Color)
		case clay.RenderCommandTypeOverlayColorEnd:
			overlays.Pop()
		case clay.RenderCommandTypeImage:
			config := &renderCommand.RenderData.Image
			img := config.ImageData.(*ebiten.Image)
			bounds := img.Bounds()
			opts := &colorm.DrawImageOptions{}
			opts.GeoM.Scale(float64(boundingBox.Width/float32(bounds.Dx())), float64(boundingBox.Height/float32(bounds.Dy())))
			opts.GeoM.Translate(float64(boundingBox.X), float64(boundingBox.Y))
			var cm colorm.ColorM
			// The image is tinted by its background color, which is untinted when unset.
			if tint := config.BackgroundColor; tint != (clay.Color{}) {
				cm.Scale(float64(tint.R/255), float64(tint.G/255), float64(tint.B/255), float64(tint.A/255))
			}
			for _, o := range overlays {
				// The color matrix works on non-premultiplied colors: rgb = rgb*(1-a) + overlay*a
				a := float64(o.A / 255)
				cm.Scale(1-a, 1-a, 1-a, 1)
				cm.Translate(float64(o.R/255)*a, float64(o.G/255)*a, float64(o.B/255)*a, 0)
			}
			colorm.DrawImage(screen, img, cm, opts)
		case clay.RenderCommandTypeBorder:
			config := &renderCommand.RenderData.Border
			config.Color = overlays.Apply(config.Color)
			width := clay.BorderWidth{
				Left:   uint16(float32(config.Width.Left) * scaleFactor),
				Right:  uint16(float32(config.Width.Right) * scaleFactor),
				Top:    uint16(float32(config.Width.Top) * scaleFactor),
				Bottom: uint16(float32(config.Width.Bottom) * scaleFactor),
			}
			drawMesh(screen, shapes.Border(boundingBox, scaleRadius(config.CornerRadius, scaleFactor), width, 0), config.Color)
		case clay.RenderCommandTypeNone:
		case clay.RenderCommandTypeCustom:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}

	return nil
}

func scaleRadius(radius clay.CornerRadius, scaleFactor float32) clay.CornerRadius {
	return clay.CornerRadius{
		TopLeft:     radius.TopLeft * scaleFactor,
		TopRight:    radius.TopRight * scaleFactor,
		BottomLeft:  radius.BottomLeft * scaleFactor,
		BottomRight: radius.BottomRight * scaleFactor,
	}
}

// drawMesh draws a shape in a single color. Ebitengine anti aliases the triangles itself,
// so the shapes are built without a feathered edge.
func drawMesh(screen *ebiten.Image, mesh shapes.Mesh, color clay.Color) {
	if len(mesh.Indices) == 0 {
		return
	}
	vertices := make([]ebiten.Vertex, len(mesh.Vertices))
	for i, v := range mesh.Vertices {
		vertices[i] = ebiten.Vertex{
			DstX: v.X, DstY: v.Y,
			SrcX: 1, SrcY: 1,
			ColorR: color.R / 255,
			ColorG: color.G / 255,
			ColorB: color.B / 255,
			ColorA: color.A / 255 * v.Alpha,
		}
	}
	screen.DrawTriangles(vertices, mesh.Indices, whiteImage, &ebiten.DrawTrianglesOptions{AntiAlias: true})
}
