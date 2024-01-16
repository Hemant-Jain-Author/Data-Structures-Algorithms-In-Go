package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func LargestPalindromicSubsequence(str string) int {
	n := len(str)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// each char is itself palindromic with length 1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for l := 1; l < n; l++ {
		for i, j := 0, l; j < n; i, j = i+1, j+1 {
			if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	str := "ABCAUCBCxxCBA"
	fmt.Println("Max Palindromic Subsequence length :", LargestPalindromicSubsequence(str))
}

/*
Max Palindromic Subsequence length : 9
*/