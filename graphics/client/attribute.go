package client

type Attribute string

type AttributeType int

const (
	FloatAttributeType AttributeType = iota
	Float2AttributeType
	Float3AttributeType
	Float4AttributeType
	Float2x2AttributeType
	Float3x3AttributeType
	Float4x4AttributeType
)

type AttributeLocation interface{}

type AttributeDeclaration struct {
	Id       Attribute
	Type     AttributeType
	Count    int
	Location AttributeLocation
}

//go:generate counterfeiter -o client_fakes/fake_attribute.go ./ AttributeClient

type AttributeClient interface {
	EnableAttribute(location AttributeLocation) error
	DisableAttribute(location AttributeLocation) error
	ConfigureAttribute(location AttributeLocation, components, strideInBytes, offsetInBytes int) error
}
