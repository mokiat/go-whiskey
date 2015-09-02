package graphics

//go:generate gostub Renderer

type Renderer interface {
	Initialize() error
	UseMaterial(Material)
	BindAttribute(AttributeName, AttributeArray)
	BindUniform(UniformName, Uniform)
	BindTexture(TextureName, Texture)
	Render(IndexArray, SequenceType)
	Flush() error
}
