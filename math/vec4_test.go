package math_test

import (
	. "github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vec4", func() {

	var nullVector Vec4
	var firstVector Vec4

	BeforeEach(func() {
		nullVector = Vec4{}
		firstVector = Vec4{
			X: 1.0,
			Y: 2.0,
			Z: 3.0,
			W: 4.0,
		}
	})

	It("#Null", func() {
		Ω(nullVector.Null()).Should(BeTrue())
		Ω(firstVector.Null()).Should(BeFalse())
	})

	It("#Mul", func() {
		result := firstVector.Mul(1.5)
		AssertVec4Equals(result, 1.5, 3.0, 4.5, 6.0)
	})

	It("#Div", func() {
		result := firstVector.Div(2.0)
		AssertVec4Equals(result, 0.5, 1.0, 1.5, 2.0)
	})

	It("NullVec4", func() {
		AssertVec4Equals(NullVec4(), 0.0, 0.0, 0.0, 0.0)
	})

	It("MakeVec4", func() {
		AssertVec4Equals(MakeVec4(1.1, 2.2, 3.3, 4.4), 1.1, 2.2, 3.3, 4.4)
	})
})
