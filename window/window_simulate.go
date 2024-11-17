package window

import (
	"errors"
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/blocks"
	"Cloblox/functools"
	"Cloblox/settings"
)

type SIMULATE_MODE int

const (
	NOT_SELECTED SIMULATE_MODE = iota
	STEP_BY_STEP
	CONTINUOUSLY
	PAUSE
	FINISHED
)

type VarButton struct {
	name string
	rect rl.Rectangle
}

func NewVarButton(name string, rect rl.Rectangle) *VarButton {
	return &VarButton{
		name: name,
		rect: rect,
	}
}

func (window *Window) simulateManager(mousePos *rl.Vector2) []error {
	if window.em.ContainsStongError() {
		return nil
	}
	window.SelectVarButtonOnClick(mousePos)

	if window.simulationMode == CONTINUOUSLY {
		window.ContinuousSimulation()
	}
	return nil
}

func (window *Window) SelectVarButtonOnClick(mousePos *rl.Vector2) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		for _, button := range window.simulationVarButtons {
			if rl.CheckCollisionPointRec(*mousePos, button.rect) {
				window.simulationVar = button.name
			}
		}
	}
}

func (window *Window) ContinuousSimulation() {
	finished, consoleMessage, err := window.diagram.MakeStep()
	if consoleMessage != "" {
		window.consoleLines = append(
			window.consoleLines,
			*NewConsoleLine(consoleMessage, settings.FONT_COLOR),
		)
	}
	current := window.diagram.GetCurrent()
	currentId := (*current).GetId()
	for i := range window.diagramShapes {
		if window.diagramShapes[i].GetBlockId() == currentId {
			window.diagramShapes[i].SetHighlight(true)
		} else {
			window.diagramShapes[i].SetHighlight(false)
		}
	}
	window.em.AppendNew(err)
	if finished {
		window.simulationMode = FINISHED
	} else {
		time.Sleep(time.Millisecond * settings.SIMULATION_TIME_STEP_MS)
	}
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
		errs = append(errs, err)
	}
	window.simulationVarButtons = window.LoadSliceVarButtons()
	if len(window.simulationVarButtons) == 0 {
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

func (window *Window) LoadSliceVarButtons() []VarButton {
	retVal := make([]VarButton, 0)
	for _, block := range window.diagram.GetAllBlocks() {
		switch variableBlock := block.(type) {
		case *blocks.VariablesBlock:
			allVars := variableBlock.GetVars()
			for key, value := range allVars {
				switch value.(type) {
				case []float64:
					retVal = append(retVal, *NewVarButton(key, rl.NewRectangle(
						settings.BUTTON_X,
						settings.BUTTON_Y+float32(len(retVal))*(settings.BUTTON_HEIGHT+settings.BUTTON_GAP),
						settings.BUTTON_WIDTH,
						settings.BUTTON_HEIGHT,
					)))
					break
				}
			}
			break
		}
	}
	return retVal
}

func (window *Window) drawAllSlicesButtons() {
	for _, val := range window.simulationVarButtons {
		window.drawSimulateButton(&val)
	}
}

func (window *Window) drawSimulateButton(vb *VarButton) {
	color := settings.BUTTON_COLOR
	if vb.name == window.simulationVar {
		color = settings.BUTTON_COLOR_SELECTED
	}
	rl.DrawRectangleRounded(vb.rect, 0.9, 10, color)
	textSize := rl.MeasureTextEx(
		settings.FONT,
		vb.name,
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
	)
	textOffsetX := vb.rect.X + vb.rect.Width/2 - textSize.X/2
	textOffsetY := vb.rect.Y + vb.rect.Height/2 - textSize.Y/2
	rl.DrawTextEx(
		settings.FONT,
		vb.name,
		rl.NewVector2(textOffsetX, textOffsetY),
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
	if !window.diagram.ContainsVar(window.simulationVar) {
		return nil
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
	window.simulationVarButtons = make([]VarButton, 0)
	window.simulationVar = ""
	window.clearConsole()
	window.em.Clear()
}

func (window *Window) clearConsole() {
	window.consoleLines = make([]ConsoleLine, 0)
}
