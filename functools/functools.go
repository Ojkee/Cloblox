package functools

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

func TextWidthEx(text string) rl.Vector2 {
	retVal := rl.MeasureTextEx(
		settings.FONT,
		text,
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
	)
	return retVal
}
