package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type ActionShape struct {
	ShapeDefault
	tiltOffset          float32
	tiltOffsetHighlight float32
}

func NewActionShape(x, y float32) *ActionShape {
	return &ActionShape{
		ShapeDefault: ShapeDefault{
			shapeType: ACTION,
			name:      "Action",
			content:   []string{},
			blockID:   -1,
			x:         x,
			y:         y,
			height:    settings.SHAPE_MIN_HEIGHT,
			width:     settings.SHAPE_MIN_WIDTH - settings.SHAPE_MIN_WIDTH/4,
			color:     settings.PRINT_COLOR,
			fontColor: settings.FONT_COLOR,
			fontSize:  settings.FONT_SIZE,
		},
		tiltOffset:          settings.SHAPE_MIN_WIDTH / 4,
		tiltOffsetHighlight: (settings.SHAPE_MIN_WIDTH + 2*settings.HIGHLIGHT_PAD) / 4,
	}
}

func (shape *ActionShape) drawShape(rect rl.Rectangle, tilt float32, color *rl.Color) {
	left_down := rl.NewVector2(rect.X, rect.Y+rect.Height)
	right_down := rl.NewVector2(rect.X+rect.Width, rect.Y+rect.Height)
	left_up := rl.NewVector2(rect.X+tilt, rect.Y)
	right_up := rl.NewVector2(rect.X+rect.Width+tilt, rect.Y)
	rl.DrawTriangle(left_up, right_down, right_up, *color)
	rl.DrawTriangle(right_down, left_up, left_down, *color)
}

func (shape *ActionShape) Draw() {
	shape.updateSize()
	if shape.isHighlighted {
		shape.drawShape(
			shape.getHighlightRect(),
			shape.tiltOffsetHighlight,
			&settings.HIGHLIGHT_COLOR,
		)
	}
	shape.drawShape(shape.GetRect(), shape.tiltOffset, &shape.color)
	shape.drawContent()
}

func (shape *ActionShape) GetOutPos() (float32, float32) {
	return shape.x + shape.width/2, shape.y + shape.height
}

func (shape *ActionShape) Info() string {
	retVal := `

  There is few types of action in this shape:
  
  - Print, prints variable to the debug console on the screen
  
  Example:
    print x
  
  - Math operations, performs operations with assignment,
  variable must be previously declared. 
  Examples:
    x++
    x--
    x += 3
    x -= x/2
    x /= t[i]*2
    x = y

  - Swap, swaps values of two variables 
  Example:
    swap t[i], x

  - Rand, randomize value in range with assignment.
  Rand is always floating number
  Example:
    x = rand 2, 5

  `
	return retVal
}
