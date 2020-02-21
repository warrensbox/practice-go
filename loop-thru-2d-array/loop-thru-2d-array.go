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
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			fmt.Print(input[i][j])
		}
		fmt.Println() //print new line
	}

	subset(input, 2, 1, 4, 3)

}

func subset(input [][]int, row1 int, col1 int, row2 int, col2 int) {

	fmt.Println(input)

	for i := row1; i <= row2; i++ {
		//fmt.Println(input[i])

		for j := col1; j <= col2; j++ {
			fmt.Print(input[i][j])
		}
		fmt.Println()
	}

}
