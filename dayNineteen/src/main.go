package main

import (
	"fmt"
	"opcode2"
)

func main() {

	// expectedReg := opcode2.Registers{6, 5, 6, 0, 0, 9}
	actualReg := opcode2.RunProgram("src/opcode2/puzzle_input.txt")
	fmt.Println(actualReg)
}
