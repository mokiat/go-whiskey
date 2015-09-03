package graphics

// MaterialFeatureMask is a flag mask that indicates a requirement
// or a usage of a given render feature depending on where
// it is used.
type MaterialFeatureMask uint64

const (
	// NoFeatures represents an empty MaterialFeatureMask
	NoFeatures MaterialFeatureMask = 1 << iota

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

// MaterialFilter is a mechanism through which users can ask the render
// pipeline to select the best tender technique based on the
// current status of the scene.
type MaterialFilter struct {
	LOD      float32
	Features MaterialFeatureMask
}

// MaterialTechnique is a mechanism through which a single material can
// be rendered differently depending on the environment in which
// it is located.
type MaterialTechnique struct {
	MinLOD       float32
	MaxLOD       float32
	Requirements MaterialFeatureMask
	Mechanism    MaterialMechanism
}

//go:generate gostub MaterialMechanism

// MaterialMechanism is an abstraction to the way through which a material
// will be relized. This could be fixed-function or some type of
// shader.
type MaterialMechanism interface {
}

// PhongMaterialMechanism is an implementation of MaterialMechanism
// that indicates a phong shader approach to rendering.
type PhongMaterialMechanism struct {
}

// MaterialDefinition is used to specify the type of material to be created.
type MaterialDefinition struct {

	// Techniques lists all techniques that can be used for this material
	Techniques []*MaterialTechnique
}

//go:generate gostub Material

// Material is used to indicate the way a given shape or mesh
// should be rendered.
type Material interface {
}
