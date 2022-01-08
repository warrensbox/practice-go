package main

import (
	"fmt"
)

func main() {

	denominations := []int{1, 2, 3}
	fmt.Println(changePossibilitiesTopDown(1, denominations, 0))

}

func changePossibilitiesTopDown(amountLeft int, denominations []int, currentIndex int) int {

	fmt.Println("amountLeft BEGIN", amountLeft)
	// base cases:
	// we hit the amount spot on. yes!
	if amountLeft == 0 {
		fmt.Println("Balance", amountLeft)
		fmt.Println()
		return 1
	}

	if amountLeft < 0 {
		return 0
	}

	// we're out of denominations
	if currentIndex == len(denominations) {
		fmt.Println("currentIndex NUL", currentIndex)
		fmt.Println()
		return 0
	}

	//fmt.Printf("checking ways to make %d with %s\n",amountLeft,"tet")

	// choose a current coin
	currentCoin := denominations[currentIndex]
	fmt.Println("currentCoin", currentCoin)
	fmt.Println("currentIndex", currentIndex)
	// see how many possibilities we can get
	// for each number of times to use currentCoin
	numPossibilities := 0
	for amountLeft >= 0 {
		fmt.Printf("amountLeft: %v , denominations : %v , currentIndex+1 : %v \n", amountLeft, denominations, currentIndex+1)
		numPossibilities += changePossibilitiesTopDown(amountLeft, denominations, currentIndex+1)
		fmt.Println("numPossibilities", numPossibilities)
		fmt.Println()
		fmt.Println("amountLeft BEFORE END", amountLeft)
		fmt.Println("currentCoin BEFORE END", currentCoin)
		amountLeft -= currentCoin
		fmt.Println("amountLeft END", amountLeft)
	}

	return numPossibilities
}
