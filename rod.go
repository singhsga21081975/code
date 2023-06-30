package main

import (
	"fmt"
	"math"
)

// Function to find the maximum obtainable price for a rod of length n
func cutRod(price []float64, target int) float64 {
	// Create a DP table to store the maximum obtainable price for each length
	dp := make([]float64, target+1)

	// Base case: if the length of the rod is 0, the price is 0
	dp[0] = 0

	var maxPrice float64
	// Solve the subproblems in a bottom-up manner
	for i := 1; i <= target; i++ {
		maxPrice = -1

		// Consider all possible ways of cutting the rod
		for j := 0; j < i; j++ {
			maxPrice = math.Max(float64(maxPrice), float64(price[j]+dp[i-j-1]))
		}

		// Store the maximum obtainable price for length i in the DP table
		dp[i] = maxPrice
	}

	return dp[target]
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example prices for different lengths of rod
	price := []float64{1, 5, 8, 9, 10, 17, 17, 20,1}

	// Length of the rod
	n := len(price)

	// Find the maximum obtainable price for the given rod length
	maxPrice := cutRod(price, n)

	fmt.Printf("Maximum obtainable price for rod length %d is %f\n", n, maxPrice)
}
