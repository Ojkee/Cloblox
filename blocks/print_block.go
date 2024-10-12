package blocks

type PrintBlock struct {
	BlockDefault
	next *Block
}

func NewPrintBlock(id int) *PrintBlock {
	return &PrintBlock{
		BlockDefault: BlockDefault{
			id:      id,
			name:    "print",
			content: "print",
		},
		next: nil,
	}
}

func (b *PrintBlock) GetContent() string {
	return b.content
}

func (b *PrintBlock) ClearNext() {
	b.next = nil
}

func (b *PrintBlock) GetNext(args ...any) *Block {
	return nil
}
