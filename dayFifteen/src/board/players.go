package board

import "fmt"

// Players on a board, used for sorting
type Players []Player

func (pl Players) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func (pl Players) Len() int { return len(pl) }

func (pl Players) Less(i, j int) bool {
	pl1 := pl[i]
	pl2 := pl[j]
	return isLess(pl1, pl2)
}

// Player either elf or goblin
type Player struct {
	id          int
	health      int
	alignment   int
	xLocation   int
	yLocation   int
	attackPower int
}

func (p Player) getX() int {
	return p.xLocation
}

func (p Player) String() string {
	var alignmentStr string
	if p.alignment == ElfAlignment {
		alignmentStr = "Elf"
	} else {
		alignmentStr = "Goblin"
	}
	return fmt.Sprintf("{id: %d, health: %d, alignment: %s, x: %d, y: %d}", p.id, p.health, alignmentStr, p.xLocation, p.yLocation)
}

func (p Player) getY() int {
	return p.yLocation
}

func (p Player) toLocation() Loc {
	return Loc{x: p.getX(), y: p.getY()}
}

func (pl Players) toLocations() Locations {
	locations := Locations{}
	for _, player := range pl {
		locations = append(locations, player.toLocation())
	}
	return locations
}

// Space on the board
type Space bool
