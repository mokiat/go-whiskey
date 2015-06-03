package client

//go:generate counterfeiter -o client_fakes/fake_shader.go ./ ShaderClient

type ShaderId interface{}
type ProgramId interface{}

type ShaderClient interface {
	CreateVertexShader() (ShaderId, error)
	CreateFragmentShader() (ShaderId, error)
	SetShaderSourceCode(id ShaderId, sourceCode string) error
	CompileShader(id ShaderId) error
	DeleteShader(id ShaderId) error

	CreateProgram() (ProgramId, error)
	AttachShaderToProgram(shaderId ShaderId, programId ProgramId) error
	LinkProgram(ProgramId) error
	UseProgram(ProgramId) error
	DeleteProgram(ProgramId) error
}
