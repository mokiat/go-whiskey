package shader_test

import (
	"github.com/momchil-atanasov/go-whiskey/graphics"
	"github.com/momchil-atanasov/go-whiskey/graphics/fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/shader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FragmentShader", func() {
	var facade *fakes.FakeFacade
	var sourceCode string
	var shader FragmentShader

	BeforeEach(func() {
		facade = new(fakes.FakeFacade)
		sourceCode = "#version 100 ..."
		shader = NewFragmentShader(facade, sourceCode)
	})

	It("is possible to get the source code", func() {
		Ω(shader.SourceCode()).Should(Equal(sourceCode))
	})

	It("is not created remotely by defualt", func() {
		Ω(shader.CreatedRemotely()).Should(BeFalse())
	})

	It("has an invalid Id by default", func() {
		Ω(shader.Id()).Should(Equal(graphics.InvalidShaderId))
	})

	Context("when shader is created remotely", func() {
		var shaderId graphics.ResourceId

		BeforeEach(func() {
			shaderId = 321
			facade.CreateFragmentShaderReturns(shaderId)

			shader.CreateRemotely()
		})

		It("is created remotely", func() {
			Ω(shader.CreatedRemotely()).Should(BeTrue())
		})

		It("has the proper shader ID", func() {
			Ω(shader.Id()).Should(Equal(shaderId))
		})

		It("has made the proper calls to the facade", func() {
			Ω(facade.CreateFragmentShaderCallCount()).Should(Equal(1))
			Ω(facade.SetShaderSourceCodeCallCount()).Should(Equal(1))
			argShaderId, argShaderSource := facade.SetShaderSourceCodeArgsForCall(0)
			Ω(argShaderId).Should(Equal(shaderId))
			Ω(argShaderSource).Should(Equal(sourceCode))
			Ω(facade.CompileShaderCallCount()).Should(Equal(1))
			argShaderId = facade.CompileShaderArgsForCall(0)
			Ω(argShaderId).Should(Equal(shaderId))
		})

		Context("when shader is deleted remotely", func() {
			BeforeEach(func() {
				shader.DeleteRemotely()
			})

			It("is not created remotely anymore", func() {
				Ω(shader.CreatedRemotely()).Should(BeFalse())
			})

			It("has an invalid ID once again", func() {
				Ω(shader.Id()).Should(Equal(graphics.InvalidShaderId))
			})

			It("has made the preper calls to the facade", func() {
				Ω(facade.DeleteShaderCallCount()).Should(Equal(1))
				argShaderId := facade.DeleteShaderArgsForCall(0)
				Ω(argShaderId).Should(Equal(shaderId))
			})
		})
	})
})
