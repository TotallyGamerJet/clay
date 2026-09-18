// SPDX-License-Identifier: Zlib

package ebitengine

import (
	"bytes"
	"testing"

	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// TestSizedFacesBelongToTheRenderer checks that the fonts made at other sizes are kept by
// the renderer that made them, and let go with it, rather than kept for the whole program
// along with the fonts they were made from.
func TestSizedFacesBelongToTheRenderer(t *testing.T) {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.RobotoRegularTTF))
	if err != nil {
		t.Fatal(err)
	}
	face := &text.GoTextFace{Source: source, Size: 16}

	first := &RendererData{Fonts: []text.Face{face}}
	sized := first.sizedFace(face, 20, 2)
	if got := sized.(*text.GoTextFace).Size; got != 40 {
		t.Errorf("a 20 point font at a scale of 2 is %v points, want 40", got)
	}
	if again := first.sizedFace(face, 20, 2); again != sized {
		t.Error("the renderer made a second font at a size it already has")
	}
	if unsized := first.sizedFace(face, 0, 2); unsized != face {
		t.Error("text without a font size did not use the font as is")
	}

	second := &RendererData{Fonts: []text.Face{face}}
	if other := second.sizedFace(face, 20, 2); other == sized {
		t.Error("two renderers share a font")
	}

	first.Close()
	if len(first.sizedFaces) != 0 {
		t.Errorf("the renderer still has %d fonts after it was closed", len(first.sizedFaces))
	}
}
