package shared_layouts

import (
	"github.com/totallygamerjet/clay"
)

const FONT_ID_BODY_16 = 0

type Document struct {
	Title    clay.String
	Contents clay.String
}
type DocumentArray struct {
	Documents []Document
	Length    uint32
}

var documentsRaw [5]Document
var documents DocumentArray = DocumentArray{Length: 5, Documents: documentsRaw[:]}

type ClayVideoDemo_Arena struct {
	Offset int64
	Memory []byte
}
type ClayVideoDemo_Data struct {
	SelectedDocumentIndex int32
	YOffset               float32
	FrameArena            ClayVideoDemo_Arena
}
type SidebarClickData struct {
	RequestedDocumentIndex int32
	SelectedDocumentIndex  *int32
}

func ClayVideoDemo_Initialize() ClayVideoDemo_Data {
	documents.Documents[0] = Document{Title: clay.ToString("Squirrels"), Contents: clay.ToString("The Secret Life of Squirrels: Nature's Clever Acrobats\n" + "Squirrels are often overlooked creatures, dismissed as mere park inhabitants or backyard nuisances. Yet, beneath their fluffy tails and twitching noses lies an intricate world of cunning, agility, and survival tactics that are nothing short of fascinating. As one of the most common mammals in North America, squirrels have adapted to a wide range of environments from bustling urban centers to tranquil forests and have developed a variety of unique behaviors that continue to intrigue scientists and nature enthusiasts alike.\n" + "\n" + "Master Tree Climbers\n" + "At the heart of a squirrel's skill set is its impressive ability to navigate trees with ease. Whether they're darting from branch to branch or leaping across wide gaps, squirrels possess an innate talent for acrobatics. Their powerful hind legs, which are longer than their front legs, give them remarkable jumping power. With a tail that acts as a counterbalance, squirrels can leap distances of up to ten times the length of their body, making them some of the best aerial acrobats in the animal kingdom.\n" + "But it's not just their agility that makes them exceptional climbers. Squirrels' sharp, curved claws allow them to grip tree bark with precision, while the soft pads on their feet provide traction on slippery surfaces. Their ability to run at high speeds and scale vertical trunks with ease is a testament to the evolutionary adaptations that have made them so successful in their arboreal habitats.\n" + "\n" + "Food Hoarders Extraordinaire\n" + "Squirrels are often seen frantically gathering nuts, seeds, and even fungi in preparation for winter. While this behavior may seem like instinctual hoarding, it is actually a survival strategy that has been honed over millions of years. Known as \"scatter hoarding,\" squirrels store their food in a variety of hidden locations, often burying it deep in the soil or stashing it in hollowed-out tree trunks.\n" + "Interestingly, squirrels have an incredible memory for the locations of their caches. Research has shown that they can remember thousands of hiding spots, often returning to them months later when food is scarce. However, they don't always recover every stash some forgotten caches eventually sprout into new trees, contributing to forest regeneration. This unintentional role as forest gardeners highlights the ecological importance of squirrels in their ecosystems.\n" + "\n" + "The Great Squirrel Debate: Urban vs. Wild\n" + "While squirrels are most commonly associated with rural or wooded areas, their adaptability has allowed them to thrive in urban environments as well. In cities, squirrels have become adept at finding food sources in places like parks, streets, and even garbage cans. However, their urban counterparts face unique challenges, including traffic, predators, and the lack of natural shelters. Despite these obstacles, squirrels in urban areas are often observed using human infrastructure such as buildings, bridges, and power lines as highways for their acrobatic escapades.\n" + "There is, however, a growing concern regarding the impact of urban life on squirrel populations. Pollution, deforestation, and the loss of natural habitats are making it more difficult for squirrels to find adequate food and shelter. As a result, conservationists are focusing on creating squirrel-friendly spaces within cities, with the goal of ensuring these resourceful creatures continue to thrive in both rural and urban landscapes.\n" + "\n" + "A Symbol of Resilience\n" + "In many cultures, squirrels are symbols of resourcefulness, adaptability, and preparation. Their ability to thrive in a variety of environments while navigating challenges with agility and grace serves as a reminder of the resilience inherent in nature. Whether you encounter them in a quiet forest, a city park, or your own backyard, squirrels are creatures that never fail to amaze with their endless energy and ingenuity.\n" + "In the end, squirrels may be small, but they are mighty in their ability to survive and thrive in a world that is constantly changing. So next time you spot one hopping across a branch or darting across your lawn, take a moment to appreciate the remarkable acrobat at work a true marvel of the natural world.\n")}
	documents.Documents[1] = Document{Title: clay.ToString("Lorem Ipsum"), Contents: clay.ToString("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")}
	documents.Documents[2] = Document{Title: clay.ToString("Vacuum Instructions"), Contents: clay.ToString("Chapter 3: Getting Started - Unpacking and Setup\n" + "\n" + "Congratulations on your new SuperClean Pro 5000 vacuum cleaner! In this section, we will guide you through the simple steps to get your vacuum up and running. Before you begin, please ensure that you have all the components listed in the \"Package Contents\" section on page 2.\n" + "\n" + "1. Unboxing Your Vacuum\n" + "Carefully remove the vacuum cleaner from the box. Avoid using sharp objects that could damage the product. Once removed, place the unit on a flat, stable surface to proceed with the setup. Inside the box, you should find:\n" + "\n" + "    The main vacuum unit\n" + "    A telescoping extension wand\n" + "    A set of specialized cleaning tools (crevice tool, upholstery brush, etc.)\n" + "    A reusable dust bag (if applicable)\n" + "    A power cord with a 3-prong plug\n" + "    A set of quick-start instructions\n" + "\n" + "2. Assembling Your Vacuum\n" + "Begin by attaching the extension wand to the main body of the vacuum cleaner. Line up the connectors and twist the wand into place until you hear a click. Next, select the desired cleaning tool and firmly attach it to the wand's end, ensuring it is securely locked in.\n" + "\n" + "For models that require a dust bag, slide the bag into the compartment at the back of the vacuum, making sure it is properly aligned with the internal mechanism. If your vacuum uses a bagless system, ensure the dust container is correctly seated and locked in place before use.\n" + "\n" + "3. Powering On\n" + "To start the vacuum, plug the power cord into a grounded electrical outlet. Once plugged in, locate the power switch, usually positioned on the side of the handle or body of the unit, depending on your model. Press the switch to the \"On\" position, and you should hear the motor begin to hum. If the vacuum does not power on, check that the power cord is securely plugged in, and ensure there are no blockages in the power switch.\n" + "\n" + "Note: Before first use, ensure that the vacuum filter (if your model has one) is properly installed. If unsure, refer to \"Section 5: Maintenance\" for filter installation instructions.")}
	documents.Documents[3] = Document{Title: clay.ToString("Article 4"), Contents: clay.ToString("Article 4")}
	documents.Documents[4] = Document{Title: clay.ToString("Article 5"), Contents: clay.ToString("Article 5")}

	var data = ClayVideoDemo_Data{
		FrameArena: ClayVideoDemo_Arena{Memory: make([]byte, 1024)},
	}
	return data
}

func ClayVideoDemo_CreateLayout(data *ClayVideoDemo_Data) {
	data.FrameArena.Offset = 0

	clay.BeginLayout()

	layoutExpand := clay.Sizing{
		Width:  clay.SizingGrow(0),
		Height: clay.SizingGrow(0),
	}

	contentBackgroundColor := clay.Color{R: 90, G: 90, B: 90, A: 255}

	clay.UI(clay.ElementDeclaration{
		Id:              clay.ID("OuterContainer"),
		BackgroundColor: clay.Color{43, 41, 51, 255},
		Layout: clay.LayoutConfig{
			LayoutDirection: clay.TOP_TO_BOTTOM,
			Sizing:          layoutExpand,
			Padding:         clay.PaddingAll(16),
			ChildGap:        16,
		},
	}, func() {
		clay.UI(clay.ElementDeclaration{
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
			clay.UI(clay.ElementDeclaration{
				Id:              clay.ID("FileButton"),
				Layout:          clay.LayoutConfig{Padding: clay.Padding{Left: 16, Right: 16, Top: 8, Bottom: 8}},
				BackgroundColor: clay.Color{R: 150, G: 150, B: 150, A: 255},
				CornerRadius:    clay.CornerRadiusAll(5),
			}, func() {
				clay.Text(clay.ToString("File"), clay.TextConfig(clay.TextElementConfig{
					FontId:    FONT_ID_BODY_16,
					FontSize:  16,
					TextColor: clay.Color{R: 255, G: 255, B: 255, A: 255},
				}))

				fileMenuVisible := clay.PointerOver(clay.GetElementId(clay.ToString("FileButton"))) ||
					clay.PointerOver(clay.GetElementId(clay.ToString("FileMenu")))

				if fileMenuVisible { // Below has been changed slightly to fix the small bug where the menu would dismiss when mousing over the top gap
					clay.UI(clay.ElementDeclaration{
						Id: clay.ID("FileMenu"),
						Floating: clay.FloatingElementConfig{
							AttachTo: clay.ATTACH_TO_PARENT,
							AttachPoints: clay.FloatingAttachPoints{
								Parent: clay.ATTACH_POINT_LEFT_BOTTOM,
							},
						},
						Layout: clay.LayoutConfig{Padding: clay.Padding{Top: 8, Bottom: 8}},
					}, func() {
						clay.UI(clay.ElementDeclaration{
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
							RenderDropdownMenuItem(clay.ToString("New"))
							RenderDropdownMenuItem(clay.ToString("Open"))
							RenderDropdownMenuItem(clay.ToString("Close"))
						})
					})
				}

				RenderHeaderButton(clay.ToString("Edit"))
				clay.UI(clay.ElementDeclaration{
					Layout: clay.LayoutConfig{
						Sizing: clay.Sizing{
							Width: clay.SizingGrow(0),
						},
					},
				}, func() {})
				RenderHeaderButton(clay.ToString("Upload"))
				RenderHeaderButton(clay.ToString("Media"))
				RenderHeaderButton(clay.ToString("Support"))
			})
		})

		clay.UI(clay.ElementDeclaration{
			Id:     clay.ID("LowerContent"),
			Layout: clay.LayoutConfig{Sizing: layoutExpand, ChildGap: 16},
		}, func() {
			// TODO: starting here...
		})
	})
}

//Clay_RenderCommandArray ClayVideoDemo_CreateLayout(ClayVideoDemo_Data *data) {
//    CLAY({ .id = CLAY_ID("OuterContainer"),
//    }) {
//        // Child elements go inside braces
//
//        CLAY({
//            .id = CLAY_ID("LowerContent"),
//            .layout = { .sizing = layoutExpand, .childGap = 16 }
//        }) {
//            CLAY({
//                .id = CLAY_ID("Sidebar"),
//                .backgroundColor = contentBackgroundColor,
//                .layout = {
//                    .layoutDirection = CLAY_TOP_TO_BOTTOM,
//                    .padding = CLAY_PADDING_ALL(16),
//                    .childGap = 8,
//                    .sizing = {
//                        .width = CLAY_SIZING_FIXED(250),
//                        .height = CLAY_SIZING_GROW(0)
//                    }
//                }
//            }) {
//                for (int i = 0; i < documents.length; i++) {
//                    Document document = documents.documents[i];
//                    Clay_LayoutConfig sidebarButtonLayout = {
//                        .sizing = { .width = CLAY_SIZING_GROW(0) },
//                        .padding = CLAY_PADDING_ALL(16)
//                    };
//
//                    if (i == data->selectedDocumentIndex) {
//                        CLAY({
//                            .layout = sidebarButtonLayout,
//                            .backgroundColor = {120, 120, 120, 255 },
//                            .cornerRadius = CLAY_CORNER_RADIUS(8)
//                        }) {
//                            CLAY_TEXT(document.title, CLAY_TEXT_CONFIG({
//                                .fontId = FONT_ID_BODY_16,
//                                .fontSize = 20,
//                                .textColor = { 255, 255, 255, 255 }
//                            }));
//                        }
//                    } else {
//                        SidebarClickData *clickData = (SidebarClickData *)(data->frameArena.memory + data->frameArena.offset);
//                        *clickData = (SidebarClickData) { .requestedDocumentIndex = i, .selectedDocumentIndex = &data->selectedDocumentIndex };
//                        data->frameArena.offset += sizeof(SidebarClickData);
//                        CLAY({ .layout = sidebarButtonLayout, .backgroundColor = (Clay_Color) { 120, 120, 120, Clay_Hovered() ? 120 : 0 }, .cornerRadius = CLAY_CORNER_RADIUS(8) }) {
//                            Clay_OnHover(HandleSidebarInteraction, (intptr_t)clickData);
//                            CLAY_TEXT(document.title, CLAY_TEXT_CONFIG({
//                                .fontId = FONT_ID_BODY_16,
//                                .fontSize = 20,
//                                .textColor = { 255, 255, 255, 255 }
//                            }));
//                        }
//                    }
//                }
//            }
//
//            CLAY({ .id = CLAY_ID("MainContent"),
//                .backgroundColor = contentBackgroundColor,
//                .scroll = { .vertical = true },
//                .layout = {
//                    .layoutDirection = CLAY_TOP_TO_BOTTOM,
//                    .childGap = 16,
//                    .padding = CLAY_PADDING_ALL(16),
//                    .sizing = layoutExpand
//                }
//            }) {
//                Document selectedDocument = documents.documents[data->selectedDocumentIndex];
//                CLAY_TEXT(selectedDocument.title, CLAY_TEXT_CONFIG({
//                    .fontId = FONT_ID_BODY_16,
//                    .fontSize = 24,
//                    .textColor = COLOR_WHITE
//                }));
//                CLAY_TEXT(selectedDocument.contents, CLAY_TEXT_CONFIG({
//                    .fontId = FONT_ID_BODY_16,
//                    .fontSize = 24,
//                    .textColor = COLOR_WHITE
//                }));
//            }
//        }
//    }
//
//    Clay_RenderCommandArray renderCommands = Clay_EndLayout();
//    for (int32_t i = 0; i < renderCommands.length; i++) {
//        Clay_RenderCommandArray_Get(&renderCommands, i)->boundingBox.y += data->yOffset;
//    }
//    return renderCommands;
//}
