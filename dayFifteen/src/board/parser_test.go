package board

import (
	"reflect"
	"sort"
	"testing"
)

func TestParserDataIsCorrect(t *testing.T) {
	board := Parse("../input1.txt")
	if reflect.DeepEqual(board, BoardData{}) {
		t.Errorf("Board is empty")
	}
	elves := board.getAdversaries(GoblinAlignment)
	if len(elves) != 4 {
		t.Errorf("Expected 4 elves, got %d", len(elves))
	}
	elfExpectedXLoc := []int{4, 1, 5, 4}
	elfExpectedYLoc := []int{1, 2, 2, 3}

	checkLocations(t, elfExpectedXLoc, elfExpectedYLoc, elves.toLocations())


	goblins := board.getAdversaries(ElfAlignment)
	if len(goblins) != 3 {
		t.Errorf("Expected 3 goblins, got %d", len(goblins))
	}
	goblinExpectedXLoc := []int{2, 3, 2}
	goblinExpectedYLoc := []int{1, 2, 3}
	checkLocations(t, goblinExpectedXLoc, goblinExpectedYLoc, goblins.toLocations())

	if len(board.spaces) != 35 {
		t.Errorf("Board had the wrong number of spaces, expected 35, got %d", len(board.spaces))
	}

	expectedOpenSpaces := []struct {
		y int
		x int
	}{
		{1, 1},
		{1, 3},
		{1, 5},
		{2, 2},
		{2, 4},
		{3, 1},
		{3, 3},
		{3, 5},
	}

	for _, val := range expectedOpenSpaces {
		if !board.spaces[Loc{y:val.y, x:val.x}] {
			t.Errorf("The space %+v should have been open", val)
		}
	}
}

func TestPlayerSort(t* testing.T){
	players := Players{
		Player{
			xLocation: 2,
			yLocation: 2,
		},
		Player{
			xLocation: 1,
			yLocation: 1,
		},
		Player{
			xLocation: 1,
			yLocation: 0,
		},
		Player{
			xLocation: 1,
			yLocation: 0,
		},
		Player{
			xLocation: 3,
			yLocation: 3,
		},
	}
	sort.Stable(players)

	playerExpectedXLoc := []int{1, 1, 1, 2, 3}
	playerExpectedYLoc := []int{0, 0, 1, 2, 3}
	checkLocations(t, playerExpectedXLoc,playerExpectedYLoc, players.toLocations())
}


func TestLocationSort_openSpots(t* testing.T){
	locations := Locations{
		Loc{
			x: 3,
			y: 3,
		},
		Loc{
			x: 3,
			y: 1,
		},
		Loc{
			x: 1,
			y: 3,
		},
		Loc{
			x: 4,
			y: 2,
		},
		Loc{
			x: 2,
			y: 2,
		},
		Loc{
			x: 1,
			y: 1,
		},
	}
	sort.Stable(locations)

	playerExpectedXLoc := []int{1, 3, 2, 4,1,3}
	playerExpectedYLoc := []int{1, 1, 2, 2,3,3}
	checkLocations(t, playerExpectedXLoc,playerExpectedYLoc, locations)
}

func checkLocations(t *testing.T, expectedXLocations, expectedYLocations []int, players []Location) {
	for index, elem := range players {
		playerXLoc := elem.getX()
		playerYLoc := elem.getY()
		expectedXLoc := expectedXLocations[index]
		expectedYLoc := expectedYLocations[index]
		if playerXLoc != expectedXLoc {
			t.Errorf(
				"Player number %d was not in the right x location, expected: %d, actual: %d",
				index, expectedXLoc, playerXLoc,
			)
		}
		if playerYLoc != expectedYLoc {
			t.Errorf(
				"Player number %d was not in the right y location, expected: %d, actual: %d",
				index, expectedYLoc, playerYLoc,
			)
		}
	}
}
