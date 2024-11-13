package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type IfShape struct {
	ShapeDefault
}

func NewIfShape(x, y float32) *IfShape {
	return &IfShape{
		ShapeDefault: ShapeDefault{
			shapeType: IF,
			name:      "If",
			content:   []string{},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    settings.SHAPE_MIN_HEIGHT,
			width:     settings.SHAPE_MIN_WIDTH,
			visible:   true,
			color:     settings.IF_COLOR,
			fontColor: settings.FONT_COLOR,
			fontSize:  settings.FONT_SIZE,
		},
	}
}

func (shape *IfShape) drawShape(rect rl.Rectangle, color *rl.Color) {
	left := rl.NewVector2(rect.X, rect.Y+rect.Height/2)
	right := rl.NewVector2(rect.X+rect.Width, rect.Y+rect.Height/2)
	up := rl.NewVector2(rect.X+rect.Width/2, rect.Y)
	down := rl.NewVector2(rect.X+rect.Width/2, rect.Y+rect.Height)
	rl.DrawTriangle(up, left, down, *color)
	rl.DrawTriangle(right, up, down, *color)
}

func (shape *IfShape) Draw() {
	shape.updateSize()
	if shape.isHighlighted {
		shape.drawShape(shape.getHighlightRect(), &settings.HIGHLIGHT_COLOR)
	}
	shape.drawShape(shape.GetRect(), &shape.color)
	shape.drawContent()
}

// Always to right
func (shape *IfShape) GetOutPosTrue() (float32, float32) {
	return shape.x + shape.width, shape.y + shape.height/2
}

// Always to left
func (shape *IfShape) GetOutPosFalse() (float32, float32) {
	return shape.x, shape.y + shape.height/2
}

// Works always upon click within Rect so Right is always
// greater or equal than mousePol, similarly left is always
// less than mousePos
func (shape *IfShape) CloserToRight(mousePos rl.Vector2) bool {
	// dR = x + w - mx          - distance to Right
	// dL = mx - x              - distance to Left
	// dR < dL => dR - dL < 0
	// x + w - mx < mx - x => x + w - mx - mx + x < 0
	// 2x - 2mx + w < 0
	return 2*shape.x-2*mousePos.X+shape.width < 0
}

func (shape *IfShape) Info() string {
	retVal := `
Evaluated logic expression. 
Accepts '&&' and "||" operators as and, or respectively.
Doesn't evaluate 'and'/'or' keywords.

Examples:
  t[i] < d || t[5] > 2
  a < y
  x >= 8

Throws error if array or variable wasn't declared.
  `
	return retVal
}
