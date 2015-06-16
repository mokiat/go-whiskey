package shader

import (
	"errors"
	"sort"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
)

//go:generate counterfeiter -o shader_fakes/fake_program.go ./ Program

type Program interface {
	VertexShader() Shader
	FragmentShader() Shader
	Id() client.ProgramId
	UniformDeclarations() []client.UniformDeclaration
	UniformDeclaration(client.Uniform) (client.UniformDeclaration, bool)
	AttributeDeclarations() []client.AttributeDeclaration
	AttributeDeclaration(client.Attribute) (client.AttributeDeclaration, bool)
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
	id                    client.ProgramId
	vertexShader          Shader
	fragmentShader        Shader
	uniformDeclarations   []client.UniformDeclaration
	attributeDeclarations []client.AttributeDeclaration
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

func (p *program) UniformDeclarations() []client.UniformDeclaration {
	return p.uniformDeclarations
}

func (p *program) UniformDeclaration(uniform client.Uniform) (client.UniformDeclaration, bool) {
	l := 0
	r := len(p.uniformDeclarations) - 1
	for l <= r {
		m := (l + r) / 2
		if p.uniformDeclarations[m].Id == uniform {
			return p.uniformDeclarations[m], true
		}
		if p.uniformDeclarations[m].Id < uniform {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return client.UniformDeclaration{}, false
}

func (p *program) AttributeDeclarations() []client.AttributeDeclaration {
	return p.attributeDeclarations
}

func (p *program) AttributeDeclaration(attribute client.Attribute) (client.AttributeDeclaration, bool) {
	l := 0
	r := len(p.attributeDeclarations) - 1
	for l <= r {
		m := (l + r) / 2
		if p.attributeDeclarations[m].Id == attribute {
			return p.attributeDeclarations[m], true
		}
		if p.attributeDeclarations[m].Id < attribute {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return client.AttributeDeclaration{}, false
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
	p.uniformDeclarations, err = shaderClient.GetProgramUniforms(p.id)
	sort.Sort(uniformList(p.uniformDeclarations))
	if err != nil {
		return err
	}
	p.attributeDeclarations, err = shaderClient.GetProgramAttributes(p.id)
	sort.Sort(attributeList(p.attributeDeclarations))
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

type uniformList []client.UniformDeclaration

func (l uniformList) Len() int {
	return len(l)
}

func (l uniformList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l uniformList) Less(i, j int) bool {
	return l[i].Id < l[j].Id
}

type attributeList []client.AttributeDeclaration

func (l attributeList) Len() int {
	return len(l)
}

func (l attributeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l attributeList) Less(i, j int) bool {
	return l[i].Id < l[j].Id
}
