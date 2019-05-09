package opcode_test

import (
	"fmt"
	"opcode"
	"testing"
)

func TestCodeDeterminer(t *testing.T) {
	dataSet := opcode.Parse("../puzz-input.txt")
	codeToSet := opcode.DetermineMappingOfOpCodeToFunc(dataSet)
	fmt.Println(codeToSet)
	if len(codeToSet) != 16 {
		fmt.Println("Should have been 16 items")
	}
	expectedVals := map[int]string{
		0:  "opcode.Muli",
		1:  "opcode.Borr",
		2:  "opcode.Gtri",
		3:  "opcode.Eqri",
		4:  "opcode.Gtrr",
		5:  "opcode.Eqir",
		6:  "opcode.Addi",
		7:  "opcode.Setr",
		8:  "opcode.Mulr",
		9:  "opcode.Addr",
		10: "opcode.Bori",
		11: "opcode.Bani",
		12: "opcode.Seti",
		13: "opcode.Eqrr",
		14: "opcode.Banr",
		15: "opcode.Gtir",
	}
	for k, v := range codeToSet {
		if expectedVals[k] != v {
			t.Errorf("Expected function for opcode %d to be %s but was %s", k, expectedVals[k], v)
		}
	}
}
