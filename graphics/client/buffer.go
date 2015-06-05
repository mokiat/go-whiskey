package client

type BufferId interface{}

//go:generate counterfeiter -o client_fakes/fake_buffer.go ./ BufferClient

type BufferClient interface {
	BindVertexBuffer(id BufferId) error
	BindIndexBuffer(id BufferId) error
}
