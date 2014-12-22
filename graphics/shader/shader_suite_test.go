package shader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shader Suite")
}
