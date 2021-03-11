package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	cards := []int{5, 4, 2, 6, 8, 9, 10, 5, 7, 9, 3, 6, 9, 12, 14, 24, 34, 50}
	fmt.Println(pickRandomArr(cards, 5))

}

func pickRandomArr(original []int, m int) []int {
	subset := make([]int, m)

	//fill in subset array with first part of original array
	for i := 0; i < m; i++ {
		subset[i] = original[i]
	}

	s1 := rand.NewSource(time.Now().UnixNano()) //truly random
	r1 := rand.New(s1)

	//go through rest of original array
	for i := m; i < len(original); i++ {
		k := r1.Intn(m)
		if k < m {
			subset[k] = original[i]
		}
	}

	return subset
}
