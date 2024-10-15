package graph_test

import (
	"testing"

	"Cloblox/blocks"
	"Cloblox/graph"
)

func TestConnection_1(t *testing.T) {
	diagram := graph.NewGraph(&[]blocks.Block{
		blocks.NewStartBlock(),
		blocks.NewPrintBlock(),
		blocks.NewStopBlock(),
	})
	if err := diagram.ConnectByIds(0, 1); err != nil {
		t.Errorf("Shouldn't return error")
	}
	if err := diagram.ConnectByIds(1, 2); err != nil {
		t.Errorf("Shouldn't return error")
	}
	if !diagram.IsFullyConnected() {
		t.Fail()
	}
}

func TestConnection_2(t *testing.T) {
	diagram := graph.NewGraph(&[]blocks.Block{
		blocks.NewStartBlock(),
		blocks.NewPrintBlock(),
		blocks.NewStopBlock(),
	})
	if err := diagram.ConnectByIds(0, 3); err == nil {
		t.Errorf("Should return error")
	}
	if err := diagram.ConnectByIds(2, 2); err == nil {
		t.Errorf("Should return error")
	}
}
