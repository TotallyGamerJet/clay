// SPDX-License-Identifier: Zlib

package software_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/renderers/software"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// The renderers are given render commands, which these tests write by hand so that each
// part of a command can be checked on its own.

func testFonts(t *testing.T) []*software.Font {
	t.Helper()
	parsed, err := opentype.Parse(fonts.RobotoRegularTTF)
	if err != nil {
		t.Fatal(err)
	}
	return []*software.Font{{
		Font:    parsed,
		Options: opentype.FaceOptions{Size: 16, DPI: 72, Hinting: font.HintingFull},
	}}
}

func render(t *testing.T, commands clay.RenderCommandArray) *image.RGBA {
	t.Helper()
	screen := image.NewRGBA(image.Rect(0, 0, 200, 100))
	if err := software.ClayRender(screen, commands, testFonts(t)); err != nil {
		t.Fatal(err)
	}
	return screen
}

func TestImageTint(t *testing.T) {
	source := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for i := range source.Pix {
		source.Pix[i] = 255 // opaque white
	}
	var img image.Image = source

	command := clay.RenderCommand{
		CommandType: clay.RenderCommandTypeImage,
		BoundingBox: clay.BoundingBox{X: 0, Y: 0, Width: 20, Height: 20},
	}
	command.RenderData.Image.ImageData = &img

	untinted := render(t, clay.RenderCommandArray{command}).RGBAAt(10, 10)
	if untinted != (color.RGBA{R: 255, G: 255, B: 255, A: 255}) {
		t.Errorf("without a tint the image was drawn as %v, want white", untinted)
	}

	// The background color of an image tints it.
	command.RenderData.Image.BackgroundColor = clay.Color{R: 255, G: 0, B: 0, A: 255}
	tinted := render(t, clay.RenderCommandArray{command}).RGBAAt(10, 10)
	if tinted != (color.RGBA{R: 255, A: 255}) {
		t.Errorf("with a red tint the image was drawn as %v, want red", tinted)
	}
}

func TestTextLetterSpacing(t *testing.T) {
	command := clay.RenderCommand{
		CommandType: clay.RenderCommandTypeText,
		BoundingBox: clay.BoundingBox{X: 0, Y: 0, Width: 200, Height: 20},
	}
	command.RenderData.Text.StringContents = "iiii"
	command.RenderData.Text.TextColor = clay.Color{R: 255, G: 255, B: 255, A: 255}

	width := func(screen *image.RGBA) int {
		rightmost := 0
		for y := range 100 {
			for x := range 200 {
				if screen.RGBAAt(x, y).R > 0 {
					rightmost = max(rightmost, x)
				}
			}
		}
		return rightmost
	}

	plain := width(render(t, clay.RenderCommandArray{command}))
	command.RenderData.Text.LetterSpacing = 10
	spaced := width(render(t, clay.RenderCommandArray{command}))
	if spaced <= plain+20 {
		t.Errorf("text is %d wide with letter spacing and %d without, want it much wider", spaced, plain)
	}
}

func TestTextLineHeight(t *testing.T) {
	command := clay.RenderCommand{
		CommandType: clay.RenderCommandTypeText,
		BoundingBox: clay.BoundingBox{X: 0, Y: 0, Width: 200, Height: 20},
	}
	command.RenderData.Text.StringContents = "Ay"
	command.RenderData.Text.TextColor = clay.Color{R: 255, G: 255, B: 255, A: 255}

	center := func(screen *image.RGBA) float32 {
		first, last := -1, -1
		for y := range 100 {
			for x := range 200 {
				if screen.RGBAAt(x, y).R > 0 {
					if first < 0 {
						first = y
					}
					last = y
					break
				}
			}
		}
		if first < 0 {
			t.Fatal("no text was drawn")
		}
		return float32(first+last) / 2
	}

	tight := center(render(t, clay.RenderCommandArray{command}))

	// A taller line box centers the text in it.
	command.BoundingBox.Height = 60
	tall := center(render(t, clay.RenderCommandArray{command}))
	if tall <= tight+10 {
		t.Errorf("text sits at %v in a 60 pixel line and %v in a 20 pixel one, want it centered", tall, tight)
	}
}

func TestRectangleCornerRadius(t *testing.T) {
	command := clay.RenderCommand{
		CommandType: clay.RenderCommandTypeRectangle,
		BoundingBox: clay.BoundingBox{X: 0, Y: 0, Width: 100, Height: 100},
	}
	command.RenderData.Rectangle.BackgroundColor = clay.Color{R: 255, G: 255, B: 255, A: 255}
	// Only the top left corner is rounded.
	command.RenderData.Rectangle.CornerRadius = clay.CornerRadius{TopLeft: 30}

	screen := render(t, clay.RenderCommandArray{command})
	if got := screen.RGBAAt(2, 2).R; got != 0 {
		t.Errorf("the rounded corner was filled in: %v", got)
	}
	for _, corner := range []image.Point{{X: 97, Y: 2}, {X: 97, Y: 97}, {X: 2, Y: 97}} {
		if got := screen.RGBAAt(corner.X, corner.Y).R; got != 255 {
			t.Errorf("the square corner at %v was rounded off: %v", corner, got)
		}
	}
}
