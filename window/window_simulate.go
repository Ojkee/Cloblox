package window

import (
	"errors"
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

type VarButton struct {
	name string
}

func NewVarButton(name string) *VarButton {
	return &VarButton{
		name: name,
	}
}

func (vButton *VarButton) GetName() string {
	return vButton.name
}

func (window *Window) simulateManager(mousePos *rl.Vector2) []error {
	if !window.simulationPrecompiled {
		errs := window.preSimulationCompile()
		window.simulationPrecompiled = true
		return errs
	}
	// if rl.CheckCollisionPointRec(*mousePos , window.consoleInnerRect) {
	//   if wheelMove := rl.GetMouseWheelMove(); wheelMove != 0 { // CONSOLE SCROLL
	//
	//   }
	// }
	return nil
}

func (window *Window) preSimulationCompile() []error {
	var errs []error
	var err error
	err = window.shipContentToBlocks()
	if err != nil {
		errs = append(errs, err)
	}
	err = window.diagram.InitIfValid()
	if err != nil {
		return append(errs, err)
	}
	window.simulationSlicesVars = window.ReadAllSliceVars()
	if len(window.simulationSlicesVars) == 0 {
		errs = append(errs, errors.New("No array to visualize"))
	}
	return errs
}

func (window *Window) shipContentToBlocks() error {
	for i := range window.diagramShapes {
		id := window.diagramShapes[i].GetBlockId()
		content := window.diagramShapes[i].GetContent()
		err := window.diagram.SetBlockContentById(&content, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (window *Window) ReadAllSliceVars() []string {
	retVal := make([]string, 0)

	return retVal
}

func (window *Window) drawAllSlicesButtons() {
	var margin float32 = 10
	var buttonWidth float32 = 40
	var buttonHeight float32 = 20
	var buttonX float32 = settings.WINDOW_WIDTH/2 - buttonWidth - margin
	var buttonY float32 = margin
	var gap float32 = 10
	for i, varName := range window.simulationSlicesVars {
		x := buttonX + float32(i)*(buttonHeight+gap)
		rect := rl.NewRectangle(x, buttonY, buttonWidth, buttonHeight)
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
	if window.simulationVar == "" {
		return errors.New("No array selected")
	}
	slices, err := window.diagram.GetAllSlicesKVP()
	if err != nil {
		return err
	}
	var highPos float64 = settings.WINDOW_HEIGHT - 10 - settings.CONSOLE_HEIGHT
	var lowPos float64 = 10

	temp := make([]float64, len(slices[window.simulationVar]))
	copy(temp, slices[window.simulationVar])
	scaledSlice, pos, neg := functools.GetScaledSlice(temp, lowPos, highPos)
	rWidth := int32(settings.WINDOW_WIDTH / 2 / len(scaledSlice))
	var basePos float64
	var textPosY float64
	if pos && neg {
		basePos = (highPos + lowPos) / 2
		textPosY = basePos
	} else if pos {
		basePos = highPos
		textPosY = basePos - float64(settings.FONT_SIZE)
	} else {
		basePos = lowPos
		textPosY = basePos
	}
	for i, val := range scaledSlice {
		posX := int32(i) * rWidth
		window.drawValueFromSlice(
			posX,
			int32(basePos),
			rWidth,
			int32(val),
			slices[window.simulationVar][i],
			textPosY,
		)
	}
	return nil
}

func (window *Window) drawValueFromSlice(
	posX, basePos, rWidth, valScaled int32,
	val, textPosY float64,
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
	textWidth := rl.MeasureTextEx(
		settings.FONT,
		fmt.Sprintf("%.2f ", val),
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
	).X
	textPosX := float32(posX+rWidth/2) - textWidth/2
	rl.DrawTextEx(
		settings.FONT,
		fmt.Sprintf("%.2f ", val),
		rl.NewVector2(textPosX, float32(textPosY)),
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		settings.FONT_COLOR,
	)
}

func (window *Window) flushSimulate() {
	window.simulationStarted = false
	window.simulationMode = NOT_SELECTED
	window.simulationSlicesVars = make([]string, 0)
	window.simulationVar = ""
	window.simulationPrecompiled = false
	window.clearConsole()
	window.em.Clear()
}

func (window *Window) clearConsole() {
	window.consoleLines = make([]ConsoleLine, 0)
}
