package board

import (
	"fmt"
	"sort"
)

// TryPowerLevels attempt different power levels up to max
func TryPowerLevels(boardLocation string, maxNumRounds, maxPowerLevel int) int {
	for power := 4; power < maxPowerLevel; power++ {
		worked := RunThroughGameWithPowerAmount(boardLocation, maxNumRounds, false, power)
		if worked {
			return power
		}
	}
	return -1
}

// RunThroughGameWithPowerAmount is similar to run through game, but instead tries a different power level
func RunThroughGameWithPowerAmount(boardLocation string, maxNumRounds int, shouldLog bool, elfPowerLevel int) bool {
	board := Parse(boardLocation)
	gameOver := false
	rounds := 0
	// Update power levels
	for index, player := range board.allPlayers {
		if player.alignment == ElfAlignment {
			board.allPlayers[index].attackPower = elfPowerLevel
		}
	}
	for i := 0; i < 10000; i++ {
		sort.Stable(board.allPlayers)
		roundOver := false
		ids := make([]int, len(board.allPlayers))
		for index, player := range board.allPlayers {
			ids[index] = player.id
		}
		// fmt.Println(ids)
		for _, id := range ids {
			_, player, err := board.getIndexAndPlayer(id)
			if err != nil {
				continue
			}
			player = movePlayer(player, board)
			playerDied, playerAlignment := determineAttackAndPerform(&board, player)
			i++
			if playerDied && playerAlignment == ElfAlignment {
				// Power level failed
				fmt.Println("Power level", elfPowerLevel, "failed")
				return false
			}
			if playerDied && (len(board.getGoblins()) == 0 || len(board.getElves()) == 0) {
				fmt.Println("Game ended!")
				gameOver = true
				if id == ids[len(ids)-1] {
					fmt.Println("Round over")
					rounds++
				}
				break
			}
			roundOver = true
		}
		if gameOver {
			break
		}
		if roundOver {
			rounds++
		}

		if shouldLog {
			fmt.Println(rounds, board.allPlayers)
			fmt.Println(board.printBoard())
		}

		if rounds >= maxNumRounds {
			break
		}
	}
	// Get health of remaining players
	health := 0
	for _, player := range board.allPlayers {
		health += player.health
	}

	fmt.Println("Number of rounds:", rounds, "| Number of HP left:", health)
	fmt.Println("Outcome:", rounds*health)
	fmt.Println("Power level", elfPowerLevel, "worked")
	return true
}
