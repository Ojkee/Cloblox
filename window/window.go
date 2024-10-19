package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/blocks"
	"Cloblox/graph"
)

type Window struct {
	name   string
	height int32
	width  int32

	backgroundColor rl.Color
	fontColor       rl.Color

	buildingShapes   []Shape
	diagramShapes    []Shape
	shapeClicked     bool
	currentShape     Shape
	currentShapeType SHAPE_TYPE

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

		backgroundColor: BACKGROUND_COLOR,
		fontColor:       FONT_COLOR,

		buildingShapes: initBuildingShapes(width, height),
		diagramShapes:  make([]Shape, 0),

		currentShape:     nil,
		shapeClicked:     false,
		currentShapeType: NONE,

		connections:       make([]Connection, 0),
		clickedConnection: false,
		currentConnection: nil,

		diagram: *graph.NewGraph(nil),
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
		window.buildNewShapeEvent(&mousePos)
	} else if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		window.currentConnectionEvent(mousePos)
	}
}

func (window *Window) buildNewShapeEvent(mousePos *rl.Vector2) {
	clickedNewShape := false
	for _, shape := range window.buildingShapes {
		if rl.CheckCollisionPointRec(*mousePos, shape.GetRect()) {
			clickedNewShape = true
			window.shapeClicked = true
			window.makeCurrentClicked(shape.GetType())
			window.updateCurrentShape(mousePos)
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

func (window *Window) placeCurrent(mx, my float32) {
	var cShape Shape
	var cBlock blocks.Block
	switch window.currentShapeType {
	case START:
		cBlock = blocks.NewStartBlock()
		cShape = NewStartShape(mx, my)
		break
	case VARIABLE:
		cBlock = blocks.NewVariableBlock()
		cShape = NewVariableShape(mx, my)
		break
	case IF:
		cBlock = blocks.NewIfBlock()
		cShape = NewIfShape(mx, my)
		break
	case PRINT:
		cBlock = blocks.NewPrintBlock()
		cShape = NewPrintShape(mx, my)
		break
	case STOP:
		cBlock = blocks.NewStopBlock()
		cShape = NewStopShape(mx, my)
		break
	default:
		window.resetClickedShape()
		panic("window.go/makeCurrentClicked fail:\n\tNot implemented shape type")
	}
	window.diagram.AppendBlock(cBlock)
	cShape.TranslateCenter()
	cShape.SetBlockId(cBlock.GetId())
	window.diagramShapes = append(window.diagramShapes, cShape)
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.MaximizeWindow()

	rl.ClearBackground(window.backgroundColor)

	rl.DrawLine(window.width/2, 0, window.width/2, window.height, FONT_COLOR)
	for _, shape := range window.buildingShapes {
		shape.Draw()
	}
	for _, shape := range window.diagramShapes {
		shape.Draw()
	}

	for _, conn := range window.connections {
		conn.Draw()
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

func (window *Window) updateCurrentShape(mousePos *rl.Vector2) {
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

func (window *Window) currentConnectionEvent(mousePos rl.Vector2) {
	clickedAnyShape := false
	for _, shape := range window.diagramShapes {
		if rl.CheckCollisionPointRec(mousePos, shape.GetRect()) {
			clickedAnyShape = true
			if !window.clickedConnection {
				shapeX, shapeY := window.getShapeOutPos(&shape, &mousePos)
				window.currentConnection = NewConnection(
					shapeX, shapeY,
					mousePos.X, mousePos.Y,
					shape.GetBlockId(), -1,
				)
				window.clickedConnection = true
			} else {
				if !window.connectionExistsOrSelf(shape.GetBlockId(), window.currentConnection.inShapeId) {
					shapePosX, shapePosY := shape.GetInPos()
					window.currentConnection.MoveOutPos(shapePosX, shapePosY)
					window.currentConnection.outShapeId = shape.GetBlockId()
					window.connections = append(window.connections, *window.currentConnection)
					window.resetCurrentConnection()
				}
			}
		}
	}
	if !clickedAnyShape {
		window.resetCurrentConnection()
	}
}

func (window *Window) updateCurrentConnection(mousePos *rl.Vector2) {
	if window.clickedConnection {
		window.currentConnection.MoveOutPos(mousePos.X, mousePos.Y)
	}
}

func (window *Window) getShapeOutPos(shape *Shape, mousePos *rl.Vector2) (float32, float32) {
	var shapeX, shapeY float32
	if shapeManyOut, ok := (*shape).(ShapeManyOut); ok {
		if shapeManyOut.CloserToRight(*mousePos) {
			shapeX, shapeY = shapeManyOut.GetOutPosTrue()
		} else {
			shapeX, shapeY = shapeManyOut.GetOutPosFalse()
		}
	} else if shapeSingleOut, ok := (*shape).(ShapeSingleOut); ok {
		shapeX, shapeY = shapeSingleOut.GetOutPos()
	} else {
		panic("window/getShapeOutPos fail:\n\tNeither Single nor Many")
	}
	return shapeX, shapeY
}

func (window *Window) resetCurrentConnection() {
	window.clickedConnection = false
	window.currentConnection = nil
}

func (window *Window) connectionExistsOrSelf(id1, id2 int) bool {
	if id1 == id2 {
		return true
	}
	for _, conn := range window.connections {
		if conn.HasIds(id1, id2) {
			return true
		}
	}
	return false
}

func (window *Window) popOutConnection(id int) *Connection {
	for _, conn := range window.connections {
		if conn.outShapeId == id {
			return &conn
		}
	}
	return nil
}
