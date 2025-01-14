package window

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/functools"
	"Cloblox/settings"
)

type ConsoleLine struct {
	isStrongError bool
	line          string
	color         rl.Color
}

func NewConsoleLine(line string, color rl.Color) *ConsoleLine {
	return &ConsoleLine{
		isStrongError: false,
		line:          line,
		color:         color,
	}
}

func (window *Window) appendErrorsToConsole(errs []error) {
	if errs != nil {
		for _, err := range errs {
			newLines := make([]ConsoleLine, 0)
			color := settings.FONT_ERROR_COLOR
			if window.errorManager.IsStrong(err) {
				color = settings.FONT_ERROR_STRONG_COLOR
				for _, line := range functools.SplitLine(err.Error(), settings.CONSOLE_MAX_LINE_WIDTH) {
					cl := *NewConsoleLine(line, color)
					cl.isStrongError = true
					newLines = append(newLines, cl)
				}
			} else {
				for _, line := range functools.SplitLine(err.Error(), settings.CONSOLE_MAX_LINE_WIDTH) {
					newLines = append(newLines, *NewConsoleLine(line, color))
				}
			}
			window.consoleLines = append(
				window.consoleLines,
				newLines...)
		}
	}
}

func (window *Window) appendTextToConsole(text string) {
	window.consoleLines = append(window.consoleLines, *NewConsoleLine(text, settings.FONT_COLOR))
}

func (window *Window) drawConsole() {
	var margin float32 = 10
	var roundFactor float32 = 0.18
	offsetY := float32(settings.WINDOW_HEIGHT - settings.CONSOLE_HEIGHT)
	rl.DrawRectangleRounded(
		rl.NewRectangle(
			margin,
			offsetY,
			settings.CONSOLE_WIDTH-2*margin,
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
		prefix := settings.CONSOLE_PREFIX
		if window.consoleLines[i].isStrongError {
			prefix = settings.CONSOLE_PREFIX_STRONG_ERROR
		}
		rl.DrawTextEx(
			settings.FONT,
			fmt.Sprintf("%s %s", prefix, window.consoleLines[i].line),
			rl.NewVector2(20, textOffset+float32((i-startIdx)*int(settings.FONT_SIZE))),
			float32(settings.FONT_SIZE),
			settings.FONT_SPACING,
			window.consoleLines[i].color,
		)
	}
}
