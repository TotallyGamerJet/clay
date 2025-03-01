package main

import (
	"fmt"
	"unsafe"

	"github.com/totallygamerjet/clay"
)

func handleClayError(errorText clay.ErrorData) {
	fmt.Println(errorText.ErrorText)
}

// TODO: CreateArenaWithCapacityAndMemory should take a slice of bytes

func main() {
	const screenWidth, screenHeight = 800, 600

	totalMemorySize := clay.MinMemorySize()
	memory := make([]byte, totalMemorySize)
	arena := clay.CreateArenaWithCapacityAndMemory(totalMemorySize, unsafe.Pointer(unsafe.SliceData(memory)))
	clay.Initialize(arena, clay.Dimensions{Width: screenWidth, Height: screenHeight}, clay.ErrorHandler{ErrorHandlerFunction: handleClayError})
	for renderLoop() {
		clay.SetLayoutDimensions(clay.Dimensions{Width: screenWidth, Height: screenHeight})
		//clay.SetPointerState(clay.Vector2{X: mousePositionX, mousePositionY}, isMouseDown)
		//clay.UpdateScrollContainers(true, clay.Vector2{X: mouseWheelX, Y: mouseWheelY}, deltaTime)

		clay.BeginLayout()

		clay.UI(clay.ElementDeclaration{Id: clay.ID("OuterContainer"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0), Height: clay.SizingGrow(0)}, Padding: clay.PaddingAll(16), ChildGap: 16}, BackgroundColor: clay.Color{R: 250, G: 250, B: 255, A: 255}},
			func() {
				clay.UI(clay.ElementDeclaration{Id: clay.ID("SideBar"), Layout: clay.LayoutConfig{LayoutDirection: clay.TOP_TO_BOTTOM, Sizing: clay.Sizing{Width: clay.SizingFixed(300), Height: clay.SizingGrow(0)}, Padding: clay.PaddingAll(16), ChildGap: 16}},
					func() {

					})
			},
		)

		var renderCommands = clay.EndLayout()
		_ = renderCommands
	}
	//		// Note: malloc is only used here as an example, any allocator that provides
	//		// a pointer to addressable memory of at least totalMemorySize will work
	//		uint64_t totalMemorySize = Clay_MinMemorySize();
	//		Clay_Arena arena = Clay_CreateArenaWithCapacityAndMemory(totalMemorySize, malloc(totalMemorySize));
	//
	//		// Note: screenWidth and screenHeight will need to come from your environment, Clay doesn't handle window related tasks
	//		Clay_Initialize(arena, (Clay_Dimensions) { screenWidth, screenHeight }, (Clay_ErrorHandler) { HandleClayErrors });
	//
	//		while(renderLoop()) { // Will be different for each renderer / environment
	//			// Optional: Update internal layout dimensions to support resizing
	//			Clay_SetLayoutDimensions((Clay_Dimensions) { screenWidth, screenHeight });
	//			// Optional: Update internal pointer position for handling mouseover / click / touch events - needed for scrolling & debug tools
	//			Clay_SetPointerState((Clay_Vector2) { mousePositionX, mousePositionY }, isMouseDown);
	//			// Optional: Update internal pointer position for handling mouseover / click / touch events - needed for scrolling and debug tools
	//			Clay_UpdateScrollContainers(true, (Clay_Vector2) { mouseWheelX, mouseWheelY }, deltaTime);
	//
	//			// All clay layouts are declared between Clay_BeginLayout and Clay_EndLayout
	//			Clay_BeginLayout();
	//
	//			// An example of laying out a UI with a fixed width sidebar and flexible width main content
	//			CLAY({ .id = CLAY_ID("OuterContainer"), .layout = { .sizing = {CLAY_SIZING_GROW(0), CLAY_SIZING_GROW(0)}, .padding = CLAY_PADDING_ALL(16), .childGap = 16 }, .backgroundColor = {250,250,255,255} }) {
	//				CLAY({
	//					.id = CLAY_ID("SideBar"),
	//					.layout = { .layoutDirection = CLAY_TOP_TO_BOTTOM, .sizing = { .width = CLAY_SIZING_FIXED(300), .height = CLAY_SIZING_GROW(0) }, .padding = CLAY_PADDING_ALL(16), .childGap = 16 },
	//					.backgroundColor = COLOR_LIGHT }
	//			}) {
	//				CLAY({ .id = CLAY_ID("ProfilePictureOuter"), .layout = { .sizing = { .width = CLAY_SIZING_GROW(0) }, .padding = CLAY_PADDING_ALL(16), .childGap = 16, .childAlignment = { .y = CLAY_ALIGN_Y_CENTER } }, .backgroundColor = COLOR_RED }) {
	//					CLAY({ .id = CLAY_ID("ProfilePicture"), .layout = { .sizing = { .width = CLAY_SIZING_FIXED(60), .height = CLAY_SIZING_FIXED(60) }}, .image = { .imageData = &profilePicture, .sourceDimensions = {60, 60} } }) {}
	//					CLAY_TEXT(CLAY_STRING("Clay - UI Library"), CLAY_TEXT_CONFIG({ .fontSize = 24, .textColor = {255, 255, 255, 255} }));
	//				}
	//
	//				// Standard C code like loops etc work inside components
	//				for (int i = 0; i < 5; i++) {
	//					SidebarItemComponent();
	//				}
	//
	//				CLAY({ .id = CLAY_ID("MainContent"), .layout = { .sizing = { .width = CLAY_SIZING_GROW(0), .height = CLAY_SIZING_GROW(0) } }, .backgroundColor = COLOR_LIGHT }) {}
	//			}
	//		}
	//
	//		// All clay layouts are declared between Clay_BeginLayout and Clay_EndLayout
	//		Clay_RenderCommandArray renderCommands = Clay_EndLayout();
	//
	//		// More comprehensive rendering examples can be found in the renderers/ directory
	//		for (int i = 0; i < renderCommands.length; i++) {
	//			Clay_RenderCommand *renderCommand = &renderCommands.internalArray[i];
	//
	//			switch (renderCommand->commandType) {
	//			case CLAY_RENDER_COMMAND_TYPE_RECTANGLE: {
	//				DrawRectangle( renderCommand->boundingBox, renderCommand->renderData.rectangle.backgroundColor);
	//			}
	//				// ... Implement handling of other command types
	//			}
	//		}
	//}

}
