package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(CountAndSay(4))
	fmt.Println(CountAndSayRecursion(4, "1"))
}

func CountAndSay(n int) string {

	finalString := "1"
	if n <= 1 {
		return finalString
	}

	characterPointer := 0
	countPointer := 0
	stringInProgress := ""

	for n > 1 {
		for countPointer < len(finalString) {
			for countPointer <= len(finalString)-1 && finalString[countPointer] == finalString[characterPointer] {
				countPointer++
			}
			stringInProgress += strconv.Itoa(countPointer - characterPointer)
			stringInProgress += string(finalString[characterPointer])
			characterPointer = countPointer
		}
		finalString = stringInProgress
		countPointer = 0
		characterPointer = 0
		stringInProgress = ""
		n--
	}

	return finalString
}

func CountAndSayRecursion(n int, finalString string) string {

	if n <= 1 {
		return finalString
	}

	characterPointer := 0
	countPointer := 0
	stringInProgress := ""

	for countPointer < len(finalString) {
		for countPointer <= len(finalString)-1 && finalString[countPointer] == finalString[characterPointer] {
			countPointer++
		}
		stringInProgress += strconv.Itoa(countPointer - characterPointer)
		stringInProgress += string(finalString[characterPointer])
		characterPointer = countPointer
	}

	//finalString = CountAndSayRecursion(n-1, stringInProgress)

	return CountAndSayRecursion(n-1, stringInProgress)
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
