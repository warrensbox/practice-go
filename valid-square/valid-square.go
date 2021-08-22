package main

import (
	"fmt"
	"sort"
)

func main() {

}
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	hash := make(map[int]int)
	listPoints := []int{
		getLength(p1, p2),
		getLength(p1, p3),
		getLength(p1, p4),
		getLength(p2, p3),
		getLength(p2, p4),
		getLength(p3, p4),
	}
	sort.Ints(listPoints)
	for i := 0; i < len(listPoints); i++ {
		fmt.Println(listPoints[i])

		if listPoints[i] != 0 {
			hash[listPoints[i]]++
		} else {
			return false
		}

	}

	if len(hash) != 2 {
		return false
	}

	fmt.Println("len(hash)", len(hash))
	for _, val := range hash {
		fmt.Println(val)
		return val == 2 || val == 4
	}

	fmt.Println("going here")
	return false
}

func getLength(a []int, b []int) int {
	return ((a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1]))
}
