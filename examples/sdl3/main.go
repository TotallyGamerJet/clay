package main

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"unsafe"

	"github.com/totallygamerjet/clay"
	sl "github.com/totallygamerjet/clay/examples/shared-layouts"
	sdl3 "github.com/totallygamerjet/clay/renderers/sdl3"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/ttf"
)

func handleClayError(errorText clay.ErrorData) {
	slog.Error(errorText.ErrorText.String(), "stacktrace", debug.Stack())
}

// TODO: CreateArenaWithCapacityAndMemory should take a slice of bytes

func main() {
	/*if err := sdl.LoadLibrary(sdl.Path()); err != nil {
		panic(err)
	}
	if err := ttf.LoadLibrary(ttf.Path()); err != nil {
		panic(err)
	}*/
	defer binsdl.Load().Unload()
	defer binttf.Load().Unload()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	font, err := ttf.OpenFont("resources/Roboto-Regular.ttf", 16)
	if err != nil {
		panic(err)
	}

	fonts := []sdl3.Font{
		sl.FONT_ID_BODY_16: {
			FontId: sl.FONT_ID_BODY_16,
			Font:   font,
		},
	}

	var (
		window   *sdl.Window
		renderer *sdl.Renderer
	)

	window, renderer, err = sdl.CreateWindowAndRenderer("SDL", 800, 600, sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}

	const enableVsync = false
	if enableVsync {
		renderer.SetVSync(1)
	}
	_ = renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND) //for alpha blending

	const screenWidth, screenHeight = 800, 600

	totalMemorySize := clay.MinMemorySize()
	memory := make([]byte, totalMemorySize)
	arena := clay.CreateArenaWithCapacityAndMemory(uint64(totalMemorySize), unsafe.Pointer(unsafe.SliceData(memory)))
	clay.Initialize(arena, clay.Dimensions{Width: screenWidth, Height: screenHeight}, clay.ErrorHandler{ErrorHandlerFunction: handleClayError})

	clay.SetMeasureTextFunction(sdl3.MeasureText, unsafe.Pointer(&fonts))

	var NOW = sdl.GetPerformanceCounter()
	var LAST uint64 = 0
	var deltaTime float64 = 0
	var demoData = sl.ClayVideoDemo_Initialize()

loop:
	for {
		scrollDelta := clay.Vector2{}
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EVENT_QUIT:
				break loop
			case sdl.EVENT_MOUSE_WHEEL:
				e := event.MouseWheelEvent()
				scrollDelta.X = float32(e.X)
				scrollDelta.Y = float32(e.Y)
			}
		}

		LAST = NOW
		NOW = sdl.GetPerformanceCounter()
		deltaTime = (float64)((NOW-LAST)*1000) / (float64)(sdl.GetPerformanceFrequency())
		fmt.Println(deltaTime)

		mouseState, mouseX, mouseY := sdl.GetMouseState()
		var mousePosition = clay.Vector2{X: float32(mouseX), Y: float32(mouseY)}
		clay.SetPointerState(mousePosition, mouseState&sdl.ButtonMask(1) != 0)

		clay.UpdateScrollContainers(
			true,
			clay.Vector2{X: scrollDelta.X, Y: scrollDelta.Y},
			float32(deltaTime),
		)

		windowWidth, windowHeight, _ := window.Size()
		clay.SetLayoutDimensions(clay.Dimensions{Width: float32(windowWidth), Height: float32(windowHeight)})

		renderCommands := sl.ClayVideoDemo_CreateLayout(&demoData)
		_ = renderer.SetDrawColor(0, 0, 0, 255)
		_ = renderer.Clear()

		sdl3.ClayRender(renderer, renderCommands, fonts)

		renderer.Present()
	}

}
