package window

import rl "github.com/gen2brain/raylib-go/raylib"

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
