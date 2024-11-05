package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/blocks"
	"Cloblox/settings"
	"Cloblox/shapes"
)

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
	if mousePos.X < settings.WINDOW_WIDTH/2 {
		window.resetClickedShape()
	}
	if window.shapeClicked && !clickedNewShape {
		window.placeCurrentShape(mousePos.X, mousePos.Y)
		window.resetClickedShape()
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
		window.resetClickedShape()
		panic("window.go/makeCurrentClicked fail:\n\tNot implemented shape type")
	}
	window.diagram.AppendBlock(cBlock)
	cShape.TranslateCenter()
	cShape.SetBlockId(cBlock.GetId())
	window.diagramShapes = append(window.diagramShapes, cShape)
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
		window.resetClickedShape()
		panic("window.go/makeCurrentClicked fail:\n\tNot implemented shape type")
	}
}

func (window *Window) updateCurrentShape(mousePos *rl.Vector2) {
	window.currentShape.MoveTo(
		mousePos.X-settings.SHAPE_WIDTH/2,
		mousePos.Y-settings.SHAPE_HEIGHT/2,
	)
}

func (window *Window) resetClickedShape() {
	window.shapeClicked = false
	window.currentShape = nil
	window.currentShapeType = shapes.NONE
}
