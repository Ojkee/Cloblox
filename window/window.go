package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/functools"
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

	errorManager functools.ErrorManager

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
	simulationMode       SIMULATE_MODE
	simulationVarButtons []VarButton
	simulateModeButton   []SimulateButton
	simulationVar        string
	consoleLines         []ConsoleLine
	consoleInnerRect     rl.Rectangle

	connections       []shapes.Connection
	clickedConnection bool
	currentConnection *shapes.Connection

	// FUNC BUTTONS
	saveStateButton FuncButton
	saveCodeButton  FuncButton

	diagram graph.Graph
}

func NewWindow(name string, height, width int32) *Window {
	return &Window{
		name:   name,
		height: height,
		width:  width,

		currentMode: BUILD,

		errorManager: *functools.NewErrorManager(nil),

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

		simulationMode:       NOT_SELECTED,
		simulationVarButtons: make([]VarButton, 0),
		simulateModeButton: []SimulateButton{
			*NewSimulateButton(
				rl.NewRectangle(
					settings.SIMULATE_BUTTON_POS_X,
					settings.SIMULATE_BUTTON_POS_Y,
					settings.SIMULATE_BUTTON_WIDTH,
					settings.SIMULATE_BUTTON_HEIGHT,
				),
				STEP_BY_STEP,
			),
			*NewSimulateButton(
				rl.NewRectangle(
					settings.SIMULATE_BUTTON_POS_X+settings.SIMULATE_BUTTON_WIDTH+10,
					settings.SIMULATE_BUTTON_POS_Y,
					settings.SIMULATE_BUTTON_WIDTH,
					settings.SIMULATE_BUTTON_HEIGHT,
				),
				CONTINUOUSLY,
			),
			*NewSimulateButton(
				rl.NewRectangle(
					settings.SIMULATE_BUTTON_POS_X+2*(settings.SIMULATE_BUTTON_WIDTH+10),
					settings.SIMULATE_BUTTON_POS_Y,
					settings.SIMULATE_BUTTON_WIDTH,
					settings.SIMULATE_BUTTON_HEIGHT,
				),
				PAUSE,
			),
		},
		simulationVar: "",
		consoleLines:  make([]ConsoleLine, 0),
		consoleInnerRect: rl.NewRectangle(
			settings.CONSOLE_MARGIN+settings.CONSOLE_BORDER_WIDTH,
			float32(settings.WINDOW_HEIGHT-settings.CONSOLE_HEIGHT)+settings.CONSOLE_BORDER_WIDTH,
			settings.WINDOW_WIDTH/2-2*(settings.CONSOLE_MARGIN+settings.CONSOLE_BORDER_WIDTH),
			settings.CONSOLE_HEIGHT-2*settings.CONSOLE_BORDER_WIDTH-settings.CONSOLE_MARGIN,
		), // TODO CLEAN

		connections:       make([]shapes.Connection, 0),
		clickedConnection: false,
		currentConnection: nil,

		saveStateButton: *NewFuncButton(
			"save blocks",
			rl.NewRectangle(
				settings.SIMULATE_BUTTON_OFFSET,
				settings.SIMULATE_BUTTON_POS_Y,
				settings.SIMULATE_BUTTON_WIDTH,
				settings.SIMULATE_BUTTON_HEIGHT,
			),
		),
		saveCodeButton: *NewFuncButton(
			"save as python",
			rl.NewRectangle(
				settings.SIMULATE_BUTTON_WIDTH+settings.SIMULATE_BUTTON_OFFSET+settings.SIMULATE_BUTTON_GAP,
				settings.SIMULATE_BUTTON_POS_Y,
				settings.SIMULATE_BUTTON_WIDTH,
				settings.SIMULATE_BUTTON_HEIGHT,
			),
		),

		diagram: *graph.NewGraph(nil),
	}
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "Cloblox")
	rl.SetExitKey(rl.KeyNull)
	defer rl.CloseWindow()
	rl.SetTargetFPS(165)
	settings.FONT = rl.LoadFont(settings.FONT_PATH)

	for !rl.WindowShouldClose() {
		window.checkEvent()
		window.draw()

		if rl.IsKeyPressed(rl.KeyF4) { // Debug
			if settings.DEBUG_SHAPE_CONTENT {
				window.debugContent()
			}
			if settings.DEBUG_BLOCKS_POINTERS {
				window.diagram.DebugLog()
			}
			if settings.DEBUG_DIAGRAM_DETAILS {
				window.diagram.DebugDiagramDetails()
			}
			if settings.DEBUG_ERRORS {
				window.errorManager.PrintAllErrors()
			}
		}
	}
	rl.UnloadFont(settings.FONT)
	rl.CloseWindow()
}

func (window *Window) checkEvent() {
	window.changeModeEvent()
	mousePos := rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		window.selectInsertShape(&mousePos)
	}
	var errs []error
	switch window.currentMode {
	case BUILD:
		errs = window.buildManager(&mousePos)
		break
	case INSERT:
		errs = window.insertManager(&mousePos)
		break
	case REMOVE:
		errs = window.removeManager(&mousePos)
		break
	case SIMULATE:
		errs = window.simulateManager(&mousePos)
		break
	}
	window.appendErrorsToConsole(errs)
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(window.backgroundColor)
	rl.DrawLine(window.width/2, 0, window.width/2, window.height, settings.FONT_COLOR)
	window.drawCurrentMode()

	mousePos := rl.GetMousePosition()
	if window.currentMode != SIMULATE {
		window.drawHelp(&mousePos)
	} else if window.currentMode == SIMULATE {
		window.drawConsole()
		window.drawCurrentSlice()
		window.drawAllSlicesButtons()
		for _, sb := range window.simulateModeButton {
			window.drawSimulateButton(&sb)
		}
		window.saveStateButton.Draw()
		window.saveCodeButton.Draw()
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
