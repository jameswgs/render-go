package gfx

import (
	"rendergo/vector"
)

type Buffer struct {
	width int
	height int
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{ width, height }
}

func (this *Buffer) Size() *vector.Size2 {
	return &vector.Size2{this.width, this.height}
}

func (this *Buffer) Count() int {
	return this.width * this.height
}