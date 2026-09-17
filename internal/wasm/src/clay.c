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

// Transitions call back into Go through function pointers that carry no user data,
// so there is a pool of trampolines, each of which knows its own slot.
// Go registers a function in a slot and stores the matching trampoline in the element.

#define GO_TRANSITION_SLOTS 16

GO_IMPORT("transitionHandler")
bool go_transition_handler(int32_t slot, Clay_TransitionCallbackArguments *arguments);

GO_IMPORT("transitionState")
void go_transition_state(int32_t slot, Clay_TransitionData *ret, Clay_TransitionData *state, Clay_TransitionProperty properties);

#define GO_TRANSITION_TRAMPOLINE(slot)                                                            \
    static bool go_handler_##slot(Clay_TransitionCallbackArguments arguments) {                   \
        return go_transition_handler(slot, &arguments);                                           \
    }                                                                                             \
    static Clay_TransitionData go_state_##slot(Clay_TransitionData state,                         \
                                               Clay_TransitionProperty properties) {              \
        Clay_TransitionData ret = {0};                                                            \
        go_transition_state(slot, &ret, &state, properties);                                      \
        return ret;                                                                               \
    }

GO_TRANSITION_TRAMPOLINE(0)
GO_TRANSITION_TRAMPOLINE(1)
GO_TRANSITION_TRAMPOLINE(2)
GO_TRANSITION_TRAMPOLINE(3)
GO_TRANSITION_TRAMPOLINE(4)
GO_TRANSITION_TRAMPOLINE(5)
GO_TRANSITION_TRAMPOLINE(6)
GO_TRANSITION_TRAMPOLINE(7)
GO_TRANSITION_TRAMPOLINE(8)
GO_TRANSITION_TRAMPOLINE(9)
GO_TRANSITION_TRAMPOLINE(10)
GO_TRANSITION_TRAMPOLINE(11)
GO_TRANSITION_TRAMPOLINE(12)
GO_TRANSITION_TRAMPOLINE(13)
GO_TRANSITION_TRAMPOLINE(14)
GO_TRANSITION_TRAMPOLINE(15)

static bool (*const go_handlers[GO_TRANSITION_SLOTS])(Clay_TransitionCallbackArguments) = {
    go_handler_0, go_handler_1, go_handler_2, go_handler_3,
    go_handler_4, go_handler_5, go_handler_6, go_handler_7,
    go_handler_8, go_handler_9, go_handler_10, go_handler_11,
    go_handler_12, go_handler_13, go_handler_14, go_handler_15,
};

static Clay_TransitionData (*const go_states[GO_TRANSITION_SLOTS])(Clay_TransitionData, Clay_TransitionProperty) = {
    go_state_0, go_state_1, go_state_2, go_state_3,
    go_state_4, go_state_5, go_state_6, go_state_7,
    go_state_8, go_state_9, go_state_10, go_state_11,
    go_state_12, go_state_13, go_state_14, go_state_15,
};

// go_transition_handler_ptr returns the trampoline of a slot, or null if there are none left.
GO_EXPORT("go_transition_handler_ptr")
void *go_transition_handler_ptr(int32_t slot) {
    return slot >= 0 && slot < GO_TRANSITION_SLOTS ? (void *)go_handlers[slot] : NULL;
}

GO_EXPORT("go_transition_state_ptr")
void *go_transition_state_ptr(int32_t slot) {
    return slot >= 0 && slot < GO_TRANSITION_SLOTS ? (void *)go_states[slot] : NULL;
}

// go_ease_out_ptr returns clay's built in transition handler, which needs no trampoline.
GO_EXPORT("go_ease_out_ptr")
void *go_ease_out_ptr(void) {
    return (void *)Clay_EaseOut;
}
