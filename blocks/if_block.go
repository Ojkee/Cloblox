package blocks

import (
	"fmt"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"

	"Cloblox/functools"
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

func (block *IfBlock) IsCloserToRigth() bool {
	panic("unimplemented")
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

func (block *IfBlock) GetNext() (*Block, error) {
	isTrue, err := block.IsEvalTrue()
	if err != nil {
		return nil, err
	}
	if isTrue {
		return block.nextTrue, nil
	}
	return block.nextFalse, nil
}

func (block *IfBlock) GetNextTrue() *Block {
	return block.nextTrue
}

func (block *IfBlock) GetNextFalse() *Block {
	return block.nextFalse
}

func (block *IfBlock) SetNextTrue(next *Block) {
	block.nextTrue = next
}

func (block *IfBlock) SetNextFalse(next *Block) {
	block.nextFalse = next
}

func (block *IfBlock) GetKeys() []string {
	return block.keys
}

func (block *IfBlock) SetConditionExpr(condition string) error {
	block.conditionExpr = condition
	exprReplaced, arrKeys := findReplaceArrayKeys(condition, block.replaceKey)
	block.arrayKeys = arrKeys

	availableOperators := []string{"==", "<=", "<", ">=", ">", "!="}
	for _, op := range availableOperators {
		if strings.Contains(exprReplaced, op) {
			block.keys = getKeysFromString(&block.conditionExpr)
			block.conditionExprReplaced = replaceLogicOps(&exprReplaced)
			return nil
		}
	}
	if condition == "" {
		condition = "If"
	}
	return functools.NewStrongError(
		fmt.Sprintf("Can't compile line: %s", condition),
		"if_block.go/SetConditionExpr fail:\n\tNo valid operator used",
	)
}

func replaceLogicOps(line *string) string {
	replaced := strings.Replace(*line, "&&", "and", -1)
	replaced = strings.Replace(replaced, "||", "or", -1)
	return replaced
}

func (block *IfBlock) SetConditionKVP(kvp *map[string]float64) {
	block.conditionKVP = *kvp
}

func (block *IfBlock) IsEvalTrue() (bool, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(block.conditionExprReplaced)
	for _, key := range block.keys {
		exprtkObj.AddDoubleVariable(key)
	}

	for key := range block.conditionKVP {
		if arrayKey, ok := block.arrayKeys[key]; ok {
			exprtkObj.AddDoubleVariable(arrayKey)
		} else {
			exprtkObj.AddDoubleVariable(key)
		}
	}

	err := exprtkObj.CompileExpression()
	if err != nil {
		consoleMessage := fmt.Sprintf("Can't compile: %s", block.conditionExpr)
		return false, functools.NewStrongError(consoleMessage, err.Error())
	}
	for key, val := range block.conditionKVP {
		if arrayKey, ok := block.arrayKeys[key]; ok {
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

func (block *IfBlock) FlushCondition() {
	block.conditionExpr = ""
	block.conditionExprReplaced = ""
	block.keys = []string{}
	block.arrayKeys = make(map[string]string, 0)
}

func (block *IfBlock) GetConditionExprString() string {
	return block.conditionExpr
}
