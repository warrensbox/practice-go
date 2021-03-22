package main

import (
	"fmt"
	"strconv"
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

	finalString = CountAndSayRecursion(n-1, stringInProgress)

	return finalString
}
