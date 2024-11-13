package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/functools"
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
	fontSizeFloat := float32(shape.fontSize)
	if shape.IsContentEmpty() {
		nameWidth := rl.MeasureText(shape.name, shape.fontSize)
		xPos := float32(shape.x + shape.width/2 - float32(nameWidth)/2)
		yPos := float32(shape.y + shape.height/2 - float32(shape.fontSize)/2)
		posVec := rl.NewVector2(xPos, yPos)
		rl.DrawTextEx(
			settings.FONT,
			shape.name,
			posVec,
			fontSizeFloat,
			settings.FONT_SPACING,
			shape.fontColor,
		)
	} else {
		offset := float32(shape.fontSize)
		for i, contentLine := range shape.content {
			contentWidth := functools.TextWidthEx(shape.content[i]).X
			xPos := float32(shape.x + shape.width/2 - contentWidth/2)
			yPos := float32(shape.y + offset*float32(i))
			posVec := rl.NewVector2(xPos, yPos)
			rl.DrawTextEx(
				settings.FONT,
				contentLine,
				posVec,
				fontSizeFloat,
				settings.FONT_SPACING,
				shape.fontColor,
			)
		}
	}
}

func (shape *ShapeDefault) updateSize() {
	maxWidth := settings.SHAPE_MIN_WIDTH
	for i := range shape.content {
		textWidth := functools.TextWidthEx(shape.content[i]).X
		maxWidth = max(maxWidth, textWidth+2*settings.MARGIN_HORIZONTAL)
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
