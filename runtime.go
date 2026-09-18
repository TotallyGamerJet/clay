// SPDX-License-Identifier: Zlib

package clay

import (
	"cmp"
	"encoding/binary"
	"math"
	"slices"
	"unsafe"

	"github.com/TotallyGamerJet/clay/internal/wasm"
)

// This file is the bridge between Go and the WebAssembly module that clay runs as.
// Nothing in it is part of the package's API.

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
	if uint32(len(b)) > scratchSize {
		// The scratch buffer is sized for the largest generated call, so this can only
		// happen if a hand written call passes something bigger.
		panic("clay: arguments are larger than the scratch buffer")
	}
	copy(wasmMemory()[scratchAddr:scratchAddr+scratchSize], b)
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
	// Hover callbacks are kept as values in a table of their own, which clay refers to
	// by index, so that binding one to an element every frame doesn't allocate.
	frameHovers [2][]onHoverFunction

	// Handles that outlive a frame, like the callbacks of a context. Each one has a slot,
	// which is reused when the same callback of the same context is set again.
	persistent      []any
	persistentFree  []uint32
	persistentSlots = map[persistentKey]uint32{}
)

// persistentKey identifies what a persistent handle is for, so that setting it again
// reuses its slot instead of leaking the old one.
type persistentKey struct {
	kind persistentKind
	ctx  uint32 // the context the handle belongs to
}

type persistentKind uint8

const (
	errorHandlerHandle persistentKind = iota
	measureTextHandle
	queryScrollOffsetHandle
)

const (
	persistentHandle = 1 << 31
	frameHandleBit   = 1 << 30
)

func nextFrame() {
	frame++
	if internFull {
		flushInterned()
	}
	stringBuffers[frame&1].reset()
	clear(frameHandles[frame&1])
	frameHandles[frame&1] = frameHandles[frame&1][:0]
	clear(frameHovers[frame&1])
	frameHovers[frame&1] = frameHovers[frame&1][:0]
}

// storeHover returns a handle to a hover callback, valid for this and the next frame.
func storeHover(v onHoverFunction) uint32 {
	t := &frameHovers[frame&1]
	*t = append(*t, v)
	return (frame&1)*frameHandleBit | uint32(len(*t))
}

func loadHover(h uint32) onHoverFunction {
	t := frameHovers[h/frameHandleBit&1]
	i := h &^ frameHandleBit
	if i == 0 || int(i) > len(t) {
		return onHoverFunction{}
	}
	return t[i-1]
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

// allocPersistent stores v in a free slot and returns the slot, which is 1 based.
func allocPersistent(v any) uint32 {
	if n := len(persistentFree); n > 0 {
		slot := persistentFree[n-1]
		persistentFree = persistentFree[:n-1]
		persistent[slot-1] = v
		return slot
	}
	persistent = append(persistent, v)
	return uint32(len(persistent))
}

// freePersistent makes the slot available to the next handle.
func freePersistent(slot uint32) {
	persistent[slot-1] = nil
	persistentFree = append(persistentFree, slot)
}

// storePersistentHandle returns a handle to v that is always valid. Storing a handle for the
// same purpose and context again reuses its slot, so that setting a callback every frame
// doesn't grow the table.
func storePersistentHandle(key persistentKey, v any) uint32 {
	slot, ok := persistentSlots[key]
	if ok {
		persistent[slot-1] = v
	} else {
		slot = allocPersistent(v)
		persistentSlots[key] = slot
	}
	return persistentHandle | slot
}

// takePersistentSlot gives the slot of an already stored handle the given key,
// freeing whatever the key held before. Initialize uses it, as the context a handle
// belongs to is only known once it has been created.
func takePersistentSlot(key persistentKey, handle uint32) {
	if old, ok := persistentSlots[key]; ok {
		freePersistent(old)
	}
	persistentSlots[key] = handle &^ persistentHandle
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

// store copies s into the buffer and returns its address.
func (b *stringBuffer) store(s string) uint32 {
	addr := b.alloc(uint32(len(s)))
	copy(wasmMemory()[addr:], s)
	if n := len(b.strings); n > 0 && b.strings[n-1].addr > addr {
		b.unsorted = true
	}
	b.strings = append(b.strings, storedString{addr: addr, s: s})
	return addr
}

// free releases the memory of the buffer and the strings it keeps.
func (b *stringBuffer) free() {
	for _, c := range b.chunks {
		module.Xgo_free(int32(c.addr))
	}
	*b = stringBuffer{}
}

// storeString copies s into this frame's buffer and returns its address, which is valid
// for this frame and the next.
func storeString(s string) uint32 {
	if len(s) == 0 {
		return 0
	}
	return stringBuffers[frame&1].store(s)
}

// Clay looks text up in a cache of its measurements by a hash of the text, which it
// computes every frame. For most strings it hashes their contents, which for long text
// like the demo's article took a third of the time a layout took. But if a string says
// its memory is statically allocated, clay hashes its address and length instead.
//
// So strings are interned: each distinct Go string is copied into the module once, at an
// address that it keeps, and passed to clay as static. The table is keyed by the string's
// data pointer and length rather than its contents, so looking a string up costs nothing
// like hashing it. That is safe because the table keeps the string alive, so Go can't
// reuse its memory for a different string while it is in the table.
//
// An address must never hold different text while any context might have it cached.
// Clay returns a cache entry whose hash matches however old the entry is, so waiting a
// few frames before reusing an address would not be enough. So interned strings are
// never freed one at a time. Once the table is over its budget, new strings are copied
// per frame instead, and at the next BeginLayout the whole table is dropped along with the
// measurement caches of every context, so that nothing measured at a freed address can
// be returned. The next frame measures its text again.
//
// Only long strings are interned. Hashing a short one costs about as much as looking it
// up in the table would, and short strings are the ones a program is likely to format anew
// every frame, like a score or a timer, which would fill the table for nothing. Element
// ids gain nothing either way, as clay hashes their contents regardless.
//
// A program whose long strings are all different every frame
// fills the table and drops it now and then. One whose strings don't fit in the budget
// would drop it every frame, so the budget doubles when that happens, up to a limit.

const (
	// Strings shorter than this are copied per frame rather than interned.
	minInternLength     = 64
	initialInternBudget = 1 << 20
	maxInternBudget     = 64 << 20
	// Dropping the table again within this many frames means the strings don't fit.
	internFlushWindow = 60
)

var (
	internMinLength        = minInternLength // variables, so tests can change them
	internBudget    uint32 = initialInternBudget
	internBuffer    stringBuffer
	interned        = map[internKey]uint32{}
	internedBytes   uint32
	internFull      bool // a string didn't fit, so the table is dropped at the next frame
	lastInternDrop  uint32
)

type internKey struct {
	data *byte
	size int
}

// internString returns the address of s in the module and whether it is stable, which
// is when it can be passed to clay as statically allocated.
func internString(s string) (addr uint32, stable bool) {
	if len(s) < internMinLength {
		return storeString(s), false
	}
	key := internKey{data: unsafe.StringData(s), size: len(s)}
	if addr, ok := interned[key]; ok {
		return addr, true
	}
	if internFull || uint64(internedBytes)+uint64(len(s)) > uint64(internBudget) {
		internFull = true
		return storeString(s), false
	}
	addr = internBuffer.store(s)
	interned[key] = addr
	internedBytes += uint32(len(s))
	return addr, true
}

// flushInterned drops every interned string, and every measurement clay might have
// cached under the address of one.
func flushInterned() {
	if frame-lastInternDrop < internFlushWindow {
		internBudget = min(internBudget*2, maxInternBudget)
	}
	lastInternDrop = frame
	internBuffer.free()
	clear(interned)
	internedBytes = 0
	internFull = false

	current := GetCurrentContext()
	for _, c := range contexts {
		SetCurrentContext(c)
		ResetMeasureTextCache()
	}
	SetCurrentContext(current)
}

// loadString returns the string at addr, avoiding a copy if it was stored by storeString.
func loadString(m []byte, addr, n uint32) string {
	if n == 0 {
		return ""
	}
	if s, ok := internBuffer.lookup(addr, n); ok {
		return s
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

func openTextElement(text string, config *TextElementConfig) {
	configOff := uint32(sizeofString+7) &^ 7
	b := wasmArgs(configOff + sizeofTextElementConfig)
	encString(b, 0, &text)
	encTextElementConfig(b, configOff, config)
	p := wasmCommit(b)
	module.Xgo_open_text_element(int32(p), int32(p+configOff))
}

// The functions registered for the transition trampolines, indexed by their slot.
var (
	transitionHandlers   []func(arguments TransitionCallbackArguments) bool
	transitionStateFuncs []func(state TransitionData, properties TransitionProperty) TransitionData
)

// measuredConfig is handed to the measure text function, and reused between calls so
// that it doesn't escape to the heap on each one. Measuring is not reentrant.
var measuredConfig TextElementConfig

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
	decTextElementConfig(m, uint32(config), &measuredConfig)
	d := f.fn(s, &measuredConfig, f.userData)
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

func (host) XtransitionHandler(slot, arguments int32) int32 {
	if int(slot) >= len(transitionHandlers) {
		return 0
	}
	var args TransitionCallbackArguments
	decTransitionCallbackArguments(wasmMemory(), uint32(arguments), &args)
	return b2i(transitionHandlers[slot](args))
}

func (host) XtransitionState(slot, ret, state int32, properties int32) {
	if int(slot) >= len(transitionStateFuncs) {
		return
	}
	var s TransitionData
	decTransitionData(wasmMemory(), uint32(state), &s)
	result := transitionStateFuncs[slot](s, TransitionProperty(properties))
	var buf [sizeofTransitionData]byte
	encTransitionData(buf[:], 0, &result)
	copy(wasmMemory()[ret:], buf[:])
}

func (host) XonHover(elementId, pointerData, userData int32) {
	f := loadHover(uint32(userData))
	if f.fn == nil {
		return
	}
	m := wasmMemory()
	var id ElementId
	decElementId(m, uint32(elementId), &id)
	var pd PointerData
	decPointerData(m, uint32(pointerData), &pd)
	f.fn(id, pd, f.userData)
}
