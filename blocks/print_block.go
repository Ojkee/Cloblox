package blocks

type PrintBlock struct {
	BlockDefault
	next *Block
}

func NewPrintBlock() *PrintBlock {
	return &PrintBlock{
		BlockDefault: BlockDefault{
			id:   -1,
			name: "Print Block",
		},
		next: nil,
	}
}

func (b *PrintBlock) GetNext(args ...float32) *Block {
	return b.next
}

func (b *PrintBlock) SetNext(next Block) {
	b.next = &next
}
