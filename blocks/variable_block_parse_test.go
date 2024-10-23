package blocks_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
)

func TestVariableParse_1(t *testing.T) { // Valid
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"x = 3.0",
		"y = 4.3",
		"z = 5",
	}
	if err := varBlock.Parse(content); err != nil {
		t.Errorf("Error: %s", err)
	}
	targetMap := map[string]any{
		"x": 3.0,
		"y": 4.3,
		"z": 5.0,
	}

	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_2(t *testing.T) { // Valid
	varBlock := blocks.NewVariableBlock()
	content := []string{}
	if err := varBlock.Parse(content); err != nil {
		t.Errorf("Error: %s", err)
	}
	targetMap := map[string]any{}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_3(t *testing.T) { // Tokens error
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"x = -9",
		"s = ",
	}
	if err := varBlock.Parse(content); err == nil {
		t.Errorf("Error should occur")
	}
	targetMap := map[string]any{
		"x": -9.0,
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_4(t *testing.T) { // Number-key error
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = 88.9",
		"3 = 4",
	}
	if err := varBlock.Parse(content); err == nil {
		t.Errorf("Error should occur")
	}
	targetMap := map[string]any{
		"z": 88.9,
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_5(t *testing.T) { // String-value error
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = 88.9",
		"k = x",
	}
	if err := varBlock.Parse(content); err == nil {
		t.Errorf("Error should occur")
	}
	targetMap := map[string]any{
		"z": 88.9,
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_6(t *testing.T) { // Token error
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = 88.9",
		"k = 4 5",
	}
	if err := varBlock.Parse(content); err == nil {
		t.Errorf("Error should occur")
	}
	targetMap := map[string]any{
		"z": 88.9,
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_7(t *testing.T) { // Token error
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = 88.9",
		"k ? 4",
	}
	if err := varBlock.Parse(content); err == nil {
		t.Errorf("Error should occur")
	}
	targetMap := map[string]any{
		"z": 88.9,
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}