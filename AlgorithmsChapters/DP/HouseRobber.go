package main

import "fmt"

func MaxRobbery(house []int) int {
	n := len(house)
	dp := make([]int, n)
	dp[0] = house[0]
	dp[1] = house[1]
	dp[2] = dp[0] + house[2]
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + house[i]
	}
	return max(dp[n-1], dp[n-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxRobbery2(house []int) int {
	n := len(house)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][1] = house[0]
	dp[0][0] = 0
	for i := 1; i < n; i++ {
		dp[i][1] = max(dp[i-1][0]+house[i], dp[i-1][1])
		dp[i][0] = dp[i-1][1]
	}
	return max(dp[n-1][1], dp[n-1][0])
}


func main() {
	arr := []int{10, 12, 9, 23, 25, 55, 49, 70}
	fmt.Println("Total cash :", MaxRobbery(arr))
	fmt.Println("Total cash :", MaxRobbery2(arr))
}

/*
Total cash : 160
Total cash : 160
*/