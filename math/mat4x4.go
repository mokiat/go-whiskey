package math

type Mat4x4 struct {
	M11, M12, M13, M14 float32
	M21, M22, M23, M24 float32
	M31, M32, M33, M34 float32
	M41, M42, M43, M44 float32
}

func (m Mat4x4) MulVec4Right(vec Vec4) Vec4 {
	return Vec4{
		m.M11*vec.X + m.M12*vec.Y + m.M13*vec.Z + m.M14*vec.W,
		m.M21*vec.X + m.M22*vec.Y + m.M23*vec.Z + m.M24*vec.W,
		m.M31*vec.X + m.M32*vec.Y + m.M33*vec.Z + m.M34*vec.W,
		m.M41*vec.X + m.M42*vec.Y + m.M43*vec.Z + m.M44*vec.W,
	}
}

func (m Mat4x4) MulMat4x4Right(other Mat4x4) Mat4x4 {
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
