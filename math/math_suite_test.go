package math_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Math Suite")
}

func AssertFloatEquals(actualValue, expectedValue float32) {
	Ω(actualValue).Should(BeNumerically("~", expectedValue))
}
