package util

import "math"

func ConvertCoordToTileIdx(xPos, yPos float64, tileSize ...int) (int, int) {
	size := 16
	if len(tileSize) > 0 {
		size = tileSize[0]
	}
	return int(math.Floor(xPos)) / 16, int(math.Floor(yPos)) / 16
}
