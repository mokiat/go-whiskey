package common_test

import (
	. "github.com/mokiat/go-whiskey/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type CustomEnum Enum

const (
	First CustomEnum = iota
	Second
	Third
	CUSTOM_ENUM_SIZE int = iota
)

var _ = Describe("Enumset", func() {

	var set EnumSet

	itBehavesLikeAnEmptySet := func() {
		It("should have size of zero", func() {
			Ω(set.Size()).Should(Equal(0))
		})

		It("should be empty", func() {
			Ω(set.Empty()).Should(BeTrue())
		})

		It("should not contain any of the enums", func() {
			Ω(set.Contains(Enum(First))).Should(BeFalse())
			Ω(set.Contains(Enum(Second))).Should(BeFalse())
			Ω(set.Contains(Enum(Third))).Should(BeFalse())
		})
	}

	BeforeEach(func() {
		set = NewEnumSet(CUSTOM_ENUM_SIZE)
	})

	itBehavesLikeAnEmptySet()

	Context("when enums are added", func() {
		BeforeEach(func() {
			set.Add(Enum(First))
			set.Add(Enum(Third))
		})

		It("should not be empty", func() {
			Ω(set.Empty()).Should(BeFalse())
		})

		It("should have changed its size", func() {
			Ω(set.Size()).Should(Equal(2))
		})

		It("should contain the added enums", func() {
			Ω(set.Contains(Enum(First))).Should(BeTrue())
			Ω(set.Contains(Enum(Third))).Should(BeTrue())
		})

		It("should still not contain the unadded enums", func() {
			Ω(set.Contains(Enum(Second))).Should(BeFalse())
		})

		It("is possible to iterate over contained enums", func() {
			iterated := []CustomEnum{}
			set.Each(func(enum Enum) {
				iterated = append(iterated, CustomEnum(enum))
			})
			Ω(iterated).Should(HaveLen(2))
			Ω(iterated[0]).Should(Equal(First))
			Ω(iterated[1]).Should(Equal(Third))
		})

		Context("when part of the enums are removed", func() {
			BeforeEach(func() {
				set.Remove(Enum(First))
			})

			It("should not be empty", func() {
				Ω(set.Empty()).Should(BeFalse())
			})

			It("should have changed its size", func() {
				Ω(set.Size()).Should(Equal(1))
			})

			It("should not contain the removed enums", func() {
				Ω(set.Contains(Enum(First))).Should(BeFalse())
			})

			It("should still contain the unremoved enums", func() {
				Ω(set.Contains(Enum(Third))).Should(BeTrue())
			})
		})

		Context("when the set is cleared", func() {
			BeforeEach(func() {
				set.Clear()
			})

			itBehavesLikeAnEmptySet()
		})

	})
})
