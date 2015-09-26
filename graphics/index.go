package graphics

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
