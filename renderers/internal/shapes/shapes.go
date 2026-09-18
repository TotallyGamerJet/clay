// SPDX-License-Identifier: Zlib

// Package shapes builds the geometry of clay's rounded rectangles and borders,
// so that the renderers that draw them with triangles all draw them the same way.
//
// A shape is built as one mesh, rather than as edges and corners drawn separately,
// so that no seams appear between its parts. Renderers that can't anti alias the
// triangles they draw can ask for a feathered edge, which fades the outline of the
// shape out over a pixel.
package shapes

import (
	"math"

	"github.com/TotallyGamerJet/clay"
)

// Vertex is a point of a shape. Alpha is 1 inside the shape and 0 on a feathered edge,
// and is meant to scale the alpha of the color the shape is drawn with.
type Vertex struct {
	X, Y  float32
	Alpha float32
}

// Mesh is a shape as triangles.
type Mesh struct {
	Vertices []Vertex
	Indices  []uint16
}

func (m *Mesh) add(v Vertex) uint16 {
	m.Vertices = append(m.Vertices, v)
	return uint16(len(m.Vertices) - 1)
}

func (m *Mesh) triangle(a, b, c uint16) {
	m.Indices = append(m.Indices, a, b, c)
}

// quad adds the two triangles of the quad a, b, c, d. Every triangle of a mesh is wound
// the same way around, as renderers may cull the ones facing away from the viewer.
func (m *Mesh) quad(a, b, c, d uint16) {
	m.triangle(a, b, c)
	m.triangle(a, c, d)
}

// point of an outline, with the direction pointing out of the shape at it.
type point struct {
	x, y   float32
	nx, ny float32
}

// segments returns how many segments an arc of the given radius is drawn with.
// Every renderer uses the same number, so that they draw the same shape.
func segments(radius float32) int {
	return max(4, min(int(radius/2)+4, 32))
}

// outline returns the points of a rounded rectangle, clockwise from its left edge.
// Corners are sampled at the same angles regardless of their radius, so that two
// outlines of the same rectangle can be matched up point by point.
func outline(x0, y0, x1, y1 float32, radius clay.CornerRadius, arcSegments int) []point {
	maxRadius := max(min(x1-x0, y1-y0)/2, 0)
	radius.TopLeft = min(max(radius.TopLeft, 0), maxRadius)
	radius.TopRight = min(max(radius.TopRight, 0), maxRadius)
	radius.BottomLeft = min(max(radius.BottomLeft, 0), maxRadius)
	radius.BottomRight = min(max(radius.BottomRight, 0), maxRadius)

	corners := []struct {
		cx, cy, r  float32
		startAngle float64
	}{
		{x0 + radius.TopLeft, y0 + radius.TopLeft, radius.TopLeft, math.Pi},
		{x1 - radius.TopRight, y0 + radius.TopRight, radius.TopRight, 3 * math.Pi / 2},
		{x1 - radius.BottomRight, y1 - radius.BottomRight, radius.BottomRight, 0},
		{x0 + radius.BottomLeft, y1 - radius.BottomLeft, radius.BottomLeft, math.Pi / 2},
	}

	points := make([]point, 0, 4*(arcSegments+1))
	for _, c := range corners {
		for i := 0; i <= arcSegments; i++ {
			angle := c.startAngle + (math.Pi/2)*float64(i)/float64(arcSegments)
			nx, ny := float32(math.Cos(angle)), float32(math.Sin(angle))
			points = append(points, point{x: c.cx + nx*c.r, y: c.cy + ny*c.r, nx: nx, ny: ny})
		}
	}
	return points
}

// Fill returns the triangles of a rounded rectangle.
// Feather is how far the edge of the shape fades out, in pixels, or 0 for a hard edge.
func Fill(box clay.BoundingBox, radius clay.CornerRadius, feather float32) Mesh {
	var m Mesh
	if box.Width <= 0 || box.Height <= 0 {
		return m
	}
	points := outline(box.X, box.Y, box.X+box.Width, box.Y+box.Height, radius, segments(maxRadius(radius)))

	// A rounded rectangle is convex, so it can be drawn as a fan around its center.
	center := m.add(Vertex{X: box.X + box.Width/2, Y: box.Y + box.Height/2, Alpha: 1})
	inner := make([]uint16, len(points))
	for i, p := range points {
		inner[i] = m.add(Vertex{X: p.x, Y: p.y, Alpha: 1})
	}
	for i := range points {
		m.triangle(center, inner[i], inner[(i+1)%len(points)])
	}
	if feather > 0 {
		outer := make([]uint16, len(points))
		for i, p := range points {
			outer[i] = m.add(Vertex{X: p.x + p.nx*feather, Y: p.y + p.ny*feather, Alpha: 0})
		}
		for i := range points {
			j := (i + 1) % len(points)
			m.quad(inner[i], outer[i], outer[j], inner[j])
		}
	}
	return m
}

// Border returns the triangles of the border drawn inside a rounded rectangle.
// A side of zero width is left open, like clay's own renderers draw it.
func Border(box clay.BoundingBox, radius clay.CornerRadius, width clay.BorderWidth, feather float32) Mesh {
	var m Mesh
	if box.Width <= 0 || box.Height <= 0 {
		return m
	}
	x0, y0 := box.X, box.Y
	x1, y1 := box.X+box.Width, box.Y+box.Height
	left, right := float32(width.Left), float32(width.Right)
	top, bottom := float32(width.Top), float32(width.Bottom)

	arcSegments := segments(maxRadius(radius))
	outerPoints := outline(x0, y0, x1, y1, radius, arcSegments)
	// The inside of the border follows the outside, inset by the width of each side.
	shrink := func(r, a, b float32) float32 { return max(r-max(a, b), 0) }
	innerPoints := outline(x0+left, y0+top, x1-right, y1-bottom, clay.CornerRadius{
		TopLeft:     shrink(radius.TopLeft, left, top),
		TopRight:    shrink(radius.TopRight, right, top),
		BottomLeft:  shrink(radius.BottomLeft, left, bottom),
		BottomRight: shrink(radius.BottomRight, right, bottom),
	}, arcSegments)

	outer := make([]uint16, len(outerPoints))
	inner := make([]uint16, len(innerPoints))
	for i := range outerPoints {
		outer[i] = m.add(Vertex{X: outerPoints[i].x, Y: outerPoints[i].y, Alpha: 1})
		inner[i] = m.add(Vertex{X: innerPoints[i].x, Y: innerPoints[i].y, Alpha: 1})
	}
	var outerEdge, innerEdge []uint16
	if feather > 0 {
		outerEdge = make([]uint16, len(outerPoints))
		innerEdge = make([]uint16, len(innerPoints))
		for i := range outerPoints {
			o, in := outerPoints[i], innerPoints[i]
			outerEdge[i] = m.add(Vertex{X: o.x + o.nx*feather, Y: o.y + o.ny*feather, Alpha: 0})
			innerEdge[i] = m.add(Vertex{X: in.x - in.nx*feather, Y: in.y - in.ny*feather, Alpha: 0})
		}
	}

	for i := range outerPoints {
		j := (i + 1) % len(outerPoints)
		if !drawn(outerPoints[i], outerPoints[j], width) {
			continue
		}
		m.quad(outer[i], outer[j], inner[j], inner[i])
		if feather > 0 {
			m.quad(outer[i], outerEdge[i], outerEdge[j], outer[j])
			// A corner of the inside of the border can collapse to a point, where the
			// border is thicker than the radius. There is no edge there to fade out.
			if in, next := innerPoints[i], innerPoints[j]; in.x != next.x || in.y != next.y {
				m.quad(inner[i], inner[j], innerEdge[j], innerEdge[i])
			}
		}
	}
	return m
}

// drawn reports whether the piece of border between two points is on a side that has width.
// The direction out of the shape says which sides a piece belongs to; pieces of an arc
// belong to the two sides that meet at that corner.
func drawn(a, b point, width clay.BorderWidth) bool {
	const epsilon = 0.0001
	for _, p := range [2]point{a, b} {
		if p.nx < -epsilon && width.Left <= 0 {
			return false
		}
		if p.nx > epsilon && width.Right <= 0 {
			return false
		}
		if p.ny < -epsilon && width.Top <= 0 {
			return false
		}
		if p.ny > epsilon && width.Bottom <= 0 {
			return false
		}
	}
	return true
}

func maxRadius(r clay.CornerRadius) float32 {
	return max(r.TopLeft, r.TopRight, r.BottomLeft, r.BottomRight)
}
