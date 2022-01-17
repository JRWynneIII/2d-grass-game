package entity

import (
	"2d-grass/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerEntity struct {
	X, Y, Height, Width float64
	Image               *ebiten.Image
}

func NewPlayerEntity(p *PlayerEntity) *Entity {
	entity := NewEntity()
	entity.Position = component.NewPositionComponent(p.X, p.Y)
	entity.Size = component.NewSizeComponent(p.Height, p.Width)
	entity.Render = component.NewRenderComponent(p.Image)
	op := &ebiten.DrawImageOptions{}

	entity.Render.Options = op
	return entity
}
