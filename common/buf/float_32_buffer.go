package buf

import "math"

type Float32Buffer []byte

func (b Float32Buffer) Length() int {
	return len(b) / 4
}

func (b Float32Buffer) Set(position int, value float32) {
	offset := position * 4
	bits := math.Float32bits(value)
	b[offset+0] = byte(bits >> 0)
	b[offset+1] = byte(bits >> 8)
	b[offset+2] = byte(bits >> 16)
	b[offset+3] = byte(bits >> 24)
}

func (b Float32Buffer) Get(position int) float32 {
	offset := position * 4
	bits := uint32(b[offset+0])<<0 +
		uint32(b[offset+1])<<8 +
		uint32(b[offset+2])<<16 +
		uint32(b[offset+3])<<24
	return math.Float32frombits(bits)
}

func (b Float32Buffer) Range(position, length int) Float32Buffer {
	offset := position * 4
	size := length * 4
	return b[offset : offset+size]
}
