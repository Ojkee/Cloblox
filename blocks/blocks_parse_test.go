package blocks_test

import (
	"testing"

	"Cloblox/blocks"
)

func TestIfParse_1(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.ParseCondition("x < 4")
	lhs, rhs := ifBlock.GetKeys()
	if lhs != "x" || rhs != "4" {
		t.Errorf("lhs: %s != %s\nrhs: %s != %s", lhs, "x", rhs, "4")
	}
	if ifBlock.Compare(3, 4) == false {
		t.Fail()
	}
	if ifBlock.Compare(4, 4) == true {
		t.Fail()
	}
	if ifBlock.Compare(5, 4) == true {
		t.Fail()
	}
}

func TestIfParse_2(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.ParseCondition("s[i] == i")
	lhs, rhs := ifBlock.GetKeys()
	if lhs != "s[i]" || rhs != "i" {
		t.Errorf("lhs: %s != %s\nrhs: %s != %s", lhs, "x", rhs, "4")
	}
	if ifBlock.Compare(3, 4) == true {
		t.Fail()
	}
	if ifBlock.Compare(4, 4) == false {
		t.Fail()
	}
	if ifBlock.Compare(5, 4) == true {
		t.Fail()
	}
}

func TestIfParse_3(t *testing.T) {
	ifBlock := blocks.NewIfBlock()
	ifBlock.ParseCondition("s[4] != 7")
	lhs, rhs := ifBlock.GetKeys()
	if lhs != "s[4]" || rhs != "7" {
		t.Errorf("lhs: %s != %s\nrhs: %s != %s", lhs, "x", rhs, "4")
	}
	if ifBlock.Compare(3, 4) == false {
		t.Fail()
	}
	if ifBlock.Compare(4, 4) == true {
		t.Fail()
	}
	if ifBlock.Compare(5, 4) == false {
		t.Fail()
	}
}
