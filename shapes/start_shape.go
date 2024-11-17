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
			name:      "Start",
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

func (shape *StartShape) drawShape(rect rl.Rectangle, color *rl.Color) {
	rl.DrawEllipse(
		int32(rect.X+rect.Width/2),
		int32(rect.Y+rect.Height/2),
		rect.Width/2,
		rect.Height/2,
		*color,
	)
}

func (shape *StartShape) Draw() {
	if shape.isHighlighted {
		shape.drawShape(shape.getHighlightRect(), &settings.HIGHLIGHT_COLOR)
	}
	shape.drawShape(shape.GetRect(), &shape.color)
	shape.drawContent()
}

func (shape *StartShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}

func (shape *StartShape) Info() string {
	retVal := `
  Every diagram must begin with this block.
  There can't be more than one start block.
  `
	return retVal
}
