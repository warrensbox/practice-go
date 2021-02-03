package main

import "fmt"

func main() {

	fmt.Println(gridPaths(3, 4))
}

func gridPaths(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	return gridPaths(n-1, m) + gridPaths(n, m-1)
}
