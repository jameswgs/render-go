package gfx

type buffer struct {
	width int
	height int
	buffer []byte
}

func NewBuffer(width int, height int) *buffer {
	return &buffer{ width, height, make( []byte, width*height*4 ) }
}

func (this *buffer) Write(x int, y int, color Colour) {
	if(x>=0 && x<this.width && y>=0 && y<this.height) {
		var slot = (y*this.width+x) * 4
		this.buffer[slot+0] = color.B
		this.buffer[slot+1] = color.G
		this.buffer[slot+2] = color.R
		this.buffer[slot+3] = color.A
	}
}

func (this *buffer) Pixels() []byte {
	return this.buffer
}