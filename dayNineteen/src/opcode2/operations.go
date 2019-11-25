package opcode2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func RunProgram(fileLocation string) Registers {
	operations, instructionPointerLocation := ParseProgram(fileLocation)
	count := 0
	instructionPointer := int64(0)
	regs := Registers{int64(0), int64(0), int64(0), int64(0), int64(0), int64(0)}
	for instructionPointer < int64(len(operations)) {
		regs[instructionPointerLocation] = int64(instructionPointer)
		currentOp := operations[instructionPointer]
		funcToRun := GetFunctionByName(currentOp.OpInstruction)
		// fmt.Println(funcToRun)
		var err error
		regs, err = funcToRun(currentOp, regs)
		if err != nil {
			panic(err)
		}
		instructionPointer = regs[instructionPointerLocation]
		instructionPointer++
		count++
		// if count%100 == 0 {
		// 	fmt.Println(instructionPointer)
		// 	fmt.Println(currentOp.OpInstruction)
		// 	fmt.Println(regs)
		// }
		// if count > 1000 {
		// 	panic("a")
		// }
		// if instructionPointer == 1 {
		// 	panic("Ah")
		// }
	}
	return regs
}

var (
	opRegex = regexp.MustCompile(`(\w+) (\d+) (\d+) (\d+)`)
)

func ParseProgram(path string) ([]Operation, int64) {
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
	var initialRegisterValue int64
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#ip") {
			initialRegisterValue = int64(toInt(text[4:]))
		} else if opRegex.MatchString(text) {
			groups := opRegex.FindStringSubmatch(text)[1:]
			FromSliceToOperation(groups)
			ops = append(ops, FromSliceToOperation(groups))
		} else {
			fmt.Println("No match for", text)
		}
	}
	return ops, initialRegisterValue
}

func FromSliceToOperation(data []string) Operation {
	operations := Operation{
		OpInstruction: data[0],
		First:         int64(toInt(data[1])),
		Second:        int64(toInt(data[2])),
		Output:        int64(toInt(data[3])),
	}
	return operations
}

func toInt(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}
