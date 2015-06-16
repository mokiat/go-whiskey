package client

type ShaderId interface{}
type ProgramId interface{}

//go:generate counterfeiter -o client_fakes/fake_shader.go ./ ShaderClient

type ShaderClient interface {
	CreateVertexShader() (ShaderId, error)
	CreateFragmentShader() (ShaderId, error)
	SetShaderSourceCode(id ShaderId, sourceCode string) error
	CompileShader(id ShaderId) error
	DeleteShader(id ShaderId) error

	CreateProgram() (ProgramId, error)
	AttachShaderToProgram(shaderId ShaderId, programId ProgramId) error
	LinkProgram(ProgramId) error
	GetProgramAttributes(ProgramId) ([]AttributeDeclaration, error)
	GetProgramUniforms(ProgramId) ([]UniformDeclaration, error)
	UseProgram(ProgramId) error
	DeleteProgram(ProgramId) error
}
