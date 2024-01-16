package main

import (
	"fmt"
	"math"
)

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

func MinCostBSTTD(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	max := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
		max[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
			max[i][j] = math.MinInt32
		}
	}

	for i := 0; i < n; i++ {
		max[i][i] = arr[i]
	}
	return minCostBSTTDUtil(dp, max, 0, n-1, arr)
}

func minCostBSTTDUtil(dp [][]int, maxVal [][]int, i int, j int, arr []int) int {
	if j <= i {
		return 0
	}
	if dp[i][j] != math.MaxInt32 {
		return dp[i][j]
	}
	for k := i; k < j; k++ {
		dp[i][j] = min(dp[i][j],
			minCostBSTTDUtil(dp, maxVal, i, k, arr)+
				minCostBSTTDUtil(dp, maxVal, k+1, j, arr)+
				findMaxVal(maxVal, i, k)*findMaxVal(maxVal, k+1, j))
	}
	return dp[i][j]
}

func findMaxVal(maxVal [][]int, i int, j int) int {
	if maxVal[i][j] != math.MinInt32 {
		return maxVal[i][j]
	}
	for k := i; k < j; k++ {
		maxVal[i][j] = max(maxVal[i][j],
			max(findMaxVal(maxVal, i, k), findMaxVal(maxVal, k+1, j)))
	}
	return maxVal[i][j]
}

func MinCostBSTBU(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	maxVal := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		maxVal[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		maxVal[i][i] = arr[i]
	}

	for l := 1; l < n; l++ { // l is length of range.
		for i, j := 0, l; j < n; i, j = i+1, j+1 {
			dp[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+maxVal[i][k]*maxVal[k+1][j])
				maxVal[i][j] = max(maxVal[i][k], maxVal[k+1][j])
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	arr := []int{6, 2, 4}
	fmt.Println("Total cost:", MinCostBSTTD(arr))
	fmt.Println("Total cost:", MinCostBSTBU(arr))
}

/*
Total cost: 32
Total cost: 32
*/
