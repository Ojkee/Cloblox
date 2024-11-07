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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_2(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, ifBlock, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2, true)
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}

func TestInitIfValid_3(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, ifBlock, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2, true)
	diagram.ConnectByIds(1, 2, false)
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_4(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	diagram.ConnectByIds(2, 3, true)
	diagram.ConnectByIds(2, 3, false)
	if err := diagram.InitIfValid(); err != nil {
		t.Error(err.Error())
	}
}

func TestInitIfValid_5(t *testing.T) {
	startBlock := blocks.NewStartBlock()
	varBlock := blocks.NewVariableBlock()
	ifBlock := blocks.NewIfBlock()
	stopBlock := blocks.NewStopBlock()

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	diagram.ConnectByIds(2, 1, true)
	diagram.ConnectByIds(2, 1, false)
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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, ifBlock2, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	diagram.ConnectByIds(2, 3, false)
	diagram.ConnectByIds(2, 1, true)
	diagram.ConnectByIds(3, 1, true)
	diagram.ConnectByIds(3, 4, false)
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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, ifBlock2, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	diagram.ConnectByIds(2, 3, false)
	diagram.ConnectByIds(2, 1, true)
	diagram.ConnectByIds(3, 4, false)
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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, ifBlock, ifBlock2, stopBlock, varBlock2,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	diagram.ConnectByIds(2, 3, false)
	diagram.ConnectByIds(2, 1, true)
	diagram.ConnectByIds(3, 4, false)
	diagram.ConnectByIds(3, 5, true)
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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 3)
	diagram.ConnectByIds(3, 4, false)
	diagram.ConnectByIds(3, 1, true)
	diagram.ConnectByIds(4, 5, false)
	diagram.ConnectByIds(4, 2, true)
	diagram.ConnectByIds(2, 5)
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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	diagram.ConnectByIds(1, 3)
	diagram.ConnectByIds(3, 4, false)
	diagram.ConnectByIds(3, 1, true)
	diagram.ConnectByIds(4, 5, false)
	diagram.ConnectByIds(4, 2, true)
	diagram.ConnectByIds(2, 5)
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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 3)
	diagram.ConnectByIds(3, 4, false)
	diagram.ConnectByIds(3, 1, true)
	diagram.ConnectByIds(4, 5, false)
	diagram.ConnectByIds(4, 2, true)
	diagram.ConnectByIds(2, 1)
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

	diagram := graph.NewGraph(&[]blocks.Block{
		startBlock, varBlock, varBlock2, ifBlock, ifBlock2, stopBlock,
	})
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 3)
	diagram.ConnectByIds(3, 4, false)
	diagram.ConnectByIds(3, 1, true)
	diagram.ConnectByIds(4, 1, false)
	diagram.ConnectByIds(4, 2, true)
	diagram.ConnectByIds(2, 1)
	if err := diagram.InitIfValid(); err == nil {
		t.Error("Should call error")
	}
}
