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

		for i := 0; i < n; i++ {
			if IsCelebrity(i) {
				return i
			}
		}

		return -1
	}
}
