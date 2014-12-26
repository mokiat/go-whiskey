package buffer_test

import (
	"errors"

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

	itHasInvalidId := func() {
		It("should have an invalid ID", func() {
			Ω(vertexBuffer.Id()).Should(Equal(graphics.InvalidBufferId))
		})
	}

	itIsNotCreatedRemotely := func() {
		It("is not created remotely initally", func() {
			Ω(vertexBuffer.CreatedRemotely()).Should(BeFalse())
		})
	}

	BeforeEach(func() {
		facade = new(FakeFacade)
		usage = graphics.StaticDraw
		size = 3
		vertexBuffer = NewVertexBuffer(facade, size, usage)
		vertexBuffer.SetValue(0, 0.01)
		vertexBuffer.SetValue(1, 0.11)
		vertexBuffer.SetValue(2, 0.21)
	})

	itHasInvalidId()

	It("should be possible to get the usage", func() {
		Ω(vertexBuffer.Usage()).Should(Equal(usage))
	})

	It("should be possible to get the size", func() {
		Ω(vertexBuffer.Size()).Should(Equal(size))
	})

	itIsNotCreatedRemotely()

	It("is possible to get values", func() {
		Ω(vertexBuffer.Value(0)).Should(Equal(float32(0.01)))
		Ω(vertexBuffer.Value(1)).Should(Equal(float32(0.11)))
		Ω(vertexBuffer.Value(2)).Should(Equal(float32(0.21)))
	})

	Describe("Remote Creation", func() {
		var bufferId graphics.ResourceId
		var createErr error

		BeforeEach(func() {
			bufferId = 876
			facade.CreateBufferReturns(bufferId, nil)
		})

		JustBeforeEach(func() {
			createErr = vertexBuffer.CreateRemotely()
		})

		Context("happy path", func() {
			It("does not return an error", func() {
				Ω(createErr).ShouldNot(HaveOccurred())
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
				Ω(argData.Length()).Should(Equal(3))
				Ω(argData.Get(0)).Should(Equal(float32(0.01)))
				Ω(argData.Get(1)).Should(Equal(float32(0.11)))
				Ω(argData.Get(2)).Should(Equal(float32(0.21)))
				Ω(argUsage).Should(Equal(usage))
			})
		})

		Context("when allocation fails", func() {
			var allocErr error

			BeforeEach(func() {
				allocErr = errors.New("Failed to allocate buffer")
				facade.CreateBufferReturns(bufferId, allocErr)
			})

			itHasInvalidId()

			itIsNotCreatedRemotely()

			It("returns an error", func() {
				Ω(createErr).Should(Equal(allocErr))
			})
		})

		Describe("Remote Binding", func() {
			JustBeforeEach(func() {
				vertexBuffer.BindRemotely()
			})

			It("should have made the proper calls to the facade", func() {
				Ω(facade.BindVertexBufferCallCount()).Should(Equal(2))
				argBufferId := facade.BindVertexBufferArgsForCall(1)
				Ω(argBufferId).Should(Equal(bufferId))
			})
		})

		Describe("Remote Deletion", func() {
			JustBeforeEach(func() {
				vertexBuffer.DeleteRemotely()
			})

			itHasInvalidId()

			itIsNotCreatedRemotely()

			It("should have made the proper calls to the facade", func() {
				Ω(facade.DeleteBufferCallCount()).Should(Equal(1))
				argBufferId := facade.DeleteBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(bufferId))
			})
		})
	})
})
