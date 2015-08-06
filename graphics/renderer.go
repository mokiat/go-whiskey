package graphics

import "github.com/momchil-atanasov/go-whiskey/math"

//go:generate gostub Renderer

type Renderer interface {
	Initialize() error
	UseMaterial(Material)
	BindFloat2AttributeArray(AttributeName, Float2AttributeArray)
	BindFloat3AttributeArray(AttributeName, Float3AttributeArray)
	BindFloat4x4Uniform(UniformName, math.Mat4x4)
	BindFloat4Uniform(UniformName, math.Vec4)
	Render(ElementType, IndexArray)
	Flush() error
}
