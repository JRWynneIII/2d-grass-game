package systems

import (
	"2d-grass/entity"
)

type System struct {
	entities []*entity.Entity
}

func (s *System) GetEntities() []*entity.Entity {
	return s.entities
}

func (s *System) AddEntity(e *entity.Entity) {
	s.entities = append(s.entities, e)
}
