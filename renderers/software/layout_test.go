// SPDX-License-Identifier: Zlib

package software_test

import (
	"image"
	"testing"

	"github.com/TotallyGamerJet/clay/internal/testlayout"
	"github.com/TotallyGamerJet/clay/renderers/software"
)

// TestLayout draws the shared test layout and checks it looks the way it should.
func TestLayout(t *testing.T) {
	fonts := testFonts(t)
	testlayout.Init(software.MeasureText, &fonts)

	screen := image.NewRGBA(image.Rect(0, 0, testlayout.Width, testlayout.Height))
	if err := software.ClayRender(screen, testlayout.Build(), fonts); err != nil {
		t.Fatal(err)
	}
	testlayout.Check(t, screen)
}
