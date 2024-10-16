package blocks

type BlockType interface {
	float32 | []float32
}

type Block interface {
	GetId() int
	SetId(id int)
	GetName() string
	GetNext(args ...float32) *Block
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
