package blocks

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"
)

type IfBlock struct {
	BlockDefault
	nextTrue  *Block
	nextFalse *Block

	replaceKey    string
	arrayKeys     map[string]string // key: array with index val: symbole
	keys          []string
	conditionKVP  map[string]float32
	conditionExpr string
}

func NewIfBlock() *IfBlock {
	return &IfBlock{
		BlockDefault: BlockDefault{
			id:   -1,
			name: "If Block",
		},
		nextTrue:  nil,
		nextFalse: nil,

		replaceKey:    "TEMPKEY",
		arrayKeys:     make(map[string]string),
		keys:          make([]string, 0),
		conditionKVP:  make(map[string]float32),
		conditionExpr: "",
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
	retVal := make([]string, 0)
	for _, key := range b.keys {
		if arrayKey, ok := b.arrayKeys[key]; ok {
			retVal = append(retVal, arrayKey)
		} else {
			retVal = append(retVal, key)
		}
	}
	return retVal
}

func (b *IfBlock) SetConditionExpr(condition string) error {
	b.conditionExpr = b.findReplaceArrayKeys(condition)

	availableOperators := []string{"==", "<=", "<", ">=", ">", "!="}
	for _, op := range availableOperators {
		tokens := strings.Split(b.conditionExpr, op)
		if len(tokens) == 2 {
			cleanKeys := cleanKeys(b.conditionExpr)
			b.keys = append(b.keys, cleanKeys...)
			return nil
		}
	}
	return errors.New("if_block.go/SetConditionExpr fail:\n\tNo valid operator used")
}

func (b *IfBlock) findReplaceArrayKeys(input string) string {
	// a[4]         // In    arr[i]       // In
	// tab[dd]      // In    [dkf]        // Out
	// [4]          // Out   myArray[45]  // In
	// dkjf[4k]     // Out   my_array[df] // In
	// t[my_x]      // In    dlsk[df sf]  // Out
	r := regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*\[(?:[a-zA-Z_][a-zA-Z0-9_]*|[0-9]+)\]`)
	arrayKeysFound := r.FindAllString(input, -1)
	for i, key := range arrayKeysFound {
		nextReplaceKey := b.replaceKey + strconv.FormatInt(int64(i), 10)
		input = strings.ReplaceAll(input, key, nextReplaceKey)
		b.arrayKeys[key] = nextReplaceKey
	}
	return input
}

func cleanKeys(input string) []string {
	// Searching variables that may has non-leading numbers
	// or/and array variables such as a[i]
	r := regexp.MustCompile(`[a-zA-Z]+(?:\[[0-9a-z]+\])?`)
	keysFound := r.FindAllString(input, -1)
	return keysFound
}

func (b *IfBlock) SetConditionKVP(kvp *map[string]float32) {
	b.conditionKVP = *kvp
}

func (b *IfBlock) IsEvalTrue() (bool, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(b.conditionExpr)
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
	b.keys = []string{}
	b.arrayKeys = make(map[string]string, 0)
}
