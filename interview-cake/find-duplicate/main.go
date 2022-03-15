package main

import "fmt"

/*
goal : find duplicate in O(n) time

*/
func main() {
	arr := []int{3, 4, 2, 3, 1, 5, 2}
	fmt.Println(findDulicate(arr))
}

func findDulicate(num []int) []int {
	output := []int{}
	for i := 0; i < len(num); i++ {

		index := Abs(num[i]) - 1

		if num[index] < 0 {
			output = append(output, index+1)
		}
		num[index] = -num[index]

	}
	return output
}

func Abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}
