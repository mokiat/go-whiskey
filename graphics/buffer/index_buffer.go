package buffer

import (
	"github.com/momchil-atanasov/go-whiskey/common/buf"
	"github.com/momchil-atanasov/go-whiskey/graphics"
)

type IndexBuffer interface {
	Id() graphics.ResourceId
	Usage() graphics.BufferUsage
	Size() int
	SetValue(position int, value uint16)
	Value(position int) uint16
	BindRemotely()
	CreateRemotely() error
	DeleteRemotely()
	CreatedRemotely() bool
}

type indexBuffer struct {
	id     graphics.ResourceId
	facade graphics.Facade
	buffer buf.UInt16Buffer
	usage  graphics.BufferUsage
	size   int
}

func NewIndexBuffer(facade graphics.Facade, usage graphics.BufferUsage, size int) IndexBuffer {
	data := make([]byte, size*2)
	return &indexBuffer{
		id:     graphics.InvalidBufferId,
		facade: facade,
		buffer: buf.UInt16Buffer(data),
		usage:  usage,
		size:   size,
	}
}

func (b *indexBuffer) Id() graphics.ResourceId {
	return b.id
}

func (b *indexBuffer) Usage() graphics.BufferUsage {
	return b.usage
}

func (b *indexBuffer) Size() int {
	return b.size
}

func (b *indexBuffer) SetValue(position int, value uint16) {
	b.buffer.Set(position, value)
}

func (b *indexBuffer) Value(position int) uint16 {
	return b.buffer.Get(position)
}

func (b *indexBuffer) CreateRemotely() error {
	var err error
	b.id, err = b.facade.CreateBuffer()
	if err != nil {
		b.id = graphics.InvalidBufferId
		return err
	}
	b.facade.BindIndexBuffer(b.id)
	b.facade.CreateIndexBufferData(b.buffer, b.usage)
	return nil
}

func (b *indexBuffer) BindRemotely() {
	b.facade.BindIndexBuffer(b.id)
}

func (b *indexBuffer) DeleteRemotely() {
	b.facade.DeleteBuffer(b.id)
	b.id = graphics.InvalidBufferId
}

func (b *indexBuffer) CreatedRemotely() bool {
	return b.id != graphics.InvalidBufferId
}
