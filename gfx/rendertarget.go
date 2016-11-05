package gfx

import ( 
	"rendergo/vector"
	"sort"
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

	fx0 := 0.0
	fx1 := 0.0

	for y:=t0.Y; y<t1.Y; y++ {
		x0 := t0.X + int(fx0)
		x1 := t0.X + int(fx1)
		for x:=x0; x<=x1; x++ {
			this.buffer.Write(x,y,tri.Colour)	
		}
		fx0 += slope0
		fx1 += slope1
	}

	fx0 = 0.0

	for y:=t1.Y; y<=t2.Y; y++ {
		x0 := t1.X + int(fx0)
		x1 := t0.X + int(fx1)
		for x:=x0; x<=x1; x++ {
			this.buffer.Write(x,y,tri.Colour)	
		}
		fx0 += slope2
		fx1 += slope1
	}

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

	xscale := float64(this.buffer.Size().X)
	yscale := float64(this.buffer.Size().Y)

	for _, f := range object.Faces {

		println(len(object.Verts), f.A, f.B, f.C)

		vA := object.Verts[f.A]
		vB := object.Verts[f.B]
		vC := object.Verts[f.C]

		v0 := vector.V2{int(vA.X*xscale),int(vA.Y*yscale)}
		v1 := vector.V2{int(vB.X*xscale),int(vB.Y*yscale)} 
		v2 := vector.V2{int(vC.X*xscale),int(vC.Y*yscale)}

		println("[",v0.X,v0.Y,"] [",v1.X,v1.Y,"] [",v2.X,v2.Y,"]")

		col := Colour{0xFF, 0x00, 0xFF, 0xFF} // Colour{ byte(rand.Intn(0x100)), byte(rand.Intn(0x100)), byte(rand.Intn(0x100)), 0xFF } 

		tri := NewTri(v0, v1, v2, col)

		this.DrawTri(tri)
	}
}
