package buf_test

import (
	. "github.com/momchil-atanasov/go-whiskey/common/buf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Uint16Buffer", func() {

	var data []byte
	var buffer UInt16Buffer

	BeforeEach(func() {
		data = make([]byte, 7)
		buffer = UInt16Buffer(data)
	})

	It("has the correct length", func() {
		Ω(buffer.Length()).Should(Equal(3))
	})

	It("has default values initially", func() {
		Ω(buffer.Get(0)).Should(Equal(uint16(0)))
		Ω(buffer.Get(1)).Should(Equal(uint16(0)))
		Ω(buffer.Get(2)).Should(Equal(uint16(0)))
	})

	Context("when values have been set", func() {
		BeforeEach(func() {
			buffer.Set(0, 65535)
			buffer.Set(1, 255)
			buffer.Set(2, 65280)
		})

		It("is possible to get the new values", func() {
			Ω(buffer.Get(0)).Should(Equal(uint16(65535)))
			Ω(buffer.Get(1)).Should(Equal(uint16(255)))
			Ω(buffer.Get(2)).Should(Equal(uint16(65280)))
		})

		It("has not changed its length", func() {
			Ω(buffer.Length()).Should(Equal(3))
		})

		It("has written the values in Little Endian to the underlying slice", func() {
			Ω(data[0]).Should(Equal(byte(255)))
			Ω(data[1]).Should(Equal(byte(255)))

			Ω(data[2]).Should(Equal(byte(255)))
			Ω(data[3]).Should(Equal(byte(0)))

			Ω(data[4]).Should(Equal(byte(0)))
			Ω(data[5]).Should(Equal(byte(255)))
		})

		Describe("Sub-Buffer", func() {
			var subBuffer UInt16Buffer

			BeforeEach(func() {
				subBuffer = buffer.Range(1, 1)
			})

			It("has proper length", func() {
				Ω(subBuffer.Length()).Should(Equal(1))
			})

			It("has the correct values", func() {
				Ω(subBuffer.Get(0)).Should(Equal(uint16(255)))
			})

			It("has the correct slice of values in Little Endian from the underlying slice", func() {
				Ω(subBuffer[0]).Should(Equal(byte(255)))
				Ω(subBuffer[1]).Should(Equal(byte(0)))
			})
		})
	})
})
