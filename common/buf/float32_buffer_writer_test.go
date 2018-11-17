package buf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/go-whiskey/common/buf"
)

var _ = Describe("Float32BufferWriter", func() {
	var buffer Float32Buffer
	var writer *Float32BufferWriter

	BeforeEach(func() {
		buffer = CreateFloat32Buffer(6)
		writer = CreateFloat32BufferWriter(buffer)
	})

	It("#PutValue", func() {
		writer.PutValue(1)
		writer.PutValue(2)
		writer.PutValue(3)

		Ω(buffer.Get(0)).Should(Equal(float32(1)))
		Ω(buffer.Get(1)).Should(Equal(float32(2)))
		Ω(buffer.Get(2)).Should(Equal(float32(3)))
	})

	It("#PutValue2", func() {
		writer.PutValue2(1, 2)
		writer.PutValue2(3, 4)

		Ω(buffer.Get(0)).Should(Equal(float32(1)))
		Ω(buffer.Get(1)).Should(Equal(float32(2)))
		Ω(buffer.Get(2)).Should(Equal(float32(3)))
		Ω(buffer.Get(3)).Should(Equal(float32(4)))
	})

	It("#PutValue3", func() {
		writer.PutValue3(1, 2, 3)
		writer.PutValue3(4, 5, 6)

		Ω(buffer.Get(0)).Should(Equal(float32(1)))
		Ω(buffer.Get(1)).Should(Equal(float32(2)))
		Ω(buffer.Get(2)).Should(Equal(float32(3)))
		Ω(buffer.Get(3)).Should(Equal(float32(4)))
		Ω(buffer.Get(4)).Should(Equal(float32(5)))
		Ω(buffer.Get(5)).Should(Equal(float32(6)))
	})
})
