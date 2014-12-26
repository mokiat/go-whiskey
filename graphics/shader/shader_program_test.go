package shader_test

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics"
	. "github.com/momchil-atanasov/go-whiskey/graphics/fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/shader"
	. "github.com/momchil-atanasov/go-whiskey/graphics/shader/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShaderProgram", func() {

	var facade *FakeFacade
	var vertexShaderId graphics.ResourceId
	var vertexShader *FakeVertexShader
	var fragmentShaderId graphics.ResourceId
	var fragmentShader *FakeFragmentShader
	var program ShaderProgram

	itHasInvalidId := func() {
		It("has invalid ID", func() {
			Ω(program.Id()).Should(Equal(graphics.InvalidProgramId))
		})
	}

	itIsNotCreatedRemotely := func() {
		It("is not created remotely initially", func() {
			Ω(program.CreatedRemotely()).Should(BeFalse())
		})
	}

	BeforeEach(func() {
		facade = new(FakeFacade)
		vertexShaderId = 123
		vertexShader = new(FakeVertexShader)
		vertexShader.IdReturns(vertexShaderId)
		fragmentShaderId = 321
		fragmentShader = new(FakeFragmentShader)
		fragmentShader.IdReturns(fragmentShaderId)
		program = NewShaderProgram(facade, vertexShader, fragmentShader)
	})

	It("is possible to get the vertex shader", func() {
		Ω(program.VertexShader()).Should(Equal(vertexShader))
	})

	It("is possible to get the fragment shader", func() {
		Ω(program.FragmentShader()).Should(Equal(fragmentShader))
	})

	itHasInvalidId()

	itIsNotCreatedRemotely()

	Describe("Remote Creation", func() {
		var programId graphics.ResourceId
		var createErr error

		BeforeEach(func() {
			programId = 987

			vertexShader.CreatedRemotelyReturns(true)
			fragmentShader.CreatedRemotelyReturns(true)

			facade.CreateProgramReturns(programId, nil)
			facade.LinkProgramReturns(nil)
		})

		JustBeforeEach(func() {
			createErr = program.CreateRemotely()
		})

		Context("happy path", func() {
			It("has correct ID", func() {
				Ω(program.Id()).Should(Equal(programId))
			})

			It("is created remotely", func() {
				Ω(program.CreatedRemotely()).Should(BeTrue())
			})

			It("should not return error", func() {
				Ω(createErr).ShouldNot(HaveOccurred())
			})

			It("has made the proper calls to the facade", func() {
				Ω(facade.CreateProgramCallCount()).Should(Equal(1))
				Ω(facade.AttachShaderToProgramCallCount()).Should(Equal(2))
				argProgramId, argShaderId := facade.AttachShaderToProgramArgsForCall(0)
				Ω(argProgramId).Should(Equal(programId))
				Ω(argShaderId).Should(Equal(vertexShaderId))
				argProgramId, argShaderId = facade.AttachShaderToProgramArgsForCall(1)
				Ω(argProgramId).Should(Equal(programId))
				Ω(argShaderId).Should(Equal(fragmentShaderId))
				Ω(facade.LinkProgramCallCount()).Should(Equal(1))
			})
		})

		Context("when program cannot be allocated", func() {
			var allocErr error

			BeforeEach(func() {
				allocErr = errors.New("Could not allocate program!")
				facade.CreateProgramReturns(programId, allocErr)
			})

			itHasInvalidId()

			itIsNotCreatedRemotely()

			It("should return error", func() {
				Ω(createErr).Should(Equal(allocErr))
			})
		})

		Context("when program cannot be linked", func() {
			var linkErr error

			BeforeEach(func() {
				linkErr = errors.New("Could not link program!")
				facade.LinkProgramReturns(linkErr)
			})

			It("should return error", func() {
				Ω(createErr).Should(Equal(linkErr))
			})
		})

		Context("when vertex shader is not created remotely", func() {
			BeforeEach(func() {
				vertexShader.CreatedRemotelyReturns(false)
			})

			It("should return an error", func() {
				Ω(createErr).Should(HaveOccurred())
			})
		})

		Context("when fragment shader is not created remotely", func() {
			BeforeEach(func() {
				fragmentShader.CreatedRemotelyReturns(false)
			})

			It("should return an error", func() {
				Ω(createErr).Should(HaveOccurred())
			})
		})

		Describe("Remote Binding", func() {
			JustBeforeEach(func() {
				program.BindRemotely()
			})

			It("has made the proper calls to the facade", func() {
				Ω(facade.UseProgramCallCount()).Should(Equal(1))
				argProgramId := facade.UseProgramArgsForCall(0)
				Ω(argProgramId).Should(Equal(programId))
			})
		})

		Describe("Remote Deletion", func() {
			JustBeforeEach(func() {
				program.DeleteRemotely()
			})

			It("has made the proper calls to the facade", func() {
				Ω(facade.DeleteProgramCallCount()).Should(Equal(1))
				argProgramId := facade.DeleteProgramArgsForCall(0)
				Ω(argProgramId).Should(Equal(programId))
			})
		})
	})
})
