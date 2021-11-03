package main

import (
	"fmt"
	"math"
	"sort"
)

// Greedy
func minCoins(coins []int, n int, val int) int {
	if val <= 0 {
		return 0
	}
	count := 0
	sort.Slice(coins, func(i, j int) bool { return coins[i] < coins[j] })
	for i := n - 1; i >= 0 && val > 0; {
		if coins[i] <= val {
			count++
			val -= coins[i]
		} else {
			i--
		}
	}
	if val == 0 {
		return count
	}
	return -1
}

// Brute Force
func minCoins2(coins []int, n int, val int) int {
	if val == 0 {
		return 0
	}

	count := math.MaxInt32
	for i := 0; i < n; i++ {
		if coins[i] <= val {
			subCount := minCoins2(coins, n, val-coins[i])
			if subCount >= 0 {
				count = min(count, subCount+1)
			}
		}
	}
	if count != math.MaxInt32 {
		return count
	}
	return -1
}

func minCoinsBU(coins []int, n int, val int) int {
	dp := make([]int, val+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= val; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				if dp[i-coins[j]] != math.MaxInt32 {
					dp[i] = min(dp[i], dp[i-coins[j]]+1)
				}
			}
		}
	}
	if dp[val] != math.MaxInt32 {
		return dp[val]
	}
	return -1
}

func MinCoinsTD(coins []int, n int, val int) int {
	dp := make([]int, val+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	return minCoinsTDUtil(dp, coins, n, val)
}

func minCoinsTDUtil(dp []int, coins []int, n int, val int) int {
	// Base Case
	if val == 0 {
		return 0
	}
	if dp[val] != math.MaxInt32 {
		return dp[val]
	}

	// Recursion
	for i := 0; i < n; i++ {
		// check validity of a sub-problem
		if coins[i] <= val {
			subCount := minCoinsTDUtil(dp, coins, n, val-coins[i])
			if subCount != math.MaxInt32 {
				dp[val] = min(dp[val], subCount+1)
			}
		}
	}
	return dp[val]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{5, 6}
	value := 16
	/*
		coins := []int{1, 5, 6, 9, 12}
		value := 15
	*/
	n := len(coins)
	fmt.Println("Count is:", minCoins(coins, n, value))
	fmt.Println("Count is:", minCoins2(coins, n, value))
	fmt.Println("Count is:", minCoinsBU(coins, n, value))
	fmt.Println("Count is:", MinCoinsTD(coins, n, value))
}

/*
Count is: -1
Count is: 3
Count is: 3
Count is: 3
*/
