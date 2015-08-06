package graphics

type ElementType int

const (
	Triangles ElementType = iota
	Lines
)

//go:generate gostub IndexArray

type IndexArray interface {
	Size() int
	PutIndex(position int, index uint16)
	Index(position int) uint16
}
