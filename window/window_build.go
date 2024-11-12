package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/blocks"
	"Cloblox/settings"
	"Cloblox/shapes"
)

func (window *Window) buildManager(mousePos *rl.Vector2) error {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) { // New Shape
		window.buildNewShapeEvent(mousePos)
	} else if rl.IsMouseButtonPressed(rl.MouseButtonRight) { // Connect
		err := window.currentConnectionEvent(mousePos)
		if err != nil {
			return err
		}
	}
	return nil
}

func initBuildingShapes(width, height int32) []shapes.Shape {
	offsetX := float32(width/2.0 + 10)
	gap := float32(settings.SHAPE_MIN_HEIGHT + 16)
	offsetY := float32(height)/2.0 - gap*2.5
	return []shapes.Shape{
		shapes.NewStartShape(offsetX, offsetY),
		shapes.NewVariableShape(offsetX, offsetY+gap),
		shapes.NewIfShape(offsetX, offsetY+2*gap),
		shapes.NewActionShape(offsetX, offsetY+3*gap),
		shapes.NewStopShape(offsetX, offsetY+4*gap),
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
