package graphics

// SequenceType is an enumeration that specifies that different
// types of primitives or shapes that can be rendered
type SequenceType int

const (
	// Triangles indicates that each three consecutive indices
	// will form a triangle
	Triangles SequenceType = iota

	// TriangleFan indicates that all the indices will form
	// a list of triangles that are connected in the shape
	// of a fan
	TriangleFan

	// TriangleStrip indicates that each index after the
	// second will be used to form a triangle will the previous
	// two
	TriangleStrip

	// Lines indicates that each two consecutive indices
	// will form a line
	Lines
)

//go:generate gostub IndexArray

// IndexArray represents an array of indices to be used
// for the construction of the mesh and its rendering
type IndexArray interface {

	// Size returns the number of indices in this IndexArray
	Size() int

	// PutIndex can be used to configure the index at the
	// specified position in this IndexArray
	PutIndex(position int, index uint16)

	// Index can be used to read the index at the specified
	// position in this IndexArray
	Index(position int) uint16
}
