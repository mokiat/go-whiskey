package graphics

import "github.com/momchil-atanasov/go-whiskey/common"

type ResourceID uint32
type BindLocation uint

const InvalidShaderId ResourceID = 0
const InvalidProgramId ResourceID = 0
const InvalidBufferId ResourceID = 0

type BufferUsage common.Enum

const (
	StreamDraw BufferUsage = iota
	StaticDraw
	DynamicDraw
	BUFFER_USAGE_COUNT int = iota
)

type Facade interface {
	CreateBuffer() ResourceID
	BindIndexBuffer(bufferId ResourceID)
	CreateIndexBufferData(data []byte, usage BufferUsage)
	DeleteBuffer(bufferId ResourceID)

	CreateVertexShader() ResourceID
	CreateFragmentShader() ResourceID
	SetShaderSourceCode(shaderId ResourceID, sourceCode string)
	CompileShader(shaderId ResourceID)
	DeleteShader(shaderId ResourceID)
}
