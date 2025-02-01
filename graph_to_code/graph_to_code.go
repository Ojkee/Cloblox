package graph_to_code

import (
	"fmt"
	"os"
	"strings"

	"Cloblox/blocks"
	"Cloblox/graph"
)

func ConvertGraphToPython(path string, g *graph.Graph) error {
	var pythonCode strings.Builder

	pythonCode.WriteString("def algorythm():\n")
	pythonCode.WriteString("    variables = {}\n\n")

	for _, block := range g.GetAllBlocks() {
		switch blockTyped := block.(type) {
		case *blocks.StartBlock:
			pythonCode.WriteString(fmt.Sprintf("    # Start block (id: %d)\n", block.GetId()))
		case *blocks.StopBlock:
			pythonCode.WriteString(fmt.Sprintf("    # Stop block (id: %d)\n    return 0\n", block.GetId()))
		case *blocks.ActionBlock:
			pythonCode.WriteString(fmt.Sprintf("    # Action block (id: %d)\n", block.GetId()))
			pythonCode.WriteString(fmt.Sprintf("    %s\n", generateActionCode(blockTyped)))
		case *blocks.IfBlock:
			pythonCode.WriteString(fmt.Sprintf("    # If block (id: %d)\n", block.GetId()))
			pythonCode.WriteString(generateIfCode(blockTyped))
		case *blocks.VariablesBlock:
			pythonCode.WriteString(fmt.Sprintf("    # Variables block (id: %d)\n", block.GetId()))
			pythonCode.WriteString(generateVariableCode(blockTyped))
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Error creating new file: %v\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(pythonCode.String())
	if err != nil {
		return fmt.Errorf("Error saving to file: %v\n", err)
	}
	return nil
}

func generateActionCode(block *blocks.ActionBlock) string {
	keys := block.GetKeys()
	var actionType blocks.ACTION_TYPE
	blockRawContent := block.GetActionInputRaw()
	actionType, err := block.GetActionType(&blockRawContent)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	var code strings.Builder

	switch actionType {
	case blocks.SWAP:
		code.WriteString(fmt.Sprintf("%s, %s = %s, %s\n", keys[0], keys[1], keys[1], keys[0]))
	case blocks.PRINT:
		code.WriteString("print(")
		for i, key := range keys {
			if i > 0 {
				code.WriteString(", ")
			}
			code.WriteString(key)
		}
		code.WriteString(")\n")
	case blocks.RAND:
		code.WriteString(fmt.Sprintf("%s = random.random()\n", keys[0]))
	case blocks.MATH_OPERATIONS:
		code.WriteString(fmt.Sprintf("%s\n", block.GetActionInputRaw()))
	default:
		code.WriteString("Unknown action type\n")
	}

	return code.String()
}

// not quite right yet but we'll get there
func generateIfCode(block *blocks.IfBlock) string {
	condition := block.GetConditionExprString()

	var code strings.Builder

	code.WriteString(fmt.Sprintf("    if %s:\n", condition))

	if block.GetNextTrue() != nil {
		code.WriteString(generateBranchCode(block.GetNextTrue(), "    "))
	} else {
		code.WriteString("        # True branch is empty\n")
		code.WriteString("        pass\n")
	}

	if block.GetNextFalse() != nil {
		code.WriteString("    else:\n")
		code.WriteString(generateBranchCode(block.GetNextFalse(), "    "))
	} else {
		code.WriteString("    else:\n")
		code.WriteString("        # False branch is empty\n")
		code.WriteString("        pass\n")
	}

	return code.String()
}

func generateBranchCode(block *blocks.Block, indent string) string {
	var code strings.Builder

	switch blockTyped := (*block).(type) {
	case *blocks.IfBlock:
		code.WriteString(indent + generateIfCode(blockTyped))
	// case *blocks.ActionBlock:
	// 	code.WriteString(indent + generateActionCode(blockTyped) + "\n")
	case *blocks.VariablesBlock:
		code.WriteString(indent + generateVariableCode(blockTyped) + "\n")
	default:
		code.WriteString(indent + "# Unsupported block type\n")
	}

	return code.String()
}

func generateVariableCode(block *blocks.VariablesBlock) string {
	vars := block.GetVars()
	var code strings.Builder
	for key, value := range vars {
		code.WriteString(fmt.Sprintf("    variables['%s'] = %v\n", key, value))
	}
	return code.String()
}
