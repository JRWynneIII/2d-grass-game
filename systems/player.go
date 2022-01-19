package systems

import (
	"2d-grass/entity"
	"2d-grass/preload"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerSystem struct {
	System
}

func NewPlayerSystem() *PlayerSystem {
	return &PlayerSystem{}
}

func (s *PlayerSystem) Init() {
	// Add a single player entity to the system
	s.System.AddEntity(entity.NewPlayerEntity(&entity.PlayerEntity{
		320.0,
		240.0,
		16,
		16,
		preload.Get("player"),
	}))
}

func (s *PlayerSystem) Update() {
	// Change the animation to 'move' here
}

func (s *PlayerSystem) Draw(screen *ebiten.Image) {
	// Just draw the player entity to the screen
	// Probably need to do the spritesheet manipulation here
	for _, entity := range s.System.entities {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(entity.Position.XPos, entity.Position.YPos)

		tileXIdx := 2
		tileSize := 16
		// We need to + 2 here because the spritesheet has padded boarders for some reason
		sx := (0%tileXIdx)*tileSize + 2
		sy := (0/tileXIdx)*tileSize + 2
		// Do the tiling of the spritesheet here
		screen.DrawImage(entity.Render.Image.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
	}
}
