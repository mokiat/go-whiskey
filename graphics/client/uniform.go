package client

type UniformLocation interface{}

//go:generate counterfeiter -o client_fakes/fake_uniform.go ./ UniformClient

type UniformClient interface {
	SetVec4Uniform(UniformLocation, []float32) error
	SetMat4x4Uniform(UniformLocation, []float32) error
	SetSamplerUniform(UniformLocation, int) error
}
