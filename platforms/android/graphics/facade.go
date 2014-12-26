package graphics

import (
	"errors"
	"fmt"

	"github.com/momchil-atanasov/go-whiskey/common/buf"
	. "github.com/momchil-atanasov/go-whiskey/graphics"
	"golang.org/x/mobile/gl"
)

type facade struct {
}

func NewFacade() Facade {
	return &facade{}
}

func (f *facade) CreateBuffer() (ResourceId, error) {
	bufferId := ResourceId(gl.GenBuffer().Value)
	if bufferId == InvalidBufferId {
		return InvalidBufferId, errors.New("Could not allocate buffer!")
	}
	return bufferId, nil
}

func (f *facade) BindIndexBuffer(bufferId ResourceId) {
	buffer := gl.Buffer{
		Value: uint32(bufferId),
	}
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffer)
}

func (f *facade) CreateIndexBufferData(data buf.UInt16Buffer, usage BufferUsage) {
	usageFlag := f.getBufferUsageFlag(usage)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, usageFlag, data)
}

func (f *facade) BindVertexBuffer(bufferId ResourceId) {
	buffer := gl.Buffer{
		Value: uint32(bufferId),
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
}

func (f *facade) CreateVertexBufferData(data buf.Float32Buffer, usage BufferUsage) {
	usageFlag := f.getBufferUsageFlag(usage)
	gl.BufferData(gl.ARRAY_BUFFER, usageFlag, data)
}

func (f *facade) DeleteBuffer(bufferId ResourceId) {
	gl.DeleteBuffer(gl.Buffer{
		Value: uint32(bufferId),
	})
}

func (f *facade) getBufferUsageFlag(usage BufferUsage) gl.Enum {
	switch usage {
	case StreamDraw:
		return gl.STREAM_DRAW
	case StaticDraw:
		return gl.STATIC_DRAW
	case DynamicDraw:
		return gl.DYNAMIC_DRAW
	default:
		panic(fmt.Sprintf("Unknown buffer usage %v", usage))
	}
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
		return errors.New(fmt.Sprintf("Could not compile shader due to '%s'!", info))
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

func (f *facade) CreateProgram() (ResourceId, error) {
	programId := ResourceId(gl.CreateProgram().Value)
	if programId == InvalidProgramId {
		return InvalidProgramId, errors.New("Could not allocate shader program!")
	}
	return programId, nil
}

func (f *facade) AttachShaderToProgram(programId ResourceId, shaderId ResourceId) {
	program := gl.Program{
		Value: uint32(programId),
	}
	shader := gl.Shader{
		Value: uint32(shaderId),
	}
	gl.AttachShader(program, shader)
}

func (f *facade) LinkProgram(programId ResourceId) error {
	program := gl.Program{
		Value: uint32(programId),
	}
	gl.LinkProgram(program)
	if !f.isProgramLinkSuccessful(program) {
		info := gl.GetProgramInfoLog(program)
		return errors.New(fmt.Sprintf("Failed to link program due to '%s'!", info))
	}
	return nil
}

func (f *facade) isProgramLinkSuccessful(program gl.Program) bool {
	flag := gl.GetProgrami(program, gl.LINK_STATUS)
	return flag != gl.FALSE
}

func (f *facade) UseProgram(programId ResourceId) {
	gl.UseProgram(gl.Program{
		Value: uint32(programId),
	})
}

func (f *facade) DeleteProgram(programId ResourceId) {
	gl.DeleteProgram(gl.Program{
		Value: uint32(programId),
	})
}
