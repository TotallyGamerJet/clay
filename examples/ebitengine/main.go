package main

import (
	"bytes"
	"log/slog"
	"runtime/debug"
	"unsafe"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/totallygamerjet/clay"
	"github.com/totallygamerjet/clay/examples/fonts"
	sl "github.com/totallygamerjet/clay/examples/shared-layouts"
	"github.com/totallygamerjet/clay/renderers/ebitengine"
)

func handleClayError(errorText clay.ErrorData) {
	slog.Error(errorText.ErrorText.String(), "stacktrace", debug.Stack())
}

// TODO: CreateArenaWithCapacityAndMemory should take a slice of bytes

const (
	winWidth, winHeight = 640, 480
)

type App struct {
	demoData     sl.ClayVideoDemo_Data
	cmds         clay.RenderCommandArray
	rendererData *ebitengine.RendererData

	width  int
	height int
}

func New(demoData sl.ClayVideoDemo_Data, rendererData *ebitengine.RendererData) *App {
	return &App{
		demoData:     demoData,
		rendererData: rendererData,
	}
}

func (a *App) Update() error {
	dx, dy := ebiten.Wheel()
	if dx != 0 || dy != 0 {
		clay.UpdateScrollContainers(true, clay.Vector2{
			X: float32(dx),
			Y: float32(dy),
		}, 0.01)
	}

	x, y := ebiten.CursorPosition()
	clay.SetPointerState(clay.Vector2{
		X: float32(x),
		Y: float32(y),
	}, ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft))

	a.cmds = sl.ClayVideoDemo_CreateLayout(&a.demoData)

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.rendererData.Screen = screen
	ebitengine.ClayRender(a.rendererData, a.cmds)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	if outsideWidth != a.width || outsideHeight != a.height {
		clay.SetLayoutDimensions(clay.Dimensions{
			Width:  float32(outsideWidth),
			Height: float32(outsideHeight),
		})
	}
	a.width = outsideWidth
	a.height = outsideHeight

	return a.width, a.height
}

func main() {
	ebiten.SetWindowTitle("Ebitengine")
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	source, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.RobotoRegularTTF))
	if err != nil {
		panic(err)
	}

	font := &text.GoTextFace{
		Source: source,
		Size:   16,
	}

	rendererData := &ebitengine.RendererData{
		Fonts: []text.Face{
			font,
		},
	}

	// Initialize Clay
	totalMemorySize := clay.MinMemorySize()
	memory := make([]byte, totalMemorySize)
	arena := clay.CreateArenaWithCapacityAndMemory(uint64(totalMemorySize), unsafe.Pointer(unsafe.SliceData(memory)))
	clay.Initialize(arena, clay.Dimensions{Width: winWidth, Height: winHeight}, clay.ErrorHandler{ErrorHandlerFunction: handleClayError})
	clay.SetMeasureTextFunction(ebitengine.MeasureText, unsafe.Pointer(&rendererData.Fonts))

	var demoData = sl.ClayVideoDemo_Initialize()

	app := New(demoData, rendererData)

	if err := ebiten.RunGame(app); err != nil {
		panic(err)
	}
}
