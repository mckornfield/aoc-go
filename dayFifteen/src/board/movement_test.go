package board

import (
	// "reflect"
	// "fmt"
	"testing"
)
func TestMovementOfElf(t *testing.T){
	board := Parse("../input3.txt")
	elf := board.getElves()[0]
	moveLoc, err := determineMoveLocation(elf, board)
	if err != nil {
		t.Error(err)
	}
	location := Loc{x:3, y:1}
	if moveLoc != location {
		t.Errorf("Closest spot should have been 1,3, was %+v", moveLoc)
	}
}

func TestMovingElf(t *testing.T){
	board := Parse("../input3.txt")
	elf := board.getElves()[0]
	oldLoc := elf.toLocation()
	movePlayer(elf, board)
	elf = board.getElves()[0]

	if !board.spaces[oldLoc]{
		t.Errorf("Old Location should have been open, %+v", oldLoc)
	}

	location := Loc{x:3, y:1}
	if location != elf.toLocation() {
		t.Errorf("Location should have been 3,1, was %+v", elf.toLocation())
	}
	elf = board.allPlayers[0]

	location = Loc{x:3, y:1}
	if location != elf.toLocation() {
		t.Errorf("Location spot should have been 3,1, was %+v", elf.toLocation())
	}

	oldLoc = elf.toLocation()
	movePlayer(elf, board)

	if !board.spaces[oldLoc]{
		t.Errorf("Old Location should have been open, %+v", oldLoc)
	}

	elf = board.getElves()[0]
	location = Loc{x:4, y:1}
	if location != elf.toLocation() {
		t.Errorf("Location spot should have been 4,1, was %+v", elf.toLocation())
	}

	
}