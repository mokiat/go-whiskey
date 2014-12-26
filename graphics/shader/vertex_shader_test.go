package shader_test

import (
	"errors"

	"github.com/momchil-atanasov/go-whiskey/graphics"
	"github.com/momchil-atanasov/go-whiskey/graphics/fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/shader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VertexShader", func() {
	var facade *fakes.FakeFacade
	var sourceCode string
	var shader VertexShader

	itIsNotCreatedRemotely := func() {
		It("is not created remotely", func() {
			Ω(shader.CreatedRemotely()).Should(BeFalse())
		})
	}

	itHasInvalidId := func() {
		It("has an invalid Id", func() {
			Ω(shader.Id()).Should(Equal(graphics.InvalidShaderId))
		})
	}

	BeforeEach(func() {
		facade = new(fakes.FakeFacade)
		sourceCode = "#version 100 ..."
		shader = NewVertexShader(facade, sourceCode)
	})

	It("is possible to get the source code", func() {
		Ω(shader.SourceCode()).Should(Equal(sourceCode))
	})

	itIsNotCreatedRemotely()

	itHasInvalidId()

	Describe("Remote Creation", func() {
		var shaderId graphics.ResourceId
		var createErr error

		BeforeEach(func() {
			shaderId = 123
			facade.CreateVertexShaderReturns(shaderId, nil)
			facade.CompileShaderReturns(nil)
		})

		JustBeforeEach(func() {
			createErr = shader.CreateRemotely()
		})

		Context("happy path", func() {
			It("is created remotely", func() {
				Ω(shader.CreatedRemotely()).Should(BeTrue())
			})

			It("has the proper shader ID", func() {
				Ω(shader.Id()).Should(Equal(shaderId))
			})

			It("does not return an error", func() {
				Ω(createErr).ShouldNot(HaveOccurred())
			})

			It("has made the proper calls to the facade", func() {
				Ω(facade.CreateVertexShaderCallCount()).Should(Equal(1))
				Ω(facade.SetShaderSourceCodeCallCount()).Should(Equal(1))
				argShaderId, argShaderSource := facade.SetShaderSourceCodeArgsForCall(0)
				Ω(argShaderId).Should(Equal(shaderId))
				Ω(argShaderSource).Should(Equal(sourceCode))
				Ω(facade.CompileShaderCallCount()).Should(Equal(1))
				argShaderId = facade.CompileShaderArgsForCall(0)
				Ω(argShaderId).Should(Equal(shaderId))
			})
		})

		Context("when shader could not be allocated", func() {
			var allocErr error

			BeforeEach(func() {
				allocErr = errors.New("Could not allocate shader.")
				facade.CreateVertexShaderReturns(shaderId, allocErr)
			})

			itIsNotCreatedRemotely()

			itHasInvalidId()

			It("returns an error", func() {
				Ω(createErr).Should(Equal(allocErr))
			})
		})

		Context("when shader could not be compiled", func() {
			var compileErr error

			BeforeEach(func() {
				compileErr = errors.New("Could not compile shader.")
				facade.CompileShaderReturns(compileErr)
			})

			It("returns an error", func() {
				Ω(createErr).Should(Equal(compileErr))
			})
		})

		Describe("Remote Deletion", func() {
			JustBeforeEach(func() {
				shader.DeleteRemotely()
			})

			itIsNotCreatedRemotely()

			itHasInvalidId()

			It("has made the preper calls to the facade", func() {
				Ω(facade.DeleteShaderCallCount()).Should(Equal(1))
				argShaderId := facade.DeleteShaderArgsForCall(0)
				Ω(argShaderId).Should(Equal(shaderId))
			})
		})
	})
})
