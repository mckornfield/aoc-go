package opcode2

import (
	"reflect"
	"testing"
)

func TestSampleState(t *testing.T) {
	expectedReg := Registers{6, 5, 6, 0, 0, 9}
	actualReg := RunProgram("sample_program.txt")
	if !reflect.DeepEqual(expectedReg, actualReg) {
		t.Error("Registers not equal, expected=", expectedReg, "actual=", actualReg)
	}
}

func TestPartOne(t *testing.T) {
	expectedReg := Registers{2072, 877, 877, 1, 256, 876}
	actualReg := RunProgram("puzzle_input.txt")
	if !reflect.DeepEqual(expectedReg, actualReg) {
		t.Error("Registers not equal, expected=", expectedReg, "actual=", actualReg)
	}
}

func TestParsing(t *testing.T) {
	operations, instructionPointerLoc := ParseProgram("sample_program.txt")
	if instructionPointerLoc != 0 {
		t.Error("Initial register value was not 0, was instead", instructionPointerLoc)
	}
	expectedOperation := Operation{"seti", 5, 0, 1}
	if operations[0] != expectedOperation {
		t.Error("Operations not equal, expected=", expectedOperation, "actual=", operations[0])
	}
	expectedOperation2 := Operation{"seti", 6, 0, 2}
	if operations[1] != expectedOperation2 {
		t.Error("Operations not equal, expected=", expectedOperation2, "actual=", operations[1])
	}
}

func TestParsingPt1(t *testing.T) {
	operations, instructionPointer := ParseProgram("puzzle_input.txt")
	if instructionPointer != 4 {
		t.Error("Initial register value was not 4, was instead", instructionPointer)
	}
	expectedOperation := Operation{"addi", 4, 16, 4}
	if operations[0] != expectedOperation {
		t.Error("Operations not equal, expected=", expectedOperation, "actual=", operations[0])
	}
	expectedOperation2 := Operation{"seti", 1, 5, 1}
	if operations[1] != expectedOperation2 {
		t.Error("Operations not equal, expected=", expectedOperation2, "actual=", operations[1])
	}
}
