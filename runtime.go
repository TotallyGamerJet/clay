package clay

import (
	"cmp"
	"encoding/binary"
	"math"
	"slices"

	"github.com/TotallyGamerJet/clay/internal/wasm"
)

// Clay runs as a WebAssembly module translated to Go (see internal/wasm).
// All of its state lives in the module's linear memory, so values are copied
// in and out of that memory by the functions generated in clay.go.

var (
	module      *wasm.Module
	memory      *[]byte
	scratchAddr uint32 // scratch buffer inside the module used to pass structs
	argBuffer   []byte
)

func init() {
	module = wasm.New(host{})
	memory = module.Xmemory().Slice()
	scratchAddr = uint32(module.Xgo_scratch())
}

// wasmMemory returns the module's memory.
// The returned slice is invalidated by calls into the module, as they might grow the memory.
func wasmMemory() []byte {
	return *memory
}

// wasmArgs returns a zeroed buffer to encode arguments in, which is then passed to wasmCommit.
// Arguments are encoded outside the module's memory because encoding strings might grow it.
func wasmArgs(n uint32) []byte {
	if uint32(cap(argBuffer)) < n {
		argBuffer = make([]byte, n)
	}
	b := argBuffer[:n]
	clear(b)
	return b
}

// wasmCommit copies the arguments into the module's scratch buffer and returns its address.
func wasmCommit(b []byte) uint32 {
	copy(wasmMemory()[scratchAddr:], b)
	return scratchAddr
}

func wasmMalloc(size uint32) uint32 {
	addr := uint32(module.Xgo_malloc(int32(size)))
	if addr == 0 {
		panic("clay: out of memory")
	}
	return addr
}

func b2i(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func putBool(b []byte, p uint32, v bool) { b[p] = byte(b2i(v)) }
func putU8(b []byte, p uint32, v uint8)  { b[p] = v }
func putU16(b []byte, p uint32, v uint16) {
	binary.LittleEndian.PutUint16(b[p:], v)
}

func putU32(b []byte, p uint32, v uint32) {
	binary.LittleEndian.PutUint32(b[p:], v)
}

func putU64(b []byte, p uint32, v uint64) {
	binary.LittleEndian.PutUint64(b[p:], v)
}

func putF32(b []byte, p uint32, v float32) {
	binary.LittleEndian.PutUint32(b[p:], math.Float32bits(v))
}

func putF64(b []byte, p uint32, v float64) {
	binary.LittleEndian.PutUint64(b[p:], math.Float64bits(v))
}

func getBool(m []byte, p uint32) bool  { return m[p] != 0 }
func getU8(m []byte, p uint32) uint8   { return m[p] }
func getU16(m []byte, p uint32) uint16 { return binary.LittleEndian.Uint16(m[p:]) }
func getU32(m []byte, p uint32) uint32 { return binary.LittleEndian.Uint32(m[p:]) }
func getU64(m []byte, p uint32) uint64 { return binary.LittleEndian.Uint64(m[p:]) }
func getF32(m []byte, p uint32) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(m[p:]))
}

func getF64(m []byte, p uint32) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(m[p:]))
}

// Clay keeps pointers to the strings and user data it is given until the next layout is done.
// Strings are copied into buffers inside the module and user data is replaced by a handle.
// Both are kept for two frames, so that data from the previous frame is still valid
// while handling pointer events before BeginLayout is called.

var (
	frame         uint32
	stringBuffers [2]stringBuffer
	frameHandles  [2][]any
	persistent    []any
)

const (
	persistentHandle = 1 << 31
	frameHandleBit   = 1 << 30
)

func nextFrame() {
	frame++
	stringBuffers[frame&1].reset()
	clear(frameHandles[frame&1])
	frameHandles[frame&1] = frameHandles[frame&1][:0]
}

// storeHandle returns a handle to v that is valid for this and the next frame.
func storeHandle(v any) uint32 {
	if v == nil {
		return 0
	}
	t := &frameHandles[frame&1]
	*t = append(*t, v)
	return (frame&1)*frameHandleBit | uint32(len(*t))
}

// storePersistentHandle returns a handle to v that is always valid.
func storePersistentHandle(v any) uint32 {
	persistent = append(persistent, v)
	return persistentHandle | uint32(len(persistent))
}

func loadHandle(h uint32) any {
	var t []any
	switch {
	case h == 0:
		return nil
	case h&persistentHandle != 0:
		t = persistent
	default:
		t = frameHandles[h/frameHandleBit&1]
	}
	i := h &^ (persistentHandle | frameHandleBit)
	if i == 0 || int(i) > len(t) {
		return nil
	}
	return t[i-1]
}

type stringBuffer struct {
	chunks []stringChunk
	// strings stored in the chunks, so that they can be returned without copying.
	strings  []storedString
	unsorted bool
}

type stringChunk struct {
	addr, size, used uint32
}

type storedString struct {
	addr uint32
	s    string
}

const minStringChunk = 64 << 10

func (b *stringBuffer) alloc(n uint32) uint32 {
	var last uint32
	if len(b.chunks) > 0 {
		c := &b.chunks[len(b.chunks)-1]
		if c.size-c.used >= n {
			addr := c.addr + c.used
			c.used += n
			return addr
		}
		last = c.size
	}
	size := max(n, minStringChunk, 2*last)
	addr := wasmMalloc(size)
	b.chunks = append(b.chunks, stringChunk{addr: addr, size: size, used: n})
	return addr
}

func (b *stringBuffer) reset() {
	if len(b.chunks) > 1 {
		// Replace the chunks with a single one that fits everything.
		var total uint32
		for _, c := range b.chunks {
			total += c.size
			module.Xgo_free(int32(c.addr))
		}
		b.chunks = append(b.chunks[:0], stringChunk{addr: wasmMalloc(total), size: total})
	} else if len(b.chunks) == 1 {
		b.chunks[0].used = 0
	}
	clear(b.strings)
	b.strings = b.strings[:0]
	b.unsorted = false
}

func (b *stringBuffer) lookup(addr, n uint32) (string, bool) {
	if len(b.strings) == 0 {
		return "", false
	}
	if b.unsorted {
		slices.SortFunc(b.strings, func(a, b storedString) int { return cmp.Compare(a.addr, b.addr) })
		b.unsorted = false
	}
	i, found := slices.BinarySearchFunc(b.strings, addr, func(s storedString, addr uint32) int { return cmp.Compare(s.addr, addr) })
	if !found {
		if i == 0 {
			return "", false
		}
		i--
	}
	s := b.strings[i]
	start := addr - s.addr
	if uint64(start)+uint64(n) > uint64(len(s.s)) {
		return "", false
	}
	return s.s[start : start+n], true
}

// storeString copies s into the module and returns its address.
func storeString(s string) uint32 {
	if len(s) == 0 {
		return 0
	}
	b := &stringBuffers[frame&1]
	addr := b.alloc(uint32(len(s)))
	copy(wasmMemory()[addr:], s)
	if n := len(b.strings); n > 0 && b.strings[n-1].addr > addr {
		b.unsorted = true
	}
	b.strings = append(b.strings, storedString{addr: addr, s: s})
	return addr
}

// loadString returns the string at addr, avoiding a copy if it was stored by storeString.
func loadString(m []byte, addr, n uint32) string {
	if n == 0 {
		return ""
	}
	for i := range stringBuffers {
		if s, ok := stringBuffers[(frame-uint32(i))&1].lookup(addr, n); ok {
			return s
		}
	}
	if uint64(addr)+uint64(n) > uint64(len(m)) {
		return ""
	}
	return string(m[addr : addr+n])
}

// Pointer points to a value inside clay's memory.
type Pointer[T any] struct {
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
	var buf [64]byte
	n := encodeValue(buf[:], 0, &v)
	copy(wasmMemory()[p.addr:], buf[:n])
}

// Context is clay's internal state.
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

type ErrorHandler struct {
	ErrorHandlerFunction func(errorData ErrorData)
	UserData             any
}

func Initialize(arena Arena, layoutDimensions Dimensions, errorHandler ErrorHandler) *Context {
	h := storePersistentHandle(&errorHandler)
	b := wasmArgs(sizeofDimensions)
	encDimensions(b, 0, &layoutDimensions)
	p := wasmCommit(b)
	return contextOf(uint32(module.Xgo_initialize(int32(arena.capacity), int32(arena.memory), int32(p), int32(h))))
}

type measureTextFunction struct {
	fn       func(text string, config *TextElementConfig, userData any) Dimensions
	userData any
}

// SetMeasureTextFunction sets the function used to measure text for the current context.
// The config passed to it must not be retained.
func SetMeasureTextFunction(measureTextFunction_ func(text string, config *TextElementConfig, userData any) Dimensions, userData any) {
	h := storePersistentHandle(&measureTextFunction{measureTextFunction_, userData})
	module.Xgo_set_measure_text_function(b2i(measureTextFunction_ != nil), int32(h))
}

type queryScrollOffsetFunction struct {
	fn       func(elementId uint32, userData any) Vector2
	userData any
}

func SetQueryScrollOffsetFunction(queryScrollOffsetFunction_ func(elementId uint32, userData any) Vector2, userData any) {
	h := storePersistentHandle(&queryScrollOffsetFunction{queryScrollOffsetFunction_, userData})
	module.Xgo_set_query_scroll_offset_function(b2i(queryScrollOffsetFunction_ != nil), int32(h))
}

type onHoverFunction struct {
	fn       func(elementId ElementId, pointerData PointerData, userData any)
	userData any
}

// OnHover binds a function that will be called when the pointer position provided by SetPointerState
// is within the current element's bounding box.
func OnHover(onHoverFunction_ func(elementId ElementId, pointerData PointerData, userData any), userData any) {
	module.Xgo_on_hover(int32(storeHandle(&onHoverFunction{onHoverFunction_, userData})))
}

func BeginLayout() {
	nextFrame()
	beginLayout()
}

func openTextElement(text string, config *TextElementConfig) {
	configOff := uint32(sizeofString+7) &^ 7
	b := wasmArgs(configOff + sizeofTextElementConfig)
	encString(b, 0, &text)
	encTextElementConfig(b, configOff, config)
	p := wasmCommit(b)
	module.Xgo_open_text_element(int32(p), int32(p+configOff))
}

// host implements the functions imported by the WebAssembly module.
type host struct{}

func (host) XerrorHandler(errorData int32) {
	var data ErrorData
	decErrorData(wasmMemory(), uint32(errorData), &data)
	h, _ := data.UserData.(*ErrorHandler)
	if h == nil || h.ErrorHandlerFunction == nil {
		return
	}
	data.UserData = h.UserData
	h.ErrorHandlerFunction(data)
}

func (host) XmeasureText(ret, text, config, userData int32) {
	f, _ := loadHandle(uint32(userData)).(*measureTextFunction)
	if f == nil || f.fn == nil {
		return
	}
	m := wasmMemory()
	var s string
	decStringSlice(m, uint32(text), &s)
	var c TextElementConfig
	decTextElementConfig(m, uint32(config), &c)
	d := f.fn(s, &c, f.userData)
	var buf [sizeofDimensions]byte
	encDimensions(buf[:], 0, &d)
	copy(wasmMemory()[ret:], buf[:])
}

func (host) XqueryScrollOffset(ret, elementId, userData int32) {
	f, _ := loadHandle(uint32(userData)).(*queryScrollOffsetFunction)
	if f == nil || f.fn == nil {
		return
	}
	v := f.fn(uint32(elementId), f.userData)
	var buf [sizeofVector2]byte
	encVector2(buf[:], 0, &v)
	copy(wasmMemory()[ret:], buf[:])
}

func (host) XonHover(elementId, pointerData, userData int32) {
	f, _ := loadHandle(uint32(userData)).(*onHoverFunction)
	if f == nil || f.fn == nil {
		return
	}
	m := wasmMemory()
	var id ElementId
	decElementId(m, uint32(elementId), &id)
	var pd PointerData
	decPointerData(m, uint32(pointerData), &pd)
	f.fn(id, pd, f.userData)
}
