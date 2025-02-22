package clay

import "unsafe"

func (s String) String() string {
	return string(unsafe.Slice(s.Chars, s.Length))
}
