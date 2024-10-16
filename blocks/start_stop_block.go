package blocks

type StartBlock struct {
	BlockDefault
	next *Block
}

func NewStartBlock() *StartBlock {
	return &StartBlock{
		BlockDefault: BlockDefault{
			id:      -1,
			name:    "start block",
			content: "Start",
		},
		next: nil,
	}
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
	return &StopBlock{
		BlockDefault: BlockDefault{
			id:      -1,
			name:    "stop block",
			content: "Stop",
		},
		next: nil,
	}
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
