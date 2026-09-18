// SPDX-License-Identifier: Zlib

package sdl3

import (
	"testing"

	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

// TestSizedFontsBelongToTheRenderer checks that the copies of a font at other sizes are
// kept by the renderer that made them, and closed with it. A *ttf.Font is SDL's pointer,
// which SDL can hand out again for a different font once one is closed, so copies kept
// for the whole program could outlive the font they were made from.
func TestSizedFontsBelongToTheRenderer(t *testing.T) {
	if err := ttf.Init(); err != nil {
		t.Skipf("SDL_ttf is not available: %v", err)
	}
	defer ttf.Quit()
	stream, err := sdl.IOFromConstMem(fonts.RobotoRegularTTF)
	if err != nil {
		t.Fatal(err)
	}
	font, err := ttf.OpenFontIO(stream, false, 16)
	if err != nil {
		t.Fatal(err)
	}
	defer font.Close()

	first := &RendererData{Fonts: []*ttf.Font{font}}
	sized, err := first.sizedFont(font, 20)
	if err != nil {
		t.Fatal(err)
	}
	if again, _ := first.sizedFont(font, 20); again != sized {
		t.Error("the renderer made a second copy of a font at a size it already has")
	}
	if unsized, _ := first.sizedFont(font, 0); unsized != font {
		t.Error("text without a font size did not use the font as is")
	}

	second := &RendererData{Fonts: []*ttf.Font{font}}
	defer second.Close()
	if other, _ := second.sizedFont(font, 20); other == sized {
		t.Error("two renderers share a copy of a font")
	}

	first.Close()
	if len(first.sizedFonts) != 0 {
		t.Errorf("the renderer still has %d copies after it was closed", len(first.sizedFonts))
	}
}
