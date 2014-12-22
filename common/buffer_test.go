package common_test

import (
	"encoding/binary"

	. "github.com/momchil-atanasov/go-whiskey/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UInt16Buffer", func() {

	var data []byte
	var array UInt16Buffer

	itBehavesLikeAnArray := func() {
		It("should have proper size", func() {
			Ω(array.Size()).Should(Equal(3))
		})

		It("is possible to get the data", func() {
			Ω(array.Bytes()).Should(Equal(data))
		})

		Context("when values are set", func() {
			BeforeEach(func() {
				array.PutUInt16(0, 10)
				array.PutUInt16(1, 11)
				array.PutUInt16(2, 12)
			})

			It("is possible to get those values back", func() {
				Ω(array.UInt16(0)).Should(Equal(uint16(10)))
				Ω(array.UInt16(1)).Should(Equal(uint16(11)))
				Ω(array.UInt16(2)).Should(Equal(uint16(12)))
			})
		})
	}

	BeforeEach(func() {
		data = make([]byte, 7)
	})

	Describe("Little Endian", func() {
		BeforeEach(func() {
			array = NewUInt16Buffer(data, binary.LittleEndian)
		})

		itBehavesLikeAnArray()
	})

	Describe("Big Endian", func() {
		BeforeEach(func() {
			array = NewUInt16Buffer(data, binary.BigEndian)
		})

		itBehavesLikeAnArray()
	})

})

var _ = Describe("Float32Buffer", func() {

	var data []byte
	var array Float32Buffer

	itBehavesLikeAnArray := func() {
		It("should have proper size", func() {
			Ω(array.Size()).Should(Equal(2))
		})

		It("is possible to get the data", func() {
			Ω(array.Bytes()).Should(Equal(data))
		})

		Context("when values are set", func() {
			BeforeEach(func() {
				array.PutFloat32(0, 1.0)
				array.PutFloat32(1, 1.1)
			})

			It("is possible to get those values back", func() {
				Ω(array.Float32(0)).Should(Equal(float32(1.0)))
				Ω(array.Float32(1)).Should(Equal(float32(1.1)))
			})
		})
	}

	BeforeEach(func() {
		data = make([]byte, 9)
	})

	Describe("Little Endian", func() {
		BeforeEach(func() {
			array = NewFloat32Buffer(data, binary.LittleEndian)
		})

		itBehavesLikeAnArray()
	})

	Describe("Big Endian", func() {
		BeforeEach(func() {
			array = NewFloat32Buffer(data, binary.BigEndian)
		})

		itBehavesLikeAnArray()
	})

})
