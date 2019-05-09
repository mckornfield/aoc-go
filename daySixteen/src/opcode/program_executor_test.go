package opcode_test

import (
	"opcode"
	"testing"
)

func TestParseProgram(t *testing.T) {
	ops := opcode.ParseProgram("../puzz-input2.txt")
	if len(ops) != 966 {
		t.Error("Should have been length 966, was instead", len(ops))
	}
}

func TestFinalProgram(t *testing.T) {
	ops := opcode.ParseProgram("../puzz-input2.txt")
	regs := opcode.RunProgram(ops)
	t.Error(regs)
}
