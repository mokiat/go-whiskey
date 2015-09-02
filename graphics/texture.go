package graphics

// TextureName is an enumeration the specifies the different possible
// binding targets for a texture.
type TextureName int

const (
	// DiffuseTexture specifies that the texture will be used as a diffuse
	// texture during rendering
	DiffuseTexture TextureName = iota

	// ReflectionTexture specifies that the texture will be used as a
	// reflection texture during rendering
	ReflectionTexture
)

// RGBAColor specifies a color made of four components - red, green, blue,
// and alpha (for opacity)
type RGBAColor struct {
	// R represents the red component of this color
	R byte

	// G represents the green component of this color
	G byte

	// B represents the blue component of this color
	B byte

	// A represents the alpha (opacity) component of this color
	A byte
}

//go:generate gostub Texture

// Texture represents a texture
type Texture interface {
}

//go:generate gostub FlatTexture

// FlatTexture represents a 2D texture
type FlatTexture interface {
	Texture

	// SetTexel configures the color in the texture at the
	// specified texel location
	SetTexel(x, y int, color RGBAColor)

	// Texel returns the color in the texture at the specified
	// texel location
	Texel(x, y int) RGBAColor
}

// CubeSide is an enumeration that specifies the possible sides of
// a CubeTexture
type CubeSide int

const (
	// NearCubeSide represents the surface of a CubeTexture that
	// is nearest the camera in an identity view
	NearCubeSide CubeSide = iota

	// FarCubeSide represents the surface of a CubeTexture that
	// is furthest the camera in an identity view
	FarCubeSide

	// LeftCubeSide represents the surface of a CubeTexture that
	// is to the left of the center of the camera in an identity view
	LeftCubeSide

	// RightCubeSide represents the surface of a CubeTexture that
	// is to the right of the center of the camera in an identity view
	RightCubeSide

	// TopCubeSide represents the surface of a CubeTexture that
	// is to the top of the center of the camera in an identity view
	TopCubeSide

	// BottomCubeSide represents the surface of a CubeTexture that
	// is to the bottom of the center of the camera in an identity view
	BottomCubeSide
)

//go:generate gostub CubeTexture

// FlatTexture represents a Cube texture that has six
// image surface sided
type CubeTexture interface {
	Texture

	// SetTexel configures the color in the texture at the
	// specified texel location and cube surface side
	SetTexel(side CubeSide, x, y int, color RGBAColor)

	// Texel returns the color in the texture at the specified
	// texel location and cube surface side
	Texel(side CubeSide, x, y int) RGBAColor
}
