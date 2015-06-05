package client

//go:generate counterfeiter -o client_fakes/fake_graphics.go ./ GraphicsClient

type GraphicsClient interface {
	ShaderClient
	BufferClient
	AttributeClient
	UniformClient
	ElementClient
}
