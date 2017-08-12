package ecs_test

import (
	. "github.com/mokiat/go-whiskey/ecs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Component", func() {
	var manager EntityManager
	var entity Entity
	var otherEntity Entity
	var locationType ComponentType
	var motionType ComponentType

	BeforeEach(func() {
		manager = NewEntityManager()
		entity = manager.CreateEntity()
		otherEntity = manager.CreateEntity()
		locationType = ComponentType("location")
		motionType = ComponentType("motion")
	})

	It("by default entities have no components", func() {
		Ω(manager.EntityHasComponent(entity, locationType)).Should(BeFalse())
		Ω(manager.EntityHasComponent(entity, motionType)).Should(BeFalse())
	})

	Context("when components are added to entity", func() {
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

		It("entity contains the components", func() {
			Ω(manager.EntityHasComponent(entity, locationType)).Should(BeTrue())
			Ω(manager.EntityHasComponent(entity, motionType)).Should(BeTrue())
		})

		It("is possible to access those components", func() {
			Ω(manager.EntityComponent(entity, locationType)).Should(Equal(location))
			Ω(manager.EntityComponent(entity, motionType)).Should(Equal(motion))
		})

		It("other entities do not have those components", func() {
			Ω(manager.EntityHasComponent(otherEntity, locationType)).Should(BeFalse())
			Ω(manager.EntityHasComponent(otherEntity, motionType)).Should(BeFalse())
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
