package opcode

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	beforeRegex = regexp.MustCompile("Before: \\[(\\d), (\\d), (\\d), (\\d)\\]")
	opRegex     = regexp.MustCompile(`(\d), (\\d), (\\d), (\\d)`)
	afterRegex  = regexp.MustCompile("After: \\[(\\d), (\\d), (\\d), (\\d)\\]")
)

func Parse(path string) OpCodeData {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	data := OpCodeData{}
	for scanner.Scan() {
		text := scanner.Text()
		if beforeRegex.MatchString(text) {
			data.Before = FromSliceToRegisters(beforeRegex.SubexpNames())
		} else if afterRegex.MatchString(text) {
			data.After = FromSliceToRegisters(afterRegex.SubexpNames())
		} else if opRegex.MatchString(text) {
			data.Operation = FromSliceToOperation(opRegex.SubexpNames())
		}
	}
	return data
}

func toInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return intVal
}

func FromSliceToRegisters(data []string) Registers {
	registers := Registers{}
	for _, element := range data {
		intVal := toInt(element)
		registers = append(registers, intVal)
	}
	return registers
}

func FromSliceToOperation(data []string) Operation {
	operations := Operation{
		OpCode: toInt(data[0]),
		First:  toInt(data[1]),
		Second: toInt(data[2]),
		Output: toInt(data[3]),
	}
	return operations
}

type OpCodeData struct {
	Before Registers
	Operation
	After Registers
}

type Registers []int

type Operation struct {
	OpCode int
	First  int
	Second int
	Output int
}
