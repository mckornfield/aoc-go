package board

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
	health    int
	alignment int
	xLocation int
	yLocation int
}

func (p Player) getX() int {
	return p.xLocation
}

func (p Player) getY() int {
	return p.yLocation
}

func (pl Players) toLocations() Locations {
	locations := Locations{}
	for _, player := range pl {
		locations = append(locations, Loc{x: player.getX(), y: player.getY()})
	}
	return locations
}

// Space on the board
type Space bool

func (p Player) getLocation() (int, int) {
	return p.xLocation, p.yLocation
}
