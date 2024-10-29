package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SHAPE_TYPE uint

const (
	NONE SHAPE_TYPE = iota
	START
	VARIABLE
	IF
	ACTION
	STOP
)

type Shape interface {
	Draw()
	GetType() SHAPE_TYPE
	GetRect() rl.Rectangle
	MoveTo(x, y float32)
	MoveToCenter(x, y float32)
	TranslateCenter()
	Resize(height, width float32)
	SetContent(content *[]string)
	GetInPos() (float32, float32)
	SetBlockId(bId int)
	GetBlockId() int
}

type ShapeSingleOut interface {
	GetOutPos() (float32, float32)
}
type ShapeManyOut interface {
	GetOutPosTrue() (float32, float32)
	GetOutPosFalse() (float32, float32)
	CloserToRight(mousePos rl.Vector2) bool
}

type ShapeDefault struct {
	shapeType SHAPE_TYPE
	content   []string
	blockID   int

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

func (shape *ShapeDefault) MoveToCenter(x, y float32) {
	shape.x = x - shape.width/2
	shape.y = y - shape.height/2
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

func (shape *ShapeDefault) GetInPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y
}

func (shape *ShapeDefault) getContentSize(idx int) int32 {
	return rl.MeasureText(shape.content[idx], shape.fontSize)
}

func (shape *ShapeDefault) SetBlockId(bId int) {
	shape.blockID = bId
}

func (shape *ShapeDefault) GetBlockId() int {
	return shape.blockID
}
