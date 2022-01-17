package systems

import (
	"2d-grass/entity"
	"2d-grass/preload"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type BuildCitySystem struct {
	System
}

func NewBuildCitySystem() *BuildCitySystem {
	return &BuildCitySystem{}
}

func (s *BuildCitySystem) Update() {
	xmin, xmax := 0, 640
	ymin, ymax := 0, 480
	x := rand.Intn(xmax-xmin) + xmin
	y := rand.Intn(ymax-ymin) + ymin
	log.Print("New City at ", x, ", ", y)
	//TODO: Change this to use some kind of size component or something
	s.System.AddEntity(entity.NewCityEntity(&entity.CityEntity{
		float64(x),
		float64(y),
		32,
		32,
		preload.Get("city"),
	}))
}

func (s *BuildCitySystem) Draw(screen *ebiten.Image) {
	for _, entity := range s.System.entities {
		op := &ebiten.DrawImageOptions{}
		//TODO: Change this to use a ScaleComponent or something  like that to scale it
		op.GeoM.Scale(0.25, 0.25)
		op.GeoM.Translate(entity.Position.XPos, entity.Position.YPos)
		screen.DrawImage(entity.Render.Image, op)
	}
}
