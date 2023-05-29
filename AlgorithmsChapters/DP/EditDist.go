package main

import "fmt"

func min(a, b, c int) int {
	// Find the minimum of three numbers.
	minVal := a
	if b < a {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	return minVal
}

func EditDist(str1, str2 string) int {
	m := len(str1)
	n := len(str2)
	return editDistUtil(str1, str2, m, n)
}

func editDistUtil(str1, str2 string, m, n int) int {
	// If any one string is empty, then empty the other string.
	if m == 0 || n == 0 {
		return m + n
	}

	// If last characters of both strings are the same, ignore the last characters.
	if str1[m-1] == str2[n-1] {
		return editDistUtil(str1, str2, m-1, n-1)
	}

	// If the last characters are not the same, consider all three operations:
	// Insert last char of the second string into the first.
	// Remove the last char of the first string.
	// Replace the last char of the first string with the second.
	return 1 + min(
		editDistUtil(str1, str2, m, n-1),   // Insert
		editDistUtil(str1, str2, m-1, n),   // Remove
		editDistUtil(str1, str2, m-1, n-1), // Replace
	)
}

func EditDistDP(str1, str2 string) int {
	m := len(str1)
	n := len(str2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill dp[][] in a bottom-up manner.
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				// If any one string is empty, then empty the other string.
				dp[i][j] = i + j
			} else if str1[i-1] == str2[j-1] {
				// If last characters of both strings are the same, ignore the last characters.
				dp[i][j] = dp[i-1][j-1]
			} else {
				// If the last characters are not the same, consider all three operations:
				// Insert the last char of the second string into the first.
				// Remove the last char of the first string.
				// Replace the last char of the first string with the second.
				dp[i][j] = 1 + min(
					dp[i][j-1],   // Insert
					dp[i-1][j],   // Remove
					dp[i-1][j-1], // Replace
				)
			}
		}
	}
	return dp[m][n]
}

func main() {
	str1 := "sunday"
	str2 := "saturday"
	fmt.Println(EditDist(str1, str2))
	fmt.Println(EditDistDP(str1, str2))
}

/*
3
3
*/
