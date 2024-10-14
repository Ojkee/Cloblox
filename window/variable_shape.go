package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type VariableShape struct {
	ShapeDefault
}

func NewVariableShape(x, y float32) *VariableShape {
	return &VariableShape{
		ShapeDefault: ShapeDefault{
			shapeType: VARIABLE,
			x:         x,
			y:         y,
			height:    SHAPE_HEIGHT,
			width:     SHAPE_WIDTH,
			color:     VARIABLE_COLOR,
			fontColor: FONT_COLOR,
			fontSize:  FONT_SIZE,
		},
	}
}

func (shape *VariableShape) Draw() {
	rl.DrawRectangle(
		int32(shape.x),
		int32(shape.y),
		int32(shape.width),
		int32(shape.height),
		shape.color,
	)
	rl.DrawText(
		"Variables",
		int32(shape.x+shape.width/2-32),
		int32(shape.y+shape.height/2-8),
		shape.fontSize,
		shape.fontColor,
	)
}
