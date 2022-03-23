package main

func minCost(colors string, neededTime []int) int {

	maxsum := 0
	length := len(neededTime)
	for i := 0; i < length; i++ {
		max := neededTime[i]
		for i < length-1 && colors[i] == colors[i+1] {
			max = Max(max, Max(neededTime[i], neededTime[i+1]))
			i++
		}
		maxsum += max
	}

	sum := 0
	for i := 0; i < length; i++ {
		sum += neededTime[i]
	}

	return sum - maxsum
}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
