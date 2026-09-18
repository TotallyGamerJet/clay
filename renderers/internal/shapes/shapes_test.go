package shapes

import (
	"testing"

	"github.com/TotallyGamerJet/clay"
)

// area of a triangle, whose sign says which way around it is wound.
func area(a, b, c Vertex) float32 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

// TestWinding checks that the triangles of a shape are all wound the same way around,
// as raylib and other renderers cull the ones facing away from the viewer.
func TestWinding(t *testing.T) {
	box := clay.BoundingBox{X: 10, Y: 20, Width: 100, Height: 60}
	radius := clay.CornerRadius{TopLeft: 20, TopRight: 0, BottomRight: 15, BottomLeft: 5}
	width := clay.BorderWidth{Left: 4, Right: 6, Top: 3, Bottom: 8}

	for _, test := range []struct {
		name string
		mesh Mesh
	}{
		{"fill", Fill(box, radius, 0)},
		{"feathered fill", Fill(box, radius, 0.5)},
		{"border", Border(box, radius, width, 0)},
		{"feathered border", Border(box, radius, width, 0.5)},
		{"square border", Border(box, clay.CornerRadius{}, width, 0)},
	} {
		mesh := test.mesh
		if len(mesh.Indices) == 0 {
			t.Errorf("%s: no triangles", test.name)
			continue
		}
		if len(mesh.Indices)%3 != 0 {
			t.Errorf("%s: %d indices, want a multiple of 3", test.name, len(mesh.Indices))
		}
		for i := 0; i < len(mesh.Indices); i += 3 {
			a := area(mesh.Vertices[mesh.Indices[i]], mesh.Vertices[mesh.Indices[i+1]], mesh.Vertices[mesh.Indices[i+2]])
			// Corners that collapse to a point make triangles with no area, which
			// have no winding to speak of and draw nothing.
			if a < -0.0001 {
				t.Fatalf("%s: triangle %d is wound the other way around", test.name, i/3)
			}
		}
	}
}

// TestBorderSides checks that a side of no width leaves that side of the border open.
func TestBorderSides(t *testing.T) {
	box := clay.BoundingBox{X: 0, Y: 0, Width: 100, Height: 100}
	full := Border(box, clay.CornerRadius{}, clay.BorderWidth{Left: 5, Right: 5, Top: 5, Bottom: 5}, 0)
	left := Border(box, clay.CornerRadius{}, clay.BorderWidth{Left: 5}, 0)
	if len(left.Indices) == 0 {
		t.Fatal("a border with only a left side has no triangles")
	}
	if len(left.Indices) >= len(full.Indices) {
		t.Errorf("a border with one side has %d indices, one with four has %d", len(left.Indices), len(full.Indices))
	}
	for _, i := range left.Indices {
		if v := left.Vertices[i]; v.X > box.Width/2 {
			t.Errorf("a left border reaches %v, past the middle of the box", v.X)
			break
		}
	}
}
