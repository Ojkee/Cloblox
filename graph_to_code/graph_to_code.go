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

	pythonCode.WriteString("def algorithm():\n")
	pythonCode.WriteString("    variables = {}\n\n")

	startBlock := g.GetHead()
	if startBlock == nil {
		return fmt.Errorf("no start block found")
	}

	visited := make(map[blocks.Block]bool)
	order := make(map[blocks.Block]int)
	err := processBlock(*startBlock, &pythonCode, "    ", visited, order, nil)
	if err != nil {
		return err
	}

	pythonCode.WriteString("    return variables\n")

	return os.WriteFile(path, []byte(pythonCode.String()), 0644)
}

func processBlock(
	block blocks.Block,
	pythonCode *strings.Builder,
	indent string,
	visited map[blocks.Block]bool,
	order map[blocks.Block]int,
	parent *blocks.Block,
) error {
	if visited[block] {
		return nil
	}
	order[block] = len(order)
	visited[block] = true

	switch blockTyped := block.(type) {
	case *blocks.StartBlock:
		pythonCode.WriteString(fmt.Sprintf("%s# Start block (id: %d)\n", indent, block.GetId()))
	case *blocks.StopBlock:
		pythonCode.WriteString(fmt.Sprintf("%s# Stop block (id: %d)\n%s\n", indent, block.GetId(), indent))
		visited[block] = false
	case *blocks.ActionBlock:
		pythonCode.WriteString(fmt.Sprintf("%s# Action block (id: %d)\n", indent, block.GetId()))
		pythonCode.WriteString(fmt.Sprintf("%s%s\n", indent, generateActionCode(blockTyped)))
	case *blocks.IfBlock:
		pythonCode.WriteString(fmt.Sprintf("%s# If block (id: %d)\n", indent, block.GetId()))
		if hasBackwardConnection(blockTyped, order) {
			pythonCode.WriteString(fmt.Sprintf("%swhile %s:\n", indent, blockTyped.GetConditionExprString()))
			indent += "    "
		} else {
			pythonCode.WriteString(fmt.Sprintf("%sif %s:\n", indent, blockTyped.GetConditionExprString()))
		}
		err := generateIfCode(blockTyped, pythonCode, order, indent, visited, parent)
		if err != nil {
			return err
		}
		indent = indent[:len(indent)-4]
	case *blocks.VariablesBlock:
		pythonCode.WriteString(fmt.Sprintf("%s# Variables block (id: %d)\n", indent, block.GetId()))
		pythonCode.WriteString(fmt.Sprintf("%s%s\n", indent, generateVariableCode(blockTyped)))
	default:
		return fmt.Errorf("unsupported block type")
	}

	nextBlock, err := block.GetNext()
	if err != nil {
		return err
	}
	if nextBlock != nil {
		return processBlock(*nextBlock, pythonCode, indent, visited, order, &block)
	}

	return nil
}

func hasBackwardConnection(block *blocks.IfBlock, order map[blocks.Block]int) bool {
	trueBlock := block.GetNextTrue()
	falseBlock := block.GetNextFalse()

	if trueBlock != nil && order[*trueBlock] < order[block] {
		return true
	}
	if falseBlock != nil && order[*falseBlock] < order[block] {
		return true
	}
	return false
}

func generateIfCode(
	block *blocks.IfBlock,
	pythonCode *strings.Builder,
	order map[blocks.Block]int,
	indent string,
	visited map[blocks.Block]bool,
	parent *blocks.Block,
) error {
	trueBlock := block.GetNextTrue()
	falseBlock := block.GetNextFalse()

	pythonCode.WriteString(fmt.Sprintf("%sif %s:\n", indent, block.GetConditionExprString()))
	if trueBlock != nil {
		err := processBlock(*trueBlock, pythonCode, indent+"    ", visited, order, parent)
		if err != nil {
			return err
		}
	}

	if falseBlock != nil {
		pythonCode.WriteString(fmt.Sprintf("%selse:\n", indent))
		err := processBlock(*falseBlock, pythonCode, indent+"    ", visited, order, parent)
		if err != nil {
			return err
		}
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

func generateVariableCode(block *blocks.VariablesBlock) string {
	vars := block.GetVars()
	var code strings.Builder
	for key, value := range vars {
		code.WriteString(fmt.Sprintf("variables['%s'] = %v\n    ", key, value))
	}
	return code.String()
}
