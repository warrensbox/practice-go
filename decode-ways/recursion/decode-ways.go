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

	if k == 0 {
		return 1
	}

	s := len(data) - k
	if data[s] == '0' {
		return 0
	}

	result := recursion(data, k-1)
	digit, _ := strconv.Atoi(data[s : s+2])
	if k >= 2 && digit <= 26 {
		result += recursion(data, k-2)
	}

	return result
}
