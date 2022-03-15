package main

func main() {

}

func maxScoreSightseeingPair(values []int) int {

	ans := 0
	best := values[0]

	for j := 1; j < len(values); j++ {
		ans = Max(ans, best+values[j]-j)
		best = Max(best, values[j]+j)
	}

	return ans

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
