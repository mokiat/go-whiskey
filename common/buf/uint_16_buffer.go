package buf

type UInt16Buffer []byte

func CreateUInt16Buffer(size int) UInt16Buffer {
	return UInt16Buffer(make([]byte, size*2))
}

func (b UInt16Buffer) Length() int {
	return len(b) / 2
}

func (b UInt16Buffer) Set(position int, value uint16) {
	offset := position * 2
	b[offset+0] = byte(value >> 0)
	b[offset+1] = byte(value >> 8)
}

func (b UInt16Buffer) Get(position int) uint16 {
	offset := position * 2
	return uint16(b[offset+0]) + uint16(b[offset+1])<<8
}

func (b UInt16Buffer) Range(position, length int) UInt16Buffer {
	offset := position * 2
	size := length * 2
	return b[offset : offset+size]
}
