package vec3_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVec3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vec3 Suite")
}
