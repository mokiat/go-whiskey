package ecs

// EntityId represents the Id number of an Entity
type EntityId uint16

// EntityVersion represents the version of the entity Id.
// It is a mechanism which allows the reuse of an EntityId.
type EntityVersion uint16

// Entity represents an object in the Game scene. It can be anything from
// the main hero to a building or a tree.
type Entity struct {

	// Id uniquely identifies the entity. There can only be one entity
	// with this Id at any given point in time.
	Id EntityId

	// Version allows the Entity's Id to be reused after it has been deleted.
	// When a new Entity is created with the same Id, it gets a new Version
	// number, which allows to distinguish it from its predecessors.
	Version EntityVersion
}

// Internal structure which holds information about a given Entity
type EntityDescriptor struct {
}
