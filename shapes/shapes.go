package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
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
	GetContent() []string
	IsContentEmpty() bool
	GetContentSize(idx int) int32
	GetInPos() (float32, float32)
	SetBlockId(bId int)
	GetBlockId() int
	GetColor() rl.Color
	SetHighlight(flag bool)
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
	name      string
	content   []string
	blockID   int

	y             float32
	x             float32
	width         float32
	height        float32
	isHighlighted bool

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

func (shape *ShapeDefault) GetContent() []string {
	return shape.content
}

func (shape *ShapeDefault) GetInPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y
}

func (shape *ShapeDefault) GetContentSize(idx int) int32 {
	return rl.MeasureText(shape.content[idx], shape.fontSize)
}

func (shape *ShapeDefault) SetBlockId(bId int) {
	shape.blockID = bId
}

func (shape *ShapeDefault) GetBlockId() int {
	return shape.blockID
}

func (shape *ShapeDefault) IsContentEmpty() bool {
	if len(shape.content) > 0 {
		return false
	}
	return true
}

func (shape *ShapeDefault) drawContent() {
	if shape.IsContentEmpty() {
		nameWidth := rl.MeasureText(shape.name, shape.fontSize)
		xPos := int32(shape.x + shape.width/2 - float32(nameWidth)/2)
		yPos := int32(shape.y + shape.height/2 - float32(settings.FONT_SIZE)/2)
		rl.DrawText(shape.name, xPos, yPos, shape.fontSize, shape.fontColor)
	} else {
		offset := int(settings.FONT_SIZE)
		for i, contentLine := range shape.content {
			contentWidth := shape.GetContentSize(i)
			xPos := int32(shape.x + shape.width/2 - float32(contentWidth)/2)
			yPos := int32(shape.y + float32(offset*i))
			rl.DrawText(contentLine, xPos, yPos, shape.fontSize, shape.fontColor)
		}
	}
}

func (shape *ShapeDefault) updateSize() {
	maxWidth := settings.SHAPE_MIN_WIDTH
	for i := range shape.content {
		maxWidth = max(maxWidth, float32(shape.GetContentSize(i))+2*settings.MARGIN_HORIZONTAL)
	}
	shape.width = maxWidth
	shape.height = max(
		settings.SHAPE_MIN_HEIGHT,
		float32(len(shape.content))*float32(settings.FONT_SIZE)+settings.MARGIN_VERTICAL/2,
	)
}

func (shape *ShapeDefault) GetColor() rl.Color {
	return shape.color
}

func (shape *ShapeDefault) SetHighlight(flag bool) {
	shape.isHighlighted = flag
}

func (shape *ShapeDefault) getHighlightRect() rl.Rectangle {
	hx := shape.x - settings.HIGHLIGHT_PAD
	hy := shape.y - settings.HIGHLIGHT_PAD
	hwidth := shape.width + 2*settings.HIGHLIGHT_PAD
	hheight := shape.height + 2*settings.HIGHLIGHT_PAD
	return rl.NewRectangle(hx, hy, hwidth, hheight)
}
