package graph_test

import (
	"testing"

	"Cloblox/blocks"
	"Cloblox/graph"
)

func TestInitIfValid_1(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(stopBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, stopBlock,
	})
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_2(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(ifBlock)
	ifBlock.SetNextTrue(stopBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, ifBlock, stopBlock,
	})
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}

func TestInitIfValid_3(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(ifBlock)
	ifBlock.SetNextTrue(stopBlock)
	ifBlock.SetNextFalse(stopBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, ifBlock, stopBlock,
	})
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_4(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextTrue(stopBlock)
	ifBlock.SetNextFalse(varBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, stopBlock,
	})
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_5(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextTrue(varBlock)
	ifBlock.SetNextFalse(varBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, stopBlock,
	})
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}

func TestInitIfValid_6(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	ifBlock2 := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextFalse(ifBlock2)
	ifBlock.SetNextTrue(varBlock)
	ifBlock2.SetNextTrue(varBlock)
	ifBlock2.SetNextFalse(stopBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, ifBlock2, stopBlock,
	})
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_7(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	ifBlock2 := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextFalse(ifBlock2)
	ifBlock.SetNextTrue(varBlock)
	ifBlock2.SetNextFalse(stopBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, ifBlock2, stopBlock,
	})
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}

func TestInitIfValid_8(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	varBlock2 := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	ifBlock2 := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextFalse(ifBlock2)
	ifBlock.SetNextTrue(varBlock)
	ifBlock2.SetNextFalse(stopBlock)
	ifBlock2.SetNextTrue(varBlock2)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, ifBlock2, stopBlock,
	})
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}

func TestInitIfValid_9(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	varBlock2 := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	ifBlock2 := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextFalse(ifBlock2)
	ifBlock.SetNextTrue(varBlock)
	ifBlock2.SetNextFalse(stopBlock)
	ifBlock2.SetNextTrue(varBlock2)
	varBlock2.SetNext(stopBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_10(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	varBlock2 := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	ifBlock2 := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	varBlock.SetNext(ifBlock)
	ifBlock.SetNextFalse(ifBlock2)
	ifBlock.SetNextTrue(varBlock)
	ifBlock2.SetNextFalse(stopBlock)
	ifBlock2.SetNextTrue(varBlock2)
	varBlock2.SetNext(stopBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}

func TestInitIfValid_11(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	varBlock2 := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	ifBlock2 := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextFalse(ifBlock2)
	ifBlock.SetNextTrue(varBlock)
	ifBlock2.SetNextFalse(stopBlock)
	ifBlock2.SetNextTrue(varBlock2)
	varBlock2.SetNext(varBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_12(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	varBlock2 := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	ifBlock2 := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	startBlock.SetNext(varBlock)
	varBlock.SetNext(ifBlock)
	ifBlock.SetNextFalse(ifBlock2)
	ifBlock.SetNextTrue(varBlock)
	ifBlock2.SetNextFalse(varBlock)
	ifBlock2.SetNextTrue(varBlock2)
	varBlock2.SetNext(varBlock)

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}
