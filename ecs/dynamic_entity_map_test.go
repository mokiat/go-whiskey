package ecs_test

import (
	. "github.com/momchil-atanasov/go-whiskey/ecs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DynamicEntityMap", func() {
	var entityMap EntityMap
	var firstEntity EntityId
	var secondEntity EntityId

	BeforeEach(func() {
		entityMap = NewDynamicEntityMap()
		firstEntity = EntityId(0)
		secondEntity = EntityId(1)
	})

	It("is empty by default", func() {
		Ω(entityMap.Size()).Should(Equal(0))
	})

	It("does not contain entries for any entity", func() {
		Ω(entityMap.Has(firstEntity)).Should(BeFalse())
		Ω(entityMap.Get(firstEntity)).Should(BeNil())
		Ω(entityMap.Has(secondEntity)).Should(BeFalse())
		Ω(entityMap.Get(secondEntity)).Should(BeNil())
	})

	Context("when entities are added to the map", func() {
		BeforeEach(func() {
			entityMap.Put(firstEntity, "first")
			entityMap.Put(secondEntity, "second")
		})

		It("size should have changed accordingly", func() {
			Ω(entityMap.Size()).Should(Equal(2))
		})

		It("is contains all the entries", func() {
			Ω(entityMap.Has(firstEntity)).Should(BeTrue())
			Ω(entityMap.Get(firstEntity)).Should(Equal("first"))
			Ω(entityMap.Has(secondEntity)).Should(BeTrue())
			Ω(entityMap.Get(secondEntity)).Should(Equal("second"))
		})

		Context("when cleared", func() {
			BeforeEach(func() {
				entityMap.Clear()
			})

			It("is empty", func() {
				Ω(entityMap.Size()).Should(Equal(0))
			})

			It("does not contain entries for any entity", func() {
				Ω(entityMap.Has(firstEntity)).Should(BeFalse())
				Ω(entityMap.Get(firstEntity)).Should(BeNil())
				Ω(entityMap.Has(secondEntity)).Should(BeFalse())
				Ω(entityMap.Get(secondEntity)).Should(BeNil())
			})
		})

		Context("when an entry is removed", func() {
			BeforeEach(func() {
				entityMap.Delete(firstEntity)
			})

			It("size should have changed accordingly", func() {
				Ω(entityMap.Size()).Should(Equal(1))
			})

			It("should not contain that entry anymore", func() {
				Ω(entityMap.Has(firstEntity)).Should(BeFalse())
				Ω(entityMap.Get(firstEntity)).Should(BeNil())
			})

			It("should still contain all other entries", func() {
				Ω(entityMap.Has(secondEntity)).Should(BeTrue())
				Ω(entityMap.Get(secondEntity)).Should(Equal("second"))
			})
		})
	})
})
