package main

import "fmt"

func main() {

	fmt.Println(countWays1(5))

}

func countWays1(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}
	return countWays1(n-1) + countWays1(n-2) + countWays1(n-2)
}
