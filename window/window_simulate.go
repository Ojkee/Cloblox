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
	"Cloblox/shapes"
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

type SimulateButton struct {
	rect         rl.Rectangle
	simulateType SIMULATE_MODE
	selected     bool
}

func NewSimulateButton(rect rl.Rectangle, simulateType SIMULATE_MODE) *SimulateButton {
	return &SimulateButton{
		rect:         rect,
		simulateType: simulateType,
		selected:     false,
	}
}

func (sb *SimulateButton) InRange(v *rl.Vector2) bool {
	return rl.CheckCollisionPointRec(*v, sb.rect)
}

func (window *Window) simulateManager(mousePos *rl.Vector2) []error {
	if window.errorManager.ContainsStongError() {
		return nil
	}
	window.SelectVarButtonOnClick(mousePos)

	for _, sb := range window.simulateModeButton {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if sb.InRange(mousePos) && window.simulationMode != FINISHED {
				window.simulationMode = sb.simulateType
				sb.selected = true
				if window.simulationMode == STEP_BY_STEP {
					err := window.SimulationStep()
					window.errorManager.AppendNew(err)
				}
			}
		}
	}

	var err error
	if window.simulationMode == CONTINUOUSLY {
		err = window.SimulationStep()
		time.Sleep(time.Millisecond * settings.SIMULATION_TIME_STEP_MS)
	} else if window.simulationMode == STEP_BY_STEP {
		if rl.IsKeyPressed(rl.KeyRight) {
			err = window.SimulationStep()
		}
	}
	if err == nil {
		return nil
	}
	return []error{err}
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

func (window *Window) SimulationStep() error {
	finished, consoleMessage, err := window.diagram.MakeStep()
	if consoleMessage != "" {
		window.appendTextToConsole(consoleMessage)
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
	window.errorManager.AppendNew(err)
	if finished {
		window.simulationMode = FINISHED
	}
	return err
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
	for _, vb := range window.simulationVarButtons {
		window.drawSimulateVariableButton(&vb)
	}
}

func (window *Window) drawSimulateVariableButton(vb *VarButton) {
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

func (window *Window) drawSimulateButton(sb *SimulateButton) {
	rl.DrawRectangleRounded(
		sb.rect,
		0.5,
		10,
		settings.SIMULATE_BUTTON_COLOR,
	)
	var color rl.Color
	if sb.selected {
		color = settings.SIMULATE_BUTTON_ATTRIB_COLOR_SELECTED
	} else {
		color = settings.SIMULATE_BUTTON_ATTRIB_COLOR
	}
	switch sb.simulateType {
	case PAUSE:
		rl.DrawRectangle(
			int32(sb.rect.X+sb.rect.Width/2-settings.SIMULATE_BUTTON_PAUSE_SIZE/2),
			int32(sb.rect.Y+sb.rect.Height/2-settings.SIMULATE_BUTTON_PAUSE_SIZE/2),
			settings.SIMULATE_BUTTON_PAUSE_SIZE,
			settings.SIMULATE_BUTTON_PAUSE_SIZE,
			color,
		)
		break
	case STEP_BY_STEP:
		window.drawTriangle(sb.rect.X+sb.rect.Width/2, sb.rect.Y+sb.rect.Height/2, color)
		break
	case CONTINUOUSLY:
		window.drawTriangle(
			sb.rect.X+sb.rect.Width/2-settings.SIMULATE_BUTTON_TRIANGLE_SIZE/4,
			sb.rect.Y+sb.rect.Height/2,
			color,
		)
		window.drawTriangle(
			sb.rect.X+sb.rect.Width/2+settings.SIMULATE_BUTTON_TRIANGLE_SIZE/4,
			sb.rect.Y+sb.rect.Height/2,
			color,
		)
		break
	default:
		break
	}
}

func (window *Window) drawTriangle(offsetX, offsetY float32, color rl.Color) {
	offsetX -= settings.SIMULATE_BUTTON_TRIANGLE_SIZE / 4
	offsetY -= settings.SIMULATE_BUTTON_TRIANGLE_SIZE / 2
	rl.DrawTriangle(
		rl.NewVector2(offsetX, offsetY),
		rl.NewVector2(offsetX, offsetY+settings.SIMULATE_BUTTON_TRIANGLE_SIZE),
		rl.NewVector2(
			offsetX+settings.SIMULATE_BUTTON_TRIANGLE_SIZE/2,
			offsetY+settings.SIMULATE_BUTTON_TRIANGLE_SIZE/2,
		),
		color,
	)
}

func (window *Window) highlightStart() {
	for _, shape := range window.diagramShapes {
		if shape.GetType() == shapes.START {
			shape.SetHighlight(true)
		} else {
			shape.SetHighlight(false)
		}
	}
}

func (window *Window) flushSimulate() {
	window.simulationMode = NOT_SELECTED
	window.simulationVarButtons = make([]VarButton, 0)
	window.simulationVar = ""
	window.clearConsole()
	window.errorManager.Clear()
}

func (window *Window) clearConsole() {
	window.consoleLines = make([]ConsoleLine, 0)
}
