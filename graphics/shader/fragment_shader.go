package shader

import "github.com/momchil-atanasov/go-whiskey/graphics"

type FragmentShader interface {
	Shader
}

type fragmentShader struct {
	shader
}

func NewFragmentShader(facade graphics.Facade, sourceCode string) FragmentShader {
	return &fragmentShader{
		shader: shader{
			id:         graphics.InvalidShaderId,
			facade:     facade,
			sourceCode: sourceCode,
		},
	}
}

func (s *fragmentShader) CreateRemotely() {
	s.id = s.facade.CreateFragmentShader()
	s.facade.SetShaderSourceCode(s.id, s.sourceCode)
	s.facade.CompileShader(s.id)
}
