package graphics

type MaterialDefinition struct {
	HasDiffuseColor bool
}

//go:generate gostub Material

type Material interface {
}
