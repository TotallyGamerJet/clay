// SPDX-License-Identifier: Zlib

package sdl2_test

import (
	"image"
	"testing"

	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/internal/testlayout"
	"github.com/TotallyGamerJet/clay/renderers/sdl2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TestLayout draws the shared test layout with SDL's software renderer, which needs no
// window, and checks it looks the way it should.
func TestLayout(t *testing.T) {
	if err := ttf.Init(); err != nil {
		t.Skipf("SDL_ttf is not available: %v", err)
	}
	defer ttf.Quit()

	stream, err := sdl.RWFromMem(fonts.RobotoRegularTTF)
	if err != nil {
		t.Fatal(err)
	}
	font, err := ttf.OpenFontRW(stream, 1, 16)
	if err != nil {
		t.Fatal(err)
	}
	defer font.Close()
	fontList := []sdl2.Font{{Font: font, Data: fonts.RobotoRegularTTF}}
	testlayout.Init(sdl2.MeasureText, &fontList)

	target, err := sdl.CreateRGBSurfaceWithFormat(0, testlayout.Width, testlayout.Height, 32, uint32(sdl.PIXELFORMAT_RGBA32))
	if err != nil {
		t.Skipf("SDL can't create a surface: %v", err)
	}
	defer target.Free()
	renderer, err := sdl.CreateSoftwareRenderer(target)
	if err != nil {
		t.Skipf("SDL can't create a software renderer: %v", err)
	}
	defer renderer.Destroy()

	if err := sdl2.ClayRender(renderer, testlayout.Build(), fontList); err != nil {
		t.Fatal(err)
	}
	renderer.Present()
	testlayout.Check(t, surfaceImage(target))
}

// surfaceImage copies a surface into an image, dropping the alpha, which is what was
// drawn onto rather than a color.
func surfaceImage(surface *sdl.Surface) *image.RGBA {
	out := image.NewRGBA(image.Rect(0, 0, int(surface.W), int(surface.H)))
	pixels := surface.Pixels()
	for y := range int(surface.H) {
		copy(out.Pix[y*out.Stride:(y+1)*out.Stride], pixels[y*int(surface.Pitch):])
	}
	for i := 3; i < len(out.Pix); i += 4 {
		out.Pix[i] = 255
	}
	return out
}
