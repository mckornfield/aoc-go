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
			case 'E', 'G':
				var alignment int
				var appendSlice *Players
				if c == 'E' {
					alignment = ElfAlignment
					appendSlice = &data.elves
				} else {
					alignment = GoblinAlignment
					appendSlice = &data.goblins
				}
				player := Player{
					health:    10,
					alignment: alignment,
					xLocation: x,
					yLocation: y,
				}
				*appendSlice = append(*appendSlice, player)
				data.allPlayers = append(data.allPlayers, player)
				fallthrough
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

// Players on a board, used for sorting
type Players []Player

// BoardData the board with elves, goblins and spaces/obstacles
type BoardData struct {
	elves      Players
	goblins    Players
	allPlayers Players
	spaces     [][]Space
}

func (pl Players) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func (pl Players) Len() int { return len(pl) }

func (pl Players) Less(i, j int) bool {
	pl1 := pl[i]
	pl2 := pl[j]
	return isLess(pl1, pl2)
}

// Player either elf or goblin
type Player struct {
	health    int
	alignment int
	xLocation int
	yLocation int
	Location
}

func (p Player) getX() int {
	return p.xLocation
}

func (p Player) getY() int {
	return p.yLocation
}

// Space on the board
type Space bool

func (p Player) getLocation() (int, int) {
	return p.xLocation, p.yLocation
}
