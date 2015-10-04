package graphics

import "fmt"

// LayerSize describes the amount of render items
// that a layer can handle.
type LayerSize int

const (
	// TinyLayerSize describes layers that can handle
	// roughly 16 render items.
	TinyLayerSize LayerSize = iota

	// SmallLayerSize describes layers that can handle
	// roughly 128 render items.
	SmallLayerSize

	// MediumLayerSize describes layers that can handle
	// roughly 512 render items.
	MediumLayerSize

	// LargeLayerSize describes layers that can handle
	// roughly 8,192 render items.
	LargeLayerSize

	// HugeLayerSize describes layers that can handle
	// roughly 131,072 render items.
	HugeLayerSize
)

func (s LayerSize) Capacity() int {
	switch s {
	case TinyLayerSize:
		return 16
	case SmallLayerSize:
		return 128
	case MediumLayerSize:
		return 512
	case LargeLayerSize:
		return 8192
	case HugeLayerSize:
		return 131072
	default:
		panic(fmt.Sprintf("Unknown layer size '%d'.", s))
	}
}

//go:generate gostub Layer

// Layer represents a surface onto which items
// can be rendered.
// Layers can be used to perform multiple render passes of a single
// scene (e.g. shadow mapping) or to use is as an alternative to
// depth testing (e.g. 2D sprite game).
type Layer interface {

	// SetClearColorEnabled determines whether this layer will
	// have it's color buffer cleared.
	SetClearColorEnabled(bool)

	// SetClearColor configures the color to be used when clearing
	// the target surface.
	// This option is only meaningful if color clearing is enabled
	// via SetClearColorEnabled
	SetClearColor(r, g, b, a float32)

	// SetClearDepthEnabled determines whether this layer will
	// have it's depth buffer cleared.
	SetClearDepthEnabled(bool)

	// ItemBuilder returns the ItemBuilder for this layer.
	ItemBuilder() ItemBuilder
}

//go:generate gostub ItemBuilder

// ItemBuilder provides the means necessary to queue items to be
// rendered on given render layer.
type ItemBuilder interface {

	// Begin method initiates the construction of a new item to be
	// queued for rendering on this layer.
	// The provided method and filter are used to configure the
	// mechanism through which the item will be rendered.
	// Other methods on the ItemBuilder can be used to further
	// configure the item.
	// Once the item has been configured, it should be queued for
	// rendering via the End method.
	// One may call Begin again only after having called End beforehand.
	Begin(Material, MaterialFilter) ItemBuilder

	// SetDepth configured the depth of the current render item.
	// The depth can be used to sort the order in which items
	// will be rendered on this layer.
	// By default items are rendered with a depth configuration of 0.
	// Usage of this setting could lead to performance issues since
	// it might prevent or hinder any internal render optimizations.
	SetDepth(depth int) ItemBuilder

	// SetAttribute configures an attribute to be used for the current
	// render item.
	// This method may be called multiple times to configure additional
	// attributes.
	SetAttribute(AttributeName, AttributeArray) ItemBuilder

	// SetUniform configures a uniform to be used for the current
	// render item.
	// This method may be called multiple times to configure additional
	// uniforms.
	SetUniform(UniformName, Uniform) ItemBuilder

	// SetTexture configures a texture to be used for the current
	// render item.
	// This method may be called multiple times to configure additional
	// textures.
	SetTexture(TextureName, Texture) ItemBuilder

	// SetIndices configures indices to be used for the description
	// of the shape of the render item.
	// If this method is not called on the current item then indices
	// will not be used and instead vertices will be iterated in order
	// from the configured attributes.
	SetIndices(IndexArray) ItemBuilder

	// End queues the render item to be rendered by this queue.
	// This method is also used to specify the Shape that will
	// rendered via the vertices in this item and the amount of
	// vertices that will be required to render it as well as
	// any offset in the vertices / index data.
	// This method will return an error if there is a problem
	// with any of the data that was configured in this ItemBuilder
	// for the current item.
	// If this method does not return an error, it is not guaranteed
	// that rendering will succeed and an error will not be returned
	// in any rendering calls.
	End(sequence SequenceType, offset, count int) error
}
