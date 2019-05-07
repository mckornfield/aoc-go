package board

import (
	"errors"
	"fmt"
)

const attackPower = 3

func determineAttackAndPerform(board *BoardData, player Player) {
	// Check nearby players of opposite alignment
	adversaries := board.getAdversaries(player.alignment)
	locToAdversary := make(map[Location]Player)
	for _, adversary := range adversaries {
		locToAdversary[adversary.toLocation()] = adversary
	}
	adversary, err := getAdjacentAdversary(player, locToAdversary)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Attack
	playerIndex := getIndex(board.allPlayers, adversary.id)
	oldHealth := board.allPlayers[playerIndex].health
	board.allPlayers[playerIndex].health = oldHealth - attackPower
	if board.allPlayers[playerIndex].health < 0 {
		// Remove player and update space
		loc := board.allPlayers[playerIndex].toLocation()
		board.allPlayers = append(board.allPlayers[:playerIndex], board.allPlayers[playerIndex+1:]...)
		board.spaces[loc] = true
	}

}

func getAdjacentAdversary(player Player, adversaryMap map[Location]Player) (Player, error) {
	var selectedAdversary Player
	noAdversaryFound := true
	for _, locDelta := range adjacentSpots {
		loc := Loc{x: player.xLocation + locDelta.x, y: player.yLocation + locDelta.y}
		if player, present := adversaryMap[loc]; present {
			if noAdversaryFound {
				selectedAdversary = player
				noAdversaryFound = false
			} else if selectedAdversary.health > player.health {
				selectedAdversary = player
			}
		}
	}
	fmt.Println(selectedAdversary)
	if noAdversaryFound {
		return Player{}, errors.New("No adversary found")
	}
	return selectedAdversary, nil
}
