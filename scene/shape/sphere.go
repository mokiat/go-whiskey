package shape

import "github.com/mokiat/go-whiskey/math"

type Sphere struct {
	Position math.Vec3
	Radius   float32
}

func (s *Sphere) Diameter() float32 {
	return s.Radius * 2.0
}

func (s *Sphere) ContainsPoint(point math.Vec3) bool {
	offset := point.DecVec3(s.Position)
	return (offset.LengthSquared() < s.Radius*s.Radius)
}
