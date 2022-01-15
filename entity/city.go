package entity

import (
	"2d-grass/component"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CityEntity struct {
	X, Y, Height, Width float64
}

func NewCityEntity(c *CityEntity) *Entity {
	entity := NewEntity()
	entity.Position = component.NewPositionComponent(c.X, c.Y)
	entity.Size = component.NewSizeComponent(c.Height, c.Width)
	img, _, err := ebitenutil.NewImageFromFile("assets/fire_0.png")
	if err != nil {
		log.Print(err)
	}
	entity.Render = component.NewRenderComponent(img)
	return entity
}
