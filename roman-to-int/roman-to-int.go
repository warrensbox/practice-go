package main

import "fmt"

func main() {

	str := romanToInt("MCMXCIV")

	fmt.Println(str)
}

func romanToInt(s string) int {

	dict := make(map[string]int)

	dict["I"] = 1
	dict["V"] = 5
	dict["X"] = 10
	dict["L"] = 50
	dict["L"] = 50
	dict["C"] = 100
	dict["D"] = 500
	dict["M"] = 1000

	total := 0

	next := dict[string(s[0])]
	prev := dict[string(s[0])]

	for i := 1; i < len(s); i++ {

		next = dict[string(s[i])]
		fmt.Println("next", next)
		fmt.Println("prev", prev)
		if prev < next {
			total -= prev
		} else {
			total += prev
		}

		prev = next
		fmt.Println("total", total)
		fmt.Println("-----")
	}

	total += prev

	return total

}
