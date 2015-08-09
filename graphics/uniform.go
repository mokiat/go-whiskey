package graphics

import "github.com/momchil-atanasov/go-whiskey/math"

type UniformName int

const (
	ProjectionMatrix UniformName = iota
	ModelViewMatrix
	ModelMatrix
	ViewMatrix
	AmbientColor
	DiffuseColor
	DiffuseTexture
	SpecularColor
	ReflectionTexture
)

type Uniform interface {
}

type Float4Uniform interface {
	Uniform
	SetValue(math.Vec4)
	Value() math.Vec4
}

type Float4x4Uniform interface {
	Uniform
	SetValue(math.Mat4x4)
	Value() math.Mat4x4
}
