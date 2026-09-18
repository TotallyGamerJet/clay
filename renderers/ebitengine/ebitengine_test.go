// SPDX-License-Identifier: Zlib

package ebitengine_test

import (
	"bytes"
	"errors"
	"image"
	"log"
	"os"
	"testing"

	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/internal/testlayout"
	"github.com/TotallyGamerJet/clay/renderers/ebitengine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Ebitengine draws with a GPU, through a window that has to be created on the main
// thread, which tests don't run on. So the tests run inside a hidden game: TestMain
// hands the main thread to ebitengine and runs the tests in a goroutine, which draw by
// handing work back to the game loop.
//
// A GPU isn't always there to draw with, so the tests only run when CLAY_TEST_GPU is set.

var (
	gpu  bool
	draw = make(chan func(), 1)
	done = make(chan struct{})
)

type game struct{}

func (game) Update() error {
	select {
	case <-done:
		return ebiten.Termination
	default:
		return nil
	}
}

func (game) Draw(*ebiten.Image) {
	select {
	case job := <-draw:
		job()
	default:
	}
}

func (game) Layout(int, int) (int, int) { return testlayout.Width, testlayout.Height }

func TestMain(m *testing.M) {
	gpu = os.Getenv("CLAY_TEST_GPU") != ""
	if !gpu {
		os.Exit(m.Run())
	}
	var code int
	go func() {
		defer close(done)
		code = m.Run()
	}()
	ebiten.SetWindowVisible(false)
	ebiten.SetWindowSize(testlayout.Width, testlayout.Height)
	if err := ebiten.RunGame(game{}); err != nil && !errors.Is(err, ebiten.Termination) {
		log.Fatal(err)
	}
	os.Exit(code)
}

// render runs a draw in the game loop and returns what it drew.
func render(t *testing.T, size image.Point, drawTo func(screen *ebiten.Image)) image.Image {
	t.Helper()
	out := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	finished := make(chan struct{})
	draw <- func() {
		defer close(finished)
		// Images can only be created and drawn to from within the game loop.
		screen := ebiten.NewImage(size.X, size.Y)
		defer screen.Deallocate()
		drawTo(screen)
		screen.ReadPixels(out.Pix)
	}
	<-finished
	return out
}

// TestLayout draws the shared test layout and checks it looks the way it should.
func TestLayout(t *testing.T) {
	if !gpu {
		t.Skip("set CLAY_TEST_GPU to run the tests that need a GPU")
	}
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.RobotoRegularTTF))
	if err != nil {
		t.Fatal(err)
	}
	data := &ebitengine.RendererData{Fonts: []text.Face{&text.GoTextFace{Source: source, Size: 16}}}
	defer data.Close()
	testlayout.Init(ebitengine.MeasureText, data)
	commands := testlayout.Build()

	screen := render(t, image.Pt(testlayout.Width, testlayout.Height), func(screen *ebiten.Image) {
		if err := ebitengine.ClayRender(screen, 1, commands, data); err != nil {
			t.Error(err)
		}
	})
	testlayout.Check(t, screen)
}
