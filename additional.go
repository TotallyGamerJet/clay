package clay

import "unsafe"

func (s String) String() string {
	return unsafe.String(s.Chars, s.Length)
}

func (s StringSlice) String() string {
	return unsafe.String(s.Chars, s.Length)
}

func ToString(s string) String {
	return String{Length: int32(len(s)), Chars: unsafe.StringData(s)}
}

func ID(label string) ElementId {
	return __HashString(ToString(label), 0, 0)
}

func PaddingAll(padding uint16) Padding {
	return Padding{
		padding,
		padding,
		padding,
		padding,
	}
}

func SizingGrow(sz float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			MinMax: SizingMinMax{sz, sz},
		},
		Type: __SIZING_TYPE_GROW}
}

func SizingFixed(sz float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			MinMax: SizingMinMax{sz, sz},
		},
		Type: __SIZING_TYPE_FIXED}
}

func CornerRadiusAll(radius float32) CornerRadius {
	return CornerRadius{
		radius,
		radius,
		radius,
		radius,
	}
}

// TODO: add generic iterator functions for types with [type]_GetValue functions that are converted into methods

func UI() func(decl ElementDeclaration, children func()) {
	__OpenElement()
	return func(decl ElementDeclaration, children func()) {
		__ConfigureOpenElement(decl)
		defer __CloseElement()
		if children != nil {
			children()
		}
	}
}

func Text(text String, config *TextElementConfig) {
	__OpenTextElement(text, config)
}

func TextConfig(config TextElementConfig) *TextElementConfig {
	return __StoreTextElementConfig(config)
}
