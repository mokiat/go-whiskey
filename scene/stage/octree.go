package stage

type Octree interface {
	Root() Node
}

type octree struct {
	root Node
}

func NewOctree(size float32, depth int) Octree {
	return &octree{
		root: NewNode(size, depth),
	}
}

func (t *octree) Root() Node {
	return t.root
}
