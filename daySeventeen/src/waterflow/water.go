package waterflow

import "errors"

func DetermineNumberOfSquaresCovered(startingWaterPoint Location, blockedPoints LocationSet) int {
	count := 0

	return count
}

/*
 There are a number of explicit behaviors that matter
 First and foremost, water always travels down first if the location below is open
 If the bottom most point is occupied, it starts to fill
 Left first, then right next if possible
 Once water has filled left and right, it can cover itself
 Water will overflow over the side, and then start filling again
 Water beyond the maximum Y coordinate of any clay does not count
*/
func DetermineNextWaterPoint(currentWaterPoint Location, blockedPoints LocationSet) (Location, error) {
	nextPoint, err := findNextOpenLocationDownLeftOrRight(currentWaterPoint, blockedPoints)
	if err != nil {
		return Location{}, nil
	}
	blockedPoints[nextPoint] = true
	return nextPoint, nil

}

func findNextOpenLocationDownLeftOrRight(currentWaterPoint Location, blockedPoints LocationSet) (Location, error) {
	belowPoint := Location{X: currentWaterPoint.X, Y: currentWaterPoint.Y + 1}
	leftPoint := Location{X: currentWaterPoint.X - 1, Y: currentWaterPoint.Y}
	rightPoint := Location{X: currentWaterPoint.X + 1, Y: currentWaterPoint.Y}
	for _, point := range []Location{belowPoint, leftPoint, rightPoint} {
		if !blockedPoints[point] {
			return point, nil
		}
	}
	return Location{}, errors.New("No available points")
}

// Stack because golang stinks
type Stack []Location

func (s *Stack) Push(loc Location) {
	*s = append(*s, loc)
}

func (s *Stack) Pop() (Location, error) {
	length := len(*s)
	if length == 0 {
		return Location{}, errors.New("Cannot pop empty stack")
	}
	val := (*s)[length-1]
	*s = (*s)[:length-1]
	return val, nil
}

func (s *Stack) Peek() (Location, error) {
	length := len(*s)
	if length == 0 {
		return Location{}, errors.New("Cannot peek empty stack")
	}
	val := (*s)[length-1]
	return val, nil
}
