package buffer

import (
	"github.com/momchil-atanasov/go-whiskey/common/buf"
	"github.com/momchil-atanasov/go-whiskey/graphics"
)

type VertexBuffer interface {
	Id() graphics.ResourceId
	Usage() graphics.BufferUsage
	Size() int
	SetValue(position int, value float32)
	Value(position int) float32
	BindRemotely()
	CreateRemotely() error
	DeleteRemotely()
	CreatedRemotely() bool
}

type vertexBuffer struct {
	id     graphics.ResourceId
	facade graphics.Facade
	buffer buf.Float32Buffer
	size   int
	usage  graphics.BufferUsage
}

func NewVertexBuffer(facade graphics.Facade, size int, usage graphics.BufferUsage) VertexBuffer {
	data := make([]byte, size*4)
	return &vertexBuffer{
		id:     graphics.InvalidBufferId,
		facade: facade,
		buffer: buf.Float32Buffer(data),
		size:   size,
		usage:  usage,
	}
}

func (b *vertexBuffer) Id() graphics.ResourceId {
	return b.id
}

func (b *vertexBuffer) Usage() graphics.BufferUsage {
	return b.usage
}

func (b *vertexBuffer) Size() int {
	return b.size
}

func (b *vertexBuffer) SetValue(position int, value float32) {
	b.buffer.Set(position, value)
}

func (b *vertexBuffer) Value(position int) float32 {
	return b.buffer.Get(position)
}

func (b *vertexBuffer) BindRemotely() {
	b.facade.BindVertexBuffer(b.id)
}

func (b *vertexBuffer) CreateRemotely() error {
	var err error
	b.id, err = b.facade.CreateBuffer()
	if err != nil {
		b.id = graphics.InvalidBufferId
		return err
	}
	b.facade.BindVertexBuffer(b.id)
	b.facade.CreateVertexBufferData(b.buffer, b.usage)
	return nil
}

func (b *vertexBuffer) DeleteRemotely() {
	b.facade.DeleteBuffer(b.id)
	b.id = graphics.InvalidBufferId
}

func (b *vertexBuffer) CreatedRemotely() bool {
	return b.id != graphics.InvalidBufferId
}
