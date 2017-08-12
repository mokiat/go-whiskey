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
		firstVector = Vec4{1.0, 2.0, 3.0, 4.0}
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

})
