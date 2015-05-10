package ecs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEcs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ecs Suite")
}
