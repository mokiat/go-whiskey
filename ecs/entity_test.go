package ecs_test

import (
	. "github.com/momchil-atanasov/go-whiskey/ecs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Entity", func() {
	var manager EntityManager
	var firstEntity Entity
	var secondEntity Entity
	var thirdEntity Entity

	BeforeEach(func() {
		manager = NewEntityManager()
		firstEntity = manager.CreateEntity()
		secondEntity = manager.CreateEntity()
		thirdEntity = manager.CreateEntity()
	})

	It("no two entities should be the same", func() {
		Ω(firstEntity).ShouldNot(Equal(secondEntity))
		Ω(secondEntity).ShouldNot(Equal(thirdEntity))
		Ω(thirdEntity).ShouldNot(Equal(firstEntity))
	})

	It("manager should contain all of the entities", func() {
		Ω(manager.HasEntity(firstEntity)).Should(BeTrue())
		Ω(manager.HasEntity(secondEntity)).Should(BeTrue())
		Ω(manager.HasEntity(thirdEntity)).Should(BeTrue())
	})

	Context("when an entity is deleted", func() {
		BeforeEach(func() {
			manager.DeleteEntity(secondEntity)
		})

		It("manager should not contain entity anymore", func() {
			Ω(manager.HasEntity(secondEntity)).Should(BeFalse())
		})

		It("manager should still contain the other entities", func() {
			Ω(manager.HasEntity(firstEntity)).Should(BeTrue())
			Ω(manager.HasEntity(thirdEntity)).Should(BeTrue())
		})
	})

	Context("when all entities are deleted", func() {
		BeforeEach(func() {
			manager.DeleteAllEntities()
		})

		It("manager should not contain any entity", func() {
			Ω(manager.HasEntity(firstEntity)).Should(BeFalse())
			Ω(manager.HasEntity(secondEntity)).Should(BeFalse())
			Ω(manager.HasEntity(thirdEntity)).Should(BeFalse())
		})
	})
})
