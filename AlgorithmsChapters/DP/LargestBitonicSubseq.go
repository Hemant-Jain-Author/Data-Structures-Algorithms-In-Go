package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func LargestBitonicSubseq(arr []int) int {
	n := len(arr)
	lis := make([]int, n)
	// Initialize LIS values for all indexes as 1.
	for i := range lis {lis[i] = 1}

	lds := make([]int, n)
	// Initialize LDS values for all indexes as 1.
	for i := range lds {lds[i] = 1}

	// Populating LIS values in bottom up manner.
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}

	// Populating LDS values in bottom up manner.
	for i := n - 1; i > 0; i-- {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[i] && lds[i] < lds[j]+1 {
				lds[i] = lds[j] + 1
			}
		}
	}

	maxVal := 0
	for i := 0; i < n; i++ {
		if maxVal < lis[i]+lds[i]-1 {
			maxVal = lis[i]+lds[i]-1
		}
	}
	return maxVal
}

func main() {
	arr := []int{1, 6, 3, 11, 1, 9, 5, 12, 3, 14, 6, 17, 3, 19, 2, 19}
	fmt.Println("Length of LBS is ", LargestBitonicSubseq(arr))
}

/*
Length of LBS is  8
*/