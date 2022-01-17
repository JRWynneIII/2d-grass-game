package entity

import (
	"2d-grass/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type CityEntity struct {
	X, Y, Height, Width float64
	Image               *ebiten.Image
}

func NewCityEntity(c *CityEntity) *Entity {
	entity := NewEntity()
	entity.Position = component.NewPositionComponent(c.X, c.Y)
	entity.Size = component.NewSizeComponent(c.Height, c.Width)
	entity.Render = component.NewRenderComponent(c.Image)
	return entity
}
