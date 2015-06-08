package render

// NewFIFOQueue is an implementation of the Queue interface
// that returns elements in the same order they were added
// to the queue.
//
// Note: This implementation results in malloc and GC calls.
func NewFIFOQueue() Queue {
	return &fifoQueue{}
}

type fifoQueue struct {
	items []Item
}

func (q *fifoQueue) Empty() bool {
	return q.Size() == 0
}

func (q *fifoQueue) Size() int {
	return len(q.items)
}

func (q *fifoQueue) Push(item Item) {
	q.items = append(q.items, item)
}

func (q *fifoQueue) Pop() (Item, bool) {
	if q.Empty() {
		return Item{}, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}
