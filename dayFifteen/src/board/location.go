package board

// Locations a slice of x y coordinates
type Locations []Location

// Location interface to define an x y location
type Location interface {
	getX() int
	getY() int
}

// Loc simple location
type Loc struct {
	x int
	y int
}

func (l Loc) getX() int {
	return l.x
}

func (l Loc) getY() int {
	return l.y
}

// Len number of players
func (locs Locations) Len() int { return len(locs) }

func (locs Locations) Less(i, j int) bool {
	loc1 := locs[i]
	loc2 := locs[j]
	return isLess(loc1, loc2)
}

func (locs Locations) Swap(i, j int) {
	locs[i], locs[j] = locs[j], locs[i]
}

func isLess(loc1, loc2 Location) bool {
	if loc1.getY() < loc2.getY() {
		return true
	} else if loc1.getY() == loc2.getY() {
		return loc1.getX() < loc2.getX()
	} else {
		return false
	}
}
