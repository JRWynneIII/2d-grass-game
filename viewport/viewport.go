package viewport

import "log"

const (
	Left = iota
	Down
	Right
	Up
)

type Viewport struct {
	X, Y int
}

func (v *Viewport) Move(direction int) {
	switch direction {
	case Left:
		v.X--
	case Right:
		v.X++
	case Up:
		v.Y--
	case Down:
		v.Y++
	default:
		log.Print("Invalid direction")
	}
}

func (v *Viewport) New() {
	v.X = 0
	v.Y = 0
}
