package main

import (
	"fmt"
	"time"
)

func main() {

	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5

	output := shipWithinDays(weights, D)

	fmt.Println(output)

}

func shipWithinDays(weights []int, D int) int {

	lo := maximumOfWeights(weights)
	hi := sum(weights)
	fmt.Println(lo)
	fmt.Println(hi)

	for lo < hi {
		fmt.Println("-----")
		// fmt.Println(lo)
		// fmt.Println(hi)

		days := 1
		load := 0
		capacity := lo + (hi-lo)/2
		fmt.Println("capacity", capacity)
		for i := 0; i < len(weights); i++ {
			fmt.Println("=====")
			fmt.Println("weights[i]", weights[i])
			fmt.Println("load", load)
			if weights[i]+load > capacity {
				load = weights[i]
				days++
				fmt.Println("daysss", days)
			} else {
				fmt.Println("w else")
				load += weights[i]
				fmt.Println("LOAD", load)
			}
			time.Sleep(8 * time.Second)
		}

		fmt.Println("days", days)
		if days <= D {
			hi = capacity
		} else {
			lo = capacity + 1
		}
		time.Sleep(3 * time.Second)
	}

	return lo
}

func sum(weights []int) int {

	sum := 0

	for i := 1; i < len(weights); i++ {
		sum += weights[i]
	}

	return sum
}

func maximumOfWeights(weights []int) int {

	max := weights[0]

	for i := 1; i < len(weights); i++ {
		if weights[i] > max {
			max = weights[i]
		}
	}

	return max
}
