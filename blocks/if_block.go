package blocks

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"
)

type IfBlock struct {
	BlockDefault
	nextTrue  *Block
	nextFalse *Block

	replaceKey            string             // key to replace every array array key for parsing
	keys                  []string           // every array with index and non-array var
	arrayKeys             map[string]string  // key: array with index val: symbole
	conditionKVP          map[string]float64 // key: array with index and non-array var val: value
	conditionExpr         string             // non edited expr
	conditionExprReplaced string             // replaced with replaceKey and set rounding
}

func NewIfBlock() *IfBlock {
	return &IfBlock{
		BlockDefault: BlockDefault{
			id:   -1,
			name: "If Block",
		},
		nextTrue:  nil,
		nextFalse: nil,

		replaceKey:            "TEMPKEY",
		keys:                  make([]string, 0),
		arrayKeys:             make(map[string]string),
		conditionKVP:          make(map[string]float64),
		conditionExpr:         "",
		conditionExprReplaced: "",
	}
}

func (b *IfBlock) GetNext() (*Block, error) {
	isTrue, err := b.IsEvalTrue()
	if err != nil {
		return nil, err
	}
	if isTrue {
		return b.nextTrue, nil
	}
	return b.nextFalse, nil
}

func (b *IfBlock) GetNextTrue() *Block {
	return b.nextTrue
}

func (b *IfBlock) GetNextFalse() *Block {
	return b.nextFalse
}

func (b *IfBlock) SetNextTrue(next Block) {
	b.nextTrue = &next
}

func (b *IfBlock) SetNextFalse(next Block) {
	b.nextFalse = &next
}

func (b *IfBlock) GetKeys() []string {
	return b.keys
}

func (b *IfBlock) SetConditionExpr(condition string) error {
	b.conditionExpr = condition
	exprReplaced := b.findReplaceArrayKeys(condition)

	availableOperators := []string{"==", "<=", "<", ">=", ">", "!="}
	for _, op := range availableOperators {
		tokens := strings.Split(exprReplaced, op)
		if len(tokens) == 2 {
			b.keys = rawKeys(b.conditionExpr)
			b.conditionExprReplaced = setRound(tokens[0]) + op + setRound(tokens[1])
			return nil
		}
	}
	return errors.New("if_block.go/SetConditionExpr fail:\n\tNo valid operator used")
}

func (b *IfBlock) SetConditionKVP(kvp *map[string]float64) {
	b.conditionKVP = *kvp
}

func (b *IfBlock) IsEvalTrue() (bool, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(b.conditionExprReplaced)
	for _, key := range b.keys {
		exprtkObj.AddDoubleVariable(key)
	}

	for key := range b.conditionKVP {
		if arrayKey, ok := b.arrayKeys[key]; ok {
			exprtkObj.AddDoubleVariable(arrayKey)
		} else {
			exprtkObj.AddDoubleVariable(key)
		}
	}

	err := exprtkObj.CompileExpression()
	if err != nil {
		return false, err
	}
	for key, val := range b.conditionKVP {
		if arrayKey, ok := b.arrayKeys[key]; ok {
			exprtkObj.SetDoubleVariableValue(arrayKey, float64(val))
		} else {
			exprtkObj.SetDoubleVariableValue(key, float64(val))
		}
	}
	if exprtkObj.GetEvaluatedValue() == 0 {
		return false, nil
	}
	return true, nil
}

func (b *IfBlock) FlushCondition() {
	b.conditionExpr = ""
	b.conditionExprReplaced = ""
	b.keys = []string{}
	b.arrayKeys = make(map[string]string, 0)
}

func (b *IfBlock) findReplaceArrayKeys(input string) string {
	// a[4]         // In    arr[i]       // In    arr[x + i] // In
	// tab[dd]      // In    [dkf]        // Out   arr[x + 2] // In
	// [4]          // Out   myArray[45]  // In
	// dkjf[4k]     // Out   my_array[df] // In
	// t[my_x]      // In    dlsk[df sf]  // Out
	r := regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*\[(?:[a-zA-Z0-9_+\-*/\s()]+)\]`)
	arrayKeysFound := r.FindAllString(input, -1)
	for i, key := range arrayKeysFound {
		nextReplaceKey := fmt.Sprintf("%s%d", b.replaceKey, i)
		input = strings.ReplaceAll(input, key, nextReplaceKey)
		b.arrayKeys[key] = nextReplaceKey
	}
	return input
}

func rawKeys(input string) []string {
	// Searching arrays that may has non-leading numbers
	// such as my_tab2[(i-1)*2], but they aren't non-array variables
	// such as x, i2, my_x etc.
	r := regexp.MustCompile(
		`[a-zA-Z_][a-zA-Z0-9_]*\[[a-zA-Z0-9_+\-*/\s()]*\]|[a-zA-Z_][a-zA-Z0-9_]*`,
	)
	keysFound := r.FindAllString(input, -1)
	return keysFound
}

func setRound(inside string) string {
	return "roundn(" + inside + ", 10)"
}
