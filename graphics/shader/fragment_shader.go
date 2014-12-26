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

func (s *fragmentShader) CreateRemotely() error {
	var err error
	s.id, err = s.facade.CreateFragmentShader()
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
