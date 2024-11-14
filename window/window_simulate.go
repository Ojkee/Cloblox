package window

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/functools"
	"Cloblox/settings"
)

type SIMULATE_MODE int

const (
	NOT_SELECTED SIMULATE_MODE = iota
	STEP_BY_STEP
	CONTINUOUSLY
)

func (window *Window) simulateManager(mousePos *rl.Vector2) {
}

func (window *Window) drawAllSlicesButtons() {
	var margin float32 = 10
	var buttonWidth float32 = 40
	var buttonHeight float32 = 20
	var buttonX float32 = settings.WINDOW_WIDTH/2 - buttonWidth - margin
	var buttonY float32 = margin
	var gap float32 = 10
	allVars := window.diagram.GetAllSliceVars()
	for i, varName := range allVars {
		x := int(buttonX) + i*(int(buttonHeight)+int(gap))
		rect := rl.NewRectangle(float32(x), buttonY, buttonWidth, buttonHeight)
		drawSimulateButton(&rect, &varName)
	}
}

func drawSimulateButton(rect *rl.Rectangle, text *string) {
	rl.DrawRectangleRounded(*rect, 0.05, 10, settings.HIGHLIGHT_COLOR)
	rl.DrawTextEx(
		settings.FONT,
		*text,
		rl.NewVector2(rect.X, rect.Y),
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		settings.FONT_COLOR,
	)
}

func (window *Window) drawCurrentSlice() error {
	slices, err := window.diagram.GetAllSlicesKVP()
	if err != nil {
		return err
	}
	var highPos float64 = settings.WINDOW_HEIGHT - 10
	var lowPos float64 = 10

	temp := make([]float64, len(slices[window.simulationVar]))
	copy(temp, slices[window.simulationVar])
	scaledSlice, pos, neg := functools.GetScaledSlice(temp, lowPos, highPos)
	rWidth := int32(settings.WINDOW_WIDTH / 2 / len(scaledSlice))
	var basePos float64
	if pos && neg {
		basePos = (highPos + lowPos) / 2
	} else if pos {
		basePos = highPos
	} else {
		basePos = lowPos
	}
	for i, val := range scaledSlice {
		posX := int32(i) * rWidth
		window.drawValueFromSlice(
			posX,
			int32(basePos),
			rWidth,
			int32(val),
			slices[window.simulationVar][i],
		)
	}
	return nil
}

func (window *Window) drawValueFromSlice(
	posX, basePos, rWidth, valScaled int32,
	val float64,
) {
	color := settings.POSITIVE_VAL_COLOR
	valScaledAbs := int32(math.Abs(float64(valScaled)))
	if valScaled < 0 {
		color = settings.NEGATIVE_VAL_COLOR
		rl.DrawRectangle(
			posX,
			basePos,
			rWidth,
			valScaledAbs,
			color,
		)
	} else {
		rl.DrawRectangle(posX, int32(basePos)-int32(valScaled), rWidth, int32(valScaled), color)
	}
	rl.DrawTextEx(
		settings.FONT,
		fmt.Sprintf("%.2f ", val),
		rl.NewVector2(float32(posX), float32(basePos)),
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		settings.FONT_COLOR,
	)
}
