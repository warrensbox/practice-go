package main

import "math"

func maxNumberOfBalloons(text string) int {

	pattern := "balloon"

	freqInPattern := make([]int, 26) //remember the character
	for _, t := range pattern {
		freqInPattern[t-'a']++
	}

	freqInText := make([]int, 26) //remember the character
	for _, t := range text {
		freqInText[t-'a']++
	}

	ans := math.MaxInt32
	for i := 0; i < 26; i++ {
		if freqInPattern[i] > 0 {
			ans = Min(ans, freqInText[i]/freqInPattern[i])
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func Min(x, y int) int {

	if x < y {
		return x
	}

	return y
}
