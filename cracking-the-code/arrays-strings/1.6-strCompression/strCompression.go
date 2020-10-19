package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "abcdde"
	strCompression(str)
}

func strCompression(str string) {

	if hasCompression(str) {
		fmt.Println("str", str)
		return
	}
	countConsecutive := 1
	var result strings.Builder

	for i := 0; i < len(str); i++ {
		countConsecutive++

		if i+1 >= len(str) || str[i] != str[i+1] {
			result.WriteString(string(str[i]) + strconv.Itoa(countConsecutive))
			countConsecutive = 0
		}

	}

	compressedStr := result.String()
	fmt.Println("compressedStr", compressedStr)

}

func hasCompression(str string) bool {

	nonconsecutive := 1
	lengthStr := len(str)

	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1] {
			nonconsecutive++
		}
	}

	fmt.Println("nonconsesecutive", nonconsecutive)
	fmt.Println("lengthStr", lengthStr)

	if nonconsecutive == lengthStr {
		return true
	}

	return false
}
