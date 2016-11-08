package vector

import (
	"math"
)

type V2 struct {
	X int
	Y int
}

type V3f struct {
	X, Y, Z float64
}

func (this V3f) Sub(other V3f) V3f {
	return V3f{ this.X-other.X, this.Y-other.Y, this.Z-other.X }
}

func (this V3f) Cross(other V3f) V3f {
	return V3f{ this.Y * other.Z - this.Z * other.Y, 
				this.Z * other.X - this.X * other.Z, 
				this.X * other.Y - this.Y * other.X } 
}

func (this V3f) Normalised() V3f {
	invmag := 1.0 / this.Magnitude()
	return V3f{ this.X*invmag, this.Y*invmag, this.Z*invmag }
}

func (this V3f) Magnitude() float64 {
	return math.Sqrt( this.X * this.X + this.Y * this.Y + this.Z * this.Z )
}
