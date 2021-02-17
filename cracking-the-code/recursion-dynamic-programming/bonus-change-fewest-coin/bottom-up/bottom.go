package main

import "fmt"

func main() {

	amount := 11
	coins := []int{1, 2, 5}
	fmt.Println(leastCoins(coins, amount))
}

func leastCoins(coins []int, amount int) int {

	// This table will store the answer to our sub problems
	dp := make([]int, amount+1)

	// Initialize the dp table
	for i := range dp {
		dp[i] = amount + 1
	}

	/*
	   The answer to making change with minimum coins for 0
	   will always be 0 coins no matter what the coins we are
	   given are
	*/
	dp[0] = 0

	fmt.Println(dp)

	// Solve every subproblem from 1 to amount
	for i := 1; i <= amount; i++ {
		// For each coin we are given
		for j := 0; j < len(coins); j++ {
			// If it is less than or equal to the sub problem amount
			if coins[j] <= i {
				// Try it. See if it gives us a more optimal solution
				dp[i] = Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	/*
	   dp[amount] has our answer. If we do not have an answer then dp[amount]
	   will be amount + 1 and hence dp[amount] > amount will be true. We then
	   return -1.

	   Otherwise, dp[amount] holds the answer
	*/

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
