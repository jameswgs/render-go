package gfx

type ZBuffer struct {
	Buffer
	z []float64
}

func NewZBuffer(width int, height int) *ZBuffer {
	return &ZBuffer{ Buffer{ width, height }, make( []float64, width*height ) }
}

func (this *ZBuffer) Clear(far float64) {
	for i:=0; i<this.Count(); i++ {
		this.z[i] = far
	}
}

func (this *ZBuffer) WriteMin(x int, y int, z float64) bool {
	min := false
	if(x>=0 && x<this.width && y>=0 && y<this.height) {
		var slot = (y*this.width+x)
		if(z<this.z[slot]) {
			this.z[slot] = z
			min = true
		}
	}
	return min
}
