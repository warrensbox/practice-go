package main

import "fmt"

func main() {

	str := "leco"
	wordDict := []string{"le", "co"}

	output := wordBreak(str, wordDict)

	fmt.Println(output)

}

func wordBreak(str string, wordDict []string) bool {

	wordMap := make(map[string]bool)
	//put the words in map for fast lookup
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// state: f[i] means whether the substring [0:i] is breakable or not, note that i is inclusive
	breakable := make([]bool, len(str)+1)
	//	fmt.Println(breakable)
	//[false false false false false false false false false]

	// init: f[0] = true since s[0:0] is empty string which default to true
	breakable[0] = true

	fmt.Println("Length", len(str))
	// func: f[i] = f[j] == true && s[j : i] is in dict for j>= 0 && j <= i - 1
	// note the state is inclusive, so f[j] means s[0:j] with j inclusive, so the rest of string is s[j:i] not s[j+1:i]
	// j's range is from 0, to i - 1 because if j = i, s[j:i] becomes empty
	for i := 1; i <= len(str); i++ {
		for j := 0; j <= i-1; j++ {
			fmt.Println("i", i)
			fmt.Println("j", j)
			substr := str[j:i]
			fmt.Println("substr", substr)

			_, ok := wordMap[substr]

			fmt.Println("wordMap[substr]", wordMap[substr])

			fmt.Println("breakable[j]", breakable[j])

			fmt.Println("breakable[i]", breakable[i])

			if breakable[j] == true && ok {
				breakable[i] = true
				fmt.Println("breakable[i]", breakable[i])
				break
			}
			fmt.Println("6666")
		}
		fmt.Println("=========")
	}

	return breakable[len(str)]

}
