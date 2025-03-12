package raylib

import (
	"strings"
	"unsafe"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/totallygamerjet/clay"
)

// #define CLAY_RECTANGLE_TO_RAYLIB_RECTANGLE(rectangle) (Rectangle) { .x = rectangle.x, .y = rectangle.y, .width = rectangle.width, .height = rectangle.height }
// #define CLAY_COLOR_TO_RAYLIB_COLOR(color) (Color) { .r = (unsigned char)roundf(color.r), .g = (unsigned char)roundf(color.g), .b = (unsigned char)roundf(color.b), .a = (unsigned char)roundf(color.a) }
//
// Camera Raylib_camera;
//
// typedef enum
//
//	{
//	   CUSTOM_LAYOUT_ELEMENT_TYPE_3D_MODEL
//	} CustomLayoutElementType;
//
// typedef struct
//
//	{
//	   Model model;
//	   float scale;
//	   Vector3 position;
//	   Matrix rotation;
//	} CustomLayoutElement_3DModel;
//
// typedef struct
//
//	{
//	   CustomLayoutElementType type;
//	   union {
//	       CustomLayoutElement_3DModel model;
//	   } customData;
//	} CustomLayoutElement;
//
// // Get a ray trace from the screen position (i.e mouse) within a specific section of the screen
// Ray GetScreenToWorldPointWithZDistance(Vector2 position, Camera camera, int screenWidth, int screenHeight, float zDistance)
//
//	{
//	   Ray ray = { 0 };
//
//	   // Calculate normalized device coordinates
//	   // NOTE: y value is negative
//	   float x = (2.0f*position.x)/(float)screenWidth - 1.0f;
//	   float y = 1.0f - (2.0f*position.y)/(float)screenHeight;
//	   float z = 1.0f;
//
//	   // Store values in a vector
//	   Vector3 deviceCoords = { x, y, z };
//
//	   // Calculate view matrix from camera look at
//	   Matrix matView = MatrixLookAt(camera.position, camera.target, camera.up);
//
//	   Matrix matProj = MatrixIdentity();
//
//	   if (camera.projection == CAMERA_PERSPECTIVE)
//	   {
//	       // Calculate projection matrix from perspective
//	       matProj = MatrixPerspective(camera.fovy*DEG2RAD, ((double)screenWidth/(double)screenHeight), 0.01f, zDistance);
//	   }
//	   else if (camera.projection == CAMERA_ORTHOGRAPHIC)
//	   {
//	       double aspect = (double)screenWidth/(double)screenHeight;
//	       double top = camera.fovy/2.0;
//	       double right = top*aspect;
//
//	       // Calculate projection matrix from orthographic
//	       matProj = MatrixOrtho(-right, right, -top, top, 0.01, 1000.0);
//	   }
//
//	   // Unproject far/near points
//	   Vector3 nearPoint = Vector3Unproject((Vector3){ deviceCoords.x, deviceCoords.y, 0.0f }, matProj, matView);
//	   Vector3 farPoint = Vector3Unproject((Vector3){ deviceCoords.x, deviceCoords.y, 1.0f }, matProj, matView);
//
//	   // Calculate normalized direction vector
//	   Vector3 direction = Vector3Normalize(Vector3Subtract(farPoint, nearPoint));
//
//	   ray.position = farPoint;
//
//	   // Apply calculated vectors to ray
//	   ray.direction = direction;
//
//	   return ray;
//	}
//
//	static inline Clay_Dimensions Raylib_MeasureText(Clay_StringSlice text, Clay_TextElementConfig *config, void *userData) {
//	   // Measure string size for Font
//	   Clay_Dimensions textSize = { 0 };
//
//	   float maxTextWidth = 0.0f;
//	   float lineTextWidth = 0;
func MeasureText(text clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	//textSize := clay.Dimensions{}
	//
	//maxTextWidth := float32(0)
	//lineTextWidth := float32(0)
	//
	//textHeight := config.FontSize
	//fonts := *(*[]rl.Font)(userData)
	//fontToUse := fonts[config.FontId]
	//
	//// Font failed to load, likely the fonts are in the wrong place relative to the execution dir.
	//// RayLib ships with a default font, so we can continue with that built in one.
	//if fontToUse.Chars == nil {
	//	fontToUse = rl.GetFontDefault()
	//}
	//scaleFactor := float32(config.FontSize) / float32(fontToUse.BaseSize)
	//
	//str := text.String()
	//for i := 0; i < len(str); i++ {
	//	if str[i] == '\n' {
	//		maxTextWidth = max(maxTextWidth, lineTextWidth)
	//		lineTextWidth = 0
	//		continue
	//	}
	//	index := str[i] - 32
	//	chars := unsafe.Slice(fontToUse.Chars, fontToUse.CharsCount)
	//	if chars[index].AdvanceX != 0 {
	//		lineTextWidth += float32(chars[index].AdvanceX)
	//	} else {
	//		lineTextWidth += fontToUse.Recs.Width + fontToUse.
	//	}
	//}
	return clay.Dimensions{}
}

//    float textHeight = config->fontSize;
//    Font* fonts = (Font*)userData;
//    Font fontToUse = fonts[config->fontId];
// Font failed to load, likely the fonts are in the wrong place relative to the execution dir.
// RayLib ships with a default font, so we can continue with that built in one.
//    if (!fontToUse.glyphs) {
//        fontToUse = GetFontDefault();
//    }
//
//    float scaleFactor = config->fontSize/(float)fontToUse.baseSize;
//
//    for (int i = 0; i < text.length; ++i)
//    {
//        if (text.chars[i] == '\n') {
//            maxTextWidth = fmax(maxTextWidth, lineTextWidth);
//            lineTextWidth = 0;
//            continue;
//        }
//        int index = text.chars[i] - 32;
//        if (fontToUse.glyphs[index].advanceX != 0) lineTextWidth += fontToUse.glyphs[index].advanceX;
//        else lineTextWidth += (fontToUse.recs[index].width + fontToUse.glyphs[index].offsetX);
//    }
//
//    maxTextWidth = fmax(maxTextWidth, lineTextWidth);
//
//    textSize.width = maxTextWidth * scaleFactor;
//    textSize.height = textHeight;
//
//    return textSize;
//}

func Initialize(width, height int32, title string, flags uint32) {
	rl.SetConfigFlags(flags)
	rl.InitWindow(width, height, title)
	rl.EnableEventWaiting()
}

func Render(renderCommands clay.RenderCommandArray, fonts []rl.Font) {
	for j := int32(0); j < renderCommands.Length; j++ {
		renderCommand := clay.RenderCommandArray_Get(&renderCommands, j)
		boundingBox := renderCommand.BoundingBox
		switch renderCommand.CommandType {
		case clay.RENDER_COMMAND_TYPE_TEXT:
			textData := &renderCommand.RenderData.Text
			cloned := strings.Clone(textData.StringContents.String())
			fontToUse := fonts[textData.FontId]
			rl.DrawTextEx(fontToUse, cloned, rl.Vector2{X: boundingBox.X, Y: boundingBox.Y}, float32(textData.FontSize), float32(textData.LetterSpacing), rl.Color{R: uint8(textData.TextColor.R), G: uint8(textData.TextColor.G), B: uint8(textData.TextColor.B), A: uint8(textData.TextColor.A)})
		case clay.RENDER_COMMAND_TYPE_IMAGE:
			imageTexture := *(*rl.Texture2D)(renderCommand.RenderData.Image.ImageData)
			tintColor := renderCommand.RenderData.Image.BackgroundColor
			if tintColor.R == 0 && tintColor.G == 0 && tintColor.B == 0 && tintColor.A == 0 {
				tintColor = clay.Color{R: 255, G: 255, B: 255, A: 255}
			}
			rl.DrawTextureEx(
				imageTexture,
				rl.Vector2{X: boundingBox.X, Y: boundingBox.Y},
				0,
				boundingBox.Width/float32(imageTexture.Width),
				rl.Color{R: uint8(tintColor.R), G: uint8(tintColor.G), B: uint8(tintColor.B), A: uint8(tintColor.A)},
			)
		case clay.RENDER_COMMAND_TYPE_NONE:
		case clay.RENDER_COMMAND_TYPE_RECTANGLE:
			config := &renderCommand.RenderData.Rectangle
			if config.CornerRadius.TopLeft > 0 {
				r := boundingBox.Width
				if boundingBox.Width > boundingBox.Height {
					r = boundingBox.Height
				}
				radius := (config.CornerRadius.TopLeft * 2) / r
				rl.DrawRectangleRounded(rl.Rectangle{X: boundingBox.X, Y: boundingBox.Y, Width: boundingBox.Width, Height: boundingBox.Height}, radius, 8, rl.Color{R: uint8(config.BackgroundColor.R), G: uint8(config.BackgroundColor.G), B: uint8(config.BackgroundColor.B), A: uint8(config.BackgroundColor.A)})
			} else {
				rl.DrawRectangle(int32(boundingBox.X), int32(boundingBox.Y), int32(boundingBox.Width), int32(boundingBox.Height), rl.Color{R: uint8(config.BackgroundColor.R), G: uint8(config.BackgroundColor.G), B: uint8(config.BackgroundColor.B), A: uint8(config.BackgroundColor.A)})
			}
		case clay.RENDER_COMMAND_TYPE_BORDER:
		case clay.RENDER_COMMAND_TYPE_SCISSOR_START:
		case clay.RENDER_COMMAND_TYPE_SCISSOR_END:
		case clay.RENDER_COMMAND_TYPE_CUSTOM:

		}
	}
}

//void Clay_Raylib_Render(Clay_RenderCommandArray renderCommands, Font* fonts)
//{
//    for (int j = 0; j < renderCommands.length; j++)
//    {
//        Clay_RenderCommand *renderCommand = Clay_RenderCommandArray_Get(&renderCommands, j);
//        Clay_BoundingBox boundingBox = renderCommand->boundingBox;
//        switch (renderCommand->commandType)
//        {
//            case CLAY_RENDER_COMMAND_TYPE_SCISSOR_START: {
//                BeginScissorMode((int)roundf(boundingBox.x), (int)roundf(boundingBox.y), (int)roundf(boundingBox.width), (int)roundf(boundingBox.height));
//                break;
//            }
//            case CLAY_RENDER_COMMAND_TYPE_SCISSOR_END: {
//                EndScissorMode();
//                break;
//            }
//            case CLAY_RENDER_COMMAND_TYPE_RECTANGLE: {
//                Clay_RectangleRenderData *config = &renderCommand->renderData.rectangle;
//                if (config->cornerRadius.topLeft > 0) {
//                    float radius = (config->cornerRadius.topLeft * 2) / (float)((boundingBox.width > boundingBox.height) ? boundingBox.height : boundingBox.width);
//                    DrawRectangleRounded((Rectangle) { boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height }, radius, 8, CLAY_COLOR_TO_RAYLIB_COLOR(config->backgroundColor));
//                } else {
//                    DrawRectangle(boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height, CLAY_COLOR_TO_RAYLIB_COLOR(config->backgroundColor));
//                }
//                break;
//            }
//            case CLAY_RENDER_COMMAND_TYPE_BORDER: {
//                Clay_BorderRenderData *config = &renderCommand->renderData.border;
//                // Left border
//                if (config->width.left > 0) {
//                    DrawRectangle((int)roundf(boundingBox.x), (int)roundf(boundingBox.y + config->cornerRadius.topLeft), (int)config->width.left, (int)roundf(boundingBox.height - config->cornerRadius.topLeft - config->cornerRadius.bottomLeft), CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                // Right border
//                if (config->width.right > 0) {
//                    DrawRectangle((int)roundf(boundingBox.x + boundingBox.width - config->width.right), (int)roundf(boundingBox.y + config->cornerRadius.topRight), (int)config->width.right, (int)roundf(boundingBox.height - config->cornerRadius.topRight - config->cornerRadius.bottomRight), CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                // Top border
//                if (config->width.top > 0) {
//                    DrawRectangle((int)roundf(boundingBox.x + config->cornerRadius.topLeft), (int)roundf(boundingBox.y), (int)roundf(boundingBox.width - config->cornerRadius.topLeft - config->cornerRadius.topRight), (int)config->width.top, CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                // Bottom border
//                if (config->width.bottom > 0) {
//                    DrawRectangle((int)roundf(boundingBox.x + config->cornerRadius.bottomLeft), (int)roundf(boundingBox.y + boundingBox.height - config->width.bottom), (int)roundf(boundingBox.width - config->cornerRadius.bottomLeft - config->cornerRadius.bottomRight), (int)config->width.bottom, CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                if (config->cornerRadius.topLeft > 0) {
//                    DrawRing((Vector2) { roundf(boundingBox.x + config->cornerRadius.topLeft), roundf(boundingBox.y + config->cornerRadius.topLeft) }, roundf(config->cornerRadius.topLeft - config->width.top), config->cornerRadius.topLeft, 180, 270, 10, CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                if (config->cornerRadius.topRight > 0) {
//                    DrawRing((Vector2) { roundf(boundingBox.x + boundingBox.width - config->cornerRadius.topRight), roundf(boundingBox.y + config->cornerRadius.topRight) }, roundf(config->cornerRadius.topRight - config->width.top), config->cornerRadius.topRight, 270, 360, 10, CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                if (config->cornerRadius.bottomLeft > 0) {
//                    DrawRing((Vector2) { roundf(boundingBox.x + config->cornerRadius.bottomLeft), roundf(boundingBox.y + boundingBox.height - config->cornerRadius.bottomLeft) }, roundf(config->cornerRadius.bottomLeft - config->width.top), config->cornerRadius.bottomLeft, 90, 180, 10, CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                if (config->cornerRadius.bottomRight > 0) {
//                    DrawRing((Vector2) { roundf(boundingBox.x + boundingBox.width - config->cornerRadius.bottomRight), roundf(boundingBox.y + boundingBox.height - config->cornerRadius.bottomRight) }, roundf(config->cornerRadius.bottomRight - config->width.bottom), config->cornerRadius.bottomRight, 0.1, 90, 10, CLAY_COLOR_TO_RAYLIB_COLOR(config->color));
//                }
//                break;
//            }
//            case CLAY_RENDER_COMMAND_TYPE_CUSTOM: {
//                Clay_CustomRenderData *config = &renderCommand->renderData.custom;
//                CustomLayoutElement *customElement = (CustomLayoutElement *)config->customData;
//                if (!customElement) continue;
//                switch (customElement->type) {
//                    case CUSTOM_LAYOUT_ELEMENT_TYPE_3D_MODEL: {
//                        Clay_BoundingBox rootBox = renderCommands.internalArray[0].boundingBox;
//                        float scaleValue = CLAY__MIN(CLAY__MIN(1, 768 / rootBox.height) * CLAY__MAX(1, rootBox.width / 1024), 1.5f);
//                        Ray positionRay = GetScreenToWorldPointWithZDistance((Vector2) { renderCommand->boundingBox.x + renderCommand->boundingBox.width / 2, renderCommand->boundingBox.y + (renderCommand->boundingBox.height / 2) + 20 }, Raylib_camera, (int)roundf(rootBox.width), (int)roundf(rootBox.height), 140);
//                        BeginMode3D(Raylib_camera);
//                            DrawModel(customElement->customData.model.model, positionRay.position, customElement->customData.model.scale * scaleValue, WHITE);        // Draw 3d model with texture
//                        EndMode3D();
//                        break;
//                    }
//                    default: break;
//                }
//                break;
//            }
//            default: {
//                printf("Error: unhandled render command.");
//                exit(1);
//            }
//        }
//    }
//}
