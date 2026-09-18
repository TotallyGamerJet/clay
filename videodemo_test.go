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

// TestVideoDemoTransition hovers a sidebar button of the demo and checks that its
// highlight fades in over several frames, rather than appearing at once.
func TestVideoDemoTransition(t *testing.T) {
	arena := clay.CreateArenaWithCapacity(clay.MinMemorySize())
	defer arena.Free()
	clay.Initialize(arena, clay.Dimensions{Width: winWidth, Height: winHeight},
		clay.ErrorHandler{ErrorHandlerFunction: func(e clay.ErrorData) { t.Errorf("clay: %v", e) }})

	parsedFont, err := opentype.Parse(fonts.RobotoRegularTTF)
	if err != nil {
		t.Fatal(err)
	}
	faces := []*software.Font{
		videodemo.FontIdBody16: {
			Font:    parsedFont,
			Options: opentype.FaceOptions{Size: fontSize, DPI: 72, Hinting: font.HintingFull},
		},
	}
	clay.SetMeasureTextFunction(software.MeasureText, &faces)
	demo := videodemo.Initialize(nil)

	// The alpha the highlight of the button under the pointer is drawn with.
	// It is not drawn at all while fully transparent.
	const buttonX, buttonY = 100, 200
	highlightAlpha := func(commands clay.RenderCommandArray) float32 {
		for _, c := range commands {
			b := c.BoundingBox
			inButton := buttonX >= b.X && buttonX <= b.X+b.Width && buttonY >= b.Y && buttonY <= b.Y+b.Height
			if c.CommandType != clay.RENDER_COMMAND_TYPE_RECTANGLE || !inButton {
				continue
			}
			// The sidebar buttons are the grey rectangles; the panel behind them is darker.
			if color := c.RenderData.Rectangle.BackgroundColor; color.R == 120 {
				return color.A
			}
		}
		return 0
	}

	const frame = 1.0 / 60
	videodemo.CreateLayout(&demo, frame)

	// Clay calls hover handlers with the pointer state of the previous call, which starts
	// out as "pressed this frame", so settle it away from the buttons first. Otherwise the
	// first hover selects the document, which is not what is being tested here.
	clay.SetPointerState(clay.Vector2{X: 600, Y: 400}, false)
	videodemo.CreateLayout(&demo, frame)

	// Hover a sidebar button. The frame the transition starts on is still drawn with the
	// old color, so the highlight fades in over the frames after it.
	clay.SetPointerState(clay.Vector2{X: buttonX, Y: buttonY}, false)
	// The transition lasts 0.4 seconds, so run a little longer than that.
	var alphas []float32
	for range 30 {
		alphas = append(alphas, highlightAlpha(videodemo.CreateLayout(&demo, frame)))
	}
	if alphas[0] != 0 {
		t.Errorf("the highlight was already visible on the first frame: %v", alphas)
	}
	if alphas[1] <= 0 || alphas[1] >= 120 {
		t.Errorf("the highlight did not fade in, alphas are %v", alphas)
	}
	for i := 2; i < len(alphas); i++ {
		if alphas[i] < alphas[i-1] {
			t.Fatalf("the highlight did not fade in steadily: %v", alphas)
		}
	}
	if alphas[len(alphas)-1] != 120 {
		t.Errorf("the highlight ended at %v, want the full 120", alphas[len(alphas)-1])
	}

	// Moving away fades it back out.
	clay.SetPointerState(clay.Vector2{X: 600, Y: 400}, false)
	var out []float32
	for range 30 {
		out = append(out, highlightAlpha(videodemo.CreateLayout(&demo, frame)))
	}
	if out[1] >= 120 || out[1] <= 0 {
		t.Errorf("the highlight did not fade out, alphas are %v", out)
	}
	if last := out[len(out)-1]; last != 0 {
		t.Errorf("the highlight ended at %v, want it gone", last)
	}
}
