package graphics

import (
	"errors"
	"fmt"

	. "github.com/momchil-atanasov/go-whiskey/graphics"
	"golang.org/x/mobile/gl"
)

type facade struct{}

func NewFacade() Facade {
	return &facade{}
}

func (f *facade) CreateBuffer() ResourceId {
	return 0
}

func (f *facade) BindIndexBuffer(bufferId ResourceId) {

}

func (f *facade) CreateIndexBufferData(data []byte, usage BufferUsage) {

}

func (f *facade) BindVertexBuffer(bufferId ResourceId) {

}

func (f *facade) CreateVertexBufferData(data []byte, usage BufferUsage) {

}

func (f *facade) DeleteBuffer(bufferId ResourceId) {

}

func (f *facade) CreateVertexShader() (ResourceId, error) {
	shaderId := ResourceId(gl.CreateShader(gl.VERTEX_SHADER).Value)
	if shaderId == InvalidShaderId {
		return InvalidShaderId, errors.New("Could not allocate vertex shader!")
	}
	return shaderId, nil
}

func (f *facade) CreateFragmentShader() (ResourceId, error) {
	shaderId := ResourceId(gl.CreateShader(gl.FRAGMENT_SHADER).Value)
	if shaderId == InvalidShaderId {
		return InvalidShaderId, errors.New("Could not allocate fragment shader!")
	}
	return shaderId, nil
}

func (f *facade) SetShaderSourceCode(shaderId ResourceId, sourceCode string) {
	shader := gl.Shader{
		Value: uint32(shaderId),
	}
	gl.ShaderSource(shader, sourceCode)
}

func (f *facade) CompileShader(shaderId ResourceId) error {
	shader := gl.Shader{
		Value: uint32(shaderId),
	}
	gl.CompileShader(shader)
	if !f.isShaderCompilationSuccessful(shader) {
		info := gl.GetShaderInfoLog(shader)
		return errors.New(fmt.Sprintf("Could not compile shader due to '%s'", info))
	}
	return nil
}

func (f *facade) isShaderCompilationSuccessful(shader gl.Shader) bool {
	flag := gl.GetShaderi(shader, gl.COMPILE_STATUS)
	return flag != gl.FALSE
}

func (f *facade) DeleteShader(shaderId ResourceId) {
	shader := gl.Shader{
		Value: uint32(shaderId),
	}
	gl.DeleteShader(shader)
}
