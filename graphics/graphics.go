package graphics

//go:generate gostub Graphics

// Graphics is an interface to the core graphics capabilities
// of the engine.
type Graphics interface {

	// CreateFloat2AttributeArray creates a new Float2AttributeArray
	// of the specified size (number of entries)
	CreateFloat2AttributeArray(size int) (Float2AttributeArray, error)

	// CreateFloat3AttributeArray creates a new Float3AttributeArray
	// of the specified size (number of entries)
	CreateFloat3AttributeArray(size int) (Float3AttributeArray, error)

	// DeleteAttributeArray deletes the specified AttributeArray.
	// Users should no longer keep a hold of the AttributeArray as
	// it might get reused.
	DeleteAttributeArray(AttributeArray) error

	// CreateIndexArray creates a new IndexArray of the specified
	// size (number of indices)
	CreateIndexArray(size int) (IndexArray, error)

	// DeleteIndexArray deletes the specified IndexArray.
	// Users should no longer keep a hold of the IndexArray as
	// it might get reused.
	DeleteIndexArray(IndexArray) error

	// CreateFloat4Uniform creates a new Float4Uniform
	CreateFloat4Uniform() (Float4Uniform, error)

	// CreateFloat4x4Uniform creates a new Float4x4Uniform
	CreateFloat4x4Uniform() (Float4x4Uniform, error)

	// DeleteUniform deletes the specified Uniform
	// Users should no longer keep a hold of the Uniform as
	// it might get reused.
	DeleteUniform(Uniform) error

	// CreateFlatTexture creates a new FlatTexture with the
	// specified 2D dimensions.
	CreateFlatTexture(width int, height int) (FlatTexture, error)

	// CreateCubeTexture creates a new CubeTexture with the
	// specified dimension for each square side.
	CreateCubeTexture(dimension int) (CubeTexture, error)

	// DeleteTexture deletes the specified texture.
	// Users should no longer keep a hold of the Texture as
	// it might get reused.
	DeleteTexture(Texture) error

	// CreateMaterial creates a new Material based off of the
	// specified MaterialDefinition.
	// Keep in mind that this is a costly operation.
	CreateMaterial(Definition) (Material, error)

	// DeleteMaterial deletes the specified Material.
	// Users should no longer keep a hold of the Material as
	// it might get reused.
	DeleteMaterial(Material) error

	// Invalidate notifies the graphics engine that the
	// graphics context was lost and that the engine needs to prepare to
	// initialize all resources during the next Initialize call
	Invalidate() error

	// Initialize should be called to initialize all the resources
	// of the graphics engine. This should be called once after the
	// graphics context has been created.
	// Some Graphics implementations might require you to call this method
	// from a specific thread.
	Initialize() error

	// Destroy should be called when the graphics engine can freely
	// delete all resources. Usually this would be called prior to
	// shutting down an application.
	// Some Graphics implementations might require you to call this method
	// from a specific thread.
	Destroy() error

	// UseMaterial indicates to the Rendering pipeline that the specified
	// material should be used for future render operations.
	UseMaterial(Material) error

	// BindAttribute binds the specified attribute array to the specified
	// attribute bind name
	BindAttribute(AttributeName, AttributeArray) error

	// BindUniform binds the specified uniform to the specified uniform
	// bind name
	BindUniform(UniformName, Uniform) error

	// BindTexture binds the specified texture to the specified texture
	// bind name
	BindTexture(TextureName, Texture) error

	// Render schedule a shape to be rendered. The actual rendering will
	// take place during the next call to Flush.
	// The IndexArray parameter can be nil in which case the vertices
	// will be consumed from the attribute arrays directly in order.
	// The offset and count parameters specify and offset relative to
	// the start of the IndexArray and a number of indices to be used
	// for the rendering of the shape.
	Render(sequence SequenceType, array IndexArray, offset, count int) error

	// Flush triggers the rendering pipeline to draw all of the scheduled
	// shapes via the Render call to be rendered to the screen.
	// Some Graphics implementations might require you to call this method
	// from a specific thread.
	Flush() error
}
