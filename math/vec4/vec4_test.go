package vec4_test

import (
	. "github.com/momchil-atanasov/go-whiskey/math/vec4"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vec4", func() {

	assertFloatEquals := func(actualValue, expectedValue float32) {
		Ω(actualValue).Should(BeNumerically("~", expectedValue))
	}

	assertVec4Equals := func(vector Vec4, expectedX, expectedY, expectedZ, expectedW float32) {
		assertFloatEquals(vector.X, expectedX)
		assertFloatEquals(vector.Y, expectedY)
		assertFloatEquals(vector.Z, expectedZ)
		assertFloatEquals(vector.W, expectedW)
	}

	It("Null", func() {
		assertVec4Equals(Null(), 0.0, 0.0, 0.0, 0.0)
	})

})
