package ecs

//go:generate gostub Event

type Event interface {
	Source() Entity
}

type EventListenerFunc func(event Event)

//go:generate gostub EventListener

type EventListener interface {
	OnEvent(event Event)
}

func NewEventListener(callback EventListenerFunc) EventListener {
	return &eventListener{
		callback: callback,
	}
}

type eventListener struct {
	callback EventListenerFunc
}

func (l *eventListener) OnEvent(event Event) {
	l.callback(event)
}
