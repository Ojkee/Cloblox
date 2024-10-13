package graph_test

import (
	"testing"

	"Cloblox/blocks"
	"Cloblox/graph"
)

func TestValid_1(t *testing.T) {
	b1 := blocks.NewStartBlock(0)
	b2 := blocks.NewPrintBlock(1)
	b3 := blocks.NewStopBlock(2)

	b1.SetNext(b2)
	b2.SetNext(b3)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3})

	if diagram.IsFullyConnected() == false {
		t.Fail()
	}
}

func TestValid_2(t *testing.T) {
	b1 := blocks.NewStartBlock(0)
	b2 := blocks.NewVariableBlock[float32](1)
	b3 := blocks.NewIfBlock[int](2)
	b4 := blocks.NewStopBlock(3)

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
	b1 := blocks.NewStartBlock(0)
	b2 := blocks.NewVariableBlock[int](1)
	b3 := blocks.NewStopBlock(2)

	b1.SetNext(b2)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3})

	if diagram.IsFullyConnected() == true {
		t.Fail()
	}
}

func TestValid_4(t *testing.T) {
	b1 := blocks.NewPrintBlock(0)
	b2 := blocks.NewVariableBlock[int](1)
	b3 := blocks.NewStopBlock(2)
	b4 := blocks.NewStartBlock(3)

	b1.SetNext(b2)
	b2.SetNext(b3)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3, b4})

	if diagram.IsFullyConnected() == true {
		t.Fail()
	}
}

func TestValid_5(t *testing.T) {
	b1 := blocks.NewStartBlock(0)
	b2 := blocks.NewVariableBlock[int](1)
	b3 := blocks.NewIfBlock[int](2)
	b4 := blocks.NewVariableBlock[int](3)
	b5 := blocks.NewStopBlock(4)

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
	b1 := blocks.NewStartBlock(0)
	b2 := blocks.NewVariableBlock[int](1)
	b3 := blocks.NewIfBlock[int](2)
	b4 := blocks.NewVariableBlock[int](3)

	b1.SetNext(b2)
	b2.SetNext(b3)
	b3.SetNextFalse(b4)
	b4.SetNext(b3)

	diagram := graph.NewGraph(&[]blocks.Block{b1, b2, b3, b4})

	if diagram.IsFullyConnected() == true {
		t.Fail()
	}
}
