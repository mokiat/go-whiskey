package buf

func CreateFloat32BufferWriter(buffer Float32Buffer) *Float32BufferWriter {
	return &Float32BufferWriter{
		buffer: buffer,
	}
}

type Float32BufferWriter struct {
	buffer Float32Buffer
	index  int
}

func (w *Float32BufferWriter) PutValue(x float32) {
	w.buffer.Set(w.index, x)
	w.index++
}

func (w *Float32BufferWriter) PutValue2(x, y float32) {
	w.buffer.Set(w.index+0, x)
	w.buffer.Set(w.index+1, y)
	w.index += 2
}

func (w *Float32BufferWriter) PutValue3(x, y, z float32) {
	w.buffer.Set(w.index+0, x)
	w.buffer.Set(w.index+1, y)
	w.buffer.Set(w.index+2, z)
	w.index += 3
}
