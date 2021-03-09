package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	fmt.Println(sortSub(arr))
}

func sortSub(orig []int) (int, int) {
	//var flag bool
	min := math.MaxInt16
	max := math.MinInt16

	for i := 1; i < len(orig)-1; i++ {
		if orig[i] < orig[i-1] {
			// 	flag = true
			// }
			// if flag {
			min = Min(min, orig[i])
		}
	}
	//flag = false //reset flag
	for j := len(orig) - 2; j >= 0; j-- {
		if orig[j] > orig[j+1] {
			fmt.Println(orig[j])
			// 	flag = true
			// }
			// if flag {
			max = Max(max, orig[j])
		}
	}

	var l, r int

	for l = 0; l < len(orig); l++ {
		if min < orig[l] {
			break
		}
	}

	for r = len(orig) - 1; r >= 0; r-- {
		if max > orig[r] {
			break
		}
	}

	return l, r
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
