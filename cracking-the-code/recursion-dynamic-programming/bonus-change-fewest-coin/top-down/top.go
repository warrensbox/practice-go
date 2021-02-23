package main

import (
	"fmt"
	"math"
)

func main() {

	amount := 11
	coins := []int{1, 2, 5}
	fmt.Println(leastCoins(coins, amount))
}

func leastCoins(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	dp := make([]int, amount+1)

	return coinChange(coins, amount, dp)
}

func coinChange(coins []int, remainder int, dp []int) int {
	fmt.Println("remainder", remainder)
	/*
	   Minimum coins to make change for a negative amount is -1.
	   This is just a base case we arbitrarily define.
	*/
	if remainder < 0 {
		return -1
	}

	/*
	   The minimum coins needed to make change for 0 is always 0
	   coins no matter what coins we have.
	*/
	if remainder == 0 {
		return 0
	}

	// We already have an answer cached. Return it.
	if dp[remainder] != 0 {
		return dp[remainder]
	}

	/*
	   No answer yet. Try each coin as the last coin in the change that
	   we make for the amount
	*/
	minimum := math.MaxInt32
	for _, coin := range coins {
		fmt.Println("coin", coin)
		fmt.Println("remainder-loop", remainder)
		changeResult := coinChange(coins, remainder-coin, dp)

		/*
		   If making change was possible (changeResult >= 0) and
		   the change result beats our present minimum, add one to
		   that smallest value

		   We accept that coin as the last coin in our change making
		   sequence up to this point since it minimizes the coins we
		   need
		*/
		fmt.Println("changeResult", changeResult)
		fmt.Println("minimum", minimum)
		if changeResult >= 0 && changeResult < minimum {
			minimum = 1 + changeResult
		}

		fmt.Println()
		fmt.Println("--------")
	}

	/*
	   If no answer is found (minimum == Integer.MAX_VALUE) then the
	   sub problem answer is just arbitrarily made to be -1, otherwise
	   the sub problem's answer is "minimum"
	*/

	if minimum == math.MaxInt32 {
		dp[remainder] = -1
	} else {
		dp[remainder] = minimum
	}

	fmt.Println("========")
	fmt.Println("remainder", remainder)
	fmt.Println("dp[remainder]", dp[remainder])
	return dp[remainder]
}
