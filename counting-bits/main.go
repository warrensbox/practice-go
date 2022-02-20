package main

import "fmt"

func main() {

	fmt.Println(countBits(5))

}

func countBits(n int) []int {

	res := []int{}
	for i := 0; i <= n; i++ {
		bit := i
		count := 0
		for bit != 0 {
			if bit&1 == 1 {
				count++
			}
			bit = bit >> 1
		}
		res = append(res, count)
	}

	return res

}
