package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
	"Cloblox/shapes"
)

func (window *Window) insertManager(mousePos *rl.Vector2) error {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		window.selectInsertShape(mousePos)
	} else if !rl.IsKeyPressed(rl.KeyZero) && window.currentInsertShape != nil {
		window.keyInsertManager()
	}
	return nil
}

func (window *Window) selectInsertShape(mousePos *rl.Vector2) {
	shapeClicked := false
	for i := range window.diagramShapes {
		rect := window.diagramShapes[i].GetRect()
		if rl.CheckCollisionPointRec(*mousePos, rect) {
			window.setCurrentInsertShape(&window.diagramShapes[i])
			if window.currentInsertShape != nil {
				window.setCursorEnd()
				shapeClicked = true
			}
		}
	}
	if !shapeClicked {
		window.flushInsertShape()
	}
}

func (window *Window) keyInsertManager() error {
	charPressed := rl.GetCharPressed()
	if charPressed >= 32 && charPressed <= 126 {
		window.insertNewBufferHandler(charPressed)
	} else if rl.IsKeyPressed(rl.KeyBackspace) {
		window.removeFromBufferHandler()
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		content := (*window.currentInsertShape).GetContent()
		if content[window.insertCursorY] != "" && len(content) < settings.MAX_CONTENT_LINES {
			newContent := window.appendNewLine(window.insertCursorY)
			(*window.currentInsertShape).SetContent(newContent)
			window.setCursorEnd()
		}
	}
	if window.currentInsertShape != nil {
		window.moveCursorHander()
	}
	return nil
}

func (window *Window) insertNewBufferHandler(keyPressed int32) {
	if !(*window.currentInsertShape).IsContentEmpty() {
		newContent := window.insertChar(string(keyPressed))
		(*window.currentInsertShape).SetContent(newContent)
	} else {
		newContent := window.appendNewLine(0)
		(*window.currentInsertShape).SetContent(newContent)
		window.setCursorEnd()
	}
}

func (window *Window) removeFromBufferHandler() {
	content := window.removeChar()
	(*window.currentInsertShape).SetContent(content)
}

func (window *Window) moveCursorHander() {
	content := (*window.currentInsertShape).GetContent()
	if rl.IsKeyPressed(rl.KeyUp) {
		if window.insertCursorY > 0 {
			window.insertCursorY--
			window.insertCursorX = min(len(content[window.insertCursorY])-1, window.insertCursorX)
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if window.insertCursorY < len(content)-1 {
			window.insertCursorY++
			window.insertCursorX = min(len(content[window.insertCursorY])-1, window.insertCursorX)
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		if window.insertCursorX > 0 {
			window.insertCursorX--
		}
	} else if rl.IsKeyPressed(rl.KeyRight) {
		if window.insertCursorX < len(content[window.insertCursorY])-1 {
			window.insertCursorX++
		}
	}
}

func (window *Window) appendNewLine(lineBeforeIdx int) *[]string {
	content := (*window.currentInsertShape).GetContent()
	if len(content) == 0 {
		return &[]string{""}
	}
	retVal := content[:lineBeforeIdx+1]
	retVal = append(retVal, "")
	retVal = append(retVal, content[lineBeforeIdx+1:]...)
	return &retVal
}

func (window *Window) insertChar(char string) *[]string {
	content := (*window.currentInsertShape).GetContent()
	currentBuffr := content[window.insertCursorY]
	if len(currentBuffr) < settings.MAX_CONTENT_CHARS {
		lhs := ""
		rhs := currentBuffr
		if window.insertCursorX > 0 {
			lhs = currentBuffr[:window.insertCursorX]
			rhs = currentBuffr[window.insertCursorX:]
		}
		newBuffr := lhs + char + rhs
		content[window.insertCursorY] = newBuffr
		window.insertCursorX++
	}
	return &content
}

func (window *Window) removeLine(content *[]string) *[]string {
	retVal := (*content)[:window.insertCursorY]
	retVal = append(retVal, (*content)[window.insertCursorY+1:]...)
	return &retVal
}

func (window *Window) removeChar() *[]string {
	content := (*window.currentInsertShape).GetContent()
	retVal := (content)[:window.insertCursorY]
	currentBuffr := []rune{}
	for i, letter := range (content)[window.insertCursorY] {
		if i != window.insertCursorX {
			currentBuffr = append(currentBuffr, letter)
		}
	}
	retVal[window.insertCursorY] = string(currentBuffr)
	retVal = append(retVal, content[window.insertCursorY+1:]...)
	window.insertCursorX = min(window.insertCursorX, len(retVal[window.insertCursorY])-1)
	return &retVal
}

func (window *Window) flushInsertShape() {
	window.currentInsertShape = nil
	window.insertCursorX = -1
	window.insertCursorY = -1
}

func (window *Window) isCursorBlank() bool {
	return window.insertCursorX == -1 && window.insertCursorY == -1
}

func (window *Window) setCursorEnd() {
	content := (*window.currentInsertShape).GetContent()
	window.insertCursorY = len(content) - 1
	window.insertCursorX = len((content)[window.insertCursorY]) - 1
}

func (window *Window) setCursorPos(y, x int) {
	window.insertCursorY = y
	window.insertCursorX = x
}

func (window *Window) debugContent() {
	fmt.Println("*************************CONTENT**************************")
	if window.currentInsertShape == nil {
		fmt.Println("No shape selected")
		fmt.Println()
		fmt.Println()
		return
	}
	fmt.Printf("Y: %d\t\tX: %d\n\n", window.insertCursorY, window.insertCursorX)
	content := (*window.currentInsertShape).GetContent()
	for i, buffr := range content {
		fmt.Printf("[ <%d>  %s ]\n", i, buffr)
	}
	fmt.Println()
	fmt.Println()
}

func (window *Window) setCurrentInsertShape(shape *shapes.Shape) {
	if shape == nil {
		window.currentInsertShape = nil
		return
	}
	if (*shape).GetType() == shapes.START || (*shape).GetType() == shapes.STOP {
		window.currentInsertShape = nil
		return
	}
	window.currentInsertShape = shape
}

func (window *Window) drawCursor() {
	if window.insertCursorX == -1 || window.insertCursorY == -1 ||
		window.currentInsertShape == nil {
		return
	}
	shape := *window.currentInsertShape
	shapeRect := shape.GetRect()
	content := shape.GetContent()
	currentChar := ""
	if window.insertCursorY < len(content[window.insertCursorY]) {
		currentChar = string(content[window.insertCursorY][window.insertCursorX])
	}

	posY := int32(shapeRect.Y) + int32(window.insertCursorY)*int32(settings.FONT_SIZE)
	offsetX := shapeRect.X + shapeRect.Width/2
	offsetText := float32(shape.GetContentSize(window.insertCursorY)) / 2
	textTillCursor := content[window.insertCursorY][:window.insertCursorX]
	offsetCurrentPosText := float32(rl.MeasureText(textTillCursor, settings.FONT_SIZE))
	posX := int32(offsetX - offsetText + offsetCurrentPosText)

	rectColor := settings.FONT_COLOR
	rl.DrawRectangle(
		int32(posX)-1,
		posY,
		settings.FONT_SIZE/2+2,
		settings.FONT_SIZE,
		rectColor,
	)
	rl.DrawText(
		currentChar,
		posX,
		posY,
		settings.FONT_SIZE,
		reverseColor(&rectColor),
	)
}

func reverseColor(color *rl.Color) rl.Color {
	return rl.NewColor(255-color.R, 255-color.G, 255-color.B, 255)
}
