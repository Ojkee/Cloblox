package blocks

import (
	"bytes"
	"errors"
	"fmt"
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

func (b *VariablesBlock) GetNext(args ...float32) *Block {
	return b.next
}

func (b *VariablesBlock) SetNext(next Block) {
	b.next = &next
}

func (b *VariablesBlock) AddVariable(name string, value any) {
	if len(strings.TrimSpace(name)) != 0 {
		b.vars[name] = value
	}
}

func (b *VariablesBlock) GetValue(variableName string) (string, error) {
	if val, ok := b.vars[variableName]; ok {
		return b.valueToString(val), nil
	}
	return "", errors.New(b.name + " fail:\t\nNo such variable")
}

func (b *VariablesBlock) Parse(input string) error {
	// TODO
	return nil
}

func (b *VariablesBlock) GetVars() map[string]any {
	return b.vars
}

func (b *VariablesBlock) varsFormatted() string {
	var buffur bytes.Buffer
	for key, value := range b.vars {
		line := fmt.Sprintf("%s = %s", key, b.valueToString(value))
		buffur.WriteString(line)
	}
	return buffur.String()
}

func (b *VariablesBlock) valueToString(v any) string {
	if floatVal, ok := v.(float32); ok {
		return fmt.Sprintf("%.2f", floatVal)
	} else if _, ok := v.([]float32); ok {
		return "TBD"
	}
	return "Invalid Type"
}
