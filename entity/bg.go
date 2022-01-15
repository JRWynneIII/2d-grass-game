package entity

import (
	"2d-grass/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type BgTileEntity struct {
	X, Y, Height, Width float64
	Image               *ebiten.Image
}

func NewBgTileEntity(b *BgTileEntity, op *ebiten.DrawImageOptions) *Entity {
	entity := NewEntity()
	entity.Position = component.NewPositionComponent(b.X, b.Y)
	entity.Size = component.NewSizeComponent(b.Height, b.Width)
	entity.Render = component.NewRenderComponent(b.Image)
	entity.Render.Options = op
	return entity
}
