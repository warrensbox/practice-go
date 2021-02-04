package main

import "fmt"

func main() {

	fmt.Println(countPartition(7, 4))
}

func countPartition(n, m int) int {

	if n == 0 {
		return 1
	} else if m == 0 || n < 0 {
		return 0
	}

	return countPartition(n-m, m) + countPartition(n, m-1)

}
