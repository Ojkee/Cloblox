package blocks_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
)

func TestActionBlock_MathOperations_Err_1(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("++")
	if err == nil {
		t.Errorf("Should Error")
	}
}

func TestActionBlock_MathOperations_Err_2(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x 2 * 2")
	if err == nil {
		t.Errorf("Should Error")
	}
}

func TestActionBlock_MathOperations_Err_3(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x /= 0")
	if err != nil {
		t.Errorf("Error: %s", err)
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
	_, _, err = actionBlock.PerformGetUpdateKVP()
	if err == nil {
		t.Errorf("Should Error")
	}
}

func TestActionBlock_MathOperations_Err_4(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x = 2/0")
	if err != nil {
		t.Errorf("Error: %s", err)
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
	_, _, err = actionBlock.PerformGetUpdateKVP()
	if err == nil {
		t.Errorf("Should Error")
	}
}

func TestActionBlock_MathOperations_Err_5(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x = ")
	if err == nil {
		t.Errorf("Should Err")
	}
}

func TestActionBlock_MathOperations_Err_6(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput(" *= 2")
	if err == nil {
		t.Errorf("Should Err")
	}
}

func TestActionBlock_MathOperations_Err_7(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput(" *= ")
	if err == nil {
		t.Errorf("Should Err")
	}
}
