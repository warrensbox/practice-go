package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	a := "11"
	b := "1"

	ans := addBinary(a, b)

	fmt.Println(ans)
}

func getInt(index int, s string) int {
	num, err := strconv.Atoi(string(s[index]))
	if err != nil {
		log.Fatal(err)
	}

	return num
}

func addBinary(a string, b string) string {

	i := len(a) - 1
	fmt.Println("i", i)
	j := len(b) - 1
	fmt.Println("j", j)
	carry := 0
	fmt.Println("carry", carry)
	//var buffer bytes.Buffer

	for i >= 0 || j >= 0 {

		sum := carry
		fmt.Println("sum", sum)

		if i >= 0 {
			partA := getInt(i, a)
			fmt.Println(partA)
		}
	}

}
