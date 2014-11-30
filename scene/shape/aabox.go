package shape

import "github.com/momchil-atanasov/go-whiskey/math"

type AABox struct {
	Position   math.Vec3
	HalfWidth  float32
	HalfHeight float32
	HalfDepth  float32
}

func (b *AABox) Width() float32 {
	return b.HalfWidth * 2.0
}

func (b *AABox) Height() float32 {
	return b.HalfHeight * 2.0
}

func (b *AABox) Depth() float32 {
	return b.HalfDepth * 2.0
}

func (b *AABox) TopLeftFrontCorner() math.Vec3 {
	return b.Position.IncCoords(-b.HalfWidth, b.HalfHeight, b.HalfDepth)
}

func (b *AABox) TopRightFrontCorner() math.Vec3 {
	return b.Position.IncCoords(b.HalfWidth, b.HalfHeight, b.HalfDepth)
}

func (b *AABox) BottomLeftFrontCorner() math.Vec3 {
	return b.Position.IncCoords(-b.HalfWidth, -b.HalfHeight, b.HalfDepth)
}

func (b *AABox) BottomRightFrontCorner() math.Vec3 {
	return b.Position.IncCoords(b.HalfWidth, -b.HalfHeight, b.HalfDepth)
}

func (b *AABox) TopLeftBackCorner() math.Vec3 {
	return b.Position.IncCoords(-b.HalfWidth, b.HalfHeight, -b.HalfDepth)
}

func (b *AABox) TopRightBackCorner() math.Vec3 {
	return b.Position.IncCoords(b.HalfWidth, b.HalfHeight, -b.HalfDepth)
}

func (b *AABox) BottomLeftBackCorner() math.Vec3 {
	return b.Position.IncCoords(-b.HalfWidth, -b.HalfHeight, -b.HalfDepth)
}

func (b *AABox) BottomRightBackCorner() math.Vec3 {
	return b.Position.IncCoords(b.HalfWidth, -b.HalfHeight, -b.HalfDepth)
}

func (b *AABox) ContainsPoint(point math.Vec3) bool {
	offset := point.DecVec3(b.Position)
	return (offset.X >= -b.HalfWidth) && (offset.X <= b.HalfWidth) &&
		(offset.Y >= -b.HalfHeight) && (offset.Y <= b.HalfHeight) &&
		(offset.Z >= -b.HalfDepth) && (offset.Z <= b.HalfDepth)
}
