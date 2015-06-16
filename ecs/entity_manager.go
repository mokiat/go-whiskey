package ecs

//go:generate counterfeiter -o ecs_fakes/fake_entity_manager.go ./ EntityManager

// EntityManager manages the lifecycle of all Entity object in the game.
type EntityManager interface {

	// CreateEntity creates a brand new Entity object.
	CreateEntity() Entity

	// HasEntity can be used to check if a given entity is part of this
	// EntityManager or has been deleted.
	HasEntity(Entity) bool

	// DeleteEntity deletes an existing Entity object.
	DeleteEntity(Entity)

	// DeleteAllEntities deletes all existing Entity objects managed by this
	// EntityManager.
	DeleteAllEntities()
}

// NewEntityManager creates a new EntityManager instance.
func NewEntityManager() EntityManager {
	return &entityManager{
		entityMap: make(map[Entity]entityDescriptor),
	}
}

type entityManager struct {
	idCounter int
	entityMap map[Entity]entityDescriptor
}

func (m *entityManager) CreateEntity() Entity {
	m.idCounter++
	entity := Entity(m.idCounter)
	m.entityMap[entity] = entityDescriptor{}
	return entity
}

func (m *entityManager) HasEntity(entity Entity) bool {
	_, contains := m.entityMap[entity]
	return contains
}

func (m *entityManager) DeleteEntity(entity Entity) {
	delete(m.entityMap, entity)
}

func (m *entityManager) DeleteAllEntities() {
	m.entityMap = make(map[Entity]entityDescriptor)
}

type entityDescriptor struct {
}
