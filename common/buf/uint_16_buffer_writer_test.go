package buf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/go-whiskey/common/buf"
)

var _ = Describe("Uint16BufferWriter", func() {
	var buffer UInt16Buffer
	var writer *UInt16BufferWriter

	BeforeEach(func() {
		buffer = CreateUInt16Buffer(2)
		writer = CreateUInt16BufferWriter(buffer)
	})

	It("#PutValue", func() {
		writer.PutValue(3)
		writer.PutValue(4)

		Ω(buffer.Get(0)).Should(Equal(uint16(3)))
		Ω(buffer.Get(1)).Should(Equal(uint16(4)))
	})
})
