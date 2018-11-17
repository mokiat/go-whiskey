package buf

func CreateUInt16BufferWriter(buffer UInt16Buffer) *UInt16BufferWriter {
	return &UInt16BufferWriter{
		buffer: buffer,
	}
}

type UInt16BufferWriter struct {
	buffer UInt16Buffer
	index  int
}

func (w *UInt16BufferWriter) PutValue(x uint16) {
	w.buffer.Set(w.index, x)
	w.index++
}
