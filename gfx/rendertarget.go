package gfx

type RenderTarget struct {
	buffer *Buffer
}

func NewRenderTarget(buffer *Buffer) *RenderTarget {
	return &RenderTarget{buffer}
}

func (this *RenderTarget) Draw(tri Tri) {
	minX := Min(tri.A.X, Min(tri.B.X, tri.C.X))
	minY := Min(tri.A.Y, Min(tri.B.Y, tri.C.Y))
	maxX := Max(tri.A.X, Max(tri.B.X, tri.C.X))
	maxY := Max(tri.A.Y, Max(tri.B.Y, tri.C.Y))

	for y := minY; y<=maxY; y++ {
		for x := minX; x<=maxX; x++ {
			if(Contains(tri,x,y)) {
				buffer.Write(x,y,tri.Colour)
			}
		}
	}
}

func Contains(tri Tri, int x, int y) bool {
	// find slope of a to (xy) = sap
	// find slope of a to b = sab
	// find slope of a to c = sac
	// return no if sap is not between sab and sac
	
}

func Max(a int, b int) int {
	if a > b { return a } else { return b }
}

func Min(a int, b int) int {
	if a < b { return a } else { return b }
}
