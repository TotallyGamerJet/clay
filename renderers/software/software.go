// SPDX-License-Identifier: Zlib

package software

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log/slog"
	"math"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/renderers/internal/overlay"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// PrintPlaygroundImage prints the image so that it displays in the Go playground. It may have to shrink the image
// if it is too big for the playground.
func PrintPlaygroundImage(m image.Image) {
	const maxPixels = math.MaxUint16
	origWidth := m.Bounds().Dx()
	origHeight := m.Bounds().Dy()
	ratio := float64(origWidth) / float64(origHeight)
	x := math.Sqrt(float64(maxPixels) / ratio)
	newHeight := int(math.Floor(x))
	newWidth := int(math.Floor(ratio * x))
	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.BiLinear.Scale(newImg, newImg.Bounds(), m, m.Bounds(), draw.Over, nil)

	var buf bytes.Buffer
	err := png.Encode(&buf, newImg)
	if err != nil {
		panic(err)
	}
	fmt.Println("IMAGE:" + base64.StdEncoding.EncodeToString(buf.Bytes()))
}

// Font is a font that can be drawn at the font sizes requested by clay.
type Font struct {
	Font *opentype.Font
	// Options are used to create the faces of each size.
	// Options.Size is used when the text doesn't set a font size.
	Options opentype.FaceOptions

	faces map[uint16]font.Face
}

// Face returns the face of the given size.
func (f *Font) Face(size uint16) (font.Face, error) {
	if face, ok := f.faces[size]; ok {
		return face, nil
	}
	opts := f.Options
	if size != 0 {
		opts.Size = float64(size)
	}
	face, err := opentype.NewFace(f.Font, &opts)
	if err != nil {
		return nil, err
	}
	if f.faces == nil {
		f.faces = map[uint16]font.Face{}
	}
	f.faces[size] = face
	return face, nil
}

// MeasureText measures text for clay. userData must be a *[]*Font indexed by font id.
func MeasureText(txt string, config *clay.TextElementConfig, userData any) clay.Dimensions {
	fonts := *userData.(*[]*Font)
	face, err := fonts[config.FontId].Face(config.FontSize)
	if err != nil {
		panic(fmt.Errorf("software: failed to create font face: %w", err))
	}
	width := font.MeasureString(face, txt).Ceil() + int(config.LetterSpacing)*max(len([]rune(txt))-1, 0)
	height := face.Metrics().Height.Ceil()
	if config.LineHeight > 0 {
		height = int(config.LineHeight)
	}
	return clay.Dimensions{
		Width: float32(width), Height: float32(height),
	}
}

func toColor(c clay.Color) color.NRGBA {
	return color.NRGBA{R: uint8(c.R), G: uint8(c.G), B: uint8(c.B), A: uint8(c.A)}
}

func ClayRender(screen draw.Image, renderCommands clay.RenderCommandArray, fonts []*Font) error {
	fullScreen := screen
	var overlays overlay.Stack
	for _, renderCommand := range renderCommands {
		boundingBox := renderCommand.BoundingBox
		rect := image.Rect(int(boundingBox.X), int(boundingBox.Y), int(boundingBox.X+boundingBox.Width), int(boundingBox.Y+boundingBox.Height))
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			src := &image.Uniform{C: toColor(overlays.Apply(config.BackgroundColor))}
			if config.CornerRadius == (clay.CornerRadius{}) {
				draw.Draw(screen, rect, src, image.Point{}, draw.Over)
			} else {
				mask := newRoundedMask(boundingBox, config.CornerRadius, clay.BorderWidth{})
				draw.DrawMask(screen, mask.bounds, src, image.Point{}, mask, mask.bounds.Min, draw.Over)
			}
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			face, err := fonts[config.FontId].Face(config.FontSize)
			if err != nil {
				return err
			}
			// Center the line in its box, which is taller than the text when a line
			// height is set, and space the characters out by the letter spacing.
			metrics := face.Metrics()
			y := fixed.I(int(boundingBox.Y)) + metrics.Ascent
			if height := (metrics.Ascent + metrics.Descent).Ceil(); float32(height) < boundingBox.Height {
				y += fixed.I(int((boundingBox.Height - float32(height)) / 2))
			}
			d := &font.Drawer{
				Dst:  screen,
				Src:  image.NewUniform(toColor(overlays.Apply(config.TextColor))),
				Face: face,
				Dot:  fixed.Point26_6{X: fixed.I(int(boundingBox.X)), Y: y},
			}
			if config.LetterSpacing == 0 {
				d.DrawString(config.StringContents)
			} else {
				for _, r := range config.StringContents {
					d.DrawString(string(r))
					d.Dot.X += fixed.I(int(config.LetterSpacing))
				}
			}
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
			screen = fullScreen.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(rect).(draw.Image)
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
			screen = fullScreen
		case clay.RENDER_COMMAND_TYPE_OVERLAY_COLOR_START:
			overlays.Push(renderCommand.RenderData.OverlayColor.Color)
		case clay.RENDER_COMMAND_TYPE_OVERLAY_COLOR_END:
			overlays.Pop()
		case clay.RENDER_COMMAND_TYPE_IMAGE:
			config := &renderCommand.RenderData.Image
			img := config.ImageData.(*image.Image)
			if img == nil {
				continue
			}
			src := *img
			// The image is tinted by its background color, which is untinted when unset.
			if tint := config.BackgroundColor; tint != (clay.Color{}) {
				src = tintedImage{Image: src, tint: tint}
			}
			if overlays.Active() {
				src = overlayImage{Image: src, overlays: overlays}
			}
			draw.ApproxBiLinear.Scale(screen, rect, src, src.Bounds(), draw.Over, nil)
		case clay.RENDER_COMMAND_TYPE_BORDER:
			config := &renderCommand.RenderData.Border
			src := &image.Uniform{C: toColor(overlays.Apply(config.Color))}
			if config.CornerRadius == (clay.CornerRadius{}) {
				w := config.Width
				for _, r := range []image.Rectangle{
					image.Rect(rect.Min.X, rect.Min.Y, rect.Min.X+int(w.Left), rect.Max.Y),
					image.Rect(rect.Max.X-int(w.Right), rect.Min.Y, rect.Max.X, rect.Max.Y),
					image.Rect(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Min.Y+int(w.Top)),
					image.Rect(rect.Min.X, rect.Max.Y-int(w.Bottom), rect.Max.X, rect.Max.Y),
				} {
					draw.Draw(screen, r, src, image.Point{}, draw.Over)
				}
			} else {
				mask := newRoundedMask(boundingBox, config.CornerRadius, config.Width)
				draw.DrawMask(screen, mask.bounds, src, image.Point{}, mask, mask.bounds.Min, draw.Over)
			}
		case clay.RENDER_COMMAND_TYPE_NONE:
		case clay.RENDER_COMMAND_TYPE_CUSTOM:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}

	return nil
}

// roundedMask is an anti-aliased mask of a rounded rectangle, or of its border when width is set.
type roundedMask struct {
	bounds image.Rectangle
	outer  roundedBox
	inner  roundedBox
	border bool
}

type roundedBox struct {
	x0, y0, x1, y1 float32
	radius         clay.CornerRadius
}

func newRoundedMask(bb clay.BoundingBox, radius clay.CornerRadius, width clay.BorderWidth) *roundedMask {
	m := &roundedMask{
		bounds: image.Rect(
			int(math.Floor(float64(bb.X))), int(math.Floor(float64(bb.Y))),
			int(math.Ceil(float64(bb.X+bb.Width))), int(math.Ceil(float64(bb.Y+bb.Height))),
		),
		outer: newRoundedBox(bb.X, bb.Y, bb.X+bb.Width, bb.Y+bb.Height, radius),
	}
	if width != (clay.BorderWidth{}) {
		m.border = true
		l, r, t, b := float32(width.Left), float32(width.Right), float32(width.Top), float32(width.Bottom)
		shrink := func(radius, a, b float32) float32 { return max(radius-max(a, b), 0) }
		m.inner = newRoundedBox(bb.X+l, bb.Y+t, bb.X+bb.Width-r, bb.Y+bb.Height-b, clay.CornerRadius{
			TopLeft:     shrink(m.outer.radius.TopLeft, l, t),
			TopRight:    shrink(m.outer.radius.TopRight, r, t),
			BottomLeft:  shrink(m.outer.radius.BottomLeft, l, b),
			BottomRight: shrink(m.outer.radius.BottomRight, r, b),
		})
	}
	return m
}

func newRoundedBox(x0, y0, x1, y1 float32, radius clay.CornerRadius) roundedBox {
	maxRadius := max(min(x1-x0, y1-y0)/2, 0)
	radius.TopLeft = min(radius.TopLeft, maxRadius)
	radius.TopRight = min(radius.TopRight, maxRadius)
	radius.BottomLeft = min(radius.BottomLeft, maxRadius)
	radius.BottomRight = min(radius.BottomRight, maxRadius)
	return roundedBox{x0: x0, y0: y0, x1: x1, y1: y1, radius: radius}
}

// coverage returns how much of the pixel centered at x, y is inside the box.
func (b *roundedBox) coverage(x, y float32) float32 {
	if b.x1 <= b.x0 || b.y1 <= b.y0 {
		return 0
	}
	cx, cy := (b.x0+b.x1)/2, (b.y0+b.y1)/2
	hw, hh := (b.x1-b.x0)/2, (b.y1-b.y0)/2
	var r float32
	switch {
	case x < cx && y < cy:
		r = b.radius.TopLeft
	case x >= cx && y < cy:
		r = b.radius.TopRight
	case x < cx:
		r = b.radius.BottomLeft
	default:
		r = b.radius.BottomRight
	}
	// signed distance to a rounded box
	qx := float32(math.Abs(float64(x-cx))) - hw + r
	qy := float32(math.Abs(float64(y-cy))) - hh + r
	outside := float32(math.Hypot(float64(max(qx, 0)), float64(max(qy, 0))))
	dist := outside + min(max(qx, qy), 0) - r
	return min(max(0.5-dist, 0), 1)
}

func (m *roundedMask) ColorModel() color.Model { return color.AlphaModel }

func (m *roundedMask) Bounds() image.Rectangle { return m.bounds }

func (m *roundedMask) At(x, y int) color.Color {
	px, py := float32(x)+0.5, float32(y)+0.5
	c := m.outer.coverage(px, py)
	if m.border && c > 0 {
		c *= 1 - m.inner.coverage(px, py)
	}
	return color.Alpha{A: uint8(c*255 + 0.5)}
}

// overlayImage blends the colors of an image towards the active overlays.
type overlayImage struct {
	image.Image
	overlays overlay.Stack
}

func (o overlayImage) ColorModel() color.Model { return color.RGBA64Model }

func (o overlayImage) At(x, y int) color.Color {
	r, g, b, a := o.Image.At(x, y).RGBA()
	fr, fg, fb, fa := float32(r), float32(g), float32(b), float32(a)
	for _, c := range o.overlays {
		oa := c.A / 255
		// the colors are premultiplied, so the overlay is scaled by the alpha of the pixel
		fr += (c.R*257*fa/0xffff - fr) * oa
		fg += (c.G*257*fa/0xffff - fg) * oa
		fb += (c.B*257*fa/0xffff - fb) * oa
	}
	return color.RGBA64{R: uint16(fr), G: uint16(fg), B: uint16(fb), A: uint16(fa)}
}

// tintedImage multiplies the colors of an image by a tint.
type tintedImage struct {
	image.Image
	tint clay.Color
}

func (t tintedImage) ColorModel() color.Model { return color.RGBA64Model }

func (t tintedImage) At(x, y int) color.Color {
	r, g, b, a := t.Image.At(x, y).RGBA()
	return color.RGBA64{
		R: uint16(float32(r) * t.tint.R / 255),
		G: uint16(float32(g) * t.tint.G / 255),
		B: uint16(float32(b) * t.tint.B / 255),
		A: uint16(float32(a) * t.tint.A / 255),
	}
}
