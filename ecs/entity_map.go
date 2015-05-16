package ecs

// EntityMap allows the mapping between an EntityId
// and an arbitrary value
type EntityMap interface {
	// Size returns the number of entries in this map
	Size() int

	// Put adds a new entry in this map
	Put(EntityId, interface{})

	// Get returns the value stored for the specified
	// EntityId. One should first use the Has method to
	// verify that this map does indeed contain a value
	// for the specified EntityId.
	Get(EntityId) interface{}

	// Has returns whether a value is stored in this map
	// for the specified EntityId
	Has(EntityId) bool

	// Delete removes an entry from this map and returns
	// whether a delete was indeed performed
	Delete(EntityId) bool

	// Clear removes all entries from this map
	Clear()
}
