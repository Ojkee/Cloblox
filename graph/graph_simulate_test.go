package graph_test

import (
	"fmt"
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

	start.SetNext(variableT)
	variableT.SetNext(variableINJKey)
	variableINJKey.SetNext(ifI)
	ifI.SetNextFalse(stop)
	ifI.SetNextTrue(actionKey)
	actionKey.SetNext(actionJI)
	actionJI.SetNext(ifJ)
	ifJ.SetNextFalse(actionT)
	actionT.SetNext(actionI)
	actionI.SetNext(ifI)
	ifJ.SetNextTrue(ifT)
	ifT.SetNextFalse(actionT)
	ifT.SetNextTrue(actionTTrue)
	actionTTrue.SetNext(actionJ)
	actionJ.SetNext(ifJ)

	diagram := graph.NewGraph(&[]blocks.Block{
		start, variableT, variableINJKey, ifI, stop, actionKey, actionJI, ifJ, ifT, actionT, actionI, actionTTrue, actionJ,
	}, 1000)
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
