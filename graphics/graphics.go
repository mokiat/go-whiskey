package graphics

//go:generate gostub Graphics

type Graphics interface {
	CreateFloat2AttributeArray(size int) (Float2AttributeArray, error)
	CreateFloat3AttributeArray(size int) (Float3AttributeArray, error)
	CreateFloat4Uniform() (Float4Uniform, error)
	CreateFloat4x4Uniform() (Float4x4Uniform, error)
	CreateIndexArray(size int) (IndexArray, error)
	CreateMaterial() (Material, error)
	Renderer() Renderer
}
