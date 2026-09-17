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
	return clay.Dimensions{
		Width:  size.X,
		Height: float32(config.FontSize),
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
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			rl.DrawTextEx(
				fontFor(fonts, config.FontId),
				config.StringContents,
				rl.Vector2{X: boundingBox.X, Y: boundingBox.Y},
				float32(config.FontSize),
				float32(config.LetterSpacing),
				toColor(config.TextColor),
			)
		case clay.RENDER_COMMAND_TYPE_IMAGE:
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
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
			rl.BeginScissorMode(
				int32(math.Round(float64(boundingBox.X))),
				int32(math.Round(float64(boundingBox.Y))),
				int32(math.Round(float64(boundingBox.Width))),
				int32(math.Round(float64(boundingBox.Height))),
			)
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
			rl.EndScissorMode()
		case clay.RENDER_COMMAND_TYPE_OVERLAY_COLOR_START:
			overlays.Push(renderCommand.RenderData.OverlayColor.Color)
			setOverlay(overlays)
		case clay.RENDER_COMMAND_TYPE_OVERLAY_COLOR_END:
			overlays.Pop()
			setOverlay(overlays)
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			if config.CornerRadius.TopLeft > 0 {
				roundness := config.CornerRadius.TopLeft * 2 / min(boundingBox.Width, boundingBox.Height)
				rl.DrawRectangleRounded(rect, roundness, 8, toColor(config.BackgroundColor))
			} else {
				rl.DrawRectangleRec(rect, toColor(config.BackgroundColor))
			}
		case clay.RENDER_COMMAND_TYPE_BORDER:
			renderBorder(boundingBox, &renderCommand.RenderData.Border)
		case clay.RENDER_COMMAND_TYPE_NONE:
		case clay.RENDER_COMMAND_TYPE_CUSTOM:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}
}

func renderBorder(bb clay.BoundingBox, config *clay.BorderRenderData) {
	col := toColor(config.Color)
	r := config.CornerRadius
	maxRadius := min(bb.Width, bb.Height) / 2
	r.TopLeft = min(r.TopLeft, maxRadius)
	r.TopRight = min(r.TopRight, maxRadius)
	r.BottomLeft = min(r.BottomLeft, maxRadius)
	r.BottomRight = min(r.BottomRight, maxRadius)
	w := config.Width

	if w.Left > 0 {
		rl.DrawRectangleV(rl.Vector2{X: bb.X, Y: bb.Y + r.TopLeft}, rl.Vector2{X: float32(w.Left), Y: bb.Height - r.TopLeft - r.BottomLeft}, col)
	}
	if w.Right > 0 {
		rl.DrawRectangleV(rl.Vector2{X: bb.X + bb.Width - float32(w.Right), Y: bb.Y + r.TopRight}, rl.Vector2{X: float32(w.Right), Y: bb.Height - r.TopRight - r.BottomRight}, col)
	}
	if w.Top > 0 {
		rl.DrawRectangleV(rl.Vector2{X: bb.X + r.TopLeft, Y: bb.Y}, rl.Vector2{X: bb.Width - r.TopLeft - r.TopRight, Y: float32(w.Top)}, col)
	}
	if w.Bottom > 0 {
		rl.DrawRectangleV(rl.Vector2{X: bb.X + r.BottomLeft, Y: bb.Y + bb.Height - float32(w.Bottom)}, rl.Vector2{X: bb.Width - r.BottomLeft - r.BottomRight, Y: float32(w.Bottom)}, col)
	}

	round := func(v float32) float32 { return float32(math.Round(float64(v))) }
	corner := func(cx, cy, radius float32, width uint16, start, end float32) {
		if radius > 0 && width > 0 {
			rl.DrawRing(rl.Vector2{X: round(cx), Y: round(cy)}, round(max(radius-float32(width), 0)), radius, start, end, 10, col)
		}
	}
	corner(bb.X+r.TopLeft, bb.Y+r.TopLeft, r.TopLeft, w.Top, 180, 270)
	corner(bb.X+bb.Width-r.TopRight, bb.Y+r.TopRight, r.TopRight, w.Top, 270, 360)
	corner(bb.X+r.BottomLeft, bb.Y+bb.Height-r.BottomLeft, r.BottomLeft, w.Bottom, 90, 180)
	corner(bb.X+bb.Width-r.BottomRight, bb.Y+bb.Height-r.BottomRight, r.BottomRight, w.Bottom, 0, 90)
}
