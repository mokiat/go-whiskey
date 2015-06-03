package shader

import "github.com/momchil-atanasov/go-whiskey/graphics/client"

//go:generate counterfeiter -o shader_fakes/fake_shader.go ./ Shader

type Shader interface {
	SourceCode() string
	Remote() RemoteShader
}

//go:generate counterfeiter -o shader_fakes/fake_remote_shader.go ./ RemoteShader

type RemoteShader interface {
	Id() client.ShaderId
	Created() bool
	Create() error
	Delete() error
}

func NewVertexShader(shaderClient client.ShaderClient, sourceCode string) Shader {
	return &vertexShader{
		shader: shader{
			shaderClient: shaderClient,
			sourceCode:   sourceCode,
		},
	}
}

func NewFragmentShader(shaderClient client.ShaderClient, sourceCode string) Shader {
	return &fragmentShader{
		shader: shader{
			shaderClient: shaderClient,
			sourceCode:   sourceCode,
		},
	}
}

type shader struct {
	shaderClient client.ShaderClient
	sourceCode   string
	id           client.ShaderId
}

type vertexShader struct {
	shader
}

type fragmentShader struct {
	shader
}

func (s *shader) SourceCode() string {
	return s.sourceCode
}

func (s *vertexShader) Remote() RemoteShader {
	return s
}

func (s *fragmentShader) Remote() RemoteShader {
	return s
}

func (s *shader) Id() client.ShaderId {
	return s.id
}

func (s *shader) Created() bool {
	return s.id != nil
}

func (s *vertexShader) Create() error {
	var err error
	s.id, err = s.shaderClient.CreateVertexShader()
	if err != nil {
		return err
	}
	err = s.shaderClient.SetShaderSourceCode(s.id, s.sourceCode)
	if err != nil {
		return err
	}
	err = s.shaderClient.CompileShader(s.id)
	if err != nil {
		return err
	}
	return nil
}

func (s *fragmentShader) Create() error {
	var err error
	s.id, err = s.shaderClient.CreateFragmentShader()
	if err != nil {
		return err
	}
	err = s.shaderClient.SetShaderSourceCode(s.id, s.sourceCode)
	if err != nil {
		return err
	}
	err = s.shaderClient.CompileShader(s.id)
	if err != nil {
		return err
	}
	return nil
}

func (s *shader) Delete() error {
	err := s.shaderClient.DeleteShader(s.id)
	if err != nil {
		return err
	}
	s.id = nil
	return nil
}
