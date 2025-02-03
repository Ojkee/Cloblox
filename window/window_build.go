package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/blocks"
	"Cloblox/graph"
	"Cloblox/iostate"
	"Cloblox/settings"
	"Cloblox/shapes"
)

func (window *Window) buildManager(mousePos *rl.Vector2) []error {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) { // New Shape
		window.buildNewShapeEvent(mousePos)
		if window.cleanButton.InRect(*mousePos) {
			window.flushAll()
		} else if window.loadButton.InRect(*mousePos) {
			errLoad := window.loadSavedState()
			if errLoad != nil {
				return []error{errLoad}
			}
		}
	} else if rl.IsMouseButtonPressed(rl.MouseButtonRight) { // Connect
		err := window.currentConnectionEvent(mousePos)
		if err != nil {
			return []error{err}
		}
	}
	return nil
}

func initBuildingShapes() []shapes.Shape {
	gap := float32(settings.SHAPE_MIN_HEIGHT + settings.SHAPE_BUILD_GAP_Y)
	offsetY := settings.SHAPE_BUILD_Y - gap*2.5
	return []shapes.Shape{
		shapes.NewStartShape(settings.SHAPE_BUILD_X, offsetY),
		shapes.NewVariableShape(settings.SHAPE_BUILD_X, offsetY+gap),
		shapes.NewIfShape(settings.SHAPE_BUILD_X, offsetY+2*gap),
		shapes.NewActionShape(settings.SHAPE_BUILD_X, offsetY+3*gap),
		shapes.NewStopShape(settings.SHAPE_BUILD_X, offsetY+4*gap),
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
	if mousePos.X < settings.WINDOW_WIDTH/2+settings.SHAPE_MIN_WIDTH+10 && !clickedNewShape {
		window.flushBuildShape()
		window.flushInsertShape()
	}
	if window.shapeClicked && !clickedNewShape {
		window.placeCurrentShape(mousePos.X, mousePos.Y)
		window.flushBuildShape()
	}
}

func (window *Window) placeCurrentShape(mx, my float32) {
	var cShape shapes.Shape
	var cBlock blocks.Block
	switch window.currentShapeType {
	case shapes.START:
		cBlock = blocks.NewStartBlock()
		cShape = shapes.NewStartShape(mx, my)
		break
	case shapes.VARIABLE:
		cBlock = blocks.NewVariableBlock()
		cShape = shapes.NewVariableShape(mx, my)
		break
	case shapes.IF:
		cBlock = blocks.NewIfBlock()
		cShape = shapes.NewIfShape(mx, my)
		break
	case shapes.ACTION:
		cBlock = blocks.NewActionBlock()
		cShape = shapes.NewActionShape(mx, my)
		break
	case shapes.STOP:
		cBlock = blocks.NewStopBlock()
		cShape = shapes.NewStopShape(mx, my)
		break
	default:
		window.flushBuildShape()
		panic("window.go/makeCurrentClicked fail:\n\tNot implemented shape type")
	}
	window.diagram.AppendBlock(cBlock)
	cShape.TranslateCenter()
	cShape.SetBlockId(cBlock.GetId())
	cShape.SetHighlight(true)
	window.diagramShapes = append(window.diagramShapes, cShape)
	for _, conn := range window.connections {
		window.connectBlocksByConnection(&conn)
	}
	window.setCurrentInsertShape(&window.diagramShapes[len(window.diagramShapes)-1])
	window.setCursorEnd()
}

func (window *Window) makeCurrentClicked(shapeType shapes.SHAPE_TYPE) {
	window.currentShapeType = shapeType
	switch shapeType {
	case shapes.START:
		window.currentShape = shapes.NewStartShape(0, 0)
		break
	case shapes.VARIABLE:
		window.currentShape = shapes.NewVariableShape(0, 0)
		break
	case shapes.IF:
		window.currentShape = shapes.NewIfShape(0, 0)
		break
	case shapes.ACTION:
		window.currentShape = shapes.NewActionShape(0, 0)
		break
	case shapes.STOP:
		window.currentShape = shapes.NewStopShape(0, 0)
		break
	default:
		window.flushBuildShape()
		panic("window.go/makeCurrentClicked fail:\n\tNot implemented shape type")
	}
}

func (window *Window) updateCurrentShape(mousePos *rl.Vector2) {
	window.currentShape.MoveTo(
		mousePos.X-settings.SHAPE_MIN_WIDTH/2,
		mousePos.Y-settings.SHAPE_MIN_HEIGHT/2,
	)
}

func (window *Window) flushBuildShape() {
	window.shapeClicked = false
	window.currentShape = nil
	window.currentShapeType = shapes.NONE
}

func (window *Window) flushDiagramShapes() {
	window.diagramShapes = make([]shapes.Shape, 0)
}

func (window *Window) loadSavedState() error {
	_blocks, _connections, err := iostate.ReadFromTxt(settings.PATH_TXT)
	// fmt.Printf("%v\n", _blocks)
	// fmt.Printf("%v\n", _connections)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	window.flushAll()
	window.diagram = *graph.NewGraph(nil)
	window.initNewGraph(_blocks, _connections)
	return nil
}

func (window *Window) initNewGraph(_blocks []blocks.Block, _connections []shapes.Connection) {
	window.diagram.SetAllBlocks(_blocks)
	for _, conn := range _connections {
		if conn.IsMultipleOut() {
			window.diagram.ConnectByIds(
				conn.GetInShapeId(),
				conn.GetOutShapeId(),
				conn.IsCloserToRigth(),
			)
		} else {
			window.diagram.ConnectByIds(
				conn.GetInShapeId(),
				conn.GetOutShapeId(),
			)
		}
	}
}
