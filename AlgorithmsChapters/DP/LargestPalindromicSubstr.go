package main

import "fmt"

func largestPalindromicSubstr(str string) int {
	n := len(str)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	maxVal := 1
	start := 0
	for l := 1; l < n; l++ {
		for i, j := 0, l; j < n; i, j = i+1, j+1 {
			if str[i] == str[j] && dp[i+1][j-1] == j-i-1 {
				dp[i][j] = dp[i+1][j-1] + 2
				if dp[i][j] > maxVal {
					maxVal = dp[i][j] // Keeping track of max length and
					start = i         // starting position of sub-string.
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	fmt.Println("Max Length Palindromic Substrings : ", str[start:start+maxVal])
	return maxVal
}

func main() {
	str := "ABCAUCBCxxCBA"
	fmt.Println("Max Palindromic Substrings len: ", largestPalindromicSubstr(str))
}

/*
Max Length Palindromic Substrings :  BCxxCB
Max Palindromic Substrings len:  6
*/
