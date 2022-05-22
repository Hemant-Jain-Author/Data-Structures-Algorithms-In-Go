package main

import (
	"fmt"
	"strconv"
)

func KS01UnboundBU(wt []int, cost []int, capacity int) int {
	n := len(wt)
	dp := make([]int, capacity+1)
	for w := 1; w <= capacity; w++ {
		for i := 1; i <= n; i++ {
			if wt[i-1] <= w {
				dp[w] = max(dp[w], dp[w-wt[i-1]]+cost[i-1])
			}
		}
	}
	return dp[capacity]
}

func GetMaxCost01(wt []int, cost []int, capacity int) int {
	n := len(wt)
	return getMaxCost01Util(wt, cost, n, capacity)
}

func getMaxCost01Util(wt []int, cost []int, n int, capacity int) int {
	// Base Case
	if n == 0 || capacity == 0 {
		return 0
	}
	// Return the maximum of two cases:
	// (1) nth item is included
	// (2) nth item is not included
	first := 0
	if wt[n-1] <= capacity {
		first = cost[n-1] + getMaxCost01Util(wt, cost, n-1, capacity-wt[n-1])
	}
	second := getMaxCost01Util(wt, cost, n-1, capacity)
	return max(first, second)
}

func GetMaxCost01TD(wt []int, cost []int, capacity int) int {
	n := len(wt)
	dp := make([][]int, capacity+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return getMaxCost01TDUtil(dp, wt, cost, n, capacity)
}

func getMaxCost01TDUtil(dp [][]int, wt []int, cost []int, i int, w int) int {
	if w == 0 || i == 0 {
		return 0
	}

	if dp[w][i] != 0 {
		return dp[w][i]
	}

	// Their are two cases:
	// (1) ith item is included
	// (2) ith item is not included
	first := 0
	if wt[i-1] <= w {
		first = getMaxCost01TDUtil(dp, wt, cost, i-1, w-wt[i-1]) + cost[i-1]
	}
	second := getMaxCost01TDUtil(dp, wt, cost, i-1, w)
	dp[w][i] = max(first, second)
	return dp[w][i]
}

func GetMaxCost01BU(wt []int, cost []int, capacity int) int {
	n := len(wt)
	dp := make([][]int, capacity+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Build table dp[][] in bottom up approach.
	// Weights considered against capacity.
	for w := 1; w <= capacity; w++ {
		for i := 1; i <= n; i++ {
			// Their are two cases:
			// (1) ith item is included
			// (2) ith item is not included
			first := 0
			if wt[i-1] <= w {
				first = dp[w-wt[i-1]][i-1] + cost[i-1]
			}
			second := dp[w][i-1]
			dp[w][i] = max(first, second)
		}
	}
	printItems(dp, wt, cost, n, capacity)
	// Number of weights considered and final capacity.
	return dp[capacity][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func printItems(dp [][]int, wt []int, cost []int, n int, capacity int) {
	totalCost := dp[capacity][n];
	output := "";
	for i := n - 1; i > 0; i-- {
		if totalCost != dp[capacity][i - 1] {
			output += " (wt:" + strconv.Itoa(wt[i]) + ", cost:" + strconv.Itoa(cost[i]) + ")";
			capacity -= wt[i];
			totalCost -= cost[i];
		}
	}
	fmt.Println("Selected items are:" + output);
}

func main() {
	wt := []int{ 10, 40, 20, 30 }
	cost := []int{ 60, 40, 90, 120 }
	capacity := 50
	maxCost := KS01UnboundBU(wt, cost, capacity)
	fmt.Println("Maximum cost obtained = ", maxCost)
	maxCost = GetMaxCost01(wt, cost, capacity)
	fmt.Println("Maximum cost obtained = ", maxCost)
	maxCost = GetMaxCost01BU(wt, cost, capacity)
	fmt.Println("Maximum cost obtained = ", maxCost)
	maxCost = GetMaxCost01TD(wt, cost, capacity)
	fmt.Println("Maximum cost obtained = ", maxCost)
}

/*
Maximum cost obtained =  300
Maximum cost obtained =  210
Selected items are: (wt:30, cost:120) (wt:20, cost:90)
Maximum cost obtained =  210
Maximum cost obtained =  210
*/
