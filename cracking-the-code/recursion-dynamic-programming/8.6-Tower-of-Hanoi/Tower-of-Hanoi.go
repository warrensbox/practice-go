package main

import "fmt"

func main() {
	hanoi(6, 1, 3)
}

func hanoi(n, start, end int) {
	if n == 1 {
		pm(start, end)
		return
	}
	other := 6 - (start + end)
	hanoi(n-1, start, other)
	pm(start, end)
	hanoi(n-1, other, end)
}

func pm(start, end int) {
	fmt.Printf("%v -> %v\n", start, end)
}
