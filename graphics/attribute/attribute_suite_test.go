package attribute_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAttribute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Attribute Suite")
}
