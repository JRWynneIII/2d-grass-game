package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"2d-grass/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	tickNum    int
	CitySystem *systems.BuildCitySystem
	BgTiles    map[*ebiten.Image]*ebiten.DrawImageOptions
}

func (g *Game) initBg() {
	for y := 0; y < 480; y += 64 {
		for x := 0; x < 640; x += 64 {
			img, _, err := ebitenutil.NewImageFromFile("assets/grass_random_grid.png")
			if err != nil {
				log.Print(err)
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			log.Print(fmt.Sprintf("tile at %d, %d", x, y))
			g.BgTiles[img] = op
		}
	}
}

func (g *Game) ConvertCoordToTileIdx(xPos, yPos float64) (int, int) {
	return int(math.Floor(xPos)) / 16, int(math.Floor(yPos)) / 16
}

func (g *Game) Draw(screen *ebiten.Image) {
	for img, op := range g.BgTiles {
		screen.DrawImage(img, op)
	}
	g.CitySystem.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
}

func (g *Game) Update() error {
	g.tickNum++
	if g.tickNum%1000 == 0 {
		g.CitySystem.Update()
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello Folks")
	game := &Game{tickNum: 0}
	img, _, err := ebitenutil.NewImageFromFile("assets_fire0.png")
	if err != nil {
		log.Print(err)
	}
	game.CitySystem = systems.NewBuildCitySystem(img)
	game.BgTiles = make(map[*ebiten.Image]*ebiten.DrawImageOptions)
	game.initBg()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
