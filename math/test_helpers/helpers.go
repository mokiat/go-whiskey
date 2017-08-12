package test_helpers

import (
	"github.com/mokiat/go-whiskey/math"
	. "github.com/onsi/gomega"
)

const FloatMargin = 0.0001

func AssertFloatEquals(actualValue, expectedValue float32) {
	Ω(actualValue).Should(BeNumerically("~", expectedValue, FloatMargin))
}

func AssertVec3Equals(vector math.Vec3, expectedX, expectedY, expectedZ float32) {
	AssertFloatEquals(vector.X, expectedX)
	AssertFloatEquals(vector.Y, expectedY)
	AssertFloatEquals(vector.Z, expectedZ)
}

func AssertVec4Equals(vector math.Vec4, expectedX, expectedY, expectedZ, expectedW float32) {
	AssertFloatEquals(vector.X, expectedX)
	AssertFloatEquals(vector.Y, expectedY)
	AssertFloatEquals(vector.Z, expectedZ)
	AssertFloatEquals(vector.W, expectedW)
}

func AssertMat4x4Equals(matrix math.Mat4x4,
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
