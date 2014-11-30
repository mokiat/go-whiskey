package shape_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShape(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shape Suite")
}
