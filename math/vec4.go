package math

type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (v Vec4) Null() bool {
	return v == Vec4{}
}

func NullVec4() Vec4 {
	return Vec4{}
}
