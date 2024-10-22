package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
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

func (c *Connection) MoveInPos(newX, newY float32) {
	c.inPosX = newX
	c.inPosY = newY
	fmt.Println(c.inPosX, c.inPosY)
}

func (c *Connection) MoveOutPos(newX, newY float32) {
	c.outPosX = newX
	c.outPosY = newY
}

func (c *Connection) Draw() {
	inPos := rl.NewVector2(c.inPosX, c.inPosY)
	outPos := rl.NewVector2(c.outPosX, c.outPosY)
	rl.DrawLineBezier(inPos, outPos, 2, FONT_COLOR)
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
