package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
* The exercise is to write the function `countAndSay` which is a sequence of
* digit strings defined by the following recursive formula:
*
* countAndSay(1): "1" // this is the base case
* countAndSay(n): is the way you would say the string returned by countAndSay(n-1)
*
* For example, if we want to know the string returned by `countAndSay(4)`:
*
* countAndSay(1) = "1"
* countAndSay(2) = "11"
* countAndSay(3) = "21"
* countAndSay(4) = "1211"
 */

func main() {
	fmt.Println(countAndSay(3))
}

func countAndSay(n int) string {
	// Solution goes here...
	str := strCounter("1", n)

	return str
}

func strCounter(str string, n int) string {
	// n = 2, str = "1"
	// n = 1, str = ""
	if n == 1 {
		return str
	}

	charPointer := 0
	countPointer := 0
	count := 0
	var sb strings.Builder
	for countPointer < len(str) {
		if str[charPointer] == str[countPointer] {
			countPointer++
		} else {
			count = countPointer - charPointer
			sb.WriteString(strconv.Itoa(count) + string(str[charPointer]))
			charPointer = countPointer
		}
	}

	count = countPointer - charPointer
	sb.WriteString(strconv.Itoa(count) + string(str[charPointer]))
	return strCounter(sb.String(), n-1)
}
