package shader

import "github.com/momchil-atanasov/go-whiskey/graphics"

type Shader interface {
	Id() int
	SourceCode() string
	CreateRemotely()
	DeleteRemotely()
	CreatedRemotely() bool
}

type shader struct {
	id         int
	facade     graphics.Facade
	sourceCode string
}

func (s *shader) Id() int {
	return s.id
}

func (s *shader) SourceCode() string {
	return s.sourceCode
}

func (s *shader) DeleteRemotely() {
	s.facade.DeleteShader(s.id)
	s.id = graphics.InvalidShaderId
}

func (s *shader) CreatedRemotely() bool {
	return s.id != graphics.InvalidShaderId
}
