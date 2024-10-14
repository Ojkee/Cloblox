package window

import rl "github.com/gen2brain/raylib-go/raylib"

type ShapeDrawable interface {
	Draw()
}

type ShapeDefault struct {
	y      float32
	x      float32
	width  float32
	height float32

	color     rl.Color
	fontColor rl.Color
	fontSize  int32
}

func (s *ShapeDefault) Move(x, y float32) {
	s.x = x
	s.y = y
}

func (s *ShapeDefault) Resize(height, width float32) {
	s.height = height
	s.width = width
}
