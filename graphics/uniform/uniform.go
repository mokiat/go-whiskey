package uniform

import "github.com/momchil-atanasov/go-whiskey/common"

type Uniform common.Enum

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
	names[ProjectionMatrix] = "projectionMatrixIn"
	names[ModelViewMatrix] = "modelViewMatrixIn"
	names[ModelMatrix] = "modelMatrixIn"
	names[ViewMatrix] = "viewMatrixIn"
	names[AmbientColor] = "ambientColorIn"
	names[DiffuseColor] = "diffuseColorIn"
	names[DiffuseTexture] = "diffuseTextureIn"
}

func (u Uniform) Name() string {
	return names[u]
}
