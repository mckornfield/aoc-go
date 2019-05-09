package opcode

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ParseProgram(path string) []Operation {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	ops := []Operation{}
	for scanner.Scan() {
		text := scanner.Text()
		if opRegex.MatchString(text) {
			groups := opRegex.FindStringSubmatch(text)[1:]
			ops = append(ops, FromSliceToOperation(groups))
		}
	}
	return ops
}

var opCodeMap = map[int]OpCodeFunction{
	0:  Muli,
	1:  Borr,
	2:  Gtri,
	3:  Eqri,
	4:  Gtrr,
	5:  Eqir,
	6:  Addi,
	7:  Setr,
	8:  Mulr,
	9:  Addr,
	10: Bori,
	11: Bani,
	12: Seti,
	13: Eqrr,
	14: Banr,
	15: Gtir,
}

func RunProgram(operations []Operation) Registers {
	// Clone register
	currentReg := Registers{0, 0, 0, 0}
	// Run program
	var err error
	for _, op := range operations {
		opToPerform := opCodeMap[op.OpCode]
		currentReg, err = opToPerform(op, currentReg)
		if err != nil {
			panic(err)
		}
	}
	return currentReg
}
