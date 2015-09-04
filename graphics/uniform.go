package graphics

import "github.com/momchil-atanasov/go-whiskey/math"

// UniformName is an enumeration the specifies the different possible
// binding targets for a uniform.
type UniformName int

const (
	// ProjectionMatrix specifies that the uniform will be used
	// as a projection matrix
	ProjectionMatrix UniformName = iota

	// ModelViewMatrix specifies that the uniform will be used
	// as a modelview matrix
	ModelViewMatrix

	// ModelMatrix specifies that the uniform will be used
	// as a model matrix
	ModelMatrix

	// ViewMatrix specifies that the uniform will be used
	// as a view matrix
	ViewMatrix

	// AmbientColor specifies that the uniform will be used
	// as an ambient color
	AmbientColor

	// DiffuseColor specifies that the uniform will be used
	// as a diffuse color
	DiffuseColor

	// SpecularColor specifies that the uniform will be used
	// as a specular color
	SpecularColor

	// LightPosition specifies the position of a light source
	// in world space
	LightPosition

	// LightDirection specifies the direction of the light source
	// in world space
	LightDirection

	// LightDiffuseColor specifies the diffuse color of the light
	// source
	LightDiffuseColor
)

//go:generate gostub Uniform

// Uniform represents a configuration that is global
// for a whole mesh structure
type Uniform interface {
}

//go:generate gostub Float3Uniform

// Float3Uniform represents a uniform that has three
// float components
type Float3Uniform interface {
	Uniform

	// SetValue configures this uniform to the specified
	// vector of three float values
	SetValue(math.Vec3)

	// Value returns the current vector that is set
	// to this uniform
	Value() math.Vec3
}

//go:generate gostub Float4Uniform

// Float4Uniform represents a uniform that has four
// float components
type Float4Uniform interface {
	Uniform

	// SetValue configures this uniform to the specified
	// vector of four float values
	SetValue(math.Vec4)

	// Value returns the current vector that is set
	// to this uniform
	Value() math.Vec4
}

//go:generate gostub Float4x4Uniform

// Float4x4Uniform represents a uniform that is a
// matrix of four by four dimension
type Float4x4Uniform interface {
	Uniform

	// SetValue configures this uniform to the specified
	// matrix of four by four float values
	SetValue(math.Mat4x4)

	// Value returns the current matrix that is set to
	// this uniform
	Value() math.Mat4x4
}
