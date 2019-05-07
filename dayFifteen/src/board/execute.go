package board

import (
	"fmt"
	"sort"
)

// RunThroughGame moves players, attacks if necessary, and exits once
// all players on one team are dead.
// Hehe this is garbage
func RunThroughGame(boardLocation string, maxNumRounds int, shouldLog bool) int {
	board := Parse(boardLocation)
	gameOver := false
	rounds := 0
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
			playerDied, _ := determineAttackAndPerform(&board, player)
			i++
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

	return health * rounds
}
