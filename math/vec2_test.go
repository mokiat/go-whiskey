package math_test

import (
	. "github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vec2", func() {

	var nullVector Vec2
	var firstVector Vec2
	var secondVector Vec2

	BeforeEach(func() {
		nullVector = Vec2{}
		firstVector = Vec2{
			X: 2.0,
			Y: 3.0,
		}
		secondVector = Vec2{
			X: -1.0,
			Y: 2.0,
		}
	})

	It("#Null", func() {
		Ω(nullVector.Null()).Should(BeTrue())
		Ω(firstVector.Null()).Should(BeFalse())
	})

	It("#Inverse", func() {
		inverted := firstVector.Inverse()
		AssertVec2Equals(inverted, -2.0, -3.0)
	})

	It("#IncCoords", func() {
		incremented := firstVector.IncCoords(1.5, -2.5)
		AssertVec2Equals(incremented, 3.5, 0.5)
	})

	It("#IncVec2", func() {
		incremented := firstVector.IncVec2(MakeVec2(1.5, -2.5))
		AssertVec2Equals(incremented, 3.5, 0.5)
	})

	It("#DecCoords", func() {
		decremented := firstVector.DecCoords(1.5, -2.5)
		AssertVec2Equals(decremented, 0.5, 5.5)
	})

	It("#DecVec2", func() {
		decremented := firstVector.DecVec2(MakeVec2(1.5, -2.5))
		AssertVec2Equals(decremented, 0.5, 5.5)
	})

	It("#Mul", func() {
		multiplied := firstVector.Mul(0.5)
		AssertVec2Equals(multiplied, 1.0, 1.5)
	})

	It("#Div", func() {
		divided := firstVector.Div(2.0)
		AssertVec2Equals(divided, 1.0, 1.5)
	})

	It("#LengthSquared", func() {
		squaredLength := firstVector.LengthSquared()
		AssertFloatEquals(squaredLength, 13.0)
	})

	It("#Length", func() {
		length := firstVector.Length()
		AssertFloatEquals(length, 3.605551275463989)
	})

	It("#Resize", func() {
		resized := firstVector.Resize(7.211102550927979)
		AssertVec2Equals(resized, 4.0, 6.0)
	})

	It("#DistanceToCoords", func() {
		distance := firstVector.DistanceToCoords(-1.0, 2.0)
		AssertFloatEquals(distance, 3.162277660168379)
	})

	It("#DistanceToVec2", func() {
		distance := firstVector.DistanceToVec2(secondVector)
		AssertFloatEquals(distance, 3.162277660168379)
	})

	It("NullVec2", func() {
		AssertVec2Equals(NullVec2(), 0.0, 0.0)
	})

	It("BaseVec2X", func() {
		AssertVec2Equals(BaseVec2X(), 1.0, 0.0)
	})

	It("BaseVec2Y", func() {
		AssertVec2Equals(BaseVec2Y(), 0.0, 1.0)
	})

	It("MakeVec2", func() {
		AssertVec2Equals(MakeVec2(1.3, 4.5), 1.3, 4.5)
	})

	It("Vec2DotProduct", func() {
		dot := Vec2DotProduct(firstVector, secondVector)
		AssertFloatEquals(dot, 4.0)
	})
})
