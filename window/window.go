package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/graph"
	"Cloblox/settings"
	"Cloblox/shapes"
)

type Window struct {
	name            string
	height          int32
	width           int32
	backgroundColor rl.Color
	fontColor       rl.Color

	currentMode MODE

	// BUILD
	buildingShapes   []shapes.Shape
	diagramShapes    []shapes.Shape
	shapeClicked     bool
	currentShape     shapes.Shape
	currentShapeType shapes.SHAPE_TYPE

	// INSERT
	currentInsertShape *shapes.Shape
	insertCursorX      int
	insertCursorY      int

	// SIMULATE
	simulationStarted bool
	simulationMode    SIMULATE_MODE
	simulationVar     string

	connections       []Connection
	clickedConnection bool
	currentConnection *Connection

	diagram graph.Graph
}

func NewWindow(name string, height, width int32) *Window {
	return &Window{
		name:   name,
		height: height,
		width:  width,

		currentMode: BUILD,

		backgroundColor: settings.BACKGROUND_COLOR,
		fontColor:       settings.FONT_COLOR,

		buildingShapes:   initBuildingShapes(width, height),
		diagramShapes:    make([]shapes.Shape, 0),
		shapeClicked:     false,
		currentShape:     nil,
		currentShapeType: shapes.NONE,

		currentInsertShape: nil,
		insertCursorX:      -1,
		insertCursorY:      -1,

		simulationStarted: false,
		simulationMode:    NOT_SELECTED,

		connections:       make([]Connection, 0),
		clickedConnection: false,
		currentConnection: nil,

		diagram: *graph.NewGraph(nil),
	}
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "Cloblox")
	rl.SetExitKey(rl.KeyNull)
	defer rl.CloseWindow()
	rl.SetTargetFPS(165)
	settings.FONT = rl.LoadFont(settings.FONT_PATH)

	// TODO REMOVE =======================================================================================
	window.diagram.SetAllVars(map[string]any{
		"x": 3,
		"y": 4.5,
		"p": []float64{3, 2, 1, 5, 3, 0, 10, 25, 13, 3, 11},
		"n": []float64{0, -10, -15, -10, -9, -3, -3, -6, -33, -23},
		"b": []float64{3, 5, 1, -1, -5, 0, 22, -25, 20, 5, -15},
	})
	window.simulationVar = "p"
	window.simulationStarted = true
	// END REMOVE ========================================================================================

	for !rl.WindowShouldClose() {
		window.checkEvent()
		window.draw()

		if rl.IsKeyPressed(rl.KeyF4) { // Debug
			if settings.DEBUG_SHAPE_CONTENT {
				window.debugContent()
			}
			if settings.DEBUG_BLOCKS_POINTERS {
				window.diagram.Log()
			}
		}
	}
}

func (window *Window) checkEvent() {
	window.changeModeEvent()
	mousePos := rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		window.selectInsertShape(&mousePos)
	}
	switch window.currentMode {
	case BUILD:
		window.buildManager(&mousePos)
		break
	case INSERT:
		window.insertManager(&mousePos)
		break
	case REMOVE:
		window.removeManager(&mousePos)
		break
	case SIMULATE:
		window.simulateManager(&mousePos)
		break
	}
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(window.backgroundColor)
	rl.DrawLine(window.width/2, 0, window.width/2, window.height, settings.FONT_COLOR)
	window.drawCurrentMode()

	mousePos := rl.GetMousePosition()
	if window.currentMode != SIMULATE {
		window.drawHelp(&mousePos)
	} else if window.currentMode == SIMULATE && !window.simulationStarted {
		window.drawAllSlicesButtons()
	} else if window.currentMode == SIMULATE {
		window.drawCurrentSlice()
	}
	for _, conn := range window.connections {
		conn.Draw()
	}
	for _, shape := range window.buildingShapes {
		shape.Draw()
	}
	for _, shape := range window.diagramShapes {
		shape.Draw()
	}

	if window.currentShape != nil {
		window.updateCurrentShape(&mousePos)
		window.currentShape.Draw()
	}
	if window.currentConnection != nil {
		window.updateCurrentConnection(&mousePos)
		window.currentConnection.Draw()
	}
	if window.currentInsertShape != nil && window.currentMode == INSERT {
		window.drawCursor()
	}

	rl.EndDrawing()
}
