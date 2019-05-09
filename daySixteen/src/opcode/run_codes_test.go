package opcode_test

import (
	"errors"
	"opcode"
	"testing"
)

func TestOpCodes(t *testing.T) {
	op := opcode.Operation{
		First:  1,
		Second: 3,
		Output: 2,
	}
	reg := opcode.Registers{1, 2, 3, 4}

	configs := []struct {
		opFunc        opcode.OpCodeFunction
		expectedReg   opcode.Registers
		expectedError error
	}{
		{opcode.Addr, opcode.Registers{1, 2, 6, 4}, nil},
		{opcode.Addi, opcode.Registers{1, 2, 5, 4}, nil},
		{opcode.Mulr, opcode.Registers{1, 2, 8, 4}, nil},
		{opcode.Muli, opcode.Registers{1, 2, 6, 4}, nil},
		{opcode.Banr, opcode.Registers{1, 2, 0, 4}, nil},
		{opcode.Bani, opcode.Registers{1, 2, 2, 4}, nil},
		{opcode.Borr, opcode.Registers{1, 2, 3, 4}, nil},
		{opcode.Bori, opcode.Registers{1, 2, 3, 4}, nil},
		{opcode.Setr, opcode.Registers{1, 2, 2, 4}, nil},
		{opcode.Seti, opcode.Registers{1, 2, 1, 4}, nil},
		{opcode.Getr, opcode.Registers{1, 2, 6, 4}, nil},
		{opcode.Gtir, opcode.Registers{1, 2, 0, 4}, nil}, // Need to add tests for greater and equal
		{opcode.Gtri, opcode.Registers{1, 2, 0, 4}, nil},
		{opcode.Gtrr, opcode.Registers{1, 2, 0, 4}, nil},
		{opcode.Eqir, opcode.Registers{1, 2, 0, 4}, nil},
		{opcode.Eqri, opcode.Registers{1, 2, 0, 4}, nil},
		{opcode.Eqrr, opcode.Registers{1, 2, 0, 4}, nil},
	}
	for _, config := range configs {
		opcodeFunc := config.opFunc
		actualReg, err := opcodeFunc(op, reg)
		if !actualReg.Equal(config.expectedReg) {
			t.Errorf("Function %s, Expected %v , Actual %v", opcodeFunc.GetFunctionName(), config.expectedReg, actualReg)
		}
		if err != config.expectedError {
			t.Errorf("Function %s, Expected %v , Actual %v", opcodeFunc.GetFunctionName(), config.expectedError, err)
		}
	}
}

func TestOpCodes_Index1Error(t *testing.T) {
	op := opcode.Operation{
		First:  5,
		Second: 3,
		Output: 2,
	}
	reg := opcode.Registers{1, 2, 3, 4}

	opcodeFuncs := []opcode.OpCodeFunction{
		opcode.Addr,
		opcode.Addi,
		opcode.Mulr,
		opcode.Muli,
		opcode.Banr,
		opcode.Bani,
		opcode.Borr,
		opcode.Bori,
		opcode.Setr,
		opcode.Seti,
		opcode.Getr,
		// opcode.Gtir,
		opcode.Gtri,
		opcode.Gtrr,
		// opcode.Eqir,
		opcode.Eqri,
		opcode.Eqrr,
	}
	expectedError := errors.New(opcode.Index1Error)
	expectedReg := opcode.Registers{}
	for _, opcodeFunc := range opcodeFuncs {
		actualReg, err := opcodeFunc(op, reg)
		if !actualReg.Equal(expectedReg) {
			t.Errorf("Function %s, Expected %v , Actual %v", opcodeFunc.GetFunctionName(), expectedReg, actualReg)
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("Function %s, Expected %v , Actual %v", opcodeFunc.GetFunctionName(), expectedError, err)
		}
	}
}

func TestOpCodes_Index2Error(t *testing.T) {
	op := opcode.Operation{
		First:  1,
		Second: 5,
		Output: 2,
	}
	reg := opcode.Registers{1, 2, 3, 4}

	opcodeFuncs := []opcode.OpCodeFunction{
		opcode.Addr,
		// opcode.Addi,
		opcode.Mulr,
		// opcode.Muli,
		opcode.Banr,
		// opcode.Bani,
		opcode.Borr,
		// opcode.Bori,
		opcode.Setr,
		// opcode.Seti,
		opcode.Getr,
		opcode.Gtir,
		// opcode.Gtri,
		opcode.Gtrr,
		opcode.Eqir,
		// opcode.Eqri,
		opcode.Eqrr,
	}
	expectedError := errors.New(opcode.Index2Error)
	expectedReg := opcode.Registers{}
	for _, opcodeFunc := range opcodeFuncs {
		actualReg, err := opcodeFunc(op, reg)
		if !actualReg.Equal(expectedReg) {
			t.Errorf("Function %s, Expected %v , Actual %v", opcodeFunc.GetFunctionName(), expectedReg, actualReg)
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("Function %s, Expected %v , Actual %v", opcodeFunc.GetFunctionName(), expectedError, err)
		}
	}
}

func areErrorsEqual(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 != nil && err2 != nil {
		return err1.Error() == err2.Error()
	}
	return false
}
