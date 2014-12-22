package uniform

type Uniform int

const (
	ProjectionMatrix Uniform = iota
	ModelViewMatrix
	ModelMatrix
	ViewMatrix
	AmbientColor
	DiffuseColor
	DiffuseTexture
	UNIFORM_COUNT int = iota
)

var names = make([]string, UNIFORM_COUNT)

func init() {
	names[ProjectionMatrix] = "projectionMatrix"
	names[ModelViewMatrix] = "modelViewMatrix"
	names[ModelMatrix] = "modelMatrix"
	names[ViewMatrix] = "viewMatrix"
	names[AmbientColor] = "ambientColor"
	names[DiffuseColor] = "diffuseColor"
	names[DiffuseTexture] = "diffuseTexture"
}

func (u Uniform) Name() string {
	return names[u]
}
