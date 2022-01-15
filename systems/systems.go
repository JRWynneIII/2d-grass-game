package systems

import (
	"2d-grass/component"
	"2d-grass/entity"
)

type System struct {
	entities []*entity.Entity
	// Does a system even need a render component?
	// Just Entities should have render components
	// What kind of components should a system have?
	//
	components []*component.Component
}

func (s *System) GetComponents() []*component.Component {
	return s.components
}

func (s *System) GetEntities() []*entity.Entity {
	return s.entities
}

func (s *System) AddEntity(e *entity.Entity) {
	s.entities = append(s.entities, e)
}

func (s *System) AddComponent(c *component.Component) {
	s.components = append(s.components, c)
}
