package ecs_test

import (
	. "github.com/mokiat/go-whiskey/ecs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EntityManager", func() {
	var manager EntityManager

	BeforeEach(func() {
		manager = NewEntityManager()
	})

	It("Entity Manager is not nil", func() {
		Ω(manager).ShouldNot(BeNil())
	})
})
