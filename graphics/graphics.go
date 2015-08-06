package graphics

//go:generate gostub Graphics

type Graphics interface {
	CreateFloat2AttributeArray(size int) (Float2AttributeArray, error)
	CreateFloat3AttributeArray(size int) (Float3AttributeArray, error)
	CreateIndexArray(size int) (IndexArray, error)
	CreateMaterial() (Material, error)
	Renderer() Renderer
}
