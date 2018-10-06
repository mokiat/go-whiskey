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
		Ω(inverted).Should(HaveVec2Coords(-2.0, -3.0))
	})

	It("#IncCoords", func() {
		incremented := firstVector.IncCoords(1.5, -2.5)
		Ω(incremented).Should(HaveVec2Coords(3.5, 0.5))
	})

	It("#IncVec2", func() {
		incremented := firstVector.IncVec2(MakeVec2(1.5, -2.5))
		Ω(incremented).Should(HaveVec2Coords(3.5, 0.5))
	})

	It("#DecCoords", func() {
		decremented := firstVector.DecCoords(1.5, -2.5)
		Ω(decremented).Should(HaveVec2Coords(0.5, 5.5))
	})

	It("#DecVec2", func() {
		decremented := firstVector.DecVec2(MakeVec2(1.5, -2.5))
		Ω(decremented).Should(HaveVec2Coords(0.5, 5.5))
	})

	It("#Mul", func() {
		multiplied := firstVector.Mul(0.5)
		Ω(multiplied).Should(HaveVec2Coords(1.0, 1.5))
	})

	It("#Div", func() {
		divided := firstVector.Div(2.0)
		Ω(divided).Should(HaveVec2Coords(1.0, 1.5))
	})

	It("#LengthSquared", func() {
		squaredLength := firstVector.LengthSquared()
		Ω(squaredLength).Should(EqualFloat32(13.0))
	})

	It("#Length", func() {
		length := firstVector.Length()
		Ω(length).Should(EqualFloat32(3.605551275463989))
	})

	It("#Resize", func() {
		resized := firstVector.Resize(7.211102550927979)
		Ω(resized).Should(HaveVec2Coords(4.0, 6.0))
	})

	It("#DistanceToCoords", func() {
		distance := firstVector.DistanceToCoords(-1.0, 2.0)
		Ω(distance).Should(EqualFloat32(3.162277660168379))
	})

	It("#DistanceToVec2", func() {
		distance := firstVector.DistanceToVec2(secondVector)
		Ω(distance).Should(EqualFloat32(3.162277660168379))
	})

	It("NullVec2", func() {
		Ω(NullVec2()).Should(HaveVec2Coords(0.0, 0.0))
	})

	It("BaseVec2X", func() {
		Ω(BaseVec2X()).Should(HaveVec2Coords(1.0, 0.0))
	})

	It("BaseVec2Y", func() {
		Ω(BaseVec2Y()).Should(HaveVec2Coords(0.0, 1.0))
	})

	It("MakeVec2", func() {
		Ω(MakeVec2(1.3, 4.5)).Should(HaveVec2Coords(1.3, 4.5))
	})

	It("Vec2DotProduct", func() {
		dot := Vec2DotProduct(firstVector, secondVector)
		Ω(dot).Should(EqualFloat32(4.0))
	})
})
