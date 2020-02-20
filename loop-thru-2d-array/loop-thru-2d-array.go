package main

import "fmt"

func main() {

	input := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	fmt.Println("length", len(input))
	fmt.Println(input)

	for i := 0; i < len(input); i++ {
		//fmt.Println(input[i])
		for j := 0; j < len(input[i]); j++ {
			fmt.Print(input[i][j])
		}
		fmt.Println("")
	}

}
