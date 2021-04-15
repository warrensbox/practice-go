package main

import "fmt"

func main() {

	s := "cbaebabacd"
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
		fmt.Println("s", string(char))
		fmt.Println("end", end)
		fmt.Println("charCount", charCount)
		charCount[char-'a']++
		fmt.Println("After charCount", charCount)
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
			fmt.Println("start", start)
		}

		end++
		fmt.Println()
	}

	return result

}
