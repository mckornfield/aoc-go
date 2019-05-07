package board

func movePlayer(player Player, board BoardData) Player {
	oldLoc := player.toLocation()
	_, adjErr := getAdjacentAdversary(board, player)
	if adjErr == nil {
		return player // Already adjacent to enemy
	}
	loc, moveErr := determineMoveLocation(player, board)
	playerIndex, _, playerErr := board.getIndexAndPlayer(player.id)
	if moveErr != nil || playerErr != nil {
		// Cannot move
	} else {
		// Update open spaces
		board.spaces[oldLoc], board.spaces[loc] = board.spaces[loc], board.spaces[oldLoc]
		board.allPlayers[playerIndex].xLocation = loc.getX()
		board.allPlayers[playerIndex].yLocation = loc.getY()
	}
	return board.allPlayers[playerIndex]
}

func determineMoveLocation(player Player, board BoardData) (Location, error) {

	spot, err := getClosestAdversarySpot(player, board)
	if err != nil {
		// fmt.Println(err)
		return Loc{}, err
	}
	// Look at board distances from other direction
	distancesFromTarget := buildBoardDistances(board, spot, player.toLocation())

	adjacentMap := make(map[Location]int)

	for _, locDelta := range adjacentSpots {
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
