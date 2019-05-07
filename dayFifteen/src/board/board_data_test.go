package board

import(
	"testing"
	"strings"
)
func TestBoardPrinting(t *testing.T){
	file := "../input1.txt"
	board := Parse(file)
	content, err := getFileAsString(file)

	if err != nil {
		t.Error(err)
	}
	
	expectedBoardWithoutNewLines := strings.Replace(string(content),"\n","",-1)
	printedBoardWithoutNewLines := strings.Replace(board.printBoard(),"\n","",-1)
	if expectedBoardWithoutNewLines != printedBoardWithoutNewLines {
		t.Errorf("Board print did not match initial board")
		t.Errorf("Expected: \n%s", expectedBoardWithoutNewLines)
		t.Errorf("Actual: \n%s", printedBoardWithoutNewLines)
	}
}