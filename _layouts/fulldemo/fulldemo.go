package fulldemo

import (
	"unsafe"

	"github.com/totallygamerjet/clay"
)

const (
	FontIdBody24 = 0
	FontIdBody16 = 1
)

const profileText = "Profile Page one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen"

var (
	orange = clay.Color{R: 255, G: 138, B: 50, A: 255}
	blue   = clay.Color{R: 111, G: 173, B: 162, A: 255}

	headerTextConfig = clay.TextElementConfig{FontId: 1, FontSize: 16, TextColor: clay.Color{0, 0, 0, 255}}
)

func headerButtonStyle(hovered bool) clay.ElementDeclaration {
	return clay.ElementDeclaration{
		Layout: clay.LayoutConfig{Padding: clay.Padding{16, 16, 8, 8}},
		BackgroundColor: func() clay.Color {
			if hovered {
				return orange
			}
			return blue
		}(),
	}
}

// Examples of re-usable "Components"
func renderHeaderButton(text string) {
	clay.UI()(headerButtonStyle(clay.Hovered()), func() {
		clay.Text(text, clay.TextConfig(headerTextConfig))
	})
}

func CreateLayout(profilePicture unsafe.Pointer) clay.RenderCommandArray {
	clay.BeginLayout()

	clay.UI()(clay.ElementDeclaration{Id: clay.ID("OuterContainer"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0), Height: clay.SizingGrow(0)}, Padding: clay.PaddingAll(16), ChildGap: 16}, BackgroundColor: clay.Color{R: 200, G: 200, B: 200, A: 255}}, func() {
		clay.UI()(clay.ElementDeclaration{Id: clay.ID("SideBar"), Layout: clay.LayoutConfig{LayoutDirection: clay.TOP_TO_BOTTOM, Sizing: clay.Sizing{Width: clay.SizingFixed(300), Height: clay.SizingGrow(0)}, Padding: clay.PaddingAll(16), ChildGap: 16}, BackgroundColor: clay.Color{R: 150, G: 150, B: 255, A: 255}}, func() {
			clay.UI()(clay.ElementDeclaration{Id: clay.ID("ProfilePictureOuter"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0)}, Padding: clay.PaddingAll(8), ChildGap: 8, ChildAlignment: clay.ChildAlignment{Y: clay.ALIGN_Y_CENTER}}, BackgroundColor: clay.Color{R: 130, G: 130, B: 255, A: 255}}, func() {
				clay.UI()(clay.ElementDeclaration{Id: clay.ID("ProfilePicture"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingFixed(60), Height: clay.SizingFixed(60)}}, Image: clay.ImageElementConfig{ImageData: unsafe.Pointer(&profilePicture), SourceDimensions: clay.Dimensions{Width: 60, Height: 60}}}, func() {})
				clay.Text(profileText, clay.TextConfig(clay.TextElementConfig{FontSize: 24, TextColor: clay.Color{A: 255}, TextAlignment: clay.TEXT_ALIGN_RIGHT}))
			})
			clay.UI()(clay.ElementDeclaration{Id: clay.ID("SidebarBlob1"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0), Height: clay.SizingFixed(50)}}, BackgroundColor: clay.Color{R: 110, G: 110, B: 255, A: 255}}, func() {})
			clay.UI()(clay.ElementDeclaration{Id: clay.ID("SidebarBlob2"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0), Height: clay.SizingFixed(50)}}, BackgroundColor: clay.Color{R: 110, G: 110, B: 255, A: 255}}, func() {})
			clay.UI()(clay.ElementDeclaration{Id: clay.ID("SidebarBlob3"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0), Height: clay.SizingFixed(50)}}, BackgroundColor: clay.Color{R: 110, G: 110, B: 255, A: 255}}, func() {})
			clay.UI()(clay.ElementDeclaration{Id: clay.ID("SidebarBlob4"), Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0), Height: clay.SizingFixed(50)}}, BackgroundColor: clay.Color{R: 110, G: 110, B: 255, A: 255}}, func() {})
		})
		clay.UI()(clay.ElementDeclaration{Id: clay.ID("RightPanel"), Layout: clay.LayoutConfig{LayoutDirection: clay.TOP_TO_BOTTOM, Sizing: clay.Sizing{Width: clay.SizingGrow(0), Height: clay.SizingGrow(0)}, ChildGap: 16}}, func() {
			clay.UI()(clay.ElementDeclaration{Layout: clay.LayoutConfig{Sizing: clay.Sizing{Width: clay.SizingGrow(0)}, ChildAlignment: clay.ChildAlignment{X: clay.ALIGN_X_RIGHT}, Padding: clay.PaddingAll(8), ChildGap: 8}, BackgroundColor: clay.Color{180, 180, 180, 255}}, func() {
				renderHeaderButton("Header Item 1")
				renderHeaderButton("Header Item 2")
				renderHeaderButton("Header Item 3")
			})
		})
	})

	return clay.EndLayout()
}
