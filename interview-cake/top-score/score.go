package main

import "fmt"

/*
goal : sort in O(n)
sol :
- can use O(1) in this case n = 100
- use values as indices

*/
func main() {

	unsortedScores := []int{37, 89, 41, 65, 91, 53}
	highestScore := 100
	fmt.Println(arangeScores(unsortedScores, highestScore))
}

func arangeScores(unsortedScores []int, highestScore int) []int {

	arranged := make([]int, highestScore+1) //plus one because index starts at 0

	for _, val := range unsortedScores {
		arranged[val]++
	}

	var sorted []int
	for score := highestScore; score >= 0; score-- {

		freq := arranged[score]

		if freq != 0 {
			for occurances := 0; occurances < freq; occurances++ {
				sorted = append(sorted, score)
			}
		}
	}

	return sorted
}
