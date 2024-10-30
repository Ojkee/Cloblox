package blocks

// This block is responsible only for declaring variables
// SYNTAX: varible = numerical value

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type VariablesBlock struct {
	BlockDefault
	vars   map[string]any // Variables with values on declarations, values should't change
	action func(v float32)
	next   *Block
}

func NewVariableBlock() *VariablesBlock {
	return &VariablesBlock{
		BlockDefault: BlockDefault{
			id:   -1,
			name: "Variable Block",
		},
		vars: make(map[string]any),
		next: nil,
	}
}

func (block *VariablesBlock) GetNext() (*Block, error) {
	return block.next, nil
}

func (block *VariablesBlock) SetNext(next Block) {
	block.next = &next
}

func (block *VariablesBlock) AddVariable(name string, value any) {
	if len(strings.TrimSpace(name)) != 0 {
		block.vars[name] = value
	}
}

func (block *VariablesBlock) GetValue(variableName string) (string, error) {
	if val, ok := block.vars[variableName]; ok {
		return block.valueToString(val), nil
	}
	return "", errors.New(block.name + " fail:\t\nNo such variable")
}

func (block *VariablesBlock) Parse(lines []string) error {
	for _, line := range lines {
		lineTokens := strings.Split(line, " ")
		if len(lineTokens) != 3 {
			return errors.New(fmt.Sprintf(
				`varible_block.go/Parse fail:
        \n\tInvalid input in line: %s`, line),
			)
		}
		if lineTokens[1] != "=" {
			return errors.New(fmt.Sprintf(
				`varible_block.go/Parse fail:
        Invalid Syntax
        Syntax: variable = values`),
			)
		}
		if parsed, err := strconv.ParseFloat(lineTokens[0], 10); err == nil {
			return errors.New(fmt.Sprintf(
				`varible_block.go/Parse fail:
        Number '%f' can't be variable
        Syntax: variable = values`, parsed),
			)
		}
		if parsed, err := strconv.ParseFloat(lineTokens[2], 10); err == nil {
			block.AddVariable(lineTokens[0], parsed)
		} else {
			return errors.New(fmt.Sprintf(
				`varible_block.go/Parse fail:
        Non-number '%s' instance can't be value
        Syntax: variable = values`, lineTokens[2]),
			)
		}
	}
	return nil
}

func (block *VariablesBlock) GetVars() map[string]any {
	return block.vars
}

func (block *VariablesBlock) varsFormatted() string {
	var buffur bytes.Buffer
	for key, value := range block.vars {
		line := fmt.Sprintf("%s = %s", key, block.valueToString(value))
		buffur.WriteString(line)
	}
	return buffur.String()
}

func (block *VariablesBlock) valueToString(v any) string {
	if floatVal, ok := v.(float32); ok {
		return fmt.Sprintf("%.2f", floatVal)
	} else if _, ok := v.([]float32); ok {
		return "TBD"
	}
	return "Invalid Type"
}
