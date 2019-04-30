package board

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ElfAlignment    = iota
	GoblinAlignment = iota
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
	spaces := make([][]Space, 5)
	for scanner.Scan() {
		text := scanner.Text()
		spaceRow := make([]Space, len(text))
		for x, c := range text {
			switch c {
			case 'E':
				player := Player{
					alignment: ElfAlignment,
					xLocation: x,
					yLocation: y,
				}
				data.elves = append(data.elves, player)
				spaceRow[x] = false
			case 'G':
				player := Player{
					alignment: GoblinAlignment,
					xLocation: x,
					yLocation: y,
				}
				data.goblins = append(data.goblins, player)
				spaceRow[x] = false
			case '#':
				spaceRow[x] = false
			case '.':
				spaceRow[x] = true
			}
		}
		spaces[y] = spaceRow
		y++
	}
	data.spaces = spaces
	return data
}

// BoardData the board with elves, goblins and spaces/obstacles
type BoardData struct {
	elves   []Player
	goblins []Player
	spaces  [][]Space
}

// Player either elf or goblin
type Player struct {
	alignment int
	xLocation int
	yLocation int
}

// Player either elf or goblin
type Space bool

func (p Player) getLocation() (int, int) {
	return p.xLocation, p.yLocation
}
