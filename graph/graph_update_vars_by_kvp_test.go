package graph_test

import (
	"reflect"
	"testing"

	"Cloblox/graph"
)

func TestUpdateVarsByKVP_1(t *testing.T) {
	diagram := graph.NewGraph(nil)
	vars := map[string]any{
		"x": 4.0,
		"s": []float64{1, 2, 3, 5},
		"i": 1,
		"a": 32,
	}
	diagram.SetAllVars(vars)
	newVars := map[string]float64{
		"x":    2.0,
		"s[2]": 8.2,
	}
	err := diagram.UpdateVarsFromKVP(&newVars)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	result := diagram.GetAllVars()
	target := map[string]any{
		"x": 2.0,
		"s": []float64{1, 2, 8.2, 5},
		"i": 1,
		"a": 32,
	}
	if !reflect.DeepEqual(target, result) {
		t.Errorf("Error:\n\t%v\n\t%v", target, result)
	}
}

func TestUpdateVarsByKVP_2(t *testing.T) {
	diagram := graph.NewGraph(nil)
	vars := map[string]any{
		"x": 4.0,
		"s": []float64{1, 2, 3, 5},
		"i": 1,
		"a": 32,
	}
	diagram.SetAllVars(vars)
	newVars := map[string]float64{
		"y":    2.0,
		"s[2]": 8.2,
	}
	err := diagram.UpdateVarsFromKVP(&newVars)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	result := diagram.GetAllVars()
	target := map[string]any{
		"x": 4.0,
		"s": []float64{1, 2, 8.2, 5},
		"i": 1,
		"a": 32,
		"y": 2.0,
	}
	if !reflect.DeepEqual(target, result) {
		t.Errorf("Error:\n\t%v\n\t%v", target, result)
	}
}

func TestUpdateVarsByKVP_3(t *testing.T) {
	diagram := graph.NewGraph(nil)
	vars := map[string]any{
		"x": 4.0,
		"s": []float64{1, 2, 3, 5},
		"i": 0.0,
		"a": 32,
	}
	diagram.SetAllVars(vars)
	newVars := map[string]float64{
		"y":    2.0,
		"s[i]": 8.2,
	}
	err := diagram.UpdateVarsFromKVP(&newVars)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	result := diagram.GetAllVars()
	target := map[string]any{
		"x": 4.0,
		"s": []float64{8.2, 2, 3, 5},
		"i": 0.0,
		"a": 32,
		"y": 2.0,
	}
	if !reflect.DeepEqual(target, result) {
		t.Errorf("Error:\n\t%v\n\t%v", target, result)
	}
}

func TestUpdateVarsByKVP_4(t *testing.T) {
	diagram := graph.NewGraph(nil)
	vars := map[string]any{
		"x": 4.0,
		"s": []float64{1, 2, 3, 5},
		"i": 0.0,
		"a": 32,
	}
	diagram.SetAllVars(vars)
	newVars := map[string]float64{
		"y": 1.0,
	}
	err := diagram.UpdateVarsFromKVP(&newVars)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	newVars2 := map[string]float64{
		"s[y]": 8.2,
	}
	err = diagram.UpdateVarsFromKVP(&newVars2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	result := diagram.GetAllVars()
	target := map[string]any{
		"x": 4.0,
		"s": []float64{1, 8.2, 3, 5},
		"i": 0.0,
		"a": 32,
		"y": 1.0,
	}
	if !reflect.DeepEqual(target, result) {
		t.Errorf("Error:\n\t%v\n\t%v", target, result)
	}
}

func TestUpdateVarsByKVP_5(t *testing.T) {
	diagram := graph.NewGraph(nil)
	vars := map[string]any{
		"x": 1,
		"s": []float64{1, 3, 4},
	}
	diagram.SetAllVars(vars)
	newVars := map[string]float64{
		"s": 2,
	}
	err := diagram.UpdateVarsFromKVP(&newVars)
	if err == nil {
		t.Errorf("Should Error")
	}
}

func TestUpdateVarsByKVP_6(t *testing.T) {
	diagram := graph.NewGraph(nil)
	vars := map[string]any{
		"x": 1,
	}
	diagram.SetAllVars(vars)
	newVars := map[string]float64{
		"s[i]": 2,
	}
	err := diagram.UpdateVarsFromKVP(&newVars)
	if err == nil {
		t.Errorf("Should Error")
	}
}

func TestUpdateVarsByKVP_7(t *testing.T) {
	diagram := graph.NewGraph(nil)
	vars := map[string]any{
		"s": []float64{1, 2, 3},
	}
	diagram.SetAllVars(vars)
	newVars := map[string]float64{
		"s[5]": 2,
	}
	err := diagram.UpdateVarsFromKVP(&newVars)
	if err == nil {
		t.Errorf("Should Error")
	}
}
