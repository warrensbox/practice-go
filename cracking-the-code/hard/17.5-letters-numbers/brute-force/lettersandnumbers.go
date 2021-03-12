package main

//important subarray
import (
	"fmt"
	"unicode"
)

func main() {
	arr := []rune{'a', '1', 'b', 'B', 'c', '3', 'a', 'A', 'b', '2', 'c', '3', 'd', '4', 'e', 'E', 'g', '6'}
	subarr := findLongestSubarray(arr)
	fmt.Println(string(subarr))
}

func findLongestSubarray(array []rune) []rune {

	//see PRINT
	lengthOfArr := len(array)
	for len := lengthOfArr; len > 1; len-- {
		lengthOfSubArr := lengthOfArr - len
		for i := 0; i <= lengthOfSubArr; i++ {
			if hasEqualLettersNumbers(array, i, i+len-1) {
				return extractSubArray(array, i, i+len-1)
			}
		}
	}

	return []rune{}
}

func hasEqualLettersNumbers(array []rune, start, end int) bool {

	counter := 0
	for i := start; i <= end; i++ {
		//fmt.Print(string(array[i]))
		if unicode.IsLetter(array[i]) {
			counter++
		} else {
			counter--
		}
	}
	//fmt.Println()

	return counter == 0
}

func extractSubArray(array []rune, start, end int) []rune {
	fmt.Println(string(array[start : end+1]))
	return array[start : end+1]
}

/* PRINT */
//a1b2c3
// inital

// a1b2c
// 1b2c3
// inital

// a1b2
// 1b2c
// b2c3
// inital

// a1b
// 1b2
// b2c
// 2c3
// inital

// a1
// 1b
// b2
// 2c
// c3
// inital
