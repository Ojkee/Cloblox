package blocks

import "errors"

type IfBlock[T BlockType] struct {
	BlockDefault
	nextTrue          *Block
	nextFalse         *Block
	conditionFunction func(x T) bool
}

func NewIfBlock[T BlockType](id int) *IfBlock[T] {
	return &IfBlock[T]{
		BlockDefault: BlockDefault{
			id:      id,
			name:    "if",
			content: "if",
		},
		nextTrue:  nil,
		nextFalse: nil,
	}
}

func (b *IfBlock[T]) GetContent() string {
	return b.content
}

func (b *IfBlock[T]) ParseCondition(condition string) error {
	// TODO
	return nil
}

func (b *IfBlock[T]) SetCondition(condition func(x T) bool) { // Devtool
	b.conditionFunction = condition
}

func (b *IfBlock[T]) SetTrue(nextBlock *Block) {
	b.nextTrue = nextBlock
}

func (b *IfBlock[T]) SetFalse(nextBlock *Block) {
	b.nextFalse = nextBlock
}

func (b *IfBlock[T]) GetNext(args ...any) *Block {
	if len(args) == 0 {
		return nil
	}
	if val, ok := args[0].(T); ok {
		if b.conditionFunction(val) {
			return b.nextTrue
		}
		return b.nextFalse
	}
	return nil
}
