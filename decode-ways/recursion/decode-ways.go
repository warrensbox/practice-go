package main

//watch https://www.youtube.com/watch?v=YcJTyrG3bZs
import (
	"fmt"
	"strconv"
)

func main() {

	input := "226"
	output := decodeWays(input)
	fmt.Println(output)

}

func decodeWays(data string) int {
	return recursion(data, len(data))
}

func recursion(data string, k int) int {
	fmt.Println("k0", k)
	if k == 0 {
		return 1
	}

	s := len(data) - k
	fmt.Println("k1", k)
	fmt.Println("s1", s)
	if data[s] == '0' {
		fmt.Println("returning zero")
		return 0
	}
	fmt.Println()
	result := recursion(data, k-1)
	fmt.Println("====")
	fmt.Println("k", k)
	fmt.Println("s", s)
	fmt.Println("len(data)-2", len(data)-2)
	if s <= len(data)-2 {

		digit, _ := strconv.Atoi(data[s : s+2])
		fmt.Println("digit", digit)
		if k >= 2 && digit <= 26 {
			result += recursion(data, k-2)
		}
	}

	return result
}
