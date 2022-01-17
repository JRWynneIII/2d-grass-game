package preload

import "github.com/hajimehoshi/ebiten/v2"

var preloadedData map[string]*ebiten.Image

func Get(key string) *ebiten.Image {
	return preloadedData[key]
}

func Set(key string, img *ebiten.Image) {
	preloadedData[key] = img
}

func GetAllData() map[string]*ebiten.Image {
	return preloadedData
}

func New() {
	preloadedData = make(map[string]*ebiten.Image)
}
