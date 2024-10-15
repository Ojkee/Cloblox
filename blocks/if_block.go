package blocks

type IfBlock[T BlockType] struct {
	BlockDefault
	nextTrue          *Block
	nextFalse         *Block
	conditionFunction func(x T) bool
}

func NewIfBlock[T BlockType]() *IfBlock[T] {
	retVal := IfBlock[T]{
		BlockDefault: BlockDefault{
			id:      blockCounter,
			name:    "if block",
			content: "if",
		},
		nextTrue:  nil,
		nextFalse: nil,
	}
	blockCounter += 1
	return &retVal
}

func (b *IfBlock[T]) GetNext(args ...any) *Block {
	if len(args) == 0 {
		return b.nextFalse
	}
	if val, ok := args[0].(T); ok {
		if b.conditionFunction(val) {
			return b.nextTrue
		}
		return b.nextFalse
	}
	return nil
}

func (b *IfBlock[T]) GetNextTrue() *Block {
	return b.nextTrue
}

func (b *IfBlock[T]) GetNextFalse() *Block {
	return b.nextFalse
}

func (b *IfBlock[T]) SetNextTrue(next Block) {
	b.nextTrue = &next
}

func (b *IfBlock[T]) SetNextFalse(next Block) {
	b.nextFalse = &next
}

func (b *IfBlock[T]) ParseCondition(condition string) error {
	// TODO
	return nil
}

func (b *IfBlock[T]) SetCondition(condition func(x T) bool) { // Devtool
	b.conditionFunction = condition
}
