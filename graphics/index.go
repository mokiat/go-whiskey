package graphics

type SequenceType int

const (
	Triangles SequenceType = iota
	TriangleFan
	TriangleStrip
	Lines
)

//go:generate gostub IndexArray

type IndexArray interface {
	Size() int
	PutIndex(position int, index uint16)
	Index(position int) uint16
}
