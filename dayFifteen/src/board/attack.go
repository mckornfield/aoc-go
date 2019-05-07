package board

import (
	"errors"
)

const attackPower = 3

func determineAttackAndPerform(board *BoardData, player Player) (bool, int) {
	// Check nearby players of opposite alignment
	adversaries := board.getAdversaries(player.alignment)
	locToAdversary := make(map[Location]Player)
	for _, adversary := range adversaries {
		locToAdversary[adversary.toLocation()] = adversary
	}
	adversary, err := getAdjacentAdversary(*board, player)
	if err != nil {
		// fmt.Println(err)
		return false, -1
	}
	// Attack
	playerIndex, _, _ := board.getIndexAndPlayer(adversary.id)
	oldHealth := board.allPlayers[playerIndex].health
	board.allPlayers[playerIndex].health = oldHealth - player.attackPower
	// fmt.Println("Attacking player", player)
	// fmt.Println("Attacked player", board.allPlayers[playerIndex])
	if board.allPlayers[playerIndex].health < 0 {
		// Remove player and update space
		loc := board.allPlayers[playerIndex].toLocation()
		alignment := board.allPlayers[playerIndex].alignment
		board.allPlayers = append(board.allPlayers[:playerIndex], board.allPlayers[playerIndex+1:]...)
		board.spaces[loc] = true
		return true, alignment
	}
	return false, -1
}

func getAdjacentAdversary(board BoardData, player Player) (Player, error) {
	adversaries := board.getAdversaries(player.alignment)
	locToAdversary := make(map[Location]Player)
	for _, adversary := range adversaries {
		locToAdversary[adversary.toLocation()] = adversary
	}

	var selectedAdversary Player
	noAdversaryFound := true
	for _, locDelta := range adjacentSpots {
		loc := Loc{x: player.xLocation + locDelta.x, y: player.yLocation + locDelta.y}
		if player, present := locToAdversary[loc]; present {
			if noAdversaryFound {
				selectedAdversary = player
				noAdversaryFound = false
			} else if selectedAdversary.health > player.health {
				selectedAdversary = player
			}
		}
	}
	// fmt.Println(selectedAdversary)
	if noAdversaryFound {
		return Player{}, errors.New("No adversary found")
	}
	return selectedAdversary, nil
}
