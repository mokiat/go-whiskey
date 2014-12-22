package attribute

import "github.com/momchil-atanasov/go-whiskey/common"

type Attribute common.Enum

const (
	Coord Attribute = iota
	Normal
	TexCoord
	ATTRIBUTE_COUNT int = iota
)

var names = make([]string, ATTRIBUTE_COUNT)

func init() {
	names[Coord] = "coordIn"
	names[Normal] = "normalIn"
	names[TexCoord] = "texCoordIn"
}

var dimensions = make([]int, ATTRIBUTE_COUNT)

func init() {
	dimensions[Coord] = 3
	dimensions[Normal] = 3
	dimensions[TexCoord] = 2
}

func (a Attribute) Name() string {
	return names[a]
}

func (a Attribute) Dimensions() int {
	return dimensions[a]
}
