package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/graph"
	"Cloblox/settings"
	"Cloblox/shapes"
)

type Window struct {
	name   string
	height int32
	width  int32

	backgroundColor rl.Color
	fontColor       rl.Color

	buildingShapes   []shapes.Shape
	diagramShapes    []shapes.Shape
	shapeClicked     bool
	currentShape     shapes.Shape
	currentShapeType shapes.SHAPE_TYPE

	connections       []Connection
	clickedConnection bool
	currentConnection *Connection

	currentMode MODE

	diagram graph.Graph
}

func NewWindow(name string, height, width int32) *Window {
	return &Window{
		name:   name,
		height: height,
		width:  width,

		backgroundColor: settings.BACKGROUND_COLOR,
		fontColor:       settings.FONT_COLOR,

		buildingShapes: initBuildingShapes(width, height),
		diagramShapes:  make([]shapes.Shape, 0),

		currentShape:     nil,
		shapeClicked:     false,
		currentShapeType: shapes.NONE,

		connections:       make([]Connection, 0),
		clickedConnection: false,
		currentConnection: nil,

		currentMode: BUILDING,

		diagram: *graph.NewGraph(nil),
	}
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "Cloblox")
	rl.SetExitKey(rl.KeyQ)
	defer rl.CloseWindow()
	rl.SetTargetFPS(165)

	// FONT = rl.LoadFont("fonts/Metropolis-Medium.otf")
	for !rl.WindowShouldClose() {
		window.checkEvent()
		window.draw()

		if rl.IsKeyPressed(rl.KeyD) { // Debug
			window.diagram.Log()
		}
	}
}

func (window *Window) checkEvent() {
	window.changeModeEvent()
	mousePos := rl.GetMousePosition()
	switch window.currentMode {
	case BUILDING:
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) { // New Shape
			window.buildNewShapeEvent(&mousePos)
		} else if rl.IsMouseButtonPressed(rl.MouseButtonRight) { // Connect
			err := window.currentConnectionEvent(mousePos)
			if err != nil {
				fmt.Println(err) // TODO write to console
			}
		}
		break
	case INSERTION:
		break
	case REMOVE:
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) ||
			rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if removeId := window.getShapeIdMousePos(&mousePos); removeId != -1 {
				window.removeShapeAndConnectionsById(removeId)
			}
		}
		break
	}
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(window.backgroundColor)
	rl.DrawLine(window.width/2, 0, window.width/2, window.height, settings.FONT_COLOR)
	window.drawCurrentMode()

	for _, conn := range window.connections {
		conn.Draw()
	}
	for _, shape := range window.buildingShapes {
		shape.Draw()
	}
	for _, shape := range window.diagramShapes {
		shape.Draw()
	}

	mousePos := rl.GetMousePosition()
	if window.currentShape != nil {
		window.updateCurrentShape(&mousePos)
		window.currentShape.Draw()
	}
	if window.currentConnection != nil {
		window.updateCurrentConnection(&mousePos)
		window.currentConnection.Draw()
	}
	rl.EndDrawing()
}
