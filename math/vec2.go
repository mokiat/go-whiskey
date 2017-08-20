package math

import "math"

type Vec2 struct {
	X float32
	Y float32
}

func (v Vec2) Null() bool {
	return v == Vec2{}
}

func (v Vec2) Inverse() Vec2 {
	return Vec2{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vec2) IncCoords(x, y float32) Vec2 {
	return Vec2{
		X: v.X + x,
		Y: v.Y + y,
	}
}

func (v Vec2) IncVec2(other Vec2) Vec2 {
	return v.IncCoords(other.X, other.Y)
}

func (v Vec2) DecCoords(x, y float32) Vec2 {
	return Vec2{
		X: v.X - x,
		Y: v.Y - y,
	}
}

func (v Vec2) DecVec2(other Vec2) Vec2 {
	return v.DecCoords(other.X, other.Y)
}

func (v Vec2) Mul(amount float32) Vec2 {
	return Vec2{
		X: v.X * amount,
		Y: v.Y * amount,
	}
}

func (v Vec2) Div(amount float32) Vec2 {
	return Vec2{
		X: v.X / amount,
		Y: v.Y / amount,
	}
}

func (v Vec2) LengthSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Length() float32 {
	return float32(math.Sqrt(float64(v.LengthSquared())))
}

func (v Vec2) Resize(desiredLength float32) Vec2 {
	ratio := desiredLength / v.Length()
	return v.Mul(ratio)
}

func (v Vec2) DistanceToCoords(x, y float32) float32 {
	delta := v.DecCoords(x, y)
	return delta.Length()
}

func (v Vec2) DistanceToVec2(other Vec2) float32 {
	return v.DistanceToCoords(other.X, other.Y)
}

func NullVec2() Vec2 {
	return Vec2{}
}

func BaseVec2X() Vec2 {
	return Vec2{
		X: 1.0,
		Y: 0.0,
	}
}

func BaseVec2Y() Vec2 {
	return Vec2{
		X: 0.0,
		Y: 1.0,
	}
}

func MakeVec2(x, y float32) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

func Vec2DotProduct(a, b Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}
