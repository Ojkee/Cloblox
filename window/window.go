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
	currentShape   Shape
	shapeClicked   bool

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
		currentShape:   nil,
	}
}

func initBuildingShapes(width, height int32) []Shape {
	offsetX := float32(width/2 + 10)
	gap := float32(SHAPE_HEIGHT + 16)
	offsetY := float32(height)/2.0 - gap*2.5
	buildingShapes := make([]Shape, 0)
	buildingShapes = append(buildingShapes, NewStartShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewVariableShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewIfShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewPrintShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewStopShape(offsetX, offsetY))
	return buildingShapes
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "Cloblox")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !(rl.WindowShouldClose() || rl.IsKeyPressed(rl.KeyQ)) {
		window.checkEvent()
		window.draw()
	}
}

func (window *Window) checkEvent() {
	mousePos := rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		for _, shape := range window.buildingShapes {
			if rl.CheckCollisionPointRec(mousePos, shape.GetRect()) {
				window.shapeClicked = true
				window.makeCurrentClicked(shape.GetType())
			}
		}
	}
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.MaximizeWindow()

	rl.ClearBackground(window.backgroundColor)

	rl.DrawLine(window.width/2, 0, window.width/2, window.height, rl.NewColor(255, 248, 231, 255))
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
	window.currentShape = nil
	window.shapeClicked = false
}
