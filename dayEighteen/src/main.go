package main

import (
	"fmt"
	"lumber"
)

func main() {
	ans, err := lumber.GetAnswer("src/lumber/part_1_input.txt", 1000000000)
	if err != nil {
		panic(err)
	}
	fmt.Println("Answer is", ans)
}
