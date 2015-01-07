package uniform

import (
	"github.com/momchil-atanasov/go-whiskey/graphics"
	"github.com/momchil-atanasov/go-whiskey/math"
)

type UniformValue interface {
	Bind(facade graphics.Facade, location graphics.BindLocation)
}

type TextureUniformValue interface {
	Bind(facade graphics.Facade, channel int, location graphics.BindLocation)
}

type Vec4UniformValue struct {
	Value math.Vec4
}

func (v Vec4UniformValue) Bind(facade graphics.Facade, location graphics.BindLocation) {
	facade.BindVec4Uniform(v.Value, location)
}
