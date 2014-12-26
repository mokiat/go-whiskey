package shader

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics"
)

type ShaderProgram interface {
	Id() graphics.ResourceId
	VertexShader() VertexShader
	FragmentShader() FragmentShader
	BindRemotely()
	CreateRemotely() error
	DeleteRemotely()
	CreatedRemotely() bool
}

type shaderProgram struct {
	id             graphics.ResourceId
	facade         graphics.Facade
	vertexShader   VertexShader
	fragmentShader FragmentShader
}

func NewShaderProgram(facade graphics.Facade, vertexShader VertexShader, fragmentShader FragmentShader) ShaderProgram {
	return &shaderProgram{
		id:             graphics.InvalidProgramId,
		facade:         facade,
		vertexShader:   vertexShader,
		fragmentShader: fragmentShader,
	}
}

func (p *shaderProgram) Id() graphics.ResourceId {
	return p.id
}

func (p *shaderProgram) VertexShader() VertexShader {
	return p.vertexShader
}

func (p *shaderProgram) FragmentShader() FragmentShader {
	return p.fragmentShader
}

func (p *shaderProgram) BindRemotely() {
	p.facade.UseProgram(p.id)
}

func (p *shaderProgram) CreateRemotely() error {
	if !p.vertexShader.CreatedRemotely() {
		return errors.New("Vertex shader is not initialized!")
	}
	if !p.fragmentShader.CreatedRemotely() {
		return errors.New("Fragment shader is not initialized!")
	}
	var err error
	p.id, err = p.facade.CreateProgram()
	if err != nil {
		p.id = graphics.InvalidProgramId
		return err
	}
	p.facade.AttachShaderToProgram(p.id, p.vertexShader.Id())
	p.facade.AttachShaderToProgram(p.id, p.fragmentShader.Id())
	err = p.facade.LinkProgram(p.id)
	if err != nil {
		return err
	}
	return nil
}

func (p *shaderProgram) DeleteRemotely() {
	p.facade.DeleteProgram(p.id)
}

func (p *shaderProgram) CreatedRemotely() bool {
	return p.id != graphics.InvalidProgramId
}
