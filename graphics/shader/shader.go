package shader

import "github.com/momchil-atanasov/go-whiskey/graphics/client"

//go:generate counterfeiter -o shader_fakes/fake_shader_data.go ./ ShaderData

type ShaderData interface {
	SetSourceCode(string)
	SourceCode() string
}

func NewShaderData() ShaderData {
	return &shaderData{}
}

type shaderData struct {
	sourceCode string
}

func (d *shaderData) SetSourceCode(sourceCode string) {
	d.sourceCode = sourceCode
}

func (d *shaderData) SourceCode() string {
	return d.sourceCode
}

//go:generate counterfeiter -o shader_fakes/fake_shader.go ./ Shader

type Shader interface {
	Data() ShaderData
	Id() client.ShaderId
	Created() bool
	Create(client.ShaderClient) error
	Delete(client.ShaderClient) error
}

func NewVertexShader(data ShaderData) Shader {
	return &vertexShader{
		shader: shader{
			data: data,
		},
	}
}

func NewFragmentShader(data ShaderData) Shader {
	return &fragmentShader{
		shader: shader{
			data: data,
		},
	}
}

type shader struct {
	id   client.ShaderId
	data ShaderData
}

type vertexShader struct {
	shader
}

type fragmentShader struct {
	shader
}

func (s *shader) Data() ShaderData {
	return s.data
}

func (s *shader) Id() client.ShaderId {
	return s.id
}

func (s *shader) Created() bool {
	return s.id != nil
}

func (s *vertexShader) Create(shaderClient client.ShaderClient) error {
	var err error
	s.id, err = shaderClient.CreateVertexShader()
	if err != nil {
		return err
	}
	err = shaderClient.SetShaderSourceCode(s.id, s.data.SourceCode())
	if err != nil {
		return err
	}
	err = shaderClient.CompileShader(s.id)
	if err != nil {
		return err
	}
	return nil
}

func (s *fragmentShader) Create(shaderClient client.ShaderClient) error {
	var err error
	s.id, err = shaderClient.CreateFragmentShader()
	if err != nil {
		return err
	}
	err = shaderClient.SetShaderSourceCode(s.id, s.data.SourceCode())
	if err != nil {
		return err
	}
	err = shaderClient.CompileShader(s.id)
	if err != nil {
		return err
	}
	return nil
}

func (s *shader) Delete(shaderClient client.ShaderClient) error {
	err := shaderClient.DeleteShader(s.id)
	if err != nil {
		return err
	}
	s.id = nil
	return nil
}
