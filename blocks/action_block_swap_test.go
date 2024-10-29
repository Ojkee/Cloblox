package blocks_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
)

func TestActionBlock_Swap_1(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("swap x, y")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"x": 5.9,
		"y": 4.2,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"x", "y"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error keys:\n\t%v\n\t%v", targetKeys, keys)
	}
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
	targetKvp := map[string]float64{
		"x": 4.2,
		"y": 5.9,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_Swap_2(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("swap x")
	if err == nil {
		t.Errorf("Should error")
	}
}

func TestActionBlock_Swap_3(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("swap x[i], y[4+5*x]")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	keys := actionBlock.GetKeys()
	targetKeys := []string{"x[i]", "y[4+5*x]"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error keys:\n\t%v\n\t%v", targetKeys, keys)
	}
	kvp := map[string]float64{
		"x[i]":     3.2,
		"y[4+5*x]": 5.1,
	}
	actionBlock.SetActionKVP(&kvp)
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
	targetKvp := map[string]float64{
		"x[i]":     5.1,
		"y[4+5*x]": 3.2,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_Swap_4(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("swap x, y, o")
	if err == nil {
		t.Errorf("Should error")
	}
}

func TestActionBlock_Swap_5(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x, y swap")
	if err == nil {
		t.Errorf("Should error")
	}
}
