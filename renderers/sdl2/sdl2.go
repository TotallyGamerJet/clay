package sdl2

import (
	"unsafe"

	"github.com/totallygamerjet/clay"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	FontId uint32
	Font   *ttf.Font
}

func MeasureText(text clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	//fonts := (*Font)(userData)
	//font := unsafe.Slice(fonts, 1)[config.FontId].Font // TODO: idk if this the best way to create the slice
	//chars := text.String()

	//width, height, err := font.SizeUTF8(chars)
	//if err != nil {
	//	panic(err)
	//}

	return clay.Dimensions{
		Width:  float32(0),
		Height: float32(0),
	}
}

func ClayRender(renderer *sdl.Renderer, renderCommands clay.RenderCommandArray, fonts []Font) {
	for i := int32(0); i < renderCommands.Length; i++ {
		renderCommand := clay.RenderCommandArray_Get(&renderCommands, i)
		boundingBox := renderCommand.BoundingBox
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			color := config.BackgroundColor
			if err := renderer.SetDrawColor(uint8(color.R), uint8(color.G), uint8(color.B), uint8(color.A)); err != nil {
				panic(err)
			}
			rect := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if config.CornerRadius.TopLeft > 0 {
				// TODO: SDL_RenderFillRoundedRect(renderer, rect, config->cornerRadius.topLeft, color);
				if err := renderer.FillRect(&rect); err != nil {
					panic(err)
				}
			} else {
				if err := renderer.FillRect(&rect); err != nil {
					panic(err)
				}
			}
		case clay.RENDER_COMMAND_TYPE_TEXT:
			config := &renderCommand.RenderData.Text
			cloned := config.StringContents.String()
			font := fonts[config.FontId].Font
			surface, err := font.RenderUTF8Blended(cloned, sdl.Color{
				R: uint8(config.TextColor.R),
				G: uint8(config.TextColor.G),
				B: uint8(config.TextColor.B),
				A: uint8(config.TextColor.A),
			})
			if err != nil {
				panic(err)
			}
			texture, err := renderer.CreateTextureFromSurface(surface)
			if err != nil {
				panic(err)
			}
			destination := sdl.Rect{
				X: int32(boundingBox.X),
				Y: int32(boundingBox.Y),
				W: int32(boundingBox.Width),
				H: int32(boundingBox.Height),
			}
			if err := renderer.Copy(texture, nil, &destination); err != nil {
				panic(err)
			}
			if err := texture.Destroy(); err != nil {
				panic(err)
			}
			surface.Free()
		case clay.RENDER_COMMAND_TYPE_NONE:
		case clay.RENDER_COMMAND_TYPE_BORDER:
		case clay.RENDER_COMMAND_TYPE_IMAGE:
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
		case clay.RENDER_COMMAND_TYPE_CUSTOM:
		}
	}
	//for (uint32_t i = 0; i < renderCommands.length; i++)
	//    {
	//        Clay_RenderCommand *renderCommand = Clay_RenderCommandArray_Get(&renderCommands, i);
	//        Clay_BoundingBox boundingBox = renderCommand->boundingBox;
	//        switch (renderCommand->commandType)
	//        {
	//            case CLAY_RENDER_COMMAND_TYPE_TEXT: {
	//                Clay_TextRenderData *config = &renderCommand->renderData.text;
	//                char *cloned = (char *)calloc(config->stringContents.length + 1, 1);
	//                memcpy(cloned, config->stringContents.chars, config->stringContents.length);
	//                TTF_Font* font = fonts[config->fontId].font;
	//                SDL_Surface *surface = TTF_RenderUTF8_Blended(font, cloned, (SDL_Color) {
	//                        .r = (Uint8)config->textColor.r,
	//                        .g = (Uint8)config->textColor.g,
	//                        .b = (Uint8)config->textColor.b,
	//                        .a = (Uint8)config->textColor.a,
	//                });
	//                SDL_Texture *texture = SDL_CreateTextureFromSurface(renderer, surface);
	//
	//                SDL_Rect destination = (SDL_Rect){
	//                        .x = boundingBox.x,
	//                        .y = boundingBox.y,
	//                        .w = boundingBox.width,
	//                        .h = boundingBox.height,
	//                };
	//                SDL_RenderCopy(renderer, texture, NULL, &destination);
	//
	//                SDL_DestroyTexture(texture);
	//                SDL_FreeSurface(surface);
	//                free(cloned);
	//                break;
	//            }
	//            case CLAY_RENDER_COMMAND_TYPE_SCISSOR_START: {
	//                currentClippingRectangle = (SDL_Rect) {
	//                        .x = boundingBox.x,
	//                        .y = boundingBox.y,
	//                        .w = boundingBox.width,
	//                        .h = boundingBox.height,
	//                };
	//                SDL_RenderSetClipRect(renderer, &currentClippingRectangle);
	//                break;
	//            }
	//            case CLAY_RENDER_COMMAND_TYPE_SCISSOR_END: {
	//                SDL_RenderSetClipRect(renderer, NULL);
	//                break;
	//            }
	//            case CLAY_RENDER_COMMAND_TYPE_IMAGE: {
	//                Clay_ImageRenderData *config = &renderCommand->renderData.image;
	//
	//                SDL_Texture *texture = SDL_CreateTextureFromSurface(renderer, config->imageData);
	//
	//                SDL_Rect destination = (SDL_Rect){
	//                    .x = boundingBox.x,
	//                    .y = boundingBox.y,
	//                    .w = boundingBox.width,
	//                    .h = boundingBox.height,
	//                };
	//
	//                SDL_RenderCopy(renderer, texture, NULL, &destination);
	//
	//                SDL_DestroyTexture(texture);
	//                break;
	//            }
	//            case CLAY_RENDER_COMMAND_TYPE_BORDER: {
	//                Clay_BorderRenderData *config = &renderCommand->renderData.border;
	//                SDL_SetRenderDrawColor(renderer, CLAY_COLOR_TO_SDL_COLOR_ARGS(config->color));
	//
	//                if(boundingBox.width > 0 & boundingBox.height > 0){
	//                    const float maxRadius = SDL_min(boundingBox.width, boundingBox.height) / 2.0f;
	//
	//                    if (config->width.left > 0) {
	//                        const float clampedRadiusTop = SDL_min((float)config->cornerRadius.topLeft, maxRadius);
	//                        const float clampedRadiusBottom = SDL_min((float)config->cornerRadius.bottomLeft, maxRadius);
	//                        SDL_FRect rect = {
	//                            boundingBox.x,
	//                            boundingBox.y + clampedRadiusTop,
	//                            (float)config->width.left,
	//                            (float)boundingBox.height - clampedRadiusTop - clampedRadiusBottom
	//                        };
	//                        SDL_RenderFillRectF(renderer, &rect);
	//                    }
	//
	//                    if (config->width.right > 0) {
	//                        const float clampedRadiusTop = SDL_min((float)config->cornerRadius.topRight, maxRadius);
	//                        const float clampedRadiusBottom = SDL_min((float)config->cornerRadius.bottomRight, maxRadius);
	//                        SDL_FRect rect = {
	//                            boundingBox.x + boundingBox.width - config->width.right,
	//                            boundingBox.y + clampedRadiusTop,
	//                            (float)config->width.right,
	//                            (float)boundingBox.height - clampedRadiusTop - clampedRadiusBottom
	//                        };
	//                        SDL_RenderFillRectF(renderer, &rect);
	//                    }
	//
	//                    if (config->width.top > 0) {
	//                        const float clampedRadiusLeft = SDL_min((float)config->cornerRadius.topLeft, maxRadius);
	//                        const float clampedRadiusRight = SDL_min((float)config->cornerRadius.topRight, maxRadius);
	//                        SDL_FRect rect = {
	//                            boundingBox.x + clampedRadiusLeft,
	//                            boundingBox.y,
	//                            boundingBox.width - clampedRadiusLeft - clampedRadiusRight,
	//                            (float)config->width.top };
	//                        SDL_RenderFillRectF(renderer, &rect);
	//                    }
	//
	//                    if (config->width.bottom > 0) {
	//                        const float clampedRadiusLeft = SDL_min((float)config->cornerRadius.bottomLeft, maxRadius);
	//                        const float clampedRadiusRight = SDL_min((float)config->cornerRadius.bottomRight, maxRadius);
	//                        SDL_FRect rect = {
	//                            boundingBox.x + clampedRadiusLeft,
	//                            boundingBox.y + boundingBox.height - config->width.bottom,
	//                            boundingBox.width - clampedRadiusLeft - clampedRadiusRight,
	//                            (float)config->width.bottom
	//                        };
	//                        SDL_RenderFillRectF(renderer, &rect);
	//                    }
	//
	//                    //corner index: 0->3 topLeft -> CW -> bottonLeft
	//                    if (config->width.top > 0 & config->cornerRadius.topLeft > 0) {
	//                        SDL_RenderCornerBorder(renderer, &boundingBox, config, 0, config->color);
	//                    }
	//
	//                    if (config->width.top > 0 & config->cornerRadius.topRight> 0) {
	//                        SDL_RenderCornerBorder(renderer, &boundingBox, config, 1, config->color);
	//                    }
	//
	//                    if (config->width.bottom > 0 & config->cornerRadius.bottomLeft > 0) {
	//                        SDL_RenderCornerBorder(renderer, &boundingBox, config, 2, config->color);
	//                    }
	//
	//                    if (config->width.bottom > 0 & config->cornerRadius.bottomLeft > 0) {
	//                        SDL_RenderCornerBorder(renderer, &boundingBox, config, 3, config->color);
	//                    }
	//                }
	//
	//                break;
	//            }
	//            default: {
	//                fprintf(stderr, "Error: unhandled render command: %d\n", renderCommand->commandType);
	//                exit(1);
	//            }
	//        }
	//    }
}
