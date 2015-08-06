package graphics

//go:generate gostub Material

type Material interface {
	EnableDiffuseColor(value bool)
	DiffuseColorEnabled() bool
}
