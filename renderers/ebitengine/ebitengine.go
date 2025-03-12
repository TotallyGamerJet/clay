package ebitengine

import (
	"fmt"
	"image"
	"image/color"
	"log/slog"
	"math"
	"strings"
	"unsafe"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/totallygamerjet/clay"
)

var (
	whiteImage *ebiten.Image
)

func init() {
	img := ebiten.NewImage(3, 3)
	img.Fill(color.White)
	whiteImage = img.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
}

type RendererData struct {
	Screen *ebiten.Image
	Fonts  []text.Face
}

func MeasureText(txt clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	fonts := *(*[]text.Face)(userData)

	font := fonts[config.FontId]

	width, height := text.Measure(txt.String(), font, font.Metrics().HLineGap)
	return clay.Dimensions{
		Width:  float32(width),
		Height: float32(height),
	}
}

func ClayRender(rendererData *RendererData, renderCommands clay.RenderCommandArray) {
	screen := rendererData.Screen
	fonts := rendererData.Fonts
	for i := int32(0); i < renderCommands.Length; i++ {
		renderCommand := clay.RenderCommandArray_Get(&renderCommands, i)
		boundingBox := renderCommand.BoundingBox
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			if config.CornerRadius.TopLeft > 0 {
				RenderFillRoundedRect(screen, boundingBox, config.CornerRadius.TopLeft, config.BackgroundColor)
			} else {
				vector.DrawFilledRect(
					screen,
					boundingBox.X,
					boundingBox.Y,
					boundingBox.Width, boundingBox.Height,
					color.RGBA{
						uint8(config.BackgroundColor.R),
						uint8(config.BackgroundColor.G),
						uint8(config.BackgroundColor.B),
						uint8(config.BackgroundColor.A),
					}, true,
				)
			}
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			cloned := strings.Clone(config.StringContents.String())
			font := fonts[config.FontId]

			opts := &text.DrawOptions{}
			opts.ColorScale.Scale(
				config.TextColor.R/255,
				config.TextColor.G/255,
				config.TextColor.B/255,
				config.TextColor.A/255,
			)
			opts.GeoM.Translate(float64(boundingBox.X), float64(boundingBox.Y))
			text.Draw(screen, cloned, font, opts)
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
			screen = rendererData.Screen.SubImage(image.Rect(
				int(boundingBox.X), int(boundingBox.Y),
				int(boundingBox.X+boundingBox.Width),
				int(boundingBox.Y+boundingBox.Height),
			)).(*ebiten.Image)
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
			screen = rendererData.Screen
		case clay.RENDER_COMMAND_TYPE_IMAGE:
			config := &renderCommand.RenderData.Image
			img := (*ebiten.Image)(config.ImageData)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(boundingBox.X), float64(boundingBox.Y))
			screen.DrawImage(img, opts)
		case clay.RENDER_COMMAND_TYPE_BORDER:
			panic("not implemented")
			//Clay_BorderRenderData *config = &renderCommand->renderData.border;
			//SDL_SetRenderDrawColor(renderer, CLAY_COLOR_TO_SDL_COLOR_ARGS(config->color));
			//
			//if(boundingBox.width > 0 & boundingBox.height > 0){
			//   const float maxRadius = SDL_min(boundingBox.width, boundingBox.height) / 2.0f;
			//
			//   if (config->width.left > 0) {
			//	   const float clampedRadiusTop = SDL_min((float)config->cornerRadius.topLeft, maxRadius);
			//	   const float clampedRadiusBottom = SDL_min((float)config->cornerRadius.bottomLeft, maxRadius);
			//	   SDL_FRect rect = {
			//		   boundingBox.x,
			//		   boundingBox.y + clampedRadiusTop,
			//		   (float)config->width.left,
			//		   (float)boundingBox.height - clampedRadiusTop - clampedRadiusBottom
			//	   };
			//	   SDL_RenderFillRectF(renderer, &rect);
			//   }
			//
			//   if (config->width.right > 0) {
			//	   const float clampedRadiusTop = SDL_min((float)config->cornerRadius.topRight, maxRadius);
			//	   const float clampedRadiusBottom = SDL_min((float)config->cornerRadius.bottomRight, maxRadius);
			//	   SDL_FRect rect = {
			//		   boundingBox.x + boundingBox.width - config->width.right,
			//		   boundingBox.y + clampedRadiusTop,
			//		   (float)config->width.right,
			//		   (float)boundingBox.height - clampedRadiusTop - clampedRadiusBottom
			//	   };
			//	   SDL_RenderFillRectF(renderer, &rect);
			//   }
			//
			//   if (config->width.top > 0) {
			//	   const float clampedRadiusLeft = SDL_min((float)config->cornerRadius.topLeft, maxRadius);
			//	   const float clampedRadiusRight = SDL_min((float)config->cornerRadius.topRight, maxRadius);
			//	   SDL_FRect rect = {
			//		   boundingBox.x + clampedRadiusLeft,
			//		   boundingBox.y,
			//		   boundingBox.width - clampedRadiusLeft - clampedRadiusRight,
			//		   (float)config->width.top };
			//	   SDL_RenderFillRectF(renderer, &rect);
			//   }
			//
			//   if (config->width.bottom > 0) {
			//	   const float clampedRadiusLeft = SDL_min((float)config->cornerRadius.bottomLeft, maxRadius);
			//	   const float clampedRadiusRight = SDL_min((float)config->cornerRadius.bottomRight, maxRadius);
			//	   SDL_FRect rect = {
			//		   boundingBox.x + clampedRadiusLeft,
			//		   boundingBox.y + boundingBox.height - config->width.bottom,
			//		   boundingBox.width - clampedRadiusLeft - clampedRadiusRight,
			//		   (float)config->width.bottom
			//	   };
			//	   SDL_RenderFillRectF(renderer, &rect);
			//   }
			//
			//   //corner index: 0->3 topLeft -> CW -> bottonLeft
			//   if (config->width.top > 0 & config->cornerRadius.topLeft > 0) {
			//	   SDL_RenderCornerBorder(renderer, &boundingBox, config, 0, config->color);
			//   }
			//
			//   if (config->width.top > 0 & config->cornerRadius.topRight> 0) {
			//	   SDL_RenderCornerBorder(renderer, &boundingBox, config, 1, config->color);
			//   }
			//
			//   if (config->width.bottom > 0 & config->cornerRadius.bottomLeft > 0) {
			//	   SDL_RenderCornerBorder(renderer, &boundingBox, config, 2, config->color);
			//   }
			//
			//   if (config->width.bottom > 0 & config->cornerRadius.bottomLeft > 0) {
			//	   SDL_RenderCornerBorder(renderer, &boundingBox, config, 3, config->color);
			//   }
			//}
		case clay.RENDER_COMMAND_TYPE_NONE:
		case clay.RENDER_COMMAND_TYPE_CUSTOM:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}
}

const NUM_CIRCLE_SEGMENTS = 16

func RenderFillRoundedRect(screen *ebiten.Image, rect clay.BoundingBox, cornerRadius float32, _color clay.Color) error {
	r := _color.R / 255
	g := _color.G / 255
	b := _color.B / 255
	a := _color.A / 255

	indexCount, vertexCount := 0, uint16(0)

	minRadius := min(rect.Width, rect.Height) / 2.0
	clampedRadius := min(cornerRadius, minRadius)

	numCircleSegments := max(NUM_CIRCLE_SEGMENTS, int(clampedRadius*0.5)) // check if needs to be clamped

	var vertices [512]ebiten.Vertex
	var indices [512]uint16

	//define center rectangle
	// 0 Center TL
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + clampedRadius,
		DstY:   rect.Y + clampedRadius,
		SrcX:   1,
		SrcY:   1,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++
	// 1 Center TR
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + rect.Width - clampedRadius,
		DstY:   rect.Y + clampedRadius,
		SrcX:   2,
		SrcY:   1,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++
	// 2 Center BR
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + rect.Width - clampedRadius,
		DstY:   rect.Y + rect.Height - clampedRadius,
		SrcX:   2,
		SrcY:   2,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++
	// 3 Center BL
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + clampedRadius,
		DstY:   rect.Y + rect.Height - clampedRadius,
		SrcX:   1,
		SrcY:   2,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++

	indices[indexCount] = 0
	indexCount++
	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = 3
	indexCount++
	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = 2
	indexCount++
	indices[indexCount] = 3
	indexCount++

	//define rounded corners as triangle fans
	step := (math.Pi / 2) / float32(numCircleSegments)
	for i := 0; i < numCircleSegments; i++ {
		angle1 := float32(i) * step
		angle2 := (float32(i) + 1) * step

		for j := uint16(0); j < 4; j++ {
			var cx, cy, signX, signY float32

			switch j {
			case 0:
				cx = rect.X + clampedRadius
				cy = rect.Y + clampedRadius
				signX = -1
				signY = -1
			case 1:
				cx = rect.X + rect.Width - clampedRadius
				cy = rect.Y + clampedRadius
				signX = 1
				signY = -1
			case 2:
				cx = rect.X + rect.Width - clampedRadius
				cy = rect.Y + rect.Height - clampedRadius
				signX = 1
				signY = 1
			case 3:
				cx = rect.X + clampedRadius
				cy = rect.Y + rect.Height - clampedRadius
				signX = -1
				signY = 1
			default:
				return fmt.Errorf("out of bounds index: %d", j)
			}

			vertices[vertexCount] = ebiten.Vertex{
				DstX:   cx + float32(math.Cos(float64(angle1)))*clampedRadius*signX,
				DstY:   cy + float32(math.Sin(float64(angle1)))*clampedRadius*signY,
				SrcX:   1,
				SrcY:   1,
				ColorR: r,
				ColorG: g,
				ColorB: b,
				ColorA: a,
			}
			vertexCount++
			vertices[vertexCount] = ebiten.Vertex{
				DstX:   cx + float32(math.Cos(float64(angle2)))*clampedRadius*signX,
				DstY:   cy + float32(math.Sin(float64(angle2)))*clampedRadius*signY,
				SrcX:   1,
				SrcY:   1,
				ColorR: r,
				ColorG: g,
				ColorB: b,
				ColorA: a,
			}
			vertexCount++

			indices[indexCount] = j // Connect to corresponding central rectangle vertex
			indexCount++
			indices[indexCount] = vertexCount - 2
			indexCount++
			indices[indexCount] = vertexCount - 1
			indexCount++
		}
	}
	//Define edge rectangles
	// Top edge
	// TL
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + clampedRadius,
		DstY:   rect.Y,
		SrcX:   0,
		SrcY:   0,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++
	// TR
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + rect.Width - clampedRadius,
		DstY:   rect.Y,
		SrcX:   1,
		SrcY:   0,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++

	indices[indexCount] = 0
	indexCount++
	indices[indexCount] = vertexCount - 2 //TL
	indexCount++
	indices[indexCount] = vertexCount - 1 //TR
	indexCount++
	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = 0
	indexCount++
	indices[indexCount] = vertexCount - 1 //TR
	indexCount++
	// Right edge
	// RT
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + rect.Width,
		DstY:   rect.Y + clampedRadius,
		SrcX:   1,
		SrcY:   0,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++
	// RB
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + rect.Width,
		DstY:   rect.Y + rect.Height - clampedRadius,
		SrcX:   1,
		SrcY:   1,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++

	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = vertexCount - 2 //RT
	indexCount++
	indices[indexCount] = vertexCount - 1 //RB
	indexCount++
	indices[indexCount] = 2
	indexCount++
	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = vertexCount - 1 //RB
	indexCount++
	// Bottom edge
	// BR
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + rect.Width - clampedRadius,
		DstY:   rect.Y + rect.Height,
		SrcX:   1,
		SrcY:   1,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++
	// BL
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X + clampedRadius,
		DstY:   rect.Y + rect.Height,
		SrcX:   0,
		SrcY:   1,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++

	indices[indexCount] = 2
	indexCount++
	indices[indexCount] = vertexCount - 2 //BR
	indexCount++
	indices[indexCount] = vertexCount - 1 //BL
	indexCount++
	indices[indexCount] = 3
	indexCount++
	indices[indexCount] = 2
	indexCount++
	indices[indexCount] = vertexCount - 1 //BL
	indexCount++
	// Left edge
	// LB
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X,
		DstY:   rect.Y + rect.Height - clampedRadius,
		SrcX:   0,
		SrcY:   1,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++
	// LT
	vertices[vertexCount] = ebiten.Vertex{
		DstX:   rect.X,
		DstY:   rect.Y + clampedRadius,
		SrcX:   0,
		SrcY:   0,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
	vertexCount++

	indices[indexCount] = 3
	indexCount++
	indices[indexCount] = vertexCount - 2 //LB
	indexCount++
	indices[indexCount] = vertexCount - 1 //LT
	indexCount++
	indices[indexCount] = 0
	indexCount++
	indices[indexCount] = 3
	indexCount++
	indices[indexCount] = vertexCount - 1 //LT
	indexCount++

	// Render everything
	screen.DrawTriangles(vertices[:vertexCount], indices[:indexCount], whiteImage, &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	})

	return nil
}
