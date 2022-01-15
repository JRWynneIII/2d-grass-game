package systems

import (
	"2d-grass/entity"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BackgroundSystem struct {
	System
}

func NewBackgroundSystem() *BackgroundSystem {
	return &BackgroundSystem{}
}

func (s *BackgroundSystem) Init() {
	for y := 0; y < 480; y += 64 {
		for x := 0; x < 640; x += 64 {
			img, _, err := ebitenutil.NewImageFromFile("assets/grass_random_grid.png")
			if err != nil {
				log.Print(err)
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			log.Print(fmt.Sprintf("tile at %d, %d", x, y))
			bgEntity := &entity.BgTileEntity{float64(x), float64(y), 16, 16, img}
			s.AddEntity(entity.NewBgTileEntity(bgEntity, op))
		}
	}
}

func (s *BackgroundSystem) Update() {
}

func (s *BackgroundSystem) Draw(screen *ebiten.Image) {
	for _, entity := range s.System.entities {
		screen.DrawImage(entity.Render.Image, entity.Render.Options)
	}
}
