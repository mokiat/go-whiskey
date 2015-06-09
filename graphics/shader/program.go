package shader

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
)

//go:generate counterfeiter -o shader_fakes/fake_program.go ./ Program

type Program interface {
	VertexShader() Shader
	FragmentShader() Shader
	Id() client.ProgramId
	Created() bool
	Create(client.ShaderClient) error
	Delete(client.ShaderClient) error
}

func NewProgram(vertexShader, fragmentShader Shader) Program {
	return &program{
		vertexShader:   vertexShader,
		fragmentShader: fragmentShader,
	}
}

type program struct {
	id             client.ProgramId
	vertexShader   Shader
	fragmentShader Shader
}

func (p *program) VertexShader() Shader {
	return p.vertexShader
}

func (p *program) FragmentShader() Shader {
	return p.fragmentShader
}

func (p *program) Id() client.ProgramId {
	return p.id
}

func (p *program) Created() bool {
	return p.id != nil
}

func (p *program) Create(shaderClient client.ShaderClient) error {
	if !p.vertexShader.Created() {
		return errors.New("Vertex shader is not created!")
	}
	if !p.fragmentShader.Created() {
		return errors.New("Fragment shader is not created!")
	}
	var err error
	p.id, err = shaderClient.CreateProgram()
	if err != nil {
		return err
	}
	err = shaderClient.AttachShaderToProgram(p.vertexShader.Id(), p.id)
	if err != nil {
		return err
	}
	err = shaderClient.AttachShaderToProgram(p.fragmentShader.Id(), p.id)
	if err != nil {
		return err
	}
	err = shaderClient.LinkProgram(p.id)
	if err != nil {
		return err
	}
	return nil
}

func (p *program) Delete(shaderClient client.ShaderClient) error {
	err := shaderClient.DeleteProgram(p.id)
	if err != nil {
		return err
	}
	p.id = nil
	return nil
}
