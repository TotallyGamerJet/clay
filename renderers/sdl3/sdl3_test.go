// SPDX-License-Identifier: Zlib

package sdl3_test

import (
	"image"
	"os"
	"testing"

	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/internal/testlayout"
	"github.com/TotallyGamerJet/clay/renderers/sdl3"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

// TestMain loads SDL once for the whole test binary. The bindings can't unload and load
// the libraries again in the same process: text stops being drawn after they are, so a
// test that loaded them itself would fail the second time it ran, as with -count=2.
func TestMain(m *testing.M) {
	sdlLib := binsdl.Load()
	ttfLib := binttf.Load()
	code := m.Run()
	ttfLib.Unload()
	sdlLib.Unload()
	os.Exit(code)
}

// TestLayout draws the shared test layout with SDL's software renderer, which needs no
// window, and checks it looks the way it should.
func TestLayout(t *testing.T) {
	if err := ttf.Init(); err != nil {
		t.Skipf("SDL_ttf is not available: %v", err)
	}
	defer ttf.Quit()

	target, err := sdl.CreateSurface(testlayout.Width, testlayout.Height, sdl.PIXELFORMAT_RGBA32)
	if err != nil {
		t.Skipf("SDL can't create a surface: %v", err)
	}
	defer target.Destroy()
	renderer, err := target.CreateSoftwareRenderer()
	if err != nil {
		t.Skipf("SDL can't create a software renderer: %v", err)
	}
	engine, err := ttf.CreateRendererTextEngine(renderer)
	if err != nil {
		t.Fatal(err)
	}
	stream, err := sdl.IOFromConstMem(fonts.RobotoRegularTTF)
	if err != nil {
		t.Fatal(err)
	}
	font, err := ttf.OpenFontIO(stream, false, 16)
	if err != nil {
		t.Fatal(err)
	}
	defer font.Close()

	data := &sdl3.RendererData{Renderer: renderer, TextEngine: engine, Fonts: []*ttf.Font{font}}
	testlayout.Init(sdl3.MeasureText, &data.Fonts)
	if err := sdl3.ClayRender(data, testlayout.Build()); err != nil {
		t.Fatal(err)
	}
	if err := renderer.Present(); err != nil {
		t.Fatal(err)
	}
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
