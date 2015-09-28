package graphics

//go:generate gostub Layer

// Layer interface represents a render layer.
// One would usually use such layers to group render objects
// into consecutive collections that need to be processed.
// TODO: In the future allow also render targets (buffers, textures)
// to be set to the layer. This would allow for shadow mapping for example.
type Layer interface {
}
