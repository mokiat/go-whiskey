package buffer_test

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/common/buf"
	. "github.com/momchil-atanasov/go-whiskey/graphics/buffer"
	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/client/client_fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VertexBufferData", func() {
	const size = 3
	const epsilon = 1e-4

	var data VertexBufferData

	BeforeEach(func() {
		data = NewVertexBufferData(size)
	})

	It("should not be nil", func() {
		Ω(data).ShouldNot(BeNil())
	})

	It("should have the correct size", func() {
		Ω(data.Size()).Should(Equal(size))
	})

	It("all values should equal 0.0", func() {
		Ω(data.Value(0)).Should(BeNumerically("~", 0.0, epsilon))
		Ω(data.Value(1)).Should(BeNumerically("~", 0.0, epsilon))
		Ω(data.Value(2)).Should(BeNumerically("~", 0.0, epsilon))
	})

	Context("when values are set", func() {
		BeforeEach(func() {
			data.SetValue(0, 3.01)
			data.SetValue(1, 2.11)
			data.SetValue(2, 1.21)
		})

		It("they are changed in the data", func() {
			Ω(data.Value(0)).Should(BeNumerically("~", 3.01, epsilon))
			Ω(data.Value(1)).Should(BeNumerically("~", 2.11, epsilon))
			Ω(data.Value(2)).Should(BeNumerically("~", 1.21, epsilon))
		})

		It("is possible to get the byte data", func() {
			content := data.Content()
			Ω(content).ShouldNot(BeNil())
			Ω(content).Should(HaveLen(3 * 4))
		})
	})
})

var _ = Describe("VertexBuffer", func() {
	const epsilon = 1e-4

	var bufferClient *client_fakes.FakeBufferClient
	var data VertexBufferData
	var usage client.BufferUsage
	var buffer VertexBuffer

	BeforeEach(func() {
		bufferClient = new(client_fakes.FakeBufferClient)
		data = NewVertexBufferData(2)
		data.SetValue(0, 2.0)
		data.SetValue(1, 1.0)
		usage = client.StaticDrawBufferUsage
		buffer = NewVertexBuffer(data, usage)
	})

	It("is not nil", func() {
		Ω(buffer).ShouldNot(BeNil())
	})

	It("is possible to get data", func() {
		Ω(buffer.Data()).Should(Equal(data))
	})

	It("is possible to get usage", func() {
		Ω(buffer.Usage()).Should(Equal(usage))
	})

	It("has nil Id by default", func() {
		Ω(buffer.Id()).Should(BeNil())
	})

	It("is not created by default", func() {
		Ω(buffer.Created()).Should(BeFalse())
	})

	Describe("Create", func() {
		var bufferId client.BufferId
		var createErr error
		var clientErr error

		BeforeEach(func() {
			bufferId = 876
			clientErr = errors.New("Client operation failed!")
			bufferClient.CreateBufferReturns(bufferId, nil)
		})

		JustBeforeEach(func() {
			createErr = buffer.Create(bufferClient)
		})

		Context("when client returns as expected", func() {
			It("it did not error on create", func() {
				Ω(createErr).ShouldNot(HaveOccurred())
			})

			It("is created", func() {
				Ω(buffer.Created()).Should(BeTrue())
			})

			It("has correct Id", func() {
				Ω(buffer.Id()).Should(Equal(bufferId))
			})

			It("made the correct calls to the client", func() {
				Ω(bufferClient.CreateBufferCallCount()).Should(Equal(1))
				Ω(bufferClient.BindVertexBufferCallCount()).Should(Equal(1))
				argBufferId := bufferClient.BindVertexBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(bufferId))
				Ω(bufferClient.CreateVertexBufferDataCallCount()).Should(Equal(1))
				argData, argUsage := bufferClient.CreateVertexBufferDataArgsForCall(0)
				argDataAsBuffer := buf.Float32Buffer(argData)
				Ω(argDataAsBuffer.Length()).Should(Equal(2))
				Ω(argDataAsBuffer.Get(0)).Should(BeNumerically("~", 2.0, epsilon))
				Ω(argDataAsBuffer.Get(1)).Should(BeNumerically("~", 1.0, epsilon))
				Ω(argUsage).Should(Equal(usage))
			})
		})

		Context("when allocation fails", func() {
			BeforeEach(func() {
				bufferClient.CreateBufferReturns(nil, clientErr)
			})

			It("errored on create", func() {
				Ω(createErr).Should(Equal(clientErr))
			})
		})

		Context("when binding fails", func() {
			BeforeEach(func() {
				bufferClient.BindVertexBufferReturns(clientErr)
			})

			It("errored on create", func() {
				Ω(createErr).Should(Equal(clientErr))
			})
		})

		Context("when data creation fails", func() {
			BeforeEach(func() {
				bufferClient.CreateVertexBufferDataReturns(clientErr)
			})

			It("errored on create", func() {
				Ω(createErr).Should(Equal(clientErr))
			})
		})

		Describe("Delete", func() {
			var deleteErr error

			JustBeforeEach(func() {
				deleteErr = buffer.Delete(bufferClient)
			})

			Context("when client returns as expected", func() {
				It("has nil Id", func() {
					Ω(buffer.Id()).Should(BeNil())
				})

				It("is no longer created", func() {
					Ω(buffer.Created()).Should(BeFalse())
				})

				It("made the proper calls to the client", func() {
					Ω(bufferClient.DeleteBufferCallCount()).Should(Equal(1))
					argBufferId := bufferClient.DeleteBufferArgsForCall(0)
					Ω(argBufferId).Should(Equal(bufferId))
				})
			})

			Context("when client errors on delete", func() {
				BeforeEach(func() {
					bufferClient.DeleteBufferReturns(clientErr)
				})

				It("errored on delete", func() {
					Ω(deleteErr).Should(Equal(clientErr))
				})
			})
		})
	})
})
