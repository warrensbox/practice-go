package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := 5
	output := countAndSay(input)

	fmt.Println(output)
}

func countAndSay(n int) string {

	var result strings.Builder

	sequence1 := "1"
	//if n is less than or = 1, return 1
	if n <= 1 {
		return sequence1
	}
	//else
	temp := sequence1
	var j int
	// i starts from 2 because we know the first sequnce already
	for i := 2; i <= n; i++ {

		for j = 0; j < len(temp); j++ {

			count := 1

			tempStr := temp[j]

			for j < len(temp)-1 && temp[j] == temp[j+1] {

				count++
				j++
			}

			result.WriteString(strconv.Itoa(count) + string(tempStr))

		}
		temp = result.String()

		result.Reset()

	}
	return temp
}
