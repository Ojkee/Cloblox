package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type VariableShape struct {
	ShapeDefault
}

func NewVariableShape(x, y float32) *VariableShape {
	return &VariableShape{
		ShapeDefault: ShapeDefault{
			shapeType: VARIABLE,
			name:      "Variables",
			content:   []string{},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    settings.SHAPE_MIN_HEIGHT,
			width:     settings.SHAPE_MIN_WIDTH,
			visible:   true,
			color:     settings.VARIABLE_COLOR,
			fontColor: settings.FONT_COLOR,
			fontSize:  settings.FONT_SIZE,
		},
	}
}

func (shape *VariableShape) drawShape(rect rl.Rectangle, color *rl.Color) {
	rl.DrawRectangle(
		rect.ToInt32().X,
		rect.ToInt32().Y,
		rect.ToInt32().Width,
		rect.ToInt32().Height,
		*color,
	)
}

func (shape *VariableShape) Draw() {
	shape.updateSize()
	if shape.isHighlighted {
		shape.drawShape(shape.getHighlightRect(), &settings.HIGHLIGHT_COLOR)
	}
	shape.drawShape(shape.GetRect(), &shape.color)
	shape.drawContent()
}

func (shape *VariableShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}

func (shape *VariableShape) Info() string {
	retVal := `
  It's the place where variables are declared. 
  Examples:
    x = -1
    y = 2.5
    t = [4, 1, -2, 82, 9.2]
  `
	return retVal
}
