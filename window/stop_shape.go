package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type StopShape struct {
	ShapeDefault
}

func NewStopShape(x, y float32) *StopShape {
	return &StopShape{
		ShapeDefault: ShapeDefault{
			shapeType: STOP,
			x:         x,
			y:         y,
			height:    SHAPE_HEIGHT,
			width:     SHAPE_WIDTH,
			color:     START_STOP_COLOR,
			fontColor: FONT_COLOR,
			fontSize:  FONT_SIZE,
		},
	}
}

func (shape *StopShape) Draw() {
	rl.DrawEllipse(
		int32(shape.x+shape.width/2),
		int32(shape.y+shape.height/2),
		shape.width/2,
		shape.height/2,
		shape.color)
	rl.DrawText(
		"Stop",
		int32(shape.x+shape.width/2-32),
		int32(shape.y+shape.height/2-8),
		shape.fontSize,
		shape.fontColor,
	)
}
