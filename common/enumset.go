package common

type EnumBlock func(Enum)

type EnumSet interface {
	Empty() bool
	Size() int
	Add(enum Enum)
	Remove(enum Enum)
	Contains(enum Enum) bool
	Each(block EnumBlock)
	Clear()
}

// XXX: Consider using bit flags instead of
// boolean array
type enumSet struct {
	size  int
	flags []bool
}

func NewEnumSet(enumSize int) EnumSet {
	return &enumSet{
		size:  0,
		flags: make([]bool, enumSize),
	}
}

func (s *enumSet) Empty() bool {
	return s.size == 0
}

func (s *enumSet) Size() int {
	return s.size
}

func (s *enumSet) Add(enum Enum) {
	if s.Contains(enum) {
		return
	}
	s.size++
	s.flags[enum] = true
}

func (s *enumSet) Remove(enum Enum) {
	if !s.Contains(enum) {
		return
	}
	s.size--
	s.flags[enum] = false
}

func (s *enumSet) Contains(enum Enum) bool {
	return s.flags[enum]
}

func (s *enumSet) Each(block EnumBlock) {
	for i, available := range s.flags {
		if available {
			block(Enum(i))
		}
	}
}

func (s *enumSet) Clear() {
	s.size = 0
	for i := range s.flags {
		s.flags[i] = false
	}
}
