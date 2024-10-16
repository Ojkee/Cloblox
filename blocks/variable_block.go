package blocks

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type VariablesBlock[T BlockType] struct {
	BlockDefault
	vars map[string]T // Variables with values on declarations, values should't change
	next *Block
}

func NewVariableBlock[T BlockType]() *VariablesBlock[T] {
	return &VariablesBlock[T]{
		BlockDefault: BlockDefault{
			id:      -1,
			name:    "variable block",
			content: "variable block",
		},
		vars: make(map[string]T),
		next: nil,
	}
}

func (b *VariablesBlock[T]) GetNext(args ...any) *Block {
	return b.next
}

func (b *VariablesBlock[T]) SetNext(next Block) {
	b.next = &next
}

func (b *VariablesBlock[T]) UpdateContent() {
	b.content = b.varsFormatted()
}

func (b *VariablesBlock[T]) AddVariable(name string, value T) {
	if len(strings.TrimSpace(name)) != 0 {
		b.vars[name] = value
	}
}

func (b *VariablesBlock[T]) GetValue(variableName string) (string, error) {
	if val, ok := b.vars[variableName]; ok {
		return b.valueToString(val), nil
	}
	return "", errors.New(b.content + " fail:\t\nNo such variable")
}

func (b *VariablesBlock[T]) varsFormatted() string {
	var buffur bytes.Buffer
	for key, value := range b.vars {
		line := fmt.Sprintf("%s = %s", key, b.valueToString(value))
		buffur.WriteString(line)
	}
	return buffur.String()
}

func (b *VariablesBlock[T]) valueToString(v T) string {
	if retVal, ok := any(v).(float32); ok {
		return fmt.Sprintf("%.2f", retVal)
	} else if retVal, ok := any(v).(int); ok {
		return fmt.Sprintf("%d", retVal)
	} else if retVal, ok := any(v).(string); ok {
		return retVal
	}
	return "Invalid Type"
}
