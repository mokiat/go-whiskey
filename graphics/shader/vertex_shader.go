package shader

import "github.com/momchil-atanasov/go-whiskey/graphics"

type VertexShader interface {
	Id() int
	SourceCode() string
	CreateRemotely()
	DeleteRemotely()
	CreatedRemotely() bool
}

type vertexShader struct {
	id         int
	facade     graphics.Facade
	sourceCode string
}

func NewVertexShader(facade graphics.Facade, sourceCode string) VertexShader {
	return &vertexShader{
		id:         graphics.InvalidShaderId,
		facade:     facade,
		sourceCode: sourceCode,
	}
}

func (s *vertexShader) Id() int {
	return s.id
}

func (s *vertexShader) SourceCode() string {
	return s.sourceCode
}

func (s *vertexShader) CreateRemotely() {
	s.id = s.facade.CreateVertexShader()
	s.facade.SetShaderSourceCode(s.id, s.sourceCode)
	s.facade.CompileShader(s.id)
}

func (s *vertexShader) DeleteRemotely() {
	s.facade.DeleteShader(s.id)
	s.id = graphics.InvalidShaderId
}

func (s *vertexShader) CreatedRemotely() bool {
	return s.id != graphics.InvalidShaderId
}
