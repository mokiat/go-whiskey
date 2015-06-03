package shader_test

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/client/client_fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/shader"
	"github.com/momchil-atanasov/go-whiskey/graphics/shader/shader_fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Program", func() {
	var shaderClient *client_fakes.FakeShaderClient
	var vertexShaderId client.ShaderId
	var vertexShader *shader_fakes.FakeRemoteShader
	var fragmentShaderId client.ShaderId
	var fragmentShader *shader_fakes.FakeRemoteShader
	var program Program

	BeforeEach(func() {
		shaderClient = new(client_fakes.FakeShaderClient)
		vertexShader = new(shader_fakes.FakeRemoteShader)
		vertexShaderId = 8989
		vertexShader.IdReturns(vertexShaderId)
		fragmentShader = new(shader_fakes.FakeRemoteShader)
		fragmentShaderId = 6767
		fragmentShader.IdReturns(fragmentShaderId)
		program = NewProgram(shaderClient, vertexShader, fragmentShader)
	})

	It("should not be nil", func() {
		Ω(program).ShouldNot(BeNil())
	})

	It("is possible to get vertex shader", func() {
		Ω(program.VertexShader()).Should(Equal(vertexShader))
	})

	It("is possible to get fragment shader", func() {
		Ω(program.FragmentShader()).Should(Equal(fragmentShader))
	})

	Describe("RemoteProgram", func() {
		var remoteProgram RemoteProgram

		BeforeEach(func() {
			remoteProgram = program.Remote()
		})

		itRemoteProgramHasNoId := func() {
			It("has nil Id", func() {
				Ω(remoteProgram.Id()).Should(BeNil())
			})
		}

		itRemoteProgramIsNotCreated := func() {
			It("is not created", func() {
				Ω(remoteProgram.Created()).Should(BeFalse())
			})
		}

		It("is not nil", func() {
			Ω(remoteProgram).ShouldNot(BeNil())
		})

		itRemoteProgramHasNoId()

		itRemoteProgramIsNotCreated()

		Describe("Creation", func() {
			var programId client.ProgramId
			var creationErr error
			var clientErr error

			BeforeEach(func() {
				programId = 543

				clientErr = errors.New("Client error!")

				vertexShader.CreatedReturns(true)
				fragmentShader.CreatedReturns(true)

				shaderClient.CreateProgramReturns(programId, nil)
				shaderClient.AttachShaderToProgramReturns(nil)
				shaderClient.LinkProgramReturns(nil)
			})

			itRemoteProgramCreationErrored := func() {
				It("should return an error", func() {
					Ω(creationErr).Should(HaveOccurred())
				})
			}

			itRemoteProgramCreationErroredWith := func(err *error) {
				It("should return an error", func() {
					Ω(creationErr).Should(Equal(*err))
				})
			}

			JustBeforeEach(func() {
				creationErr = remoteProgram.Create()
			})

			Context("when client responds as expected", func() {
				It("did not error on creation", func() {
					Ω(creationErr).ShouldNot(HaveOccurred())
				})

				It("has correct Id", func() {
					Ω(remoteProgram.Id()).Should(Equal(programId))
				})

				It("is created", func() {
					Ω(remoteProgram.Created()).Should(BeTrue())
				})

				It("made the proper calls to the client", func() {
					Ω(shaderClient.CreateProgramCallCount()).Should(Equal(1))
					Ω(shaderClient.AttachShaderToProgramCallCount()).Should(Equal(2))
					argShaderId, argProgramId := shaderClient.AttachShaderToProgramArgsForCall(0)
					Ω(argShaderId).Should(Equal(vertexShaderId))
					Ω(argProgramId).Should(Equal(programId))
					argShaderId, argProgramId = shaderClient.AttachShaderToProgramArgsForCall(1)
					Ω(argShaderId).Should(Equal(fragmentShaderId))
					Ω(argProgramId).Should(Equal(programId))
					Ω(shaderClient.LinkProgramCallCount()).Should(Equal(1))
				})
			})

			Context("when vertex shader is not created", func() {
				BeforeEach(func() {
					vertexShader.CreatedReturns(false)
				})

				itRemoteProgramCreationErrored()
			})

			Context("when fragment shader is not created", func() {
				BeforeEach(func() {
					fragmentShader.CreatedReturns(false)
				})

				itRemoteProgramCreationErrored()
			})

			Context("when program cannot be allocated", func() {
				BeforeEach(func() {
					shaderClient.CreateProgramReturns(nil, clientErr)
				})

				itRemoteProgramCreationErroredWith(&clientErr)
			})

			Context("when shader cannot be attached", func() {
				BeforeEach(func() {
					shaderClient.AttachShaderToProgramReturns(clientErr)
				})

				itRemoteProgramCreationErroredWith(&clientErr)
			})

			Context("when program cannot be linked", func() {
				BeforeEach(func() {
					shaderClient.LinkProgramReturns(clientErr)
				})

				itRemoteProgramCreationErroredWith(&clientErr)
			})

			Describe("Usage", func() {
				var usageErr error

				JustBeforeEach(func() {
					usageErr = remoteProgram.Use()
				})

				Context("when client responds as expected", func() {
					It("did not error on use", func() {
						Ω(usageErr).ShouldNot(HaveOccurred())
					})

					It("made the proper calls to the client", func() {
						Ω(shaderClient.UseProgramCallCount()).Should(Equal(1))
						argProgramId := shaderClient.UseProgramArgsForCall(0)
						Ω(argProgramId).Should(Equal(programId))
					})
				})

				Context("when program cannot be used", func() {
					BeforeEach(func() {
						shaderClient.UseProgramReturns(clientErr)
					})

					It("errored on use", func() {
						Ω(usageErr).Should(Equal(clientErr))
					})
				})
			})

			Describe("Remote Deletion", func() {
				var deletionErr error

				JustBeforeEach(func() {
					deletionErr = remoteProgram.Delete()
				})

				Context("when client responds as expected", func() {
					It("did not error on delete", func() {
						Ω(deletionErr).ShouldNot(HaveOccurred())
					})

					itRemoteProgramHasNoId()

					itRemoteProgramIsNotCreated()

					It("made the proper calls to the facade", func() {
						Ω(shaderClient.DeleteProgramCallCount()).Should(Equal(1))
						argProgramId := shaderClient.DeleteProgramArgsForCall(0)
						Ω(argProgramId).Should(Equal(programId))
					})
				})

				Context("when client responds as expected", func() {
					BeforeEach(func() {
						shaderClient.DeleteProgramReturns(clientErr)
					})

					It("errored on delete", func() {
						Ω(deletionErr).Should(Equal(clientErr))
					})
				})
			})
		})
	})
})
