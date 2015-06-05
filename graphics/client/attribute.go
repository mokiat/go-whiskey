package client

type AttributeLocation interface{}

//go:generate counterfeiter -o client_fakes/fake_attribute.go ./ AttributeClient

type AttributeClient interface {
	EnableAttribute(location AttributeLocation) error
	DisableAttribute(location AttributeLocation) error
	ConfigureAttribute(location AttributeLocation, components, strideInBytes, offsetInBytes int) error
}
