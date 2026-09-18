// SPDX-License-Identifier: Zlib

// Package raylib renders clay with raylib, using github.com/gen2brain/raylib-go.
//
// It is a port of clay's renderers/raylib/clay_renderer_raylib.c.
package raylib

import (
	"image/color"
	"log/slog"
	"math"

	"github.com/TotallyGamerJet/clay"
	"github.com/TotallyGamerJet/clay/renderers/internal/overlay"
	"github.com/TotallyGamerJet/clay/renderers/internal/shapes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func toColor(c clay.Color) color.RGBA {
	round := func(v float32) uint8 { return uint8(math.Round(float64(v))) }
	return color.RGBA{R: round(c.R), G: round(c.G), B: round(c.B), A: round(c.A)}
}

func fontFor(fonts []rl.Font, id uint16) rl.Font {
	// Raylib ships with a default font, which is used if the font failed to load.
	if int(id) >= len(fonts) || !rl.IsFontValid(fonts[id]) {
		return rl.GetFontDefault()
	}
	return fonts[id]
}

// MeasureText measures text for clay. userData must be a *[]rl.Font indexed by font id.
func MeasureText(text string, config *clay.TextElementConfig, userData any) clay.Dimensions {
	fonts := *userData.(*[]rl.Font)
	size := rl.MeasureTextEx(fontFor(fonts, config.FontId), text, float32(config.FontSize), float32(config.LetterSpacing))
	height := float32(config.FontSize)
	if config.LineHeight > 0 {
		height = float32(config.LineHeight)
	}
	return clay.Dimensions{
		Width:  size.X,
		Height: height,
	}
}

// drawMesh draws a shape in a single color. Raylib anti aliases with multisampling,
// so the shapes are built without a feathered edge.
func drawMesh(mesh shapes.Mesh, c clay.Color) {
	col := toColor(c)
	point := func(i uint16) rl.Vector2 {
		v := mesh.Vertices[i]
		return rl.Vector2{X: v.X, Y: v.Y}
	}
	for i := 0; i+2 < len(mesh.Indices); i += 3 {
		// Raylib wants the vertices of a triangle counter clockwise, which is the other
		// way around from the shapes, whose y axis points down.
		rl.DrawTriangle(point(mesh.Indices[i+2]), point(mesh.Indices[i+1]), point(mesh.Indices[i]), col)
	}
}

const overlayShaderCode = `#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

uniform sampler2D texture0;
uniform vec4 overlayColor;

out vec4 finalColor;

void main()
{
    vec4 texelColor = texture(texture0, fragTexCoord) * fragColor;

    vec3 blendedRGB = mix(texelColor.rgb, overlayColor.rgb, overlayColor.a);

    finalColor = vec4(blendedRGB, texelColor.a);
}`

var (
	overlayShader   rl.Shader
	overlayColorLoc int32
)

// setOverlay makes everything drawn blend towards the active overlays.
func setOverlay(overlays overlay.Stack) {
	if overlayShader.ID == 0 {
		overlayShader = rl.LoadShaderFromMemory("", overlayShaderCode)
		overlayColorLoc = rl.GetShaderLocation(overlayShader, "overlayColor")
	}
	rl.EndShaderMode()
	if !overlays.Active() {
		return
	}
	c := overlays.Combined()
	rl.SetShaderValue(overlayShader, overlayColorLoc, []float32{c.R / 255, c.G / 255, c.B / 255, c.A / 255}, rl.ShaderUniformVec4)
	rl.BeginShaderMode(overlayShader)
}

// ClayRender draws the render commands. It must be called between rl.BeginDrawing and rl.EndDrawing.
//
// The image data of image elements must be a *rl.Texture2D.
func ClayRender(renderCommands clay.RenderCommandArray, fonts []rl.Font) {
	var overlays overlay.Stack
	defer func() {
		if overlays.Active() {
			rl.EndShaderMode()
		}
	}()
	for _, renderCommand := range renderCommands {
		boundingBox := renderCommand.BoundingBox
		rect := rl.Rectangle{X: boundingBox.X, Y: boundingBox.Y, Width: boundingBox.Width, Height: boundingBox.Height}
		switch renderCommand.CommandType {
		case clay.RenderCommandTypeText:
			config := &renderCommand.RenderData.Text
			// Center the line in its box, which is taller than the text when a line height is set.
			y := boundingBox.Y
			if float32(config.FontSize) < boundingBox.Height {
				y += (boundingBox.Height - float32(config.FontSize)) / 2
			}
			rl.DrawTextEx(
				fontFor(fonts, config.FontId),
				config.StringContents,
				rl.Vector2{X: boundingBox.X, Y: y},
				float32(config.FontSize),
				float32(config.LetterSpacing),
				toColor(config.TextColor),
			)
		case clay.RenderCommandTypeImage:
			config := &renderCommand.RenderData.Image
			texture := *config.ImageData.(*rl.Texture2D)
			tint := config.BackgroundColor
			if tint == (clay.Color{}) {
				tint = clay.Color{R: 255, G: 255, B: 255, A: 255}
			}
			rl.DrawTexturePro(
				texture,
				rl.Rectangle{Width: float32(texture.Width), Height: float32(texture.Height)},
				rect,
				rl.Vector2{},
				0,
				toColor(tint),
			)
		case clay.RenderCommandTypeScissorStart:
			rl.BeginScissorMode(
				int32(math.Round(float64(boundingBox.X))),
				int32(math.Round(float64(boundingBox.Y))),
				int32(math.Round(float64(boundingBox.Width))),
				int32(math.Round(float64(boundingBox.Height))),
			)
		case clay.RenderCommandTypeScissorEnd:
			rl.EndScissorMode()
		case clay.RenderCommandTypeOverlayColorStart:
			overlays.Push(renderCommand.RenderData.OverlayColor.Color)
			setOverlay(overlays)
		case clay.RenderCommandTypeOverlayColorEnd:
			overlays.Pop()
			setOverlay(overlays)
		case clay.RenderCommandTypeRectangle:
			config := &renderCommand.RenderData.Rectangle
			if config.CornerRadius != (clay.CornerRadius{}) {
				drawMesh(shapes.Fill(boundingBox, config.CornerRadius, 0), config.BackgroundColor)
			} else {
				rl.DrawRectangleRec(rect, toColor(config.BackgroundColor))
			}
		case clay.RenderCommandTypeBorder:
			config := &renderCommand.RenderData.Border
			drawMesh(shapes.Border(boundingBox, config.CornerRadius, config.Width, 0), config.Color)
		case clay.RenderCommandTypeNone:
		case clay.RenderCommandTypeCustom:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}
}
