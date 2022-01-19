package systems

import (
	"2d-grass/component"
	"2d-grass/entity"
	"2d-grass/preload"
	"2d-grass/viewport"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type BackgroundSystem struct {
	System
	SizeComponent     *component.SizeComponent
	ViewportComponent *component.ViewportComponent
}

func NewBackgroundSystem(height, width int, vp *viewport.Viewport) *BackgroundSystem {
	bgsys := &BackgroundSystem{}
	bgsys.ViewportComponent = component.NewViewportComponent(vp)
	bgsys.SizeComponent = component.NewSizeComponent(float64(height), float64(width))
	return bgsys
}

func (s *BackgroundSystem) Init() {
	for y := 0; y < int(s.SizeComponent.Height); y += 64 {
		for x := 0; x < int(s.SizeComponent.Width); x += 64 {
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
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		s.ViewportComponent.Viewport.Move(viewport.Up)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		s.ViewportComponent.Viewport.Move(viewport.Left)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		s.ViewportComponent.Viewport.Move(viewport.Down)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		s.ViewportComponent.Viewport.Move(viewport.Right)
	}
}

func (s *BackgroundSystem) Draw(screen *ebiten.Image) {
	for idx, entity := range s.System.entities {
		//TODO:do offset calculation here
		//TODO: Also need to do the offset in citysystem
		offsetBegin := s.ViewportComponent.Viewport.X * s.ViewportComponent.Viewport.Y
		offsetEnd := (s.ViewportComponent.Viewport.X + 640) * (s.ViewportComponent.Viewport.Y + 480)
		if idx >= offsetBegin && idx <= offsetEnd {
			entity.Render.Options.GeoM.Translate(float64(-s.ViewportComponent.Viewport.X), float64(-s.ViewportComponent.Viewport.Y))
			screen.DrawImage(entity.Render.Image, entity.Render.Options)
		}
	}
}
