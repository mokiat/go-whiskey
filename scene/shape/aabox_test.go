package shape_test

import (
	"github.com/mokiat/go-whiskey/math"
	. "github.com/mokiat/go-whiskey/math/test_helpers"
	. "github.com/mokiat/go-whiskey/scene/shape"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AABox", func() {
	var box AABox

	BeforeEach(func() {
		box = AABox{
			Position:   math.Vec3{X: 2.0, Y: 3.0, Z: 4.0},
			HalfWidth:  2.5,
			HalfHeight: 3.0,
			HalfDepth:  3.5,
		}
	})

	It("#Width", func() {
		Ω(box.Width()).Should(EqualFloat32(5.0))
	})

	It("#Height", func() {
		Ω(box.Height()).Should(EqualFloat32(6.0))
	})

	It("#Depth", func() {
		Ω(box.Depth()).Should(EqualFloat32(7.0))
	})

	It("#TopLeftFrontCorner", func() {
		corner := box.TopLeftFrontCorner()
		Ω(corner).Should(HaveVec3Coords(-0.5, 6.0, 7.5))
	})

	It("#TopRightFrontCorner", func() {
		corner := box.TopRightFrontCorner()
		Ω(corner).Should(HaveVec3Coords(4.5, 6.0, 7.5))
	})

	It("#BottomLeftFrontCorner", func() {
		corner := box.BottomLeftFrontCorner()
		Ω(corner).Should(HaveVec3Coords(-0.5, 0.0, 7.5))
	})

	It("#BottomRightFrontCorner", func() {
		corner := box.BottomRightFrontCorner()
		Ω(corner).Should(HaveVec3Coords(4.5, 0.0, 7.5))
	})

	It("#TopLeftBackCorner", func() {
		corner := box.TopLeftBackCorner()
		Ω(corner).Should(HaveVec3Coords(-0.5, 6.0, 0.5))
	})

	It("#TopRightBackCorner", func() {
		corner := box.TopRightBackCorner()
		Ω(corner).Should(HaveVec3Coords(4.5, 6.0, 0.5))
	})

	It("#BottomLeftBackCorner", func() {
		corner := box.BottomLeftBackCorner()
		Ω(corner).Should(HaveVec3Coords(-0.5, 0.0, 0.5))
	})

	It("#BottomRightBackCorner", func() {
		corner := box.BottomRightBackCorner()
		Ω(corner).Should(HaveVec3Coords(4.5, 0.0, 0.5))
	})

	It("#ContainsPoint", func() {
		innerPoint := math.Vec3{X: 2.1, Y: 3.1, Z: 4.1}
		Ω(box.ContainsPoint(innerPoint)).Should(BeTrue())

		boundaryPoint := math.Vec3{X: 4.49, Y: 5.99, Z: 7.49}
		Ω(box.ContainsPoint(boundaryPoint)).Should(BeTrue())

		outsidePoint := math.Vec3{X: -4.0, Y: -15.0, Z: -13.0}
		Ω(box.ContainsPoint(outsidePoint)).Should(BeFalse())
	})

})
