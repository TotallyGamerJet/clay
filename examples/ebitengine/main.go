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
	"github.com/totallygamerjet/clay/examples/videodemo"
	"github.com/totallygamerjet/clay/renderers/ebitengine"
)

func handleClayError(errorText clay.ErrorData) {
	slog.Error(errorText.ErrorText.String(), "stacktrace", debug.Stack())
}

// TODO: CreateArenaWithCapacityAndMemory should take a slice of bytes

const (
	winWidth, winHeight = 640, 480
	fontSize            = 16
)

type App struct {
	demoData     videodemo.Data
	cmds         clay.RenderCommandArray
	rendererData *ebitengine.RendererData

	fontSource  *text.GoTextFaceSource
	scaleFactor float64
	width       float64
	height      float64
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

	a.cmds = videodemo.CreateLayout(&a.demoData)

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.rendererData.Screen = screen
	ebitengine.ClayRender(a.rendererData, a.cmds)
}

func (a *App) Layout(_, _ int) (int, int) {
	panic("Layout should not be called")
}

func (a *App) LayoutF(outsideWidth, outsideHeight float64) (float64, float64) {
	s := ebiten.Monitor().DeviceScaleFactor()
	if s != a.scaleFactor {
		a.rendererData.Fonts[0] = &text.GoTextFace{
			Source: a.fontSource,
			Size:   fontSize * a.scaleFactor,
		}
	}
	outsideWidth *= s
	outsideHeight *= s
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

	scaleFactor := ebiten.Monitor().DeviceScaleFactor()
	app := &App{
		rendererData: &ebitengine.RendererData{
			Fonts: []text.Face{
				&text.GoTextFace{
					Source: source,
					Size:   fontSize * scaleFactor,
				},
			},
		},

		fontSource:  source,
		scaleFactor: scaleFactor,
	}

	// Initialize Clay
	totalMemorySize := clay.MinMemorySize()
	memory := make([]byte, totalMemorySize)
	arena := clay.CreateArenaWithCapacityAndMemory(uint64(totalMemorySize), unsafe.Pointer(unsafe.SliceData(memory)))
	clay.Initialize(arena, clay.Dimensions{Width: winWidth, Height: winHeight}, clay.ErrorHandler{ErrorHandlerFunction: handleClayError})
	clay.SetMeasureTextFunction(ebitengine.MeasureText, unsafe.Pointer(&app.rendererData.Fonts))

	app.demoData = videodemo.Initialize()

	if err := ebiten.RunGame(app); err != nil {
		panic(err)
	}
}
