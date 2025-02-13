package clay

import (
	"math"
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/libc"
)

const CLAY__NULL = 0
const CLAY__MAXFLOAT = 3.4028234663852886e+38

var CLAY__ELEMENT_DEFINITION_LATCH uint8

type Clay_String struct {
	Length int32
	Chars  *byte
}
type Clay_StringSlice struct {
	Length    int32
	Chars     *byte
	BaseChars *byte
}
type Clay_Context struct {
	MaxElementCount                    int32
	MaxMeasureTextCacheWordCount       int32
	WarningsEnabled                    bool
	ErrorHandler                       Clay_ErrorHandler
	BooleanWarnings                    Clay_BooleanWarnings
	Warnings                           Clay__WarningArray
	PointerInfo                        Clay_PointerData
	LayoutDimensions                   Clay_Dimensions
	DynamicElementIndexBaseHash        Clay_ElementId
	DynamicElementIndex                uint32
	DebugModeEnabled                   bool
	DisableCulling                     bool
	ExternalScrollHandlingEnabled      bool
	DebugSelectedElementId             uint32
	Generation                         uint32
	ArenaResetOffset                   uint64
	MeasureTextUserData                unsafe.Pointer
	QueryScrollOffsetUserData          unsafe.Pointer
	InternalArena                      Clay_Arena
	LayoutElements                     Clay_LayoutElementArray
	RenderCommands                     Clay_RenderCommandArray
	OpenLayoutElementStack             Clay__int32_tArray
	LayoutElementChildren              Clay__int32_tArray
	LayoutElementChildrenBuffer        Clay__int32_tArray
	TextElementData                    Clay__TextElementDataArray
	ImageElementPointers               Clay__int32_tArray
	ReusableElementIndexBuffer         Clay__int32_tArray
	LayoutElementClipElementIds        Clay__int32_tArray
	LayoutConfigs                      Clay__LayoutConfigArray
	ElementConfigs                     Clay__ElementConfigArray
	TextElementConfigs                 Clay__TextElementConfigArray
	ImageElementConfigs                Clay__ImageElementConfigArray
	FloatingElementConfigs             Clay__FloatingElementConfigArray
	ScrollElementConfigs               Clay__ScrollElementConfigArray
	CustomElementConfigs               Clay__CustomElementConfigArray
	BorderElementConfigs               Clay__BorderElementConfigArray
	SharedElementConfigs               Clay__SharedElementConfigArray
	LayoutElementIdStrings             Clay__StringArray
	WrappedTextLines                   Clay__WrappedTextLineArray
	LayoutElementTreeNodeArray1        Clay__LayoutElementTreeNodeArray
	LayoutElementTreeRoots             Clay__LayoutElementTreeRootArray
	LayoutElementsHashMapInternal      Clay__LayoutElementHashMapItemArray
	LayoutElementsHashMap              Clay__int32_tArray
	MeasureTextHashMapInternal         Clay__MeasureTextCacheItemArray
	MeasureTextHashMapInternalFreeList Clay__int32_tArray
	MeasureTextHashMap                 Clay__int32_tArray
	MeasuredWords                      Clay__MeasuredWordArray
	MeasuredWordsFreeList              Clay__int32_tArray
	OpenClipElementStack               Clay__int32_tArray
	PointerOverIds                     Clay__ElementIdArray
	ScrollContainerDatas               Clay__ScrollContainerDataInternalArray
	TreeNodeVisited                    Clay__boolArray
	DynamicStringData                  Clay__charArray
	DebugElementData                   Clay__DebugElementDataArray
}
type Clay_Arena struct {
	NextAllocation uint64
	Capacity       uint64
	Memory         *byte
}
type Clay_Dimensions struct {
	Width  float32
	Height float32
}
type Clay_Vector2 struct {
	X float32
	Y float32
}
type Clay_Color struct {
	R float32
	G float32
	B float32
	A float32
}
type Clay_BoundingBox struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}
type Clay_ElementId struct {
	Id       uint32
	Offset   uint32
	BaseId   uint32
	StringId Clay_String
}
type Clay_CornerRadius struct {
	TopLeft     float32
	TopRight    float32
	BottomLeft  float32
	BottomRight float32
}
type Clay_LayoutDirection int64

const (
	CLAY_LEFT_TO_RIGHT = Clay_LayoutDirection(iota)
	CLAY_TOP_TO_BOTTOM
)

type Clay_LayoutAlignmentX int64

const (
	CLAY_ALIGN_X_LEFT = Clay_LayoutAlignmentX(iota)
	CLAY_ALIGN_X_RIGHT
	CLAY_ALIGN_X_CENTER
)

type Clay_LayoutAlignmentY int64

const (
	CLAY_ALIGN_Y_TOP = Clay_LayoutAlignmentY(iota)
	CLAY_ALIGN_Y_BOTTOM
	CLAY_ALIGN_Y_CENTER
)

type Clay__SizingType int64

const (
	CLAY__SIZING_TYPE_FIT = Clay__SizingType(iota)
	CLAY__SIZING_TYPE_GROW
	CLAY__SIZING_TYPE_PERCENT
	CLAY__SIZING_TYPE_FIXED
)

type Clay_ChildAlignment struct {
	X Clay_LayoutAlignmentX
	Y Clay_LayoutAlignmentY
}
type Clay_SizingMinMax struct {
	Min float32
	Max float32
}
type Clay_SizingAxis struct {
	Size struct {
		// union
		MinMax  Clay_SizingMinMax
		Percent float32
	}
	Type Clay__SizingType
}
type Clay_Sizing struct {
	Width  Clay_SizingAxis
	Height Clay_SizingAxis
}
type Clay_Padding struct {
	Left   uint16
	Right  uint16
	Top    uint16
	Bottom uint16
}
type Clay__Clay_PaddingWrapper struct {
	Wrapped Clay_Padding
}
type Clay_LayoutConfig struct {
	Sizing          Clay_Sizing
	Padding         Clay_Padding
	ChildGap        uint16
	ChildAlignment  Clay_ChildAlignment
	LayoutDirection Clay_LayoutDirection
}
type Clay__Clay_LayoutConfigWrapper struct {
	Wrapped Clay_LayoutConfig
}
type Clay_TextElementConfigWrapMode int64

const (
	CLAY_TEXT_WRAP_WORDS = Clay_TextElementConfigWrapMode(iota)
	CLAY_TEXT_WRAP_NEWLINES
	CLAY_TEXT_WRAP_NONE
)

type Clay_TextAlignment int64

const (
	CLAY_TEXT_ALIGN_LEFT = Clay_TextAlignment(iota)
	CLAY_TEXT_ALIGN_CENTER
	CLAY_TEXT_ALIGN_RIGHT
)

type Clay_TextElementConfig struct {
	TextColor          Clay_Color
	FontId             uint16
	FontSize           uint16
	LetterSpacing      uint16
	LineHeight         uint16
	WrapMode           Clay_TextElementConfigWrapMode
	TextAlignment      Clay_TextAlignment
	HashStringContents bool
}
type Clay__Clay_TextElementConfigWrapper struct {
	Wrapped Clay_TextElementConfig
}
type Clay_ImageElementConfig struct {
	ImageData        unsafe.Pointer
	SourceDimensions Clay_Dimensions
}
type Clay__Clay_ImageElementConfigWrapper struct {
	Wrapped Clay_ImageElementConfig
}
type Clay_FloatingAttachPointType int64

const (
	CLAY_ATTACH_POINT_LEFT_TOP = Clay_FloatingAttachPointType(iota)
	CLAY_ATTACH_POINT_LEFT_CENTER
	CLAY_ATTACH_POINT_LEFT_BOTTOM
	CLAY_ATTACH_POINT_CENTER_TOP
	CLAY_ATTACH_POINT_CENTER_CENTER
	CLAY_ATTACH_POINT_CENTER_BOTTOM
	CLAY_ATTACH_POINT_RIGHT_TOP
	CLAY_ATTACH_POINT_RIGHT_CENTER
	CLAY_ATTACH_POINT_RIGHT_BOTTOM
)

type Clay_FloatingAttachPoints struct {
	Element Clay_FloatingAttachPointType
	Parent  Clay_FloatingAttachPointType
}
type Clay_PointerCaptureMode int64

const (
	CLAY_POINTER_CAPTURE_MODE_CAPTURE = Clay_PointerCaptureMode(iota)
	CLAY_POINTER_CAPTURE_MODE_PASSTHROUGH
)

type Clay_FloatingAttachToElement int64

const (
	CLAY_ATTACH_TO_NONE = Clay_FloatingAttachToElement(iota)
	CLAY_ATTACH_TO_PARENT
	CLAY_ATTACH_TO_ELEMENT_WITH_ID
	CLAY_ATTACH_TO_ROOT
)

type Clay_FloatingElementConfig struct {
	Offset             Clay_Vector2
	Expand             Clay_Dimensions
	ParentId           uint32
	ZIndex             int16
	AttachPoints       Clay_FloatingAttachPoints
	PointerCaptureMode Clay_PointerCaptureMode
	AttachTo           Clay_FloatingAttachToElement
}
type Clay__Clay_FloatingElementConfigWrapper struct {
	Wrapped Clay_FloatingElementConfig
}
type Clay_CustomElementConfig struct {
	CustomData unsafe.Pointer
}
type Clay__Clay_CustomElementConfigWrapper struct {
	Wrapped Clay_CustomElementConfig
}
type Clay_ScrollElementConfig struct {
	Horizontal bool
	Vertical   bool
}
type Clay__Clay_ScrollElementConfigWrapper struct {
	Wrapped Clay_ScrollElementConfig
}
type Clay_BorderWidth struct {
	Left            uint16
	Right           uint16
	Top             uint16
	Bottom          uint16
	BetweenChildren uint16
}
type Clay_BorderElementConfig struct {
	Color Clay_Color
	Width Clay_BorderWidth
}
type Clay__Clay_BorderElementConfigWrapper struct {
	Wrapped Clay_BorderElementConfig
}
type Clay_TextRenderData struct {
	StringContents Clay_StringSlice
	TextColor      Clay_Color
	FontId         uint16
	FontSize       uint16
	LetterSpacing  uint16
	LineHeight     uint16
}
type Clay_RectangleRenderData struct {
	BackgroundColor Clay_Color
	CornerRadius    Clay_CornerRadius
}
type Clay_ImageRenderData struct {
	BackgroundColor  Clay_Color
	CornerRadius     Clay_CornerRadius
	SourceDimensions Clay_Dimensions
	ImageData        unsafe.Pointer
}
type Clay_CustomRenderData struct {
	BackgroundColor Clay_Color
	CornerRadius    Clay_CornerRadius
	CustomData      unsafe.Pointer
}
type Clay_ScrollRenderData struct {
	Horizontal bool
	Vertical   bool
}
type Clay_BorderRenderData struct {
	Color        Clay_Color
	CornerRadius Clay_CornerRadius
	Width        Clay_BorderWidth
}
type Clay_RenderData struct {
	// union
	Rectangle Clay_RectangleRenderData
	Text      Clay_TextRenderData
	Image     Clay_ImageRenderData
	Custom    Clay_CustomRenderData
	Border    Clay_BorderRenderData
	Scroll    Clay_ScrollRenderData
}
type Clay_ScrollContainerData struct {
	ScrollPosition            *Clay_Vector2
	ScrollContainerDimensions Clay_Dimensions
	ContentDimensions         Clay_Dimensions
	Config                    Clay_ScrollElementConfig
	Found                     bool
}
type Clay_ElementData struct {
	BoundingBox Clay_BoundingBox
	Found       bool
}
type Clay_RenderCommandType int64

const (
	CLAY_RENDER_COMMAND_TYPE_NONE = Clay_RenderCommandType(iota)
	CLAY_RENDER_COMMAND_TYPE_RECTANGLE
	CLAY_RENDER_COMMAND_TYPE_BORDER
	CLAY_RENDER_COMMAND_TYPE_TEXT
	CLAY_RENDER_COMMAND_TYPE_IMAGE
	CLAY_RENDER_COMMAND_TYPE_SCISSOR_START
	CLAY_RENDER_COMMAND_TYPE_SCISSOR_END
	CLAY_RENDER_COMMAND_TYPE_CUSTOM
)

type Clay_RenderCommand struct {
	BoundingBox Clay_BoundingBox
	RenderData  Clay_RenderData
	UserData    unsafe.Pointer
	Id          uint32
	ZIndex      int16
	CommandType Clay_RenderCommandType
}
type Clay_RenderCommandArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_RenderCommand
}
type Clay_PointerDataInteractionState int64

const (
	CLAY_POINTER_DATA_PRESSED_THIS_FRAME = Clay_PointerDataInteractionState(iota)
	CLAY_POINTER_DATA_PRESSED
	CLAY_POINTER_DATA_RELEASED_THIS_FRAME
	CLAY_POINTER_DATA_RELEASED
)

type Clay_PointerData struct {
	Position Clay_Vector2
	State    Clay_PointerDataInteractionState
}
type Clay_ElementDeclaration struct {
	Id              Clay_ElementId
	Layout          Clay_LayoutConfig
	BackgroundColor Clay_Color
	CornerRadius    Clay_CornerRadius
	Image           Clay_ImageElementConfig
	Floating        Clay_FloatingElementConfig
	Custom          Clay_CustomElementConfig
	Scroll          Clay_ScrollElementConfig
	Border          Clay_BorderElementConfig
	UserData        unsafe.Pointer
}
type Clay__Clay_ElementDeclarationWrapper struct {
	Wrapped Clay_ElementDeclaration
}
type Clay_ErrorType int64

const (
	CLAY_ERROR_TYPE_TEXT_MEASUREMENT_FUNCTION_NOT_PROVIDED = Clay_ErrorType(iota)
	CLAY_ERROR_TYPE_ARENA_CAPACITY_EXCEEDED
	CLAY_ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED
	CLAY_ERROR_TYPE_TEXT_MEASUREMENT_CAPACITY_EXCEEDED
	CLAY_ERROR_TYPE_DUPLICATE_ID
	CLAY_ERROR_TYPE_FLOATING_CONTAINER_PARENT_NOT_FOUND
	CLAY_ERROR_TYPE_PERCENTAGE_OVER_1
	CLAY_ERROR_TYPE_INTERNAL_ERROR
)

type Clay_ErrorData struct {
	ErrorType Clay_ErrorType
	ErrorText Clay_String
	UserData  unsafe.Pointer
}
type Clay_ErrorHandler struct {
	ErrorHandlerFunction func(errorText Clay_ErrorData)
	UserData             unsafe.Pointer
}

var CLAY_LAYOUT_DEFAULT Clay_LayoutConfig = Clay_LayoutConfig{}
var Clay__Color_DEFAULT Clay_Color = Clay_Color{}
var Clay__CornerRadius_DEFAULT Clay_CornerRadius = Clay_CornerRadius{}
var Clay__BorderWidth_DEFAULT Clay_BorderWidth = Clay_BorderWidth{}
var Clay__currentContext *Clay_Context
var Clay__defaultMaxElementCount int32 = 8192
var Clay__defaultMaxMeasureTextWordCacheCount int32 = 16384

func Clay__ErrorHandlerFunctionDefault(errorText Clay_ErrorData) {
	_ = errorText
}

var CLAY__SPACECHAR Clay_String = Clay_String{Length: 1, Chars: libc.CString(" ")}
var CLAY__STRING_DEFAULT Clay_String = Clay_String{}

type Clay_BooleanWarnings struct {
	MaxElementsExceeded           bool
	MaxRenderCommandsExceeded     bool
	MaxTextMeasureCacheExceeded   bool
	TextMeasurementFunctionNotSet bool
}
type Clay__Warning struct {
	BaseMessage    Clay_String
	DynamicMessage Clay_String
}

var CLAY__WARNING_DEFAULT Clay__Warning = Clay__Warning{}

type Clay__WarningArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__Warning
}
type Clay_SharedElementConfig struct {
	BackgroundColor Clay_Color
	CornerRadius    Clay_CornerRadius
	UserData        unsafe.Pointer
}
type Clay__Clay_SharedElementConfigWrapper struct {
	Wrapped Clay_SharedElementConfig
}
type Clay__boolArray struct {
	Capacity      int32
	Length        int32
	InternalArray *bool
}
type Clay__boolArraySlice struct {
	Length        int32
	InternalArray *bool
}

var _Bool_DEFAULT bool = false

func Clay__boolArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__boolArray {
	return Clay__boolArray{Capacity: capacity, Length: 0, InternalArray: (*bool)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(bool(false))), arena))}
}
func Clay__boolArray_Get(array *Clay__boolArray, index int32) *bool {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
	}
	return &_Bool_DEFAULT
}
func Clay__boolArray_GetValue(array *Clay__boolArray, index int32) bool {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
	}
	return _Bool_DEFAULT
}
func Clay__boolArray_Add(array *Clay__boolArray, item bool) *bool {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}())) = item
		return (*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), int64(array.Length)-1))
	}
	return &_Bool_DEFAULT
}
func Clay__boolArraySlice_Get(slice *Clay__boolArraySlice, index int32) *bool {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*bool)(unsafe.Add(unsafe.Pointer(slice.InternalArray), index))
	}
	return &_Bool_DEFAULT
}
func Clay__boolArray_RemoveSwapback(array *Clay__boolArray, index int32) bool {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed bool = *(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
		*(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = *(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), array.Length))
		return removed
	}
	return _Bool_DEFAULT
}
func Clay__boolArray_Set(array *Clay__boolArray, index int32, value bool) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*bool)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__int32_tArray struct {
	Capacity      int32
	Length        int32
	InternalArray *int32
}
type Clay__int32_tArraySlice struct {
	Length        int32
	InternalArray *int32
}

var int32_t_DEFAULT int32 = 0

func Clay__int32_tArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__int32_tArray {
	return Clay__int32_tArray{Capacity: capacity, Length: 0, InternalArray: (*int32)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(int32(0))), arena))}
}
func Clay__int32_tArray_Get(array *Clay__int32_tArray, index int32) *int32 {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
	}
	return &int32_t_DEFAULT
}
func Clay__int32_tArray_GetValue(array *Clay__int32_tArray, index int32) int32 {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
	}
	return int32_t_DEFAULT
}
func Clay__int32_tArray_Add(array *Clay__int32_tArray, item int32) *int32 {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(int64(array.Length)-1)))
	}
	return &int32_t_DEFAULT
}
func Clay__int32_tArraySlice_Get(slice *Clay__int32_tArraySlice, index int32) *int32 {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*int32)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
	}
	return &int32_t_DEFAULT
}
func Clay__int32_tArray_RemoveSwapback(array *Clay__int32_tArray, index int32) int32 {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed int32 = *(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index)))
		*(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index))) = *(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(array.Length)))
		return removed
	}
	return int32_t_DEFAULT
}
func Clay__int32_tArray_Set(array *Clay__int32_tArray, index int32, value int32) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*int32)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(int32(0))*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__charArray struct {
	Capacity      int32
	Length        int32
	InternalArray *byte
}
type Clay__charArraySlice struct {
	Length        int32
	InternalArray *byte
}

var char_DEFAULT int8 = 0

func Clay__charArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__charArray {
	return Clay__charArray{Capacity: capacity, Length: 0, InternalArray: (*byte)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(int8(0))), arena))}
}
func Clay__charArray_Get(array *Clay__charArray, index int32) *byte {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index))
	}
	return (*byte)(unsafe.Pointer(&char_DEFAULT))
}
func Clay__charArray_GetValue(array *Clay__charArray, index int32) int8 {
	if Clay__Array_RangeCheck(index, array.Length) {
		return int8(*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)))
	}
	return char_DEFAULT
}
func Clay__charArray_Add(array *Clay__charArray, item int8) *byte {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}())) = byte(item)
		return (*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), int64(array.Length)-1))
	}
	return (*byte)(unsafe.Pointer(&char_DEFAULT))
}
func Clay__charArraySlice_Get(slice *Clay__charArraySlice, index int32) *byte {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*byte)(unsafe.Add(unsafe.Pointer(slice.InternalArray), index))
	}
	return (*byte)(unsafe.Pointer(&char_DEFAULT))
}
func Clay__charArray_RemoveSwapback(array *Clay__charArray, index int32) int8 {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)))
		*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = *(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), array.Length))
		return removed
	}
	return char_DEFAULT
}
func Clay__charArray_Set(array *Clay__charArray, index int32, value int8) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*byte)(unsafe.Add(unsafe.Pointer(array.InternalArray), index)) = byte(value)
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__ElementIdArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_ElementId
}
type Clay__ElementIdArraySlice struct {
	Length        int32
	InternalArray *Clay_ElementId
}

var Clay_ElementId_DEFAULT Clay_ElementId = Clay_ElementId{}

func Clay__ElementIdArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__ElementIdArray {
	return Clay__ElementIdArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_ElementId)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_ElementId{})), arena))}
}
func Clay__ElementIdArray_Get(array *Clay__ElementIdArray, index int32) *Clay_ElementId {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(index)))
	}
	return &Clay_ElementId_DEFAULT
}
func Clay__ElementIdArray_GetValue(array *Clay__ElementIdArray, index int32) Clay_ElementId {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(index)))
	}
	return Clay_ElementId_DEFAULT
}
func Clay__ElementIdArray_Add(array *Clay__ElementIdArray, item Clay_ElementId) *Clay_ElementId {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_ElementId_DEFAULT
}
func Clay__ElementIdArraySlice_Get(slice *Clay__ElementIdArraySlice, index int32) *Clay_ElementId {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_ElementId)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(index)))
	}
	return &Clay_ElementId_DEFAULT
}
func Clay__ElementIdArray_RemoveSwapback(array *Clay__ElementIdArray, index int32) Clay_ElementId {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_ElementId = *(*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(index)))
		*(*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(index))) = *(*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(array.Length)))
		return removed
	}
	return Clay_ElementId_DEFAULT
}
func Clay__ElementIdArray_Set(array *Clay__ElementIdArray, index int32, value Clay_ElementId) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_ElementId)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementId{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__LayoutConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_LayoutConfig
}
type Clay__LayoutConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_LayoutConfig
}

var Clay_LayoutConfig_DEFAULT Clay_LayoutConfig = Clay_LayoutConfig{}

func Clay__LayoutConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__LayoutConfigArray {
	return Clay__LayoutConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_LayoutConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_LayoutConfig{})), arena))}
}
func Clay__LayoutConfigArray_Get(array *Clay__LayoutConfigArray, index int32) *Clay_LayoutConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(index)))
	}
	return &Clay_LayoutConfig_DEFAULT
}
func Clay__LayoutConfigArray_GetValue(array *Clay__LayoutConfigArray, index int32) Clay_LayoutConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(index)))
	}
	return Clay_LayoutConfig_DEFAULT
}
func Clay__LayoutConfigArray_Add(array *Clay__LayoutConfigArray, item Clay_LayoutConfig) *Clay_LayoutConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_LayoutConfig_DEFAULT
}
func Clay__LayoutConfigArraySlice_Get(slice *Clay__LayoutConfigArraySlice, index int32) *Clay_LayoutConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(index)))
	}
	return &Clay_LayoutConfig_DEFAULT
}
func Clay__LayoutConfigArray_RemoveSwapback(array *Clay__LayoutConfigArray, index int32) Clay_LayoutConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_LayoutConfig = *(*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(index)))
		*(*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(index))) = *(*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_LayoutConfig_DEFAULT
}
func Clay__LayoutConfigArray_Set(array *Clay__LayoutConfigArray, index int32, value Clay_LayoutConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_LayoutConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__TextElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_TextElementConfig
}
type Clay__TextElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_TextElementConfig
}

var Clay_TextElementConfig_DEFAULT Clay_TextElementConfig = Clay_TextElementConfig{}

func Clay__TextElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__TextElementConfigArray {
	return Clay__TextElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_TextElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_TextElementConfig{})), arena))}
}
func Clay__TextElementConfigArray_Get(array *Clay__TextElementConfigArray, index int32) *Clay_TextElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(index)))
	}
	return &Clay_TextElementConfig_DEFAULT
}
func Clay__TextElementConfigArray_GetValue(array *Clay__TextElementConfigArray, index int32) Clay_TextElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(index)))
	}
	return Clay_TextElementConfig_DEFAULT
}
func Clay__TextElementConfigArray_Add(array *Clay__TextElementConfigArray, item Clay_TextElementConfig) *Clay_TextElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_TextElementConfig_DEFAULT
}
func Clay__TextElementConfigArraySlice_Get(slice *Clay__TextElementConfigArraySlice, index int32) *Clay_TextElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(index)))
	}
	return &Clay_TextElementConfig_DEFAULT
}
func Clay__TextElementConfigArray_RemoveSwapback(array *Clay__TextElementConfigArray, index int32) Clay_TextElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_TextElementConfig = *(*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(index)))
		*(*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(index))) = *(*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_TextElementConfig_DEFAULT
}
func Clay__TextElementConfigArray_Set(array *Clay__TextElementConfigArray, index int32, value Clay_TextElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_TextElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_TextElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__ImageElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_ImageElementConfig
}
type Clay__ImageElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_ImageElementConfig
}

var Clay_ImageElementConfig_DEFAULT Clay_ImageElementConfig = Clay_ImageElementConfig{}

func Clay__ImageElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__ImageElementConfigArray {
	return Clay__ImageElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_ImageElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_ImageElementConfig{})), arena))}
}
func Clay__ImageElementConfigArray_Get(array *Clay__ImageElementConfigArray, index int32) *Clay_ImageElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(index)))
	}
	return &Clay_ImageElementConfig_DEFAULT
}
func Clay__ImageElementConfigArray_GetValue(array *Clay__ImageElementConfigArray, index int32) Clay_ImageElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(index)))
	}
	return Clay_ImageElementConfig_DEFAULT
}
func Clay__ImageElementConfigArray_Add(array *Clay__ImageElementConfigArray, item Clay_ImageElementConfig) *Clay_ImageElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_ImageElementConfig_DEFAULT
}
func Clay__ImageElementConfigArraySlice_Get(slice *Clay__ImageElementConfigArraySlice, index int32) *Clay_ImageElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(index)))
	}
	return &Clay_ImageElementConfig_DEFAULT
}
func Clay__ImageElementConfigArray_RemoveSwapback(array *Clay__ImageElementConfigArray, index int32) Clay_ImageElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_ImageElementConfig = *(*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(index)))
		*(*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(index))) = *(*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_ImageElementConfig_DEFAULT
}
func Clay__ImageElementConfigArray_Set(array *Clay__ImageElementConfigArray, index int32, value Clay_ImageElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_ImageElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ImageElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__FloatingElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_FloatingElementConfig
}
type Clay__FloatingElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_FloatingElementConfig
}

var Clay_FloatingElementConfig_DEFAULT Clay_FloatingElementConfig = Clay_FloatingElementConfig{}

func Clay__FloatingElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__FloatingElementConfigArray {
	return Clay__FloatingElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_FloatingElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_FloatingElementConfig{})), arena))}
}
func Clay__FloatingElementConfigArray_Get(array *Clay__FloatingElementConfigArray, index int32) *Clay_FloatingElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(index)))
	}
	return &Clay_FloatingElementConfig_DEFAULT
}
func Clay__FloatingElementConfigArray_GetValue(array *Clay__FloatingElementConfigArray, index int32) Clay_FloatingElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(index)))
	}
	return Clay_FloatingElementConfig_DEFAULT
}
func Clay__FloatingElementConfigArray_Add(array *Clay__FloatingElementConfigArray, item Clay_FloatingElementConfig) *Clay_FloatingElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_FloatingElementConfig_DEFAULT
}
func Clay__FloatingElementConfigArraySlice_Get(slice *Clay__FloatingElementConfigArraySlice, index int32) *Clay_FloatingElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(index)))
	}
	return &Clay_FloatingElementConfig_DEFAULT
}
func Clay__FloatingElementConfigArray_RemoveSwapback(array *Clay__FloatingElementConfigArray, index int32) Clay_FloatingElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_FloatingElementConfig = *(*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(index)))
		*(*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(index))) = *(*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_FloatingElementConfig_DEFAULT
}
func Clay__FloatingElementConfigArray_Set(array *Clay__FloatingElementConfigArray, index int32, value Clay_FloatingElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_FloatingElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_FloatingElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__CustomElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_CustomElementConfig
}
type Clay__CustomElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_CustomElementConfig
}

var Clay_CustomElementConfig_DEFAULT Clay_CustomElementConfig = Clay_CustomElementConfig{}

func Clay__CustomElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__CustomElementConfigArray {
	return Clay__CustomElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_CustomElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_CustomElementConfig{})), arena))}
}
func Clay__CustomElementConfigArray_Get(array *Clay__CustomElementConfigArray, index int32) *Clay_CustomElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(index)))
	}
	return &Clay_CustomElementConfig_DEFAULT
}
func Clay__CustomElementConfigArray_GetValue(array *Clay__CustomElementConfigArray, index int32) Clay_CustomElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(index)))
	}
	return Clay_CustomElementConfig_DEFAULT
}
func Clay__CustomElementConfigArray_Add(array *Clay__CustomElementConfigArray, item Clay_CustomElementConfig) *Clay_CustomElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_CustomElementConfig_DEFAULT
}
func Clay__CustomElementConfigArraySlice_Get(slice *Clay__CustomElementConfigArraySlice, index int32) *Clay_CustomElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(index)))
	}
	return &Clay_CustomElementConfig_DEFAULT
}
func Clay__CustomElementConfigArray_RemoveSwapback(array *Clay__CustomElementConfigArray, index int32) Clay_CustomElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_CustomElementConfig = *(*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(index)))
		*(*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(index))) = *(*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_CustomElementConfig_DEFAULT
}
func Clay__CustomElementConfigArray_Set(array *Clay__CustomElementConfigArray, index int32, value Clay_CustomElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_CustomElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_CustomElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__ScrollElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_ScrollElementConfig
}
type Clay__ScrollElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_ScrollElementConfig
}

var Clay_ScrollElementConfig_DEFAULT Clay_ScrollElementConfig = Clay_ScrollElementConfig{Horizontal: false}

func Clay__ScrollElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__ScrollElementConfigArray {
	return Clay__ScrollElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_ScrollElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_ScrollElementConfig{})), arena))}
}
func Clay__ScrollElementConfigArray_Get(array *Clay__ScrollElementConfigArray, index int32) *Clay_ScrollElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(index)))
	}
	return &Clay_ScrollElementConfig_DEFAULT
}
func Clay__ScrollElementConfigArray_GetValue(array *Clay__ScrollElementConfigArray, index int32) Clay_ScrollElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(index)))
	}
	return Clay_ScrollElementConfig_DEFAULT
}
func Clay__ScrollElementConfigArray_Add(array *Clay__ScrollElementConfigArray, item Clay_ScrollElementConfig) *Clay_ScrollElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_ScrollElementConfig_DEFAULT
}
func Clay__ScrollElementConfigArraySlice_Get(slice *Clay__ScrollElementConfigArraySlice, index int32) *Clay_ScrollElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(index)))
	}
	return &Clay_ScrollElementConfig_DEFAULT
}
func Clay__ScrollElementConfigArray_RemoveSwapback(array *Clay__ScrollElementConfigArray, index int32) Clay_ScrollElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_ScrollElementConfig = *(*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(index)))
		*(*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(index))) = *(*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_ScrollElementConfig_DEFAULT
}
func Clay__ScrollElementConfigArray_Set(array *Clay__ScrollElementConfigArray, index int32, value Clay_ScrollElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_ScrollElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ScrollElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__BorderElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_BorderElementConfig
}
type Clay__BorderElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_BorderElementConfig
}

var Clay_BorderElementConfig_DEFAULT Clay_BorderElementConfig = Clay_BorderElementConfig{}

func Clay__BorderElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__BorderElementConfigArray {
	return Clay__BorderElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_BorderElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_BorderElementConfig{})), arena))}
}
func Clay__BorderElementConfigArray_Get(array *Clay__BorderElementConfigArray, index int32) *Clay_BorderElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(index)))
	}
	return &Clay_BorderElementConfig_DEFAULT
}
func Clay__BorderElementConfigArray_GetValue(array *Clay__BorderElementConfigArray, index int32) Clay_BorderElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(index)))
	}
	return Clay_BorderElementConfig_DEFAULT
}
func Clay__BorderElementConfigArray_Add(array *Clay__BorderElementConfigArray, item Clay_BorderElementConfig) *Clay_BorderElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_BorderElementConfig_DEFAULT
}
func Clay__BorderElementConfigArraySlice_Get(slice *Clay__BorderElementConfigArraySlice, index int32) *Clay_BorderElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(index)))
	}
	return &Clay_BorderElementConfig_DEFAULT
}
func Clay__BorderElementConfigArray_RemoveSwapback(array *Clay__BorderElementConfigArray, index int32) Clay_BorderElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_BorderElementConfig = *(*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(index)))
		*(*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(index))) = *(*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_BorderElementConfig_DEFAULT
}
func Clay__BorderElementConfigArray_Set(array *Clay__BorderElementConfigArray, index int32, value Clay_BorderElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_BorderElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_BorderElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__StringArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_String
}
type Clay__StringArraySlice struct {
	Length        int32
	InternalArray *Clay_String
}

var Clay_String_DEFAULT Clay_String = Clay_String{}

func Clay__StringArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__StringArray {
	return Clay__StringArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_String)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_String{})), arena))}
}
func Clay__StringArray_Get(array *Clay__StringArray, index int32) *Clay_String {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(index)))
	}
	return &Clay_String_DEFAULT
}
func Clay__StringArray_GetValue(array *Clay__StringArray, index int32) Clay_String {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(index)))
	}
	return Clay_String_DEFAULT
}
func Clay__StringArray_Add(array *Clay__StringArray, item Clay_String) *Clay_String {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_String_DEFAULT
}
func Clay__StringArraySlice_Get(slice *Clay__StringArraySlice, index int32) *Clay_String {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_String)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(index)))
	}
	return &Clay_String_DEFAULT
}
func Clay__StringArray_RemoveSwapback(array *Clay__StringArray, index int32) Clay_String {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_String = *(*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(index)))
		*(*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(index))) = *(*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(array.Length)))
		return removed
	}
	return Clay_String_DEFAULT
}
func Clay__StringArray_Set(array *Clay__StringArray, index int32, value Clay_String) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_String)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_String{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__SharedElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_SharedElementConfig
}
type Clay__SharedElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_SharedElementConfig
}

var Clay_SharedElementConfig_DEFAULT Clay_SharedElementConfig = Clay_SharedElementConfig{}

func Clay__SharedElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__SharedElementConfigArray {
	return Clay__SharedElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_SharedElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_SharedElementConfig{})), arena))}
}
func Clay__SharedElementConfigArray_Get(array *Clay__SharedElementConfigArray, index int32) *Clay_SharedElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(index)))
	}
	return &Clay_SharedElementConfig_DEFAULT
}
func Clay__SharedElementConfigArray_GetValue(array *Clay__SharedElementConfigArray, index int32) Clay_SharedElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(index)))
	}
	return Clay_SharedElementConfig_DEFAULT
}
func Clay__SharedElementConfigArray_Add(array *Clay__SharedElementConfigArray, item Clay_SharedElementConfig) *Clay_SharedElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_SharedElementConfig_DEFAULT
}
func Clay__SharedElementConfigArraySlice_Get(slice *Clay__SharedElementConfigArraySlice, index int32) *Clay_SharedElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(index)))
	}
	return &Clay_SharedElementConfig_DEFAULT
}
func Clay__SharedElementConfigArray_RemoveSwapback(array *Clay__SharedElementConfigArray, index int32) Clay_SharedElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_SharedElementConfig = *(*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(index)))
		*(*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(index))) = *(*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_SharedElementConfig_DEFAULT
}
func Clay__SharedElementConfigArray_Set(array *Clay__SharedElementConfigArray, index int32, value Clay_SharedElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_SharedElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_SharedElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay_RenderCommandArraySlice struct {
	Length        int32
	InternalArray *Clay_RenderCommand
}

var Clay_RenderCommand_DEFAULT Clay_RenderCommand = Clay_RenderCommand{}

func Clay_RenderCommandArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay_RenderCommandArray {
	return Clay_RenderCommandArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_RenderCommand)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_RenderCommand{})), arena))}
}
func Clay_RenderCommandArray_Get(array *Clay_RenderCommandArray, index int32) *Clay_RenderCommand {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(index)))
	}
	return &Clay_RenderCommand_DEFAULT
}
func Clay_RenderCommandArray_GetValue(array *Clay_RenderCommandArray, index int32) Clay_RenderCommand {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(index)))
	}
	return Clay_RenderCommand_DEFAULT
}
func Clay_RenderCommandArray_Add(array *Clay_RenderCommandArray, item Clay_RenderCommand) *Clay_RenderCommand {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_RenderCommand_DEFAULT
}
func Clay_RenderCommandArraySlice_Get(slice *Clay_RenderCommandArraySlice, index int32) *Clay_RenderCommand {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(index)))
	}
	return &Clay_RenderCommand_DEFAULT
}
func Clay_RenderCommandArray_RemoveSwapback(array *Clay_RenderCommandArray, index int32) Clay_RenderCommand {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_RenderCommand = *(*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(index)))
		*(*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(index))) = *(*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(array.Length)))
		return removed
	}
	return Clay_RenderCommand_DEFAULT
}
func Clay_RenderCommandArray_Set(array *Clay_RenderCommandArray, index int32, value Clay_RenderCommand) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_RenderCommand)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_RenderCommand{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__ElementConfigType int64

const (
	CLAY__ELEMENT_CONFIG_TYPE_NONE = Clay__ElementConfigType(iota)
	CLAY__ELEMENT_CONFIG_TYPE_BORDER
	CLAY__ELEMENT_CONFIG_TYPE_FLOATING
	CLAY__ELEMENT_CONFIG_TYPE_SCROLL
	CLAY__ELEMENT_CONFIG_TYPE_IMAGE
	CLAY__ELEMENT_CONFIG_TYPE_TEXT
	CLAY__ELEMENT_CONFIG_TYPE_CUSTOM
	CLAY__ELEMENT_CONFIG_TYPE_SHARED
)

type Clay_ElementConfigUnion struct {
	// union
	TextElementConfig     *Clay_TextElementConfig
	ImageElementConfig    *Clay_ImageElementConfig
	FloatingElementConfig *Clay_FloatingElementConfig
	CustomElementConfig   *Clay_CustomElementConfig
	ScrollElementConfig   *Clay_ScrollElementConfig
	BorderElementConfig   *Clay_BorderElementConfig
	SharedElementConfig   *Clay_SharedElementConfig
}
type Clay_ElementConfig struct {
	Type   Clay__ElementConfigType
	Config Clay_ElementConfigUnion
}
type Clay__ElementConfigArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_ElementConfig
}
type Clay__ElementConfigArraySlice struct {
	Length        int32
	InternalArray *Clay_ElementConfig
}

var Clay_ElementConfig_DEFAULT Clay_ElementConfig = Clay_ElementConfig{}

func Clay__ElementConfigArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__ElementConfigArray {
	return Clay__ElementConfigArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_ElementConfig)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_ElementConfig{})), arena))}
}
func Clay__ElementConfigArray_Get(array *Clay__ElementConfigArray, index int32) *Clay_ElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(index)))
	}
	return &Clay_ElementConfig_DEFAULT
}
func Clay__ElementConfigArray_GetValue(array *Clay__ElementConfigArray, index int32) Clay_ElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(index)))
	}
	return Clay_ElementConfig_DEFAULT
}
func Clay__ElementConfigArray_Add(array *Clay__ElementConfigArray, item Clay_ElementConfig) *Clay_ElementConfig {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_ElementConfig_DEFAULT
}
func Clay__ElementConfigArraySlice_Get(slice *Clay__ElementConfigArraySlice, index int32) *Clay_ElementConfig {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(index)))
	}
	return &Clay_ElementConfig_DEFAULT
}
func Clay__ElementConfigArray_RemoveSwapback(array *Clay__ElementConfigArray, index int32) Clay_ElementConfig {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_ElementConfig = *(*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(index)))
		*(*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(index))) = *(*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(array.Length)))
		return removed
	}
	return Clay_ElementConfig_DEFAULT
}
func Clay__ElementConfigArray_Set(array *Clay__ElementConfigArray, index int32, value Clay_ElementConfig) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__WrappedTextLine struct {
	Dimensions Clay_Dimensions
	Line       Clay_String
}
type Clay__WrappedTextLineArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__WrappedTextLine
}
type Clay__WrappedTextLineArraySlice struct {
	Length        int32
	InternalArray *Clay__WrappedTextLine
}

var Clay__WrappedTextLine_DEFAULT Clay__WrappedTextLine = Clay__WrappedTextLine{}

func Clay__WrappedTextLineArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__WrappedTextLineArray {
	return Clay__WrappedTextLineArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__WrappedTextLine)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__WrappedTextLine{})), arena))}
}
func Clay__WrappedTextLineArray_Get(array *Clay__WrappedTextLineArray, index int32) *Clay__WrappedTextLine {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(index)))
	}
	return &Clay__WrappedTextLine_DEFAULT
}
func Clay__WrappedTextLineArray_GetValue(array *Clay__WrappedTextLineArray, index int32) Clay__WrappedTextLine {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(index)))
	}
	return Clay__WrappedTextLine_DEFAULT
}
func Clay__WrappedTextLineArray_Add(array *Clay__WrappedTextLineArray, item Clay__WrappedTextLine) *Clay__WrappedTextLine {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__WrappedTextLine_DEFAULT
}
func Clay__WrappedTextLineArraySlice_Get(slice *Clay__WrappedTextLineArraySlice, index int32) *Clay__WrappedTextLine {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(index)))
	}
	return &Clay__WrappedTextLine_DEFAULT
}
func Clay__WrappedTextLineArray_RemoveSwapback(array *Clay__WrappedTextLineArray, index int32) Clay__WrappedTextLine {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__WrappedTextLine = *(*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(index)))
		*(*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(index))) = *(*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(array.Length)))
		return removed
	}
	return Clay__WrappedTextLine_DEFAULT
}
func Clay__WrappedTextLineArray_Set(array *Clay__WrappedTextLineArray, index int32, value Clay__WrappedTextLine) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__TextElementData struct {
	Text                Clay_String
	PreferredDimensions Clay_Dimensions
	ElementIndex        int32
	WrappedLines        Clay__WrappedTextLineArraySlice
}
type Clay__TextElementDataArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__TextElementData
}
type Clay__TextElementDataArraySlice struct {
	Length        int32
	InternalArray *Clay__TextElementData
}

var Clay__TextElementData_DEFAULT Clay__TextElementData = Clay__TextElementData{}

func Clay__TextElementDataArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__TextElementDataArray {
	return Clay__TextElementDataArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__TextElementData)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__TextElementData{})), arena))}
}
func Clay__TextElementDataArray_Get(array *Clay__TextElementDataArray, index int32) *Clay__TextElementData {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(index)))
	}
	return &Clay__TextElementData_DEFAULT
}
func Clay__TextElementDataArray_GetValue(array *Clay__TextElementDataArray, index int32) Clay__TextElementData {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(index)))
	}
	return Clay__TextElementData_DEFAULT
}
func Clay__TextElementDataArray_Add(array *Clay__TextElementDataArray, item Clay__TextElementData) *Clay__TextElementData {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__TextElementData_DEFAULT
}
func Clay__TextElementDataArraySlice_Get(slice *Clay__TextElementDataArraySlice, index int32) *Clay__TextElementData {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(index)))
	}
	return &Clay__TextElementData_DEFAULT
}
func Clay__TextElementDataArray_RemoveSwapback(array *Clay__TextElementDataArray, index int32) Clay__TextElementData {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__TextElementData = *(*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(index)))
		*(*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(index))) = *(*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(array.Length)))
		return removed
	}
	return Clay__TextElementData_DEFAULT
}
func Clay__TextElementDataArray_Set(array *Clay__TextElementDataArray, index int32, value Clay__TextElementData) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__TextElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__TextElementData{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__LayoutElementChildren struct {
	Elements *int32
	Length   uint16
}
type Clay_LayoutElement struct {
	ChildrenOrTextContent struct {
		// union
		Children        Clay__LayoutElementChildren
		TextElementData *Clay__TextElementData
	}
	Dimensions     Clay_Dimensions
	MinDimensions  Clay_Dimensions
	LayoutConfig   *Clay_LayoutConfig
	ElementConfigs Clay__ElementConfigArraySlice
	Id             uint32
}
type Clay_LayoutElementArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_LayoutElement
}
type Clay_LayoutElementArraySlice struct {
	Length        int32
	InternalArray *Clay_LayoutElement
}

var Clay_LayoutElement_DEFAULT Clay_LayoutElement = Clay_LayoutElement{}

func Clay_LayoutElementArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay_LayoutElementArray {
	return Clay_LayoutElementArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_LayoutElement)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_LayoutElement{})), arena))}
}
func Clay_LayoutElementArray_Get(array *Clay_LayoutElementArray, index int32) *Clay_LayoutElement {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(index)))
	}
	return &Clay_LayoutElement_DEFAULT
}
func Clay_LayoutElementArray_GetValue(array *Clay_LayoutElementArray, index int32) Clay_LayoutElement {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(index)))
	}
	return Clay_LayoutElement_DEFAULT
}
func Clay_LayoutElementArray_Add(array *Clay_LayoutElementArray, item Clay_LayoutElement) *Clay_LayoutElement {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_LayoutElement_DEFAULT
}
func Clay_LayoutElementArraySlice_Get(slice *Clay_LayoutElementArraySlice, index int32) *Clay_LayoutElement {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(index)))
	}
	return &Clay_LayoutElement_DEFAULT
}
func Clay_LayoutElementArray_RemoveSwapback(array *Clay_LayoutElementArray, index int32) Clay_LayoutElement {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_LayoutElement = *(*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(index)))
		*(*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(index))) = *(*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(array.Length)))
		return removed
	}
	return Clay_LayoutElement_DEFAULT
}
func Clay_LayoutElementArray_Set(array *Clay_LayoutElementArray, index int32, value Clay_LayoutElement) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_LayoutElement)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElement{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__ScrollContainerDataInternal struct {
	LayoutElement       *Clay_LayoutElement
	BoundingBox         Clay_BoundingBox
	ContentSize         Clay_Dimensions
	ScrollOrigin        Clay_Vector2
	PointerOrigin       Clay_Vector2
	ScrollMomentum      Clay_Vector2
	ScrollPosition      Clay_Vector2
	PreviousDelta       Clay_Vector2
	MomentumTime        float32
	ElementId           uint32
	OpenThisFrame       bool
	PointerScrollActive bool
}
type Clay__ScrollContainerDataInternalArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__ScrollContainerDataInternal
}
type Clay__ScrollContainerDataInternalArraySlice struct {
	Length        int32
	InternalArray *Clay__ScrollContainerDataInternal
}

var Clay__ScrollContainerDataInternal_DEFAULT Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternal{}

func Clay__ScrollContainerDataInternalArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__ScrollContainerDataInternalArray {
	return Clay__ScrollContainerDataInternalArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__ScrollContainerDataInternal)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__ScrollContainerDataInternal{})), arena))}
}
func Clay__ScrollContainerDataInternalArray_Get(array *Clay__ScrollContainerDataInternalArray, index int32) *Clay__ScrollContainerDataInternal {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(index)))
	}
	return &Clay__ScrollContainerDataInternal_DEFAULT
}
func Clay__ScrollContainerDataInternalArray_GetValue(array *Clay__ScrollContainerDataInternalArray, index int32) Clay__ScrollContainerDataInternal {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(index)))
	}
	return Clay__ScrollContainerDataInternal_DEFAULT
}
func Clay__ScrollContainerDataInternalArray_Add(array *Clay__ScrollContainerDataInternalArray, item Clay__ScrollContainerDataInternal) *Clay__ScrollContainerDataInternal {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__ScrollContainerDataInternal_DEFAULT
}
func Clay__ScrollContainerDataInternalArraySlice_Get(slice *Clay__ScrollContainerDataInternalArraySlice, index int32) *Clay__ScrollContainerDataInternal {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(index)))
	}
	return &Clay__ScrollContainerDataInternal_DEFAULT
}
func Clay__ScrollContainerDataInternalArray_RemoveSwapback(array *Clay__ScrollContainerDataInternalArray, index int32) Clay__ScrollContainerDataInternal {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__ScrollContainerDataInternal = *(*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(index)))
		*(*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(index))) = *(*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(array.Length)))
		return removed
	}
	return Clay__ScrollContainerDataInternal_DEFAULT
}
func Clay__ScrollContainerDataInternalArray_Set(array *Clay__ScrollContainerDataInternalArray, index int32, value Clay__ScrollContainerDataInternal) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__ScrollContainerDataInternal)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__ScrollContainerDataInternal{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__DebugElementData struct {
	Collision bool
	Collapsed bool
}
type Clay__DebugElementDataArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__DebugElementData
}
type Clay__DebugElementDataArraySlice struct {
	Length        int32
	InternalArray *Clay__DebugElementData
}

var Clay__DebugElementData_DEFAULT Clay__DebugElementData = Clay__DebugElementData{Collision: false}

func Clay__DebugElementDataArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__DebugElementDataArray {
	return Clay__DebugElementDataArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__DebugElementData)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__DebugElementData{})), arena))}
}
func Clay__DebugElementDataArray_Get(array *Clay__DebugElementDataArray, index int32) *Clay__DebugElementData {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(index)))
	}
	return &Clay__DebugElementData_DEFAULT
}
func Clay__DebugElementDataArray_GetValue(array *Clay__DebugElementDataArray, index int32) Clay__DebugElementData {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(index)))
	}
	return Clay__DebugElementData_DEFAULT
}
func Clay__DebugElementDataArray_Add(array *Clay__DebugElementDataArray, item Clay__DebugElementData) *Clay__DebugElementData {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__DebugElementData_DEFAULT
}
func Clay__DebugElementDataArraySlice_Get(slice *Clay__DebugElementDataArraySlice, index int32) *Clay__DebugElementData {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(index)))
	}
	return &Clay__DebugElementData_DEFAULT
}
func Clay__DebugElementDataArray_RemoveSwapback(array *Clay__DebugElementDataArray, index int32) Clay__DebugElementData {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__DebugElementData = *(*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(index)))
		*(*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(index))) = *(*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(array.Length)))
		return removed
	}
	return Clay__DebugElementData_DEFAULT
}
func Clay__DebugElementDataArray_Set(array *Clay__DebugElementDataArray, index int32, value Clay__DebugElementData) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__DebugElementData)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__DebugElementData{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay_LayoutElementHashMapItem struct {
	BoundingBox           Clay_BoundingBox
	ElementId             Clay_ElementId
	LayoutElement         *Clay_LayoutElement
	OnHoverFunction       func(elementId Clay_ElementId, pointerInfo Clay_PointerData, userData int64)
	HoverFunctionUserData int64
	NextIndex             int32
	Generation            uint32
	IdAlias               uint32
	DebugData             *Clay__DebugElementData
}
type Clay__LayoutElementHashMapItemArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay_LayoutElementHashMapItem
}
type Clay__LayoutElementHashMapItemArraySlice struct {
	Length        int32
	InternalArray *Clay_LayoutElementHashMapItem
}

var Clay_LayoutElementHashMapItem_DEFAULT Clay_LayoutElementHashMapItem = Clay_LayoutElementHashMapItem{}

func Clay__LayoutElementHashMapItemArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__LayoutElementHashMapItemArray {
	return Clay__LayoutElementHashMapItemArray{Capacity: capacity, Length: 0, InternalArray: (*Clay_LayoutElementHashMapItem)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay_LayoutElementHashMapItem{})), arena))}
}
func Clay__LayoutElementHashMapItemArray_Get(array *Clay__LayoutElementHashMapItemArray, index int32) *Clay_LayoutElementHashMapItem {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(index)))
	}
	return &Clay_LayoutElementHashMapItem_DEFAULT
}
func Clay__LayoutElementHashMapItemArray_GetValue(array *Clay__LayoutElementHashMapItemArray, index int32) Clay_LayoutElementHashMapItem {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(index)))
	}
	return Clay_LayoutElementHashMapItem_DEFAULT
}
func Clay__LayoutElementHashMapItemArray_Add(array *Clay__LayoutElementHashMapItemArray, item Clay_LayoutElementHashMapItem) *Clay_LayoutElementHashMapItem {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay_LayoutElementHashMapItem_DEFAULT
}
func Clay__LayoutElementHashMapItemArraySlice_Get(slice *Clay__LayoutElementHashMapItemArraySlice, index int32) *Clay_LayoutElementHashMapItem {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(index)))
	}
	return &Clay_LayoutElementHashMapItem_DEFAULT
}
func Clay__LayoutElementHashMapItemArray_RemoveSwapback(array *Clay__LayoutElementHashMapItemArray, index int32) Clay_LayoutElementHashMapItem {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay_LayoutElementHashMapItem = *(*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(index)))
		*(*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(index))) = *(*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(array.Length)))
		return removed
	}
	return Clay_LayoutElementHashMapItem_DEFAULT
}
func Clay__LayoutElementHashMapItemArray_Set(array *Clay__LayoutElementHashMapItemArray, index int32, value Clay_LayoutElementHashMapItem) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay_LayoutElementHashMapItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay_LayoutElementHashMapItem{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__MeasuredWord struct {
	StartOffset int32
	Length      int32
	Width       float32
	Next        int32
}
type Clay__MeasuredWordArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__MeasuredWord
}
type Clay__MeasuredWordArraySlice struct {
	Length        int32
	InternalArray *Clay__MeasuredWord
}

var Clay__MeasuredWord_DEFAULT Clay__MeasuredWord = Clay__MeasuredWord{}

func Clay__MeasuredWordArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__MeasuredWordArray {
	return Clay__MeasuredWordArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__MeasuredWord)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__MeasuredWord{})), arena))}
}
func Clay__MeasuredWordArray_Get(array *Clay__MeasuredWordArray, index int32) *Clay__MeasuredWord {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(index)))
	}
	return &Clay__MeasuredWord_DEFAULT
}
func Clay__MeasuredWordArray_GetValue(array *Clay__MeasuredWordArray, index int32) Clay__MeasuredWord {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(index)))
	}
	return Clay__MeasuredWord_DEFAULT
}
func Clay__MeasuredWordArray_Add(array *Clay__MeasuredWordArray, item Clay__MeasuredWord) *Clay__MeasuredWord {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__MeasuredWord_DEFAULT
}
func Clay__MeasuredWordArraySlice_Get(slice *Clay__MeasuredWordArraySlice, index int32) *Clay__MeasuredWord {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(index)))
	}
	return &Clay__MeasuredWord_DEFAULT
}
func Clay__MeasuredWordArray_RemoveSwapback(array *Clay__MeasuredWordArray, index int32) Clay__MeasuredWord {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__MeasuredWord = *(*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(index)))
		*(*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(index))) = *(*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(array.Length)))
		return removed
	}
	return Clay__MeasuredWord_DEFAULT
}
func Clay__MeasuredWordArray_Set(array *Clay__MeasuredWordArray, index int32, value Clay__MeasuredWord) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__MeasuredWord)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasuredWord{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__MeasureTextCacheItem struct {
	UnwrappedDimensions     Clay_Dimensions
	MeasuredWordsStartIndex int32
	ContainsNewlines        bool
	Id                      uint32
	NextIndex               int32
	Generation              uint32
}
type Clay__MeasureTextCacheItemArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__MeasureTextCacheItem
}
type Clay__MeasureTextCacheItemArraySlice struct {
	Length        int32
	InternalArray *Clay__MeasureTextCacheItem
}

var Clay__MeasureTextCacheItem_DEFAULT Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItem{}

func Clay__MeasureTextCacheItemArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__MeasureTextCacheItemArray {
	return Clay__MeasureTextCacheItemArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__MeasureTextCacheItem)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__MeasureTextCacheItem{})), arena))}
}
func Clay__MeasureTextCacheItemArray_Get(array *Clay__MeasureTextCacheItemArray, index int32) *Clay__MeasureTextCacheItem {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(index)))
	}
	return &Clay__MeasureTextCacheItem_DEFAULT
}
func Clay__MeasureTextCacheItemArray_GetValue(array *Clay__MeasureTextCacheItemArray, index int32) Clay__MeasureTextCacheItem {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(index)))
	}
	return Clay__MeasureTextCacheItem_DEFAULT
}
func Clay__MeasureTextCacheItemArray_Add(array *Clay__MeasureTextCacheItemArray, item Clay__MeasureTextCacheItem) *Clay__MeasureTextCacheItem {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__MeasureTextCacheItem_DEFAULT
}
func Clay__MeasureTextCacheItemArraySlice_Get(slice *Clay__MeasureTextCacheItemArraySlice, index int32) *Clay__MeasureTextCacheItem {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(index)))
	}
	return &Clay__MeasureTextCacheItem_DEFAULT
}
func Clay__MeasureTextCacheItemArray_RemoveSwapback(array *Clay__MeasureTextCacheItemArray, index int32) Clay__MeasureTextCacheItem {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__MeasureTextCacheItem = *(*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(index)))
		*(*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(index))) = *(*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(array.Length)))
		return removed
	}
	return Clay__MeasureTextCacheItem_DEFAULT
}
func Clay__MeasureTextCacheItemArray_Set(array *Clay__MeasureTextCacheItemArray, index int32, value Clay__MeasureTextCacheItem) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__MeasureTextCacheItem)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__MeasureTextCacheItem{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__LayoutElementTreeNode struct {
	LayoutElement   *Clay_LayoutElement
	Position        Clay_Vector2
	NextChildOffset Clay_Vector2
}
type Clay__LayoutElementTreeNodeArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__LayoutElementTreeNode
}
type Clay__LayoutElementTreeNodeArraySlice struct {
	Length        int32
	InternalArray *Clay__LayoutElementTreeNode
}

var Clay__LayoutElementTreeNode_DEFAULT Clay__LayoutElementTreeNode = Clay__LayoutElementTreeNode{}

func Clay__LayoutElementTreeNodeArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__LayoutElementTreeNodeArray {
	return Clay__LayoutElementTreeNodeArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__LayoutElementTreeNode)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__LayoutElementTreeNode{})), arena))}
}
func Clay__LayoutElementTreeNodeArray_Get(array *Clay__LayoutElementTreeNodeArray, index int32) *Clay__LayoutElementTreeNode {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(index)))
	}
	return &Clay__LayoutElementTreeNode_DEFAULT
}
func Clay__LayoutElementTreeNodeArray_GetValue(array *Clay__LayoutElementTreeNodeArray, index int32) Clay__LayoutElementTreeNode {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(index)))
	}
	return Clay__LayoutElementTreeNode_DEFAULT
}
func Clay__LayoutElementTreeNodeArray_Add(array *Clay__LayoutElementTreeNodeArray, item Clay__LayoutElementTreeNode) *Clay__LayoutElementTreeNode {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__LayoutElementTreeNode_DEFAULT
}
func Clay__LayoutElementTreeNodeArraySlice_Get(slice *Clay__LayoutElementTreeNodeArraySlice, index int32) *Clay__LayoutElementTreeNode {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(index)))
	}
	return &Clay__LayoutElementTreeNode_DEFAULT
}
func Clay__LayoutElementTreeNodeArray_RemoveSwapback(array *Clay__LayoutElementTreeNodeArray, index int32) Clay__LayoutElementTreeNode {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__LayoutElementTreeNode = *(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(index)))
		*(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(index))) = *(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(array.Length)))
		return removed
	}
	return Clay__LayoutElementTreeNode_DEFAULT
}
func Clay__LayoutElementTreeNodeArray_Set(array *Clay__LayoutElementTreeNodeArray, index int32, value Clay__LayoutElementTreeNode) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}

type Clay__LayoutElementTreeRoot struct {
	LayoutElementIndex int32
	ParentId           uint32
	ClipElementId      uint32
	ZIndex             int16
	PointerOffset      Clay_Vector2
}
type Clay__LayoutElementTreeRootArray struct {
	Capacity      int32
	Length        int32
	InternalArray *Clay__LayoutElementTreeRoot
}
type Clay__LayoutElementTreeRootArraySlice struct {
	Length        int32
	InternalArray *Clay__LayoutElementTreeRoot
}

var Clay__LayoutElementTreeRoot_DEFAULT Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRoot{}

func Clay__LayoutElementTreeRootArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__LayoutElementTreeRootArray {
	return Clay__LayoutElementTreeRootArray{Capacity: capacity, Length: 0, InternalArray: (*Clay__LayoutElementTreeRoot)(Clay__Array_Allocate_Arena(capacity, uint32(unsafe.Sizeof(Clay__LayoutElementTreeRoot{})), arena))}
}
func Clay__LayoutElementTreeRootArray_Get(array *Clay__LayoutElementTreeRootArray, index int32) *Clay__LayoutElementTreeRoot {
	if Clay__Array_RangeCheck(index, array.Length) {
		return (*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(index)))
	}
	return &Clay__LayoutElementTreeRoot_DEFAULT
}
func Clay__LayoutElementTreeRootArray_GetValue(array *Clay__LayoutElementTreeRootArray, index int32) Clay__LayoutElementTreeRoot {
	if Clay__Array_RangeCheck(index, array.Length) {
		return *(*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(index)))
	}
	return Clay__LayoutElementTreeRoot_DEFAULT
}
func Clay__LayoutElementTreeRootArray_Add(array *Clay__LayoutElementTreeRootArray, item Clay__LayoutElementTreeRoot) *Clay__LayoutElementTreeRoot {
	if Clay__Array_AddCapacityCheck(array.Length, array.Capacity) {
		*(*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(int64(array.Length)-1)))
	}
	return &Clay__LayoutElementTreeRoot_DEFAULT
}
func Clay__LayoutElementTreeRootArraySlice_Get(slice *Clay__LayoutElementTreeRootArraySlice, index int32) *Clay__LayoutElementTreeRoot {
	if Clay__Array_RangeCheck(index, slice.Length) {
		return (*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(slice.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(index)))
	}
	return &Clay__LayoutElementTreeRoot_DEFAULT
}
func Clay__LayoutElementTreeRootArray_RemoveSwapback(array *Clay__LayoutElementTreeRootArray, index int32) Clay__LayoutElementTreeRoot {
	if Clay__Array_RangeCheck(index, array.Length) {
		array.Length--
		var removed Clay__LayoutElementTreeRoot = *(*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(index)))
		*(*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(index))) = *(*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(array.Length)))
		return removed
	}
	return Clay__LayoutElementTreeRoot_DEFAULT
}
func Clay__LayoutElementTreeRootArray_Set(array *Clay__LayoutElementTreeRootArray, index int32, value Clay__LayoutElementTreeRoot) {
	if Clay__Array_RangeCheck(index, array.Capacity) {
		*(*Clay__LayoutElementTreeRoot)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeRoot{})*uintptr(index))) = value
		if int64(index) < int64(array.Length) {
			array.Length = array.Length
		} else {
			array.Length = int32(int64(index) + 1)
		}
	}
}
func Clay__Context_Allocate_Arena(arena *Clay_Arena) *Clay_Context {
	var (
		totalSizeBytes  uint64 = uint64(unsafe.Sizeof(Clay_Context{}))
		memoryAddress   uint64 = uint64(uintptr(unsafe.Pointer(arena.Memory)))
		nextAllocOffset uint64 = (memoryAddress % 64)
	)
	if nextAllocOffset+totalSizeBytes > arena.Capacity {
		return nil
	}
	arena.NextAllocation = nextAllocOffset + totalSizeBytes
	return (*Clay_Context)(unsafe.Pointer(uintptr(memoryAddress + nextAllocOffset)))
}
func Clay__WriteStringToCharBuffer(buffer *Clay__charArray, string_ Clay_String) Clay_String {
	for i := int32(0); int64(i) < int64(string_.Length); i++ {
		*(*byte)(unsafe.Add(unsafe.Pointer(buffer.InternalArray), int64(buffer.Length)+int64(i))) = *(*byte)(unsafe.Add(unsafe.Pointer(string_.Chars), i))
	}
	buffer.Length += string_.Length
	return Clay_String{Length: string_.Length, Chars: ((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(buffer.InternalArray), buffer.Length))), -string_.Length)))}
}

var Clay__MeasureText func(text Clay_StringSlice, config *Clay_TextElementConfig, userData unsafe.Pointer) Clay_Dimensions
var Clay__QueryScrollOffset func(elementId uint32, userData unsafe.Pointer) Clay_Vector2

func Clay__GetOpenLayoutElement() *Clay_LayoutElement {
	var context *Clay_Context = Clay_GetCurrentContext()
	return Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&context.OpenLayoutElementStack, int32(int64(context.OpenLayoutElementStack.Length)-1)))
}
func Clay__GetParentElementId() uint32 {
	var context *Clay_Context = Clay_GetCurrentContext()
	return Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&context.OpenLayoutElementStack, int32(int64(context.OpenLayoutElementStack.Length)-2))).Id
}
func Clay__StoreLayoutConfig(config Clay_LayoutConfig) *Clay_LayoutConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &CLAY_LAYOUT_DEFAULT
	}
	return Clay__LayoutConfigArray_Add(&Clay_GetCurrentContext().LayoutConfigs, config)
}
func Clay__StoreTextElementConfig(config Clay_TextElementConfig) *Clay_TextElementConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &Clay_TextElementConfig_DEFAULT
	}
	return Clay__TextElementConfigArray_Add(&Clay_GetCurrentContext().TextElementConfigs, config)
}
func Clay__StoreImageElementConfig(config Clay_ImageElementConfig) *Clay_ImageElementConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &Clay_ImageElementConfig_DEFAULT
	}
	return Clay__ImageElementConfigArray_Add(&Clay_GetCurrentContext().ImageElementConfigs, config)
}
func Clay__StoreFloatingElementConfig(config Clay_FloatingElementConfig) *Clay_FloatingElementConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &Clay_FloatingElementConfig_DEFAULT
	}
	return Clay__FloatingElementConfigArray_Add(&Clay_GetCurrentContext().FloatingElementConfigs, config)
}
func Clay__StoreCustomElementConfig(config Clay_CustomElementConfig) *Clay_CustomElementConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &Clay_CustomElementConfig_DEFAULT
	}
	return Clay__CustomElementConfigArray_Add(&Clay_GetCurrentContext().CustomElementConfigs, config)
}
func Clay__StoreScrollElementConfig(config Clay_ScrollElementConfig) *Clay_ScrollElementConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &Clay_ScrollElementConfig_DEFAULT
	}
	return Clay__ScrollElementConfigArray_Add(&Clay_GetCurrentContext().ScrollElementConfigs, config)
}
func Clay__StoreBorderElementConfig(config Clay_BorderElementConfig) *Clay_BorderElementConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &Clay_BorderElementConfig_DEFAULT
	}
	return Clay__BorderElementConfigArray_Add(&Clay_GetCurrentContext().BorderElementConfigs, config)
}
func Clay__StoreSharedElementConfig(config Clay_SharedElementConfig) *Clay_SharedElementConfig {
	if Clay_GetCurrentContext().BooleanWarnings.MaxElementsExceeded {
		return &Clay_SharedElementConfig_DEFAULT
	}
	return Clay__SharedElementConfigArray_Add(&Clay_GetCurrentContext().SharedElementConfigs, config)
}
func Clay__AttachElementConfig(config Clay_ElementConfigUnion, type_ Clay__ElementConfigType) Clay_ElementConfig {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return Clay_ElementConfig{}
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	openLayoutElement.ElementConfigs.Length++
	return *Clay__ElementConfigArray_Add(&context.ElementConfigs, Clay_ElementConfig{Type: type_, Config: config})
}
func Clay__FindElementConfigWithType(element *Clay_LayoutElement, type_ Clay__ElementConfigType) Clay_ElementConfigUnion {
	for i := int32(0); int64(i) < int64(element.ElementConfigs.Length); i++ {
		var config *Clay_ElementConfig = Clay__ElementConfigArraySlice_Get(&element.ElementConfigs, i)
		if config.Type == type_ {
			return config.Config
		}
	}
	return Clay_ElementConfigUnion{}
}
func Clay__HashNumber(offset uint32, seed uint32) Clay_ElementId {
	var hash uint32 = seed
	hash += uint32(int32(int64(offset) + 48))
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(int32(int64(hash) << 3))
	hash ^= uint32(int32(int64(hash) >> 11))
	hash += uint32(int32(int64(hash) << 15))
	return Clay_ElementId{Id: uint32(int32(int64(hash) + 1)), Offset: offset, BaseId: seed, StringId: CLAY__STRING_DEFAULT}
}
func Clay__HashString(key Clay_String, offset uint32, seed uint32) Clay_ElementId {
	var (
		hash uint32 = 0
		base uint32 = seed
	)
	for i := int32(0); int64(i) < int64(key.Length); i++ {
		base += uint32(*(*byte)(unsafe.Add(unsafe.Pointer(key.Chars), i)))
		base += uint32(int32(int64(base) << 10))
		base ^= uint32(int32(int64(base) >> 6))
	}
	hash = base
	hash += offset
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(int32(int64(hash) << 3))
	base += uint32(int32(int64(base) << 3))
	hash ^= uint32(int32(int64(hash) >> 11))
	base ^= uint32(int32(int64(base) >> 11))
	hash += uint32(int32(int64(hash) << 15))
	base += uint32(int32(int64(base) << 15))
	return Clay_ElementId{Id: uint32(int32(int64(hash) + 1)), Offset: offset, BaseId: uint32(int32(int64(base) + 1)), StringId: key}
}
func Clay__HashTextWithConfig(text *Clay_String, config *Clay_TextElementConfig) uint32 {
	var (
		hash            uint32 = 0
		pointerAsNumber uint64 = uint64(uintptr(unsafe.Pointer(text.Chars)))
	)
	if config.HashStringContents {
		var maxLengthToHash uint32 = uint32(int32(func() int64 {
			if int64(text.Length) < 256 {
				return int64(text.Length)
			}
			return 256
		}()))
		for i := uint32(0); int64(i) < int64(maxLengthToHash); i++ {
			hash += uint32(*(*byte)(unsafe.Add(unsafe.Pointer(text.Chars), i)))
			hash += uint32(int32(int64(hash) << 10))
			hash ^= uint32(int32(int64(hash) >> 6))
		}
	} else {
		hash += uint32(pointerAsNumber)
		hash += uint32(int32(int64(hash) << 10))
		hash ^= uint32(int32(int64(hash) >> 6))
	}
	hash += uint32(text.Length)
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(config.FontId)
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(config.FontSize)
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(config.LineHeight)
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(config.LetterSpacing)
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(int32(config.WrapMode))
	hash += uint32(int32(int64(hash) << 10))
	hash ^= uint32(int32(int64(hash) >> 6))
	hash += uint32(int32(int64(hash) << 3))
	hash ^= uint32(int32(int64(hash) >> 11))
	hash += uint32(int32(int64(hash) << 15))
	return uint32(int32(int64(hash) + 1))
}
func Clay__AddMeasuredWord(word Clay__MeasuredWord, previousWord *Clay__MeasuredWord) *Clay__MeasuredWord {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int64(context.MeasuredWordsFreeList.Length) > 0 {
		var newItemIndex uint32 = uint32(Clay__int32_tArray_GetValue(&context.MeasuredWordsFreeList, int32(int64(context.MeasuredWordsFreeList.Length)-1)))
		context.MeasuredWordsFreeList.Length--
		Clay__MeasuredWordArray_Set(&context.MeasuredWords, int32(int64(newItemIndex)), word)
		previousWord.Next = int32(newItemIndex)
		return Clay__MeasuredWordArray_Get(&context.MeasuredWords, int32(int64(newItemIndex)))
	} else {
		previousWord.Next = context.MeasuredWords.Length
		return Clay__MeasuredWordArray_Add(&context.MeasuredWords, word)
	}
}
func Clay__MeasureTextCached(text *Clay_String, config *Clay_TextElementConfig) *Clay__MeasureTextCacheItem {
	var context *Clay_Context = Clay_GetCurrentContext()
	if Clay__MeasureText == nil {
		if !context.BooleanWarnings.TextMeasurementFunctionNotSet {
			context.BooleanWarnings.TextMeasurementFunctionNotSet = true
			context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_TEXT_MEASUREMENT_FUNCTION_NOT_PROVIDED, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay's internal MeasureText function is null. You may have forgotten to call Clay_SetMeasureTextFunction(), or passed a NULL function pointer by mistake.")}, UserData: context.ErrorHandler.UserData})
		}
		return &Clay__MeasureTextCacheItem_DEFAULT
	}
	var id uint32 = Clay__HashTextWithConfig(text, config)
	var hashBucket uint32 = uint32(int32(int64(id) % (int64(context.MaxMeasureTextCacheWordCount) / 32)))
	var elementIndexPrevious int32 = 0
	var elementIndex int32 = *(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	for int64(elementIndex) != 0 {
		var hashEntry *Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, elementIndex)
		if int64(hashEntry.Id) == int64(id) {
			hashEntry.Generation = context.Generation
			return hashEntry
		}
		if int64(context.Generation)-int64(hashEntry.Generation) > 2 {
			var nextWordIndex int32 = hashEntry.MeasuredWordsStartIndex
			for int64(nextWordIndex) != -1 {
				var measuredWord *Clay__MeasuredWord = Clay__MeasuredWordArray_Get(&context.MeasuredWords, nextWordIndex)
				Clay__int32_tArray_Add(&context.MeasuredWordsFreeList, nextWordIndex)
				nextWordIndex = measuredWord.Next
			}
			var nextIndex int32 = hashEntry.NextIndex
			Clay__MeasureTextCacheItemArray_Set(&context.MeasureTextHashMapInternal, elementIndex, Clay__MeasureTextCacheItem{MeasuredWordsStartIndex: -1})
			Clay__int32_tArray_Add(&context.MeasureTextHashMapInternalFreeList, elementIndex)
			if int64(elementIndexPrevious) == 0 {
				*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = nextIndex
			} else {
				var previousHashEntry *Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, elementIndexPrevious)
				previousHashEntry.NextIndex = nextIndex
			}
			elementIndex = nextIndex
		} else {
			elementIndexPrevious = elementIndex
			elementIndex = hashEntry.NextIndex
		}
	}
	var newItemIndex int32 = 0
	var newCacheItem Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItem{MeasuredWordsStartIndex: -1, Id: id, Generation: context.Generation}
	var measured *Clay__MeasureTextCacheItem = nil
	if int64(context.MeasureTextHashMapInternalFreeList.Length) > 0 {
		newItemIndex = Clay__int32_tArray_GetValue(&context.MeasureTextHashMapInternalFreeList, int32(int64(context.MeasureTextHashMapInternalFreeList.Length)-1))
		context.MeasureTextHashMapInternalFreeList.Length--
		Clay__MeasureTextCacheItemArray_Set(&context.MeasureTextHashMapInternal, newItemIndex, newCacheItem)
		measured = Clay__MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, newItemIndex)
	} else {
		if int64(context.MeasureTextHashMapInternal.Length) == int64(context.MeasureTextHashMapInternal.Capacity)-1 {
			if context.BooleanWarnings.MaxTextMeasureCacheExceeded {
				context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay ran out of capacity while attempting to measure text elements. Try using Clay_SetMaxElementCount() with a higher value.")}, UserData: context.ErrorHandler.UserData})
				context.BooleanWarnings.MaxTextMeasureCacheExceeded = true
			}
			return &Clay__MeasureTextCacheItem_DEFAULT
		}
		measured = Clay__MeasureTextCacheItemArray_Add(&context.MeasureTextHashMapInternal, newCacheItem)
		newItemIndex = int32(int64(context.MeasureTextHashMapInternal.Length) - 1)
	}
	var start int32 = 0
	var end int32 = 0
	var lineWidth float32 = 0
	var measuredWidth float32 = 0
	var measuredHeight float32 = 0
	var spaceWidth float32 = Clay__MeasureText(Clay_StringSlice{Length: 1, Chars: CLAY__SPACECHAR.Chars, BaseChars: CLAY__SPACECHAR.Chars}, config, context.MeasureTextUserData).Width
	var tempWord Clay__MeasuredWord = Clay__MeasuredWord{Next: -1}
	var previousWord *Clay__MeasuredWord = &tempWord
	for int64(end) < int64(text.Length) {
		if int64(context.MeasuredWords.Length) == int64(context.MeasuredWords.Capacity)-1 {
			if !context.BooleanWarnings.MaxTextMeasureCacheExceeded {
				context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_TEXT_MEASUREMENT_CAPACITY_EXCEEDED, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay has run out of space in it's internal text measurement cache. Try using Clay_SetMaxMeasureTextCacheWordCount() (default 16384, with 1 unit storing 1 measured word).")}, UserData: context.ErrorHandler.UserData})
				context.BooleanWarnings.MaxTextMeasureCacheExceeded = true
			}
			return &Clay__MeasureTextCacheItem_DEFAULT
		}
		var current int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(text.Chars), end)))
		if int64(current) == ' ' || int64(current) == '\n' {
			var (
				length     int32           = int32(int64(end) - int64(start))
				dimensions Clay_Dimensions = Clay__MeasureText(Clay_StringSlice{Length: length, Chars: (*byte)(unsafe.Add(unsafe.Pointer(text.Chars), start)), BaseChars: text.Chars}, config, context.MeasureTextUserData)
			)
			if measuredHeight > dimensions.Height {
				measuredHeight = measuredHeight
			} else {
				measuredHeight = dimensions.Height
			}
			if int64(current) == ' ' {
				dimensions.Width += spaceWidth
				previousWord = Clay__AddMeasuredWord(Clay__MeasuredWord{StartOffset: start, Length: int32(int64(length) + 1), Width: dimensions.Width, Next: -1}, previousWord)
				lineWidth += dimensions.Width
			}
			if int64(current) == '\n' {
				if int64(length) > 0 {
					previousWord = Clay__AddMeasuredWord(Clay__MeasuredWord{StartOffset: start, Length: length, Width: dimensions.Width, Next: -1}, previousWord)
				}
				previousWord = Clay__AddMeasuredWord(Clay__MeasuredWord{StartOffset: int32(int64(end) + 1), Length: 0, Width: 0, Next: -1}, previousWord)
				lineWidth += dimensions.Width
				if lineWidth > measuredWidth {
					measuredWidth = lineWidth
				} else {
					measuredWidth = measuredWidth
				}
				measured.ContainsNewlines = true
				lineWidth = 0
			}
			start = int32(int64(end) + 1)
		}
		end++
	}
	if int64(end)-int64(start) > 0 {
		var dimensions Clay_Dimensions = Clay__MeasureText(Clay_StringSlice{Length: int32(int64(end) - int64(start)), Chars: (*byte)(unsafe.Add(unsafe.Pointer(text.Chars), start)), BaseChars: text.Chars}, config, context.MeasureTextUserData)
		Clay__AddMeasuredWord(Clay__MeasuredWord{StartOffset: start, Length: int32(int64(end) - int64(start)), Width: dimensions.Width, Next: -1}, previousWord)
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
	if int64(elementIndexPrevious) != 0 {
		Clay__MeasureTextCacheItemArray_Get(&context.MeasureTextHashMapInternal, elementIndexPrevious).NextIndex = newItemIndex
	} else {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = newItemIndex
	}
	return measured
}
func Clay__PointIsInsideRect(point Clay_Vector2, rect Clay_BoundingBox) bool {
	return point.X >= rect.X && point.X <= rect.X+rect.Width && point.Y >= rect.Y && point.Y <= rect.Y+rect.Height
}
func Clay__AddHashMapItem(elementId Clay_ElementId, layoutElement *Clay_LayoutElement, idAlias uint32) *Clay_LayoutElementHashMapItem {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int64(context.LayoutElementsHashMapInternal.Length) == int64(context.LayoutElementsHashMapInternal.Capacity)-1 {
		return nil
	}
	var item Clay_LayoutElementHashMapItem = Clay_LayoutElementHashMapItem{ElementId: elementId, LayoutElement: layoutElement, NextIndex: -1, Generation: uint32(int32(int64(context.Generation) + 1)), IdAlias: idAlias}
	var hashBucket uint32 = uint32(int32(int64(elementId.Id) % int64(context.LayoutElementsHashMap.Capacity)))
	var hashItemPrevious int32 = -1
	var hashItemIndex int32 = *(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	for int64(hashItemIndex) != -1 {
		var hashItem *Clay_LayoutElementHashMapItem = Clay__LayoutElementHashMapItemArray_Get(&context.LayoutElementsHashMapInternal, hashItemIndex)
		if int64(hashItem.ElementId.Id) == int64(elementId.Id) {
			item.NextIndex = hashItem.NextIndex
			if int64(hashItem.Generation) <= int64(context.Generation) {
				hashItem.ElementId = elementId
				hashItem.Generation = uint32(int32(int64(context.Generation) + 1))
				hashItem.LayoutElement = layoutElement
				hashItem.DebugData.Collision = false
			} else {
				context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_DUPLICATE_ID, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("An element with this ID was already previously declared during this layout.")}, UserData: context.ErrorHandler.UserData})
				if context.DebugModeEnabled {
					hashItem.DebugData.Collision = true
				}
			}
			return hashItem
		}
		hashItemPrevious = hashItemIndex
		hashItemIndex = hashItem.NextIndex
	}
	var hashItem *Clay_LayoutElementHashMapItem = Clay__LayoutElementHashMapItemArray_Add(&context.LayoutElementsHashMapInternal, item)
	hashItem.DebugData = Clay__DebugElementDataArray_Add(&context.DebugElementData, Clay__DebugElementData{Collision: false})
	if int64(hashItemPrevious) != -1 {
		Clay__LayoutElementHashMapItemArray_Get(&context.LayoutElementsHashMapInternal, hashItemPrevious).NextIndex = int32(int64(context.LayoutElementsHashMapInternal.Length) - 1)
	} else {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = int32(int64(context.LayoutElementsHashMapInternal.Length) - 1)
	}
	return hashItem
}
func Clay__GetHashMapItem(id uint32) *Clay_LayoutElementHashMapItem {
	var (
		context      *Clay_Context = Clay_GetCurrentContext()
		hashBucket   uint32        = uint32(int32(int64(id) % int64(context.LayoutElementsHashMap.Capacity)))
		elementIndex int32         = *(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	)
	for int64(elementIndex) != -1 {
		var hashEntry *Clay_LayoutElementHashMapItem = Clay__LayoutElementHashMapItemArray_Get(&context.LayoutElementsHashMapInternal, elementIndex)
		if int64(hashEntry.ElementId.Id) == int64(id) {
			return hashEntry
		}
		elementIndex = hashEntry.NextIndex
	}
	return &Clay_LayoutElementHashMapItem_DEFAULT
}
func Clay__GenerateIdForAnonymousElement(openLayoutElement *Clay_LayoutElement) Clay_ElementId {
	var (
		context       *Clay_Context       = Clay_GetCurrentContext()
		parentElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&context.OpenLayoutElementStack, int32(int64(context.OpenLayoutElementStack.Length)-2)))
		elementId     Clay_ElementId      = Clay__HashNumber(uint32(parentElement.ChildrenOrTextContent.Children.Length), parentElement.Id)
	)
	openLayoutElement.Id = elementId.Id
	Clay__AddHashMapItem(elementId, openLayoutElement, 0)
	Clay__StringArray_Add(&context.LayoutElementIdStrings, elementId.StringId)
	return elementId
}
func Clay__ElementHasConfig(layoutElement *Clay_LayoutElement, type_ Clay__ElementConfigType) bool {
	for i := int32(0); int64(i) < int64(layoutElement.ElementConfigs.Length); i++ {
		if Clay__ElementConfigArraySlice_Get(&layoutElement.ElementConfigs, i).Type == type_ {
			return true
		}
	}
	return false
}
func Clay__CloseElement() {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	var layoutConfig *Clay_LayoutConfig = openLayoutElement.LayoutConfig
	var elementHasScrollHorizontal bool = false
	var elementHasScrollVertical bool = false
	for i := int32(0); int64(i) < int64(openLayoutElement.ElementConfigs.Length); i++ {
		var config *Clay_ElementConfig = Clay__ElementConfigArraySlice_Get(&openLayoutElement.ElementConfigs, i)
		if config.Type == CLAY__ELEMENT_CONFIG_TYPE_SCROLL {
			elementHasScrollHorizontal = config.Config.ScrollElementConfig.Horizontal
			elementHasScrollVertical = config.Config.ScrollElementConfig.Vertical
			context.OpenClipElementStack.Length--
			break
		}
	}
	openLayoutElement.ChildrenOrTextContent.Children.Elements = (*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementChildren.InternalArray), unsafe.Sizeof(int32(0))*uintptr(context.LayoutElementChildren.Length)))
	if layoutConfig.LayoutDirection == CLAY_LEFT_TO_RIGHT {
		openLayoutElement.Dimensions.Width = float32(int64(layoutConfig.Padding.Left) + int64(layoutConfig.Padding.Right))
		for i := int32(0); int64(i) < int64(openLayoutElement.ChildrenOrTextContent.Children.Length); i++ {
			var (
				childIndex int32               = Clay__int32_tArray_GetValue(&context.LayoutElementChildrenBuffer, int32(int64(context.LayoutElementChildrenBuffer.Length)-int64(openLayoutElement.ChildrenOrTextContent.Children.Length)+int64(i)))
				child      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, childIndex)
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
			Clay__int32_tArray_Add(&context.LayoutElementChildren, childIndex)
		}
		var childGap float32 = float32((func() int64 {
			if (int64(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
				return int64(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1
			}
			return 0
		}()) * int64(layoutConfig.ChildGap))
		openLayoutElement.Dimensions.Width += childGap
		openLayoutElement.MinDimensions.Width += childGap
	} else if layoutConfig.LayoutDirection == CLAY_TOP_TO_BOTTOM {
		openLayoutElement.Dimensions.Height = float32(int64(layoutConfig.Padding.Top) + int64(layoutConfig.Padding.Bottom))
		for i := int32(0); int64(i) < int64(openLayoutElement.ChildrenOrTextContent.Children.Length); i++ {
			var (
				childIndex int32               = Clay__int32_tArray_GetValue(&context.LayoutElementChildrenBuffer, int32(int64(context.LayoutElementChildrenBuffer.Length)-int64(openLayoutElement.ChildrenOrTextContent.Children.Length)+int64(i)))
				child      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, childIndex)
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
			Clay__int32_tArray_Add(&context.LayoutElementChildren, childIndex)
		}
		var childGap float32 = float32((func() int64 {
			if (int64(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
				return int64(openLayoutElement.ChildrenOrTextContent.Children.Length) - 1
			}
			return 0
		}()) * int64(layoutConfig.ChildGap))
		openLayoutElement.Dimensions.Height += childGap
		openLayoutElement.MinDimensions.Height += childGap
	}
	context.LayoutElementChildrenBuffer.Length -= int32(openLayoutElement.ChildrenOrTextContent.Children.Length)
	if layoutConfig.Sizing.Width.Type != CLAY__SIZING_TYPE_PERCENT {
		if layoutConfig.Sizing.Width.Size.MinMax.Max <= 0 {
			layoutConfig.Sizing.Width.Size.MinMax.Max = CLAY__MAXFLOAT
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
	if layoutConfig.Sizing.Height.Type != CLAY__SIZING_TYPE_PERCENT {
		if layoutConfig.Sizing.Height.Size.MinMax.Max <= 0 {
			layoutConfig.Sizing.Height.Size.MinMax.Max = CLAY__MAXFLOAT
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
	var elementIsFloating bool = Clay__ElementHasConfig(openLayoutElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING)
	var closingElementIndex int32 = Clay__int32_tArray_RemoveSwapback(&context.OpenLayoutElementStack, int32(int64(context.OpenLayoutElementStack.Length)-1))
	openLayoutElement = Clay__GetOpenLayoutElement()
	if !elementIsFloating && int64(context.OpenLayoutElementStack.Length) > 1 {
		openLayoutElement.ChildrenOrTextContent.Children.Length++
		Clay__int32_tArray_Add(&context.LayoutElementChildrenBuffer, closingElementIndex)
	}
}
func Clay__MemCmp(s1 *byte, s2 *byte, length int32) bool {
	for i := int32(0); int64(i) < int64(length); i++ {
		if *(*byte)(unsafe.Add(unsafe.Pointer(s1), i)) != *(*byte)(unsafe.Add(unsafe.Pointer(s2), i)) {
			return false
		}
	}
	return true
}
func Clay__OpenElement() {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int64(context.LayoutElements.Length) == int64(context.LayoutElements.Capacity)-1 || context.BooleanWarnings.MaxElementsExceeded {
		context.BooleanWarnings.MaxElementsExceeded = true
		return
	}
	var layoutElement Clay_LayoutElement = Clay_LayoutElement{}
	Clay_LayoutElementArray_Add(&context.LayoutElements, layoutElement)
	Clay__int32_tArray_Add(&context.OpenLayoutElementStack, int32(int64(context.LayoutElements.Length)-1))
	if int64(context.OpenClipElementStack.Length) > 0 {
		Clay__int32_tArray_Set(&context.LayoutElementClipElementIds, int32(int64(context.LayoutElements.Length)-1), Clay__int32_tArray_GetValue(&context.OpenClipElementStack, int32(int64(context.OpenClipElementStack.Length)-1)))
	} else {
		Clay__int32_tArray_Set(&context.LayoutElementClipElementIds, int32(int64(context.LayoutElements.Length)-1), 0)
	}
}
func Clay__OpenTextElement(text Clay_String, textConfig *Clay_TextElementConfig) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int64(context.LayoutElements.Length) == int64(context.LayoutElements.Capacity)-1 || context.BooleanWarnings.MaxElementsExceeded {
		context.BooleanWarnings.MaxElementsExceeded = true
		return
	}
	var parentElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	var layoutElement Clay_LayoutElement = Clay_LayoutElement{}
	var textElement *Clay_LayoutElement = Clay_LayoutElementArray_Add(&context.LayoutElements, layoutElement)
	if int64(context.OpenClipElementStack.Length) > 0 {
		Clay__int32_tArray_Set(&context.LayoutElementClipElementIds, int32(int64(context.LayoutElements.Length)-1), Clay__int32_tArray_GetValue(&context.OpenClipElementStack, int32(int64(context.OpenClipElementStack.Length)-1)))
	} else {
		Clay__int32_tArray_Set(&context.LayoutElementClipElementIds, int32(int64(context.LayoutElements.Length)-1), 0)
	}
	Clay__int32_tArray_Add(&context.LayoutElementChildrenBuffer, int32(int64(context.LayoutElements.Length)-1))
	var textMeasured *Clay__MeasureTextCacheItem = Clay__MeasureTextCached(&text, textConfig)
	var elementId Clay_ElementId = Clay__HashNumber(uint32(parentElement.ChildrenOrTextContent.Children.Length), parentElement.Id)
	textElement.Id = elementId.Id
	Clay__AddHashMapItem(elementId, textElement, 0)
	Clay__StringArray_Add(&context.LayoutElementIdStrings, elementId.StringId)
	var textDimensions Clay_Dimensions = Clay_Dimensions{Width: textMeasured.UnwrappedDimensions.Width, Height: func() float32 {
		if int64(textConfig.LineHeight) > 0 {
			return float32(textConfig.LineHeight)
		}
		return textMeasured.UnwrappedDimensions.Height
	}()}
	textElement.Dimensions = textDimensions
	textElement.MinDimensions = Clay_Dimensions{Width: textMeasured.UnwrappedDimensions.Height, Height: textDimensions.Height}
	textElement.ChildrenOrTextContent.TextElementData = Clay__TextElementDataArray_Add(&context.TextElementData, Clay__TextElementData{Text: text, PreferredDimensions: textMeasured.UnwrappedDimensions, ElementIndex: int32(int64(context.LayoutElements.Length) - 1)})
	textElement.ElementConfigs = Clay__ElementConfigArraySlice{Length: 1, InternalArray: Clay__ElementConfigArray_Add(&context.ElementConfigs, Clay_ElementConfig{Type: CLAY__ELEMENT_CONFIG_TYPE_TEXT, Config: Clay_ElementConfigUnion{TextElementConfig: textConfig}})}
	textElement.LayoutConfig = &CLAY_LAYOUT_DEFAULT
	parentElement.ChildrenOrTextContent.Children.Length++
}
func Clay__AttachId(elementId Clay_ElementId) Clay_ElementId {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return Clay_ElementId_DEFAULT
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	var idAlias uint32 = openLayoutElement.Id
	openLayoutElement.Id = elementId.Id
	Clay__AddHashMapItem(elementId, openLayoutElement, idAlias)
	Clay__StringArray_Add(&context.LayoutElementIdStrings, elementId.StringId)
	return elementId
}
func Clay__ConfigureOpenElement(declaration Clay_ElementDeclaration) {
	var (
		context           *Clay_Context       = Clay_GetCurrentContext()
		openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	)
	openLayoutElement.LayoutConfig = Clay__StoreLayoutConfig(declaration.Layout)
	if declaration.Layout.Sizing.Width.Type == CLAY__SIZING_TYPE_PERCENT && declaration.Layout.Sizing.Width.Size.Percent > 1 || declaration.Layout.Sizing.Height.Type == CLAY__SIZING_TYPE_PERCENT && declaration.Layout.Sizing.Height.Size.Percent > 1 {
		context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_PERCENTAGE_OVER_1, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("An element was configured with CLAY_SIZING_PERCENT, but the provided percentage value was over 1.0. Clay expects a value between 0 and 1, i.e. 20% is 0.2.")}, UserData: context.ErrorHandler.UserData})
	}
	var openLayoutElementId Clay_ElementId = declaration.Id
	openLayoutElement.ElementConfigs.InternalArray = (*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(context.ElementConfigs.InternalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(context.ElementConfigs.Length)))
	var sharedConfig *Clay_SharedElementConfig = nil
	if declaration.BackgroundColor.A > 0 {
		sharedConfig = Clay__StoreSharedElementConfig(Clay_SharedElementConfig{BackgroundColor: declaration.BackgroundColor})
		Clay__AttachElementConfig(Clay_ElementConfigUnion{SharedElementConfig: sharedConfig}, CLAY__ELEMENT_CONFIG_TYPE_SHARED)
	}
	if !Clay__MemCmp((*byte)(unsafe.Pointer(&declaration.CornerRadius)), (*byte)(unsafe.Pointer(&Clay__CornerRadius_DEFAULT)), int32(uint32(unsafe.Sizeof(Clay_CornerRadius{})))) {
		if sharedConfig != nil {
			sharedConfig.CornerRadius = declaration.CornerRadius
		} else {
			sharedConfig = Clay__StoreSharedElementConfig(Clay_SharedElementConfig{CornerRadius: declaration.CornerRadius})
			Clay__AttachElementConfig(Clay_ElementConfigUnion{SharedElementConfig: sharedConfig}, CLAY__ELEMENT_CONFIG_TYPE_SHARED)
		}
	}
	if declaration.UserData != nil {
		if sharedConfig != nil {
			sharedConfig.UserData = declaration.UserData
		} else {
			sharedConfig = Clay__StoreSharedElementConfig(Clay_SharedElementConfig{UserData: declaration.UserData})
			Clay__AttachElementConfig(Clay_ElementConfigUnion{SharedElementConfig: sharedConfig}, CLAY__ELEMENT_CONFIG_TYPE_SHARED)
		}
	}
	if declaration.Image.ImageData != nil {
		Clay__AttachElementConfig(Clay_ElementConfigUnion{ImageElementConfig: Clay__StoreImageElementConfig(declaration.Image)}, CLAY__ELEMENT_CONFIG_TYPE_IMAGE)
		Clay__int32_tArray_Add(&context.ImageElementPointers, int32(int64(context.LayoutElements.Length)-1))
	}
	if declaration.Floating.AttachTo != CLAY_ATTACH_TO_NONE {
		var (
			floatingConfig     Clay_FloatingElementConfig = declaration.Floating
			hierarchicalParent *Clay_LayoutElement        = Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&context.OpenLayoutElementStack, int32(int64(context.OpenLayoutElementStack.Length)-2)))
		)
		if hierarchicalParent != nil {
			var clipElementId uint32 = 0
			if declaration.Floating.AttachTo == CLAY_ATTACH_TO_PARENT {
				floatingConfig.ParentId = hierarchicalParent.Id
				if int64(context.OpenClipElementStack.Length) > 0 {
					clipElementId = uint32(Clay__int32_tArray_GetValue(&context.OpenClipElementStack, int32(int64(context.OpenClipElementStack.Length)-1)))
				}
			} else if declaration.Floating.AttachTo == CLAY_ATTACH_TO_ELEMENT_WITH_ID {
				var parentItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(floatingConfig.ParentId)
				if parentItem == nil {
					context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_FLOATING_CONTAINER_PARENT_NOT_FOUND, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("A floating element was declared with a parentId, but no element with that ID was found.")}, UserData: context.ErrorHandler.UserData})
				} else {
					clipElementId = uint32(Clay__int32_tArray_GetValue(&context.LayoutElementClipElementIds, int32(int64(uintptr(unsafe.Pointer(parentItem.LayoutElement))-uintptr(unsafe.Pointer(context.LayoutElements.InternalArray))))))
				}
			} else if declaration.Floating.AttachTo == CLAY_ATTACH_TO_ROOT {
				floatingConfig.ParentId = Clay__HashString(Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay__RootContainer")}, 0, 0).Id
			}
			if int64(openLayoutElementId.Id) == 0 {
				openLayoutElementId = Clay__HashString(Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay__FloatingContainer")}, uint32(context.LayoutElementTreeRoots.Length), 0)
			}
			Clay__LayoutElementTreeRootArray_Add(&context.LayoutElementTreeRoots, Clay__LayoutElementTreeRoot{LayoutElementIndex: Clay__int32_tArray_GetValue(&context.OpenLayoutElementStack, int32(int64(context.OpenLayoutElementStack.Length)-1)), ParentId: floatingConfig.ParentId, ClipElementId: clipElementId, ZIndex: floatingConfig.ZIndex})
			Clay__AttachElementConfig(Clay_ElementConfigUnion{FloatingElementConfig: Clay__StoreFloatingElementConfig(declaration.Floating)}, CLAY__ELEMENT_CONFIG_TYPE_FLOATING)
		}
	}
	if declaration.Custom.CustomData != nil {
		Clay__AttachElementConfig(Clay_ElementConfigUnion{CustomElementConfig: Clay__StoreCustomElementConfig(declaration.Custom)}, CLAY__ELEMENT_CONFIG_TYPE_CUSTOM)
	}
	if int64(openLayoutElementId.Id) != 0 {
		Clay__AttachId(openLayoutElementId)
	} else if int64(openLayoutElement.Id) == 0 {
		openLayoutElementId = Clay__GenerateIdForAnonymousElement(openLayoutElement)
	}
	if declaration.Scroll.Horizontal || declaration.Scroll.Vertical {
		Clay__AttachElementConfig(Clay_ElementConfigUnion{ScrollElementConfig: Clay__StoreScrollElementConfig(declaration.Scroll)}, CLAY__ELEMENT_CONFIG_TYPE_SCROLL)
		Clay__int32_tArray_Add(&context.OpenClipElementStack, int32(int64(openLayoutElement.Id)))
		var scrollOffset *Clay__ScrollContainerDataInternal = (*Clay__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(CLAY__NULL)))
		for i := int32(0); int64(i) < int64(context.ScrollContainerDatas.Length); i++ {
			var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
			if int64(openLayoutElement.Id) == int64(mapping.ElementId) {
				scrollOffset = mapping
				scrollOffset.LayoutElement = openLayoutElement
				scrollOffset.OpenThisFrame = true
			}
		}
		if scrollOffset == nil {
			scrollOffset = Clay__ScrollContainerDataInternalArray_Add(&context.ScrollContainerDatas, Clay__ScrollContainerDataInternal{LayoutElement: openLayoutElement, ScrollOrigin: Clay_Vector2{X: -1, Y: -1}, ElementId: openLayoutElement.Id, OpenThisFrame: true})
		}
		if context.ExternalScrollHandlingEnabled {
			scrollOffset.ScrollPosition = Clay__QueryScrollOffset(scrollOffset.ElementId, context.QueryScrollOffsetUserData)
		}
	}
	if !Clay__MemCmp((*byte)(unsafe.Pointer(&declaration.Border.Width)), (*byte)(unsafe.Pointer(&Clay__BorderWidth_DEFAULT)), int32(uint32(unsafe.Sizeof(Clay_BorderWidth{})))) {
		Clay__AttachElementConfig(Clay_ElementConfigUnion{BorderElementConfig: Clay__StoreBorderElementConfig(declaration.Border)}, CLAY__ELEMENT_CONFIG_TYPE_BORDER)
	}
}
func Clay__InitializeEphemeralMemory(context *Clay_Context) {
	var (
		maxElementCount int32       = context.MaxElementCount
		arena           *Clay_Arena = &context.InternalArena
	)
	arena.NextAllocation = context.ArenaResetOffset
	context.LayoutElementChildrenBuffer = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElements = Clay_LayoutElementArray_Allocate_Arena(maxElementCount, arena)
	context.Warnings = Clay__WarningArray_Allocate_Arena(100, arena)
	context.LayoutConfigs = Clay__LayoutConfigArray_Allocate_Arena(maxElementCount, arena)
	context.ElementConfigs = Clay__ElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.TextElementConfigs = Clay__TextElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.ImageElementConfigs = Clay__ImageElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.FloatingElementConfigs = Clay__FloatingElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.ScrollElementConfigs = Clay__ScrollElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.CustomElementConfigs = Clay__CustomElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.BorderElementConfigs = Clay__BorderElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.SharedElementConfigs = Clay__SharedElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementIdStrings = Clay__StringArray_Allocate_Arena(maxElementCount, arena)
	context.WrappedTextLines = Clay__WrappedTextLineArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementTreeNodeArray1 = Clay__LayoutElementTreeNodeArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementTreeRoots = Clay__LayoutElementTreeRootArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementChildren = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.OpenLayoutElementStack = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.TextElementData = Clay__TextElementDataArray_Allocate_Arena(maxElementCount, arena)
	context.ImageElementPointers = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.RenderCommands = Clay_RenderCommandArray_Allocate_Arena(maxElementCount, arena)
	context.TreeNodeVisited = Clay__boolArray_Allocate_Arena(maxElementCount, arena)
	context.TreeNodeVisited.Length = context.TreeNodeVisited.Capacity
	context.OpenClipElementStack = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.ReusableElementIndexBuffer = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementClipElementIds = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.DynamicStringData = Clay__charArray_Allocate_Arena(maxElementCount, arena)
}
func Clay__InitializePersistentMemory(context *Clay_Context) {
	var (
		maxElementCount              int32       = context.MaxElementCount
		maxMeasureTextCacheWordCount int32       = context.MaxMeasureTextCacheWordCount
		arena                        *Clay_Arena = &context.InternalArena
	)
	context.ScrollContainerDatas = Clay__ScrollContainerDataInternalArray_Allocate_Arena(10, arena)
	context.LayoutElementsHashMapInternal = Clay__LayoutElementHashMapItemArray_Allocate_Arena(maxElementCount, arena)
	context.LayoutElementsHashMap = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.MeasureTextHashMapInternal = Clay__MeasureTextCacheItemArray_Allocate_Arena(maxElementCount, arena)
	context.MeasureTextHashMapInternalFreeList = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.MeasuredWordsFreeList = Clay__int32_tArray_Allocate_Arena(maxMeasureTextCacheWordCount, arena)
	context.MeasureTextHashMap = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.MeasuredWords = Clay__MeasuredWordArray_Allocate_Arena(maxMeasureTextCacheWordCount, arena)
	context.PointerOverIds = Clay__ElementIdArray_Allocate_Arena(maxElementCount, arena)
	context.DebugElementData = Clay__DebugElementDataArray_Allocate_Arena(maxElementCount, arena)
	context.ArenaResetOffset = arena.NextAllocation
}
func Clay__CompressChildrenAlongAxis(xAxis bool, totalSizeToDistribute float32, resizableContainerBuffer Clay__int32_tArray) {
	var (
		context           *Clay_Context      = Clay_GetCurrentContext()
		largestContainers Clay__int32_tArray = context.OpenClipElementStack
	)
	for totalSizeToDistribute > 0.1 {
		largestContainers.Length = 0
		var largestSize float32 = 0
		var targetSize float32 = 0
		for i := int32(0); int64(i) < int64(resizableContainerBuffer.Length); i++ {
			var (
				childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&resizableContainerBuffer, i))
				childSize    float32
			)
			if xAxis {
				childSize = childElement.Dimensions.Width
			} else {
				childSize = childElement.Dimensions.Height
			}
			if (childSize-largestSize) < 0.1 && (childSize-largestSize) > -0.1 {
				Clay__int32_tArray_Add(&largestContainers, Clay__int32_tArray_GetValue(&resizableContainerBuffer, i))
			} else if childSize > largestSize {
				targetSize = largestSize
				largestSize = childSize
				largestContainers.Length = 0
				Clay__int32_tArray_Add(&largestContainers, Clay__int32_tArray_GetValue(&resizableContainerBuffer, i))
			} else if childSize > targetSize {
				targetSize = childSize
			}
		}
		if int64(largestContainers.Length) == 0 {
			return
		}
		targetSize = (func() float32 {
			if targetSize > ((largestSize * float32(largestContainers.Length)) - totalSizeToDistribute) {
				return targetSize
			}
			return (largestSize * float32(largestContainers.Length)) - totalSizeToDistribute
		}()) / float32(largestContainers.Length)
		for childOffset := int32(0); int64(childOffset) < int64(largestContainers.Length); childOffset++ {
			var (
				childIndex   int32               = Clay__int32_tArray_GetValue(&largestContainers, childOffset)
				childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, childIndex)
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
				for i := int32(0); int64(i) < int64(resizableContainerBuffer.Length); i++ {
					if int64(Clay__int32_tArray_GetValue(&resizableContainerBuffer, i)) == int64(childIndex) {
						Clay__int32_tArray_RemoveSwapback(&resizableContainerBuffer, i)
						break
					}
				}
			}
		}
	}
}
func Clay__SizeContainersAlongAxis(xAxis bool) {
	var (
		context                  *Clay_Context      = Clay_GetCurrentContext()
		bfsBuffer                Clay__int32_tArray = context.LayoutElementChildrenBuffer
		resizableContainerBuffer Clay__int32_tArray = context.OpenLayoutElementStack
	)
	for rootIndex := int32(0); int64(rootIndex) < int64(context.LayoutElementTreeRoots.Length); rootIndex++ {
		bfsBuffer.Length = 0
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, rootIndex)
		var rootElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, int32(int64(root.LayoutElementIndex)))
		Clay__int32_tArray_Add(&bfsBuffer, root.LayoutElementIndex)
		if Clay__ElementHasConfig(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING) {
			var (
				floatingElementConfig *Clay_FloatingElementConfig    = Clay__FindElementConfigWithType(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig
				parentItem            *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(floatingElementConfig.ParentId)
			)
			if parentItem != nil && parentItem != &Clay_LayoutElementHashMapItem_DEFAULT {
				var parentLayoutElement *Clay_LayoutElement = parentItem.LayoutElement
				if rootElement.LayoutConfig.Sizing.Width.Type == CLAY__SIZING_TYPE_GROW {
					rootElement.Dimensions.Width = parentLayoutElement.Dimensions.Width
				}
				if rootElement.LayoutConfig.Sizing.Height.Type == CLAY__SIZING_TYPE_GROW {
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
		for i := int32(0); int64(i) < int64(bfsBuffer.Length); i++ {
			var (
				parentIndex        int32               = Clay__int32_tArray_GetValue(&bfsBuffer, i)
				parent             *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, parentIndex)
				parentStyleConfig  *Clay_LayoutConfig  = parent.LayoutConfig
				growContainerCount int32               = 0
				parentSize         float32
			)
			if xAxis {
				parentSize = parent.Dimensions.Width
			} else {
				parentSize = parent.Dimensions.Height
			}
			var parentPadding float32 = float32(func() int64 {
				if xAxis {
					return int64(parent.LayoutConfig.Padding.Left) + int64(parent.LayoutConfig.Padding.Right)
				}
				return int64(parent.LayoutConfig.Padding.Top) + int64(parent.LayoutConfig.Padding.Bottom)
			}())
			var innerContentSize float32 = 0
			var growContainerContentSize float32 = 0
			var totalPaddingAndChildGaps float32 = parentPadding
			var sizingAlongAxis bool = xAxis && parentStyleConfig.LayoutDirection == CLAY_LEFT_TO_RIGHT || !xAxis && parentStyleConfig.LayoutDirection == CLAY_TOP_TO_BOTTOM
			resizableContainerBuffer.Length = 0
			var parentChildGap float32 = float32(parentStyleConfig.ChildGap)
			for childOffset := int32(0); int64(childOffset) < int64(parent.ChildrenOrTextContent.Children.Length); childOffset++ {
				var (
					childElementIndex int32               = *(*int32)(unsafe.Add(unsafe.Pointer(parent.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(childOffset)))
					childElement      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, childElementIndex)
					childSizing       Clay_SizingAxis
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
				if !Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) && int64(childElement.ChildrenOrTextContent.Children.Length) > 0 {
					Clay__int32_tArray_Add(&bfsBuffer, childElementIndex)
				}
				if childSizing.Type != CLAY__SIZING_TYPE_PERCENT && childSizing.Type != CLAY__SIZING_TYPE_FIXED && (!Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) || Clay__FindElementConfigWithType(childElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT).TextElementConfig.WrapMode == CLAY_TEXT_WRAP_WORDS) && (xAxis || !Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_IMAGE)) {
					Clay__int32_tArray_Add(&resizableContainerBuffer, childElementIndex)
				}
				if sizingAlongAxis {
					if childSizing.Type == CLAY__SIZING_TYPE_PERCENT {
						innerContentSize += 0
					} else {
						innerContentSize += childSize
					}
					if childSizing.Type == CLAY__SIZING_TYPE_GROW {
						growContainerContentSize += childSize
						growContainerCount++
					}
					if int64(childOffset) > 0 {
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
			for childOffset := int32(0); int64(childOffset) < int64(parent.ChildrenOrTextContent.Children.Length); childOffset++ {
				var (
					childElementIndex int32               = *(*int32)(unsafe.Add(unsafe.Pointer(parent.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(childOffset)))
					childElement      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, childElementIndex)
					childSizing       Clay_SizingAxis
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
				if childSizing.Type == CLAY__SIZING_TYPE_PERCENT {
					*childSize = (parentSize - totalPaddingAndChildGaps) * childSizing.Size.Percent
					if sizingAlongAxis {
						innerContentSize += *childSize
					}
				}
			}
			if sizingAlongAxis {
				var sizeToDistribute float32 = parentSize - parentPadding - innerContentSize
				if sizeToDistribute < 0 {
					var scrollElementConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(parent, CLAY__ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
					if scrollElementConfig != nil {
						if xAxis && scrollElementConfig.Horizontal || !xAxis && scrollElementConfig.Vertical {
							continue
						}
					}
					Clay__CompressChildrenAlongAxis(xAxis, -sizeToDistribute, resizableContainerBuffer)
				} else if sizeToDistribute > 0 && int64(growContainerCount) > 0 {
					var targetSize float32 = (sizeToDistribute + growContainerContentSize) / float32(growContainerCount)
					for childOffset := int32(0); int64(childOffset) < int64(resizableContainerBuffer.Length); childOffset++ {
						var (
							childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&resizableContainerBuffer, childOffset))
							childSizing  Clay_SizingAxis
						)
						if xAxis {
							childSizing = childElement.LayoutConfig.Sizing.Width
						} else {
							childSizing = childElement.LayoutConfig.Sizing.Height
						}
						if childSizing.Type == CLAY__SIZING_TYPE_GROW {
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
								Clay__int32_tArray_RemoveSwapback(&resizableContainerBuffer, childOffset)
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
				for childOffset := int32(0); int64(childOffset) < int64(resizableContainerBuffer.Length); childOffset++ {
					var (
						childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&resizableContainerBuffer, childOffset))
						childSizing  Clay_SizingAxis
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
					if !xAxis && Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_IMAGE) {
						continue
					}
					var maxSize float32 = parentSize - parentPadding
					if Clay__ElementHasConfig(parent, CLAY__ELEMENT_CONFIG_TYPE_SCROLL) {
						var scrollElementConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(parent, CLAY__ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
						if xAxis && scrollElementConfig.Horizontal || !xAxis && scrollElementConfig.Vertical {
							if maxSize > innerContentSize {
								maxSize = maxSize
							} else {
								maxSize = innerContentSize
							}
						}
					}
					if childSizing.Type == CLAY__SIZING_TYPE_FIT {
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
					} else if childSizing.Type == CLAY__SIZING_TYPE_GROW {
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
func Clay__IntToString(integer int32) Clay_String {
	if int64(integer) == 0 {
		return Clay_String{Length: 1, Chars: libc.CString("0")}
	}
	var context *Clay_Context = Clay_GetCurrentContext()
	var chars *byte = ((*byte)(unsafe.Add(unsafe.Pointer(context.DynamicStringData.InternalArray), context.DynamicStringData.Length)))
	var length int32 = 0
	var sign int32 = integer
	if int64(integer) < 0 {
		integer = -integer
	}
	for int64(integer) > 0 {
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), func() int32 {
			p := &length
			x := *p
			*p++
			return x
		}())) = byte(int8(int64(integer)%10 + '0'))
		integer /= 10
	}
	if int64(sign) < 0 {
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), func() int32 {
			p := &length
			x := *p
			*p++
			return x
		}())) = '-'
	}
	for j, k := int32(0), int32(int32(int64(length)-1)); int64(j) < int64(k); func() int32 {
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
	return Clay_String{Length: length, Chars: chars}
}
func Clay__AddRenderCommand(renderCommand Clay_RenderCommand) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int64(context.RenderCommands.Length) < int64(context.RenderCommands.Capacity)-1 {
		Clay_RenderCommandArray_Add(&context.RenderCommands, renderCommand)
	} else {
		if !context.BooleanWarnings.MaxRenderCommandsExceeded {
			context.BooleanWarnings.MaxRenderCommandsExceeded = true
			context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay ran out of capacity while attempting to create render commands. This is usually caused by a large amount of wrapping text elements while close to the max element capacity. Try using Clay_SetMaxElementCount() with a higher value.")}, UserData: context.ErrorHandler.UserData})
		}
	}
}
func Clay__ElementIsOffscreen(boundingBox *Clay_BoundingBox) bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.DisableCulling {
		return false
	}
	return boundingBox.X > context.LayoutDimensions.Width || boundingBox.Y > context.LayoutDimensions.Height || boundingBox.X+boundingBox.Width < 0 || boundingBox.Y+boundingBox.Height < 0
}
func Clay__CalculateFinalLayout() {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__SizeContainersAlongAxis(true)
	for textElementIndex := int32(0); int64(textElementIndex) < int64(context.TextElementData.Length); textElementIndex++ {
		var textElementData *Clay__TextElementData = Clay__TextElementDataArray_Get(&context.TextElementData, textElementIndex)
		textElementData.WrappedLines = Clay__WrappedTextLineArraySlice{Length: 0, InternalArray: (*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(context.WrappedTextLines.InternalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(context.WrappedTextLines.Length)))}
		var containerElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, int32(int64(textElementData.ElementIndex)))
		var textConfig *Clay_TextElementConfig = Clay__FindElementConfigWithType(containerElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT).TextElementConfig
		var measureTextCacheItem *Clay__MeasureTextCacheItem = Clay__MeasureTextCached(&textElementData.Text, textConfig)
		var lineWidth float32 = 0
		var lineHeight float32
		if int64(textConfig.LineHeight) > 0 {
			lineHeight = float32(textConfig.LineHeight)
		} else {
			lineHeight = textElementData.PreferredDimensions.Height
		}
		var lineLengthChars int32 = 0
		var lineStartOffset int32 = 0
		if !measureTextCacheItem.ContainsNewlines && textElementData.PreferredDimensions.Width <= containerElement.Dimensions.Width {
			Clay__WrappedTextLineArray_Add(&context.WrappedTextLines, Clay__WrappedTextLine{Dimensions: containerElement.Dimensions, Line: textElementData.Text})
			textElementData.WrappedLines.Length++
			continue
		}
		var spaceWidth float32 = Clay__MeasureText(Clay_StringSlice{Length: 1, Chars: CLAY__SPACECHAR.Chars, BaseChars: CLAY__SPACECHAR.Chars}, textConfig, context.MeasureTextUserData).Width
		_ = spaceWidth
		var wordIndex int32 = measureTextCacheItem.MeasuredWordsStartIndex
		for int64(wordIndex) != -1 {
			if int64(context.WrappedTextLines.Length) > int64(context.WrappedTextLines.Capacity)-1 {
				break
			}
			var measuredWord *Clay__MeasuredWord = Clay__MeasuredWordArray_Get(&context.MeasuredWords, wordIndex)
			if int64(lineLengthChars) == 0 && lineWidth+measuredWord.Width > containerElement.Dimensions.Width {
				Clay__WrappedTextLineArray_Add(&context.WrappedTextLines, Clay__WrappedTextLine{Dimensions: Clay_Dimensions{Width: measuredWord.Width, Height: lineHeight}, Line: Clay_String{Length: measuredWord.Length, Chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), measuredWord.StartOffset))}})
				textElementData.WrappedLines.Length++
				wordIndex = measuredWord.Next
				lineStartOffset = int32(int64(measuredWord.StartOffset) + int64(measuredWord.Length))
			} else if int64(measuredWord.Length) == 0 || lineWidth+measuredWord.Width > containerElement.Dimensions.Width {
				var finalCharIsSpace bool = *(*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), int64(lineStartOffset)+int64(lineLengthChars)-1)) == ' '
				Clay__WrappedTextLineArray_Add(&context.WrappedTextLines, Clay__WrappedTextLine{Dimensions: Clay_Dimensions{Width: lineWidth + (func() float32 {
					if finalCharIsSpace {
						return -spaceWidth
					}
					return 0
				}()), Height: lineHeight}, Line: Clay_String{Length: int32(int64(lineLengthChars) + (func() int64 {
					if finalCharIsSpace {
						return -1
					}
					return 0
				}())), Chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), lineStartOffset))}})
				textElementData.WrappedLines.Length++
				if int64(lineLengthChars) == 0 || int64(measuredWord.Length) == 0 {
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
		if int64(lineLengthChars) > 0 {
			Clay__WrappedTextLineArray_Add(&context.WrappedTextLines, Clay__WrappedTextLine{Dimensions: Clay_Dimensions{Width: lineWidth, Height: lineHeight}, Line: Clay_String{Length: lineLengthChars, Chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.Text.Chars), lineStartOffset))}})
			textElementData.WrappedLines.Length++
		}
		containerElement.Dimensions.Height = lineHeight * float32(textElementData.WrappedLines.Length)
	}
	for i := int32(0); int64(i) < int64(context.ImageElementPointers.Length); i++ {
		var (
			imageElement *Clay_LayoutElement      = Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&context.ImageElementPointers, i))
			config       *Clay_ImageElementConfig = Clay__FindElementConfigWithType(imageElement, CLAY__ELEMENT_CONFIG_TYPE_IMAGE).ImageElementConfig
		)
		imageElement.Dimensions.Height = (config.SourceDimensions.Height / (func() float32 {
			if config.SourceDimensions.Width > 1 {
				return config.SourceDimensions.Width
			}
			return 1
		}())) * imageElement.Dimensions.Width
	}
	var dfsBuffer Clay__LayoutElementTreeNodeArray = context.LayoutElementTreeNodeArray1
	dfsBuffer.Length = 0
	for i := int32(0); int64(i) < int64(context.LayoutElementTreeRoots.Length); i++ {
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, i)
		*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length)) = false
		Clay__LayoutElementTreeNodeArray_Add(&dfsBuffer, Clay__LayoutElementTreeNode{LayoutElement: Clay_LayoutElementArray_Get(&context.LayoutElements, int32(int64(root.LayoutElementIndex)))})
	}
	for int64(dfsBuffer.Length) > 0 {
		var (
			currentElementTreeNode *Clay__LayoutElementTreeNode = Clay__LayoutElementTreeNodeArray_Get(&dfsBuffer, int32(int64(dfsBuffer.Length)-1))
			currentElement         *Clay_LayoutElement          = currentElementTreeNode.LayoutElement
		)
		if !*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), int64(dfsBuffer.Length)-1)) {
			*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), int64(dfsBuffer.Length)-1)) = true
			if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) || int64(currentElement.ChildrenOrTextContent.Children.Length) == 0 {
				dfsBuffer.Length--
				continue
			}
			for i := int32(0); int64(i) < int64(currentElement.ChildrenOrTextContent.Children.Length); i++ {
				*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), dfsBuffer.Length)) = false
				Clay__LayoutElementTreeNodeArray_Add(&dfsBuffer, Clay__LayoutElementTreeNode{LayoutElement: Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))})
			}
			continue
		}
		dfsBuffer.Length--
		var layoutConfig *Clay_LayoutConfig = currentElement.LayoutConfig
		if layoutConfig.LayoutDirection == CLAY_LEFT_TO_RIGHT {
			for j := int32(0); int64(j) < int64(currentElement.ChildrenOrTextContent.Children.Length); j++ {
				var (
					childElement           *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(j))))
					childHeightWithPadding float32             = (func() float32 {
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
		} else if layoutConfig.LayoutDirection == CLAY_TOP_TO_BOTTOM {
			var contentHeight float32 = float32(int64(layoutConfig.Padding.Top) + int64(layoutConfig.Padding.Bottom))
			for j := int32(0); int64(j) < int64(currentElement.ChildrenOrTextContent.Children.Length); j++ {
				var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(j))))
				contentHeight += childElement.Dimensions.Height
			}
			contentHeight += float32((func() int64 {
				if (int64(currentElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
					return int64(currentElement.ChildrenOrTextContent.Children.Length) - 1
				}
				return 0
			}()) * int64(layoutConfig.ChildGap))
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
	Clay__SizeContainersAlongAxis(false)
	var sortMax int32 = int32(int64(context.LayoutElementTreeRoots.Length) - 1)
	for int64(sortMax) > 0 {
		for i := int32(0); int64(i) < int64(sortMax); i++ {
			var (
				current Clay__LayoutElementTreeRoot = *Clay__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, i)
				next    Clay__LayoutElementTreeRoot = *Clay__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, int32(int64(i)+1))
			)
			if int64(next.ZIndex) < int64(current.ZIndex) {
				Clay__LayoutElementTreeRootArray_Set(&context.LayoutElementTreeRoots, i, next)
				Clay__LayoutElementTreeRootArray_Set(&context.LayoutElementTreeRoots, int32(int64(i)+1), current)
			}
		}
		sortMax--
	}
	context.RenderCommands.Length = 0
	dfsBuffer.Length = 0
	for rootIndex := int32(0); int64(rootIndex) < int64(context.LayoutElementTreeRoots.Length); rootIndex++ {
		dfsBuffer.Length = 0
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, rootIndex)
		var rootElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, int32(int64(root.LayoutElementIndex)))
		var rootPosition Clay_Vector2 = Clay_Vector2{}
		var parentHashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(root.ParentId)
		if Clay__ElementHasConfig(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING) && parentHashMapItem != nil {
			var (
				config               *Clay_FloatingElementConfig = Clay__FindElementConfigWithType(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig
				rootDimensions       Clay_Dimensions             = rootElement.Dimensions
				parentBoundingBox    Clay_BoundingBox            = parentHashMapItem.BoundingBox
				targetAttachPosition Clay_Vector2                = Clay_Vector2{}
			)
			switch config.AttachPoints.Parent {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
				targetAttachPosition.X = parentBoundingBox.X
			case CLAY_ATTACH_POINT_CENTER_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				targetAttachPosition.X = parentBoundingBox.X + parentBoundingBox.Width/2
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.X = parentBoundingBox.X + parentBoundingBox.Width
			}
			switch config.AttachPoints.Element {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
			case CLAY_ATTACH_POINT_CENTER_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				targetAttachPosition.X -= rootDimensions.Width / 2
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.X -= rootDimensions.Width
			}
			switch config.AttachPoints.Parent {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_TOP:
				targetAttachPosition.Y = parentBoundingBox.Y
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				targetAttachPosition.Y = parentBoundingBox.Y + parentBoundingBox.Height/2
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.Y = parentBoundingBox.Y + parentBoundingBox.Height
			}
			switch config.AttachPoints.Element {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_TOP:
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				targetAttachPosition.Y -= rootDimensions.Height / 2
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.Y -= rootDimensions.Height
			}
			targetAttachPosition.X += config.Offset.X
			targetAttachPosition.Y += config.Offset.Y
			rootPosition = targetAttachPosition
		}
		if int64(root.ClipElementId) != 0 {
			var clipHashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(root.ClipElementId)
			if clipHashMapItem != nil {
				if context.ExternalScrollHandlingEnabled {
					var scrollConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(clipHashMapItem.LayoutElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
					for i := int32(0); int64(i) < int64(context.ScrollContainerDatas.Length); i++ {
						var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
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
				Clay__AddRenderCommand(Clay_RenderCommand{BoundingBox: clipHashMapItem.BoundingBox, UserData: nil, Id: Clay__HashNumber(rootElement.Id, uint32(int32(int64(rootElement.ChildrenOrTextContent.Children.Length)+10))).Id, ZIndex: root.ZIndex, CommandType: CLAY_RENDER_COMMAND_TYPE_SCISSOR_START})
			}
		}
		Clay__LayoutElementTreeNodeArray_Add(&dfsBuffer, Clay__LayoutElementTreeNode{LayoutElement: rootElement, Position: rootPosition, NextChildOffset: Clay_Vector2{X: float32(rootElement.LayoutConfig.Padding.Left), Y: float32(rootElement.LayoutConfig.Padding.Top)}})
		*context.TreeNodeVisited.InternalArray = false
		for int64(dfsBuffer.Length) > 0 {
			var (
				currentElementTreeNode *Clay__LayoutElementTreeNode = Clay__LayoutElementTreeNodeArray_Get(&dfsBuffer, int32(int64(dfsBuffer.Length)-1))
				currentElement         *Clay_LayoutElement          = currentElementTreeNode.LayoutElement
				layoutConfig           *Clay_LayoutConfig           = currentElement.LayoutConfig
				scrollOffset           Clay_Vector2                 = Clay_Vector2{}
			)
			if !*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), int64(dfsBuffer.Length)-1)) {
				*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), int64(dfsBuffer.Length)-1)) = true
				var currentElementBoundingBox Clay_BoundingBox = Clay_BoundingBox{X: currentElementTreeNode.Position.X, Y: currentElementTreeNode.Position.Y, Width: currentElement.Dimensions.Width, Height: currentElement.Dimensions.Height}
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING) {
					var (
						floatingElementConfig *Clay_FloatingElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig
						expand                Clay_Dimensions             = floatingElementConfig.Expand
					)
					currentElementBoundingBox.X -= expand.Width
					currentElementBoundingBox.Width += expand.Width * 2
					currentElementBoundingBox.Y -= expand.Height
					currentElementBoundingBox.Height += expand.Height * 2
				}
				var scrollContainerData *Clay__ScrollContainerDataInternal = (*Clay__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(CLAY__NULL)))
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL) {
					var scrollConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
					for i := int32(0); int64(i) < int64(context.ScrollContainerDatas.Length); i++ {
						var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
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
								scrollOffset = Clay_Vector2{}
							}
							break
						}
					}
				}
				var hashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(currentElement.Id)
				if hashMapItem != nil {
					hashMapItem.BoundingBox = currentElementBoundingBox
					if int64(hashMapItem.IdAlias) != 0 {
						var hashMapItemAlias *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(hashMapItem.IdAlias)
						if hashMapItemAlias != nil {
							hashMapItemAlias.BoundingBox = currentElementBoundingBox
						}
					}
				}
				var sortedConfigIndexes [20]int32
				for elementConfigIndex := int32(0); int64(elementConfigIndex) < int64(currentElement.ElementConfigs.Length); elementConfigIndex++ {
					sortedConfigIndexes[elementConfigIndex] = elementConfigIndex
				}
				sortMax = int32(int64(currentElement.ElementConfigs.Length) - 1)
				for int64(sortMax) > 0 {
					for i := int32(0); int64(i) < int64(sortMax); i++ {
						var (
							current     int32                   = sortedConfigIndexes[i]
							next        int32                   = sortedConfigIndexes[int64(i)+1]
							currentType Clay__ElementConfigType = Clay__ElementConfigArraySlice_Get(&currentElement.ElementConfigs, current).Type
							nextType    Clay__ElementConfigType = Clay__ElementConfigArraySlice_Get(&currentElement.ElementConfigs, next).Type
						)
						if nextType == CLAY__ELEMENT_CONFIG_TYPE_SCROLL || currentType == CLAY__ELEMENT_CONFIG_TYPE_BORDER {
							sortedConfigIndexes[i] = next
							sortedConfigIndexes[int64(i)+1] = current
						}
					}
					sortMax--
				}
				var emitRectangle bool = false
				var sharedConfig *Clay_SharedElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SHARED).SharedElementConfig
				if sharedConfig != nil && sharedConfig.BackgroundColor.A > 0 {
					emitRectangle = true
				} else if sharedConfig == nil {
					emitRectangle = false
					sharedConfig = &Clay_SharedElementConfig_DEFAULT
				}
				for elementConfigIndex := int32(0); int64(elementConfigIndex) < int64(currentElement.ElementConfigs.Length); elementConfigIndex++ {
					var (
						elementConfig *Clay_ElementConfig = Clay__ElementConfigArraySlice_Get(&currentElement.ElementConfigs, sortedConfigIndexes[elementConfigIndex])
						renderCommand Clay_RenderCommand  = Clay_RenderCommand{BoundingBox: currentElementBoundingBox, UserData: sharedConfig.UserData, Id: currentElement.Id}
						offscreen     bool                = Clay__ElementIsOffscreen(&currentElementBoundingBox)
						shouldRender  bool                = !offscreen
					)
					switch elementConfig.Type {
					case CLAY__ELEMENT_CONFIG_TYPE_FLOATING:
						fallthrough
					case CLAY__ELEMENT_CONFIG_TYPE_SHARED:
						fallthrough
					case CLAY__ELEMENT_CONFIG_TYPE_BORDER:
						shouldRender = false
					case CLAY__ELEMENT_CONFIG_TYPE_SCROLL:
						renderCommand.CommandType = CLAY_RENDER_COMMAND_TYPE_SCISSOR_START
						renderCommand.RenderData = Clay_RenderData{Scroll: Clay_ScrollRenderData{Horizontal: elementConfig.Config.ScrollElementConfig.Horizontal, Vertical: elementConfig.Config.ScrollElementConfig.Vertical}}
					case CLAY__ELEMENT_CONFIG_TYPE_IMAGE:
						renderCommand.CommandType = CLAY_RENDER_COMMAND_TYPE_IMAGE
						renderCommand.RenderData = Clay_RenderData{Image: Clay_ImageRenderData{BackgroundColor: sharedConfig.BackgroundColor, CornerRadius: sharedConfig.CornerRadius, SourceDimensions: elementConfig.Config.ImageElementConfig.SourceDimensions, ImageData: elementConfig.Config.ImageElementConfig.ImageData}}
						emitRectangle = false
					case CLAY__ELEMENT_CONFIG_TYPE_TEXT:
						if !shouldRender {
							break
						}
						shouldRender = false
						var configUnion Clay_ElementConfigUnion = elementConfig.Config
						var textElementConfig *Clay_TextElementConfig = configUnion.TextElementConfig
						var naturalLineHeight float32 = currentElement.ChildrenOrTextContent.TextElementData.PreferredDimensions.Height
						var finalLineHeight float32
						if int64(textElementConfig.LineHeight) > 0 {
							finalLineHeight = float32(textElementConfig.LineHeight)
						} else {
							finalLineHeight = naturalLineHeight
						}
						var lineHeightOffset float32 = (finalLineHeight - naturalLineHeight) / 2
						var yPosition float32 = lineHeightOffset
						for lineIndex := int32(0); int64(lineIndex) < int64(currentElement.ChildrenOrTextContent.TextElementData.WrappedLines.Length); lineIndex++ {
							var wrappedLine *Clay__WrappedTextLine = Clay__WrappedTextLineArraySlice_Get(&currentElement.ChildrenOrTextContent.TextElementData.WrappedLines, lineIndex)
							if int64(wrappedLine.Line.Length) == 0 {
								yPosition += finalLineHeight
								continue
							}
							var offset float32 = (currentElementBoundingBox.Width - wrappedLine.Dimensions.Width)
							if textElementConfig.TextAlignment == CLAY_TEXT_ALIGN_LEFT {
								offset = 0
							}
							if textElementConfig.TextAlignment == CLAY_TEXT_ALIGN_CENTER {
								offset /= 2
							}
							Clay__AddRenderCommand(Clay_RenderCommand{BoundingBox: Clay_BoundingBox{X: currentElementBoundingBox.X + offset, Y: currentElementBoundingBox.Y + yPosition, Width: wrappedLine.Dimensions.Width, Height: wrappedLine.Dimensions.Height}, RenderData: Clay_RenderData{Text: Clay_TextRenderData{StringContents: Clay_StringSlice{Length: wrappedLine.Line.Length, Chars: wrappedLine.Line.Chars, BaseChars: currentElement.ChildrenOrTextContent.TextElementData.Text.Chars}, TextColor: textElementConfig.TextColor, FontId: textElementConfig.FontId, FontSize: textElementConfig.FontSize, LetterSpacing: textElementConfig.LetterSpacing, LineHeight: textElementConfig.LineHeight}}, UserData: sharedConfig.UserData, Id: Clay__HashNumber(uint32(lineIndex), currentElement.Id).Id, ZIndex: root.ZIndex, CommandType: CLAY_RENDER_COMMAND_TYPE_TEXT})
							yPosition += finalLineHeight
							if !context.DisableCulling && currentElementBoundingBox.Y+yPosition > context.LayoutDimensions.Height {
								break
							}
						}
					case CLAY__ELEMENT_CONFIG_TYPE_CUSTOM:
						renderCommand.CommandType = CLAY_RENDER_COMMAND_TYPE_CUSTOM
						renderCommand.RenderData = Clay_RenderData{Custom: Clay_CustomRenderData{BackgroundColor: sharedConfig.BackgroundColor, CornerRadius: sharedConfig.CornerRadius, CustomData: elementConfig.Config.CustomElementConfig.CustomData}}
						emitRectangle = false
					default:
					}
					if shouldRender {
						Clay__AddRenderCommand(renderCommand)
					}
					if offscreen {
					}
				}
				if emitRectangle {
					Clay__AddRenderCommand(Clay_RenderCommand{BoundingBox: currentElementBoundingBox, RenderData: Clay_RenderData{Rectangle: Clay_RectangleRenderData{BackgroundColor: sharedConfig.BackgroundColor, CornerRadius: sharedConfig.CornerRadius}}, UserData: sharedConfig.UserData, Id: currentElement.Id, ZIndex: root.ZIndex, CommandType: CLAY_RENDER_COMMAND_TYPE_RECTANGLE})
				}
				if !Clay__ElementHasConfig(currentElementTreeNode.LayoutElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) {
					var contentSize Clay_Dimensions = Clay_Dimensions{}
					if layoutConfig.LayoutDirection == CLAY_LEFT_TO_RIGHT {
						for i := int32(0); int64(i) < int64(currentElement.ChildrenOrTextContent.Children.Length); i++ {
							var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
							contentSize.Width += childElement.Dimensions.Width
							if contentSize.Height > childElement.Dimensions.Height {
								contentSize.Height = contentSize.Height
							} else {
								contentSize.Height = childElement.Dimensions.Height
							}
						}
						contentSize.Width += float32((func() int64 {
							if (int64(currentElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
								return int64(currentElement.ChildrenOrTextContent.Children.Length) - 1
							}
							return 0
						}()) * int64(layoutConfig.ChildGap))
						var extraSpace float32 = currentElement.Dimensions.Width - float32(int64(layoutConfig.Padding.Left)+int64(layoutConfig.Padding.Right)) - contentSize.Width
						switch layoutConfig.ChildAlignment.X {
						case CLAY_ALIGN_X_LEFT:
							extraSpace = 0
						case CLAY_ALIGN_X_CENTER:
							extraSpace /= 2
						default:
						}
						currentElementTreeNode.NextChildOffset.X += extraSpace
					} else {
						for i := int32(0); int64(i) < int64(currentElement.ChildrenOrTextContent.Children.Length); i++ {
							var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
							if contentSize.Width > childElement.Dimensions.Width {
								contentSize.Width = contentSize.Width
							} else {
								contentSize.Width = childElement.Dimensions.Width
							}
							contentSize.Height += childElement.Dimensions.Height
						}
						contentSize.Height += float32((func() int64 {
							if (int64(currentElement.ChildrenOrTextContent.Children.Length) - 1) > 0 {
								return int64(currentElement.ChildrenOrTextContent.Children.Length) - 1
							}
							return 0
						}()) * int64(layoutConfig.ChildGap))
						var extraSpace float32 = currentElement.Dimensions.Height - float32(int64(layoutConfig.Padding.Top)+int64(layoutConfig.Padding.Bottom)) - contentSize.Height
						switch layoutConfig.ChildAlignment.Y {
						case CLAY_ALIGN_Y_TOP:
							extraSpace = 0
						case CLAY_ALIGN_Y_CENTER:
							extraSpace /= 2
						default:
						}
						currentElementTreeNode.NextChildOffset.Y += extraSpace
					}
					if scrollContainerData != nil {
						scrollContainerData.ContentSize = Clay_Dimensions{Width: contentSize.Width + float32(int64(layoutConfig.Padding.Left)+int64(layoutConfig.Padding.Right)), Height: contentSize.Height + float32(int64(layoutConfig.Padding.Top)+int64(layoutConfig.Padding.Bottom))}
					}
				}
			} else {
				var (
					closeScrollElement bool                      = false
					scrollConfig       *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
				)
				if scrollConfig != nil {
					closeScrollElement = true
					for i := int32(0); int64(i) < int64(context.ScrollContainerDatas.Length); i++ {
						var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
						if mapping.LayoutElement == currentElement {
							if scrollConfig.Horizontal {
								scrollOffset.X = mapping.ScrollPosition.X
							}
							if scrollConfig.Vertical {
								scrollOffset.Y = mapping.ScrollPosition.Y
							}
							if context.ExternalScrollHandlingEnabled {
								scrollOffset = Clay_Vector2{}
							}
							break
						}
					}
				}
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_BORDER) {
					var (
						currentElementData        *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(currentElement.Id)
						currentElementBoundingBox Clay_BoundingBox               = currentElementData.BoundingBox
					)
					if !Clay__ElementIsOffscreen(&currentElementBoundingBox) {
						var sharedConfig *Clay_SharedElementConfig
						if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SHARED) {
							sharedConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SHARED).SharedElementConfig
						} else {
							sharedConfig = &Clay_SharedElementConfig_DEFAULT
						}
						var borderConfig *Clay_BorderElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_BORDER).BorderElementConfig
						var renderCommand Clay_RenderCommand = Clay_RenderCommand{BoundingBox: currentElementBoundingBox, RenderData: Clay_RenderData{Border: Clay_BorderRenderData{Color: borderConfig.Color, CornerRadius: sharedConfig.CornerRadius, Width: borderConfig.Width}}, UserData: sharedConfig.UserData, Id: Clay__HashNumber(currentElement.Id, uint32(currentElement.ChildrenOrTextContent.Children.Length)).Id, CommandType: CLAY_RENDER_COMMAND_TYPE_BORDER}
						Clay__AddRenderCommand(renderCommand)
						if int64(borderConfig.Width.BetweenChildren) > 0 && borderConfig.Color.A > 0 {
							var (
								halfGap      float32      = float32(int64(layoutConfig.ChildGap) / 2)
								borderOffset Clay_Vector2 = Clay_Vector2{X: float32(layoutConfig.Padding.Left) - halfGap, Y: float32(layoutConfig.Padding.Top) - halfGap}
							)
							if layoutConfig.LayoutDirection == CLAY_LEFT_TO_RIGHT {
								for i := int32(0); int64(i) < int64(currentElement.ChildrenOrTextContent.Children.Length); i++ {
									var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
									if int64(i) > 0 {
										Clay__AddRenderCommand(Clay_RenderCommand{BoundingBox: Clay_BoundingBox{X: currentElementBoundingBox.X + borderOffset.X + scrollOffset.X, Y: currentElementBoundingBox.Y + scrollOffset.Y, Width: float32(borderConfig.Width.BetweenChildren), Height: currentElement.Dimensions.Height}, RenderData: Clay_RenderData{Rectangle: Clay_RectangleRenderData{BackgroundColor: borderConfig.Color}}, UserData: sharedConfig.UserData, Id: Clay__HashNumber(currentElement.Id, uint32(int32(int64(currentElement.ChildrenOrTextContent.Children.Length)+1+int64(i)))).Id, CommandType: CLAY_RENDER_COMMAND_TYPE_RECTANGLE})
									}
									borderOffset.X += childElement.Dimensions.Width + float32(layoutConfig.ChildGap)
								}
							} else {
								for i := int32(0); int64(i) < int64(currentElement.ChildrenOrTextContent.Children.Length); i++ {
									var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
									if int64(i) > 0 {
										Clay__AddRenderCommand(Clay_RenderCommand{BoundingBox: Clay_BoundingBox{X: currentElementBoundingBox.X + scrollOffset.X, Y: currentElementBoundingBox.Y + borderOffset.Y + scrollOffset.Y, Width: currentElement.Dimensions.Width, Height: float32(borderConfig.Width.BetweenChildren)}, RenderData: Clay_RenderData{Rectangle: Clay_RectangleRenderData{BackgroundColor: borderConfig.Color}}, UserData: sharedConfig.UserData, Id: Clay__HashNumber(currentElement.Id, uint32(int32(int64(currentElement.ChildrenOrTextContent.Children.Length)+1+int64(i)))).Id, CommandType: CLAY_RENDER_COMMAND_TYPE_RECTANGLE})
									}
									borderOffset.Y += childElement.Dimensions.Height + float32(layoutConfig.ChildGap)
								}
							}
						}
					}
				}
				if closeScrollElement {
					Clay__AddRenderCommand(Clay_RenderCommand{Id: Clay__HashNumber(currentElement.Id, uint32(int32(int64(rootElement.ChildrenOrTextContent.Children.Length)+11))).Id, CommandType: CLAY_RENDER_COMMAND_TYPE_SCISSOR_END})
				}
				dfsBuffer.Length--
				continue
			}
			if !Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) {
				dfsBuffer.Length += int32(currentElement.ChildrenOrTextContent.Children.Length)
				for i := int32(0); int64(i) < int64(currentElement.ChildrenOrTextContent.Children.Length); i++ {
					var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
					if layoutConfig.LayoutDirection == CLAY_LEFT_TO_RIGHT {
						currentElementTreeNode.NextChildOffset.Y = float32(currentElement.LayoutConfig.Padding.Top)
						var whiteSpaceAroundChild float32 = currentElement.Dimensions.Height - float32(int64(layoutConfig.Padding.Top)+int64(layoutConfig.Padding.Bottom)) - childElement.Dimensions.Height
						switch layoutConfig.ChildAlignment.Y {
						case CLAY_ALIGN_Y_TOP:
						case CLAY_ALIGN_Y_CENTER:
							currentElementTreeNode.NextChildOffset.Y += whiteSpaceAroundChild / 2
						case CLAY_ALIGN_Y_BOTTOM:
							currentElementTreeNode.NextChildOffset.Y += whiteSpaceAroundChild
						}
					} else {
						currentElementTreeNode.NextChildOffset.X = float32(currentElement.LayoutConfig.Padding.Left)
						var whiteSpaceAroundChild float32 = currentElement.Dimensions.Width - float32(int64(layoutConfig.Padding.Left)+int64(layoutConfig.Padding.Right)) - childElement.Dimensions.Width
						switch layoutConfig.ChildAlignment.X {
						case CLAY_ALIGN_X_LEFT:
						case CLAY_ALIGN_X_CENTER:
							currentElementTreeNode.NextChildOffset.X += whiteSpaceAroundChild / 2
						case CLAY_ALIGN_X_RIGHT:
							currentElementTreeNode.NextChildOffset.X += whiteSpaceAroundChild
						}
					}
					var childPosition Clay_Vector2 = Clay_Vector2{X: currentElementTreeNode.Position.X + currentElementTreeNode.NextChildOffset.X + scrollOffset.X, Y: currentElementTreeNode.Position.Y + currentElementTreeNode.NextChildOffset.Y + scrollOffset.Y}
					var newNodeIndex uint32 = uint32(int32(int64(dfsBuffer.Length) - 1 - int64(i)))
					*(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(dfsBuffer.InternalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(newNodeIndex))) = Clay__LayoutElementTreeNode{LayoutElement: childElement, Position: Clay_Vector2{X: childPosition.X, Y: childPosition.Y}, NextChildOffset: Clay_Vector2{X: float32(childElement.LayoutConfig.Padding.Left), Y: float32(childElement.LayoutConfig.Padding.Top)}}
					*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), newNodeIndex)) = false
					if layoutConfig.LayoutDirection == CLAY_LEFT_TO_RIGHT {
						currentElementTreeNode.NextChildOffset.X += childElement.Dimensions.Width + float32(layoutConfig.ChildGap)
					} else {
						currentElementTreeNode.NextChildOffset.Y += childElement.Dimensions.Height + float32(layoutConfig.ChildGap)
					}
				}
			}
		}
		if int64(root.ClipElementId) != 0 {
			Clay__AddRenderCommand(Clay_RenderCommand{Id: Clay__HashNumber(rootElement.Id, uint32(int32(int64(rootElement.ChildrenOrTextContent.Children.Length)+11))).Id, CommandType: CLAY_RENDER_COMMAND_TYPE_SCISSOR_END})
		}
	}
}

var Clay__debugViewWidth uint32 = 400
var Clay__debugViewHighlightColor Clay_Color = Clay_Color{R: 168, G: 66, B: 28, A: 100}

func Clay__WarningArray_Allocate_Arena(capacity int32, arena *Clay_Arena) Clay__WarningArray {
	var (
		totalSizeBytes  uint64             = uint64(uintptr(capacity) * unsafe.Sizeof(Clay_String{}))
		array           Clay__WarningArray = Clay__WarningArray{Capacity: capacity, Length: 0}
		nextAllocOffset uint64             = arena.NextAllocation + (64 - arena.NextAllocation%64)
	)
	if nextAllocOffset+totalSizeBytes <= arena.Capacity {
		array.InternalArray = (*Clay__Warning)(unsafe.Pointer(uintptr(uint64(uintptr(unsafe.Pointer(arena.Memory))) + nextAllocOffset)))
		arena.NextAllocation = nextAllocOffset + totalSizeBytes
	} else {
		Clay__currentContext.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_ARENA_CAPACITY_EXCEEDED, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to allocate memory in its arena, but ran out of capacity. Try increasing the capacity of the arena passed to Clay_Initialize()")}, UserData: Clay__currentContext.ErrorHandler.UserData})
	}
	return array
}
func Clay__WarningArray_Add(array *Clay__WarningArray, item Clay__Warning) *Clay__Warning {
	if int64(array.Length) < int64(array.Capacity) {
		*(*Clay__Warning)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__Warning{})*uintptr(func() int32 {
			p := &array.Length
			x := *p
			*p++
			return x
		}()))) = item
		return (*Clay__Warning)(unsafe.Add(unsafe.Pointer(array.InternalArray), unsafe.Sizeof(Clay__Warning{})*uintptr(int64(array.Length)-1)))
	}
	return &CLAY__WARNING_DEFAULT
}
func Clay__Array_Allocate_Arena(capacity int32, itemSize uint32, arena *Clay_Arena) unsafe.Pointer {
	var (
		totalSizeBytes  uint64 = uint64(int64(capacity) * int64(itemSize))
		nextAllocOffset uint64 = arena.NextAllocation + (64 - arena.NextAllocation%64)
	)
	if nextAllocOffset+totalSizeBytes <= arena.Capacity {
		arena.NextAllocation = nextAllocOffset + totalSizeBytes
		return unsafe.Pointer(uintptr(uint64(uintptr(unsafe.Pointer(arena.Memory))) + nextAllocOffset))
	} else {
		Clay__currentContext.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_ARENA_CAPACITY_EXCEEDED, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to allocate memory in its arena, but ran out of capacity. Try increasing the capacity of the arena passed to Clay_Initialize()")}, UserData: Clay__currentContext.ErrorHandler.UserData})
	}
	return unsafe.Pointer(uintptr(CLAY__NULL))
}
func Clay__Array_RangeCheck(index int32, length int32) bool {
	if int64(index) < int64(length) && int64(index) >= 0 {
		return true
	}
	var context *Clay_Context = Clay_GetCurrentContext()
	context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_INTERNAL_ERROR, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to make an out of bounds array access. This is an internal error and is likely a bug.")}, UserData: context.ErrorHandler.UserData})
	return false
}
func Clay__Array_AddCapacityCheck(length int32, capacity int32) bool {
	if int64(length) < int64(capacity) {
		return true
	}
	var context *Clay_Context = Clay_GetCurrentContext()
	context.ErrorHandler.ErrorHandlerFunction(Clay_ErrorData{ErrorType: CLAY_ERROR_TYPE_INTERNAL_ERROR, ErrorText: Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay attempted to make an out of bounds array access. This is an internal error and is likely a bug.")}, UserData: context.ErrorHandler.UserData})
	return false
}
func Clay_MinMemorySize() uint32 {
	var (
		fakeContext    Clay_Context  = Clay_Context{MaxElementCount: Clay__defaultMaxElementCount, MaxMeasureTextCacheWordCount: Clay__defaultMaxMeasureTextWordCacheCount, InternalArena: Clay_Arena{Capacity: math.MaxUint64, Memory: nil}}
		currentContext *Clay_Context = Clay_GetCurrentContext()
	)
	if currentContext != nil {
		fakeContext.MaxElementCount = currentContext.MaxElementCount
		fakeContext.MaxMeasureTextCacheWordCount = currentContext.MaxElementCount
	}
	Clay__Context_Allocate_Arena(&fakeContext.InternalArena)
	Clay__InitializePersistentMemory(&fakeContext)
	Clay__InitializeEphemeralMemory(&fakeContext)
	return uint32(fakeContext.InternalArena.NextAllocation + 128)
}
func Clay_CreateArenaWithCapacityAndMemory(capacity uint32, memory unsafe.Pointer) Clay_Arena {
	var arena Clay_Arena = Clay_Arena{Capacity: uint64(capacity), Memory: (*byte)(memory)}
	return arena
}
func Clay_SetMeasureTextFunction(measureTextFunction func(text Clay_StringSlice, config *Clay_TextElementConfig, userData unsafe.Pointer) Clay_Dimensions, userData unsafe.Pointer) {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__MeasureText = measureTextFunction
	context.MeasureTextUserData = userData
}
func Clay_SetQueryScrollOffsetFunction(queryScrollOffsetFunction func(elementId uint32, userData unsafe.Pointer) Clay_Vector2, userData unsafe.Pointer) {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__QueryScrollOffset = queryScrollOffsetFunction
	context.QueryScrollOffsetUserData = userData
}
func Clay_SetLayoutDimensions(dimensions Clay_Dimensions) {
	Clay_GetCurrentContext().LayoutDimensions = dimensions
}
func Clay_SetPointerState(position Clay_Vector2, isPointerDown bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return
	}
	context.PointerInfo.Position = position
	context.PointerOverIds.Length = 0
	var dfsBuffer Clay__int32_tArray = context.LayoutElementChildrenBuffer
	for rootIndex := int32(int32(int64(context.LayoutElementTreeRoots.Length) - 1)); int64(rootIndex) >= 0; rootIndex-- {
		dfsBuffer.Length = 0
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.LayoutElementTreeRoots, rootIndex)
		Clay__int32_tArray_Add(&dfsBuffer, root.LayoutElementIndex)
		*context.TreeNodeVisited.InternalArray = false
		var found bool = false
		for int64(dfsBuffer.Length) > 0 {
			if *(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), int64(dfsBuffer.Length)-1)) {
				dfsBuffer.Length--
				continue
			}
			*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), int64(dfsBuffer.Length)-1)) = true
			var currentElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, Clay__int32_tArray_GetValue(&dfsBuffer, int32(int64(dfsBuffer.Length)-1)))
			var mapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(currentElement.Id)
			var elementBox Clay_BoundingBox = mapItem.BoundingBox
			elementBox.X -= root.PointerOffset.X
			elementBox.Y -= root.PointerOffset.Y
			if mapItem != nil {
				if Clay__PointIsInsideRect(position, elementBox) {
					if mapItem.OnHoverFunction != nil {
						mapItem.OnHoverFunction(mapItem.ElementId, context.PointerInfo, mapItem.HoverFunctionUserData)
					}
					Clay__ElementIdArray_Add(&context.PointerOverIds, mapItem.ElementId)
					found = true
					if int64(mapItem.IdAlias) != 0 {
						Clay__ElementIdArray_Add(&context.PointerOverIds, Clay_ElementId{Id: mapItem.IdAlias})
					}
				}
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) {
					dfsBuffer.Length--
					continue
				}
				for i := int32(int32(int64(currentElement.ChildrenOrTextContent.Children.Length) - 1)); int64(i) >= 0; i-- {
					Clay__int32_tArray_Add(&dfsBuffer, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.ChildrenOrTextContent.Children.Elements), unsafe.Sizeof(int32(0))*uintptr(i))))
					*(*bool)(unsafe.Add(unsafe.Pointer(context.TreeNodeVisited.InternalArray), int64(dfsBuffer.Length)-1)) = false
				}
			} else {
				dfsBuffer.Length--
			}
		}
		var rootElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.LayoutElements, root.LayoutElementIndex)
		if found && Clay__ElementHasConfig(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING) && Clay__FindElementConfigWithType(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING).FloatingElementConfig.PointerCaptureMode == CLAY_POINTER_CAPTURE_MODE_CAPTURE {
			break
		}
	}
	if isPointerDown {
		if context.PointerInfo.State == CLAY_POINTER_DATA_PRESSED_THIS_FRAME {
			context.PointerInfo.State = CLAY_POINTER_DATA_PRESSED
		} else if context.PointerInfo.State != CLAY_POINTER_DATA_PRESSED {
			context.PointerInfo.State = CLAY_POINTER_DATA_PRESSED_THIS_FRAME
		}
	} else {
		if context.PointerInfo.State == CLAY_POINTER_DATA_RELEASED_THIS_FRAME {
			context.PointerInfo.State = CLAY_POINTER_DATA_RELEASED
		} else if context.PointerInfo.State != CLAY_POINTER_DATA_RELEASED {
			context.PointerInfo.State = CLAY_POINTER_DATA_RELEASED_THIS_FRAME
		}
	}
}
func Clay_Initialize(arena Clay_Arena, layoutDimensions Clay_Dimensions, errorHandler Clay_ErrorHandler) *Clay_Context {
	var context *Clay_Context = Clay__Context_Allocate_Arena(&arena)
	if context == nil {
		return nil
	}
	var oldContext *Clay_Context = Clay_GetCurrentContext()
	*context = Clay_Context{MaxElementCount: int32(func() int64 {
		if oldContext != nil {
			return int64(oldContext.MaxElementCount)
		}
		return int64(Clay__defaultMaxElementCount)
	}()), MaxMeasureTextCacheWordCount: int32(func() int64 {
		if oldContext != nil {
			return int64(oldContext.MaxMeasureTextCacheWordCount)
		}
		return int64(Clay__defaultMaxMeasureTextWordCacheCount)
	}()), ErrorHandler: func() Clay_ErrorHandler {
		if errorHandler.ErrorHandlerFunction != nil {
			return errorHandler
		}
		return Clay_ErrorHandler{ErrorHandlerFunction: Clay__ErrorHandlerFunctionDefault, UserData: nil}
	}(), LayoutDimensions: layoutDimensions, InternalArena: arena}
	Clay_SetCurrentContext(context)
	Clay__InitializePersistentMemory(context)
	Clay__InitializeEphemeralMemory(context)
	for i := int32(0); int64(i) < int64(context.LayoutElementsHashMap.Capacity); i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.LayoutElementsHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = -1
	}
	for i := int32(0); int64(i) < int64(context.MeasureTextHashMap.Capacity); i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = 0
	}
	context.MeasureTextHashMapInternal.Length = 1
	context.LayoutDimensions = layoutDimensions
	return context
}
func Clay_GetCurrentContext() *Clay_Context {
	return Clay__currentContext
}
func Clay_SetCurrentContext(context *Clay_Context) {
	Clay__currentContext = context
}
func Clay_UpdateScrollContainers(enableDragScrolling bool, scrollDelta Clay_Vector2, deltaTime float32) {
	var (
		context                     *Clay_Context                      = Clay_GetCurrentContext()
		isPointerActive             bool                               = enableDragScrolling && (context.PointerInfo.State == CLAY_POINTER_DATA_PRESSED || context.PointerInfo.State == CLAY_POINTER_DATA_PRESSED_THIS_FRAME)
		highestPriorityElementIndex int32                              = -1
		highestPriorityScrollData   *Clay__ScrollContainerDataInternal = (*Clay__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(CLAY__NULL)))
	)
	for i := int32(0); int64(i) < int64(context.ScrollContainerDatas.Length); i++ {
		var scrollData *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
		if !scrollData.OpenThisFrame {
			Clay__ScrollContainerDataInternalArray_RemoveSwapback(&context.ScrollContainerDatas, i)
			continue
		}
		scrollData.OpenThisFrame = false
		var hashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(scrollData.ElementId)
		if hashMapItem == nil {
			Clay__ScrollContainerDataInternalArray_RemoveSwapback(&context.ScrollContainerDatas, i)
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
			scrollData.PointerOrigin = Clay_Vector2{}
			scrollData.ScrollOrigin = Clay_Vector2{}
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
		for j := int32(0); int64(j) < int64(context.PointerOverIds.Length); j++ {
			if int64(scrollData.LayoutElement.Id) == int64(Clay__ElementIdArray_Get(&context.PointerOverIds, j).Id) {
				highestPriorityElementIndex = j
				highestPriorityScrollData = scrollData
			}
		}
	}
	if int64(highestPriorityElementIndex) > -1 && highestPriorityScrollData != nil {
		var (
			scrollElement         *Clay_LayoutElement       = highestPriorityScrollData.LayoutElement
			scrollConfig          *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(scrollElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig
			canScrollVertically   bool                      = scrollConfig.Vertical && highestPriorityScrollData.ContentSize.Height > scrollElement.Dimensions.Height
			canScrollHorizontally bool                      = scrollConfig.Horizontal && highestPriorityScrollData.ContentSize.Width > scrollElement.Dimensions.Width
		)
		if canScrollVertically {
			highestPriorityScrollData.ScrollPosition.Y = highestPriorityScrollData.ScrollPosition.Y + scrollDelta.Y*10
		}
		if canScrollHorizontally {
			highestPriorityScrollData.ScrollPosition.X = highestPriorityScrollData.ScrollPosition.X + scrollDelta.X*10
		}
		if isPointerActive {
			highestPriorityScrollData.ScrollMomentum = Clay_Vector2{}
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
func Clay_BeginLayout() {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__InitializeEphemeralMemory(context)
	context.Generation++
	context.DynamicElementIndex = 0
	var rootDimensions Clay_Dimensions = Clay_Dimensions{Width: context.LayoutDimensions.Width, Height: context.LayoutDimensions.Height}
	if context.DebugModeEnabled {
		rootDimensions.Width -= float32(Clay__debugViewWidth)
	}
	context.BooleanWarnings = Clay_BooleanWarnings{MaxElementsExceeded: false}
	Clay__OpenElement()
	Clay__ConfigureOpenElement(Clay_ElementDeclaration{Id: Clay__HashString(Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay__RootContainer")}, 0, 0), Layout: Clay_LayoutConfig{Sizing: Clay_Sizing{Width: Clay_SizingAxis{Size: struct {
		// union
		MinMax  Clay_SizingMinMax
		Percent float32
	}{MinMax: Clay_SizingMinMax{Min: rootDimensions.Width, Max: rootDimensions.Width}}, Type: CLAY__SIZING_TYPE_FIXED}, Height: Clay_SizingAxis{Size: struct {
		// union
		MinMax  Clay_SizingMinMax
		Percent float32
	}{MinMax: Clay_SizingMinMax{Min: rootDimensions.Height, Max: rootDimensions.Height}}, Type: CLAY__SIZING_TYPE_FIXED}}}})
	Clay__int32_tArray_Add(&context.OpenLayoutElementStack, 0)
	Clay__LayoutElementTreeRootArray_Add(&context.LayoutElementTreeRoots, Clay__LayoutElementTreeRoot{})
}
func Clay_EndLayout() Clay_RenderCommandArray {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__CloseElement()
	var elementsExceededBeforeDebugView bool = context.BooleanWarnings.MaxElementsExceeded
	if context.BooleanWarnings.MaxElementsExceeded {
		var message Clay_String
		if !elementsExceededBeforeDebugView {
			message = Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay Error: Layout elements exceeded Clay__maxElementCount after adding the debug-view to the layout.")}
		} else {
			message = Clay_String{Length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), Chars: libc.CString("Clay Error: Layout elements exceeded Clay__maxElementCount")}
		}
		Clay__AddRenderCommand(Clay_RenderCommand{BoundingBox: Clay_BoundingBox{X: context.LayoutDimensions.Width/2 - 59*4, Y: context.LayoutDimensions.Height / 2, Width: 0, Height: 0}, RenderData: Clay_RenderData{Text: Clay_TextRenderData{StringContents: Clay_StringSlice{Length: message.Length, Chars: message.Chars, BaseChars: message.Chars}, TextColor: Clay_Color{R: 255, G: 0, B: 0, A: 255}, FontSize: 16}}, CommandType: CLAY_RENDER_COMMAND_TYPE_TEXT})
	} else {
		Clay__CalculateFinalLayout()
	}
	return context.RenderCommands
}
func Clay_GetElementId(idString Clay_String) Clay_ElementId {
	return Clay__HashString(idString, 0, 0)
}
func Clay_GetElementIdWithIndex(idString Clay_String, index uint32) Clay_ElementId {
	return Clay__HashString(idString, index, 0)
}
func Clay_Hovered() bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return false
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	if int64(openLayoutElement.Id) == 0 {
		Clay__GenerateIdForAnonymousElement(openLayoutElement)
	}
	for i := int32(0); int64(i) < int64(context.PointerOverIds.Length); i++ {
		if int64(Clay__ElementIdArray_Get(&context.PointerOverIds, i).Id) == int64(openLayoutElement.Id) {
			return true
		}
	}
	return false
}
func Clay_OnHover(onHoverFunction func(elementId Clay_ElementId, pointerInfo Clay_PointerData, userData int64), userData int64) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.BooleanWarnings.MaxElementsExceeded {
		return
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	if int64(openLayoutElement.Id) == 0 {
		Clay__GenerateIdForAnonymousElement(openLayoutElement)
	}
	var hashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(openLayoutElement.Id)
	hashMapItem.OnHoverFunction = onHoverFunction
	hashMapItem.HoverFunctionUserData = userData
}
func Clay_PointerOver(elementId Clay_ElementId) bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	for i := int32(0); int64(i) < int64(context.PointerOverIds.Length); i++ {
		if int64(Clay__ElementIdArray_Get(&context.PointerOverIds, i).Id) == int64(elementId.Id) {
			return true
		}
	}
	return false
}
func Clay_GetScrollContainerData(id Clay_ElementId) Clay_ScrollContainerData {
	var context *Clay_Context = Clay_GetCurrentContext()
	for i := int32(0); int64(i) < int64(context.ScrollContainerDatas.Length); i++ {
		var scrollContainerData *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.ScrollContainerDatas, i)
		if int64(scrollContainerData.ElementId) == int64(id.Id) {
			return Clay_ScrollContainerData{ScrollPosition: &scrollContainerData.ScrollPosition, ScrollContainerDimensions: Clay_Dimensions{Width: scrollContainerData.BoundingBox.Width, Height: scrollContainerData.BoundingBox.Height}, ContentDimensions: scrollContainerData.ContentSize, Config: *Clay__FindElementConfigWithType(scrollContainerData.LayoutElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL).ScrollElementConfig, Found: true}
		}
	}
	return Clay_ScrollContainerData{}
}
func Clay_GetElementData(id Clay_ElementId) Clay_ElementData {
	var item *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(id.Id)
	if item == &Clay_LayoutElementHashMapItem_DEFAULT {
		return Clay_ElementData{}
	}
	return Clay_ElementData{BoundingBox: item.BoundingBox, Found: true}
}
func Clay_SetDebugModeEnabled(enabled bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.DebugModeEnabled = enabled
}
func Clay_IsDebugModeEnabled() bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	return context.DebugModeEnabled
}
func Clay_SetCullingEnabled(enabled bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.DisableCulling = !enabled
}
func Clay_SetExternalScrollHandlingEnabled(enabled bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.ExternalScrollHandlingEnabled = enabled
}
func Clay_GetMaxElementCount() int32 {
	var context *Clay_Context = Clay_GetCurrentContext()
	return context.MaxElementCount
}
func Clay_SetMaxElementCount(maxElementCount int32) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context != nil {
		context.MaxElementCount = maxElementCount
	} else {
		Clay__defaultMaxElementCount = maxElementCount
		Clay__defaultMaxMeasureTextWordCacheCount = int32(int64(maxElementCount) * 2)
	}
}
func Clay_GetMaxMeasureTextCacheWordCount() int32 {
	var context *Clay_Context = Clay_GetCurrentContext()
	return context.MaxMeasureTextCacheWordCount
}
func Clay_SetMaxMeasureTextCacheWordCount(maxMeasureTextCacheWordCount int32) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context != nil {
		Clay__currentContext.MaxMeasureTextCacheWordCount = maxMeasureTextCacheWordCount
	} else {
		Clay__defaultMaxMeasureTextWordCacheCount = maxMeasureTextCacheWordCount
	}
}
func Clay_ResetMeasureTextCache() {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.MeasureTextHashMapInternal.Length = 0
	context.MeasureTextHashMapInternalFreeList.Length = 0
	context.MeasureTextHashMap.Length = 0
	context.MeasuredWords.Length = 0
	context.MeasuredWordsFreeList.Length = 0
	for i := int32(0); int64(i) < int64(context.MeasureTextHashMap.Capacity); i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.MeasureTextHashMap.InternalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = 0
	}
	context.MeasureTextHashMapInternal.Length = 1
}
