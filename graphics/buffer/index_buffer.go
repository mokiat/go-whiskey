package buffer

import (
	"encoding/binary"

	"github.com/momchil-atanasov/go-whiskey/common"
	"github.com/momchil-atanasov/go-whiskey/graphics"
)

type IndexBuffer interface {
	Id() graphics.ResourceId
	Usage() graphics.BufferUsage
	Size() int
	BindRemotely()
	CreateRemotely()
	DeleteRemotely()
	CreatedRemotely() bool
}

type indexBuffer struct {
	id     graphics.ResourceId
	facade graphics.Facade
	buffer common.UInt16Buffer
	usage  graphics.BufferUsage
	size   int
}

func NewIndexBuffer(facade graphics.Facade, usage graphics.BufferUsage, size int) IndexBuffer {
	return &indexBuffer{
		id:     graphics.InvalidBufferId,
		facade: facade,
		buffer: common.NewUInt16Buffer(size, binary.LittleEndian),
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

func (b *indexBuffer) CreateRemotely() {
	b.id = b.facade.CreateBuffer()
	b.facade.BindIndexBuffer(b.id)
	b.facade.CreateIndexBufferData(b.buffer.Bytes(), b.usage)
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
