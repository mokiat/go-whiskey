package math_test

import (
	. "github.com/momchil-atanasov/go-whiskey/math"
	. "github.com/momchil-atanasov/go-whiskey/math/test_helpers"

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

	It("NullVec4", func() {
		AssertVec4Equals(NullVec4(), 0.0, 0.0, 0.0, 0.0)
	})

})
