package blocks

type StartBlock struct {
	BlockDefault
	next *Block
}

func NewStartBlock() *StartBlock {
	return &StartBlock{
		BlockDefault: BlockDefault{
			id:   -1,
			name: "Start Block",
		},
		next: nil,
	}
}

func (b *StartBlock) GetNext() (*Block, error) {
	return b.next, nil
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
			id:   -1,
			name: "Stop Block",
		},
		next: nil,
	}
}

func (b *StopBlock) GetNext() (*Block, error) {
	return nil, nil
}

func (b *StopBlock) SetNext(next Block) {
	if next != nil {
		panic("Can't content to Stop block")
	}
	b.next = nil
}
