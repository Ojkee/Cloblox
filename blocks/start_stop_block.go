package blocks

type StartBlock struct {
	BlockDefault
	next *Block
}

func NewStartBlock() *StartBlock {
	retVal := StartBlock{
		BlockDefault: BlockDefault{
			id:      blockCounter,
			name:    "start block",
			content: "Start",
		},
		next: nil,
	}
	blockCounter += 1
	return &retVal
}

func (b *StartBlock) GetNext(args ...any) *Block {
	return b.next
}

func (b *StartBlock) SetNext(next Block) {
	b.next = &next
}

type StopBlock struct {
	BlockDefault
	next *Block
}

func NewStopBlock() *StopBlock {
	retVal := StopBlock{
		BlockDefault: BlockDefault{
			id:      blockCounter,
			name:    "stop block",
			content: "Stop",
		},
		next: nil,
	}
	blockCounter += 1
	return &retVal
}

func (b *StopBlock) GetNext(args ...any) *Block {
	return nil
}

func (b *StopBlock) SetNext(next Block) {
	if next != nil {
		panic("Can't content to Stop block")
	}
	b.next = nil
}
