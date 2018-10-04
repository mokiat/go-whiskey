package math_test

import (
	. "github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
			X: -1.0,
			Y: 2.0,
			Z: -3.0,
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

	It("#IncVec3", func() {
		incremented := firstVector.IncVec3(MakeVec3(1.5, -2.5, 5.0))
		AssertVec3Equals(incremented, 3.5, 0.5, 9.0)
	})

	It("#DecCoords", func() {
		decremented := firstVector.DecCoords(1.5, -2.5, 5.0)
		AssertVec3Equals(decremented, 0.5, 5.5, -1.0)
	})

	It("#DecVec3", func() {
		decremented := firstVector.DecVec3(MakeVec3(1.5, -2.5, 5.0))
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
		Ω(squaredLength).Should(EqualFloat32(29.0))
	})

	It("#Length", func() {
		length := firstVector.Length()
		Ω(length).Should(EqualFloat32(5.385164807134504))
	})

	It("#Resize", func() {
		resized := firstVector.Resize(10.77032961426901)
		AssertVec3Equals(resized, 4.0, 6.0, 8.0)
	})

	It("#DistanceToCoords", func() {
		distance := firstVector.DistanceToCoords(-1.0, 2.0, -3.0)
		Ω(distance).Should(EqualFloat32(7.681145747868608))
	})

	It("#DistanceToVec3", func() {
		distance := firstVector.DistanceToVec3(secondVector)
		Ω(distance).Should(EqualFloat32(7.681145747868608))
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

	It("MakeVec3", func() {
		AssertVec3Equals(MakeVec3(2.6, 8.1, 9.4), 2.6, 8.1, 9.4)
	})

	It("Vec3DotProduct", func() {
		dot := Vec3DotProduct(firstVector, secondVector)
		Ω(dot).Should(EqualFloat32(-8.0))
	})

	It("Vec3CrossProduct", func() {
		cross := Vec3CrossProduct(firstVector, secondVector)
		AssertVec3Equals(cross, -17.0, 2.0, 7.0)
	})
})
