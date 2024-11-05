package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type StartShape struct {
	ShapeDefault
}

func NewStartShape(x, y float32) *StartShape {
	return &StartShape{
		ShapeDefault: ShapeDefault{
			shapeType: START,
			content:   []string{"Start"},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    settings.SHAPE_HEIGHT,
			width:     settings.SHAPE_WIDTH,
			visible:   true,
			color:     settings.START_STOP_COLOR,
			fontColor: settings.FONT_COLOR,
			fontSize:  settings.FONT_SIZE,
		},
	}
}

func (shape *StartShape) Draw() {
	rl.DrawEllipse(
		int32(shape.x+shape.width/2),
		int32(shape.y+shape.height/2),
		shape.width/2,
		shape.height/2,
		shape.color)
	contentWidth := float32(shape.getContentSize(0))
	rl.DrawText(
		shape.content[0],
		int32(shape.x+(shape.width-contentWidth)/2.0),
		int32(shape.y+shape.height/2-8),
		shape.fontSize,
		shape.fontColor,
	)
}

func (shape *StartShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}
