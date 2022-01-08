package main

import (
	"fmt"
)

/*

goal: verify meal served in order (FIFO) between take out and dine in
sol :
- start from the back of served array
- minus takeout idx and dinein idx from the back

*/

func main() {

	// takeout := []int{17, 8, 24}
	// dinein := []int{12, 19, 2}
	// served := []int{17, 8, 12, 19, 24, 2}

	takeout := []int{1, 3, 5}
	dinein := []int{2, 4, 6}
	served := []int{1, 2, 4, 6, 5, 3}

	fmt.Println(fifo(takeout, dinein, served))

}

func fifo(takeout, dinein, served []int) bool {

	idxTakeout := len(takeout) - 1
	idxDinein := len(dinein) - 1

	for i := len(served) - 1; i >= 0; i-- {

		if idxTakeout >= 0 && takeout[idxTakeout] == served[i] {
			idxTakeout--
		} else if idxDinein >= 0 && dinein[idxDinein] == served[i] {
			idxDinein--
		} else {
			return false
		}
	}

	if idxTakeout != 0 || idxDinein != 0 {
		return false
	}

	return true
}
