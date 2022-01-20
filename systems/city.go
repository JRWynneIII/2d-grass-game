package systems

import (
	"2d-grass/component"
	"2d-grass/entity"
	"2d-grass/preload"
	"2d-grass/viewport"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type BuildCitySystem struct {
	System
	ViewportComponent *component.ViewportComponent
}

func NewBuildCitySystem(vp *viewport.Viewport) *BuildCitySystem {
	sys := &BuildCitySystem{}
	sys.ViewportComponent = component.NewViewportComponent(vp)
	return sys
}

func (s *BuildCitySystem) Update() {
	xmin, xmax := 0, 640
	ymin, ymax := 0, 480
	x := rand.Intn(xmax-xmin) + xmin
	y := rand.Intn(ymax-ymin) + ymin
	log.Print("New City at ", x, ", ", y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(float64(x), float64(y))
	s.System.AddEntity(entity.NewCityEntity(&entity.CityEntity{
		float64(x),
		float64(y),
		32,
		32,
		preload.Get("city"),
	}, op))
}

func (s *BuildCitySystem) Draw(screen *ebiten.Image) {
	for _, entity := range s.System.entities {
		offsetBegin := s.ViewportComponent.Viewport.DX * s.ViewportComponent.Viewport.DY
		offsetEnd := (s.ViewportComponent.Viewport.DX + 640) * (s.ViewportComponent.Viewport.DY + 480)
		idx := (640 * entity.Position.YPos) + entity.Position.XPos
		if idx >= float64(offsetBegin) && idx <= float64(offsetEnd) {
			entity.Render.Options.GeoM.Translate(float64(-s.ViewportComponent.Viewport.DX), float64(-s.ViewportComponent.Viewport.DY))
			screen.DrawImage(entity.Render.Image, entity.Render.Options)
		}
	}
}
