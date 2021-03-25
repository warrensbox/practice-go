package main

import (
	"fmt"
)

func main() {

	//ans := lengthOfLongestSubstring("abcabcs")
	subs := slidingWindow("abcabcs")

	fmt.Println(subs)
}

func slidingWindow(subs string) int {
	aPointer := 0
	bPointer := 0
	hash := make(map[byte]bool)
	max := 0

	for bPointer < len(subs) {

		if !hash[subs[bPointer]] {
			hash[subs[bPointer]] = true
			bPointer++
			max = Max(len(hash), max)
		} else {
			delete(hash, subs[aPointer])
			aPointer++
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
