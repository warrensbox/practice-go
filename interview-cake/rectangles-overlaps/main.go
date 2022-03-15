package main

import "fmt"

/*
goal : find overllaping rectangles

sol : check for all 4 cases
1) The ranges partially overlap:

Two horizontal parallel lines. The right half of the top line overlaps the left half of the bottom line.

2) One range is completely contained in the other:

Two horizontal parallel lines. The bottom line is longer than the top line and extends farther out to the left and right.

3) The ranges don't overlap:

Two horizontal parallel lines. The right end of the bottom line is to the left of the left end of the top line.

4) The ranges "touch" at a single point:


*/

func main() {

	var rec1 Rectangle
	rec1.leftX = 1
	rec1.width = 5
	rec1.bottomY = 1
	rec1.height = 5

	var rec2 Rectangle
	rec2.leftX = 2
	rec2.width = 6
	rec2.bottomY = 2
	rec2.height = 6

	rec3 := findRectangularOverlap(rec1, rec2)

	fmt.Println(rec3)
}

type RangeOverLap struct {
	startPoint int
	length     int
}

func findRangeOverlap(point1 int, length1 int, point2 int, length2 int) *RangeOverLap {

	// find the highest start point and lowest end point.
	// the highest ("rightmost" or "upmost") start point is
	// the start point of the overlap.
	// the lowest end point is the end point of the overlap.
	highestStartPoint := Max(point1, point2)
	lowestEndPoint := Min(point1+length1, point2+length2)

	// return empty overlap if there is no overlap
	if highestStartPoint >= lowestEndPoint {
		fmt.Println("going there")
		return &RangeOverLap{0, 0}
	}

	// compute the overlap length
	overlapLength := lowestEndPoint - highestStartPoint

	return &RangeOverLap{highestStartPoint, overlapLength}
}

func findRectangularOverlap(rect1 Rectangle, rect2 Rectangle) *Rectangle {

	// get the x and y overlap points and lengths
	xOverlap := findRangeOverlap(rect1.leftX, rect1.width, rect2.leftX, rect2.width)
	yOverlap := findRangeOverlap(rect1.bottomY, rect1.height, rect2.bottomY, rect2.height)

	// return "zero" rectangle if there is no overlap
	if xOverlap.length == 0 || yOverlap.length == 0 {
		return &Rectangle{}
	}

	return &Rectangle{
		leftX:   xOverlap.startPoint,
		bottomY: yOverlap.startPoint,
		width:   xOverlap.length,
		height:  yOverlap.length,
	}
}

type Rectangle struct {
	leftX   int
	bottomY int
	width   int
	height  int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
