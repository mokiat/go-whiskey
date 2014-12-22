package graphics

type ResourceID uint32

const InvalidShaderId ResourceID = 0

type Facade interface {
	CreateVertexShader() ResourceID
	CreateFragmentShader() ResourceID
	SetShaderSourceCode(shaderId ResourceID, sourceCode string)
	CompileShader(shaderId ResourceID)
	DeleteShader(shaderId ResourceID)
}
