package clay

import (
	"strings"
	"testing"
	"unsafe"
)

// newContext returns a fresh clay context with a text measuring function that
// makes every character 10 by 16 pixels, so that layouts are predictable.
func newContext(t *testing.T) *Context {
	t.Helper()
	ctx := Initialize(CreateArenaWithCapacity(MinMemorySize()), Dimensions{Width: 800, Height: 600},
		ErrorHandler{ErrorHandlerFunction: func(e ErrorData) { t.Errorf("clay: %v", e) }})
	SetMeasureTextFunction(func(text string, config *TextElementConfig, _ any) Dimensions {
		return Dimensions{Width: float32(len(text)) * 10, Height: 16}
	}, nil)
	return ctx
}

func TestStringsRoundTrip(t *testing.T) {
	newContext(t)
	BeginLayout()

	const s = "the quick brown fox"
	addr := storeString(s)

	// A string that was stored is returned without copying it back out of the module.
	got := loadString(wasmMemory(), addr, uint32(len(s)))
	if got != s {
		t.Errorf("loadString = %q, want %q", got, s)
	}
	if unsafe.StringData(got) != unsafe.StringData(s) {
		t.Error("loadString copied the string instead of returning the stored one")
	}

	// Clay hands back slices of the strings it was given, like the words of wrapped text.
	if got := loadString(wasmMemory(), addr+4, 5); got != s[4:9] {
		t.Errorf("loadString of a slice = %q, want %q", got, s[4:9])
	}

	// Anything else, like clay's own static strings, is copied out of the module.
	const other = "not stored"
	copy(wasmMemory()[scratchAddr:], other)
	if got := loadString(wasmMemory(), scratchAddr, uint32(len(other))); got != other {
		t.Errorf("loadString of unstored memory = %q, want %q", got, other)
	}

	if got := loadString(wasmMemory(), 0, 0); got != "" {
		t.Errorf("loadString of an empty string = %q", got)
	}
}

// TestStringBuffer checks the buffers strings are copied into, which grow with more
// chunks when a frame needs them and are merged back into one when the buffer is reused.
func TestStringBuffer(t *testing.T) {
	var b stringBuffer
	first := b.alloc(minStringChunk - 16)
	second := b.alloc(16)
	if len(b.chunks) != 1 {
		t.Fatalf("allocating within a chunk used %d chunks", len(b.chunks))
	}
	if second != first+minStringChunk-16 {
		t.Errorf("second allocation at %d, want it after the first at %d", second, first)
	}

	// More than the chunk holds needs another chunk, which must not overlap the first.
	third := b.alloc(minStringChunk)
	if len(b.chunks) != 2 {
		t.Fatalf("allocating past the end of a chunk used %d chunks", len(b.chunks))
	}
	if third >= first && third < first+minStringChunk {
		t.Errorf("the second chunk at %d overlaps the first at %d", third, first)
	}
	total := b.chunks[0].size + b.chunks[1].size

	// Reusing the buffer merges the chunks, so that a frame of the same size fits in one.
	b.reset()
	if len(b.chunks) != 1 {
		t.Fatalf("reset left %d chunks", len(b.chunks))
	}
	if b.chunks[0].size < total || b.chunks[0].used != 0 {
		t.Errorf("reset chunk = %+v, want an empty chunk of at least %d bytes", b.chunks[0], total)
	}
	if len(b.strings) != 0 {
		t.Errorf("reset kept %d stored strings", len(b.strings))
	}
}

// TestStringBufferStopsGrowing checks that laying out repeatedly does not keep allocating.
func TestStringBufferStopsGrowing(t *testing.T) {
	newContext(t)
	var sizes []int
	for frame := range 40 {
		BeginLayout()
		n := 10
		if frame%4 == 0 {
			n = 400
		}
		for i := range n {
			storeString(strings.Repeat("x", 500) + string(rune('a'+i%26)))
		}
		sizes = append(sizes, len(wasmMemory()))
	}
	if sizes[len(sizes)-1] != sizes[len(sizes)/2] {
		t.Errorf("memory kept growing: %v", sizes)
	}
}

func TestHandleLifetime(t *testing.T) {
	newContext(t)
	BeginLayout()

	value := new(int)
	h := storeHandle(value)
	if got := loadHandle(h); got != any(value) {
		t.Errorf("loadHandle in the same frame = %v", got)
	}

	// Handles stay valid for one more frame, so that pointer events can be handled
	// before the next layout begins.
	BeginLayout()
	if got := loadHandle(h); got != any(value) {
		t.Errorf("loadHandle in the next frame = %v", got)
	}
	BeginLayout()
	if got := loadHandle(h); got != nil {
		t.Errorf("loadHandle two frames later = %v, want nil", got)
	}

	// Handles for things that outlive a frame, like the error handler, are always valid.
	p := storePersistentHandle(value)
	for range 4 {
		BeginLayout()
	}
	if got := loadHandle(p); got != any(value) {
		t.Errorf("loadHandle of a persistent handle = %v", got)
	}

	if got := storeHandle(nil); got != 0 {
		t.Errorf("storeHandle(nil) = %d, want 0", got)
	}
	for _, h := range []uint32{0, 1 << 29, persistentHandle | 1<<20} {
		if got := loadHandle(h); got != nil {
			t.Errorf("loadHandle(%d) = %v, want nil", h, got)
		}
	}
}

func TestPointer(t *testing.T) {
	newContext(t)

	// Any address inside the module works for a round trip; use the scratch buffer.
	small := Pointer[Vector2]{addr: scratchAddr}
	small.Set(Vector2{X: 3, Y: 4})
	if got := small.Get(); got != (Vector2{X: 3, Y: 4}) {
		t.Errorf("Vector2 round trip = %v", got)
	}

	// TransitionData is larger than a Vector2, and used to overflow the buffer Set encodes into.
	want := TransitionData{
		BoundingBox:     BoundingBox{X: 1, Y: 2, Width: 3, Height: 4},
		BackgroundColor: Color{R: 5, G: 6, B: 7, A: 8},
		OverlayColor:    Color{R: 9, G: 10, B: 11, A: 12},
		BorderColor:     Color{R: 13, G: 14, B: 15, A: 16},
		BorderWidth:     BorderWidth{Left: 17, Right: 18, Top: 19, Bottom: 20, BetweenChildren: 21},
	}
	big := Pointer[TransitionData]{addr: scratchAddr}
	big.Set(want)
	if got := big.Get(); got != want {
		t.Errorf("TransitionData round trip = %+v, want %+v", got, want)
	}
	if sizeofTransitionData > maxPointerSize {
		t.Errorf("sizeofTransitionData = %d, larger than maxPointerSize %d", sizeofTransitionData, maxPointerSize)
	}

	var nilPointer Pointer[Vector2]
	if !nilPointer.IsNil() {
		t.Error("the zero Pointer is not nil")
	}
	if got := nilPointer.Get(); got != (Vector2{}) {
		t.Errorf("Get of a nil Pointer = %v", got)
	}
}

// TestRenderCommands checks that the fields of an element declaration reach clay,
// and that each render command is decoded as the member of the union its type selects.
func TestRenderCommands(t *testing.T) {
	newContext(t)

	background := Color{R: 10, G: 20, B: 30, A: 255}
	borderColor := Color{R: 40, G: 50, B: 60, A: 255}
	image := new(int)

	BeginLayout()
	UI(ID("Root"))(ElementDeclaration{
		Layout:          LayoutConfig{Sizing: Sizing{Width: SizingFixed(400), Height: SizingFixed(300)}, LayoutDirection: TOP_TO_BOTTOM},
		BackgroundColor: background,
		CornerRadius:    CornerRadiusAll(7),
		Border:          BorderElementConfig{Color: borderColor, Width: BorderOutside(3)},
	}, func() {
		Text("hello", &TextElementConfig{FontId: 0, FontSize: 24, TextColor: background})
		UI(ID("Image"))(ElementDeclaration{
			Layout: LayoutConfig{Sizing: Sizing{Width: SizingFixed(50), Height: SizingFixed(50)}},
			Image:  ImageElementConfig{ImageData: image},
		}, nil)
	})
	commands := EndLayout(0)

	var seen []RenderCommandType
	for _, c := range commands {
		seen = append(seen, c.CommandType)
		switch c.CommandType {
		case RENDER_COMMAND_TYPE_RECTANGLE:
			if c.RenderData.Rectangle.BackgroundColor != background {
				t.Errorf("rectangle color = %v, want %v", c.RenderData.Rectangle.BackgroundColor, background)
			}
			if c.RenderData.Rectangle.CornerRadius != CornerRadiusAll(7) {
				t.Errorf("rectangle corner radius = %v", c.RenderData.Rectangle.CornerRadius)
			}
			// The other members of the union must be left alone, not decoded from the
			// rectangle's bytes, which would read colors as string pointers and user data.
			if c.RenderData.Text != (TextRenderData{}) {
				t.Errorf("rectangle decoded as text: %+v", c.RenderData.Text)
			}
			if c.RenderData.Image != (ImageRenderData{}) {
				t.Errorf("rectangle decoded as an image: %+v", c.RenderData.Image)
			}
			if c.RenderData.Custom != (CustomRenderData{}) {
				t.Errorf("rectangle decoded as custom: %+v", c.RenderData.Custom)
			}
		case RENDER_COMMAND_TYPE_TEXT:
			if got := c.RenderData.Text; got.StringContents != "hello" || got.FontSize != 24 || got.TextColor != background {
				t.Errorf("text render data = %+v", got)
			}
		case RENDER_COMMAND_TYPE_IMAGE:
			if c.RenderData.Image.ImageData != any(image) {
				t.Errorf("image data = %v, want the value passed in", c.RenderData.Image.ImageData)
			}
		case RENDER_COMMAND_TYPE_BORDER:
			if got := c.RenderData.Border; got.Color != borderColor || got.Width != BorderOutside(3) {
				t.Errorf("border render data = %+v", got)
			}
		}
	}

	for _, want := range []RenderCommandType{
		RENDER_COMMAND_TYPE_RECTANGLE, RENDER_COMMAND_TYPE_TEXT,
		RENDER_COMMAND_TYPE_IMAGE, RENDER_COMMAND_TYPE_BORDER,
	} {
		if !contains(seen, want) {
			t.Errorf("no %v command in %v", want, seen)
		}
	}
}

func contains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func TestMeasureTextCallback(t *testing.T) {
	Initialize(CreateArenaWithCapacity(MinMemorySize()), Dimensions{Width: 800, Height: 600},
		ErrorHandler{ErrorHandlerFunction: func(e ErrorData) { t.Errorf("clay: %v", e) }})

	userData := new(int)
	var texts []string
	var config TextElementConfig
	SetMeasureTextFunction(func(text string, c *TextElementConfig, u any) Dimensions {
		if u != any(userData) {
			t.Errorf("measure text user data = %v, want the value passed to SetMeasureTextFunction", u)
		}
		texts = append(texts, text)
		config = *c
		return Dimensions{Width: float32(len(text)) * 10, Height: 16}
	}, userData)

	BeginLayout()
	UI(ID("Root"))(ElementDeclaration{
		Layout: LayoutConfig{Sizing: Sizing{Width: SizingFixed(400), Height: SizingFixed(300)}},
	}, func() {
		Text("measure me", &TextElementConfig{FontSize: 33, LetterSpacing: 2})
	})
	EndLayout(0)

	if !contains(texts, "measure") || !contains(texts, "me") {
		t.Errorf("measured %q, want the words of the text", texts)
	}
	if config.FontSize != 33 || config.LetterSpacing != 2 {
		t.Errorf("measured with config %+v, want the config the text was declared with", config)
	}
}

func TestErrorHandlerCallback(t *testing.T) {
	var errors []ErrorData
	handlerData := new(int)
	Initialize(CreateArenaWithCapacity(MinMemorySize()), Dimensions{Width: 800, Height: 600},
		ErrorHandler{
			ErrorHandlerFunction: func(e ErrorData) { errors = append(errors, e) },
			UserData:             handlerData,
		})
	// Clay reports an error when it has to measure text without a measuring function.
	SetMeasureTextFunction(nil, nil)

	BeginLayout()
	UI(ID("Root"))(ElementDeclaration{}, func() {
		Text("text needs measuring", nil)
	})
	EndLayout(0)

	if len(errors) == 0 {
		t.Fatal("no error reported for a missing text measuring function")
	}
	e := errors[0]
	if e.ErrorType != ERROR_TYPE_TEXT_MEASUREMENT_FUNCTION_NOT_PROVIDED {
		t.Errorf("error type = %v", e.ErrorType)
	}
	if e.ErrorText == "" {
		t.Error("error text is empty")
	}
	if e.UserData != any(handlerData) {
		t.Errorf("error user data = %v, want the value passed to Initialize", e.UserData)
	}
	if e.Error() == "" {
		t.Error("ErrorData.Error is empty")
	}
}

func TestOnHoverCallback(t *testing.T) {
	newContext(t)
	SetLayoutDimensions(Dimensions{Width: 800, Height: 600})

	userData := new(int)
	var hovered []ElementId
	var state PointerDataInteractionState
	layout := func() {
		BeginLayout()
		UI(ID("Button"))(ElementDeclaration{
			Layout: LayoutConfig{Sizing: Sizing{Width: SizingFixed(200), Height: SizingFixed(100)}},
		}, func() {
			OnHover(func(id ElementId, data PointerData, u any) {
				if u != any(userData) {
					t.Errorf("hover user data = %v, want the value passed to OnHover", u)
				}
				hovered = append(hovered, id)
				state = data.State
			}, userData)
		})
		EndLayout(0)
	}

	// The callback is bound while laying out, and called by SetPointerState afterwards.
	layout()
	SetPointerState(Vector2{X: 50, Y: 50}, true)
	if len(hovered) != 1 || hovered[0].Id != ID("Button").Id {
		t.Errorf("hovered = %v, want the button", hovered)
	}
	if state != POINTER_DATA_PRESSED_THIS_FRAME {
		t.Errorf("pointer state = %v", state)
	}

	// Outside the element, it is not called.
	hovered = nil
	layout()
	SetPointerState(Vector2{X: 700, Y: 500}, false)
	if len(hovered) != 0 {
		t.Errorf("hovered = %v outside the element", hovered)
	}
}

func TestElementIdStrings(t *testing.T) {
	newContext(t)
	BeginLayout()

	id := ID("Sidebar")
	if id.StringId != "Sidebar" {
		t.Errorf("ID(%q).StringId = %q", "Sidebar", id.StringId)
	}
	if got := GetElementId("Sidebar"); got.Id != id.Id {
		t.Errorf("GetElementId = %d, want %d", got.Id, id.Id)
	}
	if IDI("Item", 3).Id == IDI("Item", 4).Id {
		t.Error("IDI returns the same id for different indexes")
	}
}
