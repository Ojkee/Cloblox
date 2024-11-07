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

func (block *StartBlock) GetNext() (*Block, error) {
	return block.next, nil
}

func (block *StartBlock) SetNext(next *Block) {
	block.next = next
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

func (block *StopBlock) GetNext() (*Block, error) {
	return nil, nil
}

func (block *StopBlock) SetNext(next *Block) {
	if next != nil {
		panic("Can't content to Stop block")
	}
	block.next = nil
}
