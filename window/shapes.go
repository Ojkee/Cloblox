package window

import rl "github.com/gen2brain/raylib-go/raylib"

type SHAPE_TYPE int

const (
	START SHAPE_TYPE = iota
	VARIABLE
	IF
	PRINT
	STOP
)

type Shape interface {
	Draw()
	GetType() SHAPE_TYPE
	GetRect() rl.Rectangle
	MoveTo(x, y float32)
	Resize(height, width float32)
}

type ShapeDefault struct {
	shapeType SHAPE_TYPE

	y      float32
	x      float32
	width  float32
	height float32

	color     rl.Color
	fontColor rl.Color
	fontSize  int32
}

func (s *ShapeDefault) MoveTo(x, y float32) {
	s.x = x
	s.y = y
}

func (s *ShapeDefault) Resize(height, width float32) {
	s.height = height
	s.width = width
}

func (s *ShapeDefault) GetType() SHAPE_TYPE {
	return s.shapeType
}

func (s *ShapeDefault) GetRect() rl.Rectangle {
	return rl.NewRectangle(s.x, s.y, s.width, s.height)
}
