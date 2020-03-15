package main

import "fmt"

func main() {

	s := "cbacbabacd"
	p := "abc"

	output := findAnagrams(s, p)

	fmt.Println(output)
}

func findAnagrams(s string, p string) []int {

	var result []int
	charCount := [26]int{}

	if len(s) == 0 {
		return result
	}

	for _, val := range p {
		charCount[val-'a']--
	}

	start := 0
	end := 0
	totalCountofEachCharInP := 0

	for end < len(s) {
		char := s[end]
		charCount[char-'a']++

		if charCount[char-'a'] <= 0 { //if char is in p
			totalCountofEachCharInP++
		}
		//totalCountofEachCharInP is only modified if char in p is changed
		if totalCountofEachCharInP == len(p) {
			result = append(result, start)
		}

		if end-start+1 >= len(p) {

			char = s[start]
			if charCount[char-'a'] <= 0 { //if it is a char is p
				totalCountofEachCharInP--
			}

			charCount[char-'a']--
			start++
		}

		end++

	}

	return result

}
