package shader_test

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/client/client_fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/shader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShaderData", func() {
	var data ShaderData

	BeforeEach(func() {
		data = NewShaderData()
	})

	It("is not nil", func() {
		Ω(data).ShouldNot(BeNil())
	})

	It("has no source code by default", func() {
		Ω(data.SourceCode()).Should(BeEmpty())
	})

	It("is possible to set source code", func() {
		source := "#source"
		data.SetSourceCode(source)
		Ω(data.SourceCode()).Should(Equal(source))
	})
})

var _ = Describe("Shader", func() {
	const sourceCode = "#version 100 ..."

	var shaderClient *client_fakes.FakeShaderClient
	var shaderId client.ShaderId
	var createErr error
	var shaderDeleteErr error
	var clientErr error
	var data ShaderData
	var shader Shader

	BeforeEach(func() {
		shaderClient = new(client_fakes.FakeShaderClient)

		shaderId = 123

		clientErr = errors.New("Client operation failed!")

		data = NewShaderData()
		data.SetSourceCode(sourceCode)
	})

	itIsNotNil := func() {
		It("is not nil", func() {
			Ω(shader).ShouldNot(BeNil())
		})
	}

	itIsPossibleToGetData := func() {
		It("is possible to get data", func() {
			Ω(shader.Data()).Should(Equal(data))
		})
	}

	itHasNoId := func() {
		It("has no Id set", func() {
			Ω(shader.Id()).Should(BeNil())
		})
	}

	itIsNotCreated := func() {
		It("is not created", func() {
			Ω(shader.Created()).Should(BeFalse())
		})
	}

	itIsCreated := func() {
		It("is created", func() {
			Ω(shader.Created()).Should(BeTrue())
		})
	}

	itHasCorrectId := func() {
		It("has the proper Id", func() {
			Ω(shader.Id()).Should(Equal(shaderId))
		})
	}

	itCreateDidNotError := func() {
		It("did not error on creation", func() {
			Ω(createErr).ShouldNot(HaveOccurred())
		})
	}

	itCreateErroredWith := func(expectedErr *error) {
		It("did error on creation", func() {
			Ω(createErr).Should(Equal(*expectedErr))
		})
	}

	itMadeTheProperDeletionCalls := func() {
		It("made the preper calls to the client", func() {
			Ω(shaderClient.DeleteShaderCallCount()).Should(Equal(1))
			argShaderId := shaderClient.DeleteShaderArgsForCall(0)
			Ω(argShaderId).Should(Equal(shaderId))
		})
	}

	itDeleteDidNotError := func() {
		It("did not error on deletion", func() {
			Ω(shaderDeleteErr).ShouldNot(HaveOccurred())
		})
	}

	itDeleteErroredWith := func(expectedErr *error) {
		It("did error on deletion", func() {
			Ω(shaderDeleteErr).Should(Equal(*expectedErr))
		})
	}

	Describe("VertexShader", func() {
		BeforeEach(func() {
			shader = NewVertexShader(data)
		})

		itIsNotNil()

		itIsPossibleToGetData()

		itHasNoId()

		itIsNotCreated()

		Describe("Creation", func() {
			BeforeEach(func() {
				shaderClient.CreateVertexShaderReturns(shaderId, nil)
				shaderClient.CompileShaderReturns(nil)
			})

			JustBeforeEach(func() {
				createErr = shader.Create(shaderClient)
			})

			Context("when client returns no errors", func() {
				itIsCreated()

				itHasCorrectId()

				itCreateDidNotError()

				It("made the proper calls to the client", func() {
					Ω(shaderClient.CreateVertexShaderCallCount()).Should(Equal(1))
					Ω(shaderClient.SetShaderSourceCodeCallCount()).Should(Equal(1))
					argShaderId, argShaderSource := shaderClient.SetShaderSourceCodeArgsForCall(0)
					Ω(argShaderId).Should(Equal(shaderId))
					Ω(argShaderSource).Should(Equal(sourceCode))
					Ω(shaderClient.CompileShaderCallCount()).Should(Equal(1))
					argShaderId = shaderClient.CompileShaderArgsForCall(0)
					Ω(argShaderId).Should(Equal(shaderId))
				})
			})

			Context("when a shader could not be allocated", func() {
				BeforeEach(func() {
					shaderClient.CreateVertexShaderReturns(nil, clientErr)
				})

				itHasNoId()

				itIsNotCreated()

				itCreateErroredWith(&clientErr)
			})

			Context("when shader source could not be assigned", func() {
				BeforeEach(func() {
					shaderClient.SetShaderSourceCodeReturns(clientErr)
				})

				itCreateErroredWith(&clientErr)
			})

			Context("when shader could not be compiled", func() {
				BeforeEach(func() {
					shaderClient.CompileShaderReturns(clientErr)
				})

				itCreateErroredWith(&clientErr)
			})

			Describe("Deletion", func() {
				JustBeforeEach(func() {
					shaderDeleteErr = shader.Delete(shaderClient)
				})

				itHasNoId()

				itIsNotCreated()

				itDeleteDidNotError()

				itMadeTheProperDeletionCalls()

				Context("when shader could not be deleted", func() {
					BeforeEach(func() {
						shaderClient.DeleteShaderReturns(clientErr)
					})

					itDeleteErroredWith(&clientErr)
				})
			})
		})
	})

	Describe("FragmentShader", func() {
		BeforeEach(func() {
			shader = NewFragmentShader(data)
		})

		itIsNotNil()

		itIsPossibleToGetData()

		itHasNoId()

		itIsNotCreated()

		Describe("Creation", func() {
			BeforeEach(func() {
				shaderClient.CreateFragmentShaderReturns(shaderId, nil)
				shaderClient.CompileShaderReturns(nil)
			})

			JustBeforeEach(func() {
				createErr = shader.Create(shaderClient)
			})

			Context("when client returns no errors", func() {
				itIsCreated()

				itHasCorrectId()

				itCreateDidNotError()

				It("made the proper calls to the client", func() {
					Ω(shaderClient.CreateFragmentShaderCallCount()).Should(Equal(1))
					Ω(shaderClient.SetShaderSourceCodeCallCount()).Should(Equal(1))
					argShaderId, argShaderSource := shaderClient.SetShaderSourceCodeArgsForCall(0)
					Ω(argShaderId).Should(Equal(shaderId))
					Ω(argShaderSource).Should(Equal(sourceCode))
					Ω(shaderClient.CompileShaderCallCount()).Should(Equal(1))
					argShaderId = shaderClient.CompileShaderArgsForCall(0)
					Ω(argShaderId).Should(Equal(shaderId))
				})
			})

			Context("when a shader could not be allocated", func() {
				BeforeEach(func() {
					shaderClient.CreateFragmentShaderReturns(nil, clientErr)
				})

				itHasNoId()

				itIsNotCreated()

				itCreateErroredWith(&clientErr)
			})

			Context("when shader source could not be assigned", func() {
				BeforeEach(func() {
					shaderClient.SetShaderSourceCodeReturns(clientErr)
				})

				itCreateErroredWith(&clientErr)
			})

			Context("when shader could not be compiled", func() {
				BeforeEach(func() {
					shaderClient.CompileShaderReturns(clientErr)
				})

				itCreateErroredWith(&clientErr)
			})

			Describe("Deletion", func() {
				JustBeforeEach(func() {
					shaderDeleteErr = shader.Delete(shaderClient)
				})

				itHasNoId()

				itIsNotCreated()

				itDeleteDidNotError()

				itMadeTheProperDeletionCalls()

				Context("when shader could not be deleted", func() {
					BeforeEach(func() {
						shaderClient.DeleteShaderReturns(clientErr)
					})

					itDeleteErroredWith(&clientErr)
				})
			})
		})
	})
})
