package ecs_test

import (
	. "github.com/momchil-atanasov/go-whiskey/ecs"

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

	Describe("Entity lifecycle", func() {
		var firstEntity Entity
		var secondEntity Entity
		var thirdEntity Entity

		BeforeEach(func() {
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

		Describe("Entity deletion", func() {
			Context("when a single entity is deleted", func() {
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
	})

	Describe("Component lifecycle", func() {
		var locationType ComponentType
		var motionType ComponentType
		var entity Entity

		BeforeEach(func() {
			locationType = ComponentType("location")
			motionType = ComponentType("motion")
			entity = manager.CreateEntity()
		})

		It("entity should not contain any components", func() {
			Ω(manager.EntityHasComponent(entity, locationType)).Should(BeFalse())
			Ω(manager.EntityHasComponent(entity, motionType)).Should(BeFalse())
		})

		Context("when components are added to Entity", func() {
			type LocationComponent struct {
				X, Y int
			}
			type MotionComponent struct {
				SpeedX, SpeedY, Mass int
			}

			var location LocationComponent
			var motion MotionComponent

			BeforeEach(func() {
				location = LocationComponent{
					X: 15, Y: 20,
				}
				motion = MotionComponent{
					SpeedX: 50, SpeedY: 11, Mass: 2,
				}

				manager.AddEntityComponent(entity, locationType, location)
				manager.AddEntityComponent(entity, motionType, motion)
			})

			It("is possible to check whether Entity has components", func() {
				Ω(manager.EntityHasComponent(entity, locationType)).Should(BeTrue())
				Ω(manager.EntityHasComponent(entity, motionType)).Should(BeTrue())
			})

			It("is possible to access those components", func() {
				Ω(manager.EntityComponent(entity, locationType)).Should(Equal(location))
				Ω(manager.EntityComponent(entity, motionType)).Should(Equal(motion))
			})

			Context("when component is removed", func() {
				BeforeEach(func() {
					manager.RemoveEntityComponent(entity, locationType)
				})

				It("component is no longer accessible", func() {
					Ω(manager.EntityHasComponent(entity, locationType)).Should(BeFalse())
					Ω(manager.EntityComponent(entity, locationType)).Should(BeNil())
				})

				It("other components are still accessible", func() {
					Ω(manager.EntityHasComponent(entity, motionType)).Should(BeTrue())
					Ω(manager.EntityComponent(entity, motionType)).Should(Equal(motion))
				})
			})
		})
	})
})
