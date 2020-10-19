package main

import "fmt"

func main() {

	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	printMatrix(arr)
	rotate(arr, 4)
	printMatrix(arr)

}

func rotate(arr [][]int, n int) {

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first
			//fmt.Println("offset", offset)

			//save top
			top := arr[first][i]

			//left -> top
			arr[first][i] = arr[last-offset][first]

			//bottom -> left
			arr[last-offset][first] = arr[last][last-offset]

			//right -> bottom
			arr[last][last-offset] = arr[i][last-offset]

			//top -> right
			arr[i][last] = top

		}
	}
}

func printMatrix(arr [][]int) {

	//i = row
	//j = column
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(" ")
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}

}
