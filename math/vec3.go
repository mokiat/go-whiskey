package math

import "fmt"

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func (v Vec3) Null() bool {
	return v == Vec3{}
}

func (v Vec3) Inverse() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (v Vec3) IncCoords(x, y, z float32) Vec3 {
	return Vec3{
		X: v.X + x,
		Y: v.Y + y,
		Z: v.Z + z,
	}
}

func (v Vec3) IncVec3(other Vec3) Vec3 {
	return v.IncCoords(other.X, other.Y, other.Z)
}

func (v Vec3) DecCoords(x, y, z float32) Vec3 {
	return Vec3{
		X: v.X - x,
		Y: v.Y - y,
		Z: v.Z - z,
	}
}

func (v Vec3) DecVec3(other Vec3) Vec3 {
	return v.DecCoords(other.X, other.Y, other.Z)
}

func (v Vec3) Mul(amount float32) Vec3 {
	return Vec3{
		X: v.X * amount,
		Y: v.Y * amount,
		Z: v.Z * amount,
	}
}

func (v Vec3) Div(amount float32) Vec3 {
	return Vec3{
		X: v.X / amount,
		Y: v.Y / amount,
		Z: v.Z / amount,
	}
}

func (v Vec3) LengthSquared() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float32 {
	return Sqrt32(v.LengthSquared())
}

func (v Vec3) Resize(desiredLength float32) Vec3 {
	ratio := desiredLength / v.Length()
	return v.Mul(ratio)
}

func (v Vec3) DistanceToCoords(x, y, z float32) float32 {
	delta := v.DecCoords(x, y, z)
	return delta.Length()
}

func (v Vec3) DistanceToVec3(other Vec3) float32 {
	return v.DistanceToCoords(other.X, other.Y, other.Z)
}

func (v Vec3) GoString() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

func NullVec3() Vec3 {
	return Vec3{}
}

func BaseVec3X() Vec3 {
	return Vec3{1.0, 0.0, 0.0}
}

func BaseVec3Y() Vec3 {
	return Vec3{0.0, 1.0, 0.0}
}

func BaseVec3Z() Vec3 {
	return Vec3{0.0, 0.0, 1.0}
}

func MakeVec3(x, y, z float32) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func Vec3DotProduct(a, b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Vec3CrossProduct(a, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}
