package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type StopShape struct {
	ShapeDefault
}

func NewStopShape(x, y float32) *StopShape {
	return &StopShape{
		ShapeDefault: ShapeDefault{
			shapeType: STOP,
			name:      "Stop",
			content:   []string{},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    settings.SHAPE_MIN_HEIGHT,
			width:     settings.SHAPE_MIN_WIDTH,
			visible:   true,
			color:     settings.START_STOP_COLOR,
			fontColor: settings.FONT_COLOR,
			fontSize:  settings.FONT_SIZE,
		},
	}
}

func (shape *StopShape) drawShape(rect rl.Rectangle, color *rl.Color) {
	rl.DrawEllipse(
		int32(rect.X+rect.Width/2),
		int32(rect.Y+rect.Height/2),
		rect.Width/2,
		rect.Height/2,
		*color,
	)
}

func (shape *StopShape) Draw() {
	shape.drawShape(shape.GetRect(), &shape.color)
	shape.drawContent()
}

func (shape *StopShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}
