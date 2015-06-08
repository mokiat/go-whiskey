package render_test

import (
	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	. "github.com/momchil-atanasov/go-whiskey/graphics/render"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FIFOQueue", func() {
	var queue Queue
	var firstItem Item
	var secondItem Item
	var thirdItem Item

	BeforeEach(func() {
		queue = NewFIFOQueue()

		firstItem = Item{
			ProgramId: client.ProgramId(1),
		}

		secondItem = Item{
			ProgramId: client.ProgramId(2),
		}

		thirdItem = Item{
			ProgramId: client.ProgramId(3),
		}
	})

	It("should not be nil", func() {
		Ω(queue).ShouldNot(BeNil())
	})

	It("should be empty", func() {
		Ω(queue.Size()).Should(Equal(0))
		Ω(queue.Empty()).Should(BeTrue())
	})

	Context("when items are added to the queue", func() {
		BeforeEach(func() {
			queue.Push(firstItem)
			queue.Push(secondItem)
			queue.Push(thirdItem)
		})

		It("should have changed its size", func() {
			Ω(queue.Size()).Should(Equal(3))
		})

		It("should not be empty", func() {
			Ω(queue.Empty()).Should(BeFalse())
		})

		It("items are returned in the same order they were added", func() {
			item, contains := queue.Pop()
			Ω(contains).Should(BeTrue())
			Ω(item).Should(Equal(firstItem))

			item, contains = queue.Pop()
			Ω(contains).Should(BeTrue())
			Ω(item).Should(Equal(secondItem))

			item, contains = queue.Pop()
			Ω(contains).Should(BeTrue())
			Ω(item).Should(Equal(thirdItem))
		})

		Context("when all items are removed from the queue", func() {
			BeforeEach(func() {
				queue.Pop()
				queue.Pop()
				queue.Pop()
			})

			It("should be empty", func() {
				Ω(queue.Size()).Should(Equal(0))
				Ω(queue.Empty()).Should(BeTrue())
			})
		})
	})

	Context("when removing from an empty queue", func() {
		It("an item is not returned", func() {
			_, found := queue.Pop()
			Ω(found).Should(BeFalse())
		})
	})
})
