package main

import "fmt"

func main() {
	arr := []int{5, 1, 1, 5}
	fmt.Println(adjacentDp(arr))
}

/*
	2, 4, 6, 2, 5
	0  1  2  3  4

index  0, 2,4
	   0, 3

	   1,3,
	   1,4

	   2,4
*/

//brute force way O(n^3)
func adjacent(arr []int) int {

	maxSum := 0
	for i := 0; i < len(arr); i++ {

		for k := 2; k < len(arr); k++ {
			sum := 0
			for j := i; j < len(arr); j += k {
				fmt.Println(j)
				sum += arr[j]
			}
			maxSum = Max(maxSum, sum)
			fmt.Println("-++-")
		}
		fmt.Println("----")

	}
	return maxSum
}

//dpway
/*
incl = [excl + arr[i]]
excl = Max(incl,excl)

//iterate through the arr

return Max(incl,excl)
*/

func adjacentDp(arr []int) int {

	incl := arr[0]
	excl := 0

	for i := 1; i < len(arr); i++ {

		new_excl := Max(incl, excl)
		/* current max including i */
		incl = excl + arr[i]
		excl = new_excl
	}

	return Max(incl, excl)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
