package blocks

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"
)

type IfBlock struct {
	BlockDefault
	nextTrue  *Block
	nextFalse *Block

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

	availableOperators := []string{"==", "<=", "<", ">=", ">", "!="}
	for _, op := range availableOperators {
		tokens := strings.Split(condition, op)
		if len(tokens) == 2 {
			cleanKeyLHS := cleanKeys(tokens[0])
			cleanKeyRHS := cleanKeys(tokens[1])
			b.keys = append(b.keys, cleanKeyLHS...)
			b.keys = append(b.keys, cleanKeyRHS...)
			return nil
		}
	}
	return errors.New("if_block.go/SetConditionExpr fail:\n\tNo valid operator used")
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

	err := exprtkObj.CompileExpression()
	if err != nil {
		return false, err
	}
	for key, val := range b.conditionKVP {
		exprtkObj.SetDoubleVariableValue(key, float64(val))
	}
	if exprtkObj.GetEvaluatedValue() == 0 {
		return false, nil
	}
	return true, nil
}

func (b *IfBlock) FlushCondition() {
	b.conditionExpr = ""
	b.keys = []string{}
}
