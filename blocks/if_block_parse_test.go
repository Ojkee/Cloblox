package blocks_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
)

func TestIfParse_1(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("x == 3")
	keys := ifBlock.GetKeys()
	targetKeys := []string{"x"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
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
	targetKeys := []string{"x", "m", "z"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
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
	targetKeys := []string{"si"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
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
	targetKeys := []string{"s[i]"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
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

func TestIfParse_5(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("s[i]*2.0 < 8.5")
	keys := ifBlock.GetKeys()
	targetKeys := []string{"s[i]"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
		"s[i]": 4.2,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
	kvp["s[i]"] = 7.1
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
}

func TestIfParse_6(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("s[i]+x*2.0 == 6.2")
	keys := ifBlock.GetKeys()
	targetKeys := []string{"s[i]", "x"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
		"x":    1.0,
		"s[i]": 4.2,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
	kvp["s[i]"] = 7.1
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
}

func TestIfParse_7(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("s[i+x]*2.0 == 8.4")
	keys := ifBlock.GetKeys()
	targetKeys := []string{"s[i+x]"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
		"s[i+x]": 4.2,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
	kvp["s[i+x]"] = 7.1
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
}

func TestIfParse_8(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("s[i+x]+y*2.0 == 2.2")
	keys := ifBlock.GetKeys()
	targetKeys := []string{"s[i+x]", "y"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
		"s[i+x]": 4.2,
		"y":      -1,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
	kvp["s[i+x]"] = 7.1
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
}

func TestIfParse_9(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.SetConditionExpr("s[i+x]*h+y[x-2]*2 == 24")
	keys := ifBlock.GetKeys()
	targetKeys := []string{"s[i+x]", "h", "y[x-2]"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error:\n\t'%v'\n\t'%v'", targetKeys, keys)
	}
	kvp := map[string]float64{
		"s[i+x]": 4,
		"h":      7,
		"y[x-2]": -2,
	}
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if !isTrue {
		t.Fail()
	}
	kvp["s[i+x]"] = 7
	ifBlock.SetConditionKVP(&kvp)
	if isTrue, err := ifBlock.IsEvalTrue(); err != nil {
		t.Errorf("Err: %v", err)
	} else if isTrue {
		t.Fail()
	}
}
