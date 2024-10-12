package blocks

type BlockType interface {
	int | float32 | string
}

type Block interface {
	GetContent() string
	GetNext(args ...any) *Block
}

type BlockDefault struct {
	id      int
	name    string
	content string
}

func (b *BlockDefault) GetName() string {
	return b.name
}
