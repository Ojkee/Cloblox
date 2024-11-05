package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/graph"
	"Cloblox/settings"
	"Cloblox/shapes"
)

type MODE string

const (
	BUILDING   MODE = "building"
	INSERTION       = "insertion"
	SIMULATION      = "simulation"
	REMOVE          = "remove"
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

func initBuildingShapes(width, height int32) []shapes.Shape {
	offsetX := float32(width/2.0 + 10)
	gap := float32(settings.SHAPE_HEIGHT + 16)
	offsetY := float32(height)/2.0 - gap*2.5
	return []shapes.Shape{
		shapes.NewStartShape(offsetX, offsetY),
		shapes.NewVariableShape(offsetX, offsetY+gap),
		shapes.NewIfShape(offsetX, offsetY+2*gap),
		shapes.NewActionShape(offsetX, offsetY+3*gap),
		shapes.NewStopShape(offsetX, offsetY+4*gap),
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
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) { // New Shape
		window.buildNewShapeEvent(&mousePos)
	} else if rl.IsMouseButtonPressed(rl.MouseButtonRight) { // Connect
		window.currentConnectionEvent(mousePos)
	}
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.MaximizeWindow()
	rl.ClearBackground(window.backgroundColor)
	rl.DrawLine(window.width/2, 0, window.width/2, window.height, settings.FONT_COLOR)

	window.drawMode()

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
