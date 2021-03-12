package main

//important subarray
import (
	"fmt"
	"unicode"
)

func main() {
	arr := []rune{'a', 'a', 'a', 'a', '1', '1', 'a', '1', '1', 'a', 'a', '1', 'a', 'a', '1', 'a', 'a', 'a', 'a', 'a'}
	subarr := findLongestSubarray(arr)
	fmt.Println(string(subarr))
}

func findLongestSubarray(array []rune) []rune {

	delta := aggChar(array)
	// fmt.Println(let)
	// fmt.Println(num)
	fmt.Println()
	extract(delta)
	return []rune{}
}

func aggChar(array []rune) []int {

	aggLet := make([]int, len(array))
	aggNum := make([]int, len(array))
	delta := make([]int, len(array))
	letters := 0
	numbers := 0
	for i := 0; i < len(array); i++ {
		if unicode.IsLetter(array[i]) {
			letters++
			aggLet[i] = letters
			aggNum[i] = numbers
		} else {
			numbers++
			aggLet[i] = letters
			aggNum[i] = numbers
		}
		delta[i] = Abs(aggLet[i] - aggNum[i])
	}

	fmt.Println(delta)
	fmt.Println()
	return delta
}

func extract(delta []int) {

	max := -1
	hash := make(map[int]int)
	memory := make([]int, 2)
	for i := 0; i < len(delta); i++ {

		if val, ok := hash[delta[i]]; ok {
			distance := i - val
			fmt.Println(hash[delta[i]], distance)
			if distance > max {
				max = distance
				memory[0] = hash[delta[i]] + 1
				memory[1] = i
			}
		} else {
			hash[delta[i]] = i
		}
	}

	fmt.Println(memory)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
