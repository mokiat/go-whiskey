package buf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBuf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Buf Suite")
}
