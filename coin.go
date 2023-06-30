package main

import (
	"fmt"
	"math"
)

// Function to find the minimum number of coins needed to make up the target amount using the given coins
func coinChange(coins []int, amount int) int {
	// Create a DP table to store the minimum number of coins for each amount
	dp := make([]int, amount+1)

	// Initialize the DP table with a maximum value
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// Base case: if the amount is 0, the minimum number of coins is 0
	dp[0] = 0

	// Solve the subproblems in a bottom-up manner
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			// Calculate the minimum number of coins by taking the minimum between the current minimum and 1 + the minimum for the remaining amount
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
		}
	}

	// If the value in the last cell of the DP table remains unchanged, it means no combination of coins can make up the target amount
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func main() {
	// Example coins and target amount
	coins := []int{1, 2, 5}
	amount := 11

	// Find the minimum number of coins needed to make up the target amount using the given coins
	minCoins := coinChange(coins, amount)

	fmt.Printf("Minimum number of coins needed to make %d using coins %v: %d\n", amount, coins, minCoins)
}
