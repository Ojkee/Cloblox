package blocks_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
)

// Valid 1-2
// Error 3-7
// Valid list 8-10
// Error list 11-15

func TestVariableParse_1(t *testing.T) {
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

func TestVariableParse_2(t *testing.T) {
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

func TestVariableParse_8(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"x = [88.9, 3.2, -2, -4.5, 8]",
	}
	err := varBlock.Parse(content)
	if err != nil {
		t.Errorf("Error:\n\t%s", err.Error())
	}
	targetMap := map[string]any{
		"x": []float64{88.9, 3.2, -2, -4.5, 8},
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_9(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"x = [88.9, 3.2, -2, -4.5, 8]",
		"z = [88.9,-2, -4.5, 8]",
	}
	err := varBlock.Parse(content)
	if err != nil {
		t.Errorf("Error:\n\t%s", err.Error())
	}
	targetMap := map[string]any{
		"x": []float64{88.9, 3.2, -2, -4.5, 8},
		"z": []float64{88.9, -2, -4.5, 8},
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_10(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"x = 5.5",
		"z = [88.9,-2, -4.5, 8]",
		"s_2 = -4.5",
	}
	err := varBlock.Parse(content)
	if err != nil {
		t.Errorf("Error:\n\t%s", err.Error())
	}
	targetMap := map[string]any{
		"x":   5.5,
		"z":   []float64{88.9, -2, -4.5, 8},
		"s_2": -4.5,
	}
	if !reflect.DeepEqual(targetMap, varBlock.GetVars()) {
		t.Errorf("%v  !=  %v ", targetMap, varBlock.GetVars())
	}
}

func TestVariableParse_11(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z =88.9,-2, -4.5, 8",
	}
	err := varBlock.Parse(content)
	if err == nil {
		t.Errorf("Error should occur")
	}
}

func TestVariableParse_12(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = [88.9,-2, -4.5, 8",
	}
	err := varBlock.Parse(content)
	if err == nil {
		t.Errorf("Error should occur")
	}
}

func TestVariableParse_13(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = 88.9,-2, -4.5, 8]",
	}
	err := varBlock.Parse(content)
	if err == nil {
		t.Errorf("Error should occur")
	}
}

func TestVariableParse_14(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = [x, 88.9,-2, -4.5, 8]",
	}
	err := varBlock.Parse(content)
	if err == nil {
		t.Errorf("Error should occur")
	}
}

func TestVariableParse_15(t *testing.T) {
	varBlock := blocks.NewVariableBlock()
	content := []string{
		"z = [, 88.9,-2, -4.5, 8]",
	}
	err := varBlock.Parse(content)
	if err == nil {
		t.Errorf("Error should occur")
	}
}
