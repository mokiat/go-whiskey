package graphics

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
)
