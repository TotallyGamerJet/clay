package software

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log/slog"
	"math"
	"strings"
	"unsafe"

	"github.com/TotallyGamerJet/clay"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// PrintPlaygroundImage prints the image so that it displays in the Go playground. It may have to shrink the image
// if it is too big for the playground.
func PrintPlaygroundImage(m image.Image) {
	const maxPixels = math.MaxUint16
	origWidth := m.Bounds().Dx()
	origHeight := m.Bounds().Dy()
	ratio := float64(origWidth) / float64(origHeight)
	x := math.Sqrt(float64(maxPixels) / ratio)
	newHeight := int(math.Floor(x))
	newWidth := int(math.Floor(ratio * x))
	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.BiLinear.Scale(newImg, newImg.Bounds(), m, m.Bounds(), draw.Over, nil)

	var buf bytes.Buffer
	err := png.Encode(&buf, newImg)
	if err != nil {
		panic(err)
	}
	fmt.Println("IMAGE:" + base64.StdEncoding.EncodeToString(buf.Bytes()))
}

func MeasureText(txt clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	fonts := *(*[]font.Face)(userData)
	face := fonts[config.FontId]
	width := font.MeasureString(face, txt.String()).Ceil()
	height := face.Metrics().Height.Ceil()
	return clay.Dimensions{
		Width: float32(width), Height: float32(height),
	}
}

func ClayRender(screen draw.Image, renderCommands clay.RenderCommandArray, fonts []font.Face) error {
	fullScreen := screen
	for renderCommand := range renderCommands.Iter() {
		boundingBox := renderCommand.BoundingBox
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			c := color.RGBA{
				R: uint8(config.BackgroundColor.R),
				G: uint8(config.BackgroundColor.G),
				B: uint8(config.BackgroundColor.B),
				A: uint8(config.BackgroundColor.A),
			}
			rect := image.Rect(int(boundingBox.X), int(boundingBox.Y), int(boundingBox.X+boundingBox.Width), int(boundingBox.Y+boundingBox.Height))
			if config.CornerRadius.TopLeft > 0 {
				draw.Draw(screen, rect, &image.Uniform{C: c}, image.Point{}, draw.Src)
			} else {
				draw.Draw(screen, rect, &image.Uniform{C: c}, image.Point{}, draw.Src)
			}
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			cloned := strings.Clone(config.StringContents.String())
			face := fonts[config.FontId]

			c := color.RGBA{
				R: uint8(config.TextColor.R),
				G: uint8(config.TextColor.G),
				B: uint8(config.TextColor.B),
				A: uint8(config.TextColor.A),
			}
			d := &font.Drawer{
				Dst:  screen,
				Src:  image.NewUniform(c),
				Face: face,
				Dot:  fixed.Point26_6{X: fixed.I(int(boundingBox.X)), Y: fixed.I(int(boundingBox.Y)) + face.Metrics().Ascent},
			}
			d.DrawString(cloned)
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
			rect := image.Rect(int(boundingBox.X), int(boundingBox.Y), int(boundingBox.X+boundingBox.Width), int(boundingBox.Y+boundingBox.Height))
			screen = screen.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(rect).(draw.Image)
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
			screen = fullScreen
		case clay.RENDER_COMMAND_TYPE_IMAGE:
			config := &renderCommand.RenderData.Image
			img := (*image.Image)(config.ImageData)
			if img == nil {
				continue
			}
			destRect := image.Rect(int(boundingBox.X), int(boundingBox.Y), int(boundingBox.X+boundingBox.Width), int(boundingBox.Y+boundingBox.Height))
			draw.ApproxBiLinear.Scale(screen, destRect, *img, (*img).Bounds(), draw.Over, nil)
		case clay.RENDER_COMMAND_TYPE_BORDER:
			panic("not implemented")
		case clay.RENDER_COMMAND_TYPE_NONE:
		case clay.RENDER_COMMAND_TYPE_CUSTOM:
		default:
			slog.Warn("Unknown command type", "type", renderCommand.CommandType)
		}
	}

	return nil
}
