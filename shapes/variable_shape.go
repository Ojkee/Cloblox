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
			height:    settings.SHAPE_HEIGHT,
			width:     settings.SHAPE_WIDTH,
			visible:   true,
			color:     settings.VARIABLE_COLOR,
			fontColor: settings.FONT_COLOR,
			fontSize:  settings.FONT_SIZE,
		},
	}
}

func (shape *VariableShape) Draw() {
	nLines := len(shape.content)
	if nLines == 0 {
		nLines = 1
	}
	shapeHeight := int32(shape.height+settings.SHAPE_TEXT_GAP) * int32(nLines)
	rl.DrawRectangle(
		int32(shape.x),
		int32(shape.y),
		int32(shape.width),
		shapeHeight,
		shape.color,
	)
	shape.drawContent()
}

func (shape *VariableShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}
