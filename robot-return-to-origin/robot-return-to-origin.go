package main

import "fmt"

func main() {

	input := "LURD"

	output := judgeCircle(input)

	fmt.Println(output)

}

func judgeCircle(moves string) bool {

	horizontal := 0
	vertical := 0
	for _, val := range moves {
		value := string(val)

		if value == "L" {
			horizontal++
		} else if value == "R" {
			horizontal--
		} else if value == "U" {
			vertical++
		} else if value == "D" {
			vertical--
		}

	}

	return horizontal == 0 && vertical == 0

}
