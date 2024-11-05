package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

func (window *Window) drawMode() {
	textWidth := rl.MeasureText(string(window.currentMode), settings.FONT_SIZE)
	rl.DrawText(
		string(window.currentMode),
		window.width-textWidth-4,
		0,
		settings.FONT_SIZE,
		window.fontColor,
	)
}

func (window *Window) changeModeEvent() {
	switch rl.GetKeyPressed() {
	case 0: // DEFAULT NO PRESS KEY
		break
	case rl.KeyI:
		window.currentMode = INSERTION
		break
	case rl.KeyR:
		window.currentMode = REMOVE
		break
	case rl.KeyEscape, rl.KeyB:
		window.currentMode = BUILDING
		break
	case rl.KeyS:
		window.currentMode = SIMULATION
		break
	default:
		fmt.Print("Mode not implremented")
		break
	}
}
