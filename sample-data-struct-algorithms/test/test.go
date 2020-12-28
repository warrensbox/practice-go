package main

import "fmt"

func main() {
	// Example - 1
	str := "GOLANG"
	runes := []rune(str)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	fmt.Println(result)

	// Example - 2
	s := "abcd"
	for _, r := range s {
		fmt.Printf("%c - %d\n", r, r)
		//fmt.Printf("%c - %d\n", r, r-'b')
	}

	result2 := make([]int, 3)
	result2[1] = 4
}
