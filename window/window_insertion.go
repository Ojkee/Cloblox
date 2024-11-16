package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/functools"
	"Cloblox/settings"
	"Cloblox/shapes"
)

func (window *Window) insertManager(mousePos *rl.Vector2) []error {
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
				window.diagramShapes[i].SetHighlight(true)
			}
		} else {
			window.diagramShapes[i].SetHighlight(false)
		}
	}
	if !shapeClicked {
		window.flushInsertShape()
	}
}

func (window *Window) keyInsertManager() error {
	if isPasteShortcutClicked() {
		clipboardText := rl.GetClipboardText()
		for _, letter := range clipboardText {
			window.insertChar(string(letter))
		}
		return nil
	}
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
			window.insertCursorY++
			window.insertCursorX = -1
		}
	}
	if window.currentInsertShape != nil {
		window.moveCursorHander()
	}
	return nil
}

func isPasteShortcutClicked() bool {
	if (rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl)) &&
		rl.IsKeyPressed(rl.KeyV) {
		return true
	}
	return false
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

func (window *Window) moveCursorHander() {
	content := (*window.currentInsertShape).GetContent()
	if rl.IsKeyPressed(rl.KeyUp) {
		if window.insertCursorY > 0 {
			window.insertCursorY--
			window.insertCursorX = min(len(content[window.insertCursorY]), window.insertCursorX)
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if window.insertCursorY < len(content)-1 {
			window.insertCursorY++
			window.insertCursorX = min(len(content[window.insertCursorY]), window.insertCursorX)
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		if window.insertCursorX > 0 {
			window.insertCursorX--
		}
	} else if rl.IsKeyPressed(rl.KeyRight) {
		if window.insertCursorX < len(content[window.insertCursorY]) {
			window.insertCursorX++
		}
	}
}

func (window *Window) appendNewLine(lineBeforeIdx int) *[]string {
	content := (*window.currentInsertShape).GetContent()
	if len(content) == 0 {
		return &[]string{""}
	}
	retVal := make([]string, 0)
	retVal = append(retVal, content[:lineBeforeIdx+1]...)
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
		if window.insertCursorX > -1 {
			lhs = currentBuffr[:window.insertCursorX]
			rhs = currentBuffr[window.insertCursorX:]
		}
		newBuffr := lhs + char + rhs
		if len(newBuffr) == 1 {
			window.insertCursorX++
		}
		content[window.insertCursorY] = newBuffr
		window.insertCursorX++
	}
	return &content
}

func (window *Window) removeFromBufferHandler() {
	content := (*window.currentInsertShape).GetContent()
	if len(content[window.insertCursorY]) == 0 && len(content) > 1 {
		newContent := window.removeLine()
		(*window.currentInsertShape).SetContent(newContent)
		if window.insertCursorY > 0 {
			window.insertCursorY--
		}
		window.insertCursorX = len((*newContent)[window.insertCursorY])
	} else if len(content[window.insertCursorY]) > 0 {
		newContent := window.removeChar()
		window.insertCursorX--
		(*window.currentInsertShape).SetContent(newContent)
		if len((*newContent)[window.insertCursorY]) == 0 {
			window.insertCursorX--
		}
	}
}

func (window *Window) removeLine() *[]string {
	content := (*window.currentInsertShape).GetContent()
	retVal := make([]string, 0)
	for i, line := range content {
		if i != window.insertCursorY {
			retVal = append(retVal, line)
		}
	}
	return &retVal
}

func (window *Window) removeChar() *[]string {
	retVal := (*window.currentInsertShape).GetContent()
	currentBuffr := []rune{}
	for i, letter := range retVal[window.insertCursorY] {
		if i != window.insertCursorX-1 {
			currentBuffr = append(currentBuffr, letter)
		}
	}
	retVal[window.insertCursorY] = string(currentBuffr)
	return &retVal
}

func (window *Window) isCursorBlank() bool {
	return window.insertCursorX == -1 && window.insertCursorY == -1
}

func (window *Window) setCursorEnd() {
	if window.currentInsertShape == nil {
		return
	}
	content := (*window.currentInsertShape).GetContent()
	window.insertCursorY = len(content) - 1
	if window.insertCursorY == -1 {
		window.insertCursorX = -1
		return
	}
	window.insertCursorX = len(content[window.insertCursorY]) - 1
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
		letterBuffr := ""
		for _, letter := range buffr {
			letterBuffr += "'" + string(letter) + "'"
		}
		fmt.Printf("[ <%d>  %s ]\n", i, letterBuffr)
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
	if window.insertCursorX == -1 || window.currentInsertShape == nil {
		return
	}
	shape := *window.currentInsertShape
	shapeRect := shape.GetRect()
	content := shape.GetContent()
	currentChar := ""
	if window.insertCursorY != -1 {
		if window.insertCursorX < len(content[window.insertCursorY]) && window.insertCursorX > -1 {
			currentChar = string(content[window.insertCursorY][window.insertCursorX])
		}
	}

	posY := shapeRect.Y + float32(window.insertCursorY)*float32(settings.FONT_SIZE)
	offsetX := shapeRect.X + shapeRect.Width/2
	offsetText := functools.TextWidthEx(content[window.insertCursorY]).X / 2
	textTillCursor := content[window.insertCursorY][:window.insertCursorX]
	offsetCurrentPosText := functools.TextWidthEx(textTillCursor).X
	posX := offsetX - offsetText + offsetCurrentPosText

	rectColor := settings.FONT_COLOR
	posVec := rl.NewVector2(posX, posY)
	sizeVec := rl.NewVector2(float32(settings.FONT_SIZE)/2, float32(settings.FONT_SIZE))
	rl.DrawRectangleV(
		posVec,
		sizeVec,
		rectColor,
	)
	rl.DrawTextEx(
		settings.FONT,
		currentChar,
		posVec,
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		reverseColor(&rectColor),
	)
}

func reverseColor(color *rl.Color) rl.Color {
	return rl.NewColor(255-color.R, 255-color.G, 255-color.B, 255)
}

func (window *Window) flushInsertShape() {
	window.currentInsertShape = nil
	window.flushInsertCursor()
}

func (window *Window) flushInsertCursor() {
	window.insertCursorX = -1
	window.insertCursorY = -1
}
