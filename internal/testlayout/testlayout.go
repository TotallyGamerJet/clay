// SPDX-License-Identifier: Zlib

// Package testlayout is a layout that uses the features every renderer is expected to
// draw, along with checks that a rendered image of it looks right. Each renderer runs it
// headlessly in its own test, so that they are all held to the same behavior.
//
// The layout is built out of fixed sizes, so that the positions the checks look at are
// the same whichever font a renderer measures the text with.
//
// The software renderer and both SDL renderers draw without a window, so their tests run
// anywhere. The ebitengine and raylib ones need a GPU, and a window that has to be made
// on the main thread, so they draw into a hidden window from TestMain, and only run when
// CLAY_TEST_GPU is set.
package testlayout

import (
	"image"
	"testing"

	"github.com/TotallyGamerJet/clay"
)

// The size of the layout, and the rows it is made of.
const (
	Width, Height = 400, 300

	shapesRow  = 100 // three 100x100 boxes
	spacingRow = 60  // the same text with and without letter spacing
	lineRow    = 60  // text in a line box taller than the text
)

var (
	background = clay.Color{R: 0, G: 0, B: 40, A: 255}
	fill       = clay.Color{R: 255, G: 0, B: 0, A: 255}
	white      = clay.Color{R: 255, G: 255, B: 255, A: 255}
	overlay    = clay.Color{R: 255, G: 0, B: 0, A: 128}
)

func fixed(w, h float32) clay.Sizing {
	return clay.Sizing{Width: clay.SizingFixed(w), Height: clay.SizingFixed(h)}
}

// Init sets up a context for the layout, with the renderer's own text measuring function.
func Init(measureText func(text string, config *clay.TextElementConfig, userData any) clay.Dimensions, userData any) {
	arena := clay.CreateArenaWithCapacity(clay.MinMemorySize())
	clay.Initialize(arena, clay.Dimensions{Width: Width, Height: Height},
		clay.ErrorHandler{ErrorHandlerFunction: func(e clay.ErrorData) { panic(e) }})
	clay.SetMeasureTextFunction(measureText, userData)
}

// Build declares the layout and returns the commands to draw it.
func Build() clay.RenderCommandArray {
	clay.BeginLayout()
	clay.UI(clay.ID("Root"))(clay.ElementDeclaration{
		Layout:          clay.LayoutConfig{Sizing: fixed(Width, Height), LayoutDirection: clay.TopToBottom},
		BackgroundColor: background,
	}, func() {
		clay.UI(clay.ID("Shapes"))(clay.ElementDeclaration{Layout: clay.LayoutConfig{Sizing: fixed(Width, shapesRow)}}, func() {
			// Only the top left corner is rounded, so the other three stay square.
			clay.UI(clay.ID("Corners"))(clay.ElementDeclaration{
				Layout:          clay.LayoutConfig{Sizing: fixed(100, 100)},
				BackgroundColor: fill,
				CornerRadius:    clay.CornerRadius{TopLeft: 40},
			}, nil)
			// A border with nothing inside it.
			clay.UI(clay.ID("Border"))(clay.ElementDeclaration{
				Layout:       clay.LayoutConfig{Sizing: fixed(100, 100)},
				CornerRadius: clay.CornerRadiusAll(20),
				Border:       clay.BorderElementConfig{Color: white, Width: clay.BorderOutside(10)},
			}, nil)
			// A white box under an overlay, which tints everything inside it.
			clay.UI(clay.ID("Overlay"))(clay.ElementDeclaration{
				Layout:       clay.LayoutConfig{Sizing: fixed(100, 100)},
				OverlayColor: overlay,
			}, func() {
				clay.UI(clay.ID("Overlaid"))(clay.ElementDeclaration{
					Layout:          clay.LayoutConfig{Sizing: fixed(100, 100)},
					BackgroundColor: white,
				}, nil)
			})
		})
		clay.UI(clay.ID("Spacing"))(clay.ElementDeclaration{Layout: clay.LayoutConfig{Sizing: fixed(Width, spacingRow)}}, func() {
			clay.UI(clay.ID("Plain"))(clay.ElementDeclaration{Layout: clay.LayoutConfig{Sizing: fixed(Width/2, spacingRow)}}, func() {
				clay.Text("iiiiii", &clay.TextElementConfig{FontSize: 20, TextColor: white})
			})
			clay.UI(clay.ID("Spaced"))(clay.ElementDeclaration{Layout: clay.LayoutConfig{Sizing: fixed(Width/2, spacingRow)}}, func() {
				clay.Text("iiiiii", &clay.TextElementConfig{FontSize: 20, LetterSpacing: 12, TextColor: white})
			})
		})
		// Text in a line box taller than the text is, which centers it in the box.
		clay.UI(clay.ID("LineHeight"))(clay.ElementDeclaration{Layout: clay.LayoutConfig{Sizing: fixed(Width, lineRow)}}, func() {
			clay.Text("HEIGHT", &clay.TextElementConfig{FontSize: 20, LineHeight: lineRow, TextColor: white})
		})
	})
	return clay.EndLayout(0)
}

// Check reports whether a rendered image of the layout looks the way it should.
// Renderers differ in how they anti alias and lay out glyphs, so it checks what the
// features mean rather than exact pixels.
func Check(t *testing.T, screen image.Image) {
	t.Helper()
	if b := screen.Bounds(); b.Dx() < Width || b.Dy() < Height {
		t.Fatalf("the image is %dx%d, want at least %dx%d", b.Dx(), b.Dy(), Width, Height)
	}
	// Renderers may draw at a larger scale, so work in fractions of the layout.
	scaleX := float64(screen.Bounds().Dx()) / Width
	scaleY := float64(screen.Bounds().Dy()) / Height
	at := func(x, y float64) (r, g, b float64) {
		c := screen.At(screen.Bounds().Min.X+int(x*scaleX), screen.Bounds().Min.Y+int(y*scaleY))
		ri, gi, bi, _ := c.RGBA()
		return float64(ri) / 0xffff, float64(gi) / 0xffff, float64(bi) / 0xffff
	}
	lit := func(x, y float64) bool {
		r, g, b := at(x, y)
		return r+g+b > 0.4 // brighter than the background
	}

	t.Run("corner radius", func(t *testing.T) {
		// The rounded corner is cut away, the three square ones are not.
		if r, _, _ := at(3, 3); r > 0.5 {
			t.Error("the rounded top left corner was filled in")
		}
		for _, corner := range [][2]float64{{96, 3}, {96, 96}, {3, 96}} {
			if r, _, _ := at(corner[0], corner[1]); r < 0.5 {
				t.Errorf("the square corner at %v was rounded off", corner)
			}
		}
	})

	t.Run("border", func(t *testing.T) {
		if !lit(150, 3) {
			t.Error("the top of the border is missing")
		}
		if !lit(103, 50) {
			t.Error("the left of the border is missing")
		}
		if lit(150, 50) {
			t.Error("the inside of the border was filled in")
		}
		// The rounded corner is drawn, and joins the sides without a gap. Its arc runs
		// between 10 and 20 pixels from the center of the corner, which is at 120, 20.
		if !lit(106, 6) {
			t.Error("the rounded corner of the border is missing, or has a gap where it meets a side")
		}
	})

	t.Run("overlay", func(t *testing.T) {
		r, g, b := at(250, 50)
		if r < 0.6 || g > 0.8 || b > 0.8 {
			t.Errorf("the white box under a red overlay is %.2f %.2f %.2f, want it tinted red", r, g, b)
		}
	})

	t.Run("letter spacing", func(t *testing.T) {
		rightmost := func(x0, x1 float64) float64 {
			var found float64
			for x := x0; x < x1; x++ {
				for y := float64(shapesRow); y < shapesRow+spacingRow; y++ {
					if lit(x, y) {
						found = x
						break
					}
				}
			}
			return found
		}
		plain := rightmost(0, Width/2)
		spaced := rightmost(Width/2, Width) - Width/2
		if plain == 0 || spaced == 0 {
			t.Fatalf("no text was drawn: plain ends at %v, spaced at %v", plain, spaced)
		}
		if spaced < plain+40 {
			t.Errorf("spaced text ends at %v and plain text at %v, want the spaced text much wider", spaced, plain)
		}
	})

	t.Run("line height", func(t *testing.T) {
		top, bottom := -1.0, -1.0
		for y := float64(shapesRow + spacingRow); y < shapesRow+spacingRow+lineRow; y++ {
			for x := float64(0); x < Width; x++ {
				if lit(x, y) {
					if top < 0 {
						top = y
					}
					bottom = y
					break
				}
			}
		}
		if top < 0 {
			t.Fatal("no text was drawn in the line box")
		}
		// The text sits in the middle of the box, not at the top of it.
		center := (top + bottom) / 2
		want := float64(shapesRow + spacingRow + lineRow/2)
		if center < want-12 || center > want+12 {
			t.Errorf("the text is centered at %v in its line box, want it near %v", center, want)
		}
	})
}
