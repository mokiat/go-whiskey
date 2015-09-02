package ecs_test

import (
	. "github.com/momchil-atanasov/go-whiskey/ecs"
	"github.com/momchil-atanasov/go-whiskey/ecs/ecs_stubs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventListener", func() {
	var event *ecs_stubs.EventStub
	var listenerFuncCallCount int
	var listenerFuncEventParam Event
	var listener EventListener

	BeforeEach(func() {
		event = new(ecs_stubs.EventStub)
		event.SourceReturns(Entity(123))

		listenerFuncCallCount = 0
		listenerFuncEventParam = nil
		listenerFunc := func(event Event) {
			listenerFuncCallCount++
			listenerFuncEventParam = event
		}

		listener = NewEventListener(listenerFunc)
	})

	It("is not nil", func() {
		Ω(listener).ShouldNot(BeNil())
	})

	Context("when notified", func() {
		BeforeEach(func() {
			listener.OnEvent(event)
		})

		It("event is propagated to callback", func() {
			Ω(listenerFuncCallCount).Should(Equal(1))
			Ω(listenerFuncEventParam).Should(Equal(event))
		})
	})
})
