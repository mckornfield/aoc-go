package opcode

var OpcodeFuncs = []OpCodeFunction{
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
	Gtir,
	Gtri,
	Gtrr,
	Eqir,
	Eqri,
	Eqrr,
}

func CountNumberOfMatchesOverThree(dataSet []OpCodeData) int {
	overThreeChan := make(chan int)

	for _, data := range dataSet {
		go func(data OpCodeData) {
			matches := TryCodesAndGetCountOfMatches(data)
			if matches >= 3 {
				overThreeChan <- 1
			} else {
				overThreeChan <- 0
			}
		}(data)
	}

	sum := 0
	for i := 0; i < len(dataSet); i++ {
		sum += <-overThreeChan
	}
	return sum
}

func TryCodesAndGetCountOfMatches(data OpCodeData) int {
	// op Operation, input, expectedOutput Registers
	matchChan := make(chan int)

	for _, opcodeFunc := range OpcodeFuncs {
		go func(opcodeFunc OpCodeFunction) {
			output, err := opcodeFunc(data.Operation, data.Before)
			if err == nil && data.After.Equal(output) {
				// fmt.Println(opcodeFunc.GetFunctionName(), "totally matched!")
				matchChan <- 1
			} else {
				// fmt.Println(opcodeFunc.GetFunctionName(), err)
				// fmt.Println(opcodeFunc.GetFunctionName(), "Expected", data.After, "Actual", output)
				matchChan <- 0
			}
		}(opcodeFunc)
	}

	matches := 0
	for i := 0; i < len(OpcodeFuncs); i++ {
		matches += <-matchChan
	}
	// fmt.Println(matches, data)
	return matches
}
