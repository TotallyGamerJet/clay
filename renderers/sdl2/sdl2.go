package sdl2

import (
	"fmt"
	"log/slog"
	"math"
	"strings"
	"unsafe"

	"github.com/TotallyGamerJet/clay"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	FontId uint32
	Font   *ttf.Font
}

func MeasureText(text clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	fonts := *(*[]Font)(userData)
	font := fonts[config.FontId].Font
	chars := strings.Clone(text.String())

	width, height, err := font.SizeUTF8(chars)
	if err != nil {
		panic(fmt.Errorf("sdl2: failed to measure text: %w", err))
	}

	return clay.Dimensions{
		Width:  float32(width),
		Height: float32(height),
	}
}

func ClayRender(renderer *sdl.Renderer, renderCommands clay.RenderCommandArray, fonts []Font) error {
	for renderCommand := range renderCommands.Iter() {
		boundingBox := renderCommand.BoundingBox
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			color := config.BackgroundColor
			if err := renderer.SetDrawColor(uint8(color.R), uint8(color.G), uint8(color.B), uint8(color.A)); err != nil {
				return err
			}
			rect := sdl.FRect{
				X: boundingBox.X,
				Y: boundingBox.Y,
				W: boundingBox.Width,
				H: boundingBox.Height,
			}
			if config.CornerRadius.TopLeft > 0 {
				if err := renderFillRoundedRect(renderer, rect, config.CornerRadius.TopLeft, color); err != nil {
					return err
				}
			} else {
				if err := renderer.FillRectF(&rect); err != nil {
					return err
				}
			}
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			cloned := strings.Clone(config.StringContents.String())
			font := fonts[config.FontId].Font
			surface, err := font.RenderUTF8Blended(cloned, sdl.Color{
				R: uint8(config.TextColor.R),
				G: uint8(config.TextColor.G),
				B: uint8(config.TextColor.B),
				A: uint8(config.TextColor.A),
			})
			if err != nil {
				return err
			}
			texture, err := renderer.CreateTextureFromSurface(surface)
			if err != nil {
				return err
			}
			destination := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if err := renderer.Copy(texture, nil, &destination); err != nil {
				return err
			}
			if err := texture.Destroy(); err != nil {
				return err
			}
			surface.Free()
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
			texture, err := renderer.CreateTextureFromSurface((*sdl.Surface)(config.ImageData.(unsafe.Pointer)))
			if err != nil {
				return err
			}
			destination := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if err := renderer.Copy(texture, nil, &destination); err != nil {
				return err
			}
			if err := texture.Destroy(); err != nil {
				return err
			}
		case clay.RENDER_COMMAND_TYPE_BORDER:
			config := &renderCommand.RenderData.Border
			if err := renderer.SetDrawColor(uint8(config.Color.R), uint8(config.Color.G), uint8(config.Color.B), uint8(config.Color.A)); err != nil {
				return err
			}

			if boundingBox.Width > 0 && boundingBox.Height > 0 {
				maxRadius := min(boundingBox.Width, boundingBox.Height) / 2.0

				if config.Width.Left > 0 {
					clampedRadiusTop := min(config.CornerRadius.TopLeft, maxRadius)
					clampedRadiusBottom := min(config.CornerRadius.BottomLeft, maxRadius)
					rect := sdl.FRect{
						X: boundingBox.X,
						Y: boundingBox.Y + clampedRadiusTop,
						W: float32(config.Width.Left),
						H: boundingBox.Height - clampedRadiusTop - clampedRadiusBottom,
					}
					if err := renderer.FillRectF(&rect); err != nil {
						return err
					}
				}

				if config.Width.Right > 0 {
					clampedRadiusTop := min(config.CornerRadius.TopRight, maxRadius)
					clampedRadiusBottom := min(config.CornerRadius.BottomRight, maxRadius)
					rect := sdl.FRect{
						X: boundingBox.X + boundingBox.Width - float32(config.Width.Right),
						Y: boundingBox.Y + clampedRadiusTop,
						W: float32(config.Width.Right),
						H: boundingBox.Height - clampedRadiusTop - clampedRadiusBottom,
					}
					if err := renderer.FillRectF(&rect); err != nil {
						return err
					}
				}

				if config.Width.Top > 0 {
					clampedRadiusLeft := min(config.CornerRadius.TopLeft, maxRadius)
					clampedRadiusRight := min(config.CornerRadius.TopRight, maxRadius)
					rect := sdl.FRect{
						X: boundingBox.X + clampedRadiusLeft,
						Y: boundingBox.Y,
						W: boundingBox.Width - clampedRadiusLeft - clampedRadiusRight,
						H: float32(config.Width.Top),
					}
					if err := renderer.FillRectF(&rect); err != nil {
						return err
					}
				}

				if config.Width.Bottom > 0 {
					clampedRadiusLeft := min(config.CornerRadius.BottomLeft, maxRadius)
					clampedRadiusRight := min(config.CornerRadius.BottomRight, maxRadius)
					rect := sdl.FRect{
						X: boundingBox.X + clampedRadiusLeft,
						Y: boundingBox.Y + boundingBox.Height - float32(config.Width.Bottom),
						W: boundingBox.Width - clampedRadiusLeft - clampedRadiusRight,
						H: float32(config.Width.Bottom),
					}
					if err := renderer.FillRectF(&rect); err != nil {
						return err
					}
				}

				// corner index: 0->3 topLeft -> CW -> bottonLeft
				if config.Width.Top > 0 && config.CornerRadius.TopLeft > 0 {
					renderCornerBorder(renderer, &boundingBox, config, 0, config.Color)
				}
				if config.Width.Top > 0 && config.CornerRadius.TopRight > 0 {
					renderCornerBorder(renderer, &boundingBox, config, 1, config.Color)
				}
				if config.Width.Bottom > 0 && config.CornerRadius.BottomLeft > 0 {
					renderCornerBorder(renderer, &boundingBox, config, 2, config.Color)
				}
				if config.Width.Bottom > 0 && config.CornerRadius.BottomLeft > 0 {
					renderCornerBorder(renderer, &boundingBox, config, 3, config.Color)
				}
			}
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
	color := sdl.Color{
		R: uint8(_color.R),
		G: uint8(_color.G),
		B: uint8(_color.B),
		A: uint8(_color.A),
	}

	indexCount, vertexCount := 0, int32(0)

	maxRadius := min(rect.W, rect.H) / 2.0
	clampedRadius := min(cornerRadius, maxRadius)

	numCircleSegments := int(max(numCircleSegments, clampedRadius*0.5)) // check if needs to be clamped

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
				break // Top-left
			case 1:
				cx = rect.X + rect.W - clampedRadius
				cy = rect.Y + clampedRadius
				signX = 1
				signY = -1
				break // Top-right
			case 2:
				cx = rect.X + rect.W - clampedRadius
				cy = rect.Y + rect.H - clampedRadius
				signX = 1
				signY = 1
				break // Bottom-right
			case 3:
				cx = rect.X + clampedRadius
				cy = rect.Y + rect.H - clampedRadius
				signX = -1
				signY = 1
				break // Bottom-left
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

// all rendering is performed by a single SDL call, using twi sets of arcing triangles, inner and outer, that fit together; along with two tringles to fill the end gaps.
func renderCornerBorder(renderer *sdl.Renderer, boundingBox *clay.BoundingBox, config *clay.BorderRenderData, cornerIndex int, _color clay.Color) {
	/////////////////////////////////
	//The arc is constructed of outer triangles and inner triangles (if needed).
	//First three vertices are first outer triangle's vertices
	//Each two vertices after that are the inner-middle and second-outer vertex of
	//each outer triangle after the first, because there first-outer vertex is equal to the
	//second-outer vertex of the previous triangle. Indices set accordingly.
	//The final two vertices are the missing vertices for the first and last inner triangles (if needed)
	//Everything is in clockwise order (CW).
	/////////////////////////////////

	color := sdl.Color{
		R: uint8(_color.R),
		G: uint8(_color.G),
		B: uint8(_color.B),
		A: uint8(_color.A),
	}

	var centerX, centerY, outerRadius, startAngle, borderWidth float32
	maxRadius := min(boundingBox.Width, boundingBox.Height) / 2.0

	var vertices [512]sdl.Vertex
	var indices [512]int32
	indexCount, vertexCount := int32(0), int32(0)

	switch cornerIndex {
	case 0:
		startAngle = math.Pi
		outerRadius = min(config.CornerRadius.TopLeft, maxRadius)
		centerX = boundingBox.X + outerRadius
		centerY = boundingBox.Y + outerRadius
		borderWidth = float32(config.Width.Top)
	case 1:
		startAngle = 3 * math.Pi / 2
		outerRadius = min(config.CornerRadius.TopRight, maxRadius)
		centerX = boundingBox.X + boundingBox.Width - outerRadius
		centerY = boundingBox.Y + outerRadius
		borderWidth = float32(config.Width.Top)
	case 2:
		startAngle = 0
		outerRadius = min(config.CornerRadius.BottomRight, maxRadius)
		centerX = boundingBox.X + boundingBox.Width - outerRadius
		centerY = boundingBox.Y + boundingBox.Height - outerRadius
		borderWidth = float32(config.Width.Bottom)
	case 3:
		startAngle = math.Pi / 2
		outerRadius = min(config.CornerRadius.BottomLeft, maxRadius)
		centerX = boundingBox.X + outerRadius
		centerY = boundingBox.Y + boundingBox.Height - outerRadius
		borderWidth = float32(config.Width.Bottom)
	default:
		panic("invalid corner index")
	}

	innerRadius := outerRadius - borderWidth
	minNumOuterTriangles := numCircleSegments
	numOuterTriangles := max(minNumOuterTriangles, int(math.Ceil(float64(outerRadius*0.5))))
	angleStep := math.Pi / (2.0 * float32(numOuterTriangles))

	// outer triangles, in CW order
	for i := 0; i < numOuterTriangles; i++ {
		angle1 := startAngle + float32(i)*angleStep       // first-outer vertex angle
		angle2 := startAngle + (float32(i)+0.5)*angleStep // inner-middle vertex angle
		angle3 := startAngle + float32(i+1)*angleStep     // second-outer vertex angle

		if i == 0 { // first outer triangle
			vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: centerX + float32(math.Cos(float64(angle1)))*outerRadius, Y: centerY + float32(math.Sin(float64(angle1)))*outerRadius}, Color: color} // vertex index = 0
			vertexCount++
		}
		indices[indexCount] = vertexCount - 1 // will be second-outer vertex of last outer triangle if not first outer triangle.
		indexCount++
		if innerRadius > 0 {
			vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: centerX + float32(math.Cos(float64(angle2)))*(innerRadius), Y: centerY + float32(math.Sin(float64(angle2)))*(innerRadius)}, Color: color}
			vertexCount++
		} else {
			vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{X: centerX, Y: centerY}, Color: color}
			vertexCount++
		}

		indices[indexCount] = vertexCount - 1
		indexCount++
		vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{centerX + float32(math.Cos(float64(angle3)))*outerRadius, centerY + float32(math.Sin(float64(angle3)))*outerRadius}, Color: color}
		vertexCount++
		indices[indexCount] = vertexCount - 1
		indexCount++
	}

	if innerRadius > 0 {
		// inner triangles in CW order (except the first and last)
		for i := 0; i < numOuterTriangles-1; i++ { // skip the last outer triangle
			if i == 0 { // first outer triangle -> second inner triangle
				indices[indexCount] = 1 // inner-middle vertex of first outer triangle
				indexCount++
				indices[indexCount] = 2 // second-outer vertex of first outer triangle
				indexCount++
				indices[indexCount] = 3 // innder-middle vertex of second-outer triangle
				indexCount++
			} else {
				baseIndex := 3                                   // skip first outer triangle
				indices[indexCount] = int32(baseIndex + (i-1)*2) // inner-middle vertex of current outer triangle
				indexCount++
				indices[indexCount] = int32(baseIndex + (i-1)*2 + 1) // second-outer vertex of current outer triangle
				indexCount++
				indices[indexCount] = int32(baseIndex + (i-1)*2 + 2) // inner-middle vertex of next outer triangle
				indexCount++
			}
		}

		endAngle := startAngle + math.Pi/2.0

		// last inner triangle
		indices[indexCount] = vertexCount - 2 // inner-middle vertex of last outer triangle
		indexCount++
		indices[indexCount] = vertexCount - 1 // second-outer vertex of last outer triangle
		indexCount++
		vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{centerX + float32(math.Cos(float64(endAngle)))*innerRadius, centerY + float32(math.Sin(float64(endAngle)))*innerRadius}, Color: color} // missing vertex
		vertexCount++
		indices[indexCount] = vertexCount - 1
		indexCount++

		// //first inner triangle
		indices[indexCount] = 0 // first-outer vertex of first outer triangle
		indexCount++
		indices[indexCount] = 1 // inner-middle vertex of first outer triangle
		indexCount++
		vertices[vertexCount] = sdl.Vertex{Position: sdl.FPoint{centerX + float32(math.Cos(float64(startAngle)))*innerRadius, centerY + float32(math.Sin(float64(startAngle)))*innerRadius}, Color: color} // missing vertex
		vertexCount++
		indices[indexCount] = vertexCount - 1
		indexCount++
	}

	err := renderer.RenderGeometry(nil, vertices[:vertexCount], indices[:indexCount])
	if err != nil {
		return
	}
}
