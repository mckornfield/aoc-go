package opcode_test

import (
	"opcode"
	"testing"
)

func TestOpCodeRunner(t *testing.T) {
	val := opcode.Parse("../input1.txt")
	num := opcode.TryCodesAndGetCountOfMatches(val[0])
	if num != 2 {
		t.Errorf("Should have been %d matches, got %d", 2, num)
	}
}

func TestTheFirstPart(t *testing.T) {
	dataSet := opcode.Parse("../puzz-input.txt")
	matches := opcode.CountNumberOfMatchesOverThree(dataSet)
	if matches != 563 {
		t.Errorf("Should have been %d matches, got %d", 563, matches)
	}
}
