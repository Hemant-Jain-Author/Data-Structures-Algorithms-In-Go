package main

import "fmt"

func LongestCommonSubseq(X string, Y string) int {
	m := len(X)
	n := len(Y)
	dp := make([][]int, m+1) // Dynamic programming array.
	p := make([][]int, m+1)  // For printing the substring.

	for i := range dp {
		dp[i] = make([]int, n+1)
		p[i] = make([]int, n+1)
	}

	// Fill dp array in bottom up fashion.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				p[i][j] = 0
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
					p[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1]
					p[i][j] = 2
				}
			}
		}
	}
	printLCS(p, X, m, n)
	return dp[m][n]
}

func printLCS(p [][]int, X string, i int, j int) {
	fmt.Print("LongestCommonSubseq: ")
	printLCSUtil(p, X, i, j)
	fmt.Println()
}

func printLCSUtil(p [][]int, X string, i int, j int) {
	if i == 0 || j == 0 {
		return
	}
	if p[i][j] == 0 {
		printLCSUtil(p, X, i-1, j-1)
		fmt.Print(string(X[i-1]))
	} else if p[i][j] == 1 {
		printLCSUtil(p, X, i-1, j)
	} else {
		printLCSUtil(p, X, i, j-1)
	}
}

func main() {
	X := "carpenter"
	Y := "sharpener"
	result := LongestCommonSubseq(X, Y)
	fmt.Println("LongestCommonSubseq length:", result)
}

/*
LongestCommonSubseq: arpener
LongestCommonSubseq length: 7
*/
