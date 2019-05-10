package waterflow_test

import (
	"fmt"
	"testing"
	"waterflow"
)

func TestParseOfInitialData(t *testing.T) {
	lineMap := waterflow.ParseFile("sample1.txt")
	if len(lineMap) != 34 {
		t.Errorf("length of map was %d not 34", len(lineMap))
	}
	// Check points
	val := waterflow.MakeSet([]waterflow.Location{
		waterflow.Location{X: 495, Y: 2}, waterflow.Location{X: 495, Y: 3}, waterflow.Location{X: 495, Y: 4},
		waterflow.Location{X: 495, Y: 5}, waterflow.Location{X: 495, Y: 6}, waterflow.Location{X: 495, Y: 7},
		waterflow.Location{X: 496, Y: 7}, waterflow.Location{X: 497, Y: 7}, waterflow.Location{X: 498, Y: 7},
		waterflow.Location{X: 499, Y: 7}, waterflow.Location{X: 500, Y: 7}, waterflow.Location{X: 501, Y: 7},
		waterflow.Location{X: 501, Y: 7}, waterflow.Location{X: 501, Y: 6}, waterflow.Location{X: 501, Y: 5},
		waterflow.Location{X: 501, Y: 4}, waterflow.Location{X: 501, Y: 3}, waterflow.Location{X: 506, Y: 1},
		waterflow.Location{X: 506, Y: 2}, waterflow.Location{X: 498, Y: 2}, waterflow.Location{X: 498, Y: 3},
		waterflow.Location{X: 498, Y: 4}, waterflow.Location{X: 498, Y: 10}, waterflow.Location{X: 498, Y: 11},
		waterflow.Location{X: 498, Y: 12}, waterflow.Location{X: 498, Y: 13}, waterflow.Location{X: 499, Y: 13},
		waterflow.Location{X: 500, Y: 13}, waterflow.Location{X: 501, Y: 13}, waterflow.Location{X: 502, Y: 13},
		waterflow.Location{X: 503, Y: 13}, waterflow.Location{X: 504, Y: 13}, waterflow.Location{X: 504, Y: 12},
		waterflow.Location{X: 504, Y: 11}, waterflow.Location{X: 504, Y: 10}, waterflow.Location{X: 506, Y: 1},
		waterflow.Location{X: 506, Y: 2},
	})
	for loc, _ := range val {
		if _, present := lineMap[loc]; !present {
			t.Error(loc, "was not found but should have been")
		}
	}
}

func TestSingleLineParse(t *testing.T) {
	lineData := waterflow.ParseLine("x=495, y=2..7")
	if len(lineData) != 1+7-2 {
		t.Errorf("Length of second parse should have been %d but was %d", 1+7-2, len(lineData))
	}
	for i, point := range lineData {
		if point.X != 495 {
			t.Error("X value should have been 495")
		}
		if point.Y != 2+i {
			t.Errorf("y value should have been %d was instead %d", 2+i, point.Y)
		}
	}
	lineData = waterflow.ParseLine("y=13, x=498..504")
	fmt.Println(lineData)
	if len(lineData) != 1+504-498 {
		t.Errorf("Length of second parse should have been %d but was %d", 1+504-498, len(lineData))
	}
	for i, point := range lineData {
		if point.Y != 13 {
			t.Error("X value should have been 13")
		}
		if point.X != 498+i {
			t.Errorf("y value should have been %d was instead %d", 498+i, point.Y)
		}
	}
}
