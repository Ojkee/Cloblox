package graph_test

import (
	"reflect"
	"testing"

	"Cloblox/graph"
)

// tab[]{ 1, 4, 2, 2, 4, 3, 5, 1, 0, 1 }
// tab[tab[tab[3]]] TBD

func TestGetArrayValue_1(t *testing.T) {
	diagram := graph.NewGraph(nil)
	allVars := map[string]any{
		"x":   3.0,
		"tab": []float64{3.0, 4.0, 1.0, 2.0, 3.0},
	}
	diagram.SetAllVars(allVars)
	kvp, err := diagram.GetKvpByKeys(&[]string{"x", "tab[3]"})
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	targetKvp := map[string]float64{
		"x":      3.0,
		"tab[3]": 2.0,
	}
	if !reflect.DeepEqual(targetKvp, kvp) {
		t.Errorf("Fail 1:\n\t%v\n\t%v", targetKvp, kvp)
	}
}

func TestGetArrayValue_2(t *testing.T) {
	diagram := graph.NewGraph(nil)
	allVars := map[string]any{
		"x":   3.0,
		"tab": []float64{3.0, 4.0, 1.0, 2.0, 3.0},
		"y":   0.2,
	}
	diagram.SetAllVars(allVars)
	kvp, err := diagram.GetKvpByKeys(&[]string{"tab[2 * x - 2]"})
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	targetKvp := map[string]float64{
		"tab[2 * x - 2]": 3.0,
	}
	if !reflect.DeepEqual(targetKvp, kvp) {
		t.Errorf("Fail 1:\n\t%v\n\t%v", targetKvp, kvp)
	}
}

func TestGetArrayValue_3(t *testing.T) {
	diagram := graph.NewGraph(nil)
	allVars := map[string]any{
		"x":   3.0,
		"tab": []float64{3.0, 4.0, 1.0, 2.0, 3.0},
		"y":   0.2,
	}
	diagram.SetAllVars(allVars)
	kvp, err := diagram.GetKvpByKeys(&[]string{"tab[2 * x * 2]"})
	if err == nil {
		t.Errorf("Should be error")
	}
	if len(kvp) != 0 {
		t.Errorf("Fail 1:\n\t%v", kvp)
	}
}

func TestGetArrayValue_4(t *testing.T) {
	diagram := graph.NewGraph(nil)
	allVars := map[string]any{
		"x":   3.0,
		"tab": []float64{3.0, 4.0, 1.0, 2.0, 3.0},
		"y":   0.2,
	}
	diagram.SetAllVars(allVars)
	kvp, err := diagram.GetKvpByKeys(&[]string{"tab[-2]"})
	if err == nil {
		t.Errorf("Should be error")
	}
	if len(kvp) != 0 {
		t.Errorf("Fail 1:\n\t%v", kvp)
	}
}
