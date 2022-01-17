package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"2d-grass/preload"
	"2d-grass/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	tickNum      int
	CitySystem   *systems.BuildCitySystem
	PlayerSystem *systems.PlayerSystem
	BgSystem     *systems.BackgroundSystem
}

func (g *Game) ConvertCoordToTileIdx(xPos, yPos float64) (int, int) {
	return int(math.Floor(xPos)) / 16, int(math.Floor(yPos)) / 16
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.BgSystem.Draw(screen)
	g.CitySystem.Draw(screen)
	g.PlayerSystem.Draw(screen)
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

func (g *Game) Preload() {
	preload.New()
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Print(err)
	}
	preload.Set("player", playerImg)

	cityImg, _, err := ebitenutil.NewImageFromFile("assets/fire_0.png")
	if err != nil {
		log.Print(err)
	}
	preload.Set("city", cityImg)

	bgImg, _, err := ebitenutil.NewImageFromFile("assets/grass_random_grid.png")
	if err != nil {
		log.Print(err)
	}
	preload.Set("bg", bgImg)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello Folks")
	game := &Game{tickNum: 0}
	game.Preload()

	game.BgSystem = systems.NewBackgroundSystem()
	game.BgSystem.Init()

	game.CitySystem = systems.NewBuildCitySystem()

	game.PlayerSystem = systems.NewPlayerSystem()
	game.PlayerSystem.Init()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
