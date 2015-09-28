package ecs

type Query struct {
	componentTypes []ComponentType
}

func MakeQuery(componentTypes ...ComponentType) Query {
	return Query{
		componentTypes: componentTypes,
	}
}

//go:generate gostub EntityManager

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

	// AddEntityComponent adds a component of the specified type
	// to the specified entity
	AddEntityComponent(Entity, ComponentType, interface{})

	// EntityHasComponent checks whether the specified entity has a
	// component of the specified type
	EntityHasComponent(Entity, ComponentType) bool

	// EntityComponent returns the component of the specified type contained
	// by the specified entity
	EntityComponent(Entity, ComponentType) interface{}

	// RemoveEntityComponent removes the component of the specified type from
	// the specified entity
	RemoveEntityComponent(Entity, ComponentType)

	Search(Query) []Entity
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
	m.entityMap[entity] = entityDescriptor{
		Components: make(map[ComponentType]interface{}),
	}
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

func (m *entityManager) AddEntityComponent(entity Entity, compType ComponentType, component interface{}) {
	descriptor := m.entityMap[entity]
	descriptor.Components[compType] = component
}

func (m *entityManager) EntityHasComponent(entity Entity, compType ComponentType) bool {
	descriptor := m.entityMap[entity]
	_, contains := descriptor.Components[compType]
	return contains
}

func (m *entityManager) EntityComponent(entity Entity, compType ComponentType) interface{} {
	descriptor := m.entityMap[entity]
	return descriptor.Components[compType]
}

func (m *entityManager) RemoveEntityComponent(entity Entity, compType ComponentType) {
	descriptor := m.entityMap[entity]
	delete(descriptor.Components, compType)
}

func (m *entityManager) Search(query Query) []Entity {
	result := []Entity{}
	for entity, descriptor := range m.entityMap {
		if descriptor.hasComponents(query) {
			result = append(result, entity)
		}
	}
	return result
}

type entityDescriptor struct {
	Components map[ComponentType]interface{}
}

func (d entityDescriptor) hasComponents(query Query) bool {
	for _, componentType := range query.componentTypes {
		if _, found := d.Components[componentType]; !found {
			return false
		}
	}
	return true
}
