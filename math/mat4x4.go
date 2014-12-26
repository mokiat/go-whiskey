package math

import "math"

type Mat4x4 struct {
	M11, M12, M13, M14 float32
	M21, M22, M23, M24 float32
	M31, M32, M33, M34 float32
	M41, M42, M43, M44 float32
}

func (m Mat4x4) MulVec4(vec Vec4) Vec4 {
	return Vec4{
		m.M11*vec.X + m.M12*vec.Y + m.M13*vec.Z + m.M14*vec.W,
		m.M21*vec.X + m.M22*vec.Y + m.M23*vec.Z + m.M24*vec.W,
		m.M31*vec.X + m.M32*vec.Y + m.M33*vec.Z + m.M34*vec.W,
		m.M41*vec.X + m.M42*vec.Y + m.M43*vec.Z + m.M44*vec.W,
	}
}

func (m Mat4x4) MulMat4x4(other Mat4x4) Mat4x4 {
	return Mat4x4{
		m.M11*other.M11 + m.M12*other.M21 + m.M13*other.M31 + m.M14*other.M41,
		m.M11*other.M12 + m.M12*other.M22 + m.M13*other.M32 + m.M14*other.M42,
		m.M11*other.M13 + m.M12*other.M23 + m.M13*other.M33 + m.M14*other.M43,
		m.M11*other.M14 + m.M12*other.M24 + m.M13*other.M34 + m.M14*other.M44,

		m.M21*other.M11 + m.M22*other.M21 + m.M23*other.M31 + m.M24*other.M41,
		m.M21*other.M12 + m.M22*other.M22 + m.M23*other.M32 + m.M24*other.M42,
		m.M21*other.M13 + m.M22*other.M23 + m.M23*other.M33 + m.M24*other.M43,
		m.M21*other.M14 + m.M22*other.M24 + m.M23*other.M34 + m.M24*other.M44,

		m.M31*other.M11 + m.M32*other.M21 + m.M33*other.M31 + m.M34*other.M41,
		m.M31*other.M12 + m.M32*other.M22 + m.M33*other.M32 + m.M34*other.M42,
		m.M31*other.M13 + m.M32*other.M23 + m.M33*other.M33 + m.M34*other.M43,
		m.M31*other.M14 + m.M32*other.M24 + m.M33*other.M34 + m.M34*other.M44,

		m.M41*other.M11 + m.M42*other.M21 + m.M43*other.M31 + m.M44*other.M41,
		m.M41*other.M12 + m.M42*other.M22 + m.M43*other.M32 + m.M44*other.M42,
		m.M41*other.M13 + m.M42*other.M23 + m.M43*other.M33 + m.M44*other.M43,
		m.M41*other.M14 + m.M42*other.M24 + m.M43*other.M34 + m.M44*other.M44,
	}
}

func NullMat4x4() Mat4x4 {
	return Mat4x4{}
}

func IdentityMat4x4() Mat4x4 {
	return Mat4x4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func TranslationMat4x4(x, y, z float32) Mat4x4 {
	return Mat4x4{
		1.0, 0.0, 0.0, x,
		0.0, 1.0, 0.0, y,
		0.0, 0.0, 1.0, z,
		0.0, 0.0, 0.0, 1.0,
	}
}

func ScaleMat4x4(x, y, z float32) Mat4x4 {
	return Mat4x4{
		x, 0.0, 0.0, 0.0,
		0.0, y, 0.0, 0.0,
		0.0, 0.0, z, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func RotationMat4x4(angle float64, x, y, z float32) Mat4x4 {
	radians := angle * math.Pi / 180.0
	cs := float32(math.Cos(radians))
	sn := float32(math.Sin(radians))
	vector := Vec3{x, y, z}.Resize(1.0)
	return getRotationMat4x4FromNormalizedData(cs, sn, vector.X, vector.Y, vector.Z)
}

func getRotationMat4x4FromNormalizedData(cs, sn, x, y, z float32) Mat4x4 {
	result := Mat4x4{}
	result.M11 = x*x*(1.0-cs) + cs
	result.M21 = x*y*(1.0-cs) + z*sn
	result.M31 = x*z*(1.0-cs) - y*sn
	result.M41 = 0.0

	result.M12 = y*x*(1.0-cs) - z*sn
	result.M22 = y*y*(1.0-cs) + cs
	result.M32 = y*z*(1.0-cs) + x*sn
	result.M42 = 0.0

	result.M13 = z*x*(1.0-cs) + y*sn
	result.M23 = z*y*(1.0-cs) - x*sn
	result.M33 = z*z*(1.0-cs) + cs
	result.M43 = 0.0

	result.M14 = 0.0
	result.M24 = 0.0
	result.M34 = 0.0
	result.M44 = 1.0
	return result
}

func OrthoMat4x4(left, right, top, bottom, near, far float32) Mat4x4 {
	result := Mat4x4{}
	result.M11 = 2.0 / (right - left)
	result.M12 = 0.0
	result.M13 = 0.0
	result.M14 = (right + left) / (left - right)

	result.M21 = 0.0
	result.M22 = 2.0 / (top - bottom)
	result.M23 = 0.0
	result.M24 = (top + bottom) / (bottom - top)

	result.M31 = 0.0
	result.M32 = 0.0
	result.M33 = 2.0 / (near - far)
	result.M34 = (far + near) / (near - far)

	result.M41 = 0.0
	result.M42 = 0.0
	result.M43 = 0.0
	result.M44 = 1.0
	return result
}

func PerspectiveMat4x4(left, right, top, bottom, near, far float32) Mat4x4 {
	result := Mat4x4{}

	result.M11 = 2.0 * near / (right - left)
	result.M12 = 0.0
	result.M13 = (right + left) / (right - left)
	result.M14 = 0.0

	result.M21 = 0.0
	result.M22 = 2.0 * near / (top - bottom)
	result.M23 = (top + bottom) / (top - bottom)
	result.M24 = 0.0

	result.M31 = 0.0
	result.M32 = 0.0
	result.M33 = (far + near) / (near - far)
	result.M34 = 2.0 * far * near / (near - far)

	result.M41 = 0.0
	result.M42 = 0.0
	result.M43 = -1.0
	result.M44 = 0.0
	return result
}
