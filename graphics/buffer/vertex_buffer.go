package buffer

import (
	"github.com/momchil-atanasov/go-whiskey/common/buf"
	"github.com/momchil-atanasov/go-whiskey/graphics/client"
)

//go:generate counterfeiter -o buffer_fakes/fake_vertex_buffer_data.go ./ VertexBufferData

// VertexBufferData describes the local data of a
// VertexBuffer object.
type VertexBufferData interface {
	SetValue(position int, value float32)
	Value(position int) float32
	Size() int
	Content() []byte
}

// NewVertexBufferData returns a VertexBufferData object of
// the specified size
func NewVertexBufferData(size int) VertexBufferData {
	data := make([]byte, size*4)
	return &vertexBufferData{
		size:   size,
		buffer: buf.Float32Buffer(data),
	}
}

type vertexBufferData struct {
	size   int
	buffer buf.Float32Buffer
}

func (d *vertexBufferData) SetValue(position int, value float32) {
	d.buffer.Set(position, value)
}

func (d *vertexBufferData) Value(position int) float32 {
	return d.buffer.Get(position)
}

func (d *vertexBufferData) Size() int {
	return d.size
}

func (d *vertexBufferData) Content() []byte {
	return d.buffer
}

//go:generate counterfeiter -o buffer_fakes/fake_vertex_buffer.go ./ VertexBuffer

// VertexBuffer describes a vertex buffer object, which contains
// the attribute data necessary to render a mesh.
type VertexBuffer interface {
	// Data returns the local data representation of this buffer.
	Data() VertexBufferData

	// Usage returns the usage of the buffer
	Usage() client.BufferUsage

	// Id returns the Id of the buffer
	Id() client.BufferId

	// Created returns whether this buffer has been created on
	// the graphics card
	Created() bool

	// Create creates this buffer on the graphics card
	Create(client.BufferClient) error

	// Delete deletes this buffer from the graphics card
	Delete(client.BufferClient) error
}

// NewVertexBuffer creates a new VertexBuffer instance with the
// specified data and expected usage
func NewVertexBuffer(data VertexBufferData, usage client.BufferUsage) VertexBuffer {
	return &vertexBuffer{
		data:  data,
		usage: usage,
	}
}

type vertexBuffer struct {
	id    client.BufferId
	usage client.BufferUsage
	data  VertexBufferData
}

func (b *vertexBuffer) Data() VertexBufferData {
	return b.data
}

func (b *vertexBuffer) Usage() client.BufferUsage {
	return b.usage
}

func (b *vertexBuffer) Id() client.BufferId {
	return b.id
}

func (b *vertexBuffer) Created() bool {
	return b.id != nil
}

func (b *vertexBuffer) Create(bufferClient client.BufferClient) error {
	var err error
	b.id, err = bufferClient.CreateBuffer()
	if err != nil {
		return err
	}
	err = bufferClient.BindVertexBuffer(b.id)
	if err != nil {
		return err
	}
	err = bufferClient.CreateVertexBufferData(b.data.Content(), b.usage)
	if err != nil {
		return err
	}
	return nil
}

func (b *vertexBuffer) Delete(bufferClient client.BufferClient) error {
	err := bufferClient.DeleteBuffer(b.id)
	if err != nil {
		return err
	}
	b.id = nil
	return nil
}
