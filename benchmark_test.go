// SPDX-License-Identifier: Zlib

package clay_test

import (
	"testing"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/examples/videodemo"
	"github.com/TotallyGamerJet/clay/renderers/software"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// BenchmarkVideoDemo lays out one frame of the demo, which is the work an app does
// every frame: declaring elements, measuring text, and reading the render commands back.
func BenchmarkVideoDemo(b *testing.B) {
	arena := clay.CreateArenaWithCapacity(clay.MinMemorySize())
	defer arena.Free()
	clay.Initialize(arena, clay.Dimensions{Width: 640, Height: 480}, clay.ErrorHandler{})
	parsed, _ := opentype.Parse(fonts.RobotoRegularTTF)
	faces := []*software.Font{{Font: parsed, Options: opentype.FaceOptions{Size: 16, DPI: 72, Hinting: font.HintingFull}}}
	clay.SetMeasureTextFunction(software.MeasureText, &faces)
	demo := videodemo.Initialize(nil)
	b.ReportAllocs()
	for b.Loop() {
		videodemo.CreateLayout(&demo, 1.0/60)
	}
}
