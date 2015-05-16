package ecs

// NewDynamicEntityMap creates a new EntityMap that grows
// dynamically with the addition of new entries.
func NewDynamicEntityMap() EntityMap {
	return &dynamicEntityMap{
		mapping: make(map[EntityId]interface{}),
	}
}

type dynamicEntityMap struct {
	mapping map[EntityId]interface{}
}

func (m *dynamicEntityMap) Size() int {
	return len(m.mapping)
}

func (m *dynamicEntityMap) Put(entity EntityId, value interface{}) {
	m.mapping[entity] = value
}

func (m *dynamicEntityMap) Get(entity EntityId) interface{} {
	return m.mapping[entity]
}

func (m *dynamicEntityMap) Has(entity EntityId) bool {
	_, contains := m.mapping[entity]
	return contains
}

func (m *dynamicEntityMap) Delete(entity EntityId) bool {
	_, contains := m.mapping[entity]
	delete(m.mapping, entity)
	return contains
}

func (m *dynamicEntityMap) Clear() {
	m.mapping = make(map[EntityId]interface{})
}
