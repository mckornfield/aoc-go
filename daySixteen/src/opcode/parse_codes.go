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
	beforeRegex = regexp.MustCompile(`Before: \[(\d+), (\d+), (\d+), (\d+)\]`)
	opRegex     = regexp.MustCompile(`(\d+) (\d+) (\d+) (\d+)`)
	afterRegex  = regexp.MustCompile(`After:\s+\[(\d+), (\d+), (\d+), (\d+)\]`)
)

func Parse(path string) []OpCodeData {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := []OpCodeData{}
	var data OpCodeData
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if beforeRegex.MatchString(text) {
			data = OpCodeData{}
			groups := beforeRegex.FindStringSubmatch(text)[1:]
			data.Before = FromSliceToRegisters(groups)
		} else if match := afterRegex.FindStringSubmatch(text); len(match) > 0 {
			groups := match[1:]
			data.After = FromSliceToRegisters(groups)
			result = append(result, data)
		} else if opRegex.MatchString(text) {
			groups := opRegex.FindStringSubmatch(text)[1:]
			data.Operation = FromSliceToOperation(groups)
		}
		// Skip everything else
	}
	return result
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
