package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

type ConsoleLine struct {
	line  string
	color rl.Color
}

func NewConsoleLine(line string, color rl.Color) *ConsoleLine {
	return &ConsoleLine{
		line:  line,
		color: color,
	}
}

func (window *Window) drawConsole() {
	var margin float32 = 10
	var roundFactor float32 = 0.18
	offsetY := float32(settings.WINDOW_HEIGHT - settings.CONSOLE_HEIGHT)
	rl.DrawRectangleRounded(
		rl.NewRectangle(
			margin,
			offsetY,
			settings.WINDOW_WIDTH/2-2*margin,
			settings.CONSOLE_HEIGHT-margin,
		),
		roundFactor,
		16,
		settings.HELP_OUTER_BORDER_COLOR,
	)
	rl.DrawRectangleRounded(
		window.consoleInnerRect,
		roundFactor,
		8,
		settings.HELP_INNER_BORDER_COLOR,
	)
	startIdx := max(0, len(window.consoleLines)-settings.CONSOLE_MAX_LINES)
	endCon := func(i int) bool {
		return i < startIdx+settings.CONSOLE_MAX_LINES && i < len(window.consoleLines)
	}
	textOffset := offsetY + float32(settings.FONT_SIZE)
	for i := startIdx; endCon(i); i++ {
		rl.DrawTextEx(
			settings.FONT,
			fmt.Sprintf("%s %s", settings.CONSOLE_PREFIX, window.consoleLines[i].line),
			rl.NewVector2(20, textOffset+float32((i-startIdx)*int(settings.FONT_SIZE))),
			float32(settings.FONT_SIZE),
			settings.FONT_SPACING,
			window.consoleLines[i].color,
		)
	}
}
