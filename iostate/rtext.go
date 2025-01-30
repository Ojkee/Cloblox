package iostate

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"Cloblox/blocks"
	"Cloblox/shapes"
)

// ReadFromTxt parses the file and returns blocks and connections
func ReadFromTxt(path string) ([]blocks.Block, []shapes.Connection, error) {
	// Read the file
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read file %s: %v", path, err)
	}
	inputData := string(fileContent)

	// Regex patterns
	blockPattern := regexp.MustCompile(`<([a-zA-Z0-9]+)>"(.+?)",(\d+)</([a-zA-Z0-9]+)>`)
	connectionPattern := regexp.MustCompile(`<c>([a-z])(\d+),(\d+)</c>`)

	var blocksList []blocks.Block
	var connectionsList []shapes.Connection

	// Parse blocks
	for idx, match := range blockPattern.FindAllStringSubmatch(inputData, -1) {
		if len(match) < 5 {
			return nil, nil, fmt.Errorf("invalid block format")
		}
		tagOpen := match[1]
		content := match[2]
		id, err := strconv.Atoi(match[3])
		tagClose := match[4]

		if tagOpen != tagClose {
			return nil, nil, fmt.Errorf("mismatched tags: %s and %s", tagOpen, tagClose)
		}

		// Przypisanie ID zaczynajac od 1
		id = idx + 1 // Inkrementacja zaczynajac od 1

		params := map[string]any{"id": id}
		switch tagOpen {
		case "h1": // Start block
			params["name"] = content
		case "h2": // Stop block
			params["name"] = content
		case "f": // If block
			params["condition"] = content
		case "v": // Variable block
			variables := make(map[string]float64)
			for _, pair := range strings.Split(content, ",") {
				parts := strings.Split(pair, "=")
				if len(parts) == 2 {
					key := parts[0]
					value, err := strconv.ParseFloat(parts[1], 64)
					if err == nil {
						variables[key] = value
					}
				}
			}
			params["variables"] = variables
		case "a": // Action block
			//fmt.Printf("Parsing Action block: content=%s\n", content) // Sprawdzenie czy dobrze pobiera bo wyrzucalo blad
			params["action"] = content

		default:
			return nil, nil, fmt.Errorf("unknown block type: %s", tagOpen)
		}

		block, err := NewBlockFromTag(tagOpen, params)
		if err != nil {
			return nil, nil, fmt.Errorf("error creating block (tag: %s): %v", tagOpen, err)
		}
		blocksList = append(blocksList, block)
	}

	// Parse connections
	defaultPosX, defaultPosY := float32(0.0), float32(0.0) // Default positions

	for _, match := range connectionPattern.FindAllStringSubmatch(inputData, -1) {
		direction := match[1][0]         // First character of the connection ('r' or 'l')
		id1, _ := strconv.Atoi(match[2]) // ID of the starting shape
		id2, _ := strconv.Atoi(match[3]) // ID of the ending shape

		closerToRight := false
		manyOut := false // Default value for manyOut

		if direction == 'r' {
			closerToRight = false
			manyOut = true // When direction is 'r', treat it as "manyOut" (True branch)
		}
		if direction == 'l' {
			closerToRight = true
			manyOut = false // When direction is 'r', treat it as "manyOut" (True branch)
		}

		// Create connection
		connection := *shapes.NewConnection(
			defaultPosX, defaultPosY, // inPosX, inPosY
			defaultPosX, defaultPosY, // outPosX, outPosY
			id1, id2, // inShapeId, outShapeId
			manyOut, closerToRight, // manyOut, closerToRight
		)

		connectionsList = append(connectionsList, connection)
	}

	return blocksList, connectionsList, nil
}

func NewBlockFromTag(tag string, params map[string]any) (blocks.Block, error) {
	id, ok := params["id"].(int)
	if !ok {
		return nil, fmt.Errorf("missing or invalid ID for block")
	}

	switch tag {
	case "h1":
		block := &blocks.StartBlock{}
		block.SetId(id)
		if name, ok := params["name"].(string); ok {
			block.SetName(name) // Ustawienie nazwy
		}
		return block, nil
	case "h2":
		block := &blocks.StopBlock{}
		block.SetId(id)
		if name, ok := params["name"].(string); ok {
			block.SetName(name) // Ustawienie nazwy
		}
		return block, nil
	case "f":
		condition, ok := params["condition"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid condition for If block")
		}
		block := blocks.NewIfBlock()
		block.SetId(id)
		block.SetConditionExpr(condition)
		return block, nil
	case "a":
		action, ok := params["action"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid action for Action block")
		}
		block := blocks.NewActionBlock()
		block.SetId(id)
		block.ParseFromUserInput(action)
		return block, nil
	case "v":
		variables, ok := params["variables"].(map[string]float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid variables for Variable block")
		}
		block := blocks.NewVariableBlock()
		block.SetId(id)
		for name, value := range variables {
			block.AddVariable(name, value)
		}
		return block, nil
	default:
		return nil, fmt.Errorf("unknown block type: %s", tag)
	}
}
