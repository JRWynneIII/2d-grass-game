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

	BgSystem *systems.BackgroundSystem
}

func (g *Game) ConvertCoordToTileIdx(xPos, yPos float64) (int, int) {
	return int(math.Floor(xPos)) / 16, int(math.Floor(yPos)) / 16
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.BgSystem.Draw(screen)
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
	game.BgSystem = systems.NewBackgroundSystem()
	game.BgSystem.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
