package gfx

import ( 
	"rendergo/vector"
	"sort"
	"fmt"
	"rendergo/obj"
)

type RenderTarget struct {
	buffer *Buffer
}

func NewRenderTarget(buffer *Buffer) *RenderTarget {
	return &RenderTarget{buffer}
}

func (this *RenderTarget) DrawTri(tri Tri) {

	tris := []vector.V2 { tri.A, tri.B, tri.C }

	sort.Sort(ByY(tris))

	t0 := tris[0]
	t1 := tris[1]
	t2 := tris[2]

	slope0 := float64(t0.X-t1.X)/float64(t0.Y-t1.Y)
	slope1 := float64(t2.X-t0.X)/float64(t2.Y-t0.Y)
	slope2 := float64(t2.X-t1.X)/float64(t2.Y-t1.Y)

	fx1 := this.DrawTriPart(t0.Y, t1.Y, t0.X, t0.X, tri.Colour, 0.0, 0.0, slope0, slope1)

	this.DrawTriPart(t1.Y, t2.Y, t1.X, t0.X, tri.Colour, 0.0, fx1, slope2, slope1)

	this.buffer.Write(t0.X,t0.Y,Colour{0xFF,0x00,0x00,0xFF});
	this.buffer.Write(t1.X,t1.Y,Colour{0x00,0xFF,0x00,0xFF});
	this.buffer.Write(t2.X,t2.Y,Colour{0x00,0x00,0xFF,0xFF});

}

func (this *RenderTarget) DrawTriPart(startY int, endY int, offsetX0 int, offsetX1 int, colour Colour, fx0 float64, fx1 float64, s0 float64, s1 float64) float64 {

	for y:=startY; y<endY; y++ {
		x0 := offsetX0 + int(fx0)
		x1 := offsetX1 + int(fx1)
		adder := 1
		if(x0>x1) { adder = -1 }
		for x:=x0; x!=(x1+adder); x+=adder {
			this.buffer.Write(x,y,colour)	
		}
		fx0 += s0
		fx1 += s1
	}

	return fx1
}

type ByY []vector.V2

func (a ByY) Len() int           { return len(a) }
func (a ByY) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByY) Less(i, j int) bool { return a[i].Y < a[j].Y }

func (this *RenderTarget) Fill(colour Colour) {
	size := this.buffer.Size()
	for y := 0; y<=size.Y; y++ {
		for x := 0; x<=size.X; x++ {
			this.buffer.Write(x,y,colour)
		}
	}
}

func Max(a int, b int) int {
	if a > b { return a } else { return b }
}

func Min(a int, b int) int {
	if a < b { return a } else { return b }
}

func (this *RenderTarget) DrawObject(object *obj.Object) {

	hw := this.buffer.Size().X/2
	hh := this.buffer.Size().Y/2

	xscale := float64(hw)
	yscale := float64(hh)

	xoffset := hw
	yoffset := hh

	for _, f := range object.Faces {

		vA := object.Verts[f.A]
		vB := object.Verts[f.B]
		vC := object.Verts[f.C]

		v0 := vector.V2{int(vA.X*xscale)+xoffset,int(vA.Y*yscale)+yoffset}
		v1 := vector.V2{int(vB.X*xscale)+xoffset,int(vB.Y*yscale)+yoffset} 
		v2 := vector.V2{int(vC.X*xscale)+xoffset,int(vC.Y*yscale)+yoffset}

		fmt.Printf("vector.V2{%d, %d}, vector.V2{%d, %d}, vector.V2{%d, %d}\n", v0.X, v0.Y, v1.X, v1.Y, v2.X, v2.Y)

		col := Colour{0xFF, 0xFF, 0xFF, 0xFF} // Colour{ byte(rand.Intn(0x100)), byte(rand.Intn(0x100)), byte(rand.Intn(0x100)), 0xFF } 

		tri := NewTri(v0, v1, v2, col)

		this.DrawTri(tri)
	}
}
