package main

import "fmt"

func main() {

	num1 := "56"
	num2 := "140"

	output := addStrings(num1, num2)

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
	for lengthNum1 >= 0 || lengthNum2 >= 0 {
		sum := carry
		if lengthNum1 >= 0 {
			sum += int(num1[lengthNum1] - '0')
			lengthNum1--
		}
		if lengthNum2 >= 0 {
			sum += int(num2[lengthNum2] - '0')
			lengthNum2--
		}

		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		ans[k] = byte(sum + '0')
		k--
	}

	if carry == 1 {
		ans[k] = byte(1 + '0')
		k--
	}

	return string(ans[k+1:])

}
