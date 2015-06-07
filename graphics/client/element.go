package client

//go:generate counterfeiter -o client_fakes/fake_element.go ./ ElementClient

type ElementClient interface {
	DrawTriangles(indexCount, indexOffsetInBytes int) error
	DrawLines(indexCount, indexOffsetInBytes int) error
}
