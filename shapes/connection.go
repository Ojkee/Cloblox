package shapes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type Connection struct {
	inPosX        float32
	inPosY        float32
	outPosX       float32
	outPosY       float32
	inShapeId     int
	outShapeId    int
	multipleOut   bool
	closerToRight bool
}

func NewConnection(
	inPosX, inPosY,
	outPosX, outPosY float32,
	inShapeId, outShapeId int,
	multipleOut, closerToRight bool,
) *Connection {
	return &Connection{
		inPosX:        inPosX,
		inPosY:        inPosY,
		outPosX:       outPosX,
		outPosY:       outPosY,
		inShapeId:     inShapeId,
		outShapeId:    outShapeId,
		multipleOut:   multipleOut,
		closerToRight: closerToRight,
	}
}

func (conn *Connection) MoveInPos(newX, newY float32) {
	conn.inPosX = newX
	conn.inPosY = newY
}

func (conn *Connection) MoveOutPos(newX, newY float32) {
	conn.outPosX = newX
	conn.outPosY = newY
}

func (conn *Connection) Draw() {
	inPos := rl.NewVector2(conn.inPosX, conn.inPosY)
	outPos := rl.NewVector2(conn.outPosX, conn.outPosY)
	rl.DrawLineBezier(inPos, outPos, 2, settings.FONT_COLOR)
}

func (conn *Connection) IsInId(id int) bool {
	return conn.inShapeId == id
}

func (conn *Connection) HasId(id int) bool {
	return conn.inShapeId == id || conn.outShapeId == id
}

func (conn *Connection) HasIds(id1, id2 int) bool {
	return (conn.inShapeId == id1 && conn.outShapeId == id2) ||
		(conn.inShapeId == id2 && conn.outShapeId == id1)
}

func (conn *Connection) SetCloserToRight(isCloser bool) {
	conn.closerToRight = isCloser
}

func (conn *Connection) IsCloserToRigth() bool {
	return conn.closerToRight
}

func (conn *Connection) GetInShapeId() int {
	return conn.inShapeId
}

func (conn *Connection) SetInShapeId(id int) {
	conn.inShapeId = id
}

func (conn *Connection) GetOutShapeId() int {
	return conn.outShapeId
}

func (conn *Connection) SetOutShapeId(id int) {
	conn.inShapeId = id
}

func (conn *Connection) IsMultipleOut() bool {
	return conn.multipleOut
}
