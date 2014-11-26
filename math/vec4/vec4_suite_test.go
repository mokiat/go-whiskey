package vec4_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVec4(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vec4 Suite")
}
