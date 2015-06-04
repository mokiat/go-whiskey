package uniform_test

import (
	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/client/client_fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/uniform"
	"github.com/momchil-atanasov/go-whiskey/math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	var uniformClient *client_fakes.FakeUniformClient
	var location client.UniformLocation

	BeforeEach(func() {
		uniformClient = new(client_fakes.FakeUniformClient)
		location = 10
	})

	Describe("UniformValue", func() {
		var value UniformValue

		JustBeforeEach(func() {
			value.Bind(uniformClient, location)
		})

		Describe("Vec4UniformValue", func() {
			var vector math.Vec4

			BeforeEach(func() {
				vector = math.Vec4{X: 1.0, Y: 2.0, Z: 3.0, W: 4.0}
				value = Vec4UniformValue{
					Value: vector,
				}
			})

			It("bind leads to proper client calls", func() {
				Ω(uniformClient.BindVec4UniformCallCount()).Should(Equal(1))
				argLocation, argVector := uniformClient.BindVec4UniformArgsForCall(0)
				Ω(argLocation).Should(Equal(location))
				Ω(argVector).Should(Equal(vector))
			})
		})

		Describe("Mat4x4UniformValue", func() {
			var matrix math.Mat4x4

			BeforeEach(func() {
				matrix = math.OrthoMat4x4(-1.0, 2.0, 3.0, -4.0, 1.5, 20.0)
				value = Mat4x4UniformValue{
					Value: matrix,
				}
			})

			It("bind leads to proper client calls", func() {
				Ω(uniformClient.BindMat4x4UniformCallCount()).Should(Equal(1))
				argLocation, argMatrix := uniformClient.BindMat4x4UniformArgsForCall(0)
				Ω(argLocation).Should(Equal(location))
				Ω(argMatrix).Should(Equal(matrix))
			})
		})
	})
})
