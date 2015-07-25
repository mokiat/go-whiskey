package graphics

//go:generate gostub Graphics

type Graphics interface {
	CreateFloat2AttributeArray(size int) Float2AttributeArray
	CreateFloat3AttributeArray(size int) Float3AttributeArray
}
