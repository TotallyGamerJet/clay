package main

import (
	"fmt"
	"unsafe"

	"github.com/totallygamerjet/clay"
	"github.com/totallygamerjet/clay/examples/fonts"
	"github.com/totallygamerjet/clay/examples/videodemo"
	"github.com/totallygamerjet/clay/renderers/sdl2"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func handleClayError(errorData clay.ErrorData) {
	panic(errorData)
}

// TODO: CreateArenaWithCapacityAndMemory should take a slice of bytes

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	stream, err := sdl.RWFromMem(fonts.RobotoRegularTTF)
	if err != nil {
		panic(err)
	}

	font, err := ttf.OpenFontRW(stream, 1, 16)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	fonts := []sdl2.Font{
		videodemo.FontIdBody16: {
			FontId: videodemo.FontIdBody16,
			Font:   font,
		},
	}

	var (
		window   *sdl.Window
		renderer *sdl.Renderer
	)

	sdl.SetHint(sdl.HINT_RENDER_DRIVER, "opengl")

	window, err = sdl.CreateWindow("SDL", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	_ = sdl.GLSetAttribute(sdl.GL_MULTISAMPLESAMPLES, 4) //for antialiasing

	const enableVsync = false
	if enableVsync {
		renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC) //"SDL_RENDERER_ACCELERATED" is for antialiasing
	} else {
		renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	}
	_ = renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND) //for alpha blending

	const screenWidth, screenHeight = 800, 600

	totalMemorySize := clay.MinMemorySize()
	memory := make([]byte, totalMemorySize)
	arena := clay.CreateArenaWithCapacityAndMemory(uint64(totalMemorySize), unsafe.Pointer(unsafe.SliceData(memory)))
	clay.Initialize(arena, clay.Dimensions{Width: screenWidth, Height: screenHeight}, clay.ErrorHandler{ErrorHandlerFunction: handleClayError})

	clay.SetMeasureTextFunction(sdl2.MeasureText, unsafe.Pointer(&fonts))

	var NOW = sdl.GetPerformanceCounter()
	var LAST uint64 = 0
	var deltaTime float64 = 0
	var demoData = videodemo.Initialize()

loop:
	for {
		scrollDelta := clay.Vector2{}
		for {
			event := sdl.PollEvent()
			if event == nil {
				break
			}
			switch e := event.(type) {
			case *sdl.QuitEvent:
				break loop
			case *sdl.MouseWheelEvent:
				scrollDelta.X = float32(e.X)
				scrollDelta.Y = float32(e.Y)
			}
		}

		LAST = NOW
		NOW = sdl.GetPerformanceCounter()
		deltaTime = (float64)((NOW-LAST)*1000) / (float64)(sdl.GetPerformanceFrequency())
		fmt.Println(deltaTime)

		mouseX, mouseY, mouseState := sdl.GetMouseState()
		var mousePosition = clay.Vector2{X: float32(mouseX), Y: float32(mouseY)}
		clay.SetPointerState(mousePosition, mouseState&sdl.Button(1) != 0)

		clay.UpdateScrollContainers(
			true,
			clay.Vector2{X: scrollDelta.X, Y: scrollDelta.Y},
			float32(deltaTime),
		)

		windowWidth, windowHeight := window.GetSize()
		clay.SetLayoutDimensions(clay.Dimensions{Width: float32(windowWidth), Height: float32(windowHeight)})

		renderCommands := videodemo.CreateLayout(&demoData)
		_ = renderer.SetDrawColor(0, 0, 0, 255)
		_ = renderer.Clear()

		_ = sdl2.ClayRender(renderer, renderCommands, fonts)

		renderer.Present()
	}

}
