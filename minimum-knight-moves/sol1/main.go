package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 5
	y := 5

	output := minKnightMoves(x, y)
	fmt.Println(output)
}

func minKnightMoves(x int, y int) int {

	moves := [][]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})

	var minMoves int
	seen := make(map[string]bool)

	for len(queue) > 0 {
		totalPos := len(queue) //total positions in a single hop
		for i := 0; i < totalPos; i++ {
			currPos := (queue)[0]
			queue = (queue)[1:]
			if currPos[0] == x && currPos[1] == y {
				return minMoves
			} else {
				for _, move := range moves {
					newPosX := currPos[0] + move[0]
					newPosY := currPos[1] + move[1]

					if !seen[convertToStr(newPosX, newPosY)] {
						seen[convertToStr(newPosX, newPosY)] = true
						queue = append(queue, []int{newPosX, newPosY})
					}
				}
			}
		}
		minMoves++
	}

	return minMoves
}

func convertToStr(newPosX, newPosY int) string {
	return strconv.Itoa(newPosX) + strconv.Itoa(newPosY)
}
