package graphics

const InvalidShaderId = 0

type Facade interface {
	CreateVertexShader() int
	CreateFragmentShader() int
	SetShaderSourceCode(shaderId int, sourceCode string)
	CompileShader(shaderId int)
	DeleteShader(shaderId int)
}
