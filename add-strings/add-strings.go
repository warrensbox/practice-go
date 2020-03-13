package main

import "fmt"

func main() {

	num1 := "56"
	num2 := "40"

	output := addStrings(num1)

	fmt.Println(output)

}

func addStrings(num1 string, num2 string) string {

	lengthNum1 := len(num1) - 1
	lengthNum2 := len(num2) - 1

	var ans []byte

	if lengthNum1 > lengthNum2 {
		ans = make([]byte, lengthNum1+2)
	} else {
		ans = make([]byte, lengthNum2+2)
	}

	carry := 0
	k := len(ans) - 1

}
