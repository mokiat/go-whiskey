package vec3_test

import (
	. "github.com/momchil-atanasov/go-whiskey/math/vec3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vec3", func() {

	var nullVector Vec3
	var firstVector Vec3
	var secondVector Vec3

	assertFloatEquals := func(actualValue, expectedValue float32) {
		Ω(actualValue).Should(BeNumerically("~", expectedValue))
	}

	assertVec3Equals := func(vector Vec3, expectedX, expectedY, expectedZ float32) {
		assertFloatEquals(vector.X, expectedX)
		assertFloatEquals(vector.Y, expectedY)
		assertFloatEquals(vector.Z, expectedZ)
	}

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
		assertVec3Equals(inverted, -2.0, -3.0, -4.0)
	})

	It("#IncCoords", func() {
		incremented := firstVector.IncCoords(1.5, -2.5, 5.0)
		assertVec3Equals(incremented, 3.5, 0.5, 9.0)
	})

	It("#DecCoords", func() {
		decremented := firstVector.DecCoords(1.5, -2.5, 5.0)
		assertVec3Equals(decremented, 0.5, 5.5, -1.0)
	})

	It("#Mul", func() {
		multiplied := firstVector.Mul(0.5)
		assertVec3Equals(multiplied, 1.0, 1.5, 2.0)
	})

	It("#Div", func() {
		divided := firstVector.Div(2.0)
		assertVec3Equals(divided, 1.0, 1.5, 2.0)
	})

	It("#LengthSquared", func() {
		squaredLength := firstVector.LengthSquared()
		assertFloatEquals(squaredLength, 29.0)
	})

	It("#Length", func() {
		length := firstVector.Length()
		assertFloatEquals(length, 5.385164807134504)
	})

	It("#Resize", func() {
		resized := firstVector.Resize(10.77032961426901)
		assertVec3Equals(resized, 4.0, 6.0, 8.0)
	})

	It("#DistanceToCoords", func() {
		distance := firstVector.DistanceToCoords(-1.0, 2.0, -3.0)
		assertFloatEquals(distance, 7.681145747868608)
	})

	It("#DistanceToVec3", func() {
		distance := firstVector.DistanceToVec3(secondVector)
		assertFloatEquals(distance, 7.681145747868608)
	})

	It("BaseVectorX", func() {
		assertVec3Equals(BaseVectorX(), 1.0, 0.0, 0.0)
	})

	It("Null", func() {
		assertVec3Equals(Null(), 0.0, 0.0, 0.0)
	})

	It("BaseVectorY", func() {
		assertVec3Equals(BaseVectorY(), 0.0, 1.0, 0.0)
	})

	It("BaseVectorZ", func() {
		assertVec3Equals(BaseVectorZ(), 0.0, 0.0, 1.0)
	})

	It("DotProduct", func() {
		dot := DotProduct(firstVector, secondVector)
		assertFloatEquals(dot, -8.0)
	})

	It("CrossProduct", func() {
		cross := CrossProduct(firstVector, secondVector)
		assertVec3Equals(cross, -17.0, 2.0, 7.0)
	})
})
