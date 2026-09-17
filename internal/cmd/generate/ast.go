package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// node is the subset of clang's JSON AST that the generator needs.
type node struct {
	ID                  string    `json:"id"`
	Kind                string    `json:"kind"`
	Name                string    `json:"name"`
	TagUsed             string    `json:"tagUsed"`
	CompleteDefinition  bool      `json:"completeDefinition"`
	FixedUnderlyingType *qualType `json:"fixedUnderlyingType"`
	Type                *qualType `json:"type"`
	Decl                *node     `json:"decl"`
	OwnedTagDecl        *node     `json:"ownedTagDecl"`
	Value               string    `json:"value"`
	Loc                 *loc      `json:"loc"`
	Inner               []*node   `json:"inner"`
}

// loc is a location in the source. Offsets are relative to the file the declaration is in,
// which is clay.h for everything the generator looks at.
type loc struct {
	Offset int `json:"offset"`
	Line   int `json:"line"`
	// Declarations written with a macro, like CLAY_PACKED_ENUM, are located in the macro.
	ExpansionLoc *loc `json:"expansionLoc"`
}

type qualType struct {
	QualType          string `json:"qualType"`
	DesugaredQualType string `json:"desugaredQualType"`
}

type kind int

const (
	kPrim kind = iota
	kVoid
	kEnum
	kRecord
	kPtr
	kFuncPtr
	kArray
	kOpaque // incomplete struct, like Clay_Context
)

type ctype struct {
	kind  kind
	cName string // C spelling, for records and enums
	prim  *prim
	enum  *enum
	rec   *record
	elem  *ctype
	len   uint32
}

type prim struct {
	goType string
	size   uint32
	wasm   string // i32, i64, f32 or f64
}

type enum struct {
	cName, goName string
	consts        []enumConst
	prim          *prim
	doc           []string
}

type enumConst struct {
	cName, goName string
	value         int64
	doc           []string
}

type specialRecord int

const (
	recNormal specialRecord = iota
	recString
	recArray
)

type record struct {
	cName, goName string // empty for anonymous records
	union         bool
	fields        []*field
	size, align   uint32
	special       specialRecord
	elem          *ctype // element type for recArray

	// For anonymous records: the enclosing named record and the member path used in offsetof.
	owner *record
	path  string

	doc []string
}

type field struct {
	cName, goName string
	t             *ctype
	off           uint32
	doc           []string
}

var prims = map[string]*prim{
	"bool":               {"bool", 1, "i32"},
	"_Bool":              {"bool", 1, "i32"},
	"char":               {"int8", 1, "i32"},
	"signed char":        {"int8", 1, "i32"},
	"unsigned char":      {"uint8", 1, "i32"},
	"short":              {"int16", 2, "i32"},
	"unsigned short":     {"uint16", 2, "i32"},
	"int":                {"int32", 4, "i32"},
	"unsigned int":       {"uint32", 4, "i32"},
	"long":               {"int32", 4, "i32"},
	"unsigned long":      {"uint32", 4, "i32"},
	"long long":          {"int64", 8, "i64"},
	"unsigned long long": {"uint64", 8, "i64"},
	"float":              {"float32", 4, "f32"},
	"double":             {"float64", 8, "f64"},
	"int8_t":             {"int8", 1, "i32"},
	"uint8_t":            {"uint8", 1, "i32"},
	"int16_t":            {"int16", 2, "i32"},
	"uint16_t":           {"uint16", 2, "i32"},
	"int32_t":            {"int32", 4, "i32"},
	"uint32_t":           {"uint32", 4, "i32"},
	"int64_t":            {"int64", 8, "i64"},
	"uint64_t":           {"uint64", 8, "i64"},
	"size_t":             {"uint32", 4, "i32"},
	"uintptr_t":          {"uint32", 4, "i32"},
	"intptr_t":           {"int32", 4, "i32"},
}

type model struct {
	src        []byte // clay.h, for the comments documenting each declaration
	lineStarts []int

	decls     map[string]*node // RecordDecl and EnumDecl by id
	tags      map[string]*node // complete RecordDecl by "struct Name"
	typedefs  map[string]*node
	funcs     []*node
	recs      map[*node]*record
	enums     map[*node]*enum
	aliases   map[string]*ctype // typedef name -> type, for typedefs naming an already named record
	recOrder  []*record
	enumOrder []*enum
}

// doc returns the comment documenting the declaration at the given source location,
// which is either the comment lines above it or the comment trailing it on its own line.
func (m *model) doc(l *loc) []string {
	if l != nil && l.Offset == 0 && l.ExpansionLoc != nil {
		l = l.ExpansionLoc
	}
	if l == nil || l.Offset <= 0 || l.Offset >= len(m.src) {
		return nil
	}
	line, _ := slices.BinarySearch(m.lineStarts, l.Offset)
	line-- // the line containing the offset
	text := func(i int) string {
		if i < 0 || i >= len(m.lineStarts) {
			return ""
		}
		end := len(m.src)
		if i+1 < len(m.lineStarts) {
			end = m.lineStarts[i+1]
		}
		return strings.TrimRight(string(m.src[m.lineStarts[i]:end]), "\r\n")
	}

	var doc []string
	for i := line - 1; i >= 0; i-- {
		comment, ok := strings.CutPrefix(strings.TrimSpace(text(i)), "//")
		if !ok {
			break
		}
		doc = append(doc, strings.TrimSpace(comment))
	}
	slices.Reverse(doc)

	// A comment trailing the declaration, like "uint32_t id; // The hash."
	if rest := text(line)[min(l.Offset-m.lineStarts[line], len(text(line))):]; strings.Contains(rest, "//") {
		_, comment, _ := strings.Cut(rest, "//")
		doc = append(doc, strings.TrimSpace(comment))
	}
	return doc
}

func parseAST(src, data []byte) (*model, error) {
	var root node
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}
	m := &model{
		src:      src,
		decls:    map[string]*node{},
		tags:     map[string]*node{},
		typedefs: map[string]*node{},
		recs:     map[*node]*record{},
		enums:    map[*node]*enum{},
		aliases:  map[string]*ctype{},
	}
	m.lineStarts = []int{0}
	for i, b := range src {
		if b == '\n' {
			m.lineStarts = append(m.lineStarts, i+1)
		}
	}
	var index func(n *node)
	index = func(n *node) {
		switch n.Kind {
		case "RecordDecl", "EnumDecl":
			m.decls[n.ID] = n
			if n.Kind == "RecordDecl" && n.Name != "" && n.CompleteDefinition {
				m.tags[n.TagUsed+" "+n.Name] = n
			}
		}
		for _, c := range n.Inner {
			index(c)
		}
	}
	seen := map[string]bool{}
	for _, n := range root.Inner {
		index(n)
		switch n.Kind {
		case "TypedefDecl":
			m.typedefs[n.Name] = n
		case "FunctionDecl":
			if !seen[n.Name] && strings.HasPrefix(n.Name, "Clay_") {
				seen[n.Name] = true
				m.funcs = append(m.funcs, n)
			}
		}
	}
	return m, nil
}

func stripQualifiers(s string) string {
	s = strings.TrimSpace(s)
	for _, q := range []string{"const ", "volatile ", "restrict "} {
		s = strings.TrimPrefix(s, q)
	}
	for _, q := range []string{" const", " restrict"} {
		s = strings.TrimSuffix(s, q)
	}
	return strings.TrimSpace(s)
}

// resolve resolves a type spelled by clang. anon is the anonymous record declared
// just before a field, used for fields like "union (unnamed at clay.h:1:1)".
func (m *model) resolve(spelling string, anon *node, owner *record, path string) (*ctype, error) {
	s := stripQualifiers(spelling)
	switch {
	case strings.Contains(s, "(*)"):
		return &ctype{kind: kFuncPtr}, nil
	case strings.HasSuffix(s, "*"):
		elem, err := m.resolve(strings.TrimSuffix(s, "*"), nil, nil, "")
		if err != nil {
			return nil, err
		}
		return &ctype{kind: kPtr, elem: elem}, nil
	case strings.HasSuffix(s, "]"):
		i := strings.LastIndexByte(s, '[')
		n, err := strconv.ParseUint(s[i+1:len(s)-1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("unsupported array type %q", s)
		}
		elem, err := m.resolve(s[:i], nil, nil, "")
		if err != nil {
			return nil, err
		}
		return &ctype{kind: kArray, elem: elem, len: uint32(n)}, nil
	case strings.Contains(s, "(unnamed") || strings.Contains(s, "(anonymous"):
		if anon == nil {
			return nil, fmt.Errorf("anonymous type %q without declaration", s)
		}
		rec, err := m.record(anon, "", owner, path)
		if err != nil {
			return nil, err
		}
		return &ctype{kind: kRecord, rec: rec}, nil
	case s == "void":
		return &ctype{kind: kVoid}, nil
	}
	if p, ok := prims[s]; ok {
		return &ctype{kind: kPrim, prim: p}, nil
	}
	if t, ok := m.aliases[s]; ok {
		return t, nil
	}
	if tag, name, ok := strings.Cut(s, " "); ok && (tag == "struct" || tag == "union" || tag == "enum") {
		if n, ok := m.tags[s]; ok {
			rec, err := m.record(n, name, nil, "")
			if err != nil {
				return nil, err
			}
			return &ctype{kind: kRecord, cName: s, rec: rec}, nil
		}
		if tag != "enum" {
			return &ctype{kind: kOpaque, cName: s}, nil
		}
		return nil, fmt.Errorf("unknown type %q", s)
	}
	td, ok := m.typedefs[s]
	if !ok {
		return nil, fmt.Errorf("unknown type %q", s)
	}
	var t *ctype
	if decl := typedefDecl(td); decl != nil {
		d := m.decls[decl.ID]
		switch d.Kind {
		case "EnumDecl":
			e, err := m.enum(d, s)
			if err != nil {
				return nil, err
			}
			t = &ctype{kind: kEnum, cName: s, enum: e}
		case "RecordDecl":
			if !d.CompleteDefinition {
				if c, ok := m.tags[d.TagUsed+" "+d.Name]; ok {
					d = c
				} else {
					t = &ctype{kind: kOpaque, cName: s}
					break
				}
			}
			rec, err := m.record(d, s, nil, "")
			if err != nil {
				return nil, err
			}
			t = &ctype{kind: kRecord, cName: s, rec: rec}
		}
	} else {
		under := td.Type.QualType
		if td.Type.DesugaredQualType != "" {
			under = td.Type.DesugaredQualType
		}
		var err error
		if t, err = m.resolve(under, nil, nil, ""); err != nil {
			return nil, err
		}
	}
	m.aliases[s] = t
	return t, nil
}

func typedefDecl(n *node) *node {
	for _, c := range n.Inner {
		if c.Decl != nil && (c.Decl.Kind == "RecordDecl" || c.Decl.Kind == "EnumDecl") {
			return c.Decl
		}
		if c.OwnedTagDecl != nil {
			return c.OwnedTagDecl
		}
		if d := typedefDecl(c); d != nil {
			return d
		}
	}
	return nil
}

func (m *model) enum(n *node, cName string) (*enum, error) {
	if e, ok := m.enums[n]; ok {
		return e, nil
	}
	e := &enum{cName: cName, goName: goName(cName), doc: m.doc(n.Loc)}
	m.enums[n] = e
	packed := false
	next := int64(0)
	minV, maxV := int64(0), int64(0)
	for _, c := range n.Inner {
		switch c.Kind {
		case "PackedAttr":
			packed = true
		case "EnumConstantDecl":
			if v, ok := constValue(c); ok {
				next = v
			} else if len(c.Inner) > 0 {
				return nil, fmt.Errorf("enum constant %s: can't evaluate value", c.Name)
			}
			e.consts = append(e.consts, enumConst{cName: c.Name, goName: goConstName(c.Name), value: next, doc: m.doc(c.Loc)})
			minV, maxV = min(minV, next), max(maxV, next)
			next++
		}
	}
	switch {
	case n.FixedUnderlyingType != nil:
		t, err := m.resolve(n.FixedUnderlyingType.QualType, nil, nil, "")
		if err != nil || t.kind != kPrim {
			return nil, fmt.Errorf("enum %s: unsupported underlying type", cName)
		}
		e.prim = t.prim
	case !packed:
		if minV < 0 || maxV <= 0x7fffffff {
			e.prim = prims["int32_t"]
		} else {
			e.prim = prims["uint32_t"]
		}
	case minV >= 0 && maxV <= 0xff:
		e.prim = prims["uint8_t"]
	case minV >= -0x80 && maxV <= 0x7f:
		e.prim = prims["int8_t"]
	case minV >= 0 && maxV <= 0xffff:
		e.prim = prims["uint16_t"]
	case minV >= -0x8000 && maxV <= 0x7fff:
		e.prim = prims["int16_t"]
	default:
		e.prim = prims["int32_t"]
	}
	m.enumOrder = append(m.enumOrder, e)
	return e, nil
}

func constValue(n *node) (int64, bool) {
	if n.Kind == "ConstantExpr" && n.Value != "" {
		v, err := strconv.ParseInt(n.Value, 10, 64)
		return v, err == nil
	}
	for _, c := range n.Inner {
		if v, ok := constValue(c); ok {
			return v, true
		}
	}
	return 0, false
}

func (m *model) record(n *node, cName string, owner *record, path string) (*record, error) {
	if r, ok := m.recs[n]; ok {
		return r, nil
	}
	r := &record{cName: cName, goName: goName(cName), union: n.TagUsed == "union", owner: owner, path: path, doc: m.doc(n.Loc)}
	if cName == "" {
		r.goName = ""
	}
	m.recs[n] = r
	var anon *node
	for _, c := range n.Inner {
		switch c.Kind {
		case "RecordDecl":
			anon = c
		case "FieldDecl":
			if c.Name == "" {
				return nil, fmt.Errorf("record %s: unnamed fields are not supported", cName)
			}
			fowner, fpath := r, c.Name
			if r.cName == "" {
				fowner, fpath = r.owner, r.path+"."+c.Name
			}
			t, err := m.resolve(c.Type.QualType, anon, fowner, fpath)
			if err != nil {
				return nil, fmt.Errorf("record %s field %s: %w", cName, c.Name, err)
			}
			r.fields = append(r.fields, &field{cName: c.Name, goName: exportName(c.Name), t: t, doc: m.doc(c.Loc)})
			anon = nil
		}
	}
	r.layout()
	if cName != "" {
		m.recOrder = append(m.recOrder, r)
	}
	return r, nil
}

func (t *ctype) sizeAlign() (uint32, uint32) {
	switch t.kind {
	case kPrim:
		return t.prim.size, t.prim.size
	case kEnum:
		return t.enum.prim.size, t.enum.prim.size
	case kRecord:
		return t.rec.size, t.rec.align
	case kPtr, kFuncPtr:
		return 4, 4
	case kArray:
		s, a := t.elem.sizeAlign()
		return s * t.len, a
	}
	panic(fmt.Sprintf("type of kind %d has no size", t.kind))
}

func alignUp(v, a uint32) uint32 {
	return (v + a - 1) &^ (a - 1)
}

func (r *record) layout() {
	r.align = 1
	for _, f := range r.fields {
		s, a := f.t.sizeAlign()
		r.align = max(r.align, a)
		if r.union {
			r.size = max(r.size, s)
		} else {
			f.off = alignUp(r.size, a)
			r.size = f.off + s
		}
	}
	r.size = alignUp(r.size, r.align)
}

func goName(c string) string {
	if rest, ok := strings.CutPrefix(c, "Clay__"); ok {
		return "__" + rest
	}
	return strings.TrimPrefix(c, "Clay_")
}

func goConstName(c string) string {
	if rest, ok := strings.CutPrefix(c, "CLAY__"); ok {
		return "__" + rest
	}
	return strings.TrimPrefix(c, "CLAY_")
}

func exportName(c string) string {
	return strings.ToUpper(c[:1]) + c[1:]
}
