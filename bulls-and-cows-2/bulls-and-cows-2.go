package main

import "fmt"

func getHint(secret string, guess string) string {

	tmp := make([]int, 10)
	cow := 0
	bull := 0
	for i := 0; i < len(guess); i++ {

		g := guess[i] - '0'
		s := secret[i] - '0'

		if g == s {
			bull++
			continue
		}

		if tmp[s] < 0 {
			cow++
		}
		if tmp[g] > 0 {
			cow++
		}
		tmp[s]++
		tmp[g]--
	}

	return fmt.Sprintf("%vA%vB", bull, cow)
}
