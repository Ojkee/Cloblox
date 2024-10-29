package blocks_test

import (
	"fmt"
	"reflect"
	"testing"

	"Cloblox/blocks"
)

// VALID 1-5
// ERRORS 6-8

func TestActionBlock_Rand_1(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	var mmin float64 = -10.0
	var mmax float64 = 10.0
	err := actionBlock.ParseFromUserInput(fmt.Sprintf("x = rand %f %f", mmin, mmax))
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
	for range 100 { // 100 simulations if vals is in range
		newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
		if err != nil {
			t.Errorf("Error: '%s'", err.Error())
			break
		}
		if mess != "" {
			t.Errorf("Error mess: '%s'", mess)
			break
		}
		if x, ok := newKvp["x"]; !ok || !(x >= mmin && x <= mmax) {
			t.Errorf("Error KVP:\n\t%v", newKvp)
			break
		}
	}
}

func TestActionBlock_Rand_2(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	mmin := 0
	mmax := 10
	err := actionBlock.ParseFromUserInput(fmt.Sprintf("s[c+i*2] = rand %d %d", mmin, mmax))
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s[c+i*2]": -2.0,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s[c+i*2]"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error keys:\n\t%v\n\t%v", targetKeys, keys)
	}
	for range 100 { // 100 simulations if vals is in range
		newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
		if err != nil {
			t.Errorf("Error: '%s'", err.Error())
			break
		}
		if mess != "" {
			t.Errorf("Error mess: '%s'", mess)
			break
		}
		if x, ok := newKvp["s[c+i*2]"]; !ok || !(x >= float64(mmin) && x <= float64(mmax)) {
			t.Errorf("Error KVP:\n\t%v", newKvp)
			break
		}
	}
}

func TestActionBlock_Rand_3(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	mmin := 7.0
	mmax := 10.0
	err := actionBlock.ParseFromUserInput("s = rand x y")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s": -2.0,
		"x": mmin,
		"y": mmax,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s", "x", "y"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error keys:\n\t%v\n\t%v", targetKeys, keys)
	}
	for range 100 { // 100 simulations if vals is in range
		newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
		if err != nil {
			t.Errorf("Error: '%s'", err.Error())
			break
		}
		if mess != "" {
			t.Errorf("Error mess: '%s'", mess)
			break
		}
		if x, ok := newKvp["s"]; !ok || !(x >= float64(mmin) && x <= float64(mmax)) {
			t.Errorf("Error KVP:\n\t%v", newKvp)
			break
		}
		if !(newKvp["x"] == mmin && newKvp["y"] == mmax) {
			t.Errorf("Error changed range keys:\n\t%f, %f\n\t%v", mmin, mmax, newKvp)
		}
	}
}

func TestActionBlock_Rand_4(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	mmin := -1.0
	mmax := 1.0
	err := actionBlock.ParseFromUserInput(fmt.Sprintf("s = rand %f y", mmin))
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s": -2.0,
		"y": mmax,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s", "y"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error keys:\n\t%v\n\t%v", targetKeys, keys)
	}
	for range 100 { // 100 simulations if vals is in range
		newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
		if err != nil {
			t.Errorf("Error: '%s'", err.Error())
			break
		}
		if mess != "" {
			t.Errorf("Error mess: '%s'", mess)
			break
		}
		if x, ok := newKvp["s"]; !ok || !(x >= float64(mmin) && x <= float64(mmax)) {
			t.Errorf("Error KVP:\n\t%v", newKvp)
			break
		}
		if newKvp["y"] != mmax {
			t.Errorf("Error changed range keys:\n\t%f, %f\n\t%v", mmin, mmax, newKvp)
		}
	}
}

func TestActionBlock_Rand_5(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	mmin := -5.0
	mmax := 2.0
	err := actionBlock.ParseFromUserInput(fmt.Sprintf("s = rand x, %f", mmax))
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	kvp := map[string]float64{
		"s": 7.0,
		"x": mmin,
	}
	actionBlock.SetActionKVP(&kvp)
	keys := actionBlock.GetKeys()
	targetKeys := []string{"s", "x"}
	if !reflect.DeepEqual(targetKeys, keys) {
		t.Errorf("Error keys:\n\t%v\n\t%v", targetKeys, keys)
	}
	for range 100 { // 100 simulations if vals is in range
		newKvp, mess, err := actionBlock.PerformGetUpdateKVP()
		if err != nil {
			t.Errorf("Error: '%s'", err.Error())
			break
		}
		if mess != "" {
			t.Errorf("Error mess: '%s'", mess)
			break
		}
		if x, ok := newKvp["s"]; !ok || !(x >= float64(mmin) && x <= float64(mmax)) {
			t.Errorf("Error KVP:\n\t%v", newKvp)
			break
		}
		if newKvp["x"] != mmin {
			t.Errorf("Error changed range keys:\n\t%f, %f\n\t%v", mmin, mmax, newKvp)
		}
	}
}

func TestActionBlock_Rand_6(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	mmin := 5.0
	mmax := 2.0
	actionBlock.ParseFromUserInput(fmt.Sprintf("s = rand %f, %f", mmin, mmax))
	kvp := map[string]float64{
		"s": -2.0,
	}
	actionBlock.SetActionKVP(&kvp)
	_, mess, err := actionBlock.PerformGetUpdateKVP()
	if err == nil {
		t.Errorf("Should Error")
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
}

func TestActionBlock_Rand_7(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	mmin := 5.0
	actionBlock.ParseFromUserInput(fmt.Sprintf("s = rand %f", mmin))
	kvp := map[string]float64{
		"s": -2.0,
	}
	actionBlock.SetActionKVP(&kvp)
	_, mess, err := actionBlock.PerformGetUpdateKVP()
	if err == nil {
		t.Errorf("Should Error")
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
}

func TestActionBlock_Rand_8(t *testing.T) {
	actionBlock := blocks.NewActionBlock()
	mmin := 5.0
	actionBlock.ParseFromUserInput(fmt.Sprintf("rand %f 10", mmin))
	kvp := map[string]float64{}
	actionBlock.SetActionKVP(&kvp)
	_, mess, err := actionBlock.PerformGetUpdateKVP()
	if err == nil {
		t.Errorf("Should Error")
	}
	if mess != "" {
		t.Errorf("Error mess: '%s'", mess)
	}
}
