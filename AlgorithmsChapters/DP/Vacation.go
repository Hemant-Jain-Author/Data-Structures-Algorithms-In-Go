package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// days are must travel days, costs are cost of tickets.
func MinCostTravel(days []int, costs []int) int {
	n := len(days)
	maxVal := days[n-1]
	dp := make([]int, maxVal+1)
	j := 0
	for i := 1; i <= maxVal; i++ {
		if days[j] == i {// That days is definitely travelled.
			j++
			dp[i] = dp[i-1] + costs[0]
			dp[i] = min(dp[i], dp[max(0, i-7)]+costs[1])
			dp[i] = min(dp[i], dp[max(0, i-30)]+costs[2])
		} else {
			dp[i] = dp[i-1] // day may be ignored.
		}
	}
	return dp[maxVal]
}

func main() {
	days := []int{1, 3, 5, 7, 12, 20, 30}
	costs := []int{2, 7, 20}
	fmt.Println("Min cost is :", MinCostTravel(days, costs))
}

/*
Min cost is: 13
*/