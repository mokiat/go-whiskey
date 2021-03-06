package math

import "fmt"

type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (v Vec4) Null() bool {
	return v == Vec4{}
}

func (v Vec4) Mul(amount float32) Vec4 {
	return Vec4{
		v.X * amount,
		v.Y * amount,
		v.Z * amount,
		v.W * amount,
	}
}

func (v Vec4) Div(amount float32) Vec4 {
	return Vec4{
		v.X / amount,
		v.Y / amount,
		v.Z / amount,
		v.W / amount,
	}
}

func (v Vec4) GoString() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}

func NullVec4() Vec4 {
	return Vec4{}
}

func MakeVec4(x, y, z, w float32) Vec4 {
	return Vec4{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}
