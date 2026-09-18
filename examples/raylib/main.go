// SPDX-License-Identifier: Zlib

package main

import (
	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/examples/fonts"
	"github.com/TotallyGamerJet/clay/examples/videodemo"
	"github.com/TotallyGamerJet/clay/renderers/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleClayError(errorData clay.ErrorData) {
	panic(errorData)
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowHighdpi | rl.FlagMsaa4xHint | rl.FlagVsyncHint)
	rl.InitWindow(1024, 768, "Raylib")
	defer rl.CloseWindow()

	totalMemorySize := clay.MinMemorySize()
	arena := clay.CreateArenaWithCapacity(totalMemorySize)
	defer arena.Free()
	clay.Initialize(arena, clay.Dimensions{
		Width:  float32(rl.GetScreenWidth()),
		Height: float32(rl.GetScreenHeight()),
	}, clay.ErrorHandler{ErrorHandlerFunction: handleClayError})

	// Load the font at a large size so that it stays sharp when scaled to the font sizes of the text.
	fontList := []rl.Font{
		videodemo.FontIdBody16: rl.LoadFontFromMemory(".ttf", fonts.RobotoRegularTTF, 48, nil),
	}
	rl.SetTextureFilter(fontList[videodemo.FontIdBody16].Texture, rl.FilterBilinear)
	clay.SetMeasureTextFunction(raylib.MeasureText, &fontList)

	squirrel := rl.LoadTextureFromImage(rl.NewImageFromImage(videodemo.SquirrelImage))
	defer rl.UnloadTexture(squirrel)
	demoData := videodemo.Initialize(&squirrel)

	for !rl.WindowShouldClose() {
		// Press D to toggle the debug view.
		if rl.IsKeyPressed(rl.KeyD) {
			clay.SetDebugModeEnabled(!clay.IsDebugModeEnabled())
		}

		clay.SetLayoutDimensions(clay.Dimensions{
			Width:  float32(rl.GetScreenWidth()),
			Height: float32(rl.GetScreenHeight()),
		})

		mousePosition := rl.GetMousePosition()
		scrollDelta := rl.GetMouseWheelMoveV()
		clay.SetPointerState(clay.Vector2{X: mousePosition.X, Y: mousePosition.Y}, rl.IsMouseButtonDown(rl.MouseButtonLeft))
		clay.UpdateScrollContainers(true, clay.Vector2{X: scrollDelta.X, Y: scrollDelta.Y}, rl.GetFrameTime())

		renderCommands := videodemo.CreateLayout(&demoData, rl.GetFrameTime())

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		raylib.ClayRender(renderCommands, fontList)
		rl.EndDrawing()
	}
}
