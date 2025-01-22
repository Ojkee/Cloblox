package graph_to_code

import (
	"fmt"
	"os"
	"strings"

	"Cloblox/blocks"
	"Cloblox/graph"
)

// taki loop chyba zadziala? nie jestem pewien bo nie wiem do konca jak to przetesotowac
func ConvertGraphToPython(path string, g *graph.Graph) error {
	var pythonCode strings.Builder

	pythonCode.WriteString("def algorythm():\n")
	pythonCode.WriteString("    variables = {}\n\n")

	// mapowanie blokow na pytona
	for _, block := range g.GetAllBlocks() {
		switch blockTyped := block.(type) {
		case *blocks.StartBlock:
			pythonCode.WriteString(fmt.Sprintf("    # Start block (id: %d)\n", block.GetId()))
		case *blocks.StopBlock:
			pythonCode.WriteString(fmt.Sprintf("    # Stop block (id: %d)\n    return\n", block.GetId()))
		case *blocks.ActionBlock:
			pythonCode.WriteString(fmt.Sprintf("    # Action block (id: %d)\n", block.GetId()))
			pythonCode.WriteString(fmt.Sprintf("    %s\n", generateActionCode(blockTyped)))
		case *blocks.IfBlock:
			pythonCode.WriteString(fmt.Sprintf("    # If block (id: %d)\n", block.GetId()))
			condition := generateConditionCode(blockTyped)
			pythonCode.WriteString(fmt.Sprintf("    if %s:\n", condition))
			pythonCode.WriteString("        # True branch\n")
			pythonCode.WriteString("        pass\n") // placeholder dla true
			pythonCode.WriteString("    else:\n")
			pythonCode.WriteString("        # False branch\n")
			pythonCode.WriteString("        pass\n") // placeholder dla false
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

// action bloki (to jest jakies ai generated gowno, narazie to olej xd)
func generateActionCode(block *blocks.ActionBlock) string {
	keys := block.GetKeys()
	return fmt.Sprintf("# Perform action with keys: %v", keys)
}

// if bloki (generateconditioncode i generate branchcode nie jestem pewien czy to bedzie dobrze dzialac
// jakbys mial jakas mysl jak to zrobic lepiej, albo po prostu bys widzial ze cos jest giga zle to daj znac)
func generateConditionCode(block *blocks.IfBlock) string {
	keys := block.GetKeys()
	condition := fmt.Sprintf("condition_based_on(%v)", keys)

	var code strings.Builder
	code.WriteString(fmt.Sprintf("if %s:\n", condition))

	if block.GetNextTrue() != nil {
		code.WriteString("    # True branch\n")
		code.WriteString(generateBranchCode(block.GetNextTrue(), "    "))
	} else {
		code.WriteString("    # True branch is empty\n")
	}

	if block.GetNextFalse() != nil {
		code.WriteString("else:\n")
		code.WriteString("    # False branch\n")
		code.WriteString(generateBranchCode(block.GetNextFalse(), "    "))
	} else {
		code.WriteString("else:\n")
		code.WriteString("    # False branch is empty\n")
	}

	return code.String()
}

func generateBranchCode(block *blocks.Block, indent string) string {
	var code strings.Builder

	switch blockTyped := (*block).(type) {
	case *blocks.IfBlock:
		code.WriteString(indent + generateConditionCode(blockTyped))
	case *blocks.ActionBlock:
		code.WriteString(indent + generateActionCode(blockTyped) + "\n")
	case *blocks.VariablesBlock:
		code.WriteString(indent + generateVariableCode(blockTyped) + "\n")
	default:
		code.WriteString(indent + "# Unsupported block type\n")
	}

	return code.String()
}

// variable bloki (chyba ok, prosze o review)
func generateVariableCode(block *blocks.VariablesBlock) string {
	vars := block.GetVars()
	var code strings.Builder
	for key, value := range vars {
		code.WriteString(fmt.Sprintf("    variables['%s'] = %v\n", key, value))
	}
	return code.String()
}
