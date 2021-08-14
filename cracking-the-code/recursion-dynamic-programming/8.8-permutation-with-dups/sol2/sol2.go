package main

import (
	"fmt"
)

func main() {
	permute("aabc")
}

func permute(input string) {
	countMap := make(map[rune]int)

	for _, char := range input {
		countMap[char]++
	}

	//fmt.Println("map", countMap)

	str := make([]rune, len(input)-1)
	count := make([]int, len(input)-1)
	index := 0
	for key := range countMap {
		//fmt.Println("k", string(key))
		str[index] = key
		count[index] = countMap[key]
		index++
	}

	fmt.Println("countMap", countMap)
	fmt.Println("str", string(str))
	fmt.Println("count", count)
	result := make([]rune, len(input))

	permUtil(str, count, result, 0)
}

func permUtil(str []rune, count []int, result []rune, level int) {
	fmt.Println("str", str)
	fmt.Println("count", count)
	fmt.Println("result", string(result))
	fmt.Println("level", level)

	if level == len(result) {
		for _, R := range result {
			fmt.Print(string(R))
		}
		fmt.Println("\n-----------")
		fmt.Println()
		return
	}

	for i := 0; i < len(str); i++ {
		fmt.Println("count[i]", count[i])
		if count[i] == 0 {
			continue
		}
		fmt.Println("index i back", i)
		fmt.Println("str[i]", str[i])
		result[level] = str[i]
		fmt.Println("level back", level)
		fmt.Println("result[level]", string(result[level]))
		count[i]--
		fmt.Println("count[i]", count[i])
		fmt.Println("++++++")
		permUtil(str, count, result, level+1)
		count[i]++
		fmt.Println("index i", i)
		fmt.Println("7777777")
		//fmt.Println("count[i] backtrack", count[i])
	}
}
