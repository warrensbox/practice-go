package main

import "fmt"

func maxScore(cardPoints []int, k int) int {

	n := len(cardPoints)

	frontSetOfCard := make([]int, k+1)
	rearSetOfCard := make([]int, k+1)

	for i := 0; i < k; i++ {
		frontSetOfCard[i+1] = cardPoints[i] + frontSetOfCard[i]
		rearSetOfCard[i+1] = cardPoints[n-i-1] + rearSetOfCard[i]
	}

	fmt.Println(frontSetOfCard)
	fmt.Println(rearSetOfCard)

	maxScore := 0

	//i represents the number of cards we take from the front
	for i := 0; i <= k; i++ {
		currentScore := frontSetOfCard[i] + rearSetOfCard[k-i]
		fmt.Println(currentScore)
		maxScore = Max(maxScore, currentScore)
	}

	return maxScore
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
