package graph_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
	"Cloblox/graph"
)

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

	start.SetNext(variableT)
	variableT.SetNext(variableIJ)
	variableIJ.SetNext(actionJ)
	actionJ.SetNext(ifI)
	ifI.SetNextFalse(stop)
	ifI.SetNextTrue(ifJ)
	ifJ.SetNextFalse(actionIInc)
	actionIInc.SetNext(actionJ)
	ifJ.SetNextTrue(ifSwap)
	ifSwap.SetNextTrue(actionSwap)
	actionSwap.SetNext(actionJInc)
	ifSwap.SetNextFalse(actionJInc)
	actionJInc.SetNext(ifJ)

	diagram := graph.NewGraph(&[]blocks.Block{
		start, variableT, variableIJ, actionJ, ifI, ifJ, ifSwap, actionSwap, actionJInc, actionIInc, stop,
	}, 200)
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
