package graph_test

import (
	"testing"

	"Cloblox/blocks"
	"Cloblox/graph"
)

func TestValid_1(t *testing.T) {
	b1 := blocks.NewStartBlock()
	b2 := blocks.NewPrintBlock()
	b3 := blocks.NewStopBlock()

	b1.SetNext(b2)
	b2.SetNext(b3)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3})

	if diagram.IsFullyConnected() == false {
		t.Fail()
	}
}

func TestValid_2(t *testing.T) {
	b1 := blocks.NewStartBlock()
	b2 := blocks.NewVariableBlock[float32]()
	b3 := blocks.NewIfBlock[int]()
	b4 := blocks.NewStopBlock()

	b1.SetNext(b2)
	b2.SetNext(b3)
	b3.SetNextTrue(b4)
	b3.SetNextFalse(b4)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3, b4})

	if diagram.IsFullyConnected() == false {
		t.Fail()
	}
}

func TestValid_3(t *testing.T) {
	b1 := blocks.NewStartBlock()
	b2 := blocks.NewVariableBlock[int]()
	b3 := blocks.NewStopBlock()

	b1.SetNext(b2)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3})

	if diagram.IsFullyConnected() == true {
		t.Fail()
	}
}

func TestValid_4(t *testing.T) {
	b1 := blocks.NewPrintBlock()
	b2 := blocks.NewVariableBlock[int]()
	b3 := blocks.NewStopBlock()
	b4 := blocks.NewStartBlock()

	b1.SetNext(b2)
	b2.SetNext(b3)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3, b4})

	if diagram.IsFullyConnected() == true {
		t.Fail()
	}
}

func TestValid_5(t *testing.T) {
	b1 := blocks.NewStartBlock()
	b2 := blocks.NewVariableBlock[int]()
	b3 := blocks.NewIfBlock[int]()
	b4 := blocks.NewVariableBlock[int]()
	b5 := blocks.NewStopBlock()

	b1.SetNext(b2)
	b2.SetNext(b3)
	b3.SetNextTrue(b5)
	b3.SetNextFalse(b4)
	b4.SetNext(b3)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3, b4})

	if diagram.IsFullyConnected() == false {
		t.Fail()
	}
}

func TestValid_6(t *testing.T) {
	b1 := blocks.NewStartBlock()
	b2 := blocks.NewVariableBlock[int]()
	b3 := blocks.NewIfBlock[int]()
	b4 := blocks.NewVariableBlock[int]()

	b1.SetNext(b2)
	b2.SetNext(b3)
	b3.SetNextFalse(b4)
	b4.SetNext(b3)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3, b4})

	if diagram.IsFullyConnected() == true {
		t.Fail()
	}
}
