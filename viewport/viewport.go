package viewport

import "log"

const (
	Left = iota
	Down
	Right
	Up
)

type Viewport struct {
	DX, DY int
}

func (v *Viewport) Move(direction int) {
	switch direction {
	case Left:
		v.DX--
	case Right:
		v.DX++
	case Up:
		v.DY--
	case Down:
		v.DY++
	default:
		log.Print("Invalid direction")
	}
}

func (v *Viewport) New() {
	v.DX = 0
	v.DY = 0
}
