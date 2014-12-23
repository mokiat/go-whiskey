package buffer

import (
	"encoding/binary"

	"github.com/momchil-atanasov/go-whiskey/common"
	"github.com/momchil-atanasov/go-whiskey/graphics"
)

type VertexBuffer interface {
	Id() graphics.ResourceID
	Usage() graphics.BufferUsage
	Size() int
	BindRemotely()
	CreateRemotely()
	DeleteRemotely()
	CreatedRemotely() bool
}

type vertexBuffer struct {
	id     graphics.ResourceID
	facade graphics.Facade
	buffer common.Float32Buffer
	size   int
	usage  graphics.BufferUsage
}

func NewVertexBuffer(facade graphics.Facade, size int, usage graphics.BufferUsage) VertexBuffer {
	return &vertexBuffer{
		id:     graphics.InvalidBufferId,
		facade: facade,
		buffer: common.NewFloat32Buffer(size, binary.LittleEndian),
		size:   size,
		usage:  usage,
	}
}

func (b *vertexBuffer) Id() graphics.ResourceID {
	return b.id
}

func (b *vertexBuffer) Usage() graphics.BufferUsage {
	return b.usage
}

func (b *vertexBuffer) Size() int {
	return b.size
}

func (b *vertexBuffer) BindRemotely() {
	b.facade.BindVertexBuffer(b.id)
}

func (b *vertexBuffer) CreateRemotely() {
	b.id = b.facade.CreateBuffer()
	b.facade.BindVertexBuffer(b.id)
	b.facade.CreateVertexBufferData(b.buffer.Bytes(), b.usage)
}

func (b *vertexBuffer) DeleteRemotely() {
	b.facade.DeleteBuffer(b.id)
	b.id = graphics.InvalidBufferId
}

func (b *vertexBuffer) CreatedRemotely() bool {
	return b.id != graphics.InvalidBufferId
}
