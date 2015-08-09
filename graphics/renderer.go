package graphics

//go:generate gostub Renderer

type Renderer interface {
	Initialize() error
	UseMaterial(Material)
	BindAttribute(AttributeName, AttributeArray)
	BindUniform(UniformName, Uniform)
	Render(IndexArray, SequenceType)
	Flush() error
}
