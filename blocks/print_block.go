package blocks

type PrintBlock struct {
	BlockDefault
	next *Block
}

func NewPrintBlock() *PrintBlock {
	retVal := PrintBlock{
		BlockDefault: BlockDefault{
			id:      blockCounter,
			name:    "print block",
			content: "print",
		},
		next: nil,
	}
	blockCounter += 1
	return &retVal
}

func (b *PrintBlock) GetNext(args ...any) *Block {
	return b.next
}

func (b *PrintBlock) SetNext(next Block) {
	b.next = &next
}
