package stage

import (
	"github.com/momchil-atanasov/go-whiskey/math"
	"github.com/momchil-atanasov/go-whiskey/scene/shape"
)

type NodeLocation int

const (
	TopLeftFront NodeLocation = iota
	TopRightFront
	BottomLeftFront
	BottomRightFront
	TopLeftBack
	TopRightBack
	BottomLeftBack
	BottomRightBack
)

type Node interface {
	Bounds() shape.AABox
	Child(NodeLocation) (Node, bool)
}

type node struct {
	position math.Vec3
	size     float32
	children []Node
}

func NewNode(size float32, depth int) Node {
	newNode := &node{}
	newNode.position = math.Vec3{}
	newNode.size = size
	newNode.createChildren(depth)
	return newNode
}

func (n *node) Bounds() shape.AABox {
	return shape.AABox{
		Position:   n.position,
		HalfWidth:  n.size / 2.0,
		HalfHeight: n.size / 2.0,
		HalfDepth:  n.size / 2.0,
	}
}

func (n *node) Child(position NodeLocation) (Node, bool) {
	if n.children == nil {
		return nil, false
	}
	return n.children[position], true
}

func (n *node) evaluateSize(parentNode *node) {
	n.size = parentNode.size / 2.0
}

func (n *node) evaluatePosition(parentNode *node, location NodeLocation) {
	offset := parentNode.size / 4.0
	position := parentNode.position
	position.X += offset * getLocationDirectionX(location)
	position.Y += offset * getLocationDirectionY(location)
	position.Z += offset * getLocationDirectionZ(location)
	n.position = position
}

func (n *node) createChildren(depth int) {
	if depth <= 1 {
		return
	}
	n.children = make([]Node, 8)
	n.children[TopLeftFront] = newChildNode(n, TopLeftFront, depth-1)
	n.children[TopRightFront] = newChildNode(n, TopRightFront, depth-1)
	n.children[BottomLeftFront] = newChildNode(n, BottomLeftFront, depth-1)
	n.children[BottomRightFront] = newChildNode(n, BottomRightFront, depth-1)
	n.children[TopLeftBack] = newChildNode(n, TopLeftBack, depth-1)
	n.children[TopRightBack] = newChildNode(n, TopRightBack, depth-1)
	n.children[BottomLeftBack] = newChildNode(n, BottomLeftBack, depth-1)
	n.children[BottomRightBack] = newChildNode(n, BottomRightBack, depth-1)
}

func newChildNode(parentNode *node, location NodeLocation, depth int) Node {
	childNode := &node{}
	childNode.evaluateSize(parentNode)
	childNode.evaluatePosition(parentNode, location)
	childNode.createChildren(depth)
	return childNode
}

func getLocationDirectionX(location NodeLocation) float32 {
	switch location {
	case TopLeftFront, BottomLeftFront, TopLeftBack, BottomLeftBack:
		return -1.0
	case TopRightFront, BottomRightFront, TopRightBack, BottomRightBack:
		return 1.0
	default:
		return 0.0
	}
}

func getLocationDirectionY(location NodeLocation) float32 {
	switch location {
	case TopLeftFront, TopRightFront, TopLeftBack, TopRightBack:
		return 1.0
	case BottomLeftFront, BottomRightFront, BottomLeftBack, BottomRightBack:
		return -1.0
	default:
		return 0.0
	}
}

func getLocationDirectionZ(location NodeLocation) float32 {
	switch location {
	case TopLeftFront, TopRightFront, BottomLeftFront, BottomRightFront:
		return 1.0
	case TopLeftBack, TopRightBack, BottomLeftBack, BottomRightBack:
		return -1.0
	default:
		return 0.0
	}
}
