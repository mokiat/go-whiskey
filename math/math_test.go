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

	It("Sqrt32", func() {
		AssertFloatEquals(Sqrt32(16.0), 4.0)
	})

	It("Signum32", func() {
		AssertFloatEquals(Signum32(0.1), 1.0)
		AssertFloatEquals(Signum32(-0.1), -1.0)
	})

	It("Atan32", func() {
		AssertFloatEquals(Atan32(2.0), 1.10714872)
	})

})
