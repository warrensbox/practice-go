package main

import (
	"fmt"

	"strconv"
)

func main() {

	// deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	// target := "0202"

	// deadends := []string{"8888"}
	// target := "0009"

	deadends := []string{"0000"}
	target := "8888"

	output := openLock(deadends, target)

	fmt.Println(output)
}

func openLock(deadends []string, target string) int {

	res := -1

	visited := make(map[string]bool)

	for _, val := range deadends {
		visited[val] = true
	}

	if visited["0000"] {
		return res
	}

	visited["0000"] = true
	queue := []string{}
	queue = append(queue, "0000")
	var rotate int

	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			combins := (queue)[0] //front
			queue = queue[1:]     //dequeue
			fmt.Println("combins", combins)
			if combins == target {
				return rotate
			}

			for i := 0; i < 4; i++ {
				j := (int(combins[i]-'0') + 1) % 10
				combins1 := combins[:i] + strconv.Itoa(j) + combins[i+1:]
				fmt.Println("combins1", combins1)
				j = (int(combins[i]-'0') - 1 + 10) % 10
				combins2 := combins[:i] + strconv.Itoa(j) + combins[i+1:]
				fmt.Println("combins2", combins2)
				if !visited[combins1] {
					visited[combins1] = true
					queue = append(queue, combins1)
				}

				if !visited[combins2] {
					visited[combins2] = true
					queue = append(queue, combins2)
				}
			}
			size--
		}
		rotate++
	}

	return res
}
