package main

import (
	"fmt"
	"math"
)

func MatrixChainMulBruteForce(p []int, n int) int {
	i := 1
	j := n - 1
	return MatrixChainMulBruteForceUtil(p, i, j)
}

func MatrixChainMulBruteForceUtil(p []int, i int, j int) int {
	if i == j {
		return 0
	}
	min := math.MaxInt32

	// place parenthesis at different places between
	// first and last matrix, recursively calculate
	// count of multiplications for each parenthesis
	// placement and return the minimum count
	for k := i; k < j; k++ {
		count := MatrixChainMulBruteForceUtil(p, i, k) +
			MatrixChainMulBruteForceUtil(p, k+1, j) + p[i-1]*p[k]*p[j]
		if count < min {
			min = count
		}
	}
	return min // Return minimum count
}

func MatrixChainMulTD(p []int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	return MatrixChainMulTDUtil(dp, p, 1, n-1)
}

// Function for matrix chain multiplication
func MatrixChainMulTDUtil(dp [][]int, p []int, i int, j int) int {
	// Base Case
	if i == j {
		return 0
	}
	if dp[i][j] != math.MaxInt32 {
		return dp[i][j]
	}

	// Recursion
	for k := i; k < j; k++ {
		dp[i][j] = min(dp[i][j], MatrixChainMulTDUtil(dp, p, i, k)+
			MatrixChainMulTDUtil(dp, p, k+1, j)+p[i-1]*p[k]*p[j])
	}
	return dp[i][j]
}

func MatrixChainMulBU(p []int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 1; i < n; i++ {
		dp[i][i] = 0
	}

	for l := 1; l < n; l++ {
		for i, j := 1, l; j < n; i, j = i+1, j+1 {
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+p[i-1]*p[k]*p[j]+dp[k+1][j])
			}
		}
	}
	return dp[1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 2, 3, 4}
	n := len(arr)
	fmt.Println("Matrix Chain Multiplication is:", MatrixChainMulBruteForce(arr, n))
	fmt.Println("Matrix Chain Multiplication is:", MatrixChainMulTD(arr, n))
	fmt.Println("Matrix Chain Multiplication is:", MatrixChainMulBU(arr, n))
}

/*
Matrix Chain Multiplication is: 18
Matrix Chain Multiplication is: 18
Matrix Chain Multiplication is: 18
*/
