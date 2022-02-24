package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {

	return func(n int) int {

		var IsCelebrity func(n int) bool

		IsCelebrity = func(i int) bool {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if knows(i, j) || !knows(j, i) {
					return false
				}
			}
			return true
		}

		//find celebrity
		celebrityCandidate := 0
		for i := 0; i < n; i++ {
			if knows(celebrityCandidate, i) {
				celebrityCandidate = i
			}
		}

		if IsCelebrity(celebrityCandidate) {
			return celebrityCandidate
		}

		return -1
	}
}
