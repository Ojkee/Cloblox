package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type MODE string

const (
	BUILDING   MODE = "building"
	INSERTION       = "insertion"
	REMOVE          = "remove"
	SIMULATION      = "simulation"
	DEBUG           = "debug"
)

func (window *Window) drawCurrentMode() {
	modeStr := string(window.currentMode)
	textWidth := rl.MeasureText(modeStr, settings.FONT_SIZE)
	rl.DrawText(
		modeStr,
		window.width-textWidth-4,
		0,
		settings.FONT_SIZE,
		window.fontColor,
	)
}

func (window *Window) changeModeEvent() {
	keyPressed := rl.GetKeyPressed()
	if window.currentMode == INSERTION {
		if keyPressed == rl.KeyEscape {
			window.currentMode = BUILDING
		} else {
			return
		}
	}
	switch keyPressed {
	case
		0,       // DEFAULT NO PRESS KEY
		rl.KeyD: // RESERVED FOR DEBUG
		break
	case rl.KeyEscape, rl.KeyB, rl.KeyOne:
		window.currentMode = BUILDING
		break
	case rl.KeyI, rl.KeyTwo:
		window.currentMode = INSERTION
		break
	case rl.KeyR, rl.KeyThree:
		window.currentMode = REMOVE
		break
	case rl.KeyS, rl.KeyFour:
		window.currentMode = SIMULATION
		break
	default:
		fmt.Print("Mode not implremented")
		break
	}
}
