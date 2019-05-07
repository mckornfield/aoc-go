package board

import(
	"testing"
)
func TestInput4Board_2Rounds(t *testing.T){
	result := RunThroughGame("../input4.txt",2,false)
	if result != 2 * 1170 {
		t.Errorf("Result should have been %d but was insteaad %d",1170,result)
	}
}

func TestInput4Board_28Rounds(t *testing.T){
	result := RunThroughGame("../input4.txt",28,false)
	if result != 28 * (200 + 131 + 116 + 113 + 200) {
		t.Errorf("Result should have been %d but was insteaad %d",28 * (200 + 131 + 116 + 113 + 200),result)
	}
}

func TestInput5Board_47Rounds(t *testing.T){
	result := RunThroughGame("../input4.txt",47,false)
	if result != 47 * (200+131+59+200) {
		t.Errorf("Result should have been %d but was insteaad %d",47 * (200+131+59+200),result)
	}
}

func TestInput4Board_47Rounds(t *testing.T){
	result := RunThroughGame("../input5.txt",47,false)
	if result != 36334 {
		t.Errorf("Result should have been %d but was insteaad %d",36334,result)
	}
}
