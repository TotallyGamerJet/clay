package clay

import (
	"fmt"
)

// confirm that ErrorData implements error type
var _ error = ErrorData{}

func (e ErrorData) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.ErrorText, e.ErrorType)
}

func ID(label string) ElementId {
	return __HashString(label, 0)
}

func IDI(label string, index uint32) ElementId {
	return __HashStringWithOffset(label, index, 0)
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
		Type: __SIZING_TYPE_GROW,
	}
}

func SizingFixed(sz float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			MinMax: SizingMinMax{sz, sz},
		},
		Type: __SIZING_TYPE_FIXED,
	}
}

func SizingFit(min, max float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			MinMax: SizingMinMax{min, max},
		},
		Type: __SIZING_TYPE_FIT,
	}
}

func SizingPercent(percentOfParent float32) SizingAxis {
	return SizingAxis{
		Size: struct {
			MinMax  SizingMinMax
			Percent float32
		}{
			Percent: percentOfParent,
		},
		Type: __SIZING_TYPE_PERCENT,
	}
}

func CornerRadiusAll(radius float32) CornerRadius {
	return CornerRadius{
		radius,
		radius,
		radius,
		radius,
	}
}

func BorderOutside(width uint16) BorderWidth {
	return BorderWidth{
		width,
		width,
		width,
		width,
		0,
	}
}

func BorderAll(width uint16) BorderWidth {
	return BorderWidth{
		width,
		width,
		width,
		width,
		width,
	}
}

// TODO: add generic iterator functions for types with [type]_GetValue functions that are converted into methods

func UI(id ...ElementId) func(decl ElementDeclaration, children func()) {
	if len(id) > 1 {
		panic("clay: too many element ids")
	} else if len(id) == 1 {
		__OpenElementWithId(id[0])
	} else {
		__OpenElement()
	}
	return func(decl ElementDeclaration, children func()) {
		__ConfigureOpenElement(decl)
		defer __CloseElement()
		if children != nil {
			children()
		}
	}
}

func Text(text string, config *TextElementConfig) {
	if config == nil {
		config = &TextElementConfig{}
	}
	openTextElement(text, config)
}
