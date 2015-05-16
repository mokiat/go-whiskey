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
		entityMap: NewDynamicEntityMap(),
	}
}

type entityManager struct {
	idCounter uint16
	entityMap EntityMap
}

func (m *entityManager) CreateEntity() Entity {
	m.idCounter++
	entity := Entity{
		Id:      EntityId(m.idCounter),
		Version: 0,
	}
	m.entityMap.Put(entity.Id, EntityDescriptor{})
	return entity
}

func (m *entityManager) HasEntity(entity Entity) bool {
	return m.entityMap.Has(entity.Id)
}

func (m *entityManager) DeleteEntity(entity Entity) {
	m.entityMap.Delete(entity.Id)
}

func (m *entityManager) DeleteAllEntities() {
	m.entityMap.Clear()
}
