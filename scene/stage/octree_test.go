package stage_test

import (
	. "github.com/mokiat/go-whiskey/scene/stage"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Octree", func() {

	var octree Octree

	BeforeEach(func() {
		octree = NewOctree(64.0, 4)
	})

	It("should have 4 nodes in depth", func() {
		firstDepth := octree.Root()
		Ω(firstDepth).ShouldNot(BeNil())

		secondDepth, found := firstDepth.Child(TopLeftFront)
		Ω(found).Should(BeTrue())
		Ω(secondDepth).ShouldNot(BeNil())

		thirdDepth, found := secondDepth.Child(TopLeftFront)
		Ω(found).Should(BeTrue())
		Ω(thirdDepth).ShouldNot(BeNil())

		fourthDepth, found := thirdDepth.Child(TopLeftFront)
		Ω(found).Should(BeTrue())
		Ω(fourthDepth).ShouldNot(BeNil())

		_, found = fourthDepth.Child(TopLeftFront)
		Ω(found).Should(BeFalse())
	})
})
