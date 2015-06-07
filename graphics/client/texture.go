package client

type TextureId interface{}

//go:generate counterfeiter -o client_fakes/fake_texture.go ./ TextureClient

type TextureClient interface {
	Bind2DTexture(channel int, textureId TextureId) error
	BindCubeTexture(channel int, textureId TextureId) error
}
