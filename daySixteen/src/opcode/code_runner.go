package opcode

import "fmt"

var opcodeFuncs = []OpCodeFunction{
	Addr,
	Addi,
	Mulr,
	Muli,
	Banr,
	Bani,
	Borr,
	Bori,
	Setr,
	Seti,
	Getr,
	Gtir,
	Gtri,
	Gtrr,
	Eqir,
	Eqri,
	Eqrr,
}

func TryCodesAndGetCountOfMatches(data OpCodeData) int {
	// op Operation, input, expectedOutput Registers
	matchChan := make(chan int)

	for _, opcodeFunc := range opcodeFuncs {
		go func(opcodeFunc OpCodeFunction, OpmatchChan chan int) {
			output, err := opcodeFunc(data.Operation, data.Before)
			if err == nil && data.After.Equal(output) {
				fmt.Println(opcodeFunc.GetFunctionName(), "totally matched!")
				matchChan <- 1
			}
			matchChan <- 0
		}(opcodeFunc, matchChan)
	}

	matches := 0
	for i := 0; i < len(opcodeFuncs); i++ {
		matches += <-matchChan
	}
	return matches
}
