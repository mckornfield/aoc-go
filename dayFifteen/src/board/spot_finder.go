package board

import "fmt"

func getOpenSpots(board BoardData, alignment int) Locations {
	var adversaries Players
	if alignment == ElfAlignment {
		fmt.Println("adversaries are goblins")
		adversaries = board.goblins
	} else {
		fmt.Println("adversaries are elves")
		adversaries = board.elves
	}
	fmt.Println(adversaries)

	openSpots := Locations{}
	xMax := len(board.spaces[0]) - 1
	yMax := len(board.spaces) - 1
	fmt.Println("XMax:", xMax, "yMax:", yMax)
	for _, adversary := range adversaries {
		// Get spaces around adversary
		xLoc := adversary.xLocation
		yLoc := adversary.yLocation
		if xLoc != 0 && board.spaces[yLoc][xLoc-1] {
			// Check left
			openSpots = append(openSpots, Loc{x: xLoc - 1, y: yLoc})
		}
		if xLoc != xMax && board.spaces[yLoc][xLoc+1] {
			// Check right
			openSpots = append(openSpots, Loc{x: xLoc + 1, y: yLoc})
		}
		if yLoc != yMax && board.spaces[yLoc+1][xLoc] {
			// Check down +1
			openSpots = append(openSpots, Loc{x: xLoc, y: yLoc + 1})
		}
		if yLoc != 0 && board.spaces[yLoc-1][xLoc] {
			// Check up -1
			openSpots = append(openSpots, Loc{x: xLoc, y: yLoc - 1})
		}

	}
	return openSpots
}
