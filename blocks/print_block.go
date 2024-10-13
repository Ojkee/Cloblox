package blocks

type PrintBlock struct {
	BlockDefault
	next *Block
}

func NewPrintBlock(id int) *PrintBlock {
	return &PrintBlock{
		BlockDefault: BlockDefault{
			id:      id,
			name:    "print block",
			content: "print",
		},
		next: nil,
	}
}

func (b *PrintBlock) GetNext(args ...any) *Block {
	return b.next
}

func (b *PrintBlock) SetNext(next Block) {
	b.next = &next
}
