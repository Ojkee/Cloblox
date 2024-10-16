package blocks

import (
	"errors"
	"strings"
)

type IfBlock struct {
	BlockDefault
	nextTrue  *Block
	nextFalse *Block

	keyLHS string
	keyRHS string

	conditionFunc func(lhs, rhs float32) bool
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

func (b *IfBlock) GetNext(args ...float32) *Block {
	if len(args) <= 2 {
		panic("IfBlock/GetNext fail:\n\tToo little arguments provided")
	}
	if b.conditionFunc(args[0], args[1]) {
		return b.nextTrue
	}
	return b.nextFalse
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

func (b *IfBlock) ParseCondition(condition string) error {
	exprStrings := strings.Split(condition, " ")
	if len(exprStrings) < 3 {
		return errors.New("IfBlock/ParseCondition fail:\n\tInvalid argument provided")
	}
	b.keyLHS = exprStrings[0]
	b.keyRHS = exprStrings[2]
	switch exprStrings[1] {
	case "=", "==", "eq":
		b.conditionFunc = func(lhs, rhs float32) bool { return lhs == rhs }
		break
	case "<":
		b.conditionFunc = func(lhs, rhs float32) bool { return lhs < rhs }
		break
	case "<=":
		b.conditionFunc = func(lhs, rhs float32) bool { return lhs <= rhs }
		break
	case ">":
		b.conditionFunc = func(lhs, rhs float32) bool { return lhs > rhs }
		break
	case ">=":
		b.conditionFunc = func(lhs, rhs float32) bool { return lhs >= rhs }
		break
	case "!=", "neq":
		b.conditionFunc = func(lhs, rhs float32) bool { return lhs != rhs }
		break
	default:
		return errors.New("IfBlock/ParseCondition fail:\n\tInvalid operator")
	}
	return nil
}

func (b *IfBlock) GetKeys() (string, string) {
	return b.keyLHS, b.keyRHS
}

func (b *IfBlock) SetCondition(conditionFunc func(lhs, rhs float32) bool) {
	b.conditionFunc = conditionFunc
}

func (b *IfBlock) Compare(lhs, rhs float32) bool {
	return b.conditionFunc(lhs, rhs)
}
