package board

import (
	"errors"
	"strings"
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
	// fmt.Println(alignedPlayers)
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

func (b BoardData) isPlayerDead(playerID int) bool {
	_, _, err := b.getIndexAndPlayer(playerID)
	return err != nil
}

func (b BoardData) printBoard() string {
	var bottomRightCorner Location = Loc{x: 0, y: 0}
	for location := range b.spaces {
		if location.getY() > bottomRightCorner.getY() {
			bottomRightCorner = location
		} else if location.getY() == bottomRightCorner.getY() &&
			location.getX() > bottomRightCorner.getX() {
			bottomRightCorner = location
		}
	}
	var sb strings.Builder
	locationToPlayer := make(map[Location]rune)
	for _, player := range b.allPlayers {
		var playerChar rune
		if player.alignment == ElfAlignment {
			playerChar = 'E'
		} else {
			playerChar = 'G'
		}
		locationToPlayer[player.toLocation()] = playerChar
	}

	for y := 0; y < bottomRightCorner.getY()+1; y++ {
		for x := 0; x < bottomRightCorner.getX()+1; x++ {
			loc := Loc{x: x, y: y}
			if b.spaces[loc] {
				sb.WriteRune('.')
			} else if playerChar, present := locationToPlayer[loc]; present {
				sb.WriteRune(playerChar)
			} else {
				sb.WriteRune('#')
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
