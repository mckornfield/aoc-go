package waterflow

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	lineRegex = regexp.MustCompile(`(\w)=(\d+), (\w)=(\d+)..(\d+)`)
)

// LocationSet what do you think this is?
type LocationSet map[Location]bool

// MakeSet really I have to put a comment?
func MakeSet(locs ...Location) LocationSet {
	set := make(map[Location]bool)
	for _, loc := range locs {
		set[loc] = true
	}
	return set
}

func (l LocationSet) Clone() LocationSet {
	locs := []Location{}
	for loc := range l {
		locs = append(locs, loc)
	}
	return MakeSet(locs...)
}

// ParseFile HAHhahahaa
func ParseFile(path string) (LocationSet, int) {
	scanner, close := fileToScanner(path)
	defer close()
	locsMap := make(map[Location]bool)
	scanner.Split(bufio.ScanLines)
	maxY := 0
	for scanner.Scan() {
		text := scanner.Text()
		locsForLine := ParseLine(text)
		for _, loc := range locsForLine {
			if loc.Y > maxY {
				maxY = loc.Y
			}
			locsMap[loc] = true
		}
	}
	return locsMap, maxY
}

func fileToScanner(path string) (*bufio.Scanner, func() error) {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file.Close
}

// ParseLine hahahaha
func ParseLine(line string) []Location {
	locs := []Location{}
	groups := lineRegex.FindStringSubmatch(line)
	fmt.Println(groups)
	if groups[1] == "x" {
		x := parseIntWithPanic(groups[2])
		yMin := parseIntWithPanic(groups[4])
		yMax := parseIntWithPanic(groups[5])
		for yI := yMin; yI <= yMax; yI++ {
			locs = append(locs, Location{X: x, Y: yI})
		}
	} else if groups[1] == "y" {
		y := parseIntWithPanic(groups[2])
		xMin := parseIntWithPanic(groups[4])
		xMax := parseIntWithPanic(groups[5])
		for xI := xMin; xI <= xMax; xI++ {
			locs = append(locs, Location{X: xI, Y: y})
		}
	}
	return locs
}

func parseIntWithPanic(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return intVal
}

// Location blah
type Location struct {
	X int
	Y int
}
