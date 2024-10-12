package blocks

type StartBlock struct {
	BlockDefault
	next *Block
}

func NewStartBlock(id int) *StartBlock {
	return &StartBlock{
		BlockDefault: BlockDefault{
			id:      id,
			name:    "start block",
			content: "",
		},
		next: nil,
	}
}

func (b *StartBlock) GetContent() string {
	return "Start"
}

func (b *StartBlock) GetNext(args ...any) *Block {
	return b.next
}

type StopBlock struct {
	BlockDefault
	next *Block
}

func NewStopBlock(id int) *StopBlock {
	return &StopBlock{
		BlockDefault: BlockDefault{
			id:      id,
			name:    "stop block",
			content: "",
		},
		next: nil,
	}
}

func (b *StopBlock) GetContent() string {
	return "Stop"
}

func (b *StopBlock) GetNext(args ...any) *Block {
	return nil
}
