package videodemo

import (
	"bytes"
	_ "embed"
	"image"
	"image/draw"
	"image/jpeg"
	"unsafe"

	"github.com/TotallyGamerJet/clay"
)

var white = clay.Color{R: 255, G: 255, B: 255, A: 255}

// License: CC0 Public Domain
//
// Daniel Smith has released this “Squirrel Eating Nut” image under Public Domain license.
// It means that you can use and modify it for your personal and commercial projects.
// If you intend to use an image you find here for commercial use, please be aware that some photos do require a
// model or a property release. Pictures featuring products should be used with care.

//go:embed squirrel-eating-nut.jpg
var squirrelImageJPG []byte

var SquirrelImage *image.RGBA

func init() {
	img, err := jpeg.Decode(bytes.NewReader(squirrelImageJPG))
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	SquirrelImage = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(SquirrelImage, SquirrelImage.Bounds(), img, bounds.Min, draw.Src)
}

const FontIdBody16 = 0

type document struct {
	title    string
	contents string
	image    any
}

type arena struct {
	offset int64
	memory []byte
}

func alloc[T any](arena *arena) *T {
	prev := uintptr(arena.offset)
	arena.offset = int64(prev + unsafe.Sizeof(*new(T)))
	return (*T)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(arena.memory)), prev))
}

type Data struct {
	selectedDocumentIndex int32
	yOffset               float32
	frameArena            arena
	documents             []document
}

type sidebarClickData struct {
	requestedDocumentIndex int32
	selectedDocumentIndex  *int32
	documentLen            int32
}

func Initialize(squirrelImage any) Data {
	documents := make([]document, 5)
	documents[0] = document{image: squirrelImage, title: "Squirrels", contents: "The Secret Life of Squirrels: Nature's Clever Acrobats\n" + "Squirrels are often overlooked creatures, dismissed as mere park inhabitants or backyard nuisances. Yet, beneath their fluffy tails and twitching noses lies an intricate world of cunning, agility, and survival tactics that are nothing short of fascinating. As one of the most common mammals in North America, squirrels have adapted to a wide range of environments from bustling urban centers to tranquil forests and have developed a variety of unique behaviors that continue to intrigue scientists and nature enthusiasts alike.\n" + "\n" + "Master Tree Climbers\n" + "At the heart of a squirrel's skill set is its impressive ability to navigate trees with ease. Whether they're darting from branch to branch or leaping across wide gaps, squirrels possess an innate talent for acrobatics. Their powerful hind legs, which are longer than their front legs, give them remarkable jumping power. With a tail that acts as a counterbalance, squirrels can leap distances of up to ten times the length of their body, making them some of the best aerial acrobats in the animal kingdom.\n" + "But it's not just their agility that makes them exceptional climbers. Squirrels' sharp, curved claws allow them to grip tree bark with precision, while the soft pads on their feet provide traction on slippery surfaces. Their ability to run at high speeds and scale vertical trunks with ease is a testament to the evolutionary adaptations that have made them so successful in their arboreal habitats.\n" + "\n" + "Food Hoarders Extraordinaire\n" + "Squirrels are often seen frantically gathering nuts, seeds, and even fungi in preparation for winter. While this behavior may seem like instinctual hoarding, it is actually a survival strategy that has been honed over millions of years. Known as \"scatter hoarding,\" squirrels store their food in a variety of hidden locations, often burying it deep in the soil or stashing it in hollowed-out tree trunks.\n" + "Interestingly, squirrels have an incredible memory for the locations of their caches. Research has shown that they can remember thousands of hiding spots, often returning to them months later when food is scarce. However, they don't always recover every stash some forgotten caches eventually sprout into new trees, contributing to forest regeneration. This unintentional role as forest gardeners highlights the ecological importance of squirrels in their ecosystems.\n" + "\n" + "The Great Squirrel Debate: Urban vs. Wild\n" + "While squirrels are most commonly associated with rural or wooded areas, their adaptability has allowed them to thrive in urban environments as well. In cities, squirrels have become adept at finding food sources in places like parks, streets, and even garbage cans. However, their urban counterparts face unique challenges, including traffic, predators, and the lack of natural shelters. Despite these obstacles, squirrels in urban areas are often observed using human infrastructure such as buildings, bridges, and power lines as highways for their acrobatic escapades.\n" + "There is, however, a growing concern regarding the impact of urban life on squirrel populations. Pollution, deforestation, and the loss of natural habitats are making it more difficult for squirrels to find adequate food and shelter. As a result, conservationists are focusing on creating squirrel-friendly spaces within cities, with the goal of ensuring these resourceful creatures continue to thrive in both rural and urban landscapes.\n" + "\n" + "A Symbol of Resilience\n" + "In many cultures, squirrels are symbols of resourcefulness, adaptability, and preparation. Their ability to thrive in a variety of environments while navigating challenges with agility and grace serves as a reminder of the resilience inherent in nature. Whether you encounter them in a quiet forest, a city park, or your own backyard, squirrels are creatures that never fail to amaze with their endless energy and ingenuity.\n" + "In the end, squirrels may be small, but they are mighty in their ability to survive and thrive in a world that is constantly changing. So next time you spot one hopping across a branch or darting across your lawn, take a moment to appreciate the remarkable acrobat at work a true marvel of the natural world.\n"}
	documents[1] = document{title: "Lorem Ipsum", contents: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."}
	documents[2] = document{title: "Vacuum Instructions", contents: "Chapter 3: Getting Started - Unpacking and Setup\n" + "\n" + "Congratulations on your new SuperClean Pro 5000 vacuum cleaner! In this section, we will guide you through the simple steps to get your vacuum up and running. Before you begin, please ensure that you have all the components listed in the \"Package Contents\" section on page 2.\n" + "\n" + "1. Unboxing Your Vacuum\n" + "Carefully remove the vacuum cleaner from the box. Avoid using sharp objects that could damage the product. Once removed, place the unit on a flat, stable surface to proceed with the setup. Inside the box, you should find:\n" + "\n" + "    The main vacuum unit\n" + "    A telescoping extension wand\n" + "    A set of specialized cleaning tools (crevice tool, upholstery brush, etc.)\n" + "    A reusable dust bag (if applicable)\n" + "    A power cord with a 3-prong plug\n" + "    A set of quick-start instructions\n" + "\n" + "2. Assembling Your Vacuum\n" + "Begin by attaching the extension wand to the main body of the vacuum cleaner. Line up the connectors and twist the wand into place until you hear a click. Next, select the desired cleaning tool and firmly attach it to the wand's end, ensuring it is securely locked in.\n" + "\n" + "For models that require a dust bag, slide the bag into the compartment at the back of the vacuum, making sure it is properly aligned with the internal mechanism. If your vacuum uses a bagless system, ensure the dust container is correctly seated and locked in place before use.\n" + "\n" + "3. Powering On\n" + "To start the vacuum, plug the power cord into a grounded electrical outlet. Once plugged in, locate the power switch, usually positioned on the side of the handle or body of the unit, depending on your model. Press the switch to the \"On\" position, and you should hear the motor begin to hum. If the vacuum does not power on, check that the power cord is securely plugged in, and ensure there are no blockages in the power switch.\n" + "\n" + "Note: Before first use, ensure that the vacuum filter (if your model has one) is properly installed. If unsure, refer to \"Section 5: Maintenance\" for filter installation instructions."}
	documents[3] = document{title: "Article 4", contents: "Article 4"}
	documents[4] = document{title: "Article 5", contents: "Article 5"}

	data := Data{
		frameArena: arena{memory: make([]byte, 1024)},
		documents:  documents,
	}
	return data
}

func RenderDropdownMenuItem(text string) {
	clay.UI()(clay.ElementDeclaration{
		Layout: clay.LayoutConfig{
			Padding: clay.PaddingAll(16),
		},
	}, func() {
		clay.Text(text, clay.TextConfig(clay.TextElementConfig{
			FontId:    FontIdBody16,
			FontSize:  16,
			TextColor: clay.Color{R: 255, G: 255, B: 255, A: 255},
		}))
	})
}

func RenderHeaderButton(text string) {
	clay.UI()(clay.ElementDeclaration{
		Layout: clay.LayoutConfig{
			Padding: clay.Padding{Left: 16, Right: 16, Top: 8, Bottom: 8},
		},
		Border: func() clay.BorderElementConfig {
			if !clay.Hovered() {
				return clay.BorderElementConfig{}
			}
			return clay.BorderElementConfig{
				Color: clay.Color{R: 81, G: 162, B: 255, A: 255},
				Width: clay.BorderAll(4),
			}
		}(),
		BackgroundColor: clay.Color{R: 140, G: 140, B: 140, A: 255},
		CornerRadius:    clay.CornerRadiusAll(5),
	}, func() {
		clay.Text(text, clay.TextConfig(clay.TextElementConfig{
			FontId:    FontIdBody16,
			FontSize:  16,
			TextColor: clay.Color{R: 255, G: 255, B: 255, A: 255},
		}))
	})
}

func handleSidebarInteraction(elementId clay.ElementId, pointerData clay.PointerData, userData int64) {
	clickData := (*sidebarClickData)(unsafe.Pointer(uintptr(userData)))
	// If this button was clicked
	if pointerData.State == clay.POINTER_DATA_PRESSED_THIS_FRAME {
		if clickData.requestedDocumentIndex >= 0 && clickData.requestedDocumentIndex < clickData.documentLen {
			// Select the corresponding document
			*clickData.selectedDocumentIndex = clickData.requestedDocumentIndex
		}
	}
}

func CreateLayout(data *Data) clay.RenderCommandArray {
	data.frameArena.offset = 0

	clay.BeginLayout()

	layoutExpand := clay.Sizing{
		Width:  clay.SizingGrow(0),
		Height: clay.SizingGrow(0),
	}

	contentBackgroundColor := clay.Color{R: 90, G: 90, B: 90, A: 255}

	clay.UI()(clay.ElementDeclaration{
		Id:              clay.ID("OuterContainer"),
		BackgroundColor: clay.Color{R: 43, G: 41, B: 51, A: 255},
		Layout: clay.LayoutConfig{
			LayoutDirection: clay.TOP_TO_BOTTOM,
			Sizing:          layoutExpand,
			Padding:         clay.PaddingAll(16),
			ChildGap:        16,
		},
	}, func() {
		clay.UI()(clay.ElementDeclaration{
			Id: clay.ID("HeaderBar"),
			Layout: clay.LayoutConfig{
				Sizing: clay.Sizing{
					Height: clay.SizingFixed(60),
					Width:  clay.SizingGrow(0),
				},
				Padding:  clay.Padding{Left: 16, Right: 16},
				ChildGap: 16,
				ChildAlignment: clay.ChildAlignment{
					Y: clay.ALIGN_Y_CENTER,
				},
			},
			BackgroundColor: contentBackgroundColor,
			CornerRadius:    clay.CornerRadiusAll(5),
		}, func() {
			clay.UI()(clay.ElementDeclaration{
				Id:              clay.ID("FileButton"),
				Layout:          clay.LayoutConfig{Padding: clay.Padding{Left: 16, Right: 16, Top: 8, Bottom: 8}},
				BackgroundColor: clay.Color{R: 150, G: 150, B: 150, A: 255},
				CornerRadius:    clay.CornerRadiusAll(5),
			}, func() {
				clay.Text("File", clay.TextConfig(clay.TextElementConfig{
					FontId:    FontIdBody16,
					FontSize:  16,
					TextColor: clay.Color{R: 255, G: 255, B: 255, A: 255},
				}))

				fileMenuVisible := clay.PointerOver(clay.GetElementId("FileButton")) ||
					clay.PointerOver(clay.GetElementId("FileMenu"))

				if fileMenuVisible { // Below has been changed slightly to fix the small bug where the menu would dismiss when mousing over the top gap
					clay.UI()(clay.ElementDeclaration{
						Id: clay.ID("FileMenu"),
						Floating: clay.FloatingElementConfig{
							AttachTo: clay.ATTACH_TO_PARENT,
							AttachPoints: clay.FloatingAttachPoints{
								Parent: clay.ATTACH_POINT_LEFT_BOTTOM,
							},
						},
						Layout: clay.LayoutConfig{Padding: clay.Padding{Top: 8, Bottom: 8}},
					}, func() {
						clay.UI()(clay.ElementDeclaration{
							Layout: clay.LayoutConfig{
								LayoutDirection: clay.TOP_TO_BOTTOM,
								Sizing: clay.Sizing{
									Width: clay.SizingFixed(200),
								},
							},
							BackgroundColor: clay.Color{R: 40, G: 40, B: 40, A: 255},
							CornerRadius:    clay.CornerRadiusAll(8),
						}, func() {
							// Render dropdown items here
							RenderDropdownMenuItem("New")
							RenderDropdownMenuItem("Open")
							RenderDropdownMenuItem("Close")
						})
					})
				}
			})
			RenderHeaderButton("Edit")
			clay.UI()(clay.ElementDeclaration{
				Layout: clay.LayoutConfig{
					Sizing: clay.Sizing{
						Width: clay.SizingGrow(0),
					},
				},
			}, func() {})
			RenderHeaderButton("Upload")
			RenderHeaderButton("Media")
			RenderHeaderButton("Support")
		})

		clay.UI()(clay.ElementDeclaration{
			Id:     clay.ID("LowerContent"),
			Layout: clay.LayoutConfig{Sizing: layoutExpand, ChildGap: 16},
		}, func() {
			clay.UI()(clay.ElementDeclaration{
				Id:              clay.ID("SideBar"),
				BackgroundColor: contentBackgroundColor,
				Layout: clay.LayoutConfig{
					LayoutDirection: clay.TOP_TO_BOTTOM,
					Padding:         clay.PaddingAll(16),
					ChildGap:        8,
					Sizing: clay.Sizing{
						Width:  clay.SizingFixed(250),
						Height: clay.SizingGrow(0),
					},
				},
			}, func() {
				for i, document := range data.documents {
					sidebarButtonlayout := clay.LayoutConfig{
						Sizing:  clay.Sizing{Width: clay.SizingGrow(0)},
						Padding: clay.PaddingAll(16),
					}

					if i == int(data.selectedDocumentIndex) {
						clay.UI()(clay.ElementDeclaration{
							Layout:          sidebarButtonlayout,
							BackgroundColor: clay.Color{R: 120, G: 120, B: 120, A: 255},
							CornerRadius:    clay.CornerRadiusAll(8),
						}, func() {
							clay.Text(document.title, clay.TextConfig(clay.TextElementConfig{
								FontId:    FontIdBody16,
								FontSize:  20,
								TextColor: clay.Color{R: 255, G: 255, B: 255, A: 255},
							}))
						})
					} else {
						clickData := alloc[sidebarClickData](&data.frameArena)
						*clickData = sidebarClickData{
							requestedDocumentIndex: int32(i),
							selectedDocumentIndex:  &data.selectedDocumentIndex,
							documentLen:            int32(len(data.documents)),
						}
						clay.UI()(clay.ElementDeclaration{
							Layout: sidebarButtonlayout,
							BackgroundColor: clay.Color{R: 120, G: 120, B: 120, A: func() float32 {
								if clay.Hovered() {
									return 120
								} else {
									return 0
								}
							}()},
							CornerRadius: clay.CornerRadiusAll(8),
						}, func() {
							clay.OnHover(handleSidebarInteraction, int64(uintptr(unsafe.Pointer(clickData))))
							clay.Text(document.title, clay.TextConfig(clay.TextElementConfig{
								FontId:    FontIdBody16,
								FontSize:  20,
								TextColor: clay.Color{R: 255, G: 255, B: 255, A: 255},
							}))
						})
					}
				}
			})
			clay.UI()(clay.ElementDeclaration{
				Id:              clay.ID("MainContent"),
				BackgroundColor: contentBackgroundColor,
				Clip:            clay.ClipElementConfig{Vertical: true, ChildOffset: clay.GetScrollOffset()},
				Layout: clay.LayoutConfig{
					LayoutDirection: clay.TOP_TO_BOTTOM,
					ChildGap:        16,
					Padding:         clay.PaddingAll(16),
					Sizing:          layoutExpand,
				},
			}, func() {
				selectedDocument := data.documents[data.selectedDocumentIndex]
				clay.Text(selectedDocument.title, clay.TextConfig(clay.TextElementConfig{
					FontId:    FontIdBody16,
					FontSize:  24,
					TextColor: white,
				}))
				if selectedDocument.image != nil {
					clay.UI()(clay.ElementDeclaration{
						Layout: clay.LayoutConfig{
							Sizing: clay.Sizing{
								Width:  clay.SizingFixed(160),
								Height: clay.SizingFixed(90),
							},
						},
						AspectRatio: clay.AspectRatioElementConfig{
							AspectRatio: 1280.0 / 853.0,
						},
						Image: clay.ImageElementConfig{ImageData: selectedDocument.image.(unsafe.Pointer)},
					}, func() {
					})
				}
				clay.Text(selectedDocument.contents, clay.TextConfig(clay.TextElementConfig{
					FontId:    FontIdBody16,
					FontSize:  24,
					TextColor: white,
				}))
			})
		})
	})

	renderCommands := clay.EndLayout()
	for i := int32(0); i < renderCommands.Length; i++ {
		clay.RenderCommandArray_Get(&renderCommands, i).BoundingBox.Y += data.yOffset
	}
	return renderCommands
}
