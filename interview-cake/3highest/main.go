package main

import "fmt"

func main() {
	arr := []int{1, 10, -5, 1, -100}
	fmt.Println(product3(arr))
}

func product3(arr []int) int {

	lowest := Min(arr[0], arr[1])
	highest := Max(arr[0], arr[1])

	lowestProd2 := arr[0] * arr[1]
	highestProd2 := arr[0] * arr[1]

	highestProd3 := arr[0] * arr[1] * arr[2]

	for i := 2; i < len(arr); i++ {
		current := arr[i]

		highestProd3 = Max(highestProd3, Max(current*lowestProd2, current*highestProd2))
		highestProd2 = Max(highestProd2, Max(current*lowest, current*highest))
		highest = Max(highest, current)
		lowest = Min(lowest, current)
	}

	return highestProd3
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
