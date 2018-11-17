package buf_test

import (
	. "github.com/mokiat/go-whiskey/common/buf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Float32Buffer", func() {

	var data []byte
	var buffer Float32Buffer

	BeforeEach(func() {
		data = make([]byte, 9)
		buffer = Float32Buffer(data)
	})

	It("has the correct length", func() {
		Ω(buffer.Length()).Should(Equal(2))
	})

	It("has default values initially", func() {
		Ω(buffer.Get(0)).Should(Equal(float32(0.0)))
		Ω(buffer.Get(1)).Should(Equal(float32(0.0)))
	})

	It("is possible to create one from scratch", func() {
		buffer = CreateFloat32Buffer(3)
		Ω(buffer.Length()).Should(Equal(3))
	})

	Context("when values have been set", func() {
		BeforeEach(func() {
			buffer.Set(0, 0.1)
			buffer.Set(1, 2.3)
		})

		It("is possible to get the new values", func() {
			Ω(buffer.Get(0)).Should(Equal(float32(0.1)))
			Ω(buffer.Get(1)).Should(Equal(float32(2.3)))
		})

		It("has not changed its length", func() {
			Ω(buffer.Length()).Should(Equal(2))
		})

		It("has written the values in Little Endian to the underlying data", func() {
			Ω(data[0]).Should(Equal(byte(205)))
			Ω(data[1]).Should(Equal(byte(204)))
			Ω(data[2]).Should(Equal(byte(204)))
			Ω(data[3]).Should(Equal(byte(61)))

			Ω(data[4]).Should(Equal(byte(51)))
			Ω(data[5]).Should(Equal(byte(51)))
			Ω(data[6]).Should(Equal(byte(19)))
			Ω(data[7]).Should(Equal(byte(64)))
		})

		Describe("Sub-buffer", func() {
			var subBuffer Float32Buffer

			BeforeEach(func() {
				subBuffer = buffer.Range(1, 1)
			})

			It("has proper length", func() {
				Ω(subBuffer.Length()).Should(Equal(1))
			})

			It("has the correct values", func() {
				Ω(subBuffer.Get(0)).Should(Equal(float32(2.3)))
			})

			It("has the correct slice of values in Little Endian from the underlying slice", func() {
				Ω(subBuffer[0]).Should(Equal(byte(51)))
				Ω(subBuffer[1]).Should(Equal(byte(51)))
				Ω(subBuffer[2]).Should(Equal(byte(19)))
				Ω(subBuffer[3]).Should(Equal(byte(64)))
			})
		})
	})
})
