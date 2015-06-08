package buffer

import (
	"github.com/momchil-atanasov/go-whiskey/common/buf"
	"github.com/momchil-atanasov/go-whiskey/graphics/client"
)

//go:generate counterfeiter -o buffer_fakes/fake_index_buffer_data.go ./ IndexBufferData

// IndexBufferData describes the local data of an
// IndexBuffer object.
type IndexBufferData interface {
	SetValue(position int, value uint16)
	Value(position int) uint16
	Size() int
	Content() []byte
}

// NewIndexBufferData returns a IndexBufferData object of
// the specified size
func NewIndexBufferData(size int) IndexBufferData {
	data := make([]byte, size*2)
	return &indexBufferData{
		size:   size,
		buffer: buf.UInt16Buffer(data),
	}
}

type indexBufferData struct {
	size   int
	buffer buf.UInt16Buffer
}

func (d *indexBufferData) SetValue(position int, value uint16) {
	d.buffer.Set(position, value)
}

func (d *indexBufferData) Value(position int) uint16 {
	return d.buffer.Get(position)
}

func (d *indexBufferData) Size() int {
	return d.size
}

func (d *indexBufferData) Content() []byte {
	return d.buffer
}

//go:generate counterfeiter -o buffer_fakes/fake_index_buffer.go ./ IndexBuffer

// IndexBuffer describes an index buffer object, which contains
// the index data necessary to render a mesh.
type IndexBuffer interface {
	// Data returns the local data representation of this buffer.
	Data() IndexBufferData

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

// NewIndexBuffer creates a new IndexBuffer instance with the
// specified data and expected usage
func NewIndexBuffer(data IndexBufferData, usage client.BufferUsage) IndexBuffer {
	return &indexBuffer{
		data:  data,
		usage: usage,
	}
}

type indexBuffer struct {
	id    client.BufferId
	usage client.BufferUsage
	data  IndexBufferData
}

func (b *indexBuffer) Data() IndexBufferData {
	return b.data
}

func (b *indexBuffer) Usage() client.BufferUsage {
	return b.usage
}

func (b *indexBuffer) Id() client.BufferId {
	return b.id
}

func (b *indexBuffer) Created() bool {
	return b.id != nil
}

func (b *indexBuffer) Create(bufferClient client.BufferClient) error {
	var err error
	b.id, err = bufferClient.CreateBuffer()
	if err != nil {
		return err
	}
	err = bufferClient.BindIndexBuffer(b.id)
	if err != nil {
		return err
	}
	err = bufferClient.CreateIndexBufferData(b.data.Content(), b.usage)
	if err != nil {
		return err
	}
	return nil
}

func (b *indexBuffer) Delete(bufferClient client.BufferClient) error {
	err := bufferClient.DeleteBuffer(b.id)
	if err != nil {
		return err
	}
	b.id = nil
	return nil
}
