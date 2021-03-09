package main

import "fmt"

func main() {

	solution := []int{1, 2, 4, 3}
	guess := []int{1, 2, 3, 4}
	va := estimate(guess, solution)
	fmt.Println(va)
}

func estimate(guess []int, solution []int) string {
	if len(guess) != len(solution) {
		return "incomplete"
	}

	freq := make(map[int]int)
	hit := 0
	pseudo := 0
	//compute hits and build frequency table
	for i := 0; i < len(guess); i++ {
		if guess[i] == solution[i] {
			hit++
		} else {
			freq[solution[i]]++
		}
	}

	//compute pseudo-hits
	for i := 0; i < len(guess); i++ {
		if freq[guess[i]] > 0 && guess[i] != solution[i] {
			pseudo++
			freq[guess[i]]--
		}
	}

	return fmt.Sprintf("%v hits and %v pseudo-hits", hit, pseudo)
}
