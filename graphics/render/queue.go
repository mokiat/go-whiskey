package render

//go:generate counterfeiter -o render_fakes/fake_queue.go ./ Queue

// Queue represents a collection of elements that are inserted one
// after the other and then removed one after the other.
//
// In reality, the naming Queue might not be the best but it fits
// well with rendering concepts.
type Queue interface {

	// Empty returns whether this queue is empty, hence has no more elements.
	Empty() bool

	// Size returns the number of items stored in this queue
	Size() int

	// Push adds a new element to the queue.
	Push(Item)

	// Pop removes an element from the queue.
	Pop() (Item, bool)
}
