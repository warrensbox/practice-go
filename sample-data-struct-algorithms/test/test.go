package main

import "fmt"

func main() {

	/* mute start
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
	 mute end */

	float1 := "0.13.45"
	float2 := "0.13.95"

	if float1 > float2 {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}

}
