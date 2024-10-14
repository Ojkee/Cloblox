package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PrintShape struct {
	ShapeDefault
	tiltOffset float32
}

func NewPrintShape(x, y float32) *PrintShape {
	return &PrintShape{
		ShapeDefault: ShapeDefault{
			x:         x,
			y:         y,
			height:    SHAPE_HEIGHT,
			width:     SHAPE_WIDTH - SHAPE_WIDTH/4,
			color:     PRINT_COLOR,
			fontColor: FONT_COLOR,
			fontSize:  FONT_SIZE,
		},
		tiltOffset: SHAPE_WIDTH / 4,
	}
}

func (shape *PrintShape) Draw() {
	left_down := rl.NewVector2(shape.x, shape.y+shape.height)
	right_down := rl.NewVector2(shape.x+shape.width, shape.y+shape.height)
	left_up := rl.NewVector2(shape.x+shape.tiltOffset, shape.y)
	right_up := rl.NewVector2(shape.x+shape.width+shape.tiltOffset, shape.y)
	rl.DrawTriangle(left_up, right_down, right_up, shape.color)
	rl.DrawTriangle(right_down, left_up, left_down, shape.color)
	rl.DrawText(
		"Print",
		int32(shape.x+shape.width/2-32),
		int32(shape.y+shape.height/2-8),
		shape.fontSize,
		shape.fontColor,
	)
}
