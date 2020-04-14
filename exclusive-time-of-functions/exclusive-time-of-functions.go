package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	n := 2
	intervals := []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}
	output := exclusiveTime(n, intervals)
	fmt.Println(output)
}

func exclusiveTime(n int, logs []string) []int {

	res := make([]int, n)
	stack := []int{}

	//This is used to track the next cycle based on the current log entry.
	prevTime := 0

	for _, log := range logs {
		fmt.Println("=====")
		fmt.Println(log)
		str := strings.Split(log, ":")
		id, _ := strconv.Atoi(str[0])
		state := str[1]
		time, _ := strconv.Atoi(str[2])

		if state == "start" {

			if len(stack) > 0 {
				res[stack[len(stack)-1]] += time - prevTime
			}

			stack = append(stack, id)
			prevTime = time
		} else {

			res[stack[len(stack)-1]] += time - prevTime + 1

			stack = stack[:len(stack)-1]

			prevTime = time + 1
		}

	}

	return res
}
