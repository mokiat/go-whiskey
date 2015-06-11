package client

type Uniform string

type UniformType int

func (t UniformType) IsSampler() bool {
	return (t == Sampler2DUniformType) || (t == SamplerCubeUniformType)
}

const (
	FloatUniformType UniformType = iota
	Float2UniformType
	Float3UniformType
	Float4UniformType
	IntUniformType
	Int2UniformType
	Int3UniformType
	Int4UniformType
	BoolUniformType
	Bool2UniformType
	Bool3UniformType
	Bool4UniformType
	Float2x2UniformType
	Float3x3UniformType
	Float4x4UniformType
	Sampler2DUniformType
	SamplerCubeUniformType
)

type UniformLocation interface{}

type UniformDeclaration struct {
	Id       Uniform
	Type     UniformType
	Count    int
	Location UniformLocation
}

//go:generate counterfeiter -o client_fakes/fake_uniform.go ./ UniformClient

type UniformClient interface {
	SetVec4Uniform(UniformLocation, []float32) error
	SetMat4x4Uniform(UniformLocation, []float32) error
	SetSamplerUniform(UniformLocation, int) error
}
