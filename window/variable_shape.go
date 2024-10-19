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
			content:   []string{"Variable"},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    SHAPE_HEIGHT,
			width:     SHAPE_WIDTH,
			visible:   true,
			color:     VARIABLE_COLOR,
			fontColor: FONT_COLOR,
			fontSize:  FONT_SIZE,
		},
	}
}

func (shape *VariableShape) Draw() {
	nLines := len(shape.content)
	if nLines == 0 {
		nLines = 1
	}
	shapeHeight := int32(shape.height+SHAPE_TEXT_GAP) * int32(nLines)
	rl.DrawRectangle(
		int32(shape.x),
		int32(shape.y),
		int32(shape.width),
		shapeHeight,
		shape.color,
	)
	for i, text := range shape.content {
		textWidth := float32(shape.getContentSize(i))
		rl.DrawText(
			text,
			int32(shape.x+(shape.width-textWidth)/2.0),
			int32(shape.y+float32(i*int(FONT_SIZE))),
			shape.fontSize,
			shape.fontColor,
		)
	}
}

func (shape *VariableShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}
