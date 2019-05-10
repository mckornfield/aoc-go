package waterflow_test

import (
	"fmt"
	"reflect"
	"testing"
	"waterflow"
	. "waterflow"
)

func TestFindNextPoint(t *testing.T) {
	startPoint := Location{X: 2, Y: 2}
	max := 5
	configs := []struct {
		point             Location
		blockedPoints     LocationSet
		clayPoints        LocationSet
		expectedNextPoint Location
	}{
		{startPoint, MakeSet(Location{X: 1, Y: 1}), MakeSet(Location{X: 1, Y: 1}), Location{X: 2, Y: 3}},
		{startPoint, MakeSet(Location{X: 2, Y: 3}), MakeSet(Location{X: 1, Y: 1}), Location{X: 0, Y: 0}}, // No point below
		{startPoint, MakeSet(Location{X: 2, Y: 3}), MakeSet(Location{X: 2, Y: 3}), Location{X: 1, Y: 2}}, // Run over left
		{startPoint, MakeSet(Location{X: 2, Y: 3}, Location{X: 1, Y: 3}), MakeSet(Location{X: 1, Y: 1}), Location{X: 1, Y: 2}},
		{startPoint, MakeSet(Location{X: 2, Y: 3}, Location{X: 1, Y: 2}, Location{X: 3, Y: 3}), MakeSet(Location{X: 1, Y: 1}), Location{X: 3, Y: 2}},
		{startPoint, MakeSet(Location{X: 2, Y: 3}, Location{X: 1, Y: 2}), MakeSet(Location{X: 1, Y: 1}), Location{X: 0, Y: 0}}, // No point below
		{startPoint, MakeSet(Location{X: 2, Y: 3}, Location{X: 1, Y: 2}), MakeSet(Location{X: 2, Y: 3}), Location{X: 3, Y: 2}}, // Run over right
		{startPoint, MakeSet(Location{X: 2, Y: 3}, Location{X: 1, Y: 2}, Location{X: 3, Y: 2}), MakeSet(Location{X: 1, Y: 1}), Location{X: 0, Y: 0}},
		{startPoint, MakeSet(Location{X: 2, Y: 3}, Location{X: 1, Y: 2}, Location{X: 3, Y: 2}), MakeSet(Location{X: 1, Y: 1}), Location{X: 0, Y: 0}},
		{Location{X: 3, Y: 5}, MakeSet(Location{X: 2, Y: 5}), MakeSet(Location{X: 1, Y: 1}), Location{X: 0, Y: 0}}, // Below max
	}
	for _, config := range configs {
		loc, err := DetermineNextWaterPoint(config.point, config.blockedPoints, config.clayPoints, max)
		fmt.Println(loc, err)
		if !reflect.DeepEqual(loc, config.expectedNextPoint) {
			t.Errorf("Expected: %v, Actual: %v", config.expectedNextPoint, loc)
		}
		if (loc == Location{}) && err == nil {
			t.Error("Expected an error, got nothing")
		}
	}
}

func TestStackOperations(t *testing.T) {
	stack := Stack{}
	loc, _ := stack.Pop()
	if (loc != Location{0, 0}) {
		t.Error("Expected empty value back from pop")
	}
	newLoc := Location{1, 5}
	stack.Push(newLoc)

	if val, _ := stack.Peek(); !reflect.DeepEqual(newLoc, val) {
		t.Error("Expected value to be peeked")
	}

	if val, _ := stack.Pop(); !reflect.DeepEqual(newLoc, val) {
		t.Error("Expected value to be popped")
	}

	if _, err := stack.Peek(); err == nil {
		t.Error("Expected error from peek")
	}

}

func TestNumberOfSquaresCovered(t *testing.T) {
	start := Location{X: 500, Y: 0}
	locationSet, maxY := waterflow.ParseFile("sample1.txt")
	num := DetermineNumberOfSquaresCovered(start, maxY, locationSet)
	if num != 57 {
		t.Errorf("Expected 57, got %d", num)
	}
}
