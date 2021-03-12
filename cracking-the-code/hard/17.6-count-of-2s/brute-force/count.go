package main

import "fmt"

func main() {
	fmt.Println(numOf2InRange(21))
}

func numOf2InRange(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		count += numOf2(i)
	}
	return count
}

func numOf2(n int) int {
	count := 0
	for n > 0 {
		if n%10 == 2 {
			count++
		}
		n = n / 10
	}

	return count
}
