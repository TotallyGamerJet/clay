// SPDX-License-Identifier: Zlib

package clay

import (
	"fmt"
)

// This file is the part of clay's API that is written by hand, rather than
// generated from clay.h into clay.go.

// confirm that ErrorData implements error type
var _ error = ErrorData{}

// Error implements the error interface, so that an ErrorData can be returned as an error.
func (e ErrorData) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.ErrorText, e.ErrorType)
}

// ID returns the element id for a label, for use as the id of an element or with functions
// like PointerOver and GetScrollContainerData.
func ID(label string) ElementId {
	return __HashString(label, 0)
}

// IDI returns the element id for a label and an index, for giving unique ids to the
// elements of a loop.
func IDI(label string, index uint32) ElementId {
	return __HashStringWithOffset(label, index, 0)
}

// PaddingAll returns the same padding for all four sides of an element.
func PaddingAll(padding uint16) Padding {
	return Padding{
		padding,
		padding,
		padding,
		padding,
	}
}

// SizingGrow returns a sizing that expands along its axis to fill the available space in the
// parent element, sharing it with the other growing elements. The size is clamped to sz,
// or unclamped when sz is 0.
func SizingGrow(sz float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			MinMax: SizingMinMax{sz, sz},
		},
		Type: sizingTypeGrow,
	}
}

// SizingFixed returns a sizing that clamps the axis to an exact size in pixels.
func SizingFixed(sz float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			MinMax: SizingMinMax{sz, sz},
		},
		Type: sizingTypeFixed,
	}
}

// SizingFit returns a sizing that wraps tightly to the size of the element's contents,
// clamped between min and max. This is the default sizing.
func SizingFit(min, max float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			MinMax: SizingMinMax{min, max},
		},
		Type: sizingTypeFit,
	}
}

// SizingPercent returns a sizing that clamps the axis to a percent, in the range 0-1, of the
// parent container's axis size minus its padding and child gaps.
func SizingPercent(percentOfParent float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			Percent: percentOfParent,
		},
		Type: sizingTypePercent,
	}
}

// CornerRadiusAll returns the same corner radius for all four corners of an element.
func CornerRadiusAll(radius float32) CornerRadius {
	return CornerRadius{
		radius,
		radius,
		radius,
		radius,
	}
}

// BorderOutside returns a border of the same width around an element, but not between its children.
func BorderOutside(width uint16) BorderWidth {
	return BorderWidth{
		width,
		width,
		width,
		width,
		0,
	}
}

// BorderAll returns a border of the same width around an element and between its children.
func BorderAll(width uint16) BorderWidth {
	return BorderWidth{
		width,
		width,
		width,
		width,
		width,
	}
}

// TODO: add generic iterator functions for types with [type]_GetValue functions that are converted into methods

// UI declares an element, and returns the function that configures it and declares its children.
// It takes at most one id; without one the element is given an id generated from its position
// in the layout. It is used as:
//
//	clay.UI(clay.ID("Sidebar"))(clay.ElementDeclaration{
//		BackgroundColor: clay.Color{R: 90, G: 90, B: 90, A: 255},
//	}, func() {
//		// ...children declared here
//	})
//
// The children function may be nil for an element without children.
func UI(id ...ElementId) func(decl ElementDeclaration, children func()) {
	if len(id) > 1 {
		panic("clay: too many element ids")
	} else if len(id) == 1 {
		__OpenElementWithId(id[0])
	} else {
		__OpenElement()
	}
	return func(decl ElementDeclaration, children func()) {
		__ConfigureOpenElement(decl)
		defer __CloseElement()
		if children != nil {
			children()
		}
	}
}

// Text declares a text element inside the currently open element.
// A nil config is the same as the zero TextElementConfig.
func Text(text string, config *TextElementConfig) {
	if config == nil {
		config = &TextElementConfig{}
	}
	openTextElement(text, config)
}

// Pointer points to a value of one of the [PointerTarget] types inside clay's memory.
type Pointer[T PointerTarget] struct {
	addr uint32
}

// IsNil reports whether the pointer is nil.
func (p Pointer[T]) IsNil() bool {
	return p.addr == 0
}

// Get returns the value p points to, or the zero value if p is nil.
func (p Pointer[T]) Get() T {
	var v T
	if p.addr != 0 {
		decodeValue(wasmMemory(), p.addr, &v)
	}
	return v
}

// Set sets the value p points to.
func (p Pointer[T]) Set(v T) {
	if p.addr == 0 {
		panic("clay: Set on nil Pointer")
	}
	var buf [maxPointerSize]byte
	n := encodeValue(buf[:], 0, &v)
	copy(wasmMemory()[p.addr:], buf[:n])
}

// Context is a clay context, which holds all of the state of one layout.
// Use SetCurrentContext to switch between the contexts returned by Initialize.
type Context struct {
	addr uint32
}

var contexts = map[uint32]*Context{}

func contextOf(addr uint32) *Context {
	if addr == 0 {
		return nil
	}
	c, ok := contexts[addr]
	if !ok {
		c = &Context{addr: addr}
		contexts[addr] = c
	}
	return c
}

// Clay calls transition functions through function pointers that carry no user data,
// so each Go function registered gets a slot, and with it one of the trampolines in the
// module that calls back into Go with that slot. There is a limited number of them, and a
// registration lasts for the lifetime of the program, so they are meant to be made once,
// not per frame or per element.

// TransitionHandler advances a transition, and is called once a frame while one is running.
// Use [EaseOut], or [NewTransitionHandler] to write your own.
type TransitionHandler struct {
	ptr uint32
}

// EaseOut is clay's built in transition handler, which eases towards the target state.
var EaseOut = TransitionHandler{ptr: uint32(module.Xgo_ease_out_ptr())}

// NewTransitionHandler registers fn as a transition handler, which is called with the state
// of a running transition and writes the current state of it through arguments.Current.
// It returns whether the transition is complete.
//
// It panics if there are no transition slots left.
func NewTransitionHandler(fn func(arguments TransitionCallbackArguments) bool) TransitionHandler {
	slot := len(transitionHandlers)
	ptr := uint32(module.Xgo_transition_handler_ptr(int32(slot)))
	if ptr == 0 {
		panic("clay: too many transition handlers")
	}
	transitionHandlers = append(transitionHandlers, fn)
	return TransitionHandler{ptr: ptr}
}

// TransitionStateFunc returns the state a transition starts from, or ends at.
// Use [NewTransitionStateFunc] to make one.
type TransitionStateFunc struct {
	ptr uint32
}

// NewTransitionStateFunc registers fn as the function that gives an entering element the
// state it transitions from, or an exiting element the state it transitions to.
//
// It panics if there are no transition slots left.
func NewTransitionStateFunc(fn func(state TransitionData, properties TransitionProperty) TransitionData) TransitionStateFunc {
	slot := len(transitionStateFuncs)
	ptr := uint32(module.Xgo_transition_state_ptr(int32(slot)))
	if ptr == 0 {
		panic("clay: too many transition state functions")
	}
	transitionStateFuncs = append(transitionStateFuncs, fn)
	return TransitionStateFunc{ptr: ptr}
}

// currentContextAddr returns the address of the context the callbacks are set on.
func currentContextAddr() uint32 {
	return uint32(GetCurrentContext().wasmAddr())
}

func (c *Context) wasmAddr() int32 {
	if c == nil {
		return 0
	}
	return int32(c.addr)
}

// Arena is memory inside clay's WebAssembly module used for its internal allocations.
type Arena struct {
	capacity, memory uint32
}

// CreateArenaWithCapacity allocates an arena of capacity bytes.
// Use MinMemorySize to know how big it needs to be.
func CreateArenaWithCapacity(capacity uint32) Arena {
	return Arena{capacity: capacity, memory: wasmMalloc(capacity)}
}

// Free releases the memory of the arena, which is left empty.
// The context created in the arena must not be used afterwards, so if it is the current
// context, clay is left without one until the next call to Initialize or SetCurrentContext.
func (a *Arena) Free() {
	if a.memory == 0 {
		return
	}
	if c := GetCurrentContext(); c != nil && c.addr >= a.memory && c.addr < a.memory+a.capacity {
		SetCurrentContext(nil)
		delete(contexts, c.addr)
	}
	module.Xgo_free(int32(a.memory))
	*a = Arena{}
}

// ErrorHandler is called by clay when it encounters an error, such as running out of arena space.
type ErrorHandler struct {
	// ErrorHandlerFunction is called with the error, which is also a Go error.
	ErrorHandlerFunction func(errorData ErrorData)
	// UserData is passed through to the error handler as the UserData of the ErrorData.
	UserData any
}

// Initialize creates a clay context in the arena, makes it current, and returns it.
// The dimensions are the size of the root layout element, which SetLayoutDimensions changes later.
func Initialize(arena Arena, layoutDimensions Dimensions, errorHandler ErrorHandler) *Context {
	h := persistentHandle | allocPersistent(&errorHandler)
	b := wasmArgs(sizeofDimensions)
	encDimensions(b, 0, &layoutDimensions)
	p := wasmCommit(b)
	ctx := contextOf(uint32(module.Xgo_initialize(int32(arena.capacity), int32(arena.memory), int32(p), int32(h))))
	// Initializing over a context replaces its error handler, so the old one can go.
	takePersistentSlot(persistentKey{errorHandlerHandle, ctx.addr}, h)
	return ctx
}

type measureTextFunction struct {
	fn       func(text string, config *TextElementConfig, userData any) Dimensions
	userData any
}

// SetMeasureTextFunction sets the function clay calls to measure text for the current context,
// passing it userData. Text can't be laid out without it.
//
// The config passed to the function must not be retained.
func SetMeasureTextFunction(measureText func(text string, config *TextElementConfig, userData any) Dimensions, userData any) {
	h := storePersistentHandle(persistentKey{measureTextHandle, currentContextAddr()}, &measureTextFunction{measureText, userData})
	module.Xgo_set_measure_text_function(b2i(measureText != nil), int32(h))
}

type queryScrollOffsetFunction struct {
	fn       func(elementId uint32, userData any) Vector2
	userData any
}

// SetQueryScrollOffsetFunction sets the function clay calls to get the scroll offset of an
// element from an external scrolling system, passing it userData.
// It is only used when SetExternalScrollHandlingEnabled is set.
func SetQueryScrollOffsetFunction(queryScrollOffset func(elementId uint32, userData any) Vector2, userData any) {
	h := storePersistentHandle(persistentKey{queryScrollOffsetHandle, currentContextAddr()}, &queryScrollOffsetFunction{queryScrollOffset, userData})
	module.Xgo_set_query_scroll_offset_function(b2i(queryScrollOffset != nil), int32(h))
}

type onHoverFunction struct {
	fn       func(elementId ElementId, pointerData PointerData, userData any)
	userData any
}

// OnHover binds a function that is called when the pointer position provided by SetPointerState
// is within the currently open element's bounding box, passing it userData.
// The function is called by SetPointerState, not by OnHover itself.
func OnHover(onHover func(elementId ElementId, pointerData PointerData, userData any), userData any) {
	module.Xgo_on_hover(int32(storeHandle(&onHoverFunction{onHover, userData})))
}

// BeginLayout prepares clay for the declaration of a new layout, and invalidates the strings and
// user data of the layout before last.
func BeginLayout() {
	nextFrame()
	beginLayout()
}
