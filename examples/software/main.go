package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"unsafe"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/examples/videodemo"
	"github.com/TotallyGamerJet/clay/renderers/software"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func handleClayError(errorData clay.ErrorData) {
	panic(errorData)
}

const (
	winWidth, winHeight = 640, 480
	fontSize            = 16
)

func main() {
	totalMemorySize := clay.MinMemorySize()
	memory := make([]byte, totalMemorySize)
	arena := clay.CreateArenaWithCapacityAndMemory(memory)
	clay.Initialize(arena, clay.Dimensions{Width: winWidth, Height: winHeight}, clay.ErrorHandler{ErrorHandlerFunction: handleClayError})

	parsedFont, err := opentype.Parse(fonts.RobotoRegularTTF)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     72, // Standard screen DPI
		Hinting: font.HintingFull,
	})

	faces := []font.Face{
		videodemo.FontIdBody16: face,
	}
	clay.SetMeasureTextFunction(software.MeasureText, unsafe.Pointer(&faces))
	var img image.Image = videodemo.SquirrelImage
	demoData := videodemo.Initialize(unsafe.Pointer(&img))
	window := image.NewRGBA(image.Rect(0, 0, winWidth, winHeight))
	draw.Draw(window, window.Bounds(), image.NewUniform(color.RGBA{A: 255}), image.Point{}, draw.Src)

	cmds := videodemo.CreateLayout(&demoData)

	if err := software.ClayRender(window, cmds, faces); err != nil {
		panic(err)
	}

	create, err := os.Create("out.png")
	if err != nil {
		return
	}
	defer create.Close()
	if err := png.Encode(create, window); err != nil {
		panic(err)
	}
}
