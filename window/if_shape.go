package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type IfShape struct {
	ShapeDefault
}

func NewIfShape(x, y float32) *IfShape {
	return &IfShape{
		ShapeDefault: ShapeDefault{
			shapeType: IF,
			content:   []string{"If"},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    SHAPE_HEIGHT,
			width:     SHAPE_WIDTH,
			visible:   true,
			color:     IF_COLOR,
			fontColor: FONT_COLOR,
			fontSize:  FONT_SIZE,
		},
	}
}

func (shape *IfShape) Draw() {
	left := rl.NewVector2(shape.x, shape.y+shape.height/2)
	right := rl.NewVector2(shape.x+shape.width, shape.y+shape.height/2)
	up := rl.NewVector2(shape.x+shape.width/2, shape.y)
	down := rl.NewVector2(shape.x+shape.width/2, shape.y+shape.height)
	rl.DrawTriangle(up, left, down, shape.color)
	rl.DrawTriangle(right, up, down, shape.color)
	rl.DrawText(
		shape.content[0],
		int32(shape.x+shape.width/2),
		int32(shape.y+shape.height/2-8),
		shape.fontSize,
		shape.fontColor,
	)
}

// Always to right
func (shape *IfShape) GetOutPosTrue() (float32, float32) {
	return shape.x + shape.width, shape.y + shape.height/2
}

// Always to left
func (shape *IfShape) GetOutPosFalse() (float32, float32) {
	return shape.x, shape.y + shape.height/2
}

func (shape *IfShape) CloserToRight(mousePos rl.Vector2) bool {
	// Works always upon click within Rect so Right is always
	// greater or equal than mousePol, similarly left is always
	// less than mousePos
	// dR = x + w - mx          - distance to Right
	// dL = mx - x              - distance to Left
	// dR < dL => dR - dL < 0
	// x + w - mx < mx - x => x + w - mx - mx + x < 0
	// 2x - 2mx + w < 0
	return 2*shape.x-2*mousePos.X+shape.width < 0
}
