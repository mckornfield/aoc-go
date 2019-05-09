package opcode

import (
	"errors"
	"reflect"
	"runtime"
)

const (
	Index1Error = "Index 1 Out of Range"
	Index2Error = "Index 2 Out of Range"
)

func Addr(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, copy Registers) {
		copy[op.Output] = copy[op.First] + copy[op.Second]
	}
	return runTemplate(op, input, validateRegisters, calculation)
}

func Addi(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, copy Registers) {
		copy[op.Output] = copy[op.First] + op.Second
	}
	return runTemplate(op, input, validateFirstRegister, calculation)
}

func Mulr(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, copy Registers) {
		copy[op.Output] = copy[op.First] * copy[op.Second]
	}
	return runTemplate(op, input, validateRegisters, calculation)
}

func Muli(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		input[op.Output] = input[op.First] * op.Second
	}
	return runTemplate(op, input, validateFirstRegister, calculation)
}

func Banr(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		input[op.Output] = input[op.First] & input[op.Second]
	}
	return runTemplate(op, input, validateRegisters, calculation)
}

func Bani(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		input[op.Output] = input[op.First] & op.Second
	}
	return runTemplate(op, input, validateFirstRegister, calculation)
}

func Borr(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		input[op.Output] = input[op.First] | op.Second
	}
	return runTemplate(op, input, validateRegisters, calculation)
}

func Bori(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		input[op.Output] = input[op.First] | op.Second
	}
	return runTemplate(op, input, validateFirstRegister, calculation)
}

func Setr(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		input[op.Output] = input[op.First]
	}
	return runTemplate(op, input, validateRegisters, calculation)
}

func Seti(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		input[op.Output] = op.First
	}
	return runTemplate(op, input, validateFirstRegister, calculation)
}

func Gtir(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		if op.First > input[op.Second] {
			input[op.Output] = 1
		} else {
			input[op.Output] = 0
		}
	}
	return runTemplate(op, input, validateSecondRegister, calculation)
}

func Gtri(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		if input[op.First] > op.Second {
			input[op.Output] = 1
		} else {
			input[op.Output] = 0
		}
	}
	return runTemplate(op, input, validateFirstRegister, calculation)
}

func Gtrr(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		if input[op.First] > input[op.Second] {
			input[op.Output] = 1
		} else {
			input[op.Output] = 0
		}
	}
	return runTemplate(op, input, validateRegisters, calculation)
}

func Eqir(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		if op.First == input[op.Second] {
			input[op.Output] = 1
		} else {
			input[op.Output] = 0
		}
	}
	return runTemplate(op, input, validateSecondRegister, calculation)
}

func Eqri(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		if input[op.First] == op.Second {
			input[op.Output] = 1
		} else {
			input[op.Output] = 0
		}
	}
	return runTemplate(op, input, validateFirstRegister, calculation)
}

func Eqrr(op Operation, input Registers) (Registers, error) {
	calculation := func(op Operation, input Registers) {
		if input[op.First] == input[op.Second] {
			input[op.Output] = 1
		} else {
			input[op.Output] = 0
		}
	}
	return runTemplate(op, input, validateRegisters, calculation)
}

func runTemplate(op Operation, input Registers, validation ValidationMethod, calculation Calculation) (Registers, error) {
	if invalid, err := validation(op, input); invalid {
		return Registers{}, err
	}
	copy := append(input[:0:0], input...)
	calculation(op, copy)
	return copy, nil
}

type OpCodeFunction func(op Operation, input Registers) (Registers, error)

func (opFun OpCodeFunction) GetFunctionName() string {
	return runtime.FuncForPC(reflect.ValueOf(opFun).Pointer()).Name()
}

type Calculation func(op Operation, input Registers)

type ValidationMethod func(op Operation, input Registers) (bool, error)

type OpCodeData struct {
	Before Registers
	Operation
	After Registers
}

func validateRegisters(op Operation, input Registers) (bool, error) {
	indexLength := len(input) - 1
	if indexLength < op.First {
		return true, errors.New(Index1Error)
	}
	if indexLength < op.Second {
		return true, errors.New(Index2Error)
	}
	return false, nil
}

func validateSecondRegister(op Operation, input Registers) (bool, error) {
	indexLength := len(input) - 1
	if indexLength < op.Second {
		return true, errors.New(Index2Error)
	}
	return false, nil
}

func validateFirstRegister(op Operation, input Registers) (bool, error) {
	indexLength := len(input) - 1
	if indexLength < op.First {
		return true, errors.New(Index1Error)
	}
	return false, nil
}

type Registers []int

func (r Registers) Equal(otherRegisters Registers) bool {
	if len(r) != len(otherRegisters) {
		return false
	}
	for i, val := range r {
		if val != otherRegisters[i] {
			return false
		}
	}
	return true
}

type Operation struct {
	OpCode int
	First  int
	Second int
	Output int
}
