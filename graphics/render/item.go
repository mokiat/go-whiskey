package render

import "github.com/momchil-atanasov/go-whiskey/graphics/client"

// Item holds the minimum amount of information needed to render
// a mesh on the screen.
//
// A scene would generally hold higher level objects that are subdivided
// into multiple Item objects based on Material and Shape.
type Item struct {

	// ProgramId holds the Id of the shader program
	// to be used for rendering
	ProgramId client.ProgramId

	// Attributes array holds the configuration for all attributes
	// that will be used by the shader program.
	Attributes []AttributeEntry

	// Uniforms array holds the configuration for all uniforms (except
	// textures) that will be used by the shader program.
	Uniforms []UniformEntry

	// Textures array holds the configuration for all textures that will
	// be used by the shader program.
	Textures []TextureEntry

	// Elements array describes all the shapes to be rendered
	Elements []ElementEntry
}

// AttributeEntry describes a single attribute binding
type AttributeEntry struct {

	// Location specifies the shader location of the attribute
	Location client.AttributeLocation

	// VertexBufferId specifies the Id of the vertex buffer that holds
	// this attribute's data
	VertexBufferId client.BufferId

	// Components specifies the number of components that make up an
	// attribute.
	Components int

	// StrideInBytes specifies the offset between each two consecutive attributes.
	StrideInBytes int

	// OffsetInBytes specifies the offset from the start of the buffer where
	// this attribute begins.
	OffsetInBytes int
}

// UniformType describes the type of the uniform.
type UniformType int

const (
	Vec4UniformType UniformType = iota
	Mat4x4UniformType
)

// UniformEntry describes a single uniform binding
type UniformEntry struct {

	// Location specifies the shader location of the uniform
	Location client.UniformLocation

	// Type specifies the type of the uniform
	Type UniformType

	// Values holds the float values that make up the uniform
	Values []float32
}

// TextureType describes the type of the texture.
type TextureType int

const (
	TwoDTextureType TextureType = iota
	CubeTextureType
)

// TextureEntry describes a single texture binding
type TextureEntry struct {

	// Location specifies the shader location of the uniform that
	// consumes a texture
	Location client.UniformLocation

	// Type specifies the type of texture to be bound
	Type TextureType

	// TextureId specifies the Id of the texture to be bound
	TextureId client.TextureId
}

// ElementType describes the type of shape to be rendered.
type ElementType int

const (
	TrianglesElementType ElementType = iota
	LinesElementType
)

// ElementEntry describes the mesh to be rendered.
type ElementEntry struct {

	// IndexBufferId specifies the Id of the index buffer to be used
	// for the rendering of the mesh
	IndexBufferId client.BufferId

	// Type specifies the shapes to be used for distinct elements.
	Type ElementType

	// IndexCount specifies the number of indices that make up the whole mesh.
	IndexCount int

	// IndexOffsetInBytes specifies the offset from the beginning of the
	// index buffer where the index data is located.
	IndexOffsetInBytes int
}
