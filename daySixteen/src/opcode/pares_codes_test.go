package opcode_test

import (
	"opcode"
	"reflect"
	"testing"
)

func TestFromSliceToRegisters(t *testing.T) {
	vals := []string{"0", "0", "2", "2"}
	result := opcode.FromSliceToRegisters(vals)
	if len(result) != 4 {
		t.Error("Length should have been 4")
	}
	if result[0] != 0 || result[3] != 2 {
		t.Error("Values were not the same")
	}
}

func TestFromSliceToOperation(t *testing.T) {
	vals := []string{"9", "5", "7", "7"}
	result := opcode.FromSliceToOperation(vals)
	if result.OpCode != 9 {
		t.Error("Should have been 9")
	}
	if result.First != 5 {
		t.Error("Should have been 5")
	}
	if result.Second != 7 {
		t.Error("Should have been 7")
	}
	if result.Output != 7 {
		t.Error("Should have been 7")
	}
}

func TestParse(t *testing.T) {
	val := opcode.Parse("../input1.txt")
	if !reflect.DeepEqual(val[0].Before, opcode.Registers{0, 0, 2, 2}) {
		t.Error("Before values were not equal")
	}

	if !reflect.DeepEqual(val[0].After, opcode.Registers{4, 0, 2, 2}) {
		t.Error("Before values were not equal")
	}

	if !reflect.DeepEqual(val[0].Operation, opcode.Operation{9, 2, 3, 0}) {
		t.Error("Before values were not equal")
	}
}
