package gfx

import (
	"rendergo/vector"
)

type Tri struct {
	A vector.V2
	B vector.V2
	C vector.V2
	Colour Colour
}

func NewTri( a vector.V2, b vector.V2, c vector.V2, col Colour ) Tri {
	return Tri{ a, b, c, col }
}
