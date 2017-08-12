package stage_test

import (
	. "github.com/mokiat/go-whiskey/math/test_helpers"
	. "github.com/mokiat/go-whiskey/scene/stage"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func AssertNodeBounds(node Node, x, y, z, size float32) {
	bounds := node.Bounds()
	AssertFloatEquals(bounds.Width(), size)
	AssertFloatEquals(bounds.Height(), size)
	AssertFloatEquals(bounds.Depth(), size)
}

var _ = Describe("Node", func() {

	var node Node

	BeforeEach(func() {
		node = NewNode(32.0, 3)
	})

	It("should have depth of three", func() {
		secondDepth, found := node.Child(TopLeftFront)
		Ω(found).Should(BeTrue())
		Ω(secondDepth).ShouldNot(BeNil())

		thirdDepth, found := secondDepth.Child(TopLeftFront)
		Ω(found).Should(BeTrue())
		Ω(thirdDepth).ShouldNot(BeNil())

		_, found = thirdDepth.Child(TopLeftFront)
		Ω(found).Should(BeFalse())
		_, found = thirdDepth.Child(TopRightFront)
		Ω(found).Should(BeFalse())
		_, found = thirdDepth.Child(BottomLeftFront)
		Ω(found).Should(BeFalse())
		_, found = thirdDepth.Child(BottomRightFront)
		Ω(found).Should(BeFalse())
		_, found = thirdDepth.Child(TopLeftBack)
		Ω(found).Should(BeFalse())
		_, found = thirdDepth.Child(TopRightBack)
		Ω(found).Should(BeFalse())
		_, found = thirdDepth.Child(BottomLeftBack)
		Ω(found).Should(BeFalse())
		_, found = thirdDepth.Child(BottomRightBack)
		Ω(found).Should(BeFalse())
	})

	It("should have proper bounds at level 1", func() {
		AssertNodeBounds(node, 0.0, 0.0, 0.0, 32.0)
	})

	It("should have proper bounds at level 2", func() {
		childNode, _ := node.Child(TopLeftFront)
		AssertNodeBounds(childNode, -8.0, 8.0, 8.0, 16.0)
		childNode, _ = node.Child(TopRightFront)
		AssertNodeBounds(childNode, 8.0, 8.0, 8.0, 16.0)
		childNode, _ = node.Child(BottomLeftFront)
		AssertNodeBounds(childNode, -8.0, -8.0, 8.0, 16.0)
		childNode, _ = node.Child(BottomRightFront)
		AssertNodeBounds(childNode, 8.0, -8.0, 8.0, 16.0)
		childNode, _ = node.Child(TopLeftBack)
		AssertNodeBounds(childNode, -8.0, 8.0, -8.0, 16.0)
		childNode, _ = node.Child(TopRightBack)
		AssertNodeBounds(childNode, 8.0, 8.0, -8.0, 16.0)
		childNode, _ = node.Child(BottomLeftBack)
		AssertNodeBounds(childNode, -8.0, -8.0, -8.0, 16.0)
		childNode, _ = node.Child(BottomRightBack)
		AssertNodeBounds(childNode, 8.0, -8.0, -8.0, 16.0)
	})

})
