// Package overlay tracks the overlay colors set by clay's OVERLAY_COLOR render commands.
package overlay

import "github.com/TotallyGamerJet/clay"

// Stack holds the overlay colors that are active while rendering.
// Everything drawn between RENDER_COMMAND_TYPE_OVERLAY_COLOR_START and
// RENDER_COMMAND_TYPE_OVERLAY_COLOR_END is blended towards the overlay color by its alpha.
type Stack []clay.Color

// Push handles RENDER_COMMAND_TYPE_OVERLAY_COLOR_START.
func (s *Stack) Push(c clay.Color) {
	*s = append(*s, c)
}

// Pop handles RENDER_COMMAND_TYPE_OVERLAY_COLOR_END.
func (s *Stack) Pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

// Active reports whether any overlay is active.
func (s Stack) Active() bool {
	return len(s) > 0
}

// Apply blends the color of c towards the active overlays. The alpha of c is kept.
func (s Stack) Apply(c clay.Color) clay.Color {
	for _, o := range s {
		a := o.A / 255
		c.R += (o.R - c.R) * a
		c.G += (o.G - c.G) * a
		c.B += (o.B - c.B) * a
	}
	return c
}
