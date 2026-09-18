// SPDX-License-Identifier: Zlib

package raylib_test

import (
	"image"
	"os"
	"runtime"
	"testing"

	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/internal/testlayout"
	clayraylib "github.com/TotallyGamerJet/clay/renderers/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Raylib draws with a GPU, through a window that has to be created on the main thread,
// which tests don't run on. So TestMain keeps the main thread, and the tests running in
// a goroutine hand their drawing back to it.
//
// A GPU isn't always there to draw with, so the tests only run when CLAY_TEST_GPU is set.

var (
	gpu  bool
	draw = make(chan func(), 1)
)

func TestMain(m *testing.M) {
	gpu = os.Getenv("CLAY_TEST_GPU") != ""
	if gpu {
		runtime.LockOSThread()
		rl.SetConfigFlags(rl.FlagWindowHidden)
		rl.SetTraceLogLevel(rl.LogWarning)
		rl.InitWindow(testlayout.Width, testlayout.Height, "clay test")
		if !rl.IsWindowReady() {
			gpu = false
		}
	}
	if !gpu {
		os.Exit(m.Run())
	}

	var code int
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		code = m.Run()
	}()
	for {
		select {
		case job := <-draw:
			job()
		case <-finished:
			rl.CloseWindow()
			os.Exit(code)
		}
	}
}

// onMain runs a function on the main thread, which everything that touches the GPU,
// including loading fonts, has to run on.
func onMain(t *testing.T, run func()) {
	t.Helper()
	finished := make(chan struct{})
	draw <- func() {
		defer close(finished)
		run()
	}
	<-finished
}

// TestLayout draws the shared test layout and checks it looks the way it should.
func TestLayout(t *testing.T) {
	if !gpu {
		t.Skip("set CLAY_TEST_GPU to run the tests that need a GPU")
	}
	var screen image.Image
	onMain(t, func() {
		fontList := []rl.Font{rl.LoadFontFromMemory(".ttf", fonts.RobotoRegularTTF, 48, nil)}
		rl.SetTextureFilter(fontList[0].Texture, rl.FilterBilinear)
		testlayout.Init(clayraylib.MeasureText, &fontList)
		commands := testlayout.Build()

		target := rl.LoadRenderTexture(testlayout.Width, testlayout.Height)
		defer rl.UnloadRenderTexture(target)
		rl.BeginDrawing()
		rl.BeginTextureMode(target)
		rl.ClearBackground(rl.Black)
		clayraylib.ClayRender(commands, fontList)
		rl.EndTextureMode()
		rl.EndDrawing()

		// A render texture is stored upside down.
		drawn := rl.LoadImageFromTexture(target.Texture)
		defer rl.UnloadImage(drawn)
		rl.ImageFlipVertical(drawn)
		screen = drawn.ToImage()
	})
	testlayout.Check(t, screen)
}
