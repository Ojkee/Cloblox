package blocks

type BlockType interface {
	float32 | []float32
}

type Block interface {
	GetId() int
	SetId(id int)
	GetName() string
	GetNext() (*Block, error)
}

type BlockSingleOut interface {
	Block
	SetNext(next Block)
}

type BlockManyOut interface {
	Block
	GetNextTrue() *Block
	GetNextFalse() *Block
	SetNextTrue(next Block)
	SetNextFalse(next Block)
}

type BlockWithVars interface {
	GetVars() map[string]any
}

type BlockDefault struct {
	id   int
	name string
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
