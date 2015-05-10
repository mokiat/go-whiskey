package ecs

//go:generate counterfeiter -o ecs_fakes/fake_entity_manager.go ./ EntityManager

// EntityManager is manages the lifecycle of all Entity object in the game.
type EntityManager interface {

	// CreateEntity creates a brand new Entity object.
	CreateEntity() Entity

	// HasEntity can be used to check if a given entity is part of this
	// EntityManager or has been deleted.
	HasEntity(Entity) bool

	// DeleteEntity deletes an existing Entity object.
	DeleteEntity(Entity)

	// Deletes all existing Entity objects managed by this EntityManager.
	DeleteAllEntities()
}

// Creates a new EntityManager instance
func NewEntityManager() EntityManager {
	return &entityManager{
		entitySet: make(map[Entity]struct{}),
	}
}

type entityManager struct {
	idCounter int32
	entitySet map[Entity]struct{}
}

func (m *entityManager) CreateEntity() Entity {
	m.idCounter++
	entity := Entity(m.idCounter)
	m.entitySet[entity] = struct{}{}
	return entity
}

func (m *entityManager) HasEntity(entity Entity) bool {
	_, contains := m.entitySet[entity]
	return contains
}

func (m *entityManager) DeleteEntity(entity Entity) {
	delete(m.entitySet, entity)
}

func (m *entityManager) DeleteAllEntities() {
	m.entitySet = make(map[Entity]struct{})
}
