package graphics

// AttributeName is an enumeration the specifies the different possible
// binding targets for an attribute.
type AttributeName int

const (
	// Coords specifies that the attribute array will be used
	// to specify the coordinates of the vertices
	Coords AttributeName = iota

	// TexCoords specifies that the attribute array will be used
	// to specify the texture coordinates of the vertices
	TexCoords

	// Normals specifies that the attribute array will be used
	// to specify the normals of the vertices
	Normals
)

//go:generate gostub AttributeArray

// AttributeArray represents a single feature (e.g. position, normal,
// texture coordinate) of a mesh in the form of a continuous array of entries.
type AttributeArray interface {

	// Size returns the number of entries that
	// can be stored in this attribute array
	Size() int
}

//go:generate gostub Float2AttributeArray

// Float2AttributeArray represents an AttributeArray which stores
// two-dimensional entries of float type.
type Float2AttributeArray interface {
	AttributeArray

	// PutFloat2 places a 2D float vector specified by
	// the x, y parameters at the position specified
	// by position
	PutFloat2(position int, x, y float32)

	// Float2 returns the 2D float vector positioned at
	// the specified position.
	Float2(position int) (x, y float32)
}

//go:generate gostub Float3AttributeArray

// Float3AttributeArray represents an AttributeArray which stores
// three-dimensional entries of float type.
type Float3AttributeArray interface {
	AttributeArray

	// PutFloat3 places a 3D float vector specified by
	// the x, y, z parameters at the position specified
	// by position
	PutFloat3(position int, x, y, z float32)

	// Float3 returns the 3D float vector positioned at
	// the specified position.
	Float3(position int) (x, y, z float32)
}
