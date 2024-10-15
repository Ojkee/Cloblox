package window

import rl "github.com/gen2brain/raylib-go/raylib"

type SHAPE_TYPE int

const (
	NONE SHAPE_TYPE = iota
	START
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
	TranslateCenter()
	Resize(height, width float32)
	SetContent(content *[]string)
}

type ShapeDefault struct {
	shapeType SHAPE_TYPE
	content   []string

	y      float32
	x      float32
	width  float32
	height float32

	visible   bool
	color     rl.Color
	fontColor rl.Color
	fontSize  int32
}

func (shape *ShapeDefault) GetType() SHAPE_TYPE {
	return shape.shapeType
}

func (shape *ShapeDefault) GetRect() rl.Rectangle {
	return rl.NewRectangle(shape.x, shape.y, shape.width, shape.height)
}

func (shape *ShapeDefault) MoveTo(x, y float32) {
	shape.x = x
	shape.y = y
}

func (shape *ShapeDefault) TranslateCenter() {
	shape.x -= shape.width / 2
	shape.y -= shape.height / 2
}

func (shape *ShapeDefault) Resize(height, width float32) {
	shape.height = height
	shape.width = width
}

func (shape *ShapeDefault) SetContent(content *[]string) {
	shape.content = *content
}

func (shape *ShapeDefault) getContentSize(idx int) int32 {
	return rl.MeasureText(shape.content[idx], shape.fontSize)
}
