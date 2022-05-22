package main

import (
	"fmt"
	"math"
	"strconv"
)

func MatrixChainMulBruteForce(p []int, n int) int {
	i := 1
	j := n - 1
	return matrixChainMulBruteForceUtil(p, i, j)
}

func matrixChainMulBruteForceUtil(p []int, i int, j int) int {
	if i == j {
		return 0
	}
	min := math.MaxInt32

	// place parenthesis at different places between
	// first and last matrix, recursively calculate
	// count of multiplications for each parenthesis
	// placement and return the minimum count
	for k := i; k < j; k++ {
		count := matrixChainMulBruteForceUtil(p, i, k) +
			matrixChainMulBruteForceUtil(p, k+1, j) + p[i-1]*p[k]*p[j]
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
	return matrixChainMulTDUtil(dp, p, 1, n-1)
}

// Function for matrix chain multiplication
func matrixChainMulTDUtil(dp [][]int, p []int, i int, j int) int {
	// Base Case
	if i == j {
		return 0
	}
	if dp[i][j] != math.MaxInt32 {
		return dp[i][j]
	}

	// Recursion
	for k := i; k < j; k++ {
		dp[i][j] = min(dp[i][j], matrixChainMulTDUtil(dp, p, i, k)+
			matrixChainMulTDUtil(dp, p, k+1, j)+p[i-1]*p[k]*p[j])
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

func MatrixChainMulBU2(p []int, n int) int {
	dp := make([][]int, n)
	pos := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		pos[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
			pos[i][i] = i
		}
	}

	for i := 1; i < n; i++ {
		dp[i][i] = 0
	}

	for l := 1; l < n; l++ {
		for i, j := 1, l; j < n; i, j = i+1, j+1 {
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+p[i-1]*p[k]*p[j]+dp[k+1][j])
				pos[i][j] = k
			}
		}
	}
	PrintOptimalParenthesis(n, pos)
	return dp[1][n-1]
}

func PrintOptPar(n int, pos [][]int, i int, j int) string {
	output := "";
	if i == j{
		output +=  "M" + strconv.Itoa(pos[i][i]) + " ";
	} else {
		output +=  "( ";
		output +=  PrintOptPar(n, pos, i, pos[i][j])
		output +=  PrintOptPar(n, pos, pos[i][j] + 1, j)
		output +=  ") ";
	}
	return output;
};

func PrintOptimalParenthesis(n int, pos [][]int) {
	fmt.Println("OptimalParenthesis :", PrintOptPar(n, pos , 1, n - 1));
};
func main() {
	arr := []int{1, 2, 3, 4}
	n := len(arr)
	fmt.Println("Matrix Chain Multiplication is:", MatrixChainMulBruteForce(arr, n))
	fmt.Println("Matrix Chain Multiplication is:", MatrixChainMulTD(arr, n))
	fmt.Println("Matrix Chain Multiplication is:", MatrixChainMulBU(arr, n))
	fmt.Println("Matrix Chain Multiplication is:", MatrixChainMulBU2(arr, n))
}

/*
Matrix Chain Multiplication is: 18
Matrix Chain Multiplication is: 18
Matrix Chain Multiplication is: 18
OptimalParenthesis : ( ( M1 M2 ) M3 ) 
Matrix Chain Multiplication is: 18
*/
