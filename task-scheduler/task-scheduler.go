package main

import (
	"fmt"
	"sort"
)

func main() {

	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2

	output := leastInterval(tasks, n)

	fmt.Println(output)
}

func leastInterval(tasks []byte, n int) int {

	charMap := [26]int{}
	var res []int

	for _, char := range tasks {
		charMap[char-'A']++
	}

	for _, char := range charMap {
		res = append(res, char)
	}

	sort.Slice(res, func(a, b int) bool {
		return res[a] < res[b]
	})

	fmt.Println(res)

	maxVal := res[25] - 1 // minus one because we dont need to wait on the last task
	fmt.Println("maxVal", maxVal)
	idleSlots := maxVal * n
	fmt.Println("idleSlots", idleSlots)

	for i := 24; i > 0; i-- {
		fmt.Println("----")
		fmt.Println("res[i]", res[i])
		fmt.Println("maxVal", maxVal)
		idleSlots -= Min(res[i], maxVal)
		if res[i] == 0 {
			break
		}
	}

	if idleSlots > 0 {
		idleSlots = idleSlots + len(tasks)
	} else {
		idleSlots = len(tasks)
	}

	return idleSlots
}

func Min(x, y int) int {

	if x < y {
		return x
	}

	return y
}
