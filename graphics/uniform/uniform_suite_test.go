package uniform_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUniform(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uniform Suite")
}
