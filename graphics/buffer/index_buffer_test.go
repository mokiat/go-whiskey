package buffer_test

import (
	"github.com/momchil-atanasov/go-whiskey/graphics"
	. "github.com/momchil-atanasov/go-whiskey/graphics/buffer"
	. "github.com/momchil-atanasov/go-whiskey/graphics/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IndexBuffer", func() {

	var facade *FakeFacade
	var usage graphics.BufferUsage
	var size int
	var indexBuffer IndexBuffer

	BeforeEach(func() {
		facade = new(FakeFacade)
		usage = graphics.StaticDraw
		size = 4
		indexBuffer = NewIndexBuffer(facade, usage, size)
		indexBuffer.SetValue(0, 10)
		indexBuffer.SetValue(1, 11)
		indexBuffer.SetValue(2, 21)
		indexBuffer.SetValue(3, 31)
	})

	It("should have an invalid ID initially", func() {
		Ω(indexBuffer.Id()).Should(Equal(graphics.InvalidBufferId))
	})

	It("should be possible to get the usage", func() {
		Ω(indexBuffer.Usage()).Should(Equal(usage))
	})

	It("should be possible to get the size", func() {
		Ω(indexBuffer.Size()).Should(Equal(size))
	})

	It("is not created remotely initally", func() {
		Ω(indexBuffer.CreatedRemotely()).Should(BeFalse())
	})

	It("is possible to get values", func() {
		Ω(indexBuffer.Value(0)).Should(Equal(uint16(10)))
		Ω(indexBuffer.Value(1)).Should(Equal(uint16(11)))
		Ω(indexBuffer.Value(2)).Should(Equal(uint16(21)))
		Ω(indexBuffer.Value(3)).Should(Equal(uint16(31)))
	})

	Context("when the buffer is created remotely", func() {
		var bufferId graphics.ResourceId

		BeforeEach(func() {
			bufferId = 876
			facade.CreateBufferReturns(bufferId)

			indexBuffer.CreateRemotely()
		})

		It("should be created remotelly", func() {
			Ω(indexBuffer.CreatedRemotely()).Should(BeTrue())
		})

		It("should have the proper ID", func() {
			Ω(indexBuffer.Id()).Should(Equal(bufferId))
		})

		It("should have made the correct calls to the facade", func() {
			Ω(facade.CreateBufferCallCount()).Should(Equal(1))
			Ω(facade.BindIndexBufferCallCount()).Should(Equal(1))
			argBufferId := facade.BindIndexBufferArgsForCall(0)
			Ω(argBufferId).Should(Equal(bufferId))
			Ω(facade.CreateIndexBufferDataCallCount()).Should(Equal(1))
			argData, argUsage := facade.CreateIndexBufferDataArgsForCall(0)
			Ω(argData.Length()).Should(Equal(4))
			Ω(argData.Get(0)).Should(Equal(uint16(10)))
			Ω(argData.Get(1)).Should(Equal(uint16(11)))
			Ω(argData.Get(2)).Should(Equal(uint16(21)))
			Ω(argData.Get(3)).Should(Equal(uint16(31)))
			Ω(argUsage).Should(Equal(usage))
		})

		Context("when the buffer is bound", func() {
			BeforeEach(func() {
				indexBuffer.BindRemotely()
			})

			It("should have made the proper calls to the facade", func() {
				Ω(facade.BindIndexBufferCallCount()).Should(Equal(2))
				argBufferId := facade.BindIndexBufferArgsForCall(1)
				Ω(argBufferId).Should(Equal(bufferId))
			})
		})

		Context("when the buffer is deleted remotely", func() {
			BeforeEach(func() {
				indexBuffer.DeleteRemotely()
			})

			It("should not be created remotely", func() {
				Ω(indexBuffer.CreatedRemotely()).Should(BeFalse())
			})

			It("should have an invalid ID", func() {
				Ω(indexBuffer.Id()).Should(Equal(graphics.InvalidBufferId))
			})

			It("should have made the proper calls to the facade", func() {
				Ω(facade.DeleteBufferCallCount()).Should(Equal(1))
				argBufferId := facade.DeleteBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(bufferId))
			})
		})
	})

})
