// SPDX-License-Identifier: Zlib

// Package overlay tracks the overlay colors set by clay's OVERLAY_COLOR render commands.
package overlay

import "github.com/TotallyGamerJet/clay"

// Stack holds the overlay colors that are active while rendering.
// Everything drawn between RenderCommandTypeOverlayColorStart and
// RenderCommandTypeOverlayColorEnd is blended towards the overlay color by its alpha.
type Stack []clay.Color

// Push handles RenderCommandTypeOverlayColorStart.
func (s *Stack) Push(c clay.Color) {
	*s = append(*s, c)
}

// Pop handles RenderCommandTypeOverlayColorEnd.
func (s *Stack) Pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

// Active reports whether any overlay is active.
func (s Stack) Active() bool {
	return len(s) > 0
}

// Combined returns a single overlay color that is equivalent to applying all active overlays,
// for renderers that can only apply one overlay at a time.
func (s Stack) Combined() clay.Color {
	// Blending towards o1 by a1 then o2 by a2 leaves (1-a1)(1-a2) of the original color.
	var c clay.Color
	keep := float32(1)
	for _, o := range s {
		a := o.A / 255
		c.R = c.R*(1-a) + o.R*a
		c.G = c.G*(1-a) + o.G*a
		c.B = c.B*(1-a) + o.B*a
		keep *= 1 - a
	}
	if keep >= 1 {
		return clay.Color{}
	}
	total := 1 - keep
	return clay.Color{R: c.R / total, G: c.G / total, B: c.B / total, A: total * 255}
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
