package shader

import "github.com/momchil-atanasov/go-whiskey/graphics"

type VertexShader interface {
	Shader
}

type vertexShader struct {
	shader
}

func NewVertexShader(facade graphics.Facade, sourceCode string) VertexShader {
	return &vertexShader{
		shader: shader{
			id:         graphics.InvalidShaderId,
			facade:     facade,
			sourceCode: sourceCode,
		},
	}
}

func (s *vertexShader) CreateRemotely() error {
	var err error
	s.id, err = s.facade.CreateVertexShader()
	if err != nil {
		s.id = graphics.InvalidShaderId
		return err
	}
	s.facade.SetShaderSourceCode(s.id, s.sourceCode)
	err = s.facade.CompileShader(s.id)
	if err != nil {
		return err
	}
	return nil
}
