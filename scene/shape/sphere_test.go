package shape_test

import (
	"github.com/momchil-atanasov/go-whiskey/math"
	. "github.com/momchil-atanasov/go-whiskey/math/test_helpers"
	. "github.com/momchil-atanasov/go-whiskey/scene/shape"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sphere", func() {

	var sphere Sphere

	BeforeEach(func() {
		sphere = Sphere{
			Position: math.Vec3{X: 2.0, Y: 3.0, Z: 4.0},
			Radius:   5.0,
		}
	})

	It("#Diameter", func() {
		AssertFloatEquals(sphere.Diameter(), 10.0)
	})

	It("#ContainsPoint", func() {
		innerPoint := math.Vec3{X: 2.1, Y: 3.1, Z: 4.1}
		Ω(sphere.ContainsPoint(innerPoint)).Should(BeTrue())

		boundaryPoint := math.Vec3{X: 4.88, Y: 5.88, Z: 6.88}
		Ω(sphere.ContainsPoint(boundaryPoint)).Should(BeTrue())

		outsidePoint := math.Vec3{X: 4.90, Y: 5.90, Z: 6.90}
		Ω(sphere.ContainsPoint(outsidePoint)).Should(BeFalse())
	})

})
