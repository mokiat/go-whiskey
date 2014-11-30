package math_test

import (
	. "github.com/momchil-atanasov/go-whiskey/math"

	. "github.com/onsi/ginkgo"
)

func AssertMat4x4Equals(matrix Mat4x4,
	m11, m12, m13, m14,
	m21, m22, m23, m24,
	m31, m32, m33, m34,
	m41, m42, m43, m44 float32) {

	AssertFloatEquals(matrix.M11, m11)
	AssertFloatEquals(matrix.M12, m12)
	AssertFloatEquals(matrix.M13, m13)
	AssertFloatEquals(matrix.M14, m14)

	AssertFloatEquals(matrix.M21, m21)
	AssertFloatEquals(matrix.M22, m22)
	AssertFloatEquals(matrix.M23, m23)
	AssertFloatEquals(matrix.M24, m24)

	AssertFloatEquals(matrix.M31, m31)
	AssertFloatEquals(matrix.M32, m32)
	AssertFloatEquals(matrix.M33, m33)
	AssertFloatEquals(matrix.M34, m34)

	AssertFloatEquals(matrix.M41, m41)
	AssertFloatEquals(matrix.M42, m42)
	AssertFloatEquals(matrix.M43, m43)
	AssertFloatEquals(matrix.M44, m44)
}

var _ = Describe("Mat4x4", func() {
	var vector Vec4
	var matrix Mat4x4
	var otherMatrix Mat4x4

	BeforeEach(func() {
		vector = Vec4{2.5, 1.5, 3.0, 1.0}
		matrix = Mat4x4{
			0.1, 0.2, 0.3, 0.4,
			0.5, 0.6, 0.7, 0.8,
			0.9, 1.0, 1.1, 1.2,
			1.3, 1.4, 1.5, 1.6,
		}
		otherMatrix = Mat4x4{
			0.5, 0.3, 0.2, 0.0,
			0.2, 0.8, 0.7, 0.4,
			0.1, 0.2, 0.9, 0.8,
			0.6, 0.6, 0.3, 0.1,
		}
	})

	It("#MulVec4Right", func() {
		result := matrix.MulVec4Right(vector)
		AssertVec4Equals(result, 1.85, 5.05, 8.25, 11.45)
	})

	It("#MulMat4x4Right", func() {
		result := matrix.MulMat4x4Right(otherMatrix)
		AssertMat4x4Equals(result,
			0.36, 0.49, 0.55, 0.36,
			0.92, 1.25, 1.3900001, 0.88,
			1.48, 2.0100002, 2.23, 1.4000001,
			2.04, 2.77, 3.07, 1.92)
	})

	It("NullMat4x4", func() {
		AssertMat4x4Equals(NullMat4x4(),
			0.0, 0.0, 0.0, 0.0,
			0.0, 0.0, 0.0, 0.0,
			0.0, 0.0, 0.0, 0.0,
			0.0, 0.0, 0.0, 0.0)
	})
})
