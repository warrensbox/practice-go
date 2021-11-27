package main

import "fmt"

func main() {

	theArr := []int{4, 8, 2, 4, 2, 7, 8, 8, 9, 1}
	maxVal := 8
	fmt.Println(countingSort(theArr, maxVal))

}
func countingSort(theArr []int, maxVal int) []int {

	// count the number of times each value appears.
	// counts[0] stores the number of 0's in the input
	// counts[4] stores the number of 4's in the input
	// etc.
	counts := make([]int, len(theArr)+1)

	for _, item := range theArr {
		counts[item] += 1
	}

	// overwrite counts to hold the next index an item with
	// a given value goes. so, counts[4] will now store the index
	// where the next 4 goes, not the number of 4's our
	// array has.
	numItemsBefore := 0
	for i := 0; i < len(counts); i++ {
		count := counts[i]
		counts[i] = numItemsBefore
		numItemsBefore += count
	}

	// run through the input array
	sortedArray := make([]int, len(theArr))

	for _, item := range theArr {
		// place the item in the sorted array
		sortedArray[counts[item]] = item
		// and, make sure the next item we see with the same value
		// goes after the one we just placed
		counts[item] += 1
	}

	return sortedArray

}
