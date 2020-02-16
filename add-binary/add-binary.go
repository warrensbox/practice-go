package main

import (
	"bytes"
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

func reverseString(str string) string {

	runes := []rune(str)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]

	}

	return string(runes)
}

func addBinary(a string, b string) string {

	i := len(a) - 1
	fmt.Println("i", i)
	j := len(b) - 1
	fmt.Println("j", j)
	carry := 0
	fmt.Println("carry", carry)
	var buffer bytes.Buffer

	for i >= 0 || j >= 0 {

		fmt.Println("i", i)
		fmt.Println("j", j)

		sum := carry
		fmt.Println("sum", sum)

		if i >= 0 {
			partA := getInt(i, a)
			fmt.Println("partA", partA)
			sum += partA
			fmt.Println("sum", sum)
		}

		if j >= 0 {
			partB := getInt(j, b)
			fmt.Println("partB", partB)
			sum += partB
			fmt.Println("sum", sum)
		}

		result := sum % 2
		fmt.Println("result", result)

		carry = sum / 2
		fmt.Println("carry", carry)

		buffer.WriteString(strconv.Itoa(result))
		fmt.Println(buffer.String())

		i--
		j--

		fmt.Println("--------")
	}

	// note: need to check if there is any remaining carry
	if carry != 0 {
		buffer.WriteString("1")
	}

	fmt.Println(buffer.String())

	return reverseString(buffer.String())

}
