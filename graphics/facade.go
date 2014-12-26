package graphics

import "github.com/momchil-atanasov/go-whiskey/common"

type ResourceId uint32
type BindLocation uint

const InvalidShaderId ResourceId = 0
const InvalidProgramId ResourceId = 0
const InvalidBufferId ResourceId = 0

type BufferUsage common.Enum

const (
	StreamDraw BufferUsage = iota
	StaticDraw
	DynamicDraw
	BUFFER_USAGE_COUNT int = iota
)

type Facade interface {
	CreateBuffer() ResourceId
	BindIndexBuffer(bufferId ResourceId)
	CreateIndexBufferData(data []byte, usage BufferUsage)
	BindVertexBuffer(bufferId ResourceId)
	CreateVertexBufferData(data []byte, usage BufferUsage)
	DeleteBuffer(bufferId ResourceId)

	CreateVertexShader() (ResourceId, error)
	CreateFragmentShader() (ResourceId, error)
	SetShaderSourceCode(shaderId ResourceId, sourceCode string)
	CompileShader(shaderId ResourceId) error
	DeleteShader(shaderId ResourceId)

	CreateProgram() (ResourceId, error)
	AttachShaderToProgram(programId ResourceId, shaderId ResourceId)
	LinkProgram(ResourceId) error
	UseProgram(ResourceId)
	DeleteProgram(ResourceId)
}
