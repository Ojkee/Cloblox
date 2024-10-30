package blocks_test

import (
	"reflect"
	"testing"

	"Cloblox/blocks"
)

func TestActionBlock_MathOperations_1(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x++")
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
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
	targetKvp := map[string]float64{
		"x": 6.9,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_2(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x--")
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
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
	targetKvp := map[string]float64{
		"x": 4.9,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_3(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x = 1 / 2")
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
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
	targetKvp := map[string]float64{
		"x": 0.5,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_4(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x += 1/2")
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
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
	targetKvp := map[string]float64{
		"x": 6.4,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_5(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x *= 1/2")
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
	newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
	if err != nil {
		t.Errorf("Error: '%s'", err.Error())
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
	targetKvp := map[string]float64{
		"x": 2.95,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_6(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x = y * s")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"x": 5.9,
		"y": 2.5,
		"s": -4.2,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"x", "y", "s"}
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
		"x": -10.5,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_7(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("x -= y * s")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"x": 5.9,
		"y": 2.5,
		"s": -4.2,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"x", "y", "s"}
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
		"x": 16.4,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_8(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("s[i_2] -= s[x] * i")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s[i_2]": 5.9,
		"s[x]":   2.5,
		"i":      -4.2,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s[i_2]", "s[x]", "i"}
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
		"s[i_2]": 16.4,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_9(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("s[i_2] = (s[x] - i)^2")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s[i_2]": 5.9,
		"s[x]":   2,
		"i":      5,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s[i_2]", "s[x]", "i"}
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
		"s[i_2]": 9.0,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_10(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("s[i_2] = (s[x] - 6)^2")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s[i_2]": 5.9,
		"s[x]":   2,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s[i_2]", "s[x]"}
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
		"s[i_2]": 16.0,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}

func TestActionBlock_MathOperations_11(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	err := actionBlock.ParseFromUserInput("s[i_2]--")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s[i_2]": 5.9,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s[i_2]"}
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
		"s[i_2]": 4.9,
	}
	if !reflect.DeepEqual(targetKvp, newKvp) {
		t.Errorf("Error KVP:\n\t%v\n\t%v", targetKvp, newKvp)
	}
}
