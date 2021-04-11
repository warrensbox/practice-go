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

	hash := make(map[rune]int)
	for _, val := range tasks {
		hash[rune(val)]++
	}
	values := make([]int, 0, len(hash))
	for _, v := range hash {
		values = append(values, v)
	}
	sort.Ints(values)
	maxFreq := values[len(values)-1]
	var numOfTaskMax int
	for _, v := range values {
		if v == maxFreq {
			numOfTaskMax++
		}
	}
	fmt.Println("numOfTaskMax", numOfTaskMax)

	//return (maxFreq-1)*(n+1) + numOfTaskMax //WAIT what if n=0
	return Max(len(tasks), (maxFreq-1)*(n+1)+numOfTaskMax)
}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
