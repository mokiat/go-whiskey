package uniform

import (
	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/math"
)

//go:generate counterfeiter -o uniform_fakes/fake_uniform_value.go ./ UniformValue

type UniformValue interface {
	Bind(uniformClient client.UniformClient, location client.UniformLocation)
}

type Vec4UniformValue struct {
	Value math.Vec4
}

func (v Vec4UniformValue) Bind(uniformClient client.UniformClient, location client.UniformLocation) {
	uniformClient.BindVec4Uniform(location, v.Value)
}

type Mat4x4UniformValue struct {
	Value math.Mat4x4
}

func (v Mat4x4UniformValue) Bind(uniformClient client.UniformClient, location client.UniformLocation) {
	uniformClient.BindMat4x4Uniform(location, v.Value)
}

//go:generate counterfeiter -o uniform_fakes/fake_texture_uniform_value.go ./ TextureUniformValue

type TextureUniformValue interface {
	Bind(uniformClient client.UniformClient, channel int, location client.UniformLocation)
}
