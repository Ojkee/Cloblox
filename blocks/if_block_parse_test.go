package blocks_test

import (
	"testing"

	"Cloblox/blocks"
)

func TestIfParse_1(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("x == 3")
	keys := ifBlock.GetKeys()
	if len(keys) != 1 {
		t.Errorf("Error: '%v'", keys)
	}
	kvp := map[string]float32{
		"x": 3,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
	kvp["x"] = 7
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
}

func TestIfParse_2(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("x + m < z/2")
	keys := ifBlock.GetKeys()
	if len(keys) != 3 {
		t.Errorf("Error: '%v'", keys)
	}
	kvp := map[string]float32{
		"x": 3,
		"m": 2,
		"z": 11,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
	kvp["x"] = 10
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
}

func TestIfParse_3(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("si*2 != 8")
	keys := ifBlock.GetKeys()
	if len(keys) != 1 {
		t.Errorf("Error: '%v'", keys)
	}
	kvp := map[string]float32{
		"si": 4,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
	kvp["si"] = 3
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
}

func TestIfParse_4(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("s[i]*2 != 8")
	keys := ifBlock.GetKeys()
	if len(keys) != 1 {
		t.Errorf("Error: '%v'", keys)
	}
	kvp := map[string]float32{
		"s[i]": 4,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
	kvp["s[i]"] = 3
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
}
