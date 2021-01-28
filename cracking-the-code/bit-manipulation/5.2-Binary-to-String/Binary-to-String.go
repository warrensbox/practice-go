package main

import (
	"fmt"
	"strings"
)

func main() {

	bin := PrintBinary(0.625)
	fmt.Println(bin)

}

func PrintBinary(num float32) string {

	if num >= 1 || num <= 0 {
		return "Error"
	}

	var binary strings.Builder
	binary.WriteString(".")

	for num > 0 {
		if binary.Len() >= 32 {
			return "Error"
		}

		r := num * 2
		if r >= 1 {
			binary.WriteString("1")
			num = r - 1
		} else {
			binary.WriteString("0")
			num = r
		}
	}
	return binary.String()
}
