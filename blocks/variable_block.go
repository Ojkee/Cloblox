package blocks

// This block is responsible only for declaring variables
// SYNTAX: varible = numerical value

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VariablesBlock struct {
	BlockDefault
	vars   map[string]any // Variables with values on declarations, values should't change
	action func(v float32)

	next *Block
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

func (block *VariablesBlock) SetNext(next *Block) {
	block.next = next
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

// Finds tokens of declaration expression of a variable or slice
// Syntax: [variable] = [value]
// Example x = -0.4 or x_2=3
// Syntax: [variable] = [slice]
// Example x = [2, 3,-3,    0.5, -4.3]
func (block *VariablesBlock) Parse(lines []string) error {
	for _, line := range lines {
		if name, value, ok := block.parseValue(&line); ok {
			block.vars[name] = value
		} else if name, values, ok := block.parseSlice(&line); ok {
			block.vars[name] = values
		} else {
			return errors.New(fmt.Sprintf("Invalid syntax at line:\n\t%s", line))
		}
	}
	return nil
}

func (block *VariablesBlock) parseValue(line *string) (string, float64, bool) {
	rValue := regexp.MustCompile(
		`^\s*(?<variable>[a-zA-Z_][a-zA-Z0-9_]*)\s*=\s*(?<value>-?\d+(\.\d+)?)\s*$`,
	)
	tokens := rValue.FindStringSubmatch(*line)
	if len(tokens) == 0 {
		return "", 0, false
	}
	variableName := tokens[rValue.SubexpIndex("variable")]
	valueString := tokens[rValue.SubexpIndex("value")]
	val, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return "", 0, false
	}
	return variableName, val, true
}

func (block *VariablesBlock) parseSlice(line *string) (string, []float64, bool) {
	rSlice := regexp.MustCompile(
		`^\s*(?<variable>[a-zA-Z_][a-zA-Z0-9_]*)\s*=\s*\[\s*(?<values>-?\d+(\.\d+)?(\s*,\s*-?\d+(\.\d+)?)*\s*)\]\s*$`,
	)
	tokens := rSlice.FindStringSubmatch(*line)
	if len(tokens) == 0 {
		return "", nil, false
	}
	variableName := tokens[rSlice.SubexpIndex("variable")]
	var values []float64
	valuesString := tokens[rSlice.SubexpIndex("values")]
	for _, valString := range strings.Split(valuesString, ",") {
		valTrimmed := strings.TrimSpace(valString)
		val, err := strconv.ParseFloat(valTrimmed, 64)
		if err != nil {
			return "", nil, false
		}
		values = append(values, val)
	}
	return variableName, values, true
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
	switch item := v.(type) {
	case float64:
		return fmt.Sprintf("%.2f", item)
	case float32:
		return fmt.Sprintf("%.2f", item)
	case int:
		return fmt.Sprintf("%d", item)
	case []float64:
		return fmt.Sprintf("%v", item)
	case []float32:
		return fmt.Sprintf("%v", item)
	case []int:
		return fmt.Sprintf("%v", item)
	}
	return "Invalid Type"
}
