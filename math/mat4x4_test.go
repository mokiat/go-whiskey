package math_test

import (
	. "github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mat4x4", func() {
	var vector Vec4
	var matrix Mat4x4
	var otherMatrix Mat4x4

	BeforeEach(func() {
		vector = MakeVec4(2.5, 1.5, 3.0, 1.0)
		matrix = MakeMath4x4RowOrder(
			0.1, 0.2, 0.3, 0.4,
			0.5, 0.6, 0.7, 0.8,
			0.9, 1.0, 1.1, 1.2,
			1.3, 1.4, 1.5, 1.6,
		)
		otherMatrix = MakeMath4x4RowOrder(
			0.5, 0.3, 0.2, 0.0,
			0.2, 0.8, 0.7, 0.4,
			0.1, 0.2, 0.9, 0.8,
			0.6, 0.6, 0.3, 0.1,
		)
	})

	It("#Mul", func() {
		result := matrix.Mul(2.0)
		Ω(result).Should(EqualMat4x4(MakeMath4x4RowOrder(
			0.2, 0.4, 0.6, 0.8,
			1.0, 1.2, 1.4, 1.6,
			1.8, 2.0, 2.2, 2.4,
			2.6, 2.8, 3.0, 3.2,
		)))
	})

	It("#MulVec4", func() {
		result := matrix.MulVec4(vector)
		AssertVec4Equals(result, 1.85, 5.05, 8.25, 11.45)
	})

	It("#MulMat4x4", func() {
		result := matrix.MulMat4x4(otherMatrix)
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

	It("IdentityMat4x4", func() {
		result := IdentityMat4x4().MulMat4x4(matrix)
		Ω(result).Should(Equal(matrix))
	})

	It("TranslationMat4x4", func() {
		translationMatrix := TranslationMat4x4(2.0, -3.0, 4.0)
		result := translationMatrix.MulVec4(vector)
		AssertVec4Equals(result, 4.5, -1.5, 7.0, 1.0)
	})

	It("ScaleMat4x4", func() {
		scaleMatrix := ScaleMat4x4(2.0, -3.0, 4.0)
		result := scaleMatrix.MulVec4(vector)
		AssertVec4Equals(result, 5.0, -4.5, 12.0, 1.0)
	})

	It("RotationMat4x4", func() {
		vector = MakeVec4(1.0, 0.0, 0.0, 1.0)
		rotationMatrix := RotationMat4x4(120.0, 1.0, 1.0, 1.0)
		result := rotationMatrix.MulVec4(vector)
		AssertVec4Equals(result, 0.0, 1.0, 0.0, 1.0)
	})

	It("DirectionXCoords", func() {
		matrix := NullMat4x4()
		matrix = matrix.DirectionXCoords(1.1, 2.1, 3.1)
		vector := MakeVec4(2.0, 9.0, 9.0, 9.0)
		position := matrix.MulVec4(vector)
		AssertVec4Equals(position, 2.2, 4.2, 6.2, 0.0)
	})

	It("DirectionYCoords", func() {
		matrix := NullMat4x4()
		matrix = matrix.DirectionYCoords(1.1, 2.1, 3.1)
		vector := MakeVec4(9.0, 2.0, 9.0, 9.0)
		position := matrix.MulVec4(vector)
		AssertVec4Equals(position, 2.2, 4.2, 6.2, 0.0)
	})

	It("DirectionZCoords", func() {
		matrix := NullMat4x4()
		matrix = matrix.DirectionZCoords(1.1, 2.1, 3.1)
		vector := MakeVec4(9.0, 9.0, 2.0, 9.0)
		position := matrix.MulVec4(vector)
		AssertVec4Equals(position, 2.2, 4.2, 6.2, 0.0)
	})

	It("Reposition", func() {
		matrix := IdentityMat4x4()
		matrix = matrix.RepositionCoords(1.1, 2.2, 3.3)
		vector := MakeVec4(0.0, 0.0, 0.0, 1.0)
		position := matrix.MulVec4(vector)
		AssertVec4Equals(position, 1.1, 2.2, 3.3, 1.0)
	})

	It("QuickInverse", func() {
		matrix = Mat4x4MulMany(
			TranslationMat4x4(1.5, 2.3, 3.7),
			RotationMat4x4(45.0, 0.5, 0.3, 0.2),
		)
		inverseMatrix := matrix.QuickInverse()
		productMatrix := inverseMatrix.MulMat4x4(matrix)
		Ω(productMatrix).Should(EqualMat4x4(IdentityMat4x4()))
	})

	It("Inverse", func() {
		matrix := MakeMath4x4RowOrder(
			4.0, 3.0, 2.0, 1.0,
			1.1, 4.1, 3.1, 2.1,
			2.2, 3.2, 4.2, 1.2,
			3.3, 2.3, 1.3, 4.3,
		)
		inverseMatrix := matrix.Inverse()
		productMatrix := inverseMatrix.MulMat4x4(matrix)
		Ω(productMatrix).Should(EqualMat4x4(IdentityMat4x4()))
	})

	It("VectorMat4x4", func() {
		matrix := VectorMat4x4(
			MakeVec3(-1.0, 0.0, 0.0),
			MakeVec3(0.0, -1.0, 0.0),
			MakeVec3(0.0, 0.0, -1.0),
			MakeVec3(4.4, 5.5, 6.6),
		)
		vector := MakeVec4(2.0, 3.0, 4.0, 1.0)
		position := matrix.MulVec4(vector)
		AssertVec4Equals(position, 2.4, 2.5, 2.6, 1.0)
	})

	It("OrthoMat4x4", func() {
		orthoMatrix := OrthoMat4x4(-1.1, 2.1, 1.5, -3.4, 1.7, 3.8)

		// Test two opposite corner projections
		nearCorner := MakeVec4(-1.1, -3.4, -1.7, 1.0)
		projectedNearCorner := orthoMatrix.MulVec4(nearCorner)
		projectedNearCorner = projectedNearCorner.Div(projectedNearCorner.W)
		AssertVec4Equals(projectedNearCorner, -1.0, -1.0, -1.0, 1.0)

		farCorner := MakeVec4(2.1, 1.5, -3.8, 1.0)
		projectedFarCorner := orthoMatrix.MulVec4(farCorner)
		projectedFarCorner = projectedFarCorner.Div(projectedFarCorner.W)
		AssertVec4Equals(projectedFarCorner, 1.0, 1.0, 1.0, 1.0)
	})

	It("PerspectiveMat4x4", func() {
		perspectiveMatrix := PerspectiveMat4x4(-1.1, 2.1, -3.4, 1.5, 1.7, 3.8)

		// Test two opposite corner projections
		nearCorner := MakeVec4(-1.1, -3.4, -1.7, 1.0)
		projectedNearCorner := perspectiveMatrix.MulVec4(nearCorner)
		projectedNearCorner = projectedNearCorner.Div(projectedNearCorner.W)
		AssertVec4Equals(projectedNearCorner, -1.0, -1.0, -1.0, 1.0)

		farCorner := MakeVec4(4.6941, 3.3529, -3.8, 1.0)
		projectedFarCorner := perspectiveMatrix.MulVec4(farCorner)
		projectedFarCorner = projectedFarCorner.Div(projectedFarCorner.W)
		AssertVec4Equals(projectedFarCorner, 1.0, 1.0, 1.0, 1.0)
	})

	It("Mat4x4MulMany", func() {
		matrix = Mat4x4MulMany(
			TranslationMat4x4(2.0, 3.0, 5.0),
			ScaleMat4x4(2.0, 4.0, 8.0),
		)
		vector := matrix.MulVec4(MakeVec4(1.0, 1.0, 1.0, 1.0))
		AssertVec4Equals(vector, 4.0, 7.0, 13.0, 1.0)
	})
})
