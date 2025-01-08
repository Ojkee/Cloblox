package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/shapes"
)

func (window *Window) currentConnectionEvent(mousePos *rl.Vector2) error {
	clickedAnyShape := false
	for _, shape := range window.diagramShapes {
		if !rl.CheckCollisionPointRec(*mousePos, shape.GetRect()) {
			continue
		}
		if shape.GetType() == shapes.STOP && !window.clickedConnection {
			continue
		}
		clickedAnyShape = true
		if !window.clickedConnection {
			shapeX, shapeY, multipleOut, closerToRight := window.getShapeOutPos(
				&shape,
				mousePos,
			)
			window.currentConnection = NewConnection(
				shapeX, shapeY,
				mousePos.X, mousePos.Y,
				-1, shape.GetBlockId(),
				multipleOut, closerToRight,
			)
			window.clickedConnection = true
		} else if !window.connectionExistsOrSelf(shape.GetBlockId(), window.currentConnection.inShapeId) {
			shapePosX, shapePosY := shape.GetInPos()
			window.currentConnection.MoveOutPos(shapePosX, shapePosY)
			window.currentConnection.inShapeId = shape.GetBlockId()
			if err := window.connectBlocksByConnection(window.currentConnection); err != nil {
				return err
			}
			window.connections = append(window.connections, *window.currentConnection)
			window.resetCurrentConnection()
		}
	}
	if !clickedAnyShape {
		window.resetCurrentConnection()
	}
	return nil
}

func (window *Window) updateCurrentConnection(mousePos *rl.Vector2) {
	if window.clickedConnection {
		window.currentConnection.MoveOutPos(mousePos.X, mousePos.Y)
	}
}

func (window *Window) getShapeOutPos(
	shape *shapes.Shape,
	mousePos *rl.Vector2,
) (float32, float32, bool, bool) {
	var shapeX, shapeY float32
	multipleOut := false
	closerToRight := false
	if shapeManyOut, ok := (*shape).(shapes.ShapeManyOut); ok {
		multipleOut = true
		if shapeManyOut.CloserToRight(*mousePos) {
			shapeX, shapeY = shapeManyOut.GetOutPosTrue()
			closerToRight = true
		} else {
			shapeX, shapeY = shapeManyOut.GetOutPosFalse()
		}
	} else if shapeSingleOut, ok := (*shape).(shapes.ShapeSingleOut); ok {
		shapeX, shapeY = shapeSingleOut.GetOutPos()
	} else {
		panic("window/getShapeOutPos fail:\n\tNeither Single nor Many")
	}
	return shapeX, shapeY, multipleOut, closerToRight
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

func (window *Window) connectBlocksByConnection(conn *Connection) error {
	idOut := conn.outShapeId
	idIn := conn.inShapeId
	if conn.multipleOut {
		err := window.diagram.ConnectByIds(idOut, idIn, conn.closerToRight)
		return err
	}
	err := window.diagram.ConnectByIds(idOut, idIn)
	return err
}
