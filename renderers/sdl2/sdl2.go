package sdl2

import (
	"unsafe"

	"github.com/totallygamerjet/clay"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	FontId uint32
	Font   *ttf.Font
}

// static Clay_Dimensions SDL2_MeasureText(Clay_StringSlice text, Clay_TextElementConfig *config, void *userData)
//
//	{
//	   SDL2_Font *fonts = (SDL2_Font*)userData;
//
//	   TTF_Font *font = fonts[config->fontId].font;
//	   char *chars = (char *)calloc(text.length + 1, 1);
//	   memcpy(chars, text.chars, text.length);
//	   int width = 0;
//	   int height = 0;
//	   if (TTF_SizeUTF8(font, chars, &width, &height) < 0) {
//	       fprintf(stderr, "Error: could not measure text: %s\n", TTF_GetError());
//	       exit(1);
//	   }
//	   free(chars);
//	   return (Clay_Dimensions) {
//	           .width = (float)width,
//	           .height = (float)height,
//	   };
//	}

func MeasureText(text clay.StringSlice, config *clay.TextElementConfig, userData unsafe.Pointer) clay.Dimensions {
	return clay.Dimensions{}
}
