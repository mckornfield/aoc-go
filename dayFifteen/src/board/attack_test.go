package board

import (
	"testing"
)

func TestAdversaryAttack(t *testing.T){
	elf := Player{
		id: 0,
		xLocation: 1,
		yLocation: 1,
		health: 10,
		alignment: ElfAlignment,
	}
	goblin := Player{
		id: 1,
		xLocation: 1,
		yLocation: 2,
		health: 10,
		alignment: GoblinAlignment,
	}
	goblin2 := Player{
		id: 2,
		xLocation: 0,
		yLocation: 1,
		health: 11,
		alignment: GoblinAlignment,
	}
	spaces := map[Location]Space{
		Loc{x:1,y:2}:false,
	}
	board := BoardData{
		allPlayers: Players{elf,goblin,goblin2},
		spaces: spaces,
	}
	determineAttackAndPerform(&board, elf)

	_, woundedPlayer, _ := board.getIndexAndPlayer(goblin.id)
	if woundedPlayer.health != 7 {
		t.Errorf("Goblin should have taken ten damage")
	}
	if spaces[Loc{x:1,y:2}]{
		t.Error("Space should NOT be open")
	}

	for i :=0; i < 4; i++ {
		determineAttackAndPerform(&board, elf)
	}

	_, woundedPlayer, err := board.getIndexAndPlayer(goblin.id)
	if err == nil {
		t.Errorf("Should have got an error")
	}

	if !spaces[Loc{x:1,y:2}]{
		t.Error("Space should have been open")
	}

}