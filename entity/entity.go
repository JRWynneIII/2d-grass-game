package entity

import "2d-grass/component"

var curID = 0

type Entity struct {
	id       *component.IDComponent
	Position *component.PositionComponent
	Size     *component.SizeComponent
	Render   *component.RenderComponent
}

func NewEntity() *Entity {
	entity := &Entity{
		id: component.NewIDComponent(curID),
	}
	curID++
	return entity
}
