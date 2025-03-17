package sdl3

import (
	"fmt"
	"log/slog"
	"math"
	"strings"
	"unsafe"

	"github.com/TotallyGamerJet/clay"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

type RendererData struct {
	Renderer   *sdl.Renderer
	TextEngine *ttf.TextEngine
	Fonts      []*ttf.Font
}

func MeasureText(text clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	fonts := *(*[]*ttf.Font)(userData)
	font := fonts[config.FontId]

	width, height, err := font.StringSize(text.String())
	if err != nil {
		panic(fmt.Errorf("sdl3: failed to measure text: %w", err))
	}

	return clay.Dimensions{
		Width:  float32(width),
		Height: float32(height),
	}
}

func ClayRender(rendererData *RendererData, renderCommands clay.RenderCommandArray) error {
	renderer := rendererData.Renderer
	fonts := rendererData.Fonts
	textEngine := rendererData.TextEngine
	for renderCommand := range renderCommands.Iter() {
		boundingBox := renderCommand.BoundingBox
		rect := sdl.FRect{
			X: boundingBox.X,
			Y: boundingBox.Y,
			W: boundingBox.Width,
			H: boundingBox.Height,
		}
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
			renderer.SetDrawColor(
				uint8(config.BackgroundColor.R),
				uint8(config.BackgroundColor.G),
				uint8(config.BackgroundColor.B),
				uint8(config.BackgroundColor.A),
			)
			if config.CornerRadius.TopLeft > 0 {
				err := renderFillRoundedRect(renderer, rect, config.CornerRadius.TopLeft, config.BackgroundColor)
				if err != nil {
					return err
				}
			} else {
				renderer.RenderFillRect(&rect)
			}
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			cloned := strings.Clone(config.StringContents.String())
			font := fonts[config.FontId]
			text, err := textEngine.CreateText(font, cloned)
			if err != nil {
				return err
			}
			text.SetColor(
				uint8(config.TextColor.R),
				uint8(config.TextColor.G),
				uint8(config.TextColor.B),
				uint8(config.TextColor.A),
			)
			text.DrawRenderer(rect.X, rect.Y)
			text.Destroy()
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
			currentClippingRectangle := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if err := renderer.SetClipRect(&currentClippingRectangle); err != nil {
				return err
			}
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
			if err := renderer.SetClipRect(nil); err != nil {
				return err
			}
		case clay.RENDER_COMMAND_TYPE_IMAGE:
			config := &renderCommand.RenderData.Image
			texture, err := renderer.CreateTextureFromSurface((*sdl.Surface)(config.ImageData))
			if err != nil {
				return err
			}
			destination := sdl.FRect{
				X: rect.X,
				Y: rect.Y,
				W: rect.W,
				H: rect.H,
			}
			if err := renderer.RenderTexture(texture, nil, &destination); err != nil {
				return err
			}
			texture.Destroy()
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
	return nil
}

const numCircleSegments = 16

func renderFillRoundedRect(renderer *sdl.Renderer, rect sdl.FRect, cornerRadius float32, _color clay.Color) error {
	color := sdl.FColor{
		R: _color.R / 255,
		G: _color.G / 255,
		B: _color.B / 255,
		A: _color.A / 255,
	}

	indexCount, vertexCount := 0, int32(0)

	minRadius := min(rect.W, rect.H) / 2.0
	clampedRadius := min(cornerRadius, minRadius)

	numCircleSegments := max(numCircleSegments, int(clampedRadius*0.5)) // check if needs to be clamped

	var vertices [512]sdl.Vertex
	var indices [512]int32

	// define center rectangle
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + clampedRadius, Y: rect.Y + clampedRadius}, Color: color, TexCoord: sdl.FPoint{}} // 0 center TL
	vertexCount++
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + rect.W - clampedRadius, Y: rect.Y + clampedRadius}, Color: color, TexCoord: sdl.FPoint{X: 1}} // 1 center TR
	vertexCount++
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + rect.W - clampedRadius, Y: rect.Y + rect.H - clampedRadius}, Color: color, TexCoord: sdl.FPoint{X: 1, Y: 1}} // 2 center BR
	vertexCount++
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + clampedRadius, Y: rect.Y + rect.H - clampedRadius}, Color: color, TexCoord: sdl.FPoint{Y: 1}} // 3 center BL
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

	// define rounded corners as triangle fans
	step := (math.Pi / 2) / float32(numCircleSegments)
	for i := 0; i < numCircleSegments; i++ {
		angle1 := float32(i) * step
		angle2 := (float32(i) + 1) * step

		for j := int32(0); j < 4; j++ {
			var cx, cy, signX, signY float32

			switch j {
			case 0:
				cx = rect.X + clampedRadius
				cy = rect.Y + clampedRadius
				signX = -1
				signY = -1
			case 1:
				cx = rect.X + rect.W - clampedRadius
				cy = rect.Y + clampedRadius
				signX = 1
				signY = -1
			case 2:
				cx = rect.X + rect.W - clampedRadius
				cy = rect.Y + rect.H - clampedRadius
				signX = 1
				signY = 1
			case 3:
				cx = rect.X + clampedRadius
				cy = rect.Y + rect.H - clampedRadius
				signX = -1
				signY = 1
			default:
				return fmt.Errorf("out of bounds index: %d", j)
			}

			vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: cx + float32(math.Cos(float64(angle1)))*clampedRadius*signX, Y: cy + float32(math.Sin(float64(angle1)))*clampedRadius*signY}, Color: color, TexCoord: sdl.FPoint{}}
			vertexCount++
			vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: cx + float32(math.Cos(float64(angle2)))*clampedRadius*signX, Y: cy + float32(math.Sin(float64(angle2)))*clampedRadius*signY}, Color: color, TexCoord: sdl.FPoint{}}
			vertexCount++

			indices[indexCount] = j // Connect to corresponding central rectangle vertex
			indexCount++
			indices[indexCount] = vertexCount - 2
			indexCount++
			indices[indexCount] = vertexCount - 1
			indexCount++
		}
	}
	// Define edge rectangles
	// Top edge
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + clampedRadius, Y: rect.Y}, Color: color, TexCoord: sdl.FPoint{}} // TL
	vertexCount++
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + rect.W - clampedRadius, Y: rect.Y}, Color: color, TexCoord: sdl.FPoint{X: 1}} // TR
	vertexCount++

	indices[indexCount] = 0
	indexCount++
	indices[indexCount] = vertexCount - 2 // TL
	indexCount++
	indices[indexCount] = vertexCount - 1 // TR
	indexCount++
	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = 0
	indexCount++
	indices[indexCount] = vertexCount - 1 // TR
	indexCount++
	// Right edge
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + rect.W, Y: rect.Y + clampedRadius}, Color: color, TexCoord: sdl.FPoint{X: 1}} // RT
	vertexCount++
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + rect.W, Y: rect.Y + rect.H - clampedRadius}, Color: color, TexCoord: sdl.FPoint{X: 1, Y: 1}} // RB
	vertexCount++

	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = vertexCount - 2 // RT
	indexCount++
	indices[indexCount] = vertexCount - 1 // RB
	indexCount++
	indices[indexCount] = 2
	indexCount++
	indices[indexCount] = 1
	indexCount++
	indices[indexCount] = vertexCount - 1 // RB
	indexCount++
	// Bottom edge
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + rect.W - clampedRadius, Y: rect.Y + rect.H}, Color: color, TexCoord: sdl.FPoint{X: 1, Y: 1}} // BR
	vertexCount++
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X + clampedRadius, Y: rect.Y + rect.H}, Color: color, TexCoord: sdl.FPoint{Y: 1}} // BL
	vertexCount++

	indices[indexCount] = 2
	indexCount++
	indices[indexCount] = vertexCount - 2 // BR
	indexCount++
	indices[indexCount] = vertexCount - 1 // BL
	indexCount++
	indices[indexCount] = 3
	indexCount++
	indices[indexCount] = 2
	indexCount++
	indices[indexCount] = vertexCount - 1 // BL
	indexCount++
	// Left edge
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X, Y: rect.Y + rect.H - clampedRadius}, Color: color, TexCoord: sdl.FPoint{Y: 1}} // LB
	vertexCount++
	vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: rect.X, Y: rect.Y + clampedRadius}, Color: color, TexCoord: sdl.FPoint{}} // LT
	vertexCount++

	indices[indexCount] = 3
	indexCount++
	indices[indexCount] = vertexCount - 2 // LB
	indexCount++
	indices[indexCount] = vertexCount - 1 // LT
	indexCount++
	indices[indexCount] = 0
	indexCount++
	indices[indexCount] = 3
	indexCount++
	indices[indexCount] = vertexCount - 1 // LT
	indexCount++

	// Render everything
	return renderer.RenderGeometry(nil, vertices[:vertexCount], indices[:indexCount])
}
