package board

import (
	"errors"
	"sort"
)

var adjacentSpots = []Loc{
	Loc{x: 0, y: -1},
	Loc{x: -1, y: 0},
	Loc{x: 1, y: 0},
	Loc{x: 0, y: 1},
}

func getClosestAdversarySpot(player Player, board BoardData) (Location, error) {
	reachableSpots := getReachableSpots(player, board)
	if len(reachableSpots) == 0 {
		return Loc{}, errors.New("No reachable spots found")
	}
	return getShortestReadOrderDistance(reachableSpots)
}

func getShortestReadOrderDistance(reachableSpots map[Location]int) (Location, error) {
	// Get minimum distance, set as far initially
	minDist := 5000
	locs := Locations{}
	for spot, dist := range reachableSpots {
		if minDist > dist {
			minDist = dist
			locs = Locations{} // Clear out old locations
		}
		if minDist == dist {
			locs = append(locs, spot)
		}
	}
	if len(locs) == 0 {
		return Loc{}, errors.New("No reachable spots")
	}
	sort.Stable(locs) // Put in reading order

	return locs[0], nil
}

func getReachableSpots(player Player, board BoardData) map[Location]int {
	openSpots := getOpenSpots(player.alignment, board)

	boardSpotDistances := buildBoardDistances(board, player.toLocation(), player.toLocation())

	reachableSpots := make(map[Location]int)
	for _, openSpot := range openSpots {
		distance, present := boardSpotDistances[openSpot]
		if present {
			reachableSpots[openSpot] = distance
		}
	}
	return reachableSpots
}

func buildBoardDistances(board BoardData, loc Location, targetSpot Location) map[Location]int {
	// Perform maze movement algorithm
	spotDistance := make(map[Location]int)
	currentSpot := loc
	spotDistance[currentSpot] = 0
	checkAroundSpot(board, currentSpot, spotDistance, 0)
	return spotDistance
}

// Recursively move through spots, while updaating spot distances
func checkAroundSpot(board BoardData, loc Location, spotDistance map[Location]int, distance int) {
	surroundingSpots := findNewOpenSpots(board, loc, spotDistance, distance)
	// Visit each spot in order
	sort.Stable(surroundingSpots)
	if len(surroundingSpots) == 0 {
		return
	}
	for _, spot := range surroundingSpots {
		checkAroundSpot(board, spot, spotDistance, distance+1)
	}
}

func getOpenSpots(alignment int, board BoardData) Locations {
	adversaries := board.getAdversaries(alignment)

	openSpots := Locations{}
	spotSet := make(map[Location]int)
	for _, adversary := range adversaries {
		newSpots := findNewOpenSpots(board, adversary.toLocation(), spotSet, 1)
		if len(newSpots) != 0 {
			openSpots = append(openSpots, newSpots...)
		}
	}
	return openSpots
}

func findNewOpenSpots(board BoardData, currentLoc Location,
	spotDist map[Location]int, distance int) Locations {

	xLoc := currentLoc.getX()
	yLoc := currentLoc.getY()
	openSpots := Locations{}
	for _, locDelta := range adjacentSpots {
		loc := Loc{x: xLoc + locDelta.getX(), y: yLoc + locDelta.getY()}
		if open, present := board.spaces[loc]; present && bool(open) {
			updateOpenSpotsIfNotInSet(spotDist, &openSpots, loc, distance)
		}
	}
	return openSpots
}

func updateOpenSpotsIfNotInSet(spotDist map[Location]int,
	openSpots *Locations, loc Loc, distFromPoint int) {
	if dist, present := spotDist[loc]; !present || dist > distFromPoint {
		spotDist[loc] = distFromPoint
		*openSpots = append(*openSpots, loc)
	}
}
