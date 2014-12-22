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

func (s *vertexShader) CreateRemotely() {
	s.id = s.facade.CreateVertexShader()
	s.facade.SetShaderSourceCode(s.id, s.sourceCode)
	s.facade.CompileShader(s.id)
}