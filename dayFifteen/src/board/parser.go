package board

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	ElfAlignment    = iota
	GoblinAlignment = iota
	playerHealth    = 200
)

// Parse takes a file location and parses the board to a data structure
func Parse(path string) BoardData {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	y := 0
	data := BoardData{}
	spaces := make(map[Location]Space)
	id := 0
	for scanner.Scan() {
		text := scanner.Text()
		for x, c := range text {
			switch c {
			case 'E', 'G':
				var alignment int
				if c == 'E' {
					alignment = ElfAlignment
				} else {
					alignment = GoblinAlignment
				}
				player := Player{
					id:        id,
					health:    playerHealth,
					alignment: alignment,
					xLocation: x,
					yLocation: y,
				}
				id++
				data.allPlayers = append(data.allPlayers, player)
				fallthrough
			case '#':
				spaces[Loc{x: x, y: y}] = false
			case '.':
				spaces[Loc{x: x, y: y}] = true
			}
		}
		y++
	}
	data.spaces = spaces
	return data
}

func getFileAsString(path string) ([]byte, error) {
	absPath, _ := filepath.Abs(path)
	return ioutil.ReadFile(absPath)
}
