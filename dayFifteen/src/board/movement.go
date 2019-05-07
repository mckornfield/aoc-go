package board

import "fmt"

func movePlayer(player Player, board BoardData) {
	oldLoc := player.toLocation()
	loc, err := determineMoveLocation(player, board)
	if err != nil {
		// Cannot move
	} else {
		// Update open spaces
		board.spaces[oldLoc] = true
		board.spaces[loc] = false
		playerIndex := getIndex(board.allPlayers, player.id)
		board.allPlayers[playerIndex].xLocation = loc.getX()
		board.allPlayers[playerIndex].yLocation = loc.getY()
	}
}

func getIndex(players Players, playerID int) int {
	for index, player := range players {
		if playerID == player.id {
			return index
		}
	}
	panic(fmt.Sprintf("Player %d was not found in %+v", playerID, players))
}

func determineMoveLocation(player Player, board BoardData) (Location, error) {

	spot, err := getClosestAdversarySpot(player, board)
	if err != nil {
		return Loc{}, err
	}
	// Look at board distances from other direction
	distancesFromTarget := buildBoardDistances(board, spot, player.toLocation())

	adjacentMap := make(map[Location]int)

	for _, locDelta := range []Loc{
		Loc{x: 0, y: 1}, Loc{x: 1, y: 0},
		Loc{x: 0, y: -1}, Loc{x: -1, y: 0},
	} {
		adjacentSpot := Loc{x: player.getX() + locDelta.getX(), y: player.getY() + locDelta.getY()}
		if dist, present := distancesFromTarget[adjacentSpot]; present {
			adjacentMap[adjacentSpot] = dist
		}
	}

	moveLoc, err := getShortestReadOrderDistance(adjacentMap)
	if err != nil {
		return Loc{}, err
	}
	return moveLoc, nil
}
