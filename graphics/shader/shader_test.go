package shader_test

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/client/client_fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/shader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shader", func() {
	const sourceCode = "#version 100 ..."
	var shaderId client.ShaderId
	var shaderCreateErr error
	var shaderDeleteErr error
	var shaderAllocationErr error
	var shaderSourceAssignmentErr error
	var shaderCompilationErr error
	var shaderDeletionErr error

	var shaderClient *client_fakes.FakeShaderClient
	var shader Shader
	var remoteShader RemoteShader

	BeforeEach(func() {
		shaderClient = new(client_fakes.FakeShaderClient)
		shaderId = 123
		shaderAllocationErr = errors.New("Could not allocate shader.")
		shaderSourceAssignmentErr = errors.New("Could not assign source code to shader.")
		shaderCompilationErr = errors.New("Could not compile shader.")
		shaderDeletionErr = errors.New("Could not delete shader.")
	})

	itShaderIsNotNil := func() {
		It("is not nil", func() {
			Ω(shader).ShouldNot(BeNil())
		})
	}

	itShaderHasSourceCodeGetter := func() {
		It("is possible to get source code", func() {
			Ω(shader.SourceCode()).Should(Equal(sourceCode))
		})
	}

	itRemoteShaderIsNotNil := func() {
		It("is not nil", func() {
			Ω(remoteShader).ShouldNot(BeNil())
		})
	}

	itRemoteShaderHasNoId := func() {
		It("has no Id set", func() {
			Ω(remoteShader.Id()).Should(BeNil())
		})
	}

	itRemoteShaderIsNotCreated := func() {
		It("is not created", func() {
			Ω(remoteShader.Created()).Should(BeFalse())
		})
	}

	itRemoteShaderIsCreated := func() {
		It("is created remotely", func() {
			Ω(remoteShader.Created()).Should(BeTrue())
		})
	}

	itRemoteShaderHasCorrectId := func() {
		It("has the proper Id", func() {
			Ω(remoteShader.Id()).Should(Equal(shaderId))
		})
	}

	itRemoteShaderCreationDidNotError := func() {
		It("did not error on creation", func() {
			Ω(shaderCreateErr).ShouldNot(HaveOccurred())
		})
	}

	itRemoteShaderCreationErrored := func(expectedErr *error) {
		It("did error on creation", func() {
			Ω(shaderCreateErr).Should(Equal(*expectedErr))
		})
	}

	itRemoteShaderHasMadeTheProperDeletionCalls := func() {
		It("has made the preper calls to the client", func() {
			Ω(shaderClient.DeleteShaderCallCount()).Should(Equal(1))
			argShaderId := shaderClient.DeleteShaderArgsForCall(0)
			Ω(argShaderId).Should(Equal(shaderId))
		})
	}

	itRemoteShaderDeletionDidNotError := func() {
		It("did not error on deletion", func() {
			Ω(shaderDeleteErr).ShouldNot(HaveOccurred())
		})
	}

	itRemoteShaderDeletionErrored := func(expectedErr *error) {
		It("did error on deletion", func() {
			Ω(shaderDeleteErr).Should(Equal(*expectedErr))
		})
	}

	Describe("VertexShader", func() {
		BeforeEach(func() {
			shader = NewVertexShader(shaderClient, sourceCode)
		})

		itShaderIsNotNil()

		itShaderHasSourceCodeGetter()

		Describe("RemoteVertexShader", func() {
			BeforeEach(func() {
				remoteShader = shader.Remote()
			})

			itRemoteShaderIsNotNil()

			itRemoteShaderHasNoId()

			itRemoteShaderIsNotCreated()

			Describe("Creation", func() {
				BeforeEach(func() {
					shaderClient.CreateVertexShaderReturns(shaderId, nil)
					shaderClient.CompileShaderReturns(nil)
				})

				JustBeforeEach(func() {
					shaderCreateErr = remoteShader.Create()
				})

				Context("when no client returns no errors", func() {
					itRemoteShaderIsCreated()

					itRemoteShaderHasCorrectId()

					itRemoteShaderCreationDidNotError()

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
						shaderClient.CreateVertexShaderReturns(nil, shaderAllocationErr)
					})

					itRemoteShaderHasNoId()

					itRemoteShaderIsNotCreated()

					itRemoteShaderCreationErrored(&shaderAllocationErr)
				})

				Context("when shader source could not be assigned", func() {
					BeforeEach(func() {
						shaderClient.SetShaderSourceCodeReturns(shaderSourceAssignmentErr)
					})

					itRemoteShaderCreationErrored(&shaderSourceAssignmentErr)
				})

				Context("when shader could not be compiled", func() {
					BeforeEach(func() {
						shaderClient.CompileShaderReturns(shaderCompilationErr)
					})

					itRemoteShaderCreationErrored(&shaderCompilationErr)
				})

				Describe("Remote Deletion", func() {
					JustBeforeEach(func() {
						shaderDeleteErr = remoteShader.Delete()
					})

					itRemoteShaderHasNoId()

					itRemoteShaderIsNotCreated()

					itRemoteShaderHasMadeTheProperDeletionCalls()

					Context("when shader could not be deleted", func() {
						BeforeEach(func() {
							shaderClient.DeleteShaderReturns(shaderDeletionErr)
						})

						itRemoteShaderDeletionErrored(&shaderDeleteErr)
					})
				})
			})
		})
	})

	Describe("FragmentShader", func() {
		BeforeEach(func() {
			shader = NewFragmentShader(shaderClient, sourceCode)
		})

		itShaderIsNotNil()

		itShaderHasSourceCodeGetter()

		Describe("RemoteFragmentShader", func() {
			BeforeEach(func() {
				remoteShader = shader.Remote()
			})

			itRemoteShaderIsNotNil()

			itRemoteShaderHasNoId()

			itRemoteShaderIsNotCreated()

			Describe("Creation", func() {
				BeforeEach(func() {
					shaderClient.CreateFragmentShaderReturns(shaderId, nil)
					shaderClient.CompileShaderReturns(nil)
				})

				JustBeforeEach(func() {
					shaderCreateErr = remoteShader.Create()
				})

				Context("when no client returns no errors", func() {
					itRemoteShaderIsCreated()

					itRemoteShaderHasCorrectId()

					itRemoteShaderCreationDidNotError()

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
						shaderClient.CreateFragmentShaderReturns(nil, shaderAllocationErr)
					})

					itRemoteShaderHasNoId()

					itRemoteShaderIsNotCreated()

					itRemoteShaderCreationErrored(&shaderAllocationErr)
				})

				Context("when shader source could not be assigned", func() {
					BeforeEach(func() {
						shaderClient.SetShaderSourceCodeReturns(shaderSourceAssignmentErr)
					})

					itRemoteShaderCreationErrored(&shaderSourceAssignmentErr)
				})

				Context("when shader could not be compiled", func() {
					BeforeEach(func() {
						shaderClient.CompileShaderReturns(shaderCompilationErr)
					})

					itRemoteShaderCreationErrored(&shaderCompilationErr)
				})

				Describe("Deletion", func() {
					JustBeforeEach(func() {
						shaderDeleteErr = remoteShader.Delete()
					})

					itRemoteShaderHasNoId()

					itRemoteShaderIsNotCreated()

					itRemoteShaderDeletionDidNotError()

					itRemoteShaderHasMadeTheProperDeletionCalls()

					Context("when shader could not be deleted", func() {
						BeforeEach(func() {
							shaderClient.DeleteShaderReturns(shaderDeletionErr)
						})

						itRemoteShaderDeletionErrored(&shaderDeleteErr)
					})
				})
			})
		})
	})
})
