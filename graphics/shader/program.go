package shader

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
)

//go:generate counterfeiter -o shader_fakes/fake_program.go ./ Program

type Program interface {
	VertexShader() RemoteShader
	FragmentShader() RemoteShader
	Remote() RemoteProgram
}

//go:generate counterfeiter -o shader_fakes/fake_remote_program.go ./ RemoteProgram

type RemoteProgram interface {
	Id() client.ProgramId
	Created() bool
	Create() error
	Delete() error
	Use() error
}

func NewProgram(shaderClient client.ShaderClient, vertexShader, fragmentShader RemoteShader) Program {
	return &program{
		shaderClient:   shaderClient,
		vertexShader:   vertexShader,
		fragmentShader: fragmentShader,
	}
}

type program struct {
	shaderClient   client.ShaderClient
	vertexShader   RemoteShader
	fragmentShader RemoteShader
	id             client.ProgramId
}

func (p *program) VertexShader() RemoteShader {
	return p.vertexShader
}

func (p *program) FragmentShader() RemoteShader {
	return p.fragmentShader
}

func (p *program) Remote() RemoteProgram {
	return p
}

func (p *program) Id() client.ProgramId {
	return p.id
}

func (p *program) Created() bool {
	return p.id != nil
}

func (p *program) Create() error {
	if !p.vertexShader.Created() {
		return errors.New("Vertex shader is not created!")
	}
	if !p.fragmentShader.Created() {
		return errors.New("Fragment shader is not created!")
	}
	var err error
	p.id, err = p.shaderClient.CreateProgram()
	if err != nil {
		return err
	}
	err = p.shaderClient.AttachShaderToProgram(p.vertexShader.Id(), p.id)
	if err != nil {
		return err
	}
	err = p.shaderClient.AttachShaderToProgram(p.fragmentShader.Id(), p.id)
	if err != nil {
		return err
	}
	err = p.shaderClient.LinkProgram(p.id)
	if err != nil {
		return err
	}
	return nil
}

func (p *program) Use() error {
	return p.shaderClient.UseProgram(p.id)
}

func (p *program) Delete() error {
	err := p.shaderClient.DeleteProgram(p.id)
	if err != nil {
		return err
	}
	p.id = nil
	return nil
}
