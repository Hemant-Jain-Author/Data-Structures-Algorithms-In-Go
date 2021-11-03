package main

import (
	"fmt"
	"math"
)

func OptimalBSTCost(keys []int, freq []int) int {
	n := len(freq)
	return optimalBSTCostUtil(freq, 0, n-1)
}

func optimalBSTCostUtil(freq []int, i int, j int) int {
	if i > j {
		return 0
	}
	if j == i { // one element in this subarray
		return freq[i]
	}
	minVal := math.MaxInt32
	for r := i; r <= j; r++ {
		minVal = min(minVal,
			optimalBSTCostUtil(freq, i, r-1)+
				optimalBSTCostUtil(freq, r+1, j))
	}
	return minVal + sumAll(freq, i, j)
}

func sumAll(freq []int, i int, j int) int {
	s := 0
	for k := i; k <= j; k++ {
		s += freq[k]
	}
	return s
}

func OptimalBSTCostTD(keys []int, freq []int) int {
	n := len(freq)
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		cost[i][i] = freq[i]
	}
	return OptimalBSTCostTDUtil(freq, cost, 0, n-1)
}

func OptimalBSTCostTDUtil(freq []int, cost [][]int, i int, j int) int {
	if i > j {
		return 0
	}
	if cost[i][j] != math.MaxInt32 {
		return cost[i][j]
	}
	s := sumAll(freq, i, j)
	for r := i; r <= j; r++ {
		cost[i][j] = min(cost[i][j],
			OptimalBSTCostTDUtil(freq, cost, i, r-1)+
				OptimalBSTCostTDUtil(freq, cost, r+1, j)+s)
	}
	return cost[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func OptimalBSTCostBU(keys []int, freq []int) int {
	n := len(freq)
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		cost[i][i] = freq[i]
	}
	sm := 0
	// l is length of range.
	for l := 1; l < n; l++ {
		for i, j := 0, l; j < n; i, j = i+1, j+1 {
			sm = sumAll(freq, i, j)
			for r := i; r <= j; r++ {
				first, second := 0, 0
				if r-1 >= i {
					first = cost[i][r-1]
				}
				if r+1 <= j {
					second = cost[r+1][j]
				}
				cost[i][j] = min(cost[i][j], sm+first+second)
			}
		}
	}
	return cost[0][n-1]
}

func OptimalBSTCostBU2(keys []int, freq []int) int {
	n := len(freq)
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}

	sumArr := sumInit(freq, n)
	for i := 0; i < n; i++ {
		cost[i][i] = freq[i]
	}

	sm := 0
	// l is length of range.
	for l := 1; l < n; l++ {
		for i, j := 0, l; j < n; i, j = i+1, j+1 {
			sm = sumRange(sumArr, i, j)
			for r := i; r <= j; r++ {
				first, second := 0, 0
				if r-1 >= i {
					first = cost[i][r-1]
				}
				if r+1 <= j {
					second = cost[r+1][j]
				}
				cost[i][j] = min(cost[i][j], sm+first+second)
			}
		}
	}
	return cost[0][n-1]
}
func sumInit(freq []int, n int) []int {
	sum := make([]int, n)
	sum[0] = freq[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + freq[i]
	}
	return sum
}

func sumRange(sum []int, i int, j int) int {
	if i == 0 {
		return sum[j]
	}
	return sum[j] - sum[i-1]
}

func main() {
	keys := []int{9, 15, 25}
	freq := []int{30, 10, 40}
	fmt.Println("OBST cost:", OptimalBSTCost(keys, freq))
	fmt.Println("OBST cost:", OptimalBSTCostTD(keys, freq))
	fmt.Println("OBST cost:", OptimalBSTCostBU(keys, freq))
	fmt.Println("OBST cost:", OptimalBSTCostBU2(keys, freq))
}

/*
OBST cost: 130
OBST cost: 130
OBST cost: 130
OBST cost: 130
*/
