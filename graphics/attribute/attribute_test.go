package attribute_test

import (
	. "github.com/momchil-atanasov/go-whiskey/graphics/attribute"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Attribute", func() {

	It("attribute count is not zero", func() {
		Ω(ATTRIBUTE_COUNT).Should(BeNumerically(">", 0))
	})

	It("each attribute has a name", func() {
		for i := 0; i < ATTRIBUTE_COUNT; i++ {
			Ω(Attribute(i).Name()).ShouldNot(BeEmpty())
		}
	})

	It("each attribute has non-zero dimensions", func() {
		for i := 0; i < ATTRIBUTE_COUNT; i++ {
			Ω(Attribute(i).Dimensions()).Should(BeNumerically(">", 0))
		}
	})
})
