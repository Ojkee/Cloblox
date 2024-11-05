package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type ActionShape struct {
	ShapeDefault
	tiltOffset float32
}

func NewActionShape(x, y float32) *ActionShape {
	return &ActionShape{
		ShapeDefault: ShapeDefault{
			shapeType: ACTION,
			content:   []string{},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    settings.SHAPE_HEIGHT,
			width:     settings.SHAPE_WIDTH - settings.SHAPE_WIDTH/4,
			color:     settings.PRINT_COLOR,
			fontColor: settings.FONT_COLOR,
			fontSize:  settings.FONT_SIZE,
		},
		tiltOffset: settings.SHAPE_WIDTH / 4,
	}
}

func (shape *ActionShape) Draw() {
	left_down := rl.NewVector2(shape.x, shape.y+shape.height)
	right_down := rl.NewVector2(shape.x+shape.width, shape.y+shape.height)
	left_up := rl.NewVector2(shape.x+shape.tiltOffset, shape.y)
	right_up := rl.NewVector2(shape.x+shape.width+shape.tiltOffset, shape.y)
	rl.DrawTriangle(left_up, right_down, right_up, shape.color)
	rl.DrawTriangle(right_down, left_up, left_down, shape.color)
	rl.DrawText(
		"Action",
		int32(shape.x+shape.width/2-32),
		int32(shape.y+shape.height/2-8),
		shape.fontSize,
		shape.fontColor,
	)
}

func (shape *ActionShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}