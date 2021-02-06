package main

import "fmt"

func main() {
	permute("aabc")
}

func permute(input string) {
	countMap := make(map[rune]int)

	for _, char := range input {
		countMap[char]++
	}

	str := make([]rune, len(input)-1)
	count := make([]int, len(input)-1)
	index := 0
	for k := range countMap {
		fmt.Println("k", string(k))
		str[index] = k
		count[index] = countMap[k]
		index++
	}

	fmt.Println("countMap", countMap)
	fmt.Println("str", str)
	fmt.Println("count", count)
	result := make([]rune, len(input))

	permUtil(str, count, result, 0)
}

func permUtil(str []rune, count []int, result []rune, level int) {
	if level == len(result) {
		fmt.Println(result)
		for _, R := range result {
			fmt.Println(string(R))
		}
		return
	}

	for i := 0; i < len(str); i++ {
		fmt.Println("count[i]", count[i])
		if count[i] == 0 {
			continue
		}
		result[level] = str[i]
		count[i]--
		permUtil(str, count, result, level+1)
		count[i]++
	}
}
