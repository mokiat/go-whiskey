package uniform_test

import (
	. "github.com/momchil-atanasov/go-whiskey/graphics/uniform"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Uniform", func() {

	It("uniform count is not zero", func() {
		Ω(UNIFORM_COUNT).Should(BeNumerically(">", 0))
	})

	It("each uniform has a name", func() {
		for i := 0; i < UNIFORM_COUNT; i++ {
			Ω(Uniform(i).Name()).ShouldNot(BeEmpty())
		}
	})

})
