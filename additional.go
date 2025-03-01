package clay

import "unsafe"

func (s String) String() string {
	return string(unsafe.Slice(s.Chars, s.Length))
}

func ToString(s string) String {
	return String{Length: int32(len(s)), Chars: unsafe.SliceData([]byte(s))}
}

func ID(label string) ElementId {
	return __HashString(ToString(label), 0, 0)
}
