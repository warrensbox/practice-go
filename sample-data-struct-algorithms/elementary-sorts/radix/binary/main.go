package main

import "fmt"

func main() {

	theArr := []int{4, 8, 2, 4, 2, 7, 8, 8, 9, 1}
	fmt.Println(radixSort(theArr))
}

/*
 * Returns the value of the bit at index 'bit' in 'number'
 */
func bitValue(number, bit int) int {
	mask := 1 << bit
	if (number & mask) != 0 {
		return 1
	}
	return 0
}

/*
 * Arrange the items in theArray based on the value of
 * a specific bit. This doesn't fully sort the array (it
 * just sorts by a specific bit), but we'll use it for radix
 * sort.
 */
func countingSort(mainArr []int, bit int) []int {

	// counts[0] stores the number of items with a 0 in this bit
	// counts[1] stores the number of items with a 1 in this bit
	counts := []int{0, 0}
	for _, item := range mainArr {
		counts[bitValue(item, bit)] += 1
	}

	// indices[0] stores the index where we should put the next item
	// with a 0 in this bit.
	// indices[1] stores the index where we should put the next item
	// with a 1 in this bit.
	//
	// the items with a 0 in this bit come at the beginning (index 0).
	// the items with a 1 in this bit come after all the items with a 0.
	indices := []int{0, counts[0]}

	// output array to be filled in
	sortedArray := make([]int, len(mainArr))

	for _, item := range mainArr {

		itemBitValue := bitValue(item, bit)

		// place the item at the next open index for its bit value
		sortedArray[indices[itemBitValue]] = item
		indices[itemBitValue]++
	}

	return sortedArray
}

func radixSort(mainArr []int) []int {

	for bitIndex := 0; bitIndex < 64; bitIndex++ {
		mainArr = countingSort(mainArr, bitIndex)
	}

	return mainArr
}
