package main

import "fmt"

func numOfSubarrays(arr []int) int {
	summary := make([]int, 0, len(arr))
	current := 0
	for _, v := range arr {
		current = (current + v) % 2
		summary = append(summary, current)
	}

	fmt.Println(summary)
	odd, even := 0, 0
	number := 0
	for _, v := range summary {
		fmt.Println("v", v)
		if v%2 == 0 {
			even++
			number += odd
		} else {
			odd++
			number += even + 1
		}
		number %= 1000000007
	}
	return number
}
