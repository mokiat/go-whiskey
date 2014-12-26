package shader

import "github.com/momchil-atanasov/go-whiskey/graphics"

type Shader interface {
	Id() graphics.ResourceId
	SourceCode() string
	CreateRemotely() error
	DeleteRemotely()
	CreatedRemotely() bool
}

type shader struct {
	id         graphics.ResourceId
	facade     graphics.Facade
	sourceCode string
}

func (s *shader) Id() graphics.ResourceId {
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
