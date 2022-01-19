package systems

import (
	"2d-grass/component"
	"2d-grass/entity"
	"2d-grass/preload"
	"2d-grass/viewport"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundSystem struct {
	System
	SizeComponent     *component.SizeComponent
	ViewportComponent *component.ViewportComponent
}

func NewBackgroundSystem(height, width int, vp *viewport.Viewport) *BackgroundSystem {
	bgsys := &BackgroundSystem{}
	bgsys.ViewportComponent = component.NewViewportComponent(vp)
	bgsys.SizeComponent = component.NewSizeComponent(height, width)
	return bgsys
}

func (s *BackgroundSystem) Init() {
	for y := 0; y < s.SizeComponent.Height; y += 64 {
		for x := 0; x < s.SizeComponent.Width; x += 64 {
			img := preload.Get("bg")

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
		//TODO:do offset calculation here
		//TODO: Also need to do the offset in citysystem
		screen.DrawImage(entity.Render.Image, entity.Render.Options)
	}
}
