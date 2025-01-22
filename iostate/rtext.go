package iostate

// import (
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"regexp"
// 	"strconv"
// 	"strings"
//
// 	"github.com/sqweek/dialog"
//
// 	"Cloblox/blocks"
// 	"Cloblox/graph"
// 	"Cloblox/shapes"
// )
//
// func ReadFromTxt(blocks *[]graph.Graph, connections *[]shapes.Connection) error {
// 	// Open window to choose the file
// 	filePath, err := dialog.File().Title("Select a file to read").Load()
// 	if err != nil {
// 		log.Fatalf("Failed to select file: %v", err)
// 	}
// 	fmt.Printf("Selected file: %s\n", filePath)
//
// 	// Read file
// 	fileBuffer, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		log.Fatalf("Failed to read file %s: %v", filePath, err)
// 	}
// 	inputData := string(fileBuffer)
//
// 	// Regex patterns
// 	blockPattern := regexp.MustCompile(`<(?P<tag>[a-zA-Z0-9]+)>"(?P<content>.+?)",(?P<id>\d+)</\1>`)
// 	connectionPattern := regexp.MustCompile(`<c>(?P<id1>\d+),(?P<id2>\d+)</c>`)
//
// 	// Parse blocks
// 	for _, match := range blockPattern.FindAllStringSubmatch(inputData, -1) {
// 		tag := match[1]
// 		content := match[2]
//
// 		// Create params map based on tag
// 		params := make(map[string]any)
// 		switch tag {
// 		case "Start", "Stop":
// 			// No additional parameters for Start and Stop blocks
// 		case "If":
// 			params["condition"] = content
// 		case "Action":
// 			params["action"] = content
// 		case "Variable":
// 			// Parse variables from content, e.g., "x=5,y=10"
// 			variables := make(map[string]float64)
// 			for _, pair := range strings.Split(content, ",") {
// 				parts := strings.Split(pair, "=")
// 				if len(parts) == 2 {
// 					key := parts[0]
// 					value, err := strconv.ParseFloat(parts[1], 64)
// 					if err == nil {
// 						variables[key] = value
// 					}
// 				}
// 			}
// 			params["variables"] = variables
// 		default:
// 			log.Printf("Warning: Unrecognized block type: %s\n", tag)
// 			continue
// 		}
//
// 		// Call NewBlockFromTag
// 		block1, err := NewBlockFromTag(tag, params)
// 		if err != nil {
// 			log.Printf("Error creating block (tag: %s): %v\n", tag, err)
// 			continue
// 		}
//
// 		*blocks = append(*blocks, block1)
// 	}
//
// 	// Parse connections
// 	defaultPosX, defaultPosY := float32(0.0), float32(0.0) // Default positions for connections
// 	for _, match := range connectionPattern.FindAllStringSubmatch(inputData, -1) {
// 		id1, _ := strconv.Atoi(match[1])
// 		id2, _ := strconv.Atoi(match[2])
//
// 		// Create connection with default positions and other parameters
// 		connection := shapes.NewConnection(
// 			defaultPosX, defaultPosY, // inPosX, inPosY
// 			defaultPosX, defaultPosY, // outPosX, outPosY
// 			id1, id2, // inShapeId, outShapeId
// 			false, false, // multipleOut, closerToRight
// 		)
// 		*connections = append(*connections, connection)
// 	}
//
// 	fmt.Println("File loaded successfully.")
// 	return nil
// }
//
// func NewBlockFromTag(tag string, params map[string]any) (blocks.Block, error) {
// 	if strings.TrimSpace(tag) == "" {
// 		return nil, errors.New("Tag cannot be empty")
// 	}
//
// 	switch tag {
// 	case "Start":
// 		return &blocks.StartBlock{}, nil
// 	case "Stop":
// 		return &blocks.StopBlock{}, nil
// 	case "If":
// 		conditionInput, ok := params["condition"].(string)
// 		if !ok {
// 			return nil, errors.New("Missing or invalid condition parameter for If block")
// 		}
// 		conditionBlock := blocks.NewIfBlock()
// 		for err := range conditionInput {
// 			conditionBlock.SetConditionExpr(name)
// 		}
// 		return conditionBlock, nil
// 	case "Action":
// 		actionInput, ok := params["action"].(string)
// 		if !ok {
// 			return nil, errors.New("Missing or invalid action parameter for Action block")
// 		}
// 		actionBlock := blocks.NewActionBlock()
// 		if err := actionBlock.ParseFromUserInput(actionInput); err != nil {
// 			return nil, err
// 		}
// 		return actionBlock, nil
// 	case "Variable":
// 		variables, ok := params["variables"].(map[string]float64)
// 		if !ok {
// 			return nil, errors.New("Missing or invalid variables parameter for Variable block")
// 		}
// 		variableBlock := blocks.NewVariableBlock()
// 		for name, value := range variables {
// 			variableBlock.AddVariable(name, value)
// 		}
// 		return variableBlock, nil
// 	default:
// 		return nil, fmt.Errorf("Unknown block type: %s", tag)
// 	}
// }
