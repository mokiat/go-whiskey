package graphics

// FeatureMask is a flag mask that indicates a requirement
// or a usage of a given render feature depending on where
// it is used.
type FeatureMask uint64

const (
	// NoFeatures represents an empty FeatureMask
	NoFeatures FeatureMask = 1 << iota

	// MeshSkinning indicates that the data necessary for
	// mesh joint transformation will be provided or is required
	// depending on the usage of this flag.
	MeshSkinning

	// DiffuseMapping indicates that the data necessary for
	// mapping a diffuse texture on top of a mesh will be provided
	// or is required depending on the usage of this flag.
	DiffuseMapping

	// DiffuseColoring indicates that the data necessary for
	// coloring a mesh with a diffuse color will be provided or
	// is required depending on the usage of this flag.
	DiffuseColoring

	// PointLight indicates that the data necessary for
	// a point light source will be provided or is required
	// depending on the usage of this flag.
	PointLight

	// SpotLight indicates that the data necessary for
	// a spot light source will be provided or is required
	// depending on the usage of this flag.
	SpotLight

	// DirectionalLight indicates that the data necessary for
	// a directional light source will be provided or is required
	// depending on the usage of this flag.
	DirectionalLight

	// ShadowMapping indicates that the data necessary for
	// shadow mapping will be provided or is required
	// depending on the usage of this flag.
	ShadowMapping

	// ReflectionEnvrionment indicates that the data necessary for
	// envrionment map reflection will be provided or is required
	// depending on the usage of this flag.
	ReflectionEnvironment
)

// Filter is a mechanism through which users can ask the render
// pipeline to select the best tender technique based on the
// current status of the scene.
type Filter struct {
	LOD      float32
	Features FeatureMask
}

// Technique is a mechanism through which a single material can
// be rendered differently depending on the environment in which
// it is located.
type Technique struct {
	MinLOD       float32
	MaxLOD       float32
	Requirements FeatureMask
	Mechanism    Mechanism
}

//go:generate gostub Mechanism

// Mechanism is an abstraction to the way through which a material
// will be relized. This could be fixed-function or some type of
// shader.
type Mechanism interface {
}

// PhongMechanism is an implementation of Mechanism that indicates
// a phong shader approach to rendering.
type PhongMechanism struct {
}

// Definition is used to specify the type of material to be created.
type Definition struct {

	// Techniques lists all techniques that can be used for this material
	Techniques []*Technique
}

//go:generate gostub Material

// Material is used to indicate the way a given shape or mesh
// should be rendered.
type Material interface {
}
