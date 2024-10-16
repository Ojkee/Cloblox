package blocks

type BlockType interface {
	int | float32 | string
}

type Block interface {
	GetId() int
	SetId(id int)
	GetName() string
	GetContent() string
	GetNext(args ...any) *Block
}

type SingleOutBlock interface {
	Block
	SetNext(next Block)
}

type ManyOutBlock interface {
	Block
	GetNextTrue() *Block
	GetNextFalse() *Block
	SetNextTrue(next Block)
	SetNextFalse(next Block)
}

type BlockDefault struct {
	id      int
	name    string
	content string
}

func (b *BlockDefault) GetId() int {
	return b.id
}

func (b *BlockDefault) SetId(id int) {
	b.id = id
}

func (b *BlockDefault) GetName() string {
	return b.name
}

func (b *BlockDefault) GetContent() string {
	return b.content
}
