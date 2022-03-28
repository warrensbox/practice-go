package main

import "sort"

func numRescueBoats(people []int, limit int) int {

	count := 0
	left := 0
	right := len(people) - 1

	sort.Ints(people)
	for left <= right {
		count++
		if people[left]+people[right] <= limit {
			left++
		}
		right--
	}
	return count
}
