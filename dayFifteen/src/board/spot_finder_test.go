package board

import(
	"testing"
	"fmt"
	"sort"
)

func TestSpotFinder(t *testing.T) {
	board := Parse("../input2.txt")
	openSpots := getOpenSpots(board, ElfAlignment)
	if len(openSpots) == 0 {
		t.Errorf("should have had more than one open spot")
	}
	
	fmt.Println(openSpots)
	sort.Stable(openSpots)
	playerExpectedXLoc := []int{3, 5, 2, 5, 1, 3}
	playerExpectedYLoc := []int{1, 1, 2, 2, 3, 3}
	checkLocations(t, playerExpectedXLoc, playerExpectedYLoc, openSpots)
}

func TestSpotFinder_input1(t *testing.T) {
	board := Parse("../input1.txt")
	openSpots := getOpenSpots(board, ElfAlignment)
	if len(openSpots) == 0 {
		t.Errorf("should have had more than one open spot")
	}
	
	fmt.Println(openSpots)
	sort.Stable(openSpots)
	playerExpectedXLoc := []int{3, 5, 2, 5, 1, 3}
	playerExpectedYLoc := []int{1, 1, 2, 2, 3, 3}
	checkLocations(t, playerExpectedXLoc, playerExpectedYLoc, openSpots)
}