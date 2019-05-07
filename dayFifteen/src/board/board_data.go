package board

import (
	"errors"
	"fmt"
)

// BoardData the board with elves, goblins and spaces/obstacles
type BoardData struct {
	allPlayers Players
	spaces     map[Location]Space
}

func (b BoardData) getAdversaries(alignment int) Players {
	if alignment == ElfAlignment {
		return b.getGoblins()
	}
	return b.getElves()
}

func (b BoardData) getElves() Players {
	return getPlayers(b.allPlayers, ElfAlignment)
}

func (b BoardData) getGoblins() Players {
	return getPlayers(b.allPlayers, GoblinAlignment)
}

func getPlayers(pls Players, alignment int) Players {
	alignedPlayers := Players{}
	for _, pl := range pls {
		if pl.alignment == alignment {
			alignedPlayers = append(alignedPlayers, pl)
		}
	}
	fmt.Println(alignedPlayers)
	return alignedPlayers
}

func (b BoardData) getIndexAndPlayer(playerID int) (int, Player, error) {
	for index, player := range b.allPlayers {
		if playerID == player.id {
			return index, player, nil
		}
	}
	return 0, Player{}, errors.New("No player found")
}
