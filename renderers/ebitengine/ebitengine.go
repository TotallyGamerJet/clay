package ebitengine

import (
	"fmt"
	"image"
	"image/color"
	"log/slog"
	"math"
	"strings"
	"unsafe"

	"github.com/TotallyGamerJet/clay"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var whiteImage *ebiten.Image

func init() {
	// Creating a sub-image to avoid bleeding edges
	// https://github.com/hajimehoshi/ebiten/blob/1a4237213c92be1b9c16176887d992eb4183751b/vector/util.go#L26-L29
	img := ebiten.NewImage(3, 3)
	img.Fill(color.White)
	whiteImage = img.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
}

func MeasureText(txt clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	fonts := *(*[]text.Face)(userData)

	font := fonts[config.FontId]

	width, height := text.Measure(txt.String(), font, font.Metrics().HLineGap)
	scaleFactor := ebiten.Monitor().DeviceScaleFactor() // should we be passing the scaleFactor like we do in the renderer?
	return clay.Dimensions{
		Width:  float32(width / scaleFactor),
		Height: float32(height / scaleFactor),
	}
}

func ClayRender(screen *ebiten.Image, scaleFactor float32, renderCommands clay.RenderCommandArray, fonts []text.Face) error {
	fullScreen := screen
	for renderCommand := range renderCommands.Iter() {
		boundingBox := renderCommand.BoundingBox
		boundingBox.X *= scaleFactor
		boundingBox.Y *= scaleFactor
		boundingBox.Width *= scaleFactor
		boundingBox.Height *= scaleFactor
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			if config.CornerRadius.TopLeft > 0 {
				cornerRadius := config.CornerRadius.TopLeft * scaleFactor
				if err := renderFillRoundedRect(screen, boundingBox, cornerRadius, config.BackgroundColor); err != nil {
					return err
				}
			} else {
				vector.DrawFilledRect(
					screen,
					boundingBox.X,
					boundingBox.Y,
					boundingBox.Width, boundingBox.Height,
					color.RGBA{
						R: uint8(config.BackgroundColor.R),
						G: uint8(config.BackgroundColor.G),
						B: uint8(config.BackgroundColor.B),
						A: uint8(config.BackgroundColor.A),
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
			screen = screen.SubImage(image.Rect(
				int(boundingBox.X), int(boundingBox.Y),
				int(boundingBox.X+boundingBox.Width),
				int(boundingBox.Y+boundingBox.Height),
			)).(*ebiten.Image)
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
			screen = fullScreen
		case clay.RENDER_COMMAND_TYPE_IMAGE:
			config := &renderCommand.RenderData.Image
			img := (*ebiten.Image)(config.ImageData)
			opts := &ebiten.DrawImageOptions{}
			bounds := img.Bounds()
			opts.GeoM.Scale(float64(boundingBox.Width/float32(bounds.Dx())), float64(boundingBox.Height/float32(bounds.Dy())))
			opts.GeoM.Translate(float64(boundingBox.X), float64(boundingBox.Y))
			screen.DrawImage(img, opts)
		case clay.RENDER_COMMAND_TYPE_BORDER:
			config := &renderCommand.RenderData.Border
			config.Width.Top = uint16(float32(config.Width.Top) * scaleFactor)
			config.Width.Bottom = uint16(float32(config.Width.Bottom) * scaleFactor)
			config.Width.Left = uint16(float32(config.Width.Left) * scaleFactor)
			config.Width.Right = uint16(float32(config.Width.Right) * scaleFactor)

			config.CornerRadius.TopLeft *= scaleFactor
			config.CornerRadius.BottomLeft *= scaleFactor
			config.CornerRadius.TopRight *= scaleFactor
			config.CornerRadius.BottomRight *= scaleFactor
			if boundingBox.Width > 0 && boundingBox.Height > 0 {
				maxRadius := min(boundingBox.Width, boundingBox.Height) / 2.0

				if config.Width.Left > 0 {
					clampedRadiusTop := min(config.CornerRadius.TopLeft, maxRadius)
					clampedRadiusBottom := min(config.CornerRadius.BottomLeft, maxRadius)
					vector.DrawFilledRect(
						screen,
						boundingBox.X,
						boundingBox.Y+clampedRadiusTop,
						float32(config.Width.Left), boundingBox.Height-clampedRadiusTop-clampedRadiusBottom,
						color.RGBA{
							R: uint8(config.Color.R),
							G: uint8(config.Color.G),
							B: uint8(config.Color.B),
							A: uint8(config.Color.A),
						}, true,
					)
				}

				if config.Width.Right > 0 {
					clampedRadiusTop := min(config.CornerRadius.TopRight, maxRadius)
					clampedRadiusBottom := min(config.CornerRadius.BottomRight, maxRadius)
					vector.DrawFilledRect(
						screen,
						boundingBox.X+boundingBox.Width-float32(config.Width.Right),
						boundingBox.Y+clampedRadiusTop,
						float32(config.Width.Right),
						boundingBox.Height-clampedRadiusTop-clampedRadiusBottom,
						color.RGBA{
							R: uint8(config.Color.R),
							G: uint8(config.Color.G),
							B: uint8(config.Color.B),
							A: uint8(config.Color.A),
						}, true,
					)
				}

				if config.Width.Top > 0 {
					clampedRadiusLeft := min(config.CornerRadius.TopLeft, maxRadius)
					clampedRadiusRight := min(config.CornerRadius.TopRight, maxRadius)
					vector.DrawFilledRect(
						screen,
						boundingBox.X+clampedRadiusLeft,
						boundingBox.Y,
						boundingBox.Width-clampedRadiusLeft-clampedRadiusRight,
						float32(config.Width.Top),
						color.RGBA{
							R: uint8(config.Color.R),
							G: uint8(config.Color.G),
							B: uint8(config.Color.B),
							A: uint8(config.Color.A),
						}, true,
					)
				}

				if config.Width.Bottom > 0 {
					clampedRadiusLeft := min(config.CornerRadius.BottomLeft, maxRadius)
					clampedRadiusRight := min(config.CornerRadius.BottomRight, maxRadius)
					vector.DrawFilledRect(
						screen,
						boundingBox.X+clampedRadiusLeft,
						boundingBox.Y+boundingBox.Height-float32(config.Width.Bottom),
						boundingBox.Width-clampedRadiusLeft-clampedRadiusRight,
						float32(config.Width.Bottom),
						color.RGBA{
							R: uint8(config.Color.R),
							G: uint8(config.Color.G),
							B: uint8(config.Color.B),
							A: uint8(config.Color.A),
						}, true,
					)
				}

				// corner index: 0->3 topLeft -> CW -> bottonLeft
				if config.Width.Top > 0 && config.CornerRadius.TopLeft > 0 {
					renderCornerBorder(screen, &boundingBox, config, 0, config.Color)
				}
				if config.Width.Top > 0 && config.CornerRadius.TopRight > 0 {
					renderCornerBorder(screen, &boundingBox, config, 1, config.Color)
				}
				if config.Width.Bottom > 0 && config.CornerRadius.BottomLeft > 0 {
					renderCornerBorder(screen, &boundingBox, config, 2, config.Color)
				}
				if config.Width.Bottom > 0 && config.CornerRadius.BottomLeft > 0 {
					renderCornerBorder(screen, &boundingBox, config, 3, config.Color)
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

func renderFillRoundedRect(screen *ebiten.Image, rect clay.BoundingBox, cornerRadius float32, _color clay.Color) error {
	r := _color.R / 255
	g := _color.G / 255
	b := _color.B / 255
	a := _color.A / 255

	indexCount, vertexCount := 0, uint16(0)

	minRadius := min(rect.Width, rect.Height) / 2.0
	clampedRadius := min(cornerRadius, minRadius)

	numCircleSegments := max(numCircleSegments, int(clampedRadius*0.5)) // check if it needs to be clamped

	var vertices [512]ebiten.Vertex
	var indices [512]uint16

	// define center rectangle
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

	// define rounded corners as triangle fans
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
	// Define edge rectangles
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
	screen.DrawTriangles(vertices[:vertexCount], indices[:indexCount], whiteImage, &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	})

	return nil
}

// all rendering is performed by a single ebitengine call, using two sets of arcing triangles, inner and outer, that fit together; along with two triangles to fill the end gaps.
func renderCornerBorder(screen *ebiten.Image, boundingBox *clay.BoundingBox, config *clay.BorderRenderData, cornerIndex int, _color clay.Color) {
	/////////////////////////////////
	//The arc is constructed of outer triangles and inner triangles (if needed).
	//First three vertices are first outer triangle's vertices
	//Each two vertices after that are the inner-middle and second-outer vertex of
	//each outer triangle after the first, because there first-outer vertex is equal to the
	//second-outer vertex of the previous triangle. Indices set accordingly.
	//The final two vertices are the missing vertices for the first and last inner triangles (if needed)
	//Everything is in clockwise order (CW).
	/////////////////////////////////

	r := _color.R / 255
	g := _color.G / 255
	b := _color.B / 255
	a := _color.A / 255

	var centerX, centerY, outerRadius, startAngle, borderWidth float32
	maxRadius := min(boundingBox.Width, boundingBox.Height) / 2.0

	var vertices [512]ebiten.Vertex
	var indices [512]uint16
	indexCount, vertexCount := uint16(0), uint16(0)

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
			vertices[vertexCount] = ebiten.Vertex{
				DstX:   centerX + float32(math.Cos(float64(angle1)))*outerRadius,
				DstY:   centerY + float32(math.Sin(float64(angle1)))*outerRadius,
				SrcX:   0,
				SrcY:   0,
				ColorR: r,
				ColorG: g,
				ColorB: b,
				ColorA: a,
			} // vertex index = 0
			vertexCount++
		}
		indices[indexCount] = vertexCount - 1 // will be second-outer vertex of last outer triangle if not first outer triangle.
		indexCount++
		if innerRadius > 0 {
			vertices[vertexCount] = ebiten.Vertex{
				DstX:   centerX + float32(math.Cos(float64(angle2)))*(innerRadius),
				DstY:   centerY + float32(math.Sin(float64(angle2)))*(innerRadius),
				SrcX:   0,
				SrcY:   0,
				ColorR: r,
				ColorG: g,
				ColorB: b,
				ColorA: a,
			}
			vertexCount++
		} else {
			vertices[vertexCount] = ebiten.Vertex{
				DstX:   centerX,
				DstY:   centerY,
				SrcX:   0,
				SrcY:   0,
				ColorR: r,
				ColorG: g,
				ColorB: b,
				ColorA: a,
			}
			vertexCount++
		}

		indices[indexCount] = vertexCount - 1
		indexCount++
		vertices[vertexCount] = ebiten.Vertex{
			DstX:   centerX + float32(math.Cos(float64(angle3)))*outerRadius,
			DstY:   centerY + float32(math.Sin(float64(angle3)))*outerRadius,
			SrcX:   0,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a}
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
				baseIndex := 3                                    // skip first outer triangle
				indices[indexCount] = uint16(baseIndex + (i-1)*2) // inner-middle vertex of current outer triangle
				indexCount++
				indices[indexCount] = uint16(baseIndex + (i-1)*2 + 1) // second-outer vertex of current outer triangle
				indexCount++
				indices[indexCount] = uint16(baseIndex + (i-1)*2 + 2) // inner-middle vertex of next outer triangle
				indexCount++
			}
		}

		endAngle := startAngle + math.Pi/2.0

		// last inner triangle
		indices[indexCount] = vertexCount - 2 // inner-middle vertex of last outer triangle
		indexCount++
		indices[indexCount] = vertexCount - 1 // second-outer vertex of last outer triangle
		indexCount++
		vertices[vertexCount] = ebiten.Vertex{
			DstX:   centerX + float32(math.Cos(float64(endAngle)))*innerRadius,
			DstY:   centerY + float32(math.Sin(float64(endAngle)))*innerRadius,
			SrcX:   0,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		} // missing vertex
		vertexCount++
		indices[indexCount] = vertexCount - 1
		indexCount++

		// //first inner triangle
		indices[indexCount] = 0 // first-outer vertex of first outer triangle
		indexCount++
		indices[indexCount] = 1 // inner-middle vertex of first outer triangle
		indexCount++
		vertices[vertexCount] = ebiten.Vertex{
			DstX:   centerX + float32(math.Cos(float64(startAngle)))*innerRadius,
			DstY:   centerY + float32(math.Sin(float64(startAngle)))*innerRadius,
			SrcX:   0,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		} // missing vertex
		vertexCount++
		indices[indexCount] = vertexCount - 1
		indexCount++
	}

	screen.DrawTriangles(vertices[:vertexCount], indices[:indexCount], whiteImage, &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	})
}
