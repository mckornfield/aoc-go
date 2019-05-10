package waterflow

import (
	"errors"
	"fmt"
)

func DetermineNumberOfSquaresCovered(startingWaterPoint Location, maxY int, blockedPoints LocationSet) int {
	count := 0
	waterPoint := startingWaterPoint
	clayPoints := blockedPoints.Clone()
	waterPath := Stack{}
	for i := 0; i < 5000; i++ { // Replace with proper conditional
		newPoint, err := DetermineNextWaterPoint(waterPoint, blockedPoints, clayPoints, maxY)
		if err != nil {
			// Pop the stack
			stackVal, err2 := waterPath.Pop()
			if err2 != nil {
				fmt.Println("rewound stack")
				fmt.Println(waterPoint)
				return count // Have rewound stack completely
			}
			waterPoint = stackVal
			fmt.Println("popped point", waterPoint)
		} else {
			waterPath.Push(waterPoint) // Push last point
			waterPoint = newPoint
			blockedPoints[waterPoint] = true
			count++
			fmt.Println("moved to point", waterPoint)
		}
	}
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
func DetermineNextWaterPoint(currentWaterPoint Location, blockedPoints, clayPoints LocationSet, maxY int) (Location, error) {
	nextPoint, err := findNextOpenLocationDownLeftOrRight(currentWaterPoint, blockedPoints, clayPoints, maxY)
	if err != nil {
		return Location{}, err
	}
	return nextPoint, nil

}

func findNextOpenLocationDownLeftOrRight(currentWaterPoint Location, blockedPoints, clayPoints LocationSet, maxY int) (Location, error) {
	belowPoint := Location{X: currentWaterPoint.X, Y: currentWaterPoint.Y + 1}
	leftPoint := Location{X: currentWaterPoint.X - 1, Y: currentWaterPoint.Y}
	belowLeftPoint := Location{X: currentWaterPoint.X - 1, Y: currentWaterPoint.Y + 1}
	rightPoint := Location{X: currentWaterPoint.X + 1, Y: currentWaterPoint.Y}
	belowRightPoint := Location{X: currentWaterPoint.X + 1, Y: currentWaterPoint.Y + 1}
	switch {
	case !blockedPoints[belowPoint] && belowPoint.Y <= maxY:
		return belowPoint, nil
	case !blockedPoints[leftPoint] && (blockedPoints[belowLeftPoint] || clayPoints[belowPoint]):
		return leftPoint, nil
	case !blockedPoints[rightPoint] && (blockedPoints[belowRightPoint] || clayPoints[belowPoint]):
		return rightPoint, nil
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
