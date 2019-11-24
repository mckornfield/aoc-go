package lumber

import (
	"reflect"
	"testing"
)

func TestGetAnswer(t *testing.T) {
	answer, err := GetAnswer("sampleFile.txt", 10)
	if err != nil {
		t.Error("There should have been no error")
	}
	if answer != 1147 {
		t.Error("Answer was", answer, "Should have been", 1147)
	}
}

func TestGetAnswerPartOne(t *testing.T) {
	answer, err := GetAnswer("part_1_input.txt", 10)
	if err != nil {
		t.Error("There should have been no error")
	}
	if answer != 467819 {
		t.Error("Answer was", answer, "Should have been", 1147)
	}
}

func TestGetAnswerPartTwo(t *testing.T) {
	answer, err := GetAnswer("part_1_input.txt", 1000000000)
	if err != nil {
		t.Error("There should have been no error")
	}
	if answer != 467819 {
		t.Error("Answer was", answer, "Should have been", 1147)
	}
}

func TestGetAnswer_BadFile(t *testing.T) {
	_, err := GetAnswer("sampleFile_does_not_exist", 10)
	if err == nil {
		t.Error("Expected file not found error")
	}
}

func TestFileContentsToGrid(t *testing.T) {
	grid, err := fileContentsToGrid("sampleFile.txt")
	if err != nil {
		t.Fatal("error occurred")
	}
	if len(grid) == 0 {
		t.Fatal("Length was 0")
	}
	if grid[0][0] != '.' {
		t.Error("0,0 was wrong")
	}

	if grid[0][1] != '#' {
		t.Error("0,1 was wrong")
	}

	if grid[9][3] != '#' {
		t.Error("9,3 was wrong")
	}

	if grid[9][9] != '.' {
		t.Error("9,9 was wrong")
	}

	if grid[9][8] != '|' {
		t.Error("9,8 was wrong")
	}
}

func TestRoleForwardOne(t *testing.T) {
	grid, err := fileContentsToGrid("sampleFile.txt")
	if err != nil {
		t.Fatal("error occurred")
	}
	actualGrid := grid.roleForwardOne()

	expectedGrid, err := fileContentsToGrid("sampleFile_oneMinute.txt")
	if err != nil {
		t.Fatal("error occurred")
	}
	compareGrids(expectedGrid, actualGrid, t)
}

func compareGrids(expectedGrid, actualGrid Grid, t *testing.T) {
	for index, expectedRow := range actualGrid {
		actualRow := actualGrid[index]
		if !reflect.DeepEqual(expectedRow, actualRow) {
			t.Error("Row #", index, " not equal, expected=", expectedRow, "actual=", actualRow)
		}
	}
	if !reflect.DeepEqual(expectedGrid, actualGrid) {
		t.Fatal("grids not equal, expected=", expectedGrid, "actual=", actualGrid)
	}
}

func TestRoleForwardTenTimes(t *testing.T) {
	grid, err := fileContentsToGrid("sampleFile.txt")
	if err != nil {
		t.Fatal("error occurred")
	}
	actualGrid := grid.roleForwardNTimes(int64(10))
	expectedGrid, err := fileContentsToGrid("sampleFile_tenMinute.txt")
	if err != nil {
		t.Fatal("error occurred")
	}
	compareGrids(expectedGrid, actualGrid, t)

}
