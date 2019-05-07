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

func TestInput4Board_47Rounds(t *testing.T){
	result := RunThroughGame("../input4.txt",47,false)
	if result != 47 * (200+131+59+200) {
		t.Errorf("Result should have been %d but was insteaad %d",47 * (200+131+59+200),result)
	}
}

func TestInput5Board_37Rounds(t *testing.T){
	result := RunThroughGame("../input5.txt",47,false)
	if result != 36334 {
		t.Errorf("Result should have been %d but was insteaad %d",36334,result)
	}
}

func TestInput6Board_46Rounds(t *testing.T) {
	result := RunThroughGame("../input6.txt",47,false)
	if result != 39514 {
		t.Errorf("Result should have been %d but was insteaad %d",39514,result)
	}
}


func TestInput7Board_35Rounds(t *testing.T) {
	result := RunThroughGame("../input7.txt",47,false)
	if result != 27755 {
		t.Errorf("Result should have been %d but was insteaad %d",27755,result)
	}
}


func TestInput8Board_54Rounds(t *testing.T) {
	result := RunThroughGame("../input8.txt",55,false)
	if result != 28944 {
		t.Errorf("Result should have been %d but was insteaad %d",28944,result)
	}
}


func TestInput9Board_20Rounds(t *testing.T) {
	result := RunThroughGame("../input9.txt",55,false)
	if result != 18740 {
		t.Errorf("Result should have been %d but was insteaad %d",18740,result)
	}
}


func TestInputPuzz_1000Rounds(t *testing.T) {
	result := RunThroughGame("../puzz-1-input.txt",1000,false)
	if result != 183300 {
		t.Errorf("Result should have been %d but was insteaad %d",183300,result)
	}
}
