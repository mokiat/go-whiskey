package math_test

import (
	. "github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math", func() {
	It("Abs32", func() {
		Ω(Abs32(1.1)).Should(EqualFloat32(1.1))
		Ω(Abs32(-1.1)).Should(EqualFloat32(1.1))
	})

	It("Sin32", func() {
		Ω(Sin32(Pi / 6.0)).Should(EqualFloat32(0.5))
	})

	It("Cos32", func() {
		Ω(Cos32(Pi / 3.0)).Should(EqualFloat32(0.5))
	})

	It("Sqrt32", func() {
		Ω(Sqrt32(16.0)).Should(EqualFloat32(4.0))
	})

	It("Signum32", func() {
		Ω(Signum32(0.1)).Should(EqualFloat32(1.0))
		Ω(Signum32(-0.1)).Should(EqualFloat32(-1.0))
	})

	It("Atan32", func() {
		Ω(Atan32(2.0)).Should(EqualFloat32(1.10714872))
	})
})
