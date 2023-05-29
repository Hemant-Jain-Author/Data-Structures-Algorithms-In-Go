package main

import "fmt"

func LargestIncreasingSubseq(arr []int) int {
	n := len(arr)
	lis := make([]int, n)
	maxVal := 0

	// Populating LIS values in bottom up manner.
	for i := 0; i < n; i++ {
		lis[i] = 1 // Initialize LIS values for all indexes as 1.
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
		if maxVal < lis[i] { // Max LIS values.
			maxVal = lis[i]
		}
	}
	return maxVal
}

func main() {
	arr := []int{10, 12, 9, 23, 25, 55, 49, 70}
	fmt.Println("Length of lis is", LargestIncreasingSubseq(arr))
}

/*
Length of lis is 6
*/
