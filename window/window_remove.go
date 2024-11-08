package window

import rl "github.com/gen2brain/raylib-go/raylib"

func (window *Window) removeManager(mousePos *rl.Vector2) error {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) ||
		rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		if removeId := window.getShapeIdMousePos(mousePos); removeId != -1 {
			window.removeShapeAndConnectionsById(removeId)
		}
	}
	return nil
}

func (window *Window) removeShapeAndConnectionsById(id int) {
	shapeI := 0
	for _, shape := range window.diagramShapes {
		if shape.GetBlockId() != id {
			window.diagramShapes[shapeI] = shape
			shapeI++
		}
	}
	connI := 0
	window.diagramShapes = window.diagramShapes[:shapeI]
	for _, conn := range window.connections {
		if !conn.HasId(id) {
			window.connections[connI] = conn
			connI++
		}
	}
	window.connections = window.connections[:connI]
	window.diagram.RemoveBlockById(id)
}

func (window *Window) getShapeIdMousePos(mousePos *rl.Vector2) int {
	for _, shape := range window.diagramShapes {
		if rl.CheckCollisionPointRec(*mousePos, shape.GetRect()) {
			return shape.GetBlockId()
		}
	}
	return -1
}
