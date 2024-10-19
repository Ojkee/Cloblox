package window

import rl "github.com/gen2brain/raylib-go/raylib"

type Connection struct {
	inPos         rl.Vector2
	outPos        rl.Vector2
	inShapeId     int
	outShapeId    int
	closerToRight bool
}

func NewConnection(
	inPosX, inPosY,
	outPosX, outPosY float32,
	inShapeId, outShapeId int,
) *Connection {
	return &Connection{
		inPos:         rl.NewVector2(inPosX, inPosY),
		outPos:        rl.NewVector2(outPosX, outPosY),
		inShapeId:     inShapeId,
		outShapeId:    outShapeId,
		closerToRight: false,
	}
}

func (c *Connection) MoveOutPos(newX, newY float32) {
	c.outPos.X = newX
	c.outPos.Y = newY
}

func (c *Connection) Draw() {
	rl.DrawLineBezier(c.inPos, c.outPos, 2, FONT_COLOR)
}

func (c *Connection) HasIds(id1, id2 int) bool {
	return (c.inShapeId == id1 && c.outShapeId == id2) ||
		(c.inShapeId == id2 && c.outShapeId == id1)
}

func (c *Connection) SetCloserToRight(isCloser bool) {
	c.closerToRight = isCloser
}

func (c *Connection) IsCloserToRigth() bool {
	return c.closerToRight
}
