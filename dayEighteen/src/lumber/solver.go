package lumber

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetAnswer(inputLocation string, count int) (int, error) {
	if _, err := os.Stat(inputLocation); err != nil {
		return 0, err
	}

	grid, err := fileContentsToGrid(inputLocation)
	if err != nil {
		return 0, err
	}

	changedGrid := grid.roleForwardNTimes(int64(count))
	lumberYards, trees := changedGrid.GetLumberYardsAndTrees()
	answer := trees * lumberYards

	return answer, nil
}

func fileContentsToGrid(inputLocation string) (Grid, error) {
	file, err := os.Open(inputLocation)
	if err != nil {
		return Grid{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := Grid{}
	for scanner.Scan() {
		l := scanner.Text()
		lineLength := len(l)
		row := make(Row, lineLength)
		for index, ch := range []rune(l) {
			row[index] = ch
		}
		grid = append(grid, row)
	}
	return grid, nil
}

type Grid []Row

type Pair struct {
	X int
	Y int
}

func (r Row) String() string {
	return string(r)
}

func (g Grid) String() string {
	var sb strings.Builder
	sb.WriteRune('\n')
	for _, row := range g {
		sb.WriteString(row.String())
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g Grid) GetLumberYardsAndTrees() (int, int) {
	lumberYards := 0
	trees := 0
	for _, row := range g {
		for _, ch := range row {
			if ch == '#' {
				lumberYards++
			} else if ch == '|' {
				trees++
			}
		}
	}

	return lumberYards, trees
}

func (originalGrid Grid) roleForwardNTimes(times int64) Grid {
	currentGrid := originalGrid
	gridConfigurations := make(map[string]bool, 0)
	previousGridStates := make([]Grid, 0)
	haveStartedLooping := false
	haveFiguredOutLoopingPattern := false
	previousGridPos := 0
	for i := int64(0); i < times; i++ {
		// fmt.Println(strconv.Itoa(int(times)))
		if haveFiguredOutLoopingPattern {
			// fmt.Println(len(previousGridStates))
			// for _, _ = range previousGridStates {
			// 	// fmt.Println(g.String())
			// 	fmt.Println("--------------------------------------------------------")
			// }
			// panic("abc")
			if previousGridPos > len(previousGridStates)-1 {
				previousGridPos = 0
			}
			currentGrid = previousGridStates[previousGridPos]
			previousGridPos++
			continue
		}
		currentGrid = currentGrid.roleForwardOne()

		if haveStartedLooping {
			// fmt.Println("got here")
			previousGridStates = append(previousGridStates, currentGrid)
		}
		// hasher.Reset()
		// hasher.Write([]byte(currentGrid.String()))
		gridHash := currentGrid.String()
		if _, exists := gridConfigurations[gridHash]; exists {
			if haveStartedLooping {
				fmt.Println("Figured out looping at iteration" + strconv.Itoa(int(i)))
				haveFiguredOutLoopingPattern = true
				continue
			} else {
				fmt.Println(strconv.Itoa(int(i)))
				fmt.Println("Started looping at iteration" + strconv.Itoa(int(i)))
				// previousGridStates = append(previousGridStates, currentGrid)
				haveStartedLooping = true
				gridConfigurations = make(map[string]bool, 0)
			}
		}
		gridConfigurations[gridHash] = true
	}
	fmt.Println(currentGrid.String())
	return currentGrid
}

func (g Grid) roleForwardOne() Grid {

	// An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
	// An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.
	// An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.

	newGrid := Grid{}
	for y, row := range g {
		lineLength := len(row)
		newRow := make(Row, lineLength)
		for x, spot := range row {
			adjacentTrees := 0
			adjacentLumberYards := 0
			for _, del_x := range []int{-1, 0, 1} {
				for _, del_y := range []int{-1, 0, 1} {
					if del_y == 0 && del_x == 0 {
						continue
					}
					possibleTree := g.safeGet(y+del_y, x+del_x)
					if possibleTree == '|' {
						adjacentTrees++
					} else if possibleTree == '#' {
						adjacentLumberYards++
					}
				}
			}
			newSpot := spot
			if spot == '.' && adjacentTrees > 2 {
				newSpot = '|'
			} else if spot == '|' && adjacentLumberYards > 2 {
				newSpot = '#'
			} else if spot == '#' && (adjacentLumberYards == 0 || adjacentTrees == 0) {
				newSpot = '.'
			}
			newRow[x] = newSpot
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func (g Grid) safeGet(y, x int) rune {
	maxX := len(g[0])
	maxY := len(g)
	if -1 < x && x < maxX && -1 < y && y < maxY {
		return g[y][x]
	} else {
		return '.'
	}
}

type Row []rune
