package component

import "github.com/hajimehoshi/ebiten/v2"

type IDComponent struct {
	Id int
}

func NewIDComponent(id int) *IDComponent {
	return &IDComponent{
		Id: id,
	}
}

type PositionComponent struct {
	XPos, YPos float64
}

func NewPositionComponent(x, y float64) *PositionComponent {
	return &PositionComponent{
		XPos: x,
		YPos: y,
	}
}

type SizeComponent struct {
	Height, Width float64
}

func NewSizeComponent(height, width float64) *SizeComponent {
	return &SizeComponent{
		Height: height,
		Width:  width,
	}
}

type RenderComponent struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
}

func NewRenderComponent(i *ebiten.Image) *RenderComponent {
	return &RenderComponent{
		Image: i,
	}
}

type Component struct {
	Name string
}

func NewComponent(name string) *Component {
	return &Component{
		Name: name,
	}
}
