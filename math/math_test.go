package math_test

import (
	. "github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Math", func() {

	It("Abs32", func() {
		AssertFloatEquals(Abs32(1.1), 1.1)
		AssertFloatEquals(Abs32(-1.1), 1.1)
	})

	It("Sin32", func() {
		AssertFloatEquals(Sin32(Pi/6.0), 0.5)
	})

	It("Cos32", func() {
		AssertFloatEquals(Cos32(Pi/3.0), 0.5)
	})

})
