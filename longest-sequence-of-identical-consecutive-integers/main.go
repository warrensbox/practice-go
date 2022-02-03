package main

import "fmt"

func main() {

	arr := []int{1, 1, 1, 1, 2, 2, 2, 7, 7, 7, 7, 7, 7, 3, 3, 2, 2, 2, 2}
	fmt.Println(longestSubsVal(arr))

}

func longestSubs(arr []int) int {

	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	currPointer := 0
	secondPointer := 1
	max := 1
	for secondPointer < len(arr)-1 {

		if arr[currPointer] == arr[secondPointer] {
			secondPointer++

		} else {
			length := secondPointer - currPointer
			max = Max(max, length)
			currPointer = secondPointer
		}
	}
	return max
}

/*
MapName
Key : Value
1   : 2
2   : 7
7   : 5
3   : 2


*/

func longestSubsVal(arr []int) int {

	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}

	currPointer := 0
	secondPointer := 1
	max := 1
	hashMap := make(map[int]int)

	for secondPointer < len(arr)-1 {

		if arr[currPointer] == arr[secondPointer] {
			secondPointer++

		} else {
			length := secondPointer - currPointer
			max = Max(max, length)
			hashMap[length] = arr[currPointer]
			currPointer = secondPointer
		}
	}
	return hashMap[max]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
