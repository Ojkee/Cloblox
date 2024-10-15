package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/graph"
)

type Window struct {
	name   string
	height int32
	width  int32

	backgroundColor rl.Color
	fontColor       rl.Color

	buildingShapes []Shape
	shapes         []Shape

	shapeClicked     bool
	currentShape     Shape
	currentShapeType SHAPE_TYPE

	diagram graph.Graph
}

func NewWindow(name string, height, width int32) *Window {
	return &Window{
		name:   name,
		height: height,
		width:  width,

		backgroundColor: BACKGROUND_COLOR,
		fontColor:       FONT_COLOR,

		buildingShapes: initBuildingShapes(width, height),
		shapes:         make([]Shape, 0),

		currentShape:     nil,
		shapeClicked:     false,
		currentShapeType: NONE,
	}
}

func initBuildingShapes(width, height int32) []Shape {
	offsetX := float32(width/2.0 + 10)
	gap := float32(SHAPE_HEIGHT + 16)
	offsetY := float32(height)/2.0 - gap*2.5
	return []Shape{
		NewStartShape(offsetX, offsetY),
		NewVariableShape(offsetX, offsetY+gap),
		NewIfShape(offsetX, offsetY+2*gap),
		NewPrintShape(offsetX, offsetY+3*gap),
		NewStopShape(offsetX, offsetY+4*gap),
	}
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "Cloblox")
	defer rl.CloseWindow()
	rl.SetTargetFPS(165)

	// FONT = rl.LoadFont("fonts/Metropolis-Medium.otf")
	for !(rl.WindowShouldClose() || rl.IsKeyPressed(rl.KeyQ)) {
		window.checkEvent()
		window.draw()
	}
}

func (window *Window) checkEvent() {
	mousePos := rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		clickedNewShape := false
		for _, shape := range window.buildingShapes {
			if rl.CheckCollisionPointRec(mousePos, shape.GetRect()) {
				clickedNewShape = true
				window.shapeClicked = true
				window.makeCurrentClicked(shape.GetType())
				window.updateCurrent()
			}
		}
		if mousePos.X < WINDOW_WIDTH/2 {
			window.resetClickedShape()
		}
		if window.shapeClicked && !clickedNewShape {
			window.placeCurrent(mousePos.X, mousePos.Y)
			window.resetClickedShape()
		}
	}
}

func (window *Window) placeCurrent(mx, my float32) {
	var s Shape
	switch window.currentShapeType {
	case START:
		s = NewStartShape(mx, my)
		break
	case VARIABLE:
		s = NewVariableShape(mx, my)
		break
	case IF:
		s = NewIfShape(mx, my)
		break
	case PRINT:
		s = NewPrintShape(mx, my)
		break
	case STOP:
		s = NewStopShape(mx, my)
		break
	default:
		window.resetClickedShape()
		panic("window.go/makeCurrentClicked fail:\n\tNot implemented shape type")
	}
	s.TranslateCenter()
	window.shapes = append(window.shapes, s)
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.MaximizeWindow()

	rl.ClearBackground(window.backgroundColor)

	rl.DrawLine(window.width/2, 0, window.width/2, window.height, FONT_COLOR)
	for _, shape := range window.buildingShapes {
		shape.Draw()
	}
	for _, shape := range window.shapes {
		shape.Draw()
	}
	if window.currentShape != nil {
		window.updateCurrent()
		window.currentShape.Draw()
	}

	rl.EndDrawing()
}

func (window *Window) makeCurrentClicked(shapeType SHAPE_TYPE) {
	window.currentShapeType = shapeType
	switch shapeType {
	case START:
		window.currentShape = NewStartShape(0, 0)
		break
	case VARIABLE:
		window.currentShape = NewVariableShape(0, 0)
		break
	case IF:
		window.currentShape = NewIfShape(0, 0)
		break
	case PRINT:
		window.currentShape = NewPrintShape(0, 0)
		break
	case STOP:
		window.currentShape = NewStopShape(0, 0)
		break
	default:
		window.resetClickedShape()
		panic("window.go/makeCurrentClicked fail:\n\tNot implemented shape type")
	}
}

func (window *Window) updateCurrent() {
	mousePos := rl.GetMousePosition()
	window.currentShape.MoveTo(
		mousePos.X-SHAPE_WIDTH/2,
		mousePos.Y-SHAPE_HEIGHT/2,
	)
}

func (window *Window) resetClickedShape() {
	window.shapeClicked = false
	window.currentShape = nil
	window.currentShapeType = NONE
}
