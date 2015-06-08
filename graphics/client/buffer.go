package client

type BufferId interface{}

type BufferUsage int

const (
	StreamDrawBufferUsage BufferUsage = iota
	StaticDrawBufferUsage
	DynamicDrawBufferUsage
)

//go:generate counterfeiter -o client_fakes/fake_buffer.go ./ BufferClient

type BufferClient interface {
	CreateBuffer() (BufferId, error)
	BindVertexBuffer(id BufferId) error
	BindIndexBuffer(id BufferId) error
	CreateVertexBufferData(data []byte, usage BufferUsage) error
	CreateIndexBufferData(data []byte, usage BufferUsage) error
	DeleteBuffer(bufferId BufferId) error
}
