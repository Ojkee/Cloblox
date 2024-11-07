package graph_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
	"Cloblox/graph"
)

// BUBBLE SORT
func TestGraphSimulation_1(t *testing.T) {
	start := blocks.NewStartBlock()
	variableT := blocks.NewVariableBlock()
	variableT.Parse([]string{"t = [0, -1, 4, 3, 1, 0, 5]"})
	variableIJ := blocks.NewVariableBlock()
	variableIJ.Parse([]string{"i = 0", "j = 0"})
	actionJ := blocks.NewActionBlock()
	actionJ.ParseFromUserInput("j = i")
	ifI := blocks.NewIfBlock()
	ifI.SetConditionExpr("i < 7")
	ifJ := blocks.NewIfBlock()
	ifJ.SetConditionExpr("j < 6")
	ifSwap := blocks.NewIfBlock()
	ifSwap.SetConditionExpr("t[j] > t[j+1]")
	actionSwap := blocks.NewActionBlock()
	actionSwap.ParseFromUserInput("swap t[j], t[j+1]")
	actionJInc := blocks.NewActionBlock()
	actionJInc.ParseFromUserInput("j = j + 1")
	actionIInc := blocks.NewActionBlock()
	actionIInc.ParseFromUserInput("i++")
	stop := blocks.NewStopBlock()

	diagram := graph.NewGraph(&[]blocks.Block{
		start, variableT, variableIJ, actionJ, ifI, ifJ, ifSwap, actionSwap, actionJInc, actionIInc, stop,
	}, 200)
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	diagram.ConnectByIds(2, 3)
	diagram.ConnectByIds(3, 4)
	diagram.ConnectByIds(4, 10, false)
	diagram.ConnectByIds(4, 5, true)
	diagram.ConnectByIds(5, 9, false)
	diagram.ConnectByIds(9, 3)
	diagram.ConnectByIds(5, 6, true)
	diagram.ConnectByIds(6, 7, true)
	diagram.ConnectByIds(7, 8)
	diagram.ConnectByIds(6, 8, false)
	diagram.ConnectByIds(8, 5)

	err := diagram.InitIfValid()
	if err != nil {
		t.Fatal(err)
	}

	finished, _, err := diagram.MakeStep()
	for !finished && err == nil {
		finished, _, err = diagram.MakeStep()
	}
	kvp := diagram.GetAllVars()
	result, ok := kvp["t"]
	if !ok {
		t.Fail()
	}
	target := []float64{-1, 0, 0, 1, 3, 4, 5}
	if !reflect.DeepEqual(target, result) {
		t.Errorf("Error:\n\t%v\n\t%v", target, result)
	}
}

// INSERTION SORT
func TestGraphSimulation_2(t *testing.T) {
	start := blocks.NewStartBlock()
	variableT := blocks.NewVariableBlock()
	variableT.Parse([]string{"t = [0, -1, 4, 3, 1, 0, 5]"})
	variableINJKey := blocks.NewVariableBlock()
	variableINJKey.Parse([]string{"i = 1", "n = 7", "j = 0", "key = 0"})
	ifI := blocks.NewIfBlock()
	ifI.SetConditionExpr("i < n")
	stop := blocks.NewStopBlock()
	actionKey := blocks.NewActionBlock()
	actionKey.ParseFromUserInput("key = t[i]")
	actionJI := blocks.NewActionBlock()
	actionJI.ParseFromUserInput("j = i - 1")
	ifJ := blocks.NewIfBlock()
	ifJ.SetConditionExpr("j >= 0")
	ifT := blocks.NewIfBlock()
	ifT.SetConditionExpr("t[j] > key")
	actionT := blocks.NewActionBlock()
	actionT.ParseFromUserInput("t[j+1] = key")
	actionI := blocks.NewActionBlock()
	actionI.ParseFromUserInput("i++")
	actionTTrue := blocks.NewActionBlock()
	actionTTrue.ParseFromUserInput("t[j+1] = t[j]")
	actionJ := blocks.NewActionBlock()
	actionJ.ParseFromUserInput("j--")

	diagram := graph.NewGraph(&[]blocks.Block{
		start, variableT, variableINJKey, ifI, stop, actionKey, actionJI, ifJ, ifT, actionT, actionI, actionTTrue, actionJ,
	}, 1000)
	diagram.ConnectByIds(0, 1)
	diagram.ConnectByIds(1, 2)
	diagram.ConnectByIds(2, 3)
	diagram.ConnectByIds(3, 4, false)
	diagram.ConnectByIds(3, 5, true)
	diagram.ConnectByIds(5, 6)
	diagram.ConnectByIds(6, 7)
	diagram.ConnectByIds(7, 9, false)
	diagram.ConnectByIds(9, 10)
	diagram.ConnectByIds(10, 3)
	diagram.ConnectByIds(7, 8, true)
	diagram.ConnectByIds(8, 9, false)
	diagram.ConnectByIds(8, 11, true)
	diagram.ConnectByIds(11, 12)
	diagram.ConnectByIds(12, 7)

	err := diagram.InitIfValid()
	if err != nil {
		t.Fatal(err)
	}

	finished, _, err := diagram.MakeStep()
	for !finished && err == nil {
		finished, _, err = diagram.MakeStep()
	}
	kvp := diagram.GetAllVars()
	result, ok := kvp["t"]
	if !ok {
		t.Fail()
	}
	target := []float64{-1, 0, 0, 1, 3, 4, 5}
	if !reflect.DeepEqual(target, result) {
		t.Errorf("Error:\n\t%v\n\t%v", target, result)
	}
}
