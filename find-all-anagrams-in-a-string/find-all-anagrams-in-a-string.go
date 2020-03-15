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

	fmt.Println(charCount)

	start := 0
	end := 0
	totalCountofEachCharInP := 0

	for end < len(s) {
		fmt.Println("0000000000000")
		char := s[end]
		fmt.Println("char LLOP", string(char))
		charCount[char-'a']++
		fmt.Println("charCount", charCount)

		if charCount[char-'a'] <= 0 { //if char is in p
			totalCountofEachCharInP++
		}
		fmt.Println("totalCountofEachCharInP", totalCountofEachCharInP)
		//totalCountofEachCharInP is only modified if char in p is changed
		if totalCountofEachCharInP == len(p) {
			result = append(result, start)
		}
		fmt.Println("--------")
		fmt.Println("result", result)
		fmt.Println("end", end)
		fmt.Println("start", start)
		if end-start+1 >= len(p) {
			fmt.Println("########")
			fmt.Println("TOTTALL", end-start+1)
			char = s[start]
			fmt.Println("char IF", string(char))
			if charCount[char-'a'] <= 0 { //if it is a char is p
				totalCountofEachCharInP--
			}
			fmt.Println("totalCountofEachCharInPLATER", totalCountofEachCharInP)

			charCount[char-'a']--
			start++
			fmt.Println(charCount)
		}

		end++

	}

	return result

}
