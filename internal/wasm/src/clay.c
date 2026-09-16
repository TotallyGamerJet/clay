// This file is the entry point of the Wasm build of clay.h.
// It is compiled together with bindings generated from clay.h by internal/cmd/generate
// and then translated to Go by wasm2go. Run `go generate` in the repository root to rebuild.
//
// It contains the bits that can't be generated: trampolines for the function pointers
// clay calls back into and a memory allocator for the Go side.

#define CLAY_IMPLEMENTATION
#include "clay.h"

#include <stdlib.h>

#define GO_EXPORT(name) __attribute__((export_name(name)))
#define GO_IMPORT(name) __attribute__((import_module("go"), import_name(name)))

GO_EXPORT("go_malloc")
void *go_malloc(size_t size) {
    return calloc(1, size);
}

GO_EXPORT("go_free")
void go_free(void *ptr) {
    free(ptr);
}

GO_IMPORT("errorHandler")
void go_error_handler(Clay_ErrorData *errorData);

static void error_handler(Clay_ErrorData errorData) {
    go_error_handler(&errorData);
}

GO_EXPORT("go_initialize")
Clay_Context *go_initialize(size_t capacity, void *memory, Clay_Dimensions *layoutDimensions, void *userData) {
    Clay_Arena arena = Clay_CreateArenaWithCapacityAndMemory(capacity, memory);
    return Clay_Initialize(arena, *layoutDimensions, (Clay_ErrorHandler) {
        .errorHandlerFunction = error_handler,
        .userData = userData,
    });
}

GO_IMPORT("measureText")
void go_measure_text(Clay_Dimensions *ret, Clay_StringSlice *text, Clay_TextElementConfig *config, void *userData);

static Clay_Dimensions measure_text(Clay_StringSlice text, Clay_TextElementConfig *config, void *userData) {
    Clay_Dimensions ret = {0};
    go_measure_text(&ret, &text, config, userData);
    return ret;
}

GO_EXPORT("go_set_measure_text_function")
void go_set_measure_text_function(bool enabled, void *userData) {
    Clay_SetMeasureTextFunction(enabled ? measure_text : NULL, userData);
}

GO_IMPORT("queryScrollOffset")
void go_query_scroll_offset(Clay_Vector2 *ret, uint32_t elementId, void *userData);

static Clay_Vector2 query_scroll_offset(uint32_t elementId, void *userData) {
    Clay_Vector2 ret = {0};
    go_query_scroll_offset(&ret, elementId, userData);
    return ret;
}

GO_EXPORT("go_set_query_scroll_offset_function")
void go_set_query_scroll_offset_function(bool enabled, void *userData) {
    Clay_SetQueryScrollOffsetFunction(enabled ? query_scroll_offset : NULL, userData);
}

GO_IMPORT("onHover")
void go_on_hover_function(Clay_ElementId *elementId, Clay_PointerData *pointerData, void *userData);

static void on_hover(Clay_ElementId elementId, Clay_PointerData pointerData, void *userData) {
    go_on_hover_function(&elementId, &pointerData, userData);
}

GO_EXPORT("go_on_hover")
void go_on_hover(void *userData) {
    Clay_OnHover(on_hover, userData);
}

GO_EXPORT("go_open_text_element")
void go_open_text_element(Clay_String *text, Clay_TextElementConfig *config) {
    // Use the macros, which are more stable than the functions they expand to.
    Clay_String t = *text;
    Clay_TextElementConfig c = *config;
    CLAY_TEXT(t, CLAY_TEXT_CONFIG(c));
}
