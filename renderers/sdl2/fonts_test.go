// SPDX-License-Identifier: Zlib

package sdl2

import (
	"testing"

	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TestSizedFontsBelongToTheRenderer checks that the fonts opened at other sizes are kept
// by the renderer that opened them, and closed with it, rather than kept for the whole
// program with no way to free them.
func TestSizedFontsBelongToTheRenderer(t *testing.T) {
	if err := ttf.Init(); err != nil {
		t.Skipf("SDL_ttf is not available: %v", err)
	}
	defer ttf.Quit()
	stream, err := sdl.RWFromMem(fonts.RobotoRegularTTF)
	if err != nil {
		t.Fatal(err)
	}
	base, err := ttf.OpenFontRW(stream, 1, 16)
	if err != nil {
		t.Fatal(err)
	}
	defer base.Close()
	font := Font{Font: base, Data: fonts.RobotoRegularTTF}

	first := &RendererData{Fonts: []Font{font}}
	sized, err := first.sizedFont(&first.Fonts[0], 20)
	if err != nil {
		t.Fatal(err)
	}
	if again, _ := first.sizedFont(&first.Fonts[0], 20); again != sized {
		t.Error("the renderer opened a font a second time at a size it already has")
	}
	if unsized, _ := first.sizedFont(&first.Fonts[0], 0); unsized != base {
		t.Error("text without a font size did not use the font as is")
	}

	second := &RendererData{Fonts: []Font{font}}
	defer second.Close()
	if other, _ := second.sizedFont(&second.Fonts[0], 20); other == sized {
		t.Error("two renderers share a font")
	}

	first.Close()
	if len(first.sizedFonts) != 0 {
		t.Errorf("the renderer still has %d fonts after it was closed", len(first.sizedFonts))
	}
}
