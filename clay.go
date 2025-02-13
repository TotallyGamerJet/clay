package clay

import "unsafe"

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
