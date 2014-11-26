package vec4

type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func Null() Vec4 {
	return Vec4{}
}
