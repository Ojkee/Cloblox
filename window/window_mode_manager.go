package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type MODE string

const (
	BUILD    MODE = "build"
	INSERT        = "insert"
	REMOVE        = "remove"
	SIMULATE      = "simulate"
	DEBUG         = "debug"
)

func (window *Window) drawCurrentMode() {
	modeStr := string(window.currentMode)
	textWidth := rl.MeasureTextEx(
		settings.FONT,
		modeStr,
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
	).X
	rl.DrawTextEx(
		settings.FONT,
		modeStr,
		rl.NewVector2(float32(window.width)-textWidth-4, 0),
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		window.fontColor,
	)
}

func (window *Window) changeModeEvent() {
	keyPressed := rl.GetKeyPressed()
	if window.currentMode == INSERT {
		if keyPressed == rl.KeyEscape {
			window.currentMode = BUILD
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
		window.currentMode = BUILD
		break
	case rl.KeyI, rl.KeyTwo:
		window.currentMode = INSERT
		break
	case rl.KeyR, rl.KeyThree:
		window.currentMode = REMOVE
		break
	case rl.KeyS, rl.KeyFour:
		window.currentMode = SIMULATE
		break
	default:
		// fmt.Print("Mode not implemented")
		break
	}
}
