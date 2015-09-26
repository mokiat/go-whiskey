package graphics

// SequenceType is an enumeration that specifies that different
// types of primitives or shapes that can be rendered
type SequenceType int

const (
	// Points indicates that each consecutive element will form a
	// point to be rendered.
	Points SequenceType = iota

	// LineStrips indicates that the first two elements will be used to
	// form a single line and each consecutive element will extend on the
	// end of each previous line.
	LineStrips

	// LineLoops is similar to LineStrips, except that the end of the last
	// line is connected with the beginning of the first line to form
	// one last line.
	LineLoops

	// Lines indicates that each two consecutive indices
	// will form a line.
	Lines

	// TriangleStrip indicates that the first three elements will be used
	// to construct a triangle and each consecutive element will be used
	// to extend a new triangle on the last edge of the previous one.
	TriangleStrip

	// TriangleFan indicates that the first three elements will be used
	// to construct a triangle and each consecutive element will be used
	// to extend on the previous triangle by forming a fan shape.
	TriangleFan

	// Triangles indicates that each three consecutive indices
	// will form a separate triangle.
	Triangles
)
