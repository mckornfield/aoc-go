package board

import(
	"testing"
	"sort"
)

func TestSpotFinder(t *testing.T) {
	board := Parse("../input2.txt")
	openSpots := getOpenSpots(ElfAlignment, board)
	if len(openSpots) == 0 {
		t.Errorf("should have had more than one open spot")
	}
	
	// fmt.Println(openSpots)
	sort.Stable(openSpots)
	playerExpectedXLoc := []int{3, 5, 2, 5, 1, 3}
	playerExpectedYLoc := []int{1, 1, 2, 2, 3, 3}
	checkLocations(t, playerExpectedXLoc, playerExpectedYLoc, openSpots)
}

func TestSpotFinder_input1(t *testing.T) {
	board := Parse("../input1.txt")
	openSpots := getOpenSpots(ElfAlignment, board)
	if len(openSpots) == 0 {
		t.Errorf("should have had more than one open spot")
	}
	
	sort.Stable(openSpots)
	// fmt.Println(openSpots)
	playerExpectedXLoc := []int{1, 3, 2, 4, 1, 3}
	playerExpectedYLoc := []int{1, 1, 2, 2, 3, 3}
	checkLocations(t, playerExpectedXLoc, playerExpectedYLoc, openSpots)
}

func TestSpotFinder_reachableSpots_input2(t *testing.T) {
	board := Parse("../input2.txt")
	elf := board.getElves()[0]
	reachableSpots := getReachableSpots(elf, board)

	if len(reachableSpots) != 4 {
		t.Errorf("should have had more than one reachable spot")
	}
	for _, expected := range []struct{
		x int
		y int
		dist int
	}{
		{1,3,1},
		{2,2,1},
		{3,1,1},
		{3,3,3},
	}{
		expectedLoc := Loc{x:expected.x,y:expected.y}
		expectedDist := expected.dist
		dist, present := reachableSpots[expectedLoc]
		if present && dist != dist {
			t.Errorf("Distance for point %+v should have been 1, was instead %d", expectedLoc, expectedDist)
		} else if !present {
			t.Errorf("point %+v should have been present", expectedLoc)
		}
	}
}

func TestSpotFinder_getClosestChosenSpot(t *testing.T){
	board := Parse("../input2.txt")
	elf := board.getElves()[0]
	closestSpot, err := getClosestAdversarySpot(elf, board)
	if err != nil {
		t.Errorf(err.Error())
	}
	location := Loc{x:3, y:1}
	if closestSpot != location {
		t.Errorf("Closest spot should have been 1,3, was %+v", closestSpot)
	}
}