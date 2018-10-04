package math_test

import (
	. "github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vec4", func() {
	var nullVector Vec4
	var preconfiguredVector Vec4

	BeforeEach(func() {
		nullVector = Vec4{}
		preconfiguredVector = Vec4{
			X: 1.0,
			Y: 2.0,
			Z: 3.0,
			W: 4.0,
		}
	})

	It("#Null", func() {
		Ω(nullVector.Null()).Should(BeTrue())
		Ω(preconfiguredVector.Null()).Should(BeFalse())
	})

	It("#Mul", func() {
		result := preconfiguredVector.Mul(1.5)
		Ω(result).Should(HaveVec4Coords(1.5, 3.0, 4.5, 6.0))
	})

	It("#Div", func() {
		result := preconfiguredVector.Div(2.0)
		Ω(result).Should(HaveVec4Coords(0.5, 1.0, 1.5, 2.0))
	})

	It("NullVec4", func() {
		Ω(NullVec4()).Should(HaveVec4Coords(0.0, 0.0, 0.0, 0.0))
	})

	It("MakeVec4", func() {
		vector := MakeVec4(1.1, 2.2, 3.3, 4.4)
		Ω(vector).Should(HaveVec4Coords(1.1, 2.2, 3.3, 4.4))
	})
})
