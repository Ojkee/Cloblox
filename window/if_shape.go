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
			x:         x,
			y:         y,
			height:    SHAPE_HEIGHT,
			width:     SHAPE_WIDTH,
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
		"If",
		int32(shape.x+shape.width/2),
		int32(shape.y+shape.height/2-8),
		shape.fontSize,
		shape.fontColor,
	)
}
