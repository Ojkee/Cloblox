package blocks_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
)

func TestActionBlock_Print_1(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("print x")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	kvp := map[string]float64{
		"x": 5.9,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"x"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error keys:\n\t%v\n\t%v", targetKeys, keys)
	}
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}

	targetMess := "x = 5.900"
	if reflect.DeepEqual(targetMess, mess) {
		t.Errorf("Error mess:\n\t%s\n\t%s", targetMess, mess)
	}
	if newKvp != nil {
		t.Error("KVP should be nil")
	}
}

func TestActionBlock_Print_2(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("print ")
	if err == nil {
		t.Errorf("Should Error")
	}
}

func TestActionBlock_Print_3(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x print ")
	if err == nil {
		t.Errorf("Should Error")
	}
}
