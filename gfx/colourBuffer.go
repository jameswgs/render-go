package gfx

type ColourBuffer struct {
	Buffer
	pixels []byte
}

func NewColourBuffer(width int, height int) *ColourBuffer {
	return &ColourBuffer{ Buffer{ width, height }, make( []byte, width*height*4 ) }
}

func (this *ColourBuffer) Write(x int, y int, color Colour) {
	if(x>=0 && x<this.width && y>=0 && y<this.height) {
		var slot = (y*this.width+x) * 4
		this.pixels[slot+0] = color.B
		this.pixels[slot+1] = color.G
		this.pixels[slot+2] = color.R
		this.pixels[slot+3] = color.A
	}
}

func (this *ColourBuffer) Pixels() []byte {
	return this.pixels
}
