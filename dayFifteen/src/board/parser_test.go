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
	if len(board.elves) != 4 {
		t.Errorf("Expected 4 elves, got %d", len(board.elves))
	}
	elfExpectedXLoc := []int{4, 1, 5, 4}
	elfExpectedYLoc := []int{1, 2, 2, 3}

	checkLocations(t, elfExpectedXLoc, elfExpectedYLoc, board.elves.toLocations())

	if len(board.goblins) != 3 {
		t.Errorf("Expected 3 goblins, got %d", len(board.goblins))
	}
	goblinExpectedXLoc := []int{2, 3, 2}
	goblinExpectedYLoc := []int{1, 2, 3}
	checkLocations(t, goblinExpectedXLoc, goblinExpectedYLoc, board.goblins.toLocations())

	if len(board.spaces) != 5 {
		t.Errorf("Board had the wrong number of rows, expected 5, got %d", len(board.spaces))
	}

	for index, val := range board.spaces {
		length := len(val)
		if length != 7 {
			t.Errorf("Row number %d had the wrong number of items, expected 7, got %d", index, length)
		}
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
		if !board.spaces[val.y][val.x] {
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
