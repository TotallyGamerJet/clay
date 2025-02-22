package clay

import (
	"math"
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/libc"
)

const __NULL = 0
const __MAXFLOAT = 3.4028234663852886e+38

var __ELEMENT_DEFINITION_LATCH uint8

type String struct {
	Length int32
	Chars  *byte
}
type StringSlice struct {
	Length    int32
	Chars     *byte
	BaseChars *byte
}
type Context struct {
	MaxElementCount                    int32
	MaxMeasureTextCacheWordCount       int32
	WarningsEnabled                    bool
	ErrorHandler                       ErrorHandler
	BooleanWarnings                    BooleanWarnings
	Warnings                           __WarningArray
	PointerInfo                        PointerData
	LayoutDimensions                   Dimensions
	DynamicElementIndexBaseHash        ElementId
	DynamicElementIndex                uint32
	DebugModeEnabled                   bool
	DisableCulling                     bool
	ExternalScrollHandlingEnabled      bool
	DebugSelectedElementId             uint32
	Generation                         uint32
	ArenaResetOffset                   uint64
	MeasureTextUserData                any
	QueryScrollOffsetUserData          any
	InternalArena                      Arena
	LayoutElements                     LayoutElementArray
	RenderCommands                     RenderCommandArray
	OpenLayoutElementStack             __int32_tArray
	LayoutElementChildren              __int32_tArray
	LayoutElementChildrenBuffer        __int32_tArray
	TextElementData                    __TextElementDataArray
	ImageElementPointers               __int32_tArray
	ReusableElementIndexBuffer         __int32_tArray
	LayoutElementClipElementIds        __int32_tArray
	LayoutConfigs                      __LayoutConfigArray
	ElementConfigs                     __ElementConfigArray
	TextElementConfigs                 __TextElementConfigArray
	ImageElementConfigs                __ImageElementConfigArray
	FloatingElementConfigs             __FloatingElementConfigArray
	ScrollElementConfigs               __ScrollElementConfigArray
	CustomElementConfigs               __CustomElementConfigArray
	BorderElementConfigs               __BorderElementConfigArray
	SharedElementConfigs               __SharedElementConfigArray
	LayoutElementIdStrings             __StringArray
	WrappedTextLines                   __WrappedTextLineArray
	LayoutElementTreeNodeArray1        __LayoutElementTreeNodeArray
	LayoutElementTreeRoots             __LayoutElementTreeRootArray
	LayoutElementsHashMapInternal      __LayoutElementHashMapItemArray
	LayoutElementsHashMap              __int32_tArray
	MeasureTextHashMapInternal         __MeasureTextCacheItemArray
	MeasureTextHashMapInternalFreeList __int32_tArray
	MeasureTextHashMap                 __int32_tArray
	MeasuredWords                      __MeasuredWordArray
	MeasuredWordsFreeList              __int32_tArray
	OpenClipElementStack               __int32_tArray
	PointerOverIds                     __ElementIdArray
	ScrollContainerDatas               __ScrollContainerDataInternalArray
	TreeNodeVisited                    __boolArray
	DynamicStringData                  __charArray
	DebugElementData                   __DebugElementDataArray
}
type Arena struct {
	NextAllocation uint64
	Capacity       uint64
	Memory         *byte
}
type Dimensions struct {
	Width  float32
	Height float32
}
type Vector2 struct {
	X float32
	Y float32
}
type Color struct {
	R float32
	G float32
	B float32
	A float32
}
type BoundingBox struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}
type ElementId struct {
	Id       uint32
	Offset   uint32
	BaseId   uint32
	StringId String
}
type CornerRadius struct {
	TopLeft     float32
	TopRight    float32
	BottomLeft  float32
	BottomRight float32
}
type LayoutDirection int32

const (
	LEFT_TO_RIGHT = LayoutDirection(iota)
	TOP_TO_BOTTOM
)

type LayoutAlignmentX int32

const (
	ALIGN_X_LEFT = LayoutAlignmentX(iota)
	ALIGN_X_RIGHT
	ALIGN_X_CENTER
)

type LayoutAlignmentY int32

const (
	ALIGN_Y_TOP = LayoutAlignmentY(iota)
	ALIGN_Y_BOTTOM
	ALIGN_Y_CENTER
)

type __SizingType int32

const (
	__SIZING_TYPE_FIT = __SizingType(iota)
	__SIZING_TYPE_GROW
	__SIZING_TYPE_PERCENT
	__SIZING_TYPE_FIXED
)

type ChildAlignment struct {
	X LayoutAlignmentX
	Y LayoutAlignmentY
}
type SizingMinMax struct {
	Min float32
	Max float32
}
type SizingAxis struct {
	Size struct {
		// union
		MinMax  SizingMinMax
		Percent float32
	}
	Type __SizingType
}
type Sizing struct {
	Width  SizingAxis
	Height SizingAxis
}
type Padding struct {
	Left   uint16
	Right  uint16
	Top    uint16
	Bottom uint16
}
type __PaddingWrapper struct {
	Wrapped Padding
}
type LayoutConfig struct {
	Sizing          Sizing
	Padding         Padding
	ChildGap        uint16
	ChildAlignment  ChildAlignment
	LayoutDirection LayoutDirection
}
type __LayoutConfigWrapper struct {
	Wrapped LayoutConfig
}
type TextElementConfigWrapMode int32

const (
	TEXT_WRAP_WORDS = TextElementConfigWrapMode(iota)
	TEXT_WRAP_NEWLINES
	TEXT_WRAP_NONE
)

type TextAlignment int32

const (
	TEXT_ALIGN_LEFT = TextAlignment(iota)
	TEXT_ALIGN_CENTER
	TEXT_ALIGN_RIGHT
)

type TextElementConfig struct {
	TextColor          Color
	FontId             uint16
	FontSize           uint16
	LetterSpacing      uint16
	LineHeight         uint16
	WrapMode           TextElementConfigWrapMode
	TextAlignment      TextAlignment
	HashStringContents bool
}
type __TextElementConfigWrapper struct {
	Wrapped TextElementConfig
}
type ImageElementConfig struct {
	ImageData        unsafe.Pointer
	SourceDimensions Dimensions
}
type __ImageElementConfigWrapper struct {
	Wrapped ImageElementConfig
}
type FloatingAttachPointType int32

const (
	ATTACH_POINT_LEFT_TOP = FloatingAttachPointType(iota)
	ATTACH_POINT_LEFT_CENTER
	ATTACH_POINT_LEFT_BOTTOM
	ATTACH_POINT_CENTER_TOP
	ATTACH_POINT_CENTER_CENTER
	ATTACH_POINT_CENTER_BOTTOM
	ATTACH_POINT_RIGHT_TOP
	ATTACH_POINT_RIGHT_CENTER
	ATTACH_POINT_RIGHT_BOTTOM
)

type FloatingAttachPoints struct {
	Element FloatingAttachPointType
	Parent  FloatingAttachPointType
}
type PointerCaptureMode int32

const (
	POINTER_CAPTURE_MODE_CAPTURE = PointerCaptureMode(iota)
	POINTER_CAPTURE_MODE_PASSTHROUGH
)

type FloatingAttachToElement int32

const (
	ATTACH_TO_NONE = FloatingAttachToElement(iota)
	ATTACH_TO_PARENT
	ATTACH_TO_ELEMENT_WITH_ID
	ATTACH_TO_ROOT
)

type FloatingElementConfig struct {
	Offset             Vector2
	Expand             Dimensions
	ParentId           uint32
	ZIndex             int16
	AttachPoints       FloatingAttachPoints
	PointerCaptureMode PointerCaptureMode
	AttachTo           FloatingAttachToElement
}
type __FloatingElementConfigWrapper struct {
	Wrapped FloatingElementConfig
}
type CustomElementConfig struct {
	CustomData unsafe.Pointer
}
type __CustomElementConfigWrapper struct {
	Wrapped CustomElementConfig
}
type ScrollElementConfig struct {
	Horizontal bool
	Vertical   bool
}
type __ScrollElementConfigWrapper struct {
	Wrapped ScrollElementConfig
}
type BorderWidth struct {
	Left            uint16
	Right           uint16
	Top             uint16
	Bottom          uint16
	BetweenChildren uint16
}
type BorderElementConfig struct {
	Color Color
	Width BorderWidth
}
type __BorderElementConfigWrapper struct {
	Wrapped BorderElementConfig
}
type TextRenderData struct {
	StringContents StringSlice
	TextColor      Color
	FontId         uint16
	FontSize       uint16
	LetterSpacing  uint16
	LineHeight     uint16
}
type RectangleRenderData struct {
	BackgroundColor Color
	CornerRadius    CornerRadius
}
type ImageRenderData struct {
	BackgroundColor  Color
	CornerRadius     CornerRadius
	SourceDimensions Dimensions
	ImageData        unsafe.Pointer
}
type CustomRenderData struct {
	BackgroundColor Color
	CornerRadius    CornerRadius
	CustomData      unsafe.Pointer
}
type ScrollRenderData struct {
	Horizontal bool
	Vertical   bool
}
type BorderRenderData struct {
	Color        Color
	CornerRadius CornerRadius
	Width        BorderWidth
}
type RenderData struct {
	// union
	Rectangle RectangleRenderData
	Text      TextRenderData
	Image     ImageRenderData
	Custom    CustomRenderData
	Border    BorderRenderData
	Scroll    ScrollRenderData
}
type ScrollContainerData struct {
	ScrollPosition            *Vector2
	ScrollContainerDimensions Dimensions
	ContentDimensions         Dimensions
	Config                    ScrollElementConfig
	Found                     bool
}
type ElementData struct {
	BoundingBox BoundingBox
	Found       bool
}
type RenderCommandType int32

const (
	RENDER_COMMAND_TYPE_NONE = RenderCommandType(iota)
	RENDER_COMMAND_TYPE_RECTANGLE
	RENDER_COMMAND_TYPE_BORDER
	RENDER_COMMAND_TYPE_TEXT
	RENDER_COMMAND_TYPE_IMAGE
	RENDER_COMMAND_TYPE_SCISSOR_START
	RENDER_COMMAND_TYPE_SCISSOR_END
	RENDER_COMMAND_TYPE_CUSTOM
)

type RenderCommand struct {
	BoundingBox BoundingBox
	RenderData  RenderData
	UserData    unsafe.Pointer
	Id          uint32
	ZIndex      int16
	CommandType RenderCommandType
}
type RenderCommandArray struct {
	Capacity      int32
	Length        int32
	InternalArray *RenderCommand
}
type PointerDataInteractionState int32

const (
	POINTER_DATA_PRESSED_THIS_FRAME = PointerDataInteractionState(iota)
	POINTER_DATA_PRESSED
	POINTER_DATA_RELEASED_THIS_FRAME
	POINTER_DATA_RELEASED
)

type PointerData struct {
	Position Vector2
	State    PointerDataInteractionState
}
type ElementDeclaration struct {
	Id              ElementId
	Layout          LayoutConfig
	BackgroundColor Color
	CornerRadius    CornerRadius
	Image           ImageElementConfig
	Floating        FloatingElementConfig
	Custom          CustomElementConfig
	Scroll          ScrollElementConfig
	Border          BorderElementConfig
	UserData        unsafe.Pointer
}
type __ElementDeclarationWrapper struct {
	Wrapped ElementDeclaration
}
type ErrorType int32

const (
	ERROR_TYPE_TEXT_MEASUREMENT_FUNCTION_NOT_PROVIDED = ErrorType(iota)
	ERROR_TYPE_ARENA_CAPACITY_EXCEEDED
	ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED
	ERROR_TYPE_TEXT_MEASUREMENT_CAPACITY_EXCEEDED
	ERROR_TYPE_DUPLICATE_ID
	ERROR_TYPE_FLOATING_CONTAINER_PARENT_NOT_FOUND
	ERROR_TYPE_PERCENTAGE_OVER_1
	ERROR_TYPE_INTERNAL_ERROR
)

type ErrorData struct {
	ErrorType ErrorType
	ErrorText String
	UserData  any
}
type ErrorHandler struct {
	ErrorHandlerFunction func(errorText ErrorData)
	UserData             any
}

var LAYOUT_DEFAULT LayoutConfig = LayoutConfig{}
var __Color_DEFAULT Color = Color{}
var __CornerRadius_DEFAULT CornerRadius = CornerRadius{}
var __BorderWidth_DEFAULT BorderWidth = BorderWidth{}
var __currentContext *Context
var __defaultMaxElementCount int32 = 8192
var __defaultMaxMeasureTextWordCacheCount int32 = 16384

func __ErrorHandlerFunctionDefault(errorText ErrorData) {
	_ = errorText
}

var __SPACECHAR String = String{Length: 1, Chars: libc.CString(" ")}
var __STRING_DEFAULT String = String{}

type BooleanWarnings struct {
	MaxElementsExceeded           bool
	MaxRenderCommandsExceeded     bool
	MaxTextMeasureCacheExceeded   bool
	TextMeasurementFunctionNotSet bool
}
type __Warning struct {
	BaseMessage    String
	DynamicMessage String
}

var __WARNING_DEFAULT __Warning = __Warning{}

type __WarningArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__Warning
}
type SharedElementConfig struct {
	BackgroundColor Color
	CornerRadius    CornerRadius
	UserData        unsafe.Pointer
}
type __SharedElementConfigWrapper struct {
	Wrapped SharedElementConfig
}
type __boolArray struct {
	Capacity      int32
	Length        int32
	InternalArray *bool
}
type __boolArraySlice struct {
	Length        int32
	InternalArray *bool
}

var _Bool_DEFAULT bool = false

func __boolArray_Allocate_Arena(capacity int32, arena *Arena) __boolArray {
	return __boolArray{Capacity: capacity, Length: 0, InternalArray: (*bool)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(bool(false))), arena))}
}
func __boolArray_Get(array *__boolArray, index int32) *bool {
	if __Array_RangeCheck(index, array.Length) {
		return (*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
	}
	return &_Bool_DEFAULT
}
func __boolArray_GetValue(array *__boolArray, index int32) bool {
	if __Array_RangeCheck(index, array.Length) {
		return *(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
	}
	return _Bool_DEFAULT
}
func __boolArray_Add(array *__boolArray, item bool) *bool {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}())) = item
		return (*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), array.Length-1))
	}
	return &_Bool_DEFAULT
}
func __boolArraySlice_Get(slice *__boolArraySlice, index int32) *bool {
	if __Array_RangeCheck(index, slice.Length) {
		return (*bool)(unsafe.Add(unsafe.Pointer(slice.InternalArray), index))
	}
	return &_Bool_DEFAULT
}
func __boolArray_RemoveSwapback(array *__boolArray, index int32) bool {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed bool = *(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
		*(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = *(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), array.Length))
		return removed
	}
	return _Bool_DEFAULT
}
func __boolArray_Set(array *__boolArray, index int32, value bool) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __int32_tArray struct {
	Capacity      int32
	Length        int32
	InternalArray *int32
}
type __int32_tArraySlice struct {
	Length        int32
	InternalArray *int32
}

var int32_t_DEFAULT int32 = 0

func __int32_tArray_Allocate_Arena(capacity int32, arena *Arena) __int32_tArray {
	return __int32_tArray{Capacity: capacity, Length: 0, InternalArray: (*int32)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(int32(0))), arena))}
}
func __int32_tArray_Get(array *__int32_tArray, index int32) *int32 {
	if __Array_RangeCheck(index, array.Length) {
		return (*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
	}
	return &int32_t_DEFAULT
}
func __int32_tArray_GetValue(array *__int32_tArray, index int32) int32 {
	if __Array_RangeCheck(index, array.Length) {
		return *(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
	}
	return int32_t_DEFAULT
}
func __int32_tArray_Add(array *__int32_tArray, item int32) *int32 {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(array.Length-1)))
	}
	return &int32_t_DEFAULT
}
func __int32_tArraySlice_Get(slice *__int32_tArraySlice, index int32) *int32 {
	if __Array_RangeCheck(index, slice.Length) {
		return (*int32)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
	}
	return &int32_t_DEFAULT
}
func __int32_tArray_RemoveSwapback(array *__int32_tArray, index int32) int32 {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed int32 = *(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
		*(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index))) = *(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(array.Length)))
		return removed
	}
	return int32_t_DEFAULT
}
func __int32_tArray_Set(array *__int32_tArray, index int32, value int32) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __charArray struct {
	Capacity      int32
	Length        int32
	InternalArray *byte
}
type __charArraySlice struct {
	Length        int32
	InternalArray *byte
}

var char_DEFAULT int8 = 0

func __charArray_Allocate_Arena(capacity int32, arena *Arena) __charArray {
	return __charArray{Capacity: capacity, Length: 0, InternalArray: (*byte)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(int8(0))), arena))}
}
func __charArray_Get(array *__charArray, index int32) *byte {
	if __Array_RangeCheck(index, array.Length) {
		return (*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
	}
	return (*byte)(unsafe.Pointer(&char_DEFAULT))
}
func __charArray_GetValue(array *__charArray, index int32) int8 {
	if __Array_RangeCheck(index, array.Length) {
		return int8(*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)))
	}
	return char_DEFAULT
}
func __charArray_Add(array *__charArray, item int8) *byte {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}())) = byte(item)
		return (*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), array.Length-1))
	}
	return (*byte)(unsafe.Pointer(&char_DEFAULT))
}
func __charArraySlice_Get(slice *__charArraySlice, index int32) *byte {
	if __Array_RangeCheck(index, slice.Length) {
		return (*byte)(unsafe.Add(unsafe.Pointer(slice.InternalArray), index))
	}
	return (*byte)(unsafe.Pointer(&char_DEFAULT))
}
func __charArray_RemoveSwapback(array *__charArray, index int32) int8 {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)))
		*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = *(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), array.Length))
		return removed
	}
	return char_DEFAULT
}
func __charArray_Set(array *__charArray, index int32, value int8) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = byte(value)
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __ElementIdArray struct {
	Capacity      int32
	Length        int32
	InternalArray *ElementId
}
type __ElementIdArraySlice struct {
	Length        int32
	InternalArray *ElementId
}

var ElementId_DEFAULT ElementId = ElementId{}

func __ElementIdArray_Allocate_Arena(capacity int32, arena *Arena) __ElementIdArray {
	return __ElementIdArray{Capacity: capacity, Length: 0, InternalArray: (*ElementId)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(ElementId{})), arena))}
}
func __ElementIdArray_Get(array *__ElementIdArray, index int32) *ElementId {
	if __Array_RangeCheck(index, array.Length) {
		return (*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(index)))
	}
	return &ElementId_DEFAULT
}
func __ElementIdArray_GetValue(array *__ElementIdArray, index int32) ElementId {
	if __Array_RangeCheck(index, array.Length) {
		return *(*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(index)))
	}
	return ElementId_DEFAULT
}
func __ElementIdArray_Add(array *__ElementIdArray, item ElementId) *ElementId {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(array.Length-1)))
	}
	return &ElementId_DEFAULT
}
func __ElementIdArraySlice_Get(slice *__ElementIdArraySlice, index int32) *ElementId {
	if __Array_RangeCheck(index, slice.Length) {
		return (*ElementId)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(index)))
	}
	return &ElementId_DEFAULT
}
func __ElementIdArray_RemoveSwapback(array *__ElementIdArray, index int32) ElementId {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed ElementId = *(*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(index)))
		*(*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(index))) = *(*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(array.Length)))
		return removed
	}
	return ElementId_DEFAULT
}
func __ElementIdArray_Set(array *__ElementIdArray, index int32, value ElementId) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementId{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __LayoutConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *LayoutConfig
}
type __LayoutConfigArraySlice struct {
	Length        int32
	InternalArray *LayoutConfig
}

var LayoutConfig_DEFAULT LayoutConfig = LayoutConfig{}

func __LayoutConfigArray_Allocate_Arena(capacity int32, arena *Arena) __LayoutConfigArray {
	return __LayoutConfigArray{Capacity: capacity, Length: 0, InternalArray: (*LayoutConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(LayoutConfig{})), arena))}
}
func __LayoutConfigArray_Get(array *__LayoutConfigArray, index int32) *LayoutConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(index)))
	}
	return &LayoutConfig_DEFAULT
}
func __LayoutConfigArray_GetValue(array *__LayoutConfigArray, index int32) LayoutConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(index)))
	}
	return LayoutConfig_DEFAULT
}
func __LayoutConfigArray_Add(array *__LayoutConfigArray, item LayoutConfig) *LayoutConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(array.Length-1)))
	}
	return &LayoutConfig_DEFAULT
}
func __LayoutConfigArraySlice_Get(slice *__LayoutConfigArraySlice, index int32) *LayoutConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*LayoutConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(index)))
	}
	return &LayoutConfig_DEFAULT
}
func __LayoutConfigArray_RemoveSwapback(array *__LayoutConfigArray, index int32) LayoutConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed LayoutConfig = *(*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(index)))
		*(*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(index))) = *(*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(array.Length)))
		return removed
	}
	return LayoutConfig_DEFAULT
}
func __LayoutConfigArray_Set(array *__LayoutConfigArray, index int32, value LayoutConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __TextElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *TextElementConfig
}
type __TextElementConfigArraySlice struct {
	Length        int32
	InternalArray *TextElementConfig
}

var TextElementConfig_DEFAULT TextElementConfig = TextElementConfig{}

func __TextElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __TextElementConfigArray {
	return __TextElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*TextElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(TextElementConfig{})), arena))}
}
func __TextElementConfigArray_Get(array *__TextElementConfigArray, index int32) *TextElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(index)))
	}
	return &TextElementConfig_DEFAULT
}
func __TextElementConfigArray_GetValue(array *__TextElementConfigArray, index int32) TextElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(index)))
	}
	return TextElementConfig_DEFAULT
}
func __TextElementConfigArray_Add(array *__TextElementConfigArray, item TextElementConfig) *TextElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(array.Length-1)))
	}
	return &TextElementConfig_DEFAULT
}
func __TextElementConfigArraySlice_Get(slice *__TextElementConfigArraySlice, index int32) *TextElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*TextElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(index)))
	}
	return &TextElementConfig_DEFAULT
}
func __TextElementConfigArray_RemoveSwapback(array *__TextElementConfigArray, index int32) TextElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed TextElementConfig = *(*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(index)))
		*(*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(index))) = *(*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return TextElementConfig_DEFAULT
}
func __TextElementConfigArray_Set(array *__TextElementConfigArray, index int32, value TextElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(TextElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __ImageElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *ImageElementConfig
}
type __ImageElementConfigArraySlice struct {
	Length        int32
	InternalArray *ImageElementConfig
}

var ImageElementConfig_DEFAULT ImageElementConfig = ImageElementConfig{}

func __ImageElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __ImageElementConfigArray {
	return __ImageElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*ImageElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(ImageElementConfig{})), arena))}
}
func __ImageElementConfigArray_Get(array *__ImageElementConfigArray, index int32) *ImageElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(index)))
	}
	return &ImageElementConfig_DEFAULT
}
func __ImageElementConfigArray_GetValue(array *__ImageElementConfigArray, index int32) ImageElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(index)))
	}
	return ImageElementConfig_DEFAULT
}
func __ImageElementConfigArray_Add(array *__ImageElementConfigArray, item ImageElementConfig) *ImageElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(array.Length-1)))
	}
	return &ImageElementConfig_DEFAULT
}
func __ImageElementConfigArraySlice_Get(slice *__ImageElementConfigArraySlice, index int32) *ImageElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*ImageElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(index)))
	}
	return &ImageElementConfig_DEFAULT
}
func __ImageElementConfigArray_RemoveSwapback(array *__ImageElementConfigArray, index int32) ImageElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed ImageElementConfig = *(*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(index)))
		*(*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(index))) = *(*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return ImageElementConfig_DEFAULT
}
func __ImageElementConfigArray_Set(array *__ImageElementConfigArray, index int32, value ImageElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ImageElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __FloatingElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *FloatingElementConfig
}
type __FloatingElementConfigArraySlice struct {
	Length        int32
	InternalArray *FloatingElementConfig
}

var FloatingElementConfig_DEFAULT FloatingElementConfig = FloatingElementConfig{}

func __FloatingElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __FloatingElementConfigArray {
	return __FloatingElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*FloatingElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(FloatingElementConfig{})), arena))}
}
func __FloatingElementConfigArray_Get(array *__FloatingElementConfigArray, index int32) *FloatingElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(index)))
	}
	return &FloatingElementConfig_DEFAULT
}
func __FloatingElementConfigArray_GetValue(array *__FloatingElementConfigArray, index int32) FloatingElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(index)))
	}
	return FloatingElementConfig_DEFAULT
}
func __FloatingElementConfigArray_Add(array *__FloatingElementConfigArray, item FloatingElementConfig) *FloatingElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(array.Length-1)))
	}
	return &FloatingElementConfig_DEFAULT
}
func __FloatingElementConfigArraySlice_Get(slice *__FloatingElementConfigArraySlice, index int32) *FloatingElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(index)))
	}
	return &FloatingElementConfig_DEFAULT
}
func __FloatingElementConfigArray_RemoveSwapback(array *__FloatingElementConfigArray, index int32) FloatingElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed FloatingElementConfig = *(*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(index)))
		*(*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(index))) = *(*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return FloatingElementConfig_DEFAULT
}
func __FloatingElementConfigArray_Set(array *__FloatingElementConfigArray, index int32, value FloatingElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(FloatingElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __CustomElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *CustomElementConfig
}
type __CustomElementConfigArraySlice struct {
	Length        int32
	InternalArray *CustomElementConfig
}

var CustomElementConfig_DEFAULT CustomElementConfig = CustomElementConfig{}

func __CustomElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __CustomElementConfigArray {
	return __CustomElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*CustomElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(CustomElementConfig{})), arena))}
}
func __CustomElementConfigArray_Get(array *__CustomElementConfigArray, index int32) *CustomElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(index)))
	}
	return &CustomElementConfig_DEFAULT
}
func __CustomElementConfigArray_GetValue(array *__CustomElementConfigArray, index int32) CustomElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(index)))
	}
	return CustomElementConfig_DEFAULT
}
func __CustomElementConfigArray_Add(array *__CustomElementConfigArray, item CustomElementConfig) *CustomElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(array.Length-1)))
	}
	return &CustomElementConfig_DEFAULT
}
func __CustomElementConfigArraySlice_Get(slice *__CustomElementConfigArraySlice, index int32) *CustomElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*CustomElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(index)))
	}
	return &CustomElementConfig_DEFAULT
}
func __CustomElementConfigArray_RemoveSwapback(array *__CustomElementConfigArray, index int32) CustomElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed CustomElementConfig = *(*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(index)))
		*(*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(index))) = *(*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return CustomElementConfig_DEFAULT
}
func __CustomElementConfigArray_Set(array *__CustomElementConfigArray, index int32, value CustomElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(CustomElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __ScrollElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *ScrollElementConfig
}
type __ScrollElementConfigArraySlice struct {
	Length        int32
	InternalArray *ScrollElementConfig
}

var ScrollElementConfig_DEFAULT ScrollElementConfig = ScrollElementConfig{Horizontal: false}

func __ScrollElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __ScrollElementConfigArray {
	return __ScrollElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*ScrollElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(ScrollElementConfig{})), arena))}
}
func __ScrollElementConfigArray_Get(array *__ScrollElementConfigArray, index int32) *ScrollElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(index)))
	}
	return &ScrollElementConfig_DEFAULT
}
func __ScrollElementConfigArray_GetValue(array *__ScrollElementConfigArray, index int32) ScrollElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(index)))
	}
	return ScrollElementConfig_DEFAULT
}
func __ScrollElementConfigArray_Add(array *__ScrollElementConfigArray, item ScrollElementConfig) *ScrollElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(array.Length-1)))
	}
	return &ScrollElementConfig_DEFAULT
}
func __ScrollElementConfigArraySlice_Get(slice *__ScrollElementConfigArraySlice, index int32) *ScrollElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(index)))
	}
	return &ScrollElementConfig_DEFAULT
}
func __ScrollElementConfigArray_RemoveSwapback(array *__ScrollElementConfigArray, index int32) ScrollElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed ScrollElementConfig = *(*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(index)))
		*(*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(index))) = *(*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return ScrollElementConfig_DEFAULT
}
func __ScrollElementConfigArray_Set(array *__ScrollElementConfigArray, index int32, value ScrollElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ScrollElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __BorderElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *BorderElementConfig
}
type __BorderElementConfigArraySlice struct {
	Length        int32
	InternalArray *BorderElementConfig
}

var BorderElementConfig_DEFAULT BorderElementConfig = BorderElementConfig{}

func __BorderElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __BorderElementConfigArray {
	return __BorderElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*BorderElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(BorderElementConfig{})), arena))}
}
func __BorderElementConfigArray_Get(array *__BorderElementConfigArray, index int32) *BorderElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(index)))
	}
	return &BorderElementConfig_DEFAULT
}
func __BorderElementConfigArray_GetValue(array *__BorderElementConfigArray, index int32) BorderElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(index)))
	}
	return BorderElementConfig_DEFAULT
}
func __BorderElementConfigArray_Add(array *__BorderElementConfigArray, item BorderElementConfig) *BorderElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(array.Length-1)))
	}
	return &BorderElementConfig_DEFAULT
}
func __BorderElementConfigArraySlice_Get(slice *__BorderElementConfigArraySlice, index int32) *BorderElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*BorderElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(index)))
	}
	return &BorderElementConfig_DEFAULT
}
func __BorderElementConfigArray_RemoveSwapback(array *__BorderElementConfigArray, index int32) BorderElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed BorderElementConfig = *(*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(index)))
		*(*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(index))) = *(*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return BorderElementConfig_DEFAULT
}
func __BorderElementConfigArray_Set(array *__BorderElementConfigArray, index int32, value BorderElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(BorderElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __StringArray struct {
	Capacity      int32
	Length        int32
	InternalArray *String
}
type __StringArraySlice struct {
	Length        int32
	InternalArray *String
}

var String_DEFAULT String = String{}

func __StringArray_Allocate_Arena(capacity int32, arena *Arena) __StringArray {
	return __StringArray{Capacity: capacity, Length: 0, InternalArray: (*String)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(String{})), arena))}
}
func __StringArray_Get(array *__StringArray, index int32) *String {
	if __Array_RangeCheck(index, array.Length) {
		return (*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(index)))
	}
	return &String_DEFAULT
}
func __StringArray_GetValue(array *__StringArray, index int32) String {
	if __Array_RangeCheck(index, array.Length) {
		return *(*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(index)))
	}
	return String_DEFAULT
}
func __StringArray_Add(array *__StringArray, item String) *String {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(array.Length-1)))
	}
	return &String_DEFAULT
}
func __StringArraySlice_Get(slice *__StringArraySlice, index int32) *String {
	if __Array_RangeCheck(index, slice.Length) {
		return (*String)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(String{})*uintptr(index)))
	}
	return &String_DEFAULT
}
func __StringArray_RemoveSwapback(array *__StringArray, index int32) String {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed String = *(*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(index)))
		*(*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(index))) = *(*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(array.Length)))
		return removed
	}
	return String_DEFAULT
}
func __StringArray_Set(array *__StringArray, index int32, value String) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(String{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __SharedElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *SharedElementConfig
}
type __SharedElementConfigArraySlice struct {
	Length        int32
	InternalArray *SharedElementConfig
}

var SharedElementConfig_DEFAULT SharedElementConfig = SharedElementConfig{}

func __SharedElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __SharedElementConfigArray {
	return __SharedElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*SharedElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(SharedElementConfig{})), arena))}
}
func __SharedElementConfigArray_Get(array *__SharedElementConfigArray, index int32) *SharedElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(index)))
	}
	return &SharedElementConfig_DEFAULT
}
func __SharedElementConfigArray_GetValue(array *__SharedElementConfigArray, index int32) SharedElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(index)))
	}
	return SharedElementConfig_DEFAULT
}
func __SharedElementConfigArray_Add(array *__SharedElementConfigArray, item SharedElementConfig) *SharedElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(array.Length-1)))
	}
	return &SharedElementConfig_DEFAULT
}
func __SharedElementConfigArraySlice_Get(slice *__SharedElementConfigArraySlice, index int32) *SharedElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*SharedElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(index)))
	}
	return &SharedElementConfig_DEFAULT
}
func __SharedElementConfigArray_RemoveSwapback(array *__SharedElementConfigArray, index int32) SharedElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed SharedElementConfig = *(*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(index)))
		*(*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(index))) = *(*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return SharedElementConfig_DEFAULT
}
func __SharedElementConfigArray_Set(array *__SharedElementConfigArray, index int32, value SharedElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(SharedElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type RenderCommandArraySlice struct {
	Length        int32
	InternalArray *RenderCommand
}

var RenderCommand_DEFAULT RenderCommand = RenderCommand{}

func RenderCommandArray_Allocate_Arena(capacity int32, arena *Arena) RenderCommandArray {
	return RenderCommandArray{Capacity: capacity, Length: 0, InternalArray: (*RenderCommand)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(RenderCommand{})), arena))}
}
func RenderCommandArray_Get(array *RenderCommandArray, index int32) *RenderCommand {
	if __Array_RangeCheck(index, array.Length) {
		return (*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(index)))
	}
	return &RenderCommand_DEFAULT
}
func RenderCommandArray_GetValue(array *RenderCommandArray, index int32) RenderCommand {
	if __Array_RangeCheck(index, array.Length) {
		return *(*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(index)))
	}
	return RenderCommand_DEFAULT
}
func RenderCommandArray_Add(array *RenderCommandArray, item RenderCommand) *RenderCommand {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(array.Length-1)))
	}
	return &RenderCommand_DEFAULT
}
func RenderCommandArraySlice_Get(slice *RenderCommandArraySlice, index int32) *RenderCommand {
	if __Array_RangeCheck(index, slice.Length) {
		return (*RenderCommand)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(index)))
	}
	return &RenderCommand_DEFAULT
}
func RenderCommandArray_RemoveSwapback(array *RenderCommandArray, index int32) RenderCommand {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed RenderCommand = *(*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(index)))
		*(*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(index))) = *(*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(array.Length)))
		return removed
	}
	return RenderCommand_DEFAULT
}
func RenderCommandArray_Set(array *RenderCommandArray, index int32, value RenderCommand) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(RenderCommand{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __ElementConfigType int32

const (
	__ELEMENT_CONFIG_TYPE_NONE = __ElementConfigType(iota)
	__ELEMENT_CONFIG_TYPE_BORDER
	__ELEMENT_CONFIG_TYPE_FLOATING
	__ELEMENT_CONFIG_TYPE_SCROLL
	__ELEMENT_CONFIG_TYPE_IMAGE
	__ELEMENT_CONFIG_TYPE_TEXT
	__ELEMENT_CONFIG_TYPE_CUSTOM
	__ELEMENT_CONFIG_TYPE_SHARED
)

type ElementConfigUnion struct {
	// union
	TextElementConfig     *TextElementConfig
	ImageElementConfig    *ImageElementConfig
	FloatingElementConfig *FloatingElementConfig
	CustomElementConfig   *CustomElementConfig
	ScrollElementConfig   *ScrollElementConfig
	BorderElementConfig   *BorderElementConfig
	SharedElementConfig   *SharedElementConfig
}
type ElementConfig struct {
	Type   __ElementConfigType
	Config ElementConfigUnion
}
type __ElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *ElementConfig
}
type __ElementConfigArraySlice struct {
	Length        int32
	InternalArray *ElementConfig
}

var ElementConfig_DEFAULT ElementConfig = ElementConfig{}

func __ElementConfigArray_Allocate_Arena(capacity int32, arena *Arena) __ElementConfigArray {
	return __ElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*ElementConfig)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(ElementConfig{})), arena))}
}
func __ElementConfigArray_Get(array *__ElementConfigArray, index int32) *ElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return (*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(index)))
	}
	return &ElementConfig_DEFAULT
}
func __ElementConfigArray_GetValue(array *__ElementConfigArray, index int32) ElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		return *(*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(index)))
	}
	return ElementConfig_DEFAULT
}
func __ElementConfigArray_Add(array *__ElementConfigArray, item ElementConfig) *ElementConfig {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(array.Length-1)))
	}
	return &ElementConfig_DEFAULT
}
func __ElementConfigArraySlice_Get(slice *__ElementConfigArraySlice, index int32) *ElementConfig {
	if __Array_RangeCheck(index, slice.Length) {
		return (*ElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(index)))
	}
	return &ElementConfig_DEFAULT
}
func __ElementConfigArray_RemoveSwapback(array *__ElementConfigArray, index int32) ElementConfig {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed ElementConfig = *(*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(index)))
		*(*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(index))) = *(*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return ElementConfig_DEFAULT
}
func __ElementConfigArray_Set(array *__ElementConfigArray, index int32, value ElementConfig) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __WrappedTextLine struct {
	Dimensions Dimensions
	Line       String
}
type __WrappedTextLineArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__WrappedTextLine
}
type __WrappedTextLineArraySlice struct {
	Length        int32
	InternalArray *__WrappedTextLine
}

var __WrappedTextLine_DEFAULT __WrappedTextLine = __WrappedTextLine{}

func __WrappedTextLineArray_Allocate_Arena(capacity int32, arena *Arena) __WrappedTextLineArray {
	return __WrappedTextLineArray{Capacity: capacity, Length: 0, InternalArray: (*__WrappedTextLine)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__WrappedTextLine{})), arena))}
}
func __WrappedTextLineArray_Get(array *__WrappedTextLineArray, index int32) *__WrappedTextLine {
	if __Array_RangeCheck(index, array.Length) {
		return (*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(index)))
	}
	return &__WrappedTextLine_DEFAULT
}
func __WrappedTextLineArray_GetValue(array *__WrappedTextLineArray, index int32) __WrappedTextLine {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(index)))
	}
	return __WrappedTextLine_DEFAULT
}
func __WrappedTextLineArray_Add(array *__WrappedTextLineArray, item __WrappedTextLine) *__WrappedTextLine {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(array.Length-1)))
	}
	return &__WrappedTextLine_DEFAULT
}
func __WrappedTextLineArraySlice_Get(slice *__WrappedTextLineArraySlice, index int32) *__WrappedTextLine {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(index)))
	}
	return &__WrappedTextLine_DEFAULT
}
func __WrappedTextLineArray_RemoveSwapback(array *__WrappedTextLineArray, index int32) __WrappedTextLine {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __WrappedTextLine = *(*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(index)))
		*(*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(index))) = *(*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(array.Length)))
		return removed
	}
	return __WrappedTextLine_DEFAULT
}
func __WrappedTextLineArray_Set(array *__WrappedTextLineArray, index int32, value __WrappedTextLine) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __TextElementData struct {
	Text                String
	PreferredDimensions Dimensions
	ElementIndex        int32
	WrappedLines        __WrappedTextLineArraySlice
}
type __TextElementDataArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__TextElementData
}
type __TextElementDataArraySlice struct {
	Length        int32
	InternalArray *__TextElementData
}

var __TextElementData_DEFAULT __TextElementData = __TextElementData{}

func __TextElementDataArray_Allocate_Arena(capacity int32, arena *Arena) __TextElementDataArray {
	return __TextElementDataArray{Capacity: capacity, Length: 0, InternalArray: (*__TextElementData)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__TextElementData{})), arena))}
}
func __TextElementDataArray_Get(array *__TextElementDataArray, index int32) *__TextElementData {
	if __Array_RangeCheck(index, array.Length) {
		return (*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(index)))
	}
	return &__TextElementData_DEFAULT
}
func __TextElementDataArray_GetValue(array *__TextElementDataArray, index int32) __TextElementData {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(index)))
	}
	return __TextElementData_DEFAULT
}
func __TextElementDataArray_Add(array *__TextElementDataArray, item __TextElementData) *__TextElementData {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(array.Length-1)))
	}
	return &__TextElementData_DEFAULT
}
func __TextElementDataArraySlice_Get(slice *__TextElementDataArraySlice, index int32) *__TextElementData {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__TextElementData)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(index)))
	}
	return &__TextElementData_DEFAULT
}
func __TextElementDataArray_RemoveSwapback(array *__TextElementDataArray, index int32) __TextElementData {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __TextElementData = *(*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(index)))
		*(*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(index))) = *(*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(array.Length)))
		return removed
	}
	return __TextElementData_DEFAULT
}
func __TextElementDataArray_Set(array *__TextElementDataArray, index int32, value __TextElementData) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__TextElementData{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __LayoutElementChildren struct {
	Elements *int32
	Length   uint16
}
type LayoutElement struct {
	ChildrenOrTextContent struct {
		// union
		Children        __LayoutElementChildren
		TextElementData *__TextElementData
	}
	Dimensions     Dimensions
	MinDimensions  Dimensions
	LayoutConfig   *LayoutConfig
	ElementConfigs __ElementConfigArraySlice
	Id             uint32
}
type LayoutElementArray struct {
	Capacity      int32
	Length        int32
	InternalArray *LayoutElement
}
type LayoutElementArraySlice struct {
	Length        int32
	InternalArray *LayoutElement
}

var LayoutElement_DEFAULT LayoutElement = LayoutElement{}

func LayoutElementArray_Allocate_Arena(capacity int32, arena *Arena) LayoutElementArray {
	return LayoutElementArray{Capacity: capacity, Length: 0, InternalArray: (*LayoutElement)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(LayoutElement{})), arena))}
}
func LayoutElementArray_Get(array *LayoutElementArray, index int32) *LayoutElement {
	if __Array_RangeCheck(index, array.Length) {
		return (*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(index)))
	}
	return &LayoutElement_DEFAULT
}
func LayoutElementArray_GetValue(array *LayoutElementArray, index int32) LayoutElement {
	if __Array_RangeCheck(index, array.Length) {
		return *(*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(index)))
	}
	return LayoutElement_DEFAULT
}
func LayoutElementArray_Add(array *LayoutElementArray, item LayoutElement) *LayoutElement {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(array.Length-1)))
	}
	return &LayoutElement_DEFAULT
}
func LayoutElementArraySlice_Get(slice *LayoutElementArraySlice, index int32) *LayoutElement {
	if __Array_RangeCheck(index, slice.Length) {
		return (*LayoutElement)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(index)))
	}
	return &LayoutElement_DEFAULT
}
func LayoutElementArray_RemoveSwapback(array *LayoutElementArray, index int32) LayoutElement {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed LayoutElement = *(*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(index)))
		*(*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(index))) = *(*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(array.Length)))
		return removed
	}
	return LayoutElement_DEFAULT
}
func LayoutElementArray_Set(array *LayoutElementArray, index int32, value LayoutElement) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElement{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __ScrollContainerDataInternal struct {
	LayoutElement       *LayoutElement
	BoundingBox         BoundingBox
	ContentSize         Dimensions
	ScrollOrigin        Vector2
	PointerOrigin       Vector2
	ScrollMomentum      Vector2
	ScrollPosition      Vector2
	PreviousDelta       Vector2
	MomentumTime        float32
	ElementId           uint32
	OpenThisFrame       bool
	PointerScrollActive bool
}
type __ScrollContainerDataInternalArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__ScrollContainerDataInternal
}
type __ScrollContainerDataInternalArraySlice struct {
	Length        int32
	InternalArray *__ScrollContainerDataInternal
}

var __ScrollContainerDataInternal_DEFAULT __ScrollContainerDataInternal = __ScrollContainerDataInternal{}

func __ScrollContainerDataInternalArray_Allocate_Arena(capacity int32, arena *Arena) __ScrollContainerDataInternalArray {
	return __ScrollContainerDataInternalArray{Capacity: capacity, Length: 0, InternalArray: (*__ScrollContainerDataInternal)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__ScrollContainerDataInternal{})), arena))}
}
func __ScrollContainerDataInternalArray_Get(array *__ScrollContainerDataInternalArray, index int32) *__ScrollContainerDataInternal {
	if __Array_RangeCheck(index, array.Length) {
		return (*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(index)))
	}
	return &__ScrollContainerDataInternal_DEFAULT
}
func __ScrollContainerDataInternalArray_GetValue(array *__ScrollContainerDataInternalArray, index int32) __ScrollContainerDataInternal {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(index)))
	}
	return __ScrollContainerDataInternal_DEFAULT
}
func __ScrollContainerDataInternalArray_Add(array *__ScrollContainerDataInternalArray, item __ScrollContainerDataInternal) *__ScrollContainerDataInternal {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(array.Length-1)))
	}
	return &__ScrollContainerDataInternal_DEFAULT
}
func __ScrollContainerDataInternalArraySlice_Get(slice *__ScrollContainerDataInternalArraySlice, index int32) *__ScrollContainerDataInternal {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(index)))
	}
	return &__ScrollContainerDataInternal_DEFAULT
}
func __ScrollContainerDataInternalArray_RemoveSwapback(array *__ScrollContainerDataInternalArray, index int32) __ScrollContainerDataInternal {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __ScrollContainerDataInternal = *(*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(index)))
		*(*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(index))) = *(*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(array.Length)))
		return removed
	}
	return __ScrollContainerDataInternal_DEFAULT
}
func __ScrollContainerDataInternalArray_Set(array *__ScrollContainerDataInternalArray, index int32, value __ScrollContainerDataInternal) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__ScrollContainerDataInternal{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __DebugElementData struct {
	Collision bool
	Collapsed bool
}
type __DebugElementDataArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__DebugElementData
}
type __DebugElementDataArraySlice struct {
	Length        int32
	InternalArray *__DebugElementData
}

var __DebugElementData_DEFAULT __DebugElementData = __DebugElementData{Collision: false}

func __DebugElementDataArray_Allocate_Arena(capacity int32, arena *Arena) __DebugElementDataArray {
	return __DebugElementDataArray{Capacity: capacity, Length: 0, InternalArray: (*__DebugElementData)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__DebugElementData{})), arena))}
}
func __DebugElementDataArray_Get(array *__DebugElementDataArray, index int32) *__DebugElementData {
	if __Array_RangeCheck(index, array.Length) {
		return (*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(index)))
	}
	return &__DebugElementData_DEFAULT
}
func __DebugElementDataArray_GetValue(array *__DebugElementDataArray, index int32) __DebugElementData {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(index)))
	}
	return __DebugElementData_DEFAULT
}
func __DebugElementDataArray_Add(array *__DebugElementDataArray, item __DebugElementData) *__DebugElementData {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(array.Length-1)))
	}
	return &__DebugElementData_DEFAULT
}
func __DebugElementDataArraySlice_Get(slice *__DebugElementDataArraySlice, index int32) *__DebugElementData {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__DebugElementData)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(index)))
	}
	return &__DebugElementData_DEFAULT
}
func __DebugElementDataArray_RemoveSwapback(array *__DebugElementDataArray, index int32) __DebugElementData {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __DebugElementData = *(*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(index)))
		*(*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(index))) = *(*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(array.Length)))
		return removed
	}
	return __DebugElementData_DEFAULT
}
func __DebugElementDataArray_Set(array *__DebugElementDataArray, index int32, value __DebugElementData) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__DebugElementData{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type LayoutElementHashMapItem struct {
	BoundingBox           BoundingBox
	ElementId             ElementId
	LayoutElement         *LayoutElement
	OnHoverFunction       func(elementId ElementId, pointerInfo PointerData, userData int64)
	HoverFunctionUserData int64
	NextIndex             int32
	Generation            uint32
	IdAlias               uint32
	DebugData             *__DebugElementData
}
type __LayoutElementHashMapItemArray struct {
	Capacity      int32
	Length        int32
	InternalArray *LayoutElementHashMapItem
}
type __LayoutElementHashMapItemArraySlice struct {
	Length        int32
	InternalArray *LayoutElementHashMapItem
}

var LayoutElementHashMapItem_DEFAULT LayoutElementHashMapItem = LayoutElementHashMapItem{}

func __LayoutElementHashMapItemArray_Allocate_Arena(capacity int32, arena *Arena) __LayoutElementHashMapItemArray {
	return __LayoutElementHashMapItemArray{Capacity: capacity, Length: 0, InternalArray: (*LayoutElementHashMapItem)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(LayoutElementHashMapItem{})), arena))}
}
func __LayoutElementHashMapItemArray_Get(array *__LayoutElementHashMapItemArray, index int32) *LayoutElementHashMapItem {
	if __Array_RangeCheck(index, array.Length) {
		return (*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(index)))
	}
	return &LayoutElementHashMapItem_DEFAULT
}
func __LayoutElementHashMapItemArray_GetValue(array *__LayoutElementHashMapItemArray, index int32) LayoutElementHashMapItem {
	if __Array_RangeCheck(index, array.Length) {
		return *(*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(index)))
	}
	return LayoutElementHashMapItem_DEFAULT
}
func __LayoutElementHashMapItemArray_Add(array *__LayoutElementHashMapItemArray, item LayoutElementHashMapItem) *LayoutElementHashMapItem {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(array.Length-1)))
	}
	return &LayoutElementHashMapItem_DEFAULT
}
func __LayoutElementHashMapItemArraySlice_Get(slice *__LayoutElementHashMapItemArraySlice, index int32) *LayoutElementHashMapItem {
	if __Array_RangeCheck(index, slice.Length) {
		return (*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(index)))
	}
	return &LayoutElementHashMapItem_DEFAULT
}
func __LayoutElementHashMapItemArray_RemoveSwapback(array *__LayoutElementHashMapItemArray, index int32) LayoutElementHashMapItem {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed LayoutElementHashMapItem = *(*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(index)))
		*(*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(index))) = *(*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(array.Length)))
		return removed
	}
	return LayoutElementHashMapItem_DEFAULT
}
func __LayoutElementHashMapItemArray_Set(array *__LayoutElementHashMapItemArray, index int32, value LayoutElementHashMapItem) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(LayoutElementHashMapItem{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __MeasuredWord struct {
	StartOffset int32
	Length      int32
	Width       float32
	Next        int32
}
type __MeasuredWordArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__MeasuredWord
}
type __MeasuredWordArraySlice struct {
	Length        int32
	InternalArray *__MeasuredWord
}

var __MeasuredWord_DEFAULT __MeasuredWord = __MeasuredWord{}

func __MeasuredWordArray_Allocate_Arena(capacity int32, arena *Arena) __MeasuredWordArray {
	return __MeasuredWordArray{Capacity: capacity, Length: 0, InternalArray: (*__MeasuredWord)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__MeasuredWord{})), arena))}
}
func __MeasuredWordArray_Get(array *__MeasuredWordArray, index int32) *__MeasuredWord {
	if __Array_RangeCheck(index, array.Length) {
		return (*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(index)))
	}
	return &__MeasuredWord_DEFAULT
}
func __MeasuredWordArray_GetValue(array *__MeasuredWordArray, index int32) __MeasuredWord {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(index)))
	}
	return __MeasuredWord_DEFAULT
}
func __MeasuredWordArray_Add(array *__MeasuredWordArray, item __MeasuredWord) *__MeasuredWord {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(array.Length-1)))
	}
	return &__MeasuredWord_DEFAULT
}
func __MeasuredWordArraySlice_Get(slice *__MeasuredWordArraySlice, index int32) *__MeasuredWord {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__MeasuredWord)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(index)))
	}
	return &__MeasuredWord_DEFAULT
}
func __MeasuredWordArray_RemoveSwapback(array *__MeasuredWordArray, index int32) __MeasuredWord {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __MeasuredWord = *(*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(index)))
		*(*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(index))) = *(*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(array.Length)))
		return removed
	}
	return __MeasuredWord_DEFAULT
}
func __MeasuredWordArray_Set(array *__MeasuredWordArray, index int32, value __MeasuredWord) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasuredWord{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __MeasureTextCacheItem struct {
	UnwrappedDimensions     Dimensions
	MeasuredWordsStartIndex int32
	ContainsNewlines        bool
	Id                      uint32
	NextIndex               int32
	Generation              uint32
}
type __MeasureTextCacheItemArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__MeasureTextCacheItem
}
type __MeasureTextCacheItemArraySlice struct {
	Length        int32
	InternalArray *__MeasureTextCacheItem
}

var __MeasureTextCacheItem_DEFAULT __MeasureTextCacheItem = __MeasureTextCacheItem{}

func __MeasureTextCacheItemArray_Allocate_Arena(capacity int32, arena *Arena) __MeasureTextCacheItemArray {
	return __MeasureTextCacheItemArray{Capacity: capacity, Length: 0, InternalArray: (*__MeasureTextCacheItem)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__MeasureTextCacheItem{})), arena))}
}
func __MeasureTextCacheItemArray_Get(array *__MeasureTextCacheItemArray, index int32) *__MeasureTextCacheItem {
	if __Array_RangeCheck(index, array.Length) {
		return (*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(index)))
	}
	return &__MeasureTextCacheItem_DEFAULT
}
func __MeasureTextCacheItemArray_GetValue(array *__MeasureTextCacheItemArray, index int32) __MeasureTextCacheItem {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(index)))
	}
	return __MeasureTextCacheItem_DEFAULT
}
func __MeasureTextCacheItemArray_Add(array *__MeasureTextCacheItemArray, item __MeasureTextCacheItem) *__MeasureTextCacheItem {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(array.Length-1)))
	}
	return &__MeasureTextCacheItem_DEFAULT
}
func __MeasureTextCacheItemArraySlice_Get(slice *__MeasureTextCacheItemArraySlice, index int32) *__MeasureTextCacheItem {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(index)))
	}
	return &__MeasureTextCacheItem_DEFAULT
}
func __MeasureTextCacheItemArray_RemoveSwapback(array *__MeasureTextCacheItemArray, index int32) __MeasureTextCacheItem {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __MeasureTextCacheItem = *(*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(index)))
		*(*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(index))) = *(*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(array.Length)))
		return removed
	}
	return __MeasureTextCacheItem_DEFAULT
}
func __MeasureTextCacheItemArray_Set(array *__MeasureTextCacheItemArray, index int32, value __MeasureTextCacheItem) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__MeasureTextCacheItem{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __LayoutElementTreeNode struct {
	LayoutElement   *LayoutElement
	Position        Vector2
	NextChildOffset Vector2
}
type __LayoutElementTreeNodeArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__LayoutElementTreeNode
}
type __LayoutElementTreeNodeArraySlice struct {
	Length        int32
	InternalArray *__LayoutElementTreeNode
}

var __LayoutElementTreeNode_DEFAULT __LayoutElementTreeNode = __LayoutElementTreeNode{}

func __LayoutElementTreeNodeArray_Allocate_Arena(capacity int32, arena *Arena) __LayoutElementTreeNodeArray {
	return __LayoutElementTreeNodeArray{Capacity: capacity, Length: 0, InternalArray: (*__LayoutElementTreeNode)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__LayoutElementTreeNode{})), arena))}
}
func __LayoutElementTreeNodeArray_Get(array *__LayoutElementTreeNodeArray, index int32) *__LayoutElementTreeNode {
	if __Array_RangeCheck(index, array.Length) {
		return (*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(index)))
	}
	return &__LayoutElementTreeNode_DEFAULT
}
func __LayoutElementTreeNodeArray_GetValue(array *__LayoutElementTreeNodeArray, index int32) __LayoutElementTreeNode {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(index)))
	}
	return __LayoutElementTreeNode_DEFAULT
}
func __LayoutElementTreeNodeArray_Add(array *__LayoutElementTreeNodeArray, item __LayoutElementTreeNode) *__LayoutElementTreeNode {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(array.Length-1)))
	}
	return &__LayoutElementTreeNode_DEFAULT
}
func __LayoutElementTreeNodeArraySlice_Get(slice *__LayoutElementTreeNodeArraySlice, index int32) *__LayoutElementTreeNode {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(index)))
	}
	return &__LayoutElementTreeNode_DEFAULT
}
func __LayoutElementTreeNodeArray_RemoveSwapback(array *__LayoutElementTreeNodeArray, index int32) __LayoutElementTreeNode {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __LayoutElementTreeNode = *(*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(index)))
		*(*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(index))) = *(*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(array.Length)))
		return removed
	}
	return __LayoutElementTreeNode_DEFAULT
}
func __LayoutElementTreeNodeArray_Set(array *__LayoutElementTreeNodeArray, index int32, value __LayoutElementTreeNode) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}

type __LayoutElementTreeRoot struct {
	LayoutElementIndex int32
	ParentId           uint32
	ClipElementId      uint32
	ZIndex             int16
	PointerOffset      Vector2
}
type __LayoutElementTreeRootArray struct {
	Capacity      int32
	Length        int32
	InternalArray *__LayoutElementTreeRoot
}
type __LayoutElementTreeRootArraySlice struct {
	Length        int32
	InternalArray *__LayoutElementTreeRoot
}

var __LayoutElementTreeRoot_DEFAULT __LayoutElementTreeRoot = __LayoutElementTreeRoot{}

func __LayoutElementTreeRootArray_Allocate_Arena(capacity int32, arena *Arena) __LayoutElementTreeRootArray {
	return __LayoutElementTreeRootArray{Capacity: capacity, Length: 0, InternalArray: (*__LayoutElementTreeRoot)(__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(__LayoutElementTreeRoot{})), arena))}
}
func __LayoutElementTreeRootArray_Get(array *__LayoutElementTreeRootArray, index int32) *__LayoutElementTreeRoot {
	if __Array_RangeCheck(index, array.Length) {
		return (*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(index)))
	}
	return &__LayoutElementTreeRoot_DEFAULT
}
func __LayoutElementTreeRootArray_GetValue(array *__LayoutElementTreeRootArray, index int32) __LayoutElementTreeRoot {
	if __Array_RangeCheck(index, array.Length) {
		return *(*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(index)))
	}
	return __LayoutElementTreeRoot_DEFAULT
}
func __LayoutElementTreeRootArray_Add(array *__LayoutElementTreeRootArray, item __LayoutElementTreeRoot) *__LayoutElementTreeRoot {
	if __Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(array.Length-1)))
	}
	return &__LayoutElementTreeRoot_DEFAULT
}
func __LayoutElementTreeRootArraySlice_Get(slice *__LayoutElementTreeRootArraySlice, index int32) *__LayoutElementTreeRoot {
	if __Array_RangeCheck(index, slice.Length) {
		return (*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(index)))
	}
	return &__LayoutElementTreeRoot_DEFAULT
}
func __LayoutElementTreeRootArray_RemoveSwapback(array *__LayoutElementTreeRootArray, index int32) __LayoutElementTreeRoot {
	if __Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed __LayoutElementTreeRoot = *(*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(index)))
		*(*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(index))) = *(*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(array.Length)))
		return removed
	}
	return __LayoutElementTreeRoot_DEFAULT
}
func __LayoutElementTreeRootArray_Set(array *__LayoutElementTreeRootArray, index int32, value __LayoutElementTreeRoot) {
	if __Array_RangeCheck(index, array.Capacity) {
		*(*__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__LayoutElementTreeRoot{})*uintptr(index))) = value
		if index < array.Length {
			array.Length = array.Length
		} else {
			array.Length = index + 1
		}
	}
}
func __Context_Allocate_Arena(arena *Arena) *Context {
	var (
		totalSizeBytes  uint64 = uint64(unsafe.Sizeof(Context{}))
		memoryAddress   uint64 = uint64(uintptr(unsafe.Pointer(arena.Memory)))
		nextAllocOffset uint64 = (memoryAddress % 64)
	)
	if nextAllocOffset+totalSizeBytes > arena.Capacity {
		return nil
	}
	arena.NextAllocation = nextAllocOffset + totalSizeBytes
	return (*Context)(unsafe.Pointer(uintptr(memoryAddress + nextAllocOffset)))
}
func __WriteStringToCharBuffer(buffer *__charArray, string_ String) String {
	for i := int32(0); i < string_.Length; i++ {
		*(*byte)(unsafe.Add(unsafe.Pointer(buffer.InternalArray), buffer.Length+i)) = *(*byte)(unsafe.Add(unsafe.Pointer(string_.Chars), i))
	}
	buffer.Length += string_.Length
	return String{Length: string_.Length, Chars: ((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(buffer.InternalArray), buffer.Length))), -string_.Length)))}
}

var __MeasureText func(text StringSlice, config *TextElementConfig, userData unsafe.Pointer) Dimensions
var __QueryScrollOffset func(elementId uint32, userData unsafe.Pointer) Vector2

func __GetOpenLayoutElement() *LayoutElement {
	var context *Context = GetCurrentContext()
	return LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&context.OpenLayoutElementStack, context.OpenLayoutElementStack.Length-1))
}
func __GetParentElementId() uint32 {
	var context *Context = GetCurrentContext()
	return LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&context.OpenLayoutElementStack, context.OpenLayoutElementStack.Length-2)).Id
}
func __StoreLayoutConfig(config LayoutConfig) *LayoutConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &LAYOUT_DEFAULT
	}
	return __LayoutConfigArray_Add(&GetCurrentContext().LayoutConfigs, config)
}
func __StoreTextElementConfig(config TextElementConfig) *TextElementConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &TextElementConfig_DEFAULT
	}
	return __TextElementConfigArray_Add(&GetCurrentContext().TextElementConfigs, config)
}
func __StoreImageElementConfig(config ImageElementConfig) *ImageElementConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &ImageElementConfig_DEFAULT
	}
	return __ImageElementConfigArray_Add(&GetCurrentContext().ImageElementConfigs, config)
}
func __StoreFloatingElementConfig(config FloatingElementConfig) *FloatingElementConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &FloatingElementConfig_DEFAULT
	}
	return __FloatingElementConfigArray_Add(&GetCurrentContext().FloatingElementConfigs, config)
}
func __StoreCustomElementConfig(config CustomElementConfig) *CustomElementConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &CustomElementConfig_DEFAULT
	}
	return __CustomElementConfigArray_Add(&GetCurrentContext().CustomElementConfigs, config)
}
func __StoreScrollElementConfig(config ScrollElementConfig) *ScrollElementConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &ScrollElementConfig_DEFAULT
	}
	return __ScrollElementConfigArray_Add(&GetCurrentContext().ScrollElementConfigs, config)
}
func __StoreBorderElementConfig(config BorderElementConfig) *BorderElementConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &BorderElementConfig_DEFAULT
	}
	return __BorderElementConfigArray_Add(&GetCurrentContext().BorderElementConfigs, config)
}
func __StoreSharedElementConfig(config SharedElementConfig) *SharedElementConfig {
	if GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &SharedElementConfig_DEFAULT
	}
	return __SharedElementConfigArray_Add(&GetCurrentContext().SharedElementConfigs, config)
}
func __AttachElementConfig(config ElementConfigUnion, type_ __ElementConfigType) ElementConfig {
	var context *Context = GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return ElementConfig{}
	}
	var openLayoutElement *LayoutElement = __GetOpenLayoutElement()
	openLayoutElement.ElementConfigs.Length++
	return *__ElementConfigArray_Add(&context.ElementConfigs, ElementConfig{Type: type_, Config: config})
}
func __FindElementConfigWithType(element *LayoutElement, type_ __ElementConfigType) ElementConfigUnion {
	for i := int32(0); i < element.ElementConfigs.Length; i++ {
		var config *ElementConfig = __ElementConfigArraySlice_Get(&element.ElementConfigs, i)
		if config.Type == type_ {
			return config.Config
		}
	}
	return ElementConfigUnion{}
}
func __HashNumber(offset uint32, seed uint32) ElementId {
	var hash uint32 = seed
	hash += offset + 48
	hash += hash << 10
	hash ^= hash >> 6
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return ElementId{Id: hash + 1, Offset: offset, BaseId: seed, StringId: __STRING_DEFAULT}
}
func __HashString(key String, offset uint32, seed uint32) ElementId {
	var (
		hash uint32 = 0
		base uint32 = seed
	)
	for i := int32(0); i < key.Length; i++ {
		base += uint32(*(*byte)(unsafe.Add(unsafe.Pointer(key.Chars), i)))
		base += base << 10
		base ^= base >> 6
	}
	hash = base
	hash += offset
	hash += hash << 10
	hash ^= hash >> 6
	hash += hash << 3
	base += base << 3
	hash ^= hash >> 11
	base ^= base >> 11
	hash += hash << 15
	base += base << 15
	return ElementId{Id: hash + 1, Offset: offset, BaseId: base + 1, StringId: key}
}
func __HashTextWithConfig(text *String, config *TextElementConfig) uint32 {
	var (
		hash            uint32 = 0
		pointerAsNumber uint64 = uint64(uintptr(unsafe.Pointer(text.Chars)))
	)
	if config.HashStringContents {
		var maxLengthToHash uint32 = uint32(func() int32 {
			if text.Length < 256 {
				return text.Length
			}
			return 256
		}())
		for i := uint32(0); i < maxLengthToHash; i++ {
			hash += uint32(*(*byte)(unsafe.Add(unsafe.Pointer(text.Chars), i)))
			hash += hash << 10
			hash ^= hash >> 6
		}
	} else {
		hash += uint32(pointerAsNumber)
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += uint32(text.Length)
	hash += hash << 10
	hash ^= hash >> 6
	hash += uint32(config.FontId)
	hash += hash << 10
	hash ^= hash >> 6
	hash += uint32(config.FontSize)
	hash += hash << 10
	hash ^= hash >> 6
	hash += uint32(config.LineHeight)
	hash += hash << 10
	hash ^= hash >> 6
	hash += uint32(config.LetterSpacing)
	hash += hash << 10
	hash ^= hash >> 6
	hash += uint32(config.WrapMode)
	hash += hash << 10
	hash ^= hash >> 6
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return hash + 1
}
func __AddMeasuredWord(word __MeasuredWord, previousWord *__MeasuredWord) *__MeasuredWord {
	var context *Context = GetCurrentContext()
	if context.MeasuredWordsFreeList.Length > 0 {
		var newItemIndex uint32 = uint32(__int32_tArray_GetValue(&context.MeasuredWordsFreeList, context.MeasuredWordsFreeList.Length-1))
		context.MeasuredWordsFreeList.Length--
		__MeasuredWordArray_Set(&context.MeasuredWords, int32(newItemIndex), word)
		previousWord.Next = int32(newItemIndex)
		return __MeasuredWordArray_Get(&context.MeasuredWords, int32(newItemIndex))
	} else {
		previousWord.Next = context.MeasuredWords.Length
		return __MeasuredWordArray_Add(&context.MeasuredWords, word)
	}
}
func __MeasureTextCached(text *String, config *TextElementConfig) *__MeasureTextCacheItem {
	var context *Context = GetCurrentContext()
	if __MeasureText == nil {
		if !context.BooleanWarnings.TextMeasurementFunctionNotSet {
			context.BooleanWarnings.TextMeasurementFunctionNotSet = true
			context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_TEXT_MEASUREMENT_FUNCTION_NOT_PROVIDED, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay's internal MeasureText function is null. You may have forgotten to call SetMeasureTextFunction(), or passed a NULL function pointer by mistake.")}, UserData: context.ErrorHandler.UserData.(any)})
		}
		return &__MeasureTextCacheItem_DEFAULT
	}
	var id uint32 = __HashTextWithConfig(text, config)
	var hashBucket uint32 = id % uint32(context.MaxMeasureTextCacheWordCount/32)
	var elementIndexPrevious int32 = 0
	var elementIndex int32 = *(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	for elementIndex != 0 {
		var hashEntry *__MeasureTextCacheItem = __MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, elementIndex)
		if hashEntry.Id == id {
			hashEntry.Generation = context.Generation
			return hashEntry
		}
		if context.Generation-hashEntry.Generation > 2 {
			var nextWordIndex int32 = hashEntry.MeasuredWordsStartIndex
			for nextWordIndex != -1 {
				var measuredWord *__MeasuredWord = __MeasuredWordArray_Get(&context.MeasuredWords, nextWordIndex)
				__int32_tArray_Add(&context.MeasuredWordsFreeList, nextWordIndex)
				nextWordIndex = measuredWord.Next
			}
			var nextIndex int32 = hashEntry.NextIndex
			__MeasureTextCacheItemArray_Set(&context.MeasureTextHashMapInternal, elementIndex, __MeasureTextCacheItem{MeasuredWordsStartIndex: -1})
			__int32_tArray_Add(&context.MeasureTextHashMapInternalFreeList, elementIndex)
			if elementIndexPrevious == 0 {
				*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = nextIndex
			} else {
				var previousHashEntry *__MeasureTextCacheItem = __MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, elementIndexPrevious)
				previousHashEntry.NextIndex = nextIndex
			}
			elementIndex = nextIndex
		} else {
			elementIndexPrevious = elementIndex
			elementIndex = hashEntry.NextIndex
		}
	}
	var newItemIndex int32 = 0
	var newCacheItem __MeasureTextCacheItem = __MeasureTextCacheItem{MeasuredWordsStartIndex: -1, Id: id, Generation: context.Generation}
	var measured *__MeasureTextCacheItem = nil
	if context.MeasureTextHashMapInternalFreeList.Length > 0 {
		newItemIndex = __int32_tArray_GetValue(&context.MeasureTextHashMapInternalFreeList, context.MeasureTextHashMapInternalFreeList.Length-1)
		context.MeasureTextHashMapInternalFreeList.Length--
		__MeasureTextCacheItemArray_Set(&context.MeasureTextHashMapInternal, newItemIndex, newCacheItem)
		measured = __MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, newItemIndex)
	} else {
		if context.MeasureTextHashMapInternal.Length == context.MeasureTextHashMapInternal.Capacity-1 {
			if context.BooleanWarnings.MaxTextMeasureCacheExceeded {
				context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay ran out of capacity while attempting to measure text elements. Try using SetMaxElementCount() with a higher value.")}, UserData: context.ErrorHandler.UserData.(any)})
				context.BooleanWarnings.MaxTextMeasureCacheExceeded = true
			}
			return &__MeasureTextCacheItem_DEFAULT
		}
		measured = __MeasureTextCacheItemArray_Add(&context.MeasureTextHashMapInternal, newCacheItem)
		newItemIndex = context.MeasureTextHashMapInternal.Length - 1
	}
	var start int32 = 0
	var end int32 = 0
	var lineWidth float32 = 0
	var measuredWidth float32 = 0
	var measuredHeight float32 = 0
	var spaceWidth float32 = __MeasureText(StringSlice{Length: 1, Chars: __SPACECHAR.Chars, BaseChars: __SPACECHAR.Chars}, config, context.MeasureTextUserData.(unsafe.Pointer)).Width
	var tempWord __MeasuredWord = __MeasuredWord{Next: -1}
	var previousWord *__MeasuredWord = &tempWord
	for end < text.Length {
		if context.MeasuredWords.Length == context.MeasuredWords.Capacity-1 {
			if !context.BooleanWarnings.MaxTextMeasureCacheExceeded {
				context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_TEXT_MEASUREMENT_CAPACITY_EXCEEDED, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay has run out of space in it's internal text measurement cache. Try using SetMaxMeasureTextCacheWordCount() (default 16384, with 1 unit storing 1 measured word).")}, UserData: context.ErrorHandler.UserData.(any)})
				context.BooleanWarnings.MaxTextMeasureCacheExceeded = true
			}
			return &__MeasureTextCacheItem_DEFAULT
		}
		var current int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(text.Chars), end)))
		if int32(current) == ' ' || int32(current) == '\n' {
			var (
				length     int32      = end - start
				dimensions Dimensions = __MeasureText(StringSlice{Length: length, Chars: (*byte)(unsafe.Add(unsafe.Pointer(text.Chars), start)), BaseChars: text.Chars}, config, context.MeasureTextUserData.(unsafe.Pointer))
			)
			if measuredHeight > dimensions.Height {
				measuredHeight = measuredHeight
			} else {
				measuredHeight = dimensions.Height
			}
			if int32(current) == ' ' {
				dimensions.Width += spaceWidth
				previousWord = __AddMeasuredWord(__MeasuredWord{StartOffset: start, Length: length + 1, Width: dimensions.Width, Next: -1}, previousWord)
				lineWidth += dimensions.Width
			}
			if int32(current) == '\n' {
				if length > 0 {
					previousWord = __AddMeasuredWord(__MeasuredWord{StartOffset: start, Length: length, Width: dimensions.Width, Next: -1}, previousWord)
				}
				previousWord = __AddMeasuredWord(__MeasuredWord{StartOffset: end + 1, Length: 0, Width: 0, Next: -1}, previousWord)
				lineWidth += dimensions.Width
				if lineWidth > measuredWidth {
					measuredWidth = lineWidth
				} else {
					measuredWidth = measuredWidth
				}
				measured.ContainsNewlines = true
				lineWidth = 0
			}
			start = end + 1
		}
		end++
	}
	if end-start > 0 {
		var dimensions Dimensions = __MeasureText(StringSlice{Length: end - start, Chars: (*byte)(unsafe.Add(unsafe.Pointer(text.Chars), start)), BaseChars: text.Chars}, config, context.MeasureTextUserData.(unsafe.Pointer))
		__AddMeasuredWord(__MeasuredWord{StartOffset: start, Length: end - start, Width: dimensions.Width, Next: -1}, previousWord)
		lineWidth += dimensions.Width
		if measuredHeight > dimensions.Height {
			measuredHeight = measuredHeight
		} else {
			measuredHeight = dimensions.Height
		}
	}
	if lineWidth > measuredWidth {
		measuredWidth = lineWidth
	} else {
		measuredWidth = measuredWidth
	}
	measured.MeasuredWordsStartIndex = tempWord.Next
	measured.UnwrappedDimensions.Width = measuredWidth
	measured.UnwrappedDimensions.Height = measuredHeight
	if elementIndexPrevious != 0 {
		__MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, elementIndexPrevious).NextIndex = newItemIndex
	} else {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = newItemIndex
	}
	return measured
}
func __PointIsInsideRect(point Vector2, rect BoundingBox) bool {
	return point.X >= rect.X && point.X <= rect.X+rect.Width && point.Y >= rect.Y && point.Y <= rect.Y+rect.Height
}
func __AddHashMapItem(elementId ElementId, layoutElement *LayoutElement, idAlias uint32) *LayoutElementHashMapItem {
	var context *Context = GetCurrentContext()
	if context.LayoutElementsHashMapInternal.Length == context.LayoutElementsHashMapInternal.Capacity-1 {
		return nil
	}
	var item LayoutElementHashMapItem = LayoutElementHashMapItem{ElementId: elementId, LayoutElement: layoutElement, NextIndex: -1, Generation: context.Generation + 1, IdAlias: idAlias}
	var hashBucket uint32 = elementId.Id % uint32(context.LayoutElementsHashMap.Capacity)
	var hashItemPrevious int32 = -1
	var hashItemIndex int32 = *(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	for hashItemIndex != -1 {
		var hashItem *LayoutElementHashMapItem = __LayoutElementHashMapItemArray_Get(&context.LayoutElementsHashMapInternal, hashItemIndex)
		if hashItem.ElementId.Id == elementId.Id {
			item.NextIndex = hashItem.NextIndex
			if hashItem.Generation <= context.Generation {
				hashItem.ElementId = elementId
				hashItem.Generation = context.Generation + 1
				hashItem.LayoutElement = layoutElement
				hashItem.DebugData.Collision = false
			} else {
				context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_DUPLICATE_ID, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("An element with this ID was already previously declared during this layout.")}, UserData: context.ErrorHandler.UserData.(any)})
				if context.DebugModeEnabled {
					hashItem.DebugData.Collision = true
				}
			}
			return hashItem
		}
		hashItemPrevious = hashItemIndex
		hashItemIndex = hashItem.NextIndex
	}
	var hashItem *LayoutElementHashMapItem = __LayoutElementHashMapItemArray_Add(&context.LayoutElementsHashMapInternal, item)
	hashItem.DebugData = __DebugElementDataArray_Add(&context.DebugElementData, __DebugElementData{Collision: false})
	if hashItemPrevious != -1 {
		__LayoutElementHashMapItemArray_Get(&context.LayoutElementsHashMapInternal, hashItemPrevious).NextIndex = context.LayoutElementsHashMapInternal.Length - 1
	} else {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = context.LayoutElementsHashMapInternal.Length - 1
	}
	return hashItem
}
func __GetHashMapItem(id uint32) *LayoutElementHashMapItem {
	var (
		context      *Context = GetCurrentContext()
		hashBucket   uint32   = id % uint32(context.LayoutElementsHashMap.Capacity)
		elementIndex int32    = *(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	)
	for elementIndex != -1 {
		var hashEntry *LayoutElementHashMapItem = __LayoutElementHashMapItemArray_Get(&context.LayoutElementsHashMapInternal, elementIndex)
		if hashEntry.ElementId.Id == id {
			return hashEntry
		}
		elementIndex = hashEntry.NextIndex
	}
	return &LayoutElementHashMapItem_DEFAULT
}
func __GenerateIdForAnonymousElement(openLayoutElement *LayoutElement) ElementId {
	var (
		context       *Context       = GetCurrentContext()
		parentElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&context.OpenLayoutElementStack, context.OpenLayoutElementStack.Length-2))
		elementId     ElementId      = __HashNumber(uint32(parentElement.ChildrenOrTextContent.Children.Length), parentElement.Id)
	)
	openLayoutElement.Id = elementId.Id
	__AddHashMapItem(elementId, openLayoutElement, 0)
	__StringArray_Add(&context.LayoutElementIdStrings, elementId.StringId)
	return elementId
}
func __ElementHasConfig(layoutElement *LayoutElement, type_ __ElementConfigType) bool {
	for i := int32(0); i < layoutElement.ElementConfigs.Length; i++ {
		if __ElementConfigArraySlice_Get(&layoutElement.ElementConfigs, i).Type == type_ {
			return true
		}
	}
	return false
}
func __CloseElement() {
	var context *Context = GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return
	}
	var openLayoutElement *LayoutElement = __GetOpenLayoutElement()
	var layoutConfig *LayoutConfig = openLayoutElement.LayoutConfig
	var elementHasScrollHorizontal bool = false
	var elementHasScrollVertical bool = false
	for i := int32(0); i < openLayoutElement.ElementConfigs.Length; i++ {
		var config *ElementConfig = __ElementConfigArraySlice_Get(&openLayoutElement.ElementConfigs, i)
		if config.Type == __ELEMENT_CONFIG_TYPE_SCROLL {
			elementHasScrollHorizontal = config.Config.ScrollElementConfig.Horizontal
			elementHasScrollVertical = config.Config.ScrollElementConfig.Vertical
			context.OpenClipElementStack.Length--
			break
		}
	}
	openLayoutElement.ChildrenOrTextContent.Children.Elements = (*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementChildren.InternalArray), unsafe.Sizeof(int32(0))*uintptr(context.LayoutElementChildren.Length)))
	if layoutConfig.LayoutDirection == LEFT_TO_RIGHT {
		openLayoutElement.Dimensions.Width = float32(int32(layoutConfig.Padding.Left) + int32(layoutConfig.Padding.Right))
		for i := int32(0); i < int32(openLayoutElement.ChildrenOrTextContent.Children.Length); i++ {
			var (
				childIndex int32          = __int32_tArray_GetValue(&context.LayoutElementChildrenBuffer, context.LayoutElementChildrenBuffer.Length-int32(openLayoutElement.ChildrenOrTextContent.Children.Length)+i)
				child      *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, childIndex)
			)
			openLayoutElement.Dimensions.Width += child.Dimensions.Width
			if openLayoutElement.Dimensions.Height > (child.Dimensions.Height + float32(layoutConfig.Padding.Top) + float32(layoutConfig.Padding.Bottom)) {
				openLayoutElement.Dimensions.Height = openLayoutElement.Dimensions.Height
			} else {
				openLayoutElement.Dimensions.Height = child.Dimensions.Height + float32(layoutConfig.Padding.Top) + float32(layoutConfig.Padding.Bottom)
			}
			if !elementHasScrollHorizontal {
				openLayoutElement.MinDimensions.Width += child.MinDimensions.Width
			}
			if !elementHasScrollVertical {
				if openLayoutElement.MinDimensions.Height > (child.MinDimensions.Height + float32(layoutConfig.Padding.Top) + float32(layoutConfig.Padding.Bottom)) {
					openLayoutElement.MinDimensions.Height = openLayoutElement.MinDimensions.Height
				} else {
					openLayoutElement.MinDimensions.Height = child.MinDimensions.Height + float32(layoutConfig.Padding.Top) + float32(layoutConfig.Padding.Bottom)
				}
			}
			__int32_tArray_Add(&context.LayoutElementChildren, childIndex)
		}
		var childGap float32 = float32((func() int32 {
			if (int32(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
				return int32(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1
			}
			return 0
		}()) * int32(layoutConfig.ChildGap))
		openLayoutElement.Dimensions.Width += childGap
		openLayoutElement.MinDimensions.Width += childGap
	} else if layoutConfig.LayoutDirection == TOP_TO_BOTTOM {
		openLayoutElement.Dimensions.Height = float32(int32(layoutConfig.Padding.Top) + int32(layoutConfig.Padding.Bottom))
		for i := int32(0); i < int32(openLayoutElement.ChildrenOrTextContent.Children.Length); i++ {
			var (
				childIndex int32          = __int32_tArray_GetValue(&context.LayoutElementChildrenBuffer, context.LayoutElementChildrenBuffer.Length-int32(openLayoutElement.ChildrenOrTextContent.Children.Length)+i)
				child      *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, childIndex)
			)
			openLayoutElement.Dimensions.Height += child.Dimensions.Height
			if openLayoutElement.Dimensions.Width > (child.Dimensions.Width + float32(layoutConfig.Padding.Left) + float32(layoutConfig.Padding.Right)) {
				openLayoutElement.Dimensions.Width = openLayoutElement.Dimensions.Width
			} else {
				openLayoutElement.Dimensions.Width = child.Dimensions.Width + float32(layoutConfig.Padding.Left) + float32(layoutConfig.Padding.Right)
			}
			if !elementHasScrollVertical {
				openLayoutElement.MinDimensions.Height += child.MinDimensions.Height
			}
			if !elementHasScrollHorizontal {
				if openLayoutElement.MinDimensions.Width > (child.MinDimensions.Width + float32(layoutConfig.Padding.Left) + float32(layoutConfig.Padding.Right)) {
					openLayoutElement.MinDimensions.Width = openLayoutElement.MinDimensions.Width
				} else {
					openLayoutElement.MinDimensions.Width = child.MinDimensions.Width + float32(layoutConfig.Padding.Left) + float32(layoutConfig.Padding.Right)
				}
			}
			__int32_tArray_Add(&context.LayoutElementChildren, childIndex)
		}
		var childGap float32 = float32((func() int32 {
			if (int32(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
				return int32(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1
			}
			return 0
		}()) * int32(layoutConfig.ChildGap))
		openLayoutElement.Dimensions.Height += childGap
		openLayoutElement.MinDimensions.Height += childGap
	}
	context.LayoutElementChildrenBuffer.Length -= int32(openLayoutElement.ChildrenOrTextContent.Children.Length)
	if layoutConfig.Sizing.Width.Type != __SIZING_TYPE_PERCENT {
		if layoutConfig.Sizing.Width.Size.MinMax.Max <= 0 {
			layoutConfig.Sizing.Width.Size.MinMax.Max = __MAXFLOAT
		}
		if (func() float32 {
			if openLayoutElement.Dimensions.Width > layoutConfig.Sizing.Width.Size.MinMax.Min {
				return openLayoutElement.Dimensions.Width
			}
			return layoutConfig.Sizing.Width.Size.MinMax.Min
		}()) < layoutConfig.Sizing.Width.Size.MinMax.Max {
			if openLayoutElement.Dimensions.Width > layoutConfig.Sizing.Width.Size.MinMax.Min {
				openLayoutElement.Dimensions.Width = openLayoutElement.Dimensions.Width
			} else {
				openLayoutElement.Dimensions.Width = layoutConfig.Sizing.Width.Size.MinMax.Min
			}
		} else {
			openLayoutElement.Dimensions.Width = layoutConfig.Sizing.Width.Size.MinMax.Max
		}
		if (func() float32 {
			if openLayoutElement.MinDimensions.Width > layoutConfig.Sizing.Width.Size.MinMax.Min {
				return openLayoutElement.MinDimensions.Width
			}
			return layoutConfig.Sizing.Width.Size.MinMax.Min
		}()) < layoutConfig.Sizing.Width.Size.MinMax.Max {
			if openLayoutElement.MinDimensions.Width > layoutConfig.Sizing.Width.Size.MinMax.Min {
				openLayoutElement.MinDimensions.Width = openLayoutElement.MinDimensions.Width
			} else {
				openLayoutElement.MinDimensions.Width = layoutConfig.Sizing.Width.Size.MinMax.Min
			}
		} else {
			openLayoutElement.MinDimensions.Width = layoutConfig.Sizing.Width.Size.MinMax.Max
		}
	} else {
		openLayoutElement.Dimensions.Width = 0
	}
	if layoutConfig.Sizing.Height.Type != __SIZING_TYPE_PERCENT {
		if layoutConfig.Sizing.Height.Size.MinMax.Max <= 0 {
			layoutConfig.Sizing.Height.Size.MinMax.Max = __MAXFLOAT
		}
		if (func() float32 {
			if openLayoutElement.Dimensions.Height > layoutConfig.Sizing.Height.Size.MinMax.Min {
				return openLayoutElement.Dimensions.Height
			}
			return layoutConfig.Sizing.Height.Size.MinMax.Min
		}()) < layoutConfig.Sizing.Height.Size.MinMax.Max {
			if openLayoutElement.Dimensions.Height > layoutConfig.Sizing.Height.Size.MinMax.Min {
				openLayoutElement.Dimensions.Height = openLayoutElement.Dimensions.Height
			} else {
				openLayoutElement.Dimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Min
			}
		} else {
			openLayoutElement.Dimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Max
		}
		if (func() float32 {
			if openLayoutElement.MinDimensions.Height > layoutConfig.Sizing.Height.Size.MinMax.Min {
				return openLayoutElement.MinDimensions.Height
			}
			return layoutConfig.Sizing.Height.Size.MinMax.Min
		}()) < layoutConfig.Sizing.Height.Size.MinMax.Max {
			if openLayoutElement.MinDimensions.Height > layoutConfig.Sizing.Height.Size.MinMax.Min {
				openLayoutElement.MinDimensions.Height = openLayoutElement.MinDimensions.Height
			} else {
				openLayoutElement.MinDimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Min
			}
		} else {
			openLayoutElement.MinDimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Max
		}
	} else {
		openLayoutElement.Dimensions.Height = 0
	}
	var elementIsFloating bool = __ElementHasConfig(openLayoutElement, __ELEMENT_CONFIG_TYPE_FLOATING)
	var closingElementIndex int32 = __int32_tArray_RemoveSwapback(&context.OpenLayoutElementStack, context.OpenLayoutElementStack.Length-1)
	openLayoutElement = __GetOpenLayoutElement()
	if !elementIsFloating && context.OpenLayoutElementStack.Length > 1 {
		openLayoutElement.ChildrenOrTextContent.Children.Length++
		__int32_tArray_Add(&context.LayoutElementChildrenBuffer, closingElementIndex)
	}
}
func __MemCmp(s1 *byte, s2 *byte, length int32) bool {
	for i := int32(0); i < length; i++ {
		if *(*byte)(unsafe.Add(unsafe.Pointer(s1), i)) != *(*byte)(unsafe.Add(unsafe.Pointer(s2), i)) {
			return false
		}
	}
	return true
}
func __OpenElement() {
	var context *Context = GetCurrentContext()
	if context.LayoutElements.Length == context.LayoutElements.Capacity-1 || context.BooleanWarnings.MaxElementsExceeded {
		context.BooleanWarnings.MaxElementsExceeded = true
		return
	}
	var layoutElement LayoutElement = LayoutElement{}
	LayoutElementArray_Add(&context.LayoutElements, layoutElement)
	__int32_tArray_Add(&context.OpenLayoutElementStack, context.LayoutElements.Length-1)
	if context.OpenClipElementStack.Length > 0 {
		__int32_tArray_Set(&context.LayoutElementClipElementIds, context.LayoutElements.Length-1, __int32_tArray_GetValue(&context.OpenClipElementStack, context.OpenClipElementStack.Length-1))
	} else {
		__int32_tArray_Set(&context.LayoutElementClipElementIds, context.LayoutElements.Length-1, 0)
	}
}
func __OpenTextElement(text String, textConfig *TextElementConfig) {
	var context *Context = GetCurrentContext()
	if context.LayoutElements.Length == context.LayoutElements.Capacity-1 || context.BooleanWarnings.MaxElementsExceeded {
		context.BooleanWarnings.MaxElementsExceeded = true
		return
	}
	var parentElement *LayoutElement = __GetOpenLayoutElement()
	var layoutElement LayoutElement = LayoutElement{}
	var textElement *LayoutElement = LayoutElementArray_Add(&context.LayoutElements, layoutElement)
	if context.OpenClipElementStack.Length > 0 {
		__int32_tArray_Set(&context.LayoutElementClipElementIds, context.LayoutElements.Length-1, __int32_tArray_GetValue(&context.OpenClipElementStack, context.OpenClipElementStack.Length-1))
	} else {
		__int32_tArray_Set(&context.LayoutElementClipElementIds, context.LayoutElements.Length-1, 0)
	}
	__int32_tArray_Add(&context.LayoutElementChildrenBuffer, context.LayoutElements.Length-1)
	var textMeasured *__MeasureTextCacheItem = __MeasureTextCached(&text, textConfig)
	var elementId ElementId = __HashNumber(uint32(parentElement.ChildrenOrTextContent.Children.Length), parentElement.Id)
	textElement.Id = elementId.Id
	__AddHashMapItem(elementId, textElement, 0)
	__StringArray_Add(&context.LayoutElementIdStrings, elementId.StringId)
	var textDimensions Dimensions = Dimensions{Width: textMeasured.UnwrappedDimensions.Width, Height: func() float32 {
		if int32(textConfig.LineHeight) > 0 {
			return float32(textConfig.LineHeight)
		}
		return textMeasured.UnwrappedDimensions.Height
	}()}
	textElement.Dimensions = textDimensions
	textElement.MinDimensions = Dimensions{Width: textMeasured.UnwrappedDimensions.Height, Height: textDimensions.Height}
	textElement.ChildrenOrTextContent.TextElementData = __TextElementDataArray_Add(&context.TextElementData, __TextElementData{Text: text, PreferredDimensions: textMeasured.UnwrappedDimensions, ElementIndex: context.LayoutElements.Length - 1})
	textElement.ElementConfigs = __ElementConfigArraySlice{Length: 1, InternalArray: __ElementConfigArray_Add(&context.ElementConfigs, ElementConfig{Type: __ELEMENT_CONFIG_TYPE_TEXT, Config: ElementConfigUnion{TextElementConfig: textConfig}})}
	textElement.LayoutConfig = &LAYOUT_DEFAULT
	parentElement.ChildrenOrTextContent.Children.Length++
}
func __AttachId(elementId ElementId) ElementId {
	var context *Context = GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return ElementId_DEFAULT
	}
	var openLayoutElement *LayoutElement = __GetOpenLayoutElement()
	var idAlias uint32 = openLayoutElement.Id
	openLayoutElement.Id = elementId.Id
	__AddHashMapItem(elementId, openLayoutElement, idAlias)
	__StringArray_Add(&context.LayoutElementIdStrings, elementId.StringId)
	return elementId
}
func __ConfigureOpenElement(declaration ElementDeclaration) {
	var (
		context           *Context       = GetCurrentContext()
		openLayoutElement *LayoutElement = __GetOpenLayoutElement()
	)
	openLayoutElement.LayoutConfig = __StoreLayoutConfig(declaration.Layout)
	if declaration.Layout.Sizing.Width.Type == __SIZING_TYPE_PERCENT && declaration.Layout.Sizing.Width.Size.Percent > 1 || declaration.Layout.Sizing.Height.Type == __SIZING_TYPE_PERCENT && declaration.Layout.Sizing.Height.Size.Percent > 1 {
		context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_PERCENTAGE_OVER_1, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("An element was configured with SIZING_PERCENT, but the provided percentage value was over 1.0. Clay expects a value between 0 and 1, i.e. 20% is 0.2.")}, UserData: context.ErrorHandler.UserData.(any)})
	}
	var openLayoutElementId ElementId = declaration.Id
	openLayoutElement.ElementConfigs.InternalArray = (*ElementConfig)(unsafe.Add(unsafe.Pointer(context.ElementConfigs.InternalArray), unsafe.Sizeof(ElementConfig{})*uintptr(context.ElementConfigs.Length)))
	var sharedConfig *SharedElementConfig = nil
	if declaration.BackgroundColor.A > 0 {
		sharedConfig = __StoreSharedElementConfig(SharedElementConfig{BackgroundColor: declaration.BackgroundColor})
		__AttachElementConfig(ElementConfigUnion{SharedElementConfig: sharedConfig}, __ELEMENT_CONFIG_TYPE_SHARED)
	}
	if !__MemCmp((*byte)(unsafe.Pointer(&declaration.CornerRadius)), (*byte)(unsafe.Pointer(&__CornerRadius_DEFAULT)), int32(uint32(unsafe.Sizeof(CornerRadius{})))) {
		if sharedConfig != nil {
			sharedConfig.CornerRadius = declaration.CornerRadius
		} else {
			sharedConfig = __StoreSharedElementConfig(SharedElementConfig{CornerRadius: declaration.CornerRadius})
			__AttachElementConfig(ElementConfigUnion{SharedElementConfig: sharedConfig}, __ELEMENT_CONFIG_TYPE_SHARED)
		}
	}
	if declaration.UserData != nil {
		if sharedConfig != nil {
			sharedConfig.UserData = declaration.UserData
		} else {
			sharedConfig = __StoreSharedElementConfig(SharedElementConfig{UserData: declaration.UserData})
			__AttachElementConfig(ElementConfigUnion{SharedElementConfig: sharedConfig}, __ELEMENT_CONFIG_TYPE_SHARED)
		}
	}
	if declaration.Image.ImageData != nil {
		__AttachElementConfig(ElementConfigUnion{ImageElementConfig: __StoreImageElementConfig(declaration.Image)}, __ELEMENT_CONFIG_TYPE_IMAGE)
		__int32_tArray_Add(&context.ImageElementPointers, context.LayoutElements.Length-1)
	}
	if declaration.Floating.AttachTo != ATTACH_TO_NONE {
		var (
			floatingConfig     FloatingElementConfig = declaration.Floating
			hierarchicalParent *LayoutElement        = LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&context.OpenLayoutElementStack, context.OpenLayoutElementStack.Length-2))
		)
		if hierarchicalParent != nil {
			var clipElementId uint32 = 0
			if declaration.Floating.AttachTo == ATTACH_TO_PARENT {
				floatingConfig.ParentId = hierarchicalParent.Id
				if context.OpenClipElementStack.Length > 0 {
					clipElementId = uint32(__int32_tArray_GetValue(&context.OpenClipElementStack, context.OpenClipElementStack.Length-1))
				}
			} else if declaration.Floating.AttachTo == ATTACH_TO_ELEMENT_WITH_ID {
				var parentItem *LayoutElementHashMapItem = __GetHashMapItem(floatingConfig.ParentId)
				if parentItem == nil {
					context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_FLOATING_CONTAINER_PARENT_NOT_FOUND, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("A floating element was declared with a parentId, but no element with that ID was found.")}, UserData: context.ErrorHandler.UserData.(any)})
				} else {
					clipElementId = uint32(__int32_tArray_GetValue(&context.LayoutElementClipElementIds, int32(int64(uintptr(unsafe.Pointer(parentItem.LayoutElement))-uintptr(unsafe.Pointer(context.LayoutElements.InternalArray))))))
				}
			} else if declaration.Floating.AttachTo == ATTACH_TO_ROOT {
				floatingConfig.ParentId = __HashString(String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("__RootContainer")}, 0, 0).Id
			}
			if openLayoutElementId.Id == 0 {
				openLayoutElementId = __HashString(String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("__FloatingContainer")}, uint32(context.LayoutElementTreeRoots.Length), 0)
			}
			__LayoutElementTreeRootArray_Add(&context.LayoutElementTreeRoots, __LayoutElementTreeRoot{LayoutElementIndex: __int32_tArray_GetValue(&context.OpenLayoutElementStack, context.OpenLayoutElementStack.Length-1), ParentId: floatingConfig.ParentId, ClipElementId: clipElementId, ZIndex: floatingConfig.ZIndex})
			__AttachElementConfig(ElementConfigUnion{FloatingElementConfig: __StoreFloatingElementConfig(declaration.Floating)}, __ELEMENT_CONFIG_TYPE_FLOATING)
		}
	}
	if declaration.Custom.CustomData != nil {
		__AttachElementConfig(ElementConfigUnion{CustomElementConfig: __StoreCustomElementConfig(declaration.Custom)}, __ELEMENT_CONFIG_TYPE_CUSTOM)
	}
	if openLayoutElementId.Id != 0 {
		__AttachId(openLayoutElementId)
	} else if openLayoutElement.Id == 0 {
		openLayoutElementId = __GenerateIdForAnonymousElement(openLayoutElement)
	}
	if declaration.Scroll.Horizontal || declaration.Scroll.Vertical {
		__AttachElementConfig(ElementConfigUnion{ScrollElementConfig: __StoreScrollElementConfig(declaration.Scroll)}, __ELEMENT_CONFIG_TYPE_SCROLL)
		__int32_tArray_Add(&context.OpenClipElementStack, int32(openLayoutElement.Id))
		var scrollOffset *__ScrollContainerDataInternal = (*__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(__NULL)))
		for i := int32(0); i < context.ScrollContainerDatas.Length; i++ {
			var mapping *__ScrollContainerDataInternal = __ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
			if openLayoutElement.Id == mapping.ElementId {
				scrollOffset = mapping
				scrollOffset.LayoutElement = openLayoutElement
				scrollOffset.OpenThisFrame = true
			}
		}
		if scrollOffset == nil {
			scrollOffset = __ScrollContainerDataInternalArray_Add(&context.ScrollContainerDatas, __ScrollContainerDataInternal{LayoutElement: openLayoutElement, ScrollOrigin: Vector2{X: -1, Y: -1}, ElementId: openLayoutElement.Id, OpenThisFrame: true})
		}
		if context.ExternalScrollHandlingEnabled {
			scrollOffset.ScrollPosition = __QueryScrollOffset(scrollOffset.ElementId, context.QueryScrollOffsetUserData.(unsafe.Pointer))
		}
	}
	if !__MemCmp((*byte)(unsafe.Pointer(&declaration.Border.Width)), (*byte)(unsafe.Pointer(&__BorderWidth_DEFAULT)), int32(uint32(unsafe.Sizeof(BorderWidth{})))) {
		__AttachElementConfig(ElementConfigUnion{BorderElementConfig: __StoreBorderElementConfig(declaration.Border)}, __ELEMENT_CONFIG_TYPE_BORDER)
	}
}
func __InitializeEphemeralMemory(context *Context) {
	var (
		maxElementCount int32  = context.MaxElementCount
		arena           *Arena = &context.InternalArena
	)
	arena.NextAllocation = context.ArenaResetOffset
	context.LayoutElementChildrenBuffer = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElements = LayoutElementArray_Allocate_Arena(maxElementCount, arena)
	context.Warnings = __WarningArray_Allocate_Arena(100, arena)
	context.LayoutConfigs = __LayoutConfigArray_Allocate_Arena(maxElementCount, arena)
	context.ElementConfigs = __ElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.TextElementConfigs = __TextElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.ImageElementConfigs = __ImageElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.FloatingElementConfigs = __FloatingElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.ScrollElementConfigs = __ScrollElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.CustomElementConfigs = __CustomElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.BorderElementConfigs = __BorderElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.SharedElementConfigs = __SharedElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementIdStrings = __StringArray_Allocate_Arena(maxElementCount, arena)
	context.WrappedTextLines = __WrappedTextLineArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementTreeNodeArray1 = __LayoutElementTreeNodeArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementTreeRoots = __LayoutElementTreeRootArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementChildren = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.OpenLayoutElementStack = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.TextElementData = __TextElementDataArray_Allocate_Arena(maxElementCount, arena)
	context.ImageElementPointers = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.RenderCommands = RenderCommandArray_Allocate_Arena(maxElementCount, arena)
	context.TreeNodeVisited = __boolArray_Allocate_Arena(maxElementCount, arena)
	context.TreeNodeVisited.Length = context.TreeNodeVisited.Capacity
	context.OpenClipElementStack = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.ReusableElementIndexBuffer = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementClipElementIds = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.DynamicStringData = __charArray_Allocate_Arena(maxElementCount, arena)
}
func __InitializePersistentMemory(context *Context) {
	var (
		maxElementCount              int32  = context.MaxElementCount
		maxMeasureTextCacheWordCount int32  = context.MaxMeasureTextCacheWordCount
		arena                        *Arena = &context.InternalArena
	)
	context.ScrollContainerDatas = __ScrollContainerDataInternalArray_Allocate_Arena(10, arena)
	context.LayoutElementsHashMapInternal = __LayoutElementHashMapItemArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementsHashMap = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.MeasureTextHashMapInternal = __MeasureTextCacheItemArray_Allocate_Arena(maxElementCount, arena)
	context.MeasureTextHashMapInternalFreeList = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.MeasuredWordsFreeList = __int32_tArray_Allocate_Arena(maxMeasureTextCacheWordCount, arena)
	context.MeasureTextHashMap = __int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.MeasuredWords = __MeasuredWordArray_Allocate_Arena(maxMeasureTextCacheWordCount, arena)
	context.PointerOverIds = __ElementIdArray_Allocate_Arena(maxElementCount, arena)
	context.DebugElementData = __DebugElementDataArray_Allocate_Arena(maxElementCount, arena)
	context.ArenaResetOffset = arena.NextAllocation
}
func __CompressChildrenAlongAxis(xAxis bool, totalSizeToDistribute float32, resizableContainerBuffer __int32_tArray) {
	var (
		context           *Context       = GetCurrentContext()
		largestContainers __int32_tArray = context.OpenClipElementStack
	)
	for totalSizeToDistribute > 0.1 {
		largestContainers.Length = 0
		var largestSize float32 = 0
		var targetSize float32 = 0
		for i := int32(0); i < resizableContainerBuffer.Length; i++ {
			var (
				childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&resizableContainerBuffer, i))
				childSize    float32
			)
			if xAxis {
				childSize = childElement.Dimensions.Width
			} else {
				childSize = childElement.Dimensions.Height
			}
			if (childSize-largestSize) < 0.1 && (childSize-largestSize) > -0.1 {
				__int32_tArray_Add(&largestContainers, __int32_tArray_GetValue(&resizableContainerBuffer, i))
			} else if childSize > largestSize {
				targetSize = largestSize
				largestSize = childSize
				largestContainers.Length = 0
				__int32_tArray_Add(&largestContainers, __int32_tArray_GetValue(&resizableContainerBuffer, i))
			} else if childSize > targetSize {
				targetSize = childSize
			}
		}
		if largestContainers.Length == 0 {
			return
		}
		targetSize = (func() float32 {
			if targetSize > ((largestSize * float32(largestContainers.Length)) - totalSizeToDistribute) {
				return targetSize
			}
			return (largestSize * float32(largestContainers.Length)) - totalSizeToDistribute
		}()) / float32(largestContainers.Length)
		for childOffset := int32(0); childOffset < largestContainers.Length; childOffset++ {
			var (
				childIndex   int32          = __int32_tArray_GetValue(&largestContainers, childOffset)
				childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, childIndex)
				childSize    *float32
			)
			if xAxis {
				childSize = &childElement.Dimensions.Width
			} else {
				childSize = &childElement.Dimensions.Height
			}
			var childMinSize float32
			if xAxis {
				childMinSize = childElement.MinDimensions.Width
			} else {
				childMinSize = childElement.MinDimensions.Height
			}
			var oldChildSize float32 = *childSize
			if childMinSize > targetSize {
				*childSize = childMinSize
			} else {
				*childSize = targetSize
			}
			totalSizeToDistribute -= oldChildSize - *childSize
			if *childSize == childMinSize {
				for i := int32(0); i < resizableContainerBuffer.Length; i++ {
					if __int32_tArray_GetValue(&resizableContainerBuffer, i) == childIndex {
						__int32_tArray_RemoveSwapback(&resizableContainerBuffer, i)
						break
					}
				}
			}
		}
	}
}
func __SizeContainersAlongAxis(xAxis bool) {
	var (
		context                  *Context       = GetCurrentContext()
		bfsBuffer                __int32_tArray = context.LayoutElementChildrenBuffer
		resizableContainerBuffer __int32_tArray = context.OpenLayoutElementStack
	)
	for rootIndex := int32(0); rootIndex < context.LayoutElementTreeRoots.Length; rootIndex++ {
		bfsBuffer.Length = 0
		var root *__LayoutElementTreeRoot = __LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, rootIndex)
		var rootElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, root.LayoutElementIndex)
		__int32_tArray_Add(&bfsBuffer, root.LayoutElementIndex)
		if __ElementHasConfig(rootElement, __ELEMENT_CONFIG_TYPE_FLOATING) {
			var (
				floatingElementConfig *FloatingElementConfig    = __FindElementConfigWithType(rootElement, __ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig
				parentItem            *LayoutElementHashMapItem = __GetHashMapItem(floatingElementConfig.ParentId)
			)
			if parentItem != nil && parentItem != &LayoutElementHashMapItem_DEFAULT {
				var parentLayoutElement *LayoutElement = parentItem.LayoutElement
				if rootElement.LayoutConfig.Sizing.Width.Type == __SIZING_TYPE_GROW {
					rootElement.Dimensions.Width = parentLayoutElement.Dimensions.Width
				}
				if rootElement.LayoutConfig.Sizing.Height.Type == __SIZING_TYPE_GROW {
					rootElement.Dimensions.Height = parentLayoutElement.Dimensions.Height
				}
			}
		}
		if (func() float32 {
			if rootElement.Dimensions.Width > rootElement.LayoutConfig.Sizing.Width.Size.MinMax.Min {
				return rootElement.Dimensions.Width
			}
			return rootElement.LayoutConfig.Sizing.Width.Size.MinMax.Min
		}()) < rootElement.LayoutConfig.Sizing.Width.Size.MinMax.Max {
			if rootElement.Dimensions.Width > rootElement.LayoutConfig.Sizing.Width.Size.MinMax.Min {
				rootElement.Dimensions.Width = rootElement.Dimensions.Width
			} else {
				rootElement.Dimensions.Width = rootElement.LayoutConfig.Sizing.Width.Size.MinMax.Min
			}
		} else {
			rootElement.Dimensions.Width = rootElement.LayoutConfig.Sizing.Width.Size.MinMax.Max
		}
		if (func() float32 {
			if rootElement.Dimensions.Height > rootElement.LayoutConfig.Sizing.Height.Size.MinMax.Min {
				return rootElement.Dimensions.Height
			}
			return rootElement.LayoutConfig.Sizing.Height.Size.MinMax.Min
		}()) < rootElement.LayoutConfig.Sizing.Height.Size.MinMax.Max {
			if rootElement.Dimensions.Height > rootElement.LayoutConfig.Sizing.Height.Size.MinMax.Min {
				rootElement.Dimensions.Height = rootElement.Dimensions.Height
			} else {
				rootElement.Dimensions.Height = rootElement.LayoutConfig.Sizing.Height.Size.MinMax.Min
			}
		} else {
			rootElement.Dimensions.Height = rootElement.LayoutConfig.Sizing.Height.Size.MinMax.Max
		}
		for i := int32(0); i < bfsBuffer.Length; i++ {
			var (
				parentIndex        int32          = __int32_tArray_GetValue(&bfsBuffer, i)
				parent             *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, parentIndex)
				parentStyleConfig  *LayoutConfig  = parent.LayoutConfig
				growContainerCount int32          = 0
				parentSize         float32
			)
			if xAxis {
				parentSize = parent.Dimensions.Width
			} else {
				parentSize = parent.Dimensions.Height
			}
			var parentPadding float32 = float32(func() int32 {
				if xAxis {
					return int32(parent.LayoutConfig.Padding.Left) + int32(parent.LayoutConfig.Padding.Right)
				}
				return int32(parent.LayoutConfig.Padding.Top) + int32(parent.LayoutConfig.Padding.Bottom)
			}())
			var innerContentSize float32 = 0
			var growContainerContentSize float32 = 0
			var totalPaddingAndChildGaps float32 = parentPadding
			var sizingAlongAxis bool = xAxis && parentStyleConfig.LayoutDirection == LEFT_TO_RIGHT || !xAxis && parentStyleConfig.LayoutDirection == TOP_TO_BOTTOM
			resizableContainerBuffer.Length = 0
			var parentChildGap float32 = float32(parentStyleConfig.ChildGap)
			for childOffset := int32(0); childOffset < int32(parent.ChildrenOrTextContent.Children.Length); childOffset++ {
				var (
					childElementIndex int32          = *(*int32)(unsafe.Add(unsafe.Pointer(parent.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(childOffset)))
					childElement      *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, childElementIndex)
					childSizing       SizingAxis
				)
				if xAxis {
					childSizing = childElement.LayoutConfig.Sizing.Width
				} else {
					childSizing = childElement.LayoutConfig.Sizing.Height
				}
				var childSize float32
				if xAxis {
					childSize = childElement.Dimensions.Width
				} else {
					childSize = childElement.Dimensions.Height
				}
				if !__ElementHasConfig(childElement, __ELEMENT_CONFIG_TYPE_TEXT) && int32(childElement.ChildrenOrTextContent.Children.Length) > 0 {
					__int32_tArray_Add(&bfsBuffer, childElementIndex)
				}
				if childSizing.Type != __SIZING_TYPE_PERCENT && childSizing.Type != __SIZING_TYPE_FIXED && (!__ElementHasConfig(childElement, __ELEMENT_CONFIG_TYPE_TEXT) || __FindElementConfigWithType(childElement, __ELEMENT_CONFIG_TYPE_TEXT).TextElementConfig.WrapMode == TEXT_WRAP_WORDS) && (xAxis || !__ElementHasConfig(childElement, __ELEMENT_CONFIG_TYPE_IMAGE)) {
					__int32_tArray_Add(&resizableContainerBuffer, childElementIndex)
				}
				if sizingAlongAxis {
					if childSizing.Type == __SIZING_TYPE_PERCENT {
						innerContentSize += 0
					} else {
						innerContentSize += childSize
					}
					if childSizing.Type == __SIZING_TYPE_GROW {
						growContainerContentSize += childSize
						growContainerCount++
					}
					if childOffset > 0 {
						innerContentSize += parentChildGap
						totalPaddingAndChildGaps += parentChildGap
					}
				} else {
					if childSize > innerContentSize {
						innerContentSize = childSize
					} else {
						innerContentSize = innerContentSize
					}
				}
			}
			for childOffset := int32(0); childOffset < int32(parent.ChildrenOrTextContent.Children.Length); childOffset++ {
				var (
					childElementIndex int32          = *(*int32)(unsafe.Add(unsafe.Pointer(parent.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(childOffset)))
					childElement      *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, childElementIndex)
					childSizing       SizingAxis
				)
				if xAxis {
					childSizing = childElement.LayoutConfig.Sizing.Width
				} else {
					childSizing = childElement.LayoutConfig.Sizing.Height
				}
				var childSize *float32
				if xAxis {
					childSize = &childElement.Dimensions.Width
				} else {
					childSize = &childElement.Dimensions.Height
				}
				if childSizing.Type == __SIZING_TYPE_PERCENT {
					*childSize = (parentSize - totalPaddingAndChildGaps) * childSizing.Size.Percent
					if sizingAlongAxis {
						innerContentSize += *childSize
					}
				}
			}
			if sizingAlongAxis {
				var sizeToDistribute float32 = parentSize - parentPadding - innerContentSize
				if sizeToDistribute < 0 {
					var scrollElementConfig *ScrollElementConfig = __FindElementConfigWithType(parent, __ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
					if scrollElementConfig != nil {
						if xAxis && scrollElementConfig.Horizontal || !xAxis && scrollElementConfig.Vertical {
							continue
						}
					}
					__CompressChildrenAlongAxis(xAxis, -sizeToDistribute, resizableContainerBuffer)
				} else if sizeToDistribute > 0 && growContainerCount > 0 {
					var targetSize float32 = (sizeToDistribute + growContainerContentSize) / float32(growContainerCount)
					for childOffset := int32(0); childOffset < resizableContainerBuffer.Length; childOffset++ {
						var (
							childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&resizableContainerBuffer, childOffset))
							childSizing  SizingAxis
						)
						if xAxis {
							childSizing = childElement.LayoutConfig.Sizing.Width
						} else {
							childSizing = childElement.LayoutConfig.Sizing.Height
						}
						if childSizing.Type == __SIZING_TYPE_GROW {
							var childSize *float32
							_ = childSize
							if xAxis {
								childSize = &childElement.Dimensions.Width
							} else {
								childSize = &childElement.Dimensions.Height
							}
							var minSize *float32
							if xAxis {
								minSize = &childElement.MinDimensions.Width
							} else {
								minSize = &childElement.MinDimensions.Height
							}
							if targetSize < *minSize {
								growContainerContentSize -= *minSize
								__int32_tArray_RemoveSwapback(&resizableContainerBuffer, childOffset)
								growContainerCount--
								targetSize = (sizeToDistribute + growContainerContentSize) / float32(growContainerCount)
								childOffset = -1
								continue
							}
							*childSize = targetSize
						}
					}
				}
			} else {
				for childOffset := int32(0); childOffset < resizableContainerBuffer.Length; childOffset++ {
					var (
						childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&resizableContainerBuffer, childOffset))
						childSizing  SizingAxis
					)
					if xAxis {
						childSizing = childElement.LayoutConfig.Sizing.Width
					} else {
						childSizing = childElement.LayoutConfig.Sizing.Height
					}
					var childSize *float32
					if xAxis {
						childSize = &childElement.Dimensions.Width
					} else {
						childSize = &childElement.Dimensions.Height
					}
					if !xAxis && __ElementHasConfig(childElement, __ELEMENT_CONFIG_TYPE_IMAGE) {
						continue
					}
					var maxSize float32 = parentSize - parentPadding
					if __ElementHasConfig(parent, __ELEMENT_CONFIG_TYPE_SCROLL) {
						var scrollElementConfig *ScrollElementConfig = __FindElementConfigWithType(parent, __ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
						if xAxis && scrollElementConfig.Horizontal || !xAxis && scrollElementConfig.Vertical {
							if maxSize > innerContentSize {
								maxSize = maxSize
							} else {
								maxSize = innerContentSize
							}
						}
					}
					if childSizing.Type == __SIZING_TYPE_FIT {
						if childSizing.Size.MinMax.Min > (func() float32 {
							if (*childSize) < maxSize {
								return *childSize
							}
							return maxSize
						}()) {
							*childSize = childSizing.Size.MinMax.Min
						} else if (*childSize) < maxSize {
							*childSize = *childSize
						} else {
							*childSize = maxSize
						}
					} else if childSizing.Type == __SIZING_TYPE_GROW {
						if maxSize < childSizing.Size.MinMax.Max {
							*childSize = maxSize
						} else {
							*childSize = childSizing.Size.MinMax.Max
						}
					}
				}
			}
		}
	}
}
func __IntToString(integer int32) String {
	if integer == 0 {
		return String{Length: 1, Chars: libc.CString("0")}
	}
	var context *Context = GetCurrentContext()
	var chars *byte = ((*byte)(unsafe.Add(unsafe.Pointer(context.DynamicStringData.InternalArray), context.DynamicStringData.Length)))
	var length int32 = 0
	var sign int32 = integer
	if integer < 0 {
		integer = -integer
	}
	for integer > 0 {
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), func() int32 {
			p := &length
			x := *p
			*p++
			return x
		}())) = byte(int8(integer%10 + '0'))
		integer /= 10
	}
	if sign < 0 {
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), func() int32 {
			p := &length
			x := *p
			*p++
			return x
		}())) = '-'
	}
	for j, k := int32(0), int32(length-1); j < k; func() int32 {
		j++
		return func() int32 {
			p := &k
			x := *p
			*p--
			return x
		}()
	}() {
		var temp int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(chars), j)))
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), j)) = *(*byte)(unsafe.Add(unsafe.Pointer(chars), k))
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), k)) = byte(temp)
	}
	context.DynamicStringData.Length += length
	return String{Length: length, Chars: chars}
}
func __AddRenderCommand(renderCommand RenderCommand) {
	var context *Context = GetCurrentContext()
	if context.RenderCommands.Length < context.RenderCommands.Capacity-1 {
		RenderCommandArray_Add(&context.RenderCommands, renderCommand)
	} else {
		if !context.BooleanWarnings.MaxRenderCommandsExceeded {
			context.BooleanWarnings.MaxRenderCommandsExceeded = true
			context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay ran out of capacity while attempting to create render commands. This is usually caused by a large amount of wrapping text elements while close to the max element capacity. Try using SetMaxElementCount() with a higher value.")}, UserData: context.ErrorHandler.UserData.(any)})
		}
	}
}
func __ElementIsOffscreen(boundingBox *BoundingBox) bool {
	var context *Context = GetCurrentContext()
	if context.DisableCulling {
		return false
	}
	return boundingBox.X > context.LayoutDimensions.Width || boundingBox.Y > context.LayoutDimensions.Height || boundingBox.X+boundingBox.Width < 0 || boundingBox.Y+boundingBox.Height < 0
}
func __CalculateFinalLayout() {
	var context *Context = GetCurrentContext()
	__SizeContainersAlongAxis(true)
	for textElementIndex := int32(0); textElementIndex < context.TextElementData.Length; textElementIndex++ {
		var textElementData *__TextElementData = __TextElementDataArray_Get(&context.TextElementData, textElementIndex)
		textElementData.WrappedLines = __WrappedTextLineArraySlice{Length: 0, InternalArray: (*__WrappedTextLine)(unsafe.Add(unsafe.Pointer(context.WrappedTextLines.InternalArray), unsafe.Sizeof(__WrappedTextLine{})*uintptr(context.WrappedTextLines.Length)))}
		var containerElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, textElementData.ElementIndex)
		var textConfig *TextElementConfig = __FindElementConfigWithType(containerElement, __ELEMENT_CONFIG_TYPE_TEXT).TextElementConfig
		var measureTextCacheItem *__MeasureTextCacheItem = __MeasureTextCached(&textElementData.Text, textConfig)
		var lineWidth float32 = 0
		var lineHeight float32
		if int32(textConfig.LineHeight) > 0 {
			lineHeight = float32(textConfig.LineHeight)
		} else {
			lineHeight = textElementData.PreferredDimensions.Height
		}
		var lineLengthChars int32 = 0
		var lineStartOffset int32 = 0
		if !measureTextCacheItem.ContainsNewlines && textElementData.PreferredDimensions.Width <= containerElement.Dimensions.Width {
			__WrappedTextLineArray_Add(&context.WrappedTextLines, __WrappedTextLine{Dimensions: containerElement.Dimensions, Line: textElementData.Text})
			textElementData.WrappedLines.Length++
			continue
		}
		var spaceWidth float32 = __MeasureText(StringSlice{Length: 1, Chars: __SPACECHAR.Chars, BaseChars: __SPACECHAR.Chars}, textConfig, context.MeasureTextUserData.(unsafe.Pointer)).Width
		_ = spaceWidth
		var wordIndex int32 = measureTextCacheItem.MeasuredWordsStartIndex
		for wordIndex != -1 {
			if context.WrappedTextLines.Length > context.WrappedTextLines.Capacity-1 {
				break
			}
			var measuredWord *__MeasuredWord = __MeasuredWordArray_Get(&context.MeasuredWords, wordIndex)
			if lineLengthChars == 0 && lineWidth+measuredWord.Width > containerElement.Dimensions.Width {
				__WrappedTextLineArray_Add(&context.WrappedTextLines, __WrappedTextLine{Dimensions: Dimensions{Width: measuredWord.Width, Height: lineHeight}, Line: String{Length: measuredWord.Length, Chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), measuredWord.StartOffset))}})
				textElementData.WrappedLines.Length++
				wordIndex = measuredWord.Next
				lineStartOffset = measuredWord.StartOffset + measuredWord.Length
			} else if measuredWord.Length == 0 || lineWidth+measuredWord.Width > containerElement.Dimensions.Width {
				var finalCharIsSpace bool = *(*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), lineStartOffset+lineLengthChars-1)) == ' '
				__WrappedTextLineArray_Add(&context.WrappedTextLines, __WrappedTextLine{Dimensions: Dimensions{Width: lineWidth + (func() float32 {
					if finalCharIsSpace {
						return -spaceWidth
					}
					return 0
				}()), Height: lineHeight}, Line: String{Length: lineLengthChars + (func() int32 {
					if finalCharIsSpace {
						return -1
					}
					return 0
				}()), Chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), lineStartOffset))}})
				textElementData.WrappedLines.Length++
				if lineLengthChars == 0 || measuredWord.Length == 0 {
					wordIndex = measuredWord.Next
				}
				lineWidth = 0
				lineLengthChars = 0
				lineStartOffset = measuredWord.StartOffset
			} else {
				lineWidth += measuredWord.Width
				lineLengthChars += measuredWord.Length
				wordIndex = measuredWord.Next
			}
		}
		if lineLengthChars > 0 {
			__WrappedTextLineArray_Add(&context.WrappedTextLines, __WrappedTextLine{Dimensions: Dimensions{Width: lineWidth, Height: lineHeight}, Line: String{Length: lineLengthChars, Chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), lineStartOffset))}})
			textElementData.WrappedLines.Length++
		}
		containerElement.Dimensions.Height = lineHeight * float32(textElementData.WrappedLines.Length)
	}
	for i := int32(0); i < context.ImageElementPointers.Length; i++ {
		var (
			imageElement *LayoutElement      = LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&context.ImageElementPointers, i))
			config       *ImageElementConfig = __FindElementConfigWithType(imageElement, __ELEMENT_CONFIG_TYPE_IMAGE).ImageElementConfig
		)
		imageElement.Dimensions.Height = (config.SourceDimensions.Height / (func() float32 {
			if config.SourceDimensions.Width > 1 {
				return config.SourceDimensions.Width
			}
			return 1
		}())) * imageElement.Dimensions.Width
	}
	var dfsBuffer __LayoutElementTreeNodeArray = context.LayoutElementTreeNodeArray1
	dfsBuffer.Length = 0
	for i := int32(0); i < context.LayoutElementTreeRoots.Length; i++ {
		var root *__LayoutElementTreeRoot = __LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, i)
		*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length)) = false
		__LayoutElementTreeNodeArray_Add(&dfsBuffer, __LayoutElementTreeNode{LayoutElement: LayoutElementArray_Get(&context.LayoutElements, root.LayoutElementIndex)})
	}
	for dfsBuffer.Length > 0 {
		var (
			currentElementTreeNode *__LayoutElementTreeNode = __LayoutElementTreeNodeArray_Get(&dfsBuffer, dfsBuffer.Length-1)
			currentElement         *LayoutElement           = currentElementTreeNode.LayoutElement
		)
		if !*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length-1)) {
			*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length-1)) = true
			if __ElementHasConfig(currentElement, __ELEMENT_CONFIG_TYPE_TEXT) || int32(currentElement.ChildrenOrTextContent.Children.Length) == 0 {
				dfsBuffer.Length--
				continue
			}
			for i := int32(0); i < int32(currentElement.ChildrenOrTextContent.Children.Length); i++ {
				*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length)) = false
				__LayoutElementTreeNodeArray_Add(&dfsBuffer, __LayoutElementTreeNode{LayoutElement: LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))})
			}
			continue
		}
		dfsBuffer.Length--
		var layoutConfig *LayoutConfig = currentElement.LayoutConfig
		if layoutConfig.LayoutDirection == LEFT_TO_RIGHT {
			for j := int32(0); j < int32(currentElement.ChildrenOrTextContent.Children.Length); j++ {
				var (
					childElement           *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(j))))
					childHeightWithPadding float32        = (func() float32 {
						if (childElement.Dimensions.Height + float32(layoutConfig.Padding.Top) + float32(layoutConfig.Padding.Bottom)) > currentElement.Dimensions.Height {
							return childElement.Dimensions.Height + float32(layoutConfig.Padding.Top) + float32(layoutConfig.Padding.Bottom)
						}
						return currentElement.Dimensions.Height
					}())
				)
				if (func() float32 {
					if childHeightWithPadding > layoutConfig.Sizing.Height.Size.MinMax.Min {
						return childHeightWithPadding
					}
					return layoutConfig.Sizing.Height.Size.MinMax.Min
				}()) < layoutConfig.Sizing.Height.Size.MinMax.Max {
					if childHeightWithPadding > layoutConfig.Sizing.Height.Size.MinMax.Min {
						currentElement.Dimensions.Height = childHeightWithPadding
					} else {
						currentElement.Dimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Min
					}
				} else {
					currentElement.Dimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Max
				}
			}
		} else if layoutConfig.LayoutDirection == TOP_TO_BOTTOM {
			var contentHeight float32 = float32(int32(layoutConfig.Padding.Top) + int32(layoutConfig.Padding.Bottom))
			for j := int32(0); j < int32(currentElement.ChildrenOrTextContent.Children.Length); j++ {
				var childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(j))))
				contentHeight += childElement.Dimensions.Height
			}
			contentHeight += float32((func() int32 {
				if (int32(currentElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
					return int32(currentElement.ChildrenOrTextContent.Children.Length) - 1
				}
				return 0
			}()) * int32(layoutConfig.ChildGap))
			if (func() float32 {
				if contentHeight > layoutConfig.Sizing.Height.Size.MinMax.Min {
					return contentHeight
				}
				return layoutConfig.Sizing.Height.Size.MinMax.Min
			}()) < layoutConfig.Sizing.Height.Size.MinMax.Max {
				if contentHeight > layoutConfig.Sizing.Height.Size.MinMax.Min {
					currentElement.Dimensions.Height = contentHeight
				} else {
					currentElement.Dimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Min
				}
			} else {
				currentElement.Dimensions.Height = layoutConfig.Sizing.Height.Size.MinMax.Max
			}
		}
	}
	__SizeContainersAlongAxis(false)
	var sortMax int32 = context.LayoutElementTreeRoots.Length - 1
	for sortMax > 0 {
		for i := int32(0); i < sortMax; i++ {
			var (
				current __LayoutElementTreeRoot = *__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, i)
				next    __LayoutElementTreeRoot = *__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, i+1)
			)
			if int32(next.ZIndex) < int32(current.ZIndex) {
				__LayoutElementTreeRootArray_Set(&context.LayoutElementTreeRoots, i, next)
				__LayoutElementTreeRootArray_Set(&context.LayoutElementTreeRoots, i+1, current)
			}
		}
		sortMax--
	}
	context.RenderCommands.Length = 0
	dfsBuffer.Length = 0
	for rootIndex := int32(0); rootIndex < context.LayoutElementTreeRoots.Length; rootIndex++ {
		dfsBuffer.Length = 0
		var root *__LayoutElementTreeRoot = __LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, rootIndex)
		var rootElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, root.LayoutElementIndex)
		var rootPosition Vector2 = Vector2{}
		var parentHashMapItem *LayoutElementHashMapItem = __GetHashMapItem(root.ParentId)
		if __ElementHasConfig(rootElement, __ELEMENT_CONFIG_TYPE_FLOATING) && parentHashMapItem != nil {
			var (
				config               *FloatingElementConfig = __FindElementConfigWithType(rootElement, __ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig
				rootDimensions       Dimensions             = rootElement.Dimensions
				parentBoundingBox    BoundingBox            = parentHashMapItem.BoundingBox
				targetAttachPosition Vector2                = Vector2{}
			)
			switch config.AttachPoints.Parent {
			case ATTACH_POINT_LEFT_TOP:
				fallthrough
			case ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case ATTACH_POINT_LEFT_BOTTOM:
				targetAttachPosition.X = parentBoundingBox.X
			case ATTACH_POINT_CENTER_TOP:
				fallthrough
			case ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case ATTACH_POINT_CENTER_BOTTOM:
				targetAttachPosition.X = parentBoundingBox.X + parentBoundingBox.Width/2
			case ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case ATTACH_POINT_RIGHT_CENTER:
				fallthrough
			case ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.X = parentBoundingBox.X + parentBoundingBox.Width
			}
			switch config.AttachPoints.Element {
			case ATTACH_POINT_LEFT_TOP:
				fallthrough
			case ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case ATTACH_POINT_LEFT_BOTTOM:
			case ATTACH_POINT_CENTER_TOP:
				fallthrough
			case ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case ATTACH_POINT_CENTER_BOTTOM:
				targetAttachPosition.X -= rootDimensions.Width / 2
			case ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case ATTACH_POINT_RIGHT_CENTER:
				fallthrough
			case ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.X -= rootDimensions.Width
			}
			switch config.AttachPoints.Parent {
			case ATTACH_POINT_LEFT_TOP:
				fallthrough
			case ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case ATTACH_POINT_CENTER_TOP:
				targetAttachPosition.Y = parentBoundingBox.Y
			case ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case ATTACH_POINT_RIGHT_CENTER:
				targetAttachPosition.Y = parentBoundingBox.Y + parentBoundingBox.Height/2
			case ATTACH_POINT_LEFT_BOTTOM:
				fallthrough
			case ATTACH_POINT_CENTER_BOTTOM:
				fallthrough
			case ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.Y = parentBoundingBox.Y + parentBoundingBox.Height
			}
			switch config.AttachPoints.Element {
			case ATTACH_POINT_LEFT_TOP:
				fallthrough
			case ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case ATTACH_POINT_CENTER_TOP:
			case ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case ATTACH_POINT_RIGHT_CENTER:
				targetAttachPosition.Y -= rootDimensions.Height / 2
			case ATTACH_POINT_LEFT_BOTTOM:
				fallthrough
			case ATTACH_POINT_CENTER_BOTTOM:
				fallthrough
			case ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.Y -= rootDimensions.Height
			}
			targetAttachPosition.X += config.Offset.X
			targetAttachPosition.Y += config.Offset.Y
			rootPosition = targetAttachPosition
		}
		if root.ClipElementId != 0 {
			var clipHashMapItem *LayoutElementHashMapItem = __GetHashMapItem(root.ClipElementId)
			if clipHashMapItem != nil {
				if context.ExternalScrollHandlingEnabled {
					var scrollConfig *ScrollElementConfig = __FindElementConfigWithType(clipHashMapItem.LayoutElement, __ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
					for i := int32(0); i < context.ScrollContainerDatas.Length; i++ {
						var mapping *__ScrollContainerDataInternal = __ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
						if mapping.LayoutElement == clipHashMapItem.LayoutElement {
							root.PointerOffset = mapping.ScrollPosition
							if scrollConfig.Horizontal {
								rootPosition.X += mapping.ScrollPosition.X
							}
							if scrollConfig.Vertical {
								rootPosition.Y += mapping.ScrollPosition.Y
							}
							break
						}
					}
				}
				__AddRenderCommand(RenderCommand{BoundingBox: clipHashMapItem.BoundingBox, UserData: nil, Id: __HashNumber(rootElement.Id, uint32(int32(rootElement.ChildrenOrTextContent.Children.Length)+10)).Id, ZIndex: root.ZIndex, CommandType: RENDER_COMMAND_TYPE_SCISSOR_START})
			}
		}
		__LayoutElementTreeNodeArray_Add(&dfsBuffer, __LayoutElementTreeNode{LayoutElement: rootElement, Position: rootPosition, NextChildOffset: Vector2{X: float32(rootElement.LayoutConfig.Padding.Left), Y: float32(rootElement.LayoutConfig.Padding.Top)}})
		*context.TreeNodeVisited.InternalArray = false
		for dfsBuffer.Length > 0 {
			var (
				currentElementTreeNode *__LayoutElementTreeNode = __LayoutElementTreeNodeArray_Get(&dfsBuffer, dfsBuffer.Length-1)
				currentElement         *LayoutElement           = currentElementTreeNode.LayoutElement
				layoutConfig           *LayoutConfig            = currentElement.LayoutConfig
				scrollOffset           Vector2                  = Vector2{}
			)
			if !*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length-1)) {
				*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length-1)) = true
				var currentElementBoundingBox BoundingBox = BoundingBox{X: currentElementTreeNode.Position.X, Y: currentElementTreeNode.Position.Y, Width: currentElement.Dimensions.Width, Height: currentElement.Dimensions.Height}
				if __ElementHasConfig(currentElement, __ELEMENT_CONFIG_TYPE_FLOATING) {
					var (
						floatingElementConfig *FloatingElementConfig = __FindElementConfigWithType(currentElement, __ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig
						expand                Dimensions             = floatingElementConfig.Expand
					)
					currentElementBoundingBox.X -= expand.Width
					currentElementBoundingBox.Width += expand.Width * 2
					currentElementBoundingBox.Y -= expand.Height
					currentElementBoundingBox.Height += expand.Height * 2
				}
				var scrollContainerData *__ScrollContainerDataInternal = (*__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(__NULL)))
				if __ElementHasConfig(currentElement, __ELEMENT_CONFIG_TYPE_SCROLL) {
					var scrollConfig *ScrollElementConfig = __FindElementConfigWithType(currentElement, __ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
					for i := int32(0); i < context.ScrollContainerDatas.Length; i++ {
						var mapping *__ScrollContainerDataInternal = __ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
						if mapping.LayoutElement == currentElement {
							scrollContainerData = mapping
							mapping.BoundingBox = currentElementBoundingBox
							if scrollConfig.Horizontal {
								scrollOffset.X = mapping.ScrollPosition.X
							}
							if scrollConfig.Vertical {
								scrollOffset.Y = mapping.ScrollPosition.Y
							}
							if context.ExternalScrollHandlingEnabled {
								scrollOffset = Vector2{}
							}
							break
						}
					}
				}
				var hashMapItem *LayoutElementHashMapItem = __GetHashMapItem(currentElement.Id)
				if hashMapItem != nil {
					hashMapItem.BoundingBox = currentElementBoundingBox
					if hashMapItem.IdAlias != 0 {
						var hashMapItemAlias *LayoutElementHashMapItem = __GetHashMapItem(hashMapItem.IdAlias)
						if hashMapItemAlias != nil {
							hashMapItemAlias.BoundingBox = currentElementBoundingBox
						}
					}
				}
				var sortedConfigIndexes [20]int32
				for elementConfigIndex := int32(0); elementConfigIndex < currentElement.ElementConfigs.Length; elementConfigIndex++ {
					sortedConfigIndexes[elementConfigIndex] = elementConfigIndex
				}
				sortMax = currentElement.ElementConfigs.Length - 1
				for sortMax > 0 {
					for i := int32(0); i < sortMax; i++ {
						var (
							current     int32               = sortedConfigIndexes[i]
							next        int32               = sortedConfigIndexes[i+1]
							currentType __ElementConfigType = __ElementConfigArraySlice_Get(&currentElement.ElementConfigs, current).Type
							nextType    __ElementConfigType = __ElementConfigArraySlice_Get(&currentElement.ElementConfigs, next).Type
						)
						if nextType == __ELEMENT_CONFIG_TYPE_SCROLL || currentType == __ELEMENT_CONFIG_TYPE_BORDER {
							sortedConfigIndexes[i] = next
							sortedConfigIndexes[i+1] = current
						}
					}
					sortMax--
				}
				var emitRectangle bool = false
				var sharedConfig *SharedElementConfig = __FindElementConfigWithType(currentElement, __ELEMENT_CONFIG_TYPE_SHARED).SharedElementConfig
				if sharedConfig != nil && sharedConfig.BackgroundColor.A > 0 {
					emitRectangle = true
				} else if sharedConfig == nil {
					emitRectangle = false
					sharedConfig = &SharedElementConfig_DEFAULT
				}
				for elementConfigIndex := int32(0); elementConfigIndex < currentElement.ElementConfigs.Length; elementConfigIndex++ {
					var (
						elementConfig *ElementConfig = __ElementConfigArraySlice_Get(&currentElement.ElementConfigs, sortedConfigIndexes[elementConfigIndex])
						renderCommand RenderCommand  = RenderCommand{BoundingBox: currentElementBoundingBox, UserData: sharedConfig.UserData, Id: currentElement.Id}
						offscreen     bool           = __ElementIsOffscreen(&currentElementBoundingBox)
						shouldRender  bool           = !offscreen
					)
					switch elementConfig.Type {
					case __ELEMENT_CONFIG_TYPE_FLOATING:
						fallthrough
					case __ELEMENT_CONFIG_TYPE_SHARED:
						fallthrough
					case __ELEMENT_CONFIG_TYPE_BORDER:
						shouldRender = false
					case __ELEMENT_CONFIG_TYPE_SCROLL:
						renderCommand.CommandType = RENDER_COMMAND_TYPE_SCISSOR_START
						renderCommand.RenderData = RenderData{Scroll: ScrollRenderData{Horizontal: elementConfig.Config.ScrollElementConfig.Horizontal, Vertical: elementConfig.Config.ScrollElementConfig.Vertical}}
					case __ELEMENT_CONFIG_TYPE_IMAGE:
						renderCommand.CommandType = RENDER_COMMAND_TYPE_IMAGE
						renderCommand.RenderData = RenderData{Image: ImageRenderData{BackgroundColor: sharedConfig.BackgroundColor, CornerRadius: sharedConfig.CornerRadius, SourceDimensions: elementConfig.Config.ImageElementConfig.SourceDimensions, ImageData: elementConfig.Config.ImageElementConfig.ImageData}}
						emitRectangle = false
					case __ELEMENT_CONFIG_TYPE_TEXT:
						if !shouldRender {
							break
						}
						shouldRender = false
						var configUnion ElementConfigUnion = elementConfig.Config
						var textElementConfig *TextElementConfig = configUnion.TextElementConfig
						var naturalLineHeight float32 = currentElement.ChildrenOrTextContent.TextElementData.PreferredDimensions.Height
						var finalLineHeight float32
						if int32(textElementConfig.LineHeight) > 0 {
							finalLineHeight = float32(textElementConfig.LineHeight)
						} else {
							finalLineHeight = naturalLineHeight
						}
						var lineHeightOffset float32 = (finalLineHeight - naturalLineHeight) / 2
						var yPosition float32 = lineHeightOffset
						for lineIndex := int32(0); lineIndex < currentElement.ChildrenOrTextContent.TextElementData.WrappedLines.Length; lineIndex++ {
							var wrappedLine *__WrappedTextLine = __WrappedTextLineArraySlice_Get(&currentElement.ChildrenOrTextContent.TextElementData.WrappedLines, lineIndex)
							if wrappedLine.Line.Length == 0 {
								yPosition += finalLineHeight
								continue
							}
							var offset float32 = (currentElementBoundingBox.Width - wrappedLine.Dimensions.Width)
							if textElementConfig.TextAlignment == TEXT_ALIGN_LEFT {
								offset = 0
							}
							if textElementConfig.TextAlignment == TEXT_ALIGN_CENTER {
								offset /= 2
							}
							__AddRenderCommand(RenderCommand{BoundingBox: BoundingBox{X: currentElementBoundingBox.X + offset, Y: currentElementBoundingBox.Y + yPosition, Width: wrappedLine.Dimensions.Width, Height: wrappedLine.Dimensions.Height}, RenderData: RenderData{Text: TextRenderData{StringContents: StringSlice{Length: wrappedLine.Line.Length, Chars: wrappedLine.Line.Chars, BaseChars: currentElement.ChildrenOrTextContent.TextElementData.Text.Chars}, TextColor: textElementConfig.TextColor, FontId: textElementConfig.FontId, FontSize: textElementConfig.FontSize, LetterSpacing: textElementConfig.LetterSpacing, LineHeight: textElementConfig.LineHeight}}, UserData: sharedConfig.UserData, Id: __HashNumber(uint32(lineIndex), currentElement.Id).Id, ZIndex: root.ZIndex, CommandType: RENDER_COMMAND_TYPE_TEXT})
							yPosition += finalLineHeight
							if !context.DisableCulling && currentElementBoundingBox.Y+yPosition > context.LayoutDimensions.Height {
								break
							}
						}
					case __ELEMENT_CONFIG_TYPE_CUSTOM:
						renderCommand.CommandType = RENDER_COMMAND_TYPE_CUSTOM
						renderCommand.RenderData = RenderData{Custom: CustomRenderData{BackgroundColor: sharedConfig.BackgroundColor, CornerRadius: sharedConfig.CornerRadius, CustomData: elementConfig.Config.CustomElementConfig.CustomData}}
						emitRectangle = false
					default:
					}
					if shouldRender {
						__AddRenderCommand(renderCommand)
					}
					if offscreen {
					}
				}
				if emitRectangle {
					__AddRenderCommand(RenderCommand{BoundingBox: currentElementBoundingBox, RenderData: RenderData{Rectangle: RectangleRenderData{BackgroundColor: sharedConfig.BackgroundColor, CornerRadius: sharedConfig.CornerRadius}}, UserData: sharedConfig.UserData, Id: currentElement.Id, ZIndex: root.ZIndex, CommandType: RENDER_COMMAND_TYPE_RECTANGLE})
				}
				if !__ElementHasConfig(currentElementTreeNode.LayoutElement, __ELEMENT_CONFIG_TYPE_TEXT) {
					var contentSize Dimensions = Dimensions{}
					if layoutConfig.LayoutDirection == LEFT_TO_RIGHT {
						for i := int32(0); i < int32(currentElement.ChildrenOrTextContent.Children.Length); i++ {
							var childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
							contentSize.Width += childElement.Dimensions.Width
							if contentSize.Height > childElement.Dimensions.Height {
								contentSize.Height = contentSize.Height
							} else {
								contentSize.Height = childElement.Dimensions.Height
							}
						}
						contentSize.Width += float32((func() int32 {
							if (int32(currentElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
								return int32(currentElement.ChildrenOrTextContent.Children.Length) - 1
							}
							return 0
						}()) * int32(layoutConfig.ChildGap))
						var extraSpace float32 = currentElement.Dimensions.Width - float32(int32(layoutConfig.Padding.Left)+int32(layoutConfig.Padding.Right)) - contentSize.Width
						switch layoutConfig.ChildAlignment.X {
						case ALIGN_X_LEFT:
							extraSpace = 0
						case ALIGN_X_CENTER:
							extraSpace /= 2
						default:
						}
						currentElementTreeNode.NextChildOffset.X += extraSpace
					} else {
						for i := int32(0); i < int32(currentElement.ChildrenOrTextContent.Children.Length); i++ {
							var childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
							if contentSize.Width > childElement.Dimensions.Width {
								contentSize.Width = contentSize.Width
							} else {
								contentSize.Width = childElement.Dimensions.Width
							}
							contentSize.Height += childElement.Dimensions.Height
						}
						contentSize.Height += float32((func() int32 {
							if (int32(currentElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
								return int32(currentElement.ChildrenOrTextContent.Children.Length) - 1
							}
							return 0
						}()) * int32(layoutConfig.ChildGap))
						var extraSpace float32 = currentElement.Dimensions.Height - float32(int32(layoutConfig.Padding.Top)+int32(layoutConfig.Padding.Bottom)) - contentSize.Height
						switch layoutConfig.ChildAlignment.Y {
						case ALIGN_Y_TOP:
							extraSpace = 0
						case ALIGN_Y_CENTER:
							extraSpace /= 2
						default:
						}
						currentElementTreeNode.NextChildOffset.Y += extraSpace
					}
					if scrollContainerData != nil {
						scrollContainerData.ContentSize = Dimensions{Width: contentSize.Width + float32(int32(layoutConfig.Padding.Left)+int32(layoutConfig.Padding.Right)), Height: contentSize.Height + float32(int32(layoutConfig.Padding.Top)+int32(layoutConfig.Padding.Bottom))}
					}
				}
			} else {
				var (
					closeScrollElement bool                 = false
					scrollConfig       *ScrollElementConfig = __FindElementConfigWithType(currentElement, __ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
				)
				if scrollConfig != nil {
					closeScrollElement = true
					for i := int32(0); i < context.ScrollContainerDatas.Length; i++ {
						var mapping *__ScrollContainerDataInternal = __ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
						if mapping.LayoutElement == currentElement {
							if scrollConfig.Horizontal {
								scrollOffset.X = mapping.ScrollPosition.X
							}
							if scrollConfig.Vertical {
								scrollOffset.Y = mapping.ScrollPosition.Y
							}
							if context.ExternalScrollHandlingEnabled {
								scrollOffset = Vector2{}
							}
							break
						}
					}
				}
				if __ElementHasConfig(currentElement, __ELEMENT_CONFIG_TYPE_BORDER) {
					var (
						currentElementData        *LayoutElementHashMapItem = __GetHashMapItem(currentElement.Id)
						currentElementBoundingBox BoundingBox               = currentElementData.BoundingBox
					)
					if !__ElementIsOffscreen(&currentElementBoundingBox) {
						var sharedConfig *SharedElementConfig
						if __ElementHasConfig(currentElement, __ELEMENT_CONFIG_TYPE_SHARED) {
							sharedConfig = __FindElementConfigWithType(currentElement, __ELEMENT_CONFIG_TYPE_SHARED).SharedElementConfig
						} else {
							sharedConfig = &SharedElementConfig_DEFAULT
						}
						var borderConfig *BorderElementConfig = __FindElementConfigWithType(currentElement, __ELEMENT_CONFIG_TYPE_BORDER).BorderElementConfig
						var renderCommand RenderCommand = RenderCommand{BoundingBox: currentElementBoundingBox, RenderData: RenderData{Border: BorderRenderData{Color: borderConfig.Color, CornerRadius: sharedConfig.CornerRadius, Width: borderConfig.Width}}, UserData: sharedConfig.UserData, Id: __HashNumber(currentElement.Id, uint32(currentElement.ChildrenOrTextContent.Children.Length)).Id, CommandType: RENDER_COMMAND_TYPE_BORDER}
						__AddRenderCommand(renderCommand)
						if int32(borderConfig.Width.BetweenChildren) > 0 && borderConfig.Color.A > 0 {
							var (
								halfGap      float32 = float32(int32(layoutConfig.ChildGap) / 2)
								borderOffset Vector2 = Vector2{X: float32(layoutConfig.Padding.Left) - halfGap, Y: float32(layoutConfig.Padding.Top) - halfGap}
							)
							if layoutConfig.LayoutDirection == LEFT_TO_RIGHT {
								for i := int32(0); i < int32(currentElement.ChildrenOrTextContent.Children.Length); i++ {
									var childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
									if i > 0 {
										__AddRenderCommand(RenderCommand{BoundingBox: BoundingBox{X: currentElementBoundingBox.X + borderOffset.X + scrollOffset.X, Y: currentElementBoundingBox.Y + scrollOffset.Y, Width: float32(borderConfig.Width.BetweenChildren), Height: currentElement.Dimensions.Height}, RenderData: RenderData{Rectangle: RectangleRenderData{BackgroundColor: borderConfig.Color}}, UserData: sharedConfig.UserData, Id: __HashNumber(currentElement.Id, uint32(int32(currentElement.ChildrenOrTextContent.Children.Length)+1+i)).Id, CommandType: RENDER_COMMAND_TYPE_RECTANGLE})
									}
									borderOffset.X += childElement.Dimensions.Width + float32(layoutConfig.ChildGap)
								}
							} else {
								for i := int32(0); i < int32(currentElement.ChildrenOrTextContent.Children.Length); i++ {
									var childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
									if i > 0 {
										__AddRenderCommand(RenderCommand{BoundingBox: BoundingBox{X: currentElementBoundingBox.X + scrollOffset.X, Y: currentElementBoundingBox.Y + borderOffset.Y + scrollOffset.Y, Width: currentElement.Dimensions.Width, Height: float32(borderConfig.Width.BetweenChildren)}, RenderData: RenderData{Rectangle: RectangleRenderData{BackgroundColor: borderConfig.Color}}, UserData: sharedConfig.UserData, Id: __HashNumber(currentElement.Id, uint32(int32(currentElement.ChildrenOrTextContent.Children.Length)+1+i)).Id, CommandType: RENDER_COMMAND_TYPE_RECTANGLE})
									}
									borderOffset.Y += childElement.Dimensions.Height + float32(layoutConfig.ChildGap)
								}
							}
						}
					}
				}
				if closeScrollElement {
					__AddRenderCommand(RenderCommand{Id: __HashNumber(currentElement.Id, uint32(int32(rootElement.ChildrenOrTextContent.Children.Length)+11)).Id, CommandType: RENDER_COMMAND_TYPE_SCISSOR_END})
				}
				dfsBuffer.Length--
				continue
			}
			if !__ElementHasConfig(currentElement, __ELEMENT_CONFIG_TYPE_TEXT) {
				dfsBuffer.Length += int32(currentElement.ChildrenOrTextContent.Children.Length)
				for i := int32(0); i < int32(currentElement.ChildrenOrTextContent.Children.Length); i++ {
					var childElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
					if layoutConfig.LayoutDirection == LEFT_TO_RIGHT {
						currentElementTreeNode.NextChildOffset.Y = float32(currentElement.LayoutConfig.Padding.Top)
						var whiteSpaceAroundChild float32 = currentElement.Dimensions.Height - float32(int32(layoutConfig.Padding.Top)+int32(layoutConfig.Padding.Bottom)) - childElement.Dimensions.Height
						switch layoutConfig.ChildAlignment.Y {
						case ALIGN_Y_TOP:
						case ALIGN_Y_CENTER:
							currentElementTreeNode.NextChildOffset.Y += whiteSpaceAroundChild / 2
						case ALIGN_Y_BOTTOM:
							currentElementTreeNode.NextChildOffset.Y += whiteSpaceAroundChild
						}
					} else {
						currentElementTreeNode.NextChildOffset.X = float32(currentElement.LayoutConfig.Padding.Left)
						var whiteSpaceAroundChild float32 = currentElement.Dimensions.Width - float32(int32(layoutConfig.Padding.Left)+int32(layoutConfig.Padding.Right)) - childElement.Dimensions.Width
						switch layoutConfig.ChildAlignment.X {
						case ALIGN_X_LEFT:
						case ALIGN_X_CENTER:
							currentElementTreeNode.NextChildOffset.X += whiteSpaceAroundChild / 2
						case ALIGN_X_RIGHT:
							currentElementTreeNode.NextChildOffset.X += whiteSpaceAroundChild
						}
					}
					var childPosition Vector2 = Vector2{X: currentElementTreeNode.Position.X + currentElementTreeNode.NextChildOffset.X + scrollOffset.X, Y: currentElementTreeNode.Position.Y + currentElementTreeNode.NextChildOffset.Y + scrollOffset.Y}
					var newNodeIndex uint32 = uint32(dfsBuffer.Length - 1 - i)
					*(*__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(dfsBuffer.InternalArray), unsafe.Sizeof(__LayoutElementTreeNode{})*uintptr(newNodeIndex))) = __LayoutElementTreeNode{LayoutElement: childElement, Position: Vector2{X: childPosition.X, Y: childPosition.Y}, NextChildOffset: Vector2{X: float32(childElement.LayoutConfig.Padding.Left), Y: float32(childElement.LayoutConfig.Padding.Top)}}
					*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), newNodeIndex)) = false
					if layoutConfig.LayoutDirection == LEFT_TO_RIGHT {
						currentElementTreeNode.NextChildOffset.X += childElement.Dimensions.Width + float32(layoutConfig.ChildGap)
					} else {
						currentElementTreeNode.NextChildOffset.Y += childElement.Dimensions.Height + float32(layoutConfig.ChildGap)
					}
				}
			}
		}
		if root.ClipElementId != 0 {
			__AddRenderCommand(RenderCommand{Id: __HashNumber(rootElement.Id, uint32(int32(rootElement.ChildrenOrTextContent.Children.Length)+11)).Id, CommandType: RENDER_COMMAND_TYPE_SCISSOR_END})
		}
	}
}

var __debugViewWidth uint32 = 400
var __debugViewHighlightColor Color = Color{R: 168, G: 66, B: 28, A: 100}

func __WarningArray_Allocate_Arena(capacity int32, arena *Arena) __WarningArray {
	var (
		totalSizeBytes  uint64         = uint64(uintptr(capacity) * unsafe.Sizeof(String{}))
		array           __WarningArray = __WarningArray{Capacity: capacity, Length: 0}
		nextAllocOffset uint64         = arena.NextAllocation + (64 - arena.NextAllocation%64)
	)
	if nextAllocOffset+totalSizeBytes <= arena.Capacity {
		array.InternalArray = (*__Warning)(unsafe.Pointer(uintptr(uint64(uintptr(unsafe.Pointer(arena.Memory))) + nextAllocOffset)))
		arena.NextAllocation = nextAllocOffset + totalSizeBytes
	} else {
		__currentContext.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_ARENA_CAPACITY_EXCEEDED, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to allocate memory in its arena, but ran out of capacity. Try increasing the capacity of the arena passed to Initialize()")}, UserData: __currentContext.ErrorHandler.UserData.(any)})
	}
	return array
}
func __WarningArray_Add(array *__WarningArray, item __Warning) *__Warning {
	if array.Length < array.Capacity {
		*(*__Warning)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__Warning{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*__Warning)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(__Warning{})*uintptr(array.Length-1)))
	}
	return &__WARNING_DEFAULT
}
func __Array_Allocate_Arena(capacity int32, itemSize uint32, arena *Arena) unsafe.Pointer {
	var (
		totalSizeBytes  uint64 = uint64(uint32(capacity) * itemSize)
		nextAllocOffset uint64 = arena.NextAllocation + (64 - arena.NextAllocation%64)
	)
	if nextAllocOffset+totalSizeBytes <= arena.Capacity {
		arena.NextAllocation = nextAllocOffset + totalSizeBytes
		return unsafe.Pointer(uintptr(uint64(uintptr(unsafe.Pointer(arena.Memory))) + nextAllocOffset))
	} else {
		__currentContext.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_ARENA_CAPACITY_EXCEEDED, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to allocate memory in its arena, but ran out of capacity. Try increasing the capacity of the arena passed to Initialize()")}, UserData: __currentContext.ErrorHandler.UserData.(any)})
	}
	return unsafe.Pointer(uintptr(__NULL))
}
func __Array_RangeCheck(index int32, length int32) bool {
	if index < length && index >= 0 {
		return true
	}
	var context *Context = GetCurrentContext()
	context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_INTERNAL_ERROR, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to make an out of bounds array access. This is an internal error and is likely a bug.")}, UserData: context.ErrorHandler.UserData.(any)})
	return false
}
func __Array_AddCapacityCheck(length int32, capacity int32) bool {
	if length < capacity {
		return true
	}
	var context *Context = GetCurrentContext()
	context.ErrorHandler.ErrorHandlerFunction(ErrorData{ErrorType: ERROR_TYPE_INTERNAL_ERROR, ErrorText: String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to make an out of bounds array access. This is an internal error and is likely a bug.")}, UserData: context.ErrorHandler.UserData.(any)})
	return false
}
func MinMemorySize() uint32 {
	var (
		fakeContext    Context  = Context{MaxElementCount: __defaultMaxElementCount, MaxMeasureTextCacheWordCount: __defaultMaxMeasureTextWordCacheCount, InternalArena: Arena{Capacity: math.MaxUint64, Memory: nil}}
		currentContext *Context = GetCurrentContext()
	)
	if currentContext != nil {
		fakeContext.MaxElementCount = currentContext.MaxElementCount
		fakeContext.MaxMeasureTextCacheWordCount = currentContext.MaxElementCount
	}
	__Context_Allocate_Arena(&fakeContext.InternalArena)
	__InitializePersistentMemory(&fakeContext)
	__InitializeEphemeralMemory(&fakeContext)
	return uint32(fakeContext.InternalArena.NextAllocation + 128)
}
func CreateArenaWithCapacityAndMemory(capacity uint32, memory unsafe.Pointer) Arena {
	var arena Arena = Arena{Capacity: uint64(capacity), Memory: (*byte)(memory)}
	return arena
}
func SetMeasureTextFunction(measureTextFunction func(text StringSlice, config *TextElementConfig, userData unsafe.Pointer) Dimensions, userData any) {
	var context *Context = GetCurrentContext()
	__MeasureText = measureTextFunction
	context.MeasureTextUserData = userData.(any).(any)
}
func SetQueryScrollOffsetFunction(queryScrollOffsetFunction func(elementId uint32, userData unsafe.Pointer) Vector2, userData any) {
	var context *Context = GetCurrentContext()
	__QueryScrollOffset = queryScrollOffsetFunction
	context.QueryScrollOffsetUserData = userData.(any).(any)
}
func SetLayoutDimensions(dimensions Dimensions) {
	GetCurrentContext().LayoutDimensions = dimensions
}
func SetPointerState(position Vector2, isPointerDown bool) {
	var context *Context = GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return
	}
	context.PointerInfo.Position = position
	context.PointerOverIds.Length = 0
	var dfsBuffer __int32_tArray = context.LayoutElementChildrenBuffer
	for rootIndex := int32(context.LayoutElementTreeRoots.Length - 1); rootIndex >= 0; rootIndex-- {
		dfsBuffer.Length = 0
		var root *__LayoutElementTreeRoot = __LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, rootIndex)
		__int32_tArray_Add(&dfsBuffer, root.LayoutElementIndex)
		*context.TreeNodeVisited.InternalArray = false
		var found bool = false
		for dfsBuffer.Length > 0 {
			if *(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length-1)) {
				dfsBuffer.Length--
				continue
			}
			*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length-1)) = true
			var currentElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, __int32_tArray_GetValue(&dfsBuffer, dfsBuffer.Length-1))
			var mapItem *LayoutElementHashMapItem = __GetHashMapItem(currentElement.Id)
			var elementBox BoundingBox = mapItem.BoundingBox
			elementBox.X -= root.PointerOffset.X
			elementBox.Y -= root.PointerOffset.Y
			if mapItem != nil {
				if __PointIsInsideRect(position, elementBox) {
					if mapItem.OnHoverFunction != nil {
						mapItem.OnHoverFunction(mapItem.ElementId, context.PointerInfo, mapItem.HoverFunctionUserData)
					}
					__ElementIdArray_Add(&context.PointerOverIds, mapItem.ElementId)
					found = true
					if mapItem.IdAlias != 0 {
						__ElementIdArray_Add(&context.PointerOverIds, ElementId{Id: mapItem.IdAlias})
					}
				}
				if __ElementHasConfig(currentElement, __ELEMENT_CONFIG_TYPE_TEXT) {
					dfsBuffer.Length--
					continue
				}
				for i := int32(int32(currentElement.ChildrenOrTextContent.Children.Length) - 1); i >= 0; i-- {
					__int32_tArray_Add(&dfsBuffer, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
					*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length-1)) = false
				}
			} else {
				dfsBuffer.Length--
			}
		}
		var rootElement *LayoutElement = LayoutElementArray_Get(&context.LayoutElements, root.LayoutElementIndex)
		if found && __ElementHasConfig(rootElement, __ELEMENT_CONFIG_TYPE_FLOATING) && __FindElementConfigWithType(rootElement, __ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig.PointerCaptureMode == POINTER_CAPTURE_MODE_CAPTURE {
			break
		}
	}
	if isPointerDown {
		if context.PointerInfo.State == POINTER_DATA_PRESSED_THIS_FRAME {
			context.PointerInfo.State = POINTER_DATA_PRESSED
		} else if context.PointerInfo.State != POINTER_DATA_PRESSED {
			context.PointerInfo.State = POINTER_DATA_PRESSED_THIS_FRAME
		}
	} else {
		if context.PointerInfo.State == POINTER_DATA_RELEASED_THIS_FRAME {
			context.PointerInfo.State = POINTER_DATA_RELEASED
		} else if context.PointerInfo.State != POINTER_DATA_RELEASED {
			context.PointerInfo.State = POINTER_DATA_RELEASED_THIS_FRAME
		}
	}
}
func Initialize(arena Arena, layoutDimensions Dimensions, errorHandler ErrorHandler) *Context {
	var context *Context = __Context_Allocate_Arena(&arena)
	if context == nil {
		return nil
	}
	var oldContext *Context = GetCurrentContext()
	*context = Context{MaxElementCount: func() int32 {
		if oldContext != nil {
			return oldContext.MaxElementCount
		}
		return __defaultMaxElementCount
	}(), MaxMeasureTextCacheWordCount: func() int32 {
		if oldContext != nil {
			return oldContext.MaxMeasureTextCacheWordCount
		}
		return __defaultMaxMeasureTextWordCacheCount
	}(), ErrorHandler: func() ErrorHandler {
		if errorHandler.ErrorHandlerFunction != nil {
			return errorHandler
		}
		return ErrorHandler{ErrorHandlerFunction: __ErrorHandlerFunctionDefault, UserData: 0}
	}(), LayoutDimensions: layoutDimensions, InternalArena: arena}
	SetCurrentContext(context)
	__InitializePersistentMemory(context)
	__InitializeEphemeralMemory(context)
	for i := int32(0); i < context.LayoutElementsHashMap.Capacity; i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = -1
	}
	for i := int32(0); i < context.MeasureTextHashMap.Capacity; i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = 0
	}
	context.MeasureTextHashMapInternal.Length = 1
	context.LayoutDimensions = layoutDimensions
	return context
}
func GetCurrentContext() *Context {
	return __currentContext
}
func SetCurrentContext(context *Context) {
	__currentContext = context
}
func UpdateScrollContainers(enableDragScrolling bool, scrollDelta Vector2, deltaTime float32) {
	var (
		context                     *Context                       = GetCurrentContext()
		isPointerActive             bool                           = enableDragScrolling && (context.PointerInfo.State == POINTER_DATA_PRESSED || context.PointerInfo.State == POINTER_DATA_PRESSED_THIS_FRAME)
		highestPriorityElementIndex int32                          = -1
		highestPriorityScrollData   *__ScrollContainerDataInternal = (*__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(__NULL)))
	)
	for i := int32(0); i < context.ScrollContainerDatas.Length; i++ {
		var scrollData *__ScrollContainerDataInternal = __ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
		if !scrollData.OpenThisFrame {
			__ScrollContainerDataInternalArray_RemoveSwapback(&context.ScrollContainerDatas, i)
			continue
		}
		scrollData.OpenThisFrame = false
		var hashMapItem *LayoutElementHashMapItem = __GetHashMapItem(scrollData.ElementId)
		if hashMapItem == nil {
			__ScrollContainerDataInternalArray_RemoveSwapback(&context.ScrollContainerDatas, i)
			continue
		}
		if !isPointerActive && scrollData.PointerScrollActive {
			var xDiff float32 = scrollData.ScrollPosition.X - scrollData.ScrollOrigin.X
			if xDiff < -10 || xDiff > 10 {
				scrollData.ScrollMomentum.X = (scrollData.ScrollPosition.X - scrollData.ScrollOrigin.X) / (scrollData.MomentumTime * 25)
			}
			var yDiff float32 = scrollData.ScrollPosition.Y - scrollData.ScrollOrigin.Y
			if yDiff < -10 || yDiff > 10 {
				scrollData.ScrollMomentum.Y = (scrollData.ScrollPosition.Y - scrollData.ScrollOrigin.Y) / (scrollData.MomentumTime * 25)
			}
			scrollData.PointerScrollActive = false
			scrollData.PointerOrigin = Vector2{}
			scrollData.ScrollOrigin = Vector2{}
			scrollData.MomentumTime = 0
		}
		scrollData.ScrollPosition.X += scrollData.ScrollMomentum.X
		scrollData.ScrollMomentum.X *= 0.95
		var scrollOccurred bool = scrollDelta.X != 0 || scrollDelta.Y != 0
		if scrollData.ScrollMomentum.X > -0.1 && scrollData.ScrollMomentum.X < 0.1 || scrollOccurred {
			scrollData.ScrollMomentum.X = 0
		}
		if (func() float32 {
			if scrollData.ScrollPosition.X > (-func() float32 {
				if (scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width) > 0 {
					return scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width
				}
				return 0
			}()) {
				return scrollData.ScrollPosition.X
			}
			return -func() float32 {
				if (scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width) > 0 {
					return scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width
				}
				return 0
			}()
		}()) < 0 {
			if scrollData.ScrollPosition.X > (-func() float32 {
				if (scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width) > 0 {
					return scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width
				}
				return 0
			}()) {
				scrollData.ScrollPosition.X = scrollData.ScrollPosition.X
			} else if (scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width) > 0 {
				scrollData.ScrollPosition.X = -(scrollData.ContentSize.Width - scrollData.LayoutElement.Dimensions.Width)
			} else {
				scrollData.ScrollPosition.X = 0
			}
		} else {
			scrollData.ScrollPosition.X = 0
		}
		scrollData.ScrollPosition.Y += scrollData.ScrollMomentum.Y
		scrollData.ScrollMomentum.Y *= 0.95
		if scrollData.ScrollMomentum.Y > -0.1 && scrollData.ScrollMomentum.Y < 0.1 || scrollOccurred {
			scrollData.ScrollMomentum.Y = 0
		}
		if (func() float32 {
			if scrollData.ScrollPosition.Y > (-func() float32 {
				if (scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height) > 0 {
					return scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height
				}
				return 0
			}()) {
				return scrollData.ScrollPosition.Y
			}
			return -func() float32 {
				if (scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height) > 0 {
					return scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height
				}
				return 0
			}()
		}()) < 0 {
			if scrollData.ScrollPosition.Y > (-func() float32 {
				if (scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height) > 0 {
					return scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height
				}
				return 0
			}()) {
				scrollData.ScrollPosition.Y = scrollData.ScrollPosition.Y
			} else if (scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height) > 0 {
				scrollData.ScrollPosition.Y = -(scrollData.ContentSize.Height - scrollData.LayoutElement.Dimensions.Height)
			} else {
				scrollData.ScrollPosition.Y = 0
			}
		} else {
			scrollData.ScrollPosition.Y = 0
		}
		for j := int32(0); j < context.PointerOverIds.Length; j++ {
			if scrollData.LayoutElement.Id == __ElementIdArray_Get(&context.PointerOverIds, j).Id {
				highestPriorityElementIndex = j
				highestPriorityScrollData = scrollData
			}
		}
	}
	if highestPriorityElementIndex > -1 && highestPriorityScrollData != nil {
		var (
			scrollElement         *LayoutElement       = highestPriorityScrollData.LayoutElement
			scrollConfig          *ScrollElementConfig = __FindElementConfigWithType(scrollElement, __ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
			canScrollVertically   bool                 = scrollConfig.Vertical && highestPriorityScrollData.ContentSize.Height > scrollElement.Dimensions.Height
			canScrollHorizontally bool                 = scrollConfig.Horizontal && highestPriorityScrollData.ContentSize.Width > scrollElement.Dimensions.Width
		)
		if canScrollVertically {
			highestPriorityScrollData.ScrollPosition.Y = highestPriorityScrollData.ScrollPosition.Y + scrollDelta.Y*10
		}
		if canScrollHorizontally {
			highestPriorityScrollData.ScrollPosition.X = highestPriorityScrollData.ScrollPosition.X + scrollDelta.X*10
		}
		if isPointerActive {
			highestPriorityScrollData.ScrollMomentum = Vector2{}
			if !highestPriorityScrollData.PointerScrollActive {
				highestPriorityScrollData.PointerOrigin = context.PointerInfo.Position
				highestPriorityScrollData.ScrollOrigin = highestPriorityScrollData.ScrollPosition
				highestPriorityScrollData.PointerScrollActive = true
			} else {
				var (
					scrollDeltaX float32 = 0
					scrollDeltaY float32 = 0
				)
				if canScrollHorizontally {
					var oldXScrollPosition float32 = highestPriorityScrollData.ScrollPosition.X
					highestPriorityScrollData.ScrollPosition.X = highestPriorityScrollData.ScrollOrigin.X + (context.PointerInfo.Position.X - highestPriorityScrollData.PointerOrigin.X)
					if (func() float32 {
						if highestPriorityScrollData.ScrollPosition.X < 0 {
							return highestPriorityScrollData.ScrollPosition.X
						}
						return 0
					}()) > (-(highestPriorityScrollData.ContentSize.Width - highestPriorityScrollData.BoundingBox.Width)) {
						if highestPriorityScrollData.ScrollPosition.X < 0 {
							highestPriorityScrollData.ScrollPosition.X = highestPriorityScrollData.ScrollPosition.X
						} else {
							highestPriorityScrollData.ScrollPosition.X = 0
						}
					} else {
						highestPriorityScrollData.ScrollPosition.X = -(highestPriorityScrollData.ContentSize.Width - highestPriorityScrollData.BoundingBox.Width)
					}
					scrollDeltaX = highestPriorityScrollData.ScrollPosition.X - oldXScrollPosition
				}
				if canScrollVertically {
					var oldYScrollPosition float32 = highestPriorityScrollData.ScrollPosition.Y
					highestPriorityScrollData.ScrollPosition.Y = highestPriorityScrollData.ScrollOrigin.Y + (context.PointerInfo.Position.Y - highestPriorityScrollData.PointerOrigin.Y)
					if (func() float32 {
						if highestPriorityScrollData.ScrollPosition.Y < 0 {
							return highestPriorityScrollData.ScrollPosition.Y
						}
						return 0
					}()) > (-(highestPriorityScrollData.ContentSize.Height - highestPriorityScrollData.BoundingBox.Height)) {
						if highestPriorityScrollData.ScrollPosition.Y < 0 {
							highestPriorityScrollData.ScrollPosition.Y = highestPriorityScrollData.ScrollPosition.Y
						} else {
							highestPriorityScrollData.ScrollPosition.Y = 0
						}
					} else {
						highestPriorityScrollData.ScrollPosition.Y = -(highestPriorityScrollData.ContentSize.Height - highestPriorityScrollData.BoundingBox.Height)
					}
					scrollDeltaY = highestPriorityScrollData.ScrollPosition.Y - oldYScrollPosition
				}
				if scrollDeltaX > -0.1 && scrollDeltaX < 0.1 && scrollDeltaY > -0.1 && scrollDeltaY < 0.1 && highestPriorityScrollData.MomentumTime > 0.15 {
					highestPriorityScrollData.MomentumTime = 0
					highestPriorityScrollData.PointerOrigin = context.PointerInfo.Position
					highestPriorityScrollData.ScrollOrigin = highestPriorityScrollData.ScrollPosition
				} else {
					highestPriorityScrollData.MomentumTime += deltaTime
				}
			}
		}
		if canScrollVertically {
			if (func() float32 {
				if highestPriorityScrollData.ScrollPosition.Y < 0 {
					return highestPriorityScrollData.ScrollPosition.Y
				}
				return 0
			}()) > (-(highestPriorityScrollData.ContentSize.Height - scrollElement.Dimensions.Height)) {
				if highestPriorityScrollData.ScrollPosition.Y < 0 {
					highestPriorityScrollData.ScrollPosition.Y = highestPriorityScrollData.ScrollPosition.Y
				} else {
					highestPriorityScrollData.ScrollPosition.Y = 0
				}
			} else {
				highestPriorityScrollData.ScrollPosition.Y = -(highestPriorityScrollData.ContentSize.Height - scrollElement.Dimensions.Height)
			}
		}
		if canScrollHorizontally {
			if (func() float32 {
				if highestPriorityScrollData.ScrollPosition.X < 0 {
					return highestPriorityScrollData.ScrollPosition.X
				}
				return 0
			}()) > (-(highestPriorityScrollData.ContentSize.Width - scrollElement.Dimensions.Width)) {
				if highestPriorityScrollData.ScrollPosition.X < 0 {
					highestPriorityScrollData.ScrollPosition.X = highestPriorityScrollData.ScrollPosition.X
				} else {
					highestPriorityScrollData.ScrollPosition.X = 0
				}
			} else {
				highestPriorityScrollData.ScrollPosition.X = -(highestPriorityScrollData.ContentSize.Width - scrollElement.Dimensions.Width)
			}
		}
	}
}
func BeginLayout() {
	var context *Context = GetCurrentContext()
	__InitializeEphemeralMemory(context)
	context.Generation++
	context.DynamicElementIndex = 0
	var rootDimensions Dimensions = Dimensions{Width: context.LayoutDimensions.Width, Height: context.LayoutDimensions.Height}
	if context.DebugModeEnabled {
		rootDimensions.Width -= float32(__debugViewWidth)
	}
	context.BooleanWarnings = BooleanWarnings{MaxElementsExceeded: false}
	__OpenElement()
	__ConfigureOpenElement(ElementDeclaration{Id: __HashString(String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("__RootContainer")}, 0, 0), Layout: LayoutConfig{Sizing: Sizing{Width: SizingAxis{Size: struct {
		// union
		MinMax  SizingMinMax
		Percent float32
	}{MinMax: SizingMinMax{Min: rootDimensions.Width, Max: rootDimensions.Width}}, Type: __SIZING_TYPE_FIXED}, Height: SizingAxis{Size: struct {
		// union
		MinMax  SizingMinMax
		Percent float32
	}{MinMax: SizingMinMax{Min: rootDimensions.Height, Max: rootDimensions.Height}}, Type: __SIZING_TYPE_FIXED}}}})
	__int32_tArray_Add(&context.OpenLayoutElementStack, 0)
	__LayoutElementTreeRootArray_Add(&context.LayoutElementTreeRoots, __LayoutElementTreeRoot{})
}
func EndLayout() RenderCommandArray {
	var context *Context = GetCurrentContext()
	__CloseElement()
	var elementsExceededBeforeDebugView bool = context.BooleanWarnings.MaxElementsExceeded
	if context.BooleanWarnings.MaxElementsExceeded {
		var message String
		if !elementsExceededBeforeDebugView {
			message = String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay Error: Layout elements exceeded __maxElementCount after adding the debug-view to the layout.")}
		} else {
			message = String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay Error: Layout elements exceeded __maxElementCount")}
		}
		__AddRenderCommand(RenderCommand{BoundingBox: BoundingBox{X: context.LayoutDimensions.Width/2 - 59*4, Y: context.LayoutDimensions.Height / 2, Width: 0, Height: 0}, RenderData: RenderData{Text: TextRenderData{StringContents: StringSlice{Length: message.Length, Chars: message.Chars, BaseChars: message.Chars}, TextColor: Color{R: 255, G: 0, B: 0, A: 255}, FontSize: 16}}, CommandType: RENDER_COMMAND_TYPE_TEXT})
	} else {
		__CalculateFinalLayout()
	}
	return context.RenderCommands
}
func GetElementId(idString String) ElementId {
	return __HashString(idString, 0, 0)
}
func GetElementIdWithIndex(idString String, index uint32) ElementId {
	return __HashString(idString, index, 0)
}
func Hovered() bool {
	var context *Context = GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return false
	}
	var openLayoutElement *LayoutElement = __GetOpenLayoutElement()
	if openLayoutElement.Id == 0 {
		__GenerateIdForAnonymousElement(openLayoutElement)
	}
	for i := int32(0); i < context.PointerOverIds.Length; i++ {
		if __ElementIdArray_Get(&context.PointerOverIds, i).Id == openLayoutElement.Id {
			return true
		}
	}
	return false
}
func OnHover(onHoverFunction func(elementId ElementId, pointerInfo PointerData, userData int64), userData int64) {
	var context *Context = GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return
	}
	var openLayoutElement *LayoutElement = __GetOpenLayoutElement()
	if openLayoutElement.Id == 0 {
		__GenerateIdForAnonymousElement(openLayoutElement)
	}
	var hashMapItem *LayoutElementHashMapItem = __GetHashMapItem(openLayoutElement.Id)
	hashMapItem.OnHoverFunction = onHoverFunction
	hashMapItem.HoverFunctionUserData = userData
}
func PointerOver(elementId ElementId) bool {
	var context *Context = GetCurrentContext()
	for i := int32(0); i < context.PointerOverIds.Length; i++ {
		if __ElementIdArray_Get(&context.PointerOverIds, i).Id == elementId.Id {
			return true
		}
	}
	return false
}
func GetScrollContainerData(id ElementId) ScrollContainerData {
	var context *Context = GetCurrentContext()
	for i := int32(0); i < context.ScrollContainerDatas.Length; i++ {
		var scrollContainerData *__ScrollContainerDataInternal = __ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
		if scrollContainerData.ElementId == id.Id {
			return ScrollContainerData{ScrollPosition: &scrollContainerData.ScrollPosition, ScrollContainerDimensions: Dimensions{Width: scrollContainerData.BoundingBox.Width, Height: scrollContainerData.BoundingBox.Height}, ContentDimensions: scrollContainerData.ContentSize, Config: *__FindElementConfigWithType(scrollContainerData.LayoutElement, __ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig, Found: true}
		}
	}
	return ScrollContainerData{}
}
func GetElementData(id ElementId) ElementData {
	var item *LayoutElementHashMapItem = __GetHashMapItem(id.Id)
	if item == &LayoutElementHashMapItem_DEFAULT {
		return ElementData{}
	}
	return ElementData{BoundingBox: item.BoundingBox, Found: true}
}
func SetDebugModeEnabled(enabled bool) {
	var context *Context = GetCurrentContext()
	context.DebugModeEnabled = enabled
}
func IsDebugModeEnabled() bool {
	var context *Context = GetCurrentContext()
	return context.DebugModeEnabled
}
func SetCullingEnabled(enabled bool) {
	var context *Context = GetCurrentContext()
	context.DisableCulling = !enabled
}
func SetExternalScrollHandlingEnabled(enabled bool) {
	var context *Context = GetCurrentContext()
	context.ExternalScrollHandlingEnabled = enabled
}
func GetMaxElementCount() int32 {
	var context *Context = GetCurrentContext()
	return context.MaxElementCount
}
func SetMaxElementCount(maxElementCount int32) {
	var context *Context = GetCurrentContext()
	if context != nil {
		context.MaxElementCount = maxElementCount
	} else {
		__defaultMaxElementCount = maxElementCount
		__defaultMaxMeasureTextWordCacheCount = maxElementCount * 2
	}
}
func GetMaxMeasureTextCacheWordCount() int32 {
	var context *Context = GetCurrentContext()
	return context.MaxMeasureTextCacheWordCount
}
func SetMaxMeasureTextCacheWordCount(maxMeasureTextCacheWordCount int32) {
	var context *Context = GetCurrentContext()
	if context != nil {
		__currentContext.MaxMeasureTextCacheWordCount = maxMeasureTextCacheWordCount
	} else {
		__defaultMaxMeasureTextWordCacheCount = maxMeasureTextCacheWordCount
	}
}
func ResetMeasureTextCache() {
	var context *Context = GetCurrentContext()
	context.MeasureTextHashMapInternal.Length = 0
	context.MeasureTextHashMapInternalFreeList.Length = 0
	context.MeasureTextHashMap.Length = 0
	context.MeasuredWords.Length = 0
	context.MeasuredWordsFreeList.Length = 0
	for i := int32(0); i < context.MeasureTextHashMap.Capacity; i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = 0
	}
	context.MeasureTextHashMapInternal.Length = 1
}
