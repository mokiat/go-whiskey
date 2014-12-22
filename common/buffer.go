package common

import (
	"encoding/binary"
	"math"
)

type UInt16Buffer interface {
	Size() int
	PutUInt16(position int, value uint16)
	UInt16(position int) uint16
	Bytes() []byte
}

type buffer struct {
	littleEndian bool
	data         []byte
	size         int
}

func (a *buffer) Size() int {
	return a.size
}

func (a *buffer) Bytes() []byte {
	return a.data
}

type uInt16Buffer struct {
	buffer
}

func NewUInt16Buffer(data []byte, order binary.ByteOrder) UInt16Buffer {
	return &uInt16Buffer{
		buffer: buffer{
			littleEndian: order == binary.LittleEndian,
			data:         data,
			size:         cap(data) / 2,
		},
	}
}

func (a *uInt16Buffer) PutUInt16(position int, value uint16) {
	offset := position * 2
	if a.littleEndian {
		a.data[offset+0] = byte(value >> 0)
		a.data[offset+1] = byte(value >> 8)
	} else {
		a.data[offset+0] = byte(value >> 8)
		a.data[offset+1] = byte(value >> 0)
	}
}

func (a *uInt16Buffer) UInt16(position int) uint16 {
	offset := position * 2
	if a.littleEndian {
		return uint16(a.data[offset+0]) + uint16(a.data[offset+1])<<8
	} else {
		return uint16(a.data[offset+0])<<8 + uint16(a.data[offset+1])
	}
}

type Float32Buffer interface {
	Size() int
	PutFloat32(position int, value float32)
	Float32(position int) float32
	Bytes() []byte
}

type float32Buffer struct {
	buffer
}

func NewFloat32Buffer(data []byte, order binary.ByteOrder) Float32Buffer {
	return &float32Buffer{
		buffer: buffer{
			littleEndian: order == binary.LittleEndian,
			data:         data,
			size:         cap(data) / 4,
		},
	}
}

func (a *float32Buffer) PutFloat32(position int, value float32) {
	offset := position * 4
	bits := math.Float32bits(value)
	if a.littleEndian {
		a.data[offset+0] = byte(bits >> 0)
		a.data[offset+1] = byte(bits >> 8)
		a.data[offset+2] = byte(bits >> 16)
		a.data[offset+3] = byte(bits >> 24)
	} else {
		a.data[offset+0] = byte(bits >> 24)
		a.data[offset+1] = byte(bits >> 16)
		a.data[offset+2] = byte(bits >> 8)
		a.data[offset+3] = byte(bits >> 0)
	}
}

func (a *float32Buffer) Float32(position int) float32 {
	offset := position * 4
	var bits uint32
	if a.littleEndian {
		bits = uint32(a.data[offset+0])<<0 +
			uint32(a.data[offset+1])<<8 +
			uint32(a.data[offset+2])<<16 +
			uint32(a.data[offset+3])<<24
	} else {
		bits = uint32(a.data[offset+0])<<24 +
			uint32(a.data[offset+1])<<16 +
			uint32(a.data[offset+2])<<8 +
			uint32(a.data[offset+3])<<0
	}
	return math.Float32frombits(bits)
}
