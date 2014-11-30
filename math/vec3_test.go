package math_test

import (
	. "github.com/momchil-atanasov/go-whiskey/math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func AssertVec3Equals(vector Vec3, expectedX, expectedY, expectedZ float32) {
	AssertFloatEquals(vector.X, expectedX)
	AssertFloatEquals(vector.Y, expectedY)
	AssertFloatEquals(vector.Z, expectedZ)
}

var _ = Describe("Vec3", func() {

	var nullVector Vec3
	var firstVector Vec3
	var secondVector Vec3

	BeforeEach(func() {
		nullVector = Vec3{}
		firstVector = Vec3{
			X: 2.0,
			Y: 3.0,
			Z: 4.0,
		}
		secondVector = Vec3{
			-1.0,
			2.0,
			-3.0,
		}
	})

	It("#Null", func() {
		Ω(nullVector.Null()).Should(BeTrue())
		Ω(firstVector.Null()).Should(BeFalse())
	})

	It("#Inverse", func() {
		inverted := firstVector.Inverse()
		AssertVec3Equals(inverted, -2.0, -3.0, -4.0)
	})

	It("#IncCoords", func() {
		incremented := firstVector.IncCoords(1.5, -2.5, 5.0)
		AssertVec3Equals(incremented, 3.5, 0.5, 9.0)
	})

	It("#DecCoords", func() {
		decremented := firstVector.DecCoords(1.5, -2.5, 5.0)
		AssertVec3Equals(decremented, 0.5, 5.5, -1.0)
	})

	It("#Mul", func() {
		multiplied := firstVector.Mul(0.5)
		AssertVec3Equals(multiplied, 1.0, 1.5, 2.0)
	})

	It("#Div", func() {
		divided := firstVector.Div(2.0)
		AssertVec3Equals(divided, 1.0, 1.5, 2.0)
	})

	It("#LengthSquared", func() {
		squaredLength := firstVector.LengthSquared()
		AssertFloatEquals(squaredLength, 29.0)
	})

	It("#Length", func() {
		length := firstVector.Length()
		AssertFloatEquals(length, 5.385164807134504)
	})

	It("#Resize", func() {
		resized := firstVector.Resize(10.77032961426901)
		AssertVec3Equals(resized, 4.0, 6.0, 8.0)
	})

	It("#DistanceToCoords", func() {
		distance := firstVector.DistanceToCoords(-1.0, 2.0, -3.0)
		AssertFloatEquals(distance, 7.681145747868608)
	})

	It("#DistanceToVec3", func() {
		distance := firstVector.DistanceToVec3(secondVector)
		AssertFloatEquals(distance, 7.681145747868608)
	})

	It("NullVec3", func() {
		AssertVec3Equals(NullVec3(), 0.0, 0.0, 0.0)
	})

	It("BaseVec3X", func() {
		AssertVec3Equals(BaseVec3X(), 1.0, 0.0, 0.0)
	})

	It("BaseVec3Y", func() {
		AssertVec3Equals(BaseVec3Y(), 0.0, 1.0, 0.0)
	})

	It("BaseVec3Z", func() {
		AssertVec3Equals(BaseVec3Z(), 0.0, 0.0, 1.0)
	})

	It("Vec3DotProduct", func() {
		dot := Vec3DotProduct(firstVector, secondVector)
		AssertFloatEquals(dot, -8.0)
	})

	It("Vec3CrossProduct", func() {
		cross := Vec3CrossProduct(firstVector, secondVector)
		AssertVec3Equals(cross, -17.0, 2.0, 7.0)
	})
})
