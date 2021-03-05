package main

import (
	"fmt"
)

func main() {
	arrA := []int{1, 4, 2, 1, 1, 2}
	arrB := []int{3, 6, 3, 3}
	fmt.Println(sumSwap(arrA, arrB))
}

func sumSwap(arrA, arrB []int) (int, int) {

	//target
	//a-b = (sumA-sumB)/2
	//a-b = target

	sumA := 0
	for _, val := range arrA {
		sumA += val
	}

	sumB := 0
	hash := make(map[int]bool)
	for _, val := range arrB {
		hash[val] = true
		sumB += val
	}

	//get target
	target := getTarget(sumA, sumB)

	for _, val := range arrA {
		if ok, _ := hash[target+val]; ok {
			return val, target + val
		}
	}

	return -1, -1
}

func getTarget(sumA, sumB int) int {
	if (sumA-sumB)%2 == 0 {
		return Abs((sumA - sumB) / 2)
	}
	return -1
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
