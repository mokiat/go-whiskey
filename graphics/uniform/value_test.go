package uniform_test

import (
	"github.com/momchil-atanasov/go-whiskey/graphics"
	"github.com/momchil-atanasov/go-whiskey/graphics/fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/uniform"
	"github.com/momchil-atanasov/go-whiskey/math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	var facade *fakes.FakeFacade
	var location graphics.BindLocation

	BeforeEach(func() {
		facade = new(fakes.FakeFacade)
		location = 10
	})

	Describe("Vec4UniformValue", func() {
		var value Vec4UniformValue
		var vector math.Vec4

		BeforeEach(func() {
			vector = math.Vec4{X: 1.0, Y: 2.0, Z: 3.0, W: 4.0}
			value.Value = vector
		})

		It("bind leads to proper calls to the facade", func() {
			value.Bind(facade, location)

			Ω(facade.BindVec4UniformCallCount()).Should(Equal(1))
			argVector, argLocation := facade.BindVec4UniformArgsForCall(0)
			Ω(argVector).Should(Equal(vector))
			Ω(argLocation).Should(Equal(location))
		})
	})
})
