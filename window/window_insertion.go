package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (window *Window) insertManager(mousePos *rl.Vector2) error {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		window.selectInsertShape(mousePos)
	}

	if !rl.IsKeyPressed(rl.KeyZero) && window.currentInsertShape != nil {
		window.keyInsertManager()
	}

	return nil
}

func (window *Window) selectInsertShape(mousePos *rl.Vector2) {
	shapeClicked := false
	for i := range window.diagramShapes {
		rect := window.diagramShapes[i].GetRect()
		if rl.CheckCollisionPointRec(*mousePos, rect) {
			window.currentInsertShape = &window.diagramShapes[i]
			shapeClicked = true
		}
	}
	if !shapeClicked {
		window.currentInsertShape = nil
	}
}

func (window *Window) keyInsertManager() error {
	keyPressed := rl.GetCharPressed()
	if keyPressed >= 32 && keyPressed <= 126 {
		buffr := (*window.currentInsertShape).GetContent()
		cursorRow := len(buffr) - 1
		cursorCol := len(buffr[cursorRow]) - 1
		fmt.Println(cursorRow, cursorCol)
		// char := string(keyPressed)
		// fmt.Println(char)
	}
	if rl.IsKeyPressed(rl.KeyBackspace) {
		fmt.Println("BACK")
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		fmt.Println("Enter")
	}
	return nil
}

func (window *Window) flushInsertShape() {
	window.currentInsertShape = nil
	window.insertBuffer = make([]string, 0)
}
