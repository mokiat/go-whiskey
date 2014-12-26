package buffer_test

import (
	"github.com/momchil-atanasov/go-whiskey/graphics"
	. "github.com/momchil-atanasov/go-whiskey/graphics/buffer"
	. "github.com/momchil-atanasov/go-whiskey/graphics/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VertexBuffer", func() {

	var facade *FakeFacade
	var usage graphics.BufferUsage
	var size int
	var vertexBuffer VertexBuffer

	BeforeEach(func() {
		facade = new(FakeFacade)
		usage = graphics.StaticDraw
		size = 3
		vertexBuffer = NewVertexBuffer(facade, size, usage)
	})

	It("should have an invalid ID initially", func() {
		Ω(vertexBuffer.Id()).Should(Equal(graphics.InvalidBufferId))
	})

	It("should be possible to get the usage", func() {
		Ω(vertexBuffer.Usage()).Should(Equal(usage))
	})

	It("should be possible to get the size", func() {
		Ω(vertexBuffer.Size()).Should(Equal(size))
	})

	It("is not created remotely initally", func() {
		Ω(vertexBuffer.CreatedRemotely()).Should(BeFalse())
	})

	Context("when the buffer is created remotely", func() {
		var bufferId graphics.ResourceId

		BeforeEach(func() {
			bufferId = 876
			facade.CreateBufferReturns(bufferId)

			vertexBuffer.CreateRemotely()
		})

		It("should be created remotelly", func() {
			Ω(vertexBuffer.CreatedRemotely()).Should(BeTrue())
		})

		It("should have the proper ID", func() {
			Ω(vertexBuffer.Id()).Should(Equal(bufferId))
		})

		It("should have made the correct calls to the facade", func() {
			Ω(facade.CreateBufferCallCount()).Should(Equal(1))
			Ω(facade.BindVertexBufferCallCount()).Should(Equal(1))
			argBufferId := facade.BindVertexBufferArgsForCall(0)
			Ω(argBufferId).Should(Equal(bufferId))
			Ω(facade.CreateVertexBufferDataCallCount()).Should(Equal(1))
			argData, argUsage := facade.CreateVertexBufferDataArgsForCall(0)
			Ω(argData).Should(Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
			Ω(argUsage).Should(Equal(usage))
		})

		Context("when the buffer is bound", func() {
			BeforeEach(func() {
				vertexBuffer.BindRemotely()
			})

			It("should have made the proper calls to the facade", func() {
				Ω(facade.BindVertexBufferCallCount()).Should(Equal(2))
				argBufferId := facade.BindVertexBufferArgsForCall(1)
				Ω(argBufferId).Should(Equal(bufferId))
			})
		})

		Context("when the buffer is deleted remotely", func() {
			BeforeEach(func() {
				vertexBuffer.DeleteRemotely()
			})

			It("should not be created remotely", func() {
				Ω(vertexBuffer.CreatedRemotely()).Should(BeFalse())
			})

			It("should have an invalid ID", func() {
				Ω(vertexBuffer.Id()).Should(Equal(graphics.InvalidBufferId))
			})

			It("should have made the proper calls to the facade", func() {
				Ω(facade.DeleteBufferCallCount()).Should(Equal(1))
				argBufferId := facade.DeleteBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(bufferId))
			})
		})
	})
})
