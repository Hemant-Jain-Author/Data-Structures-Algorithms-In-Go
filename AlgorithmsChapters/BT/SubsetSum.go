package main

import "fmt"

func SubsetSum(arr []int, n int, target int) {
	flags := make([]bool, n)
	subsetSumUtil(arr, n, flags, 0, 0, target)
}

func subsetSumUtil(arr []int, n int, flags []bool, sum int, curr int, target int) {
	if target == sum {
		printSubset(flags, arr, n)
		return
	}
	if curr >= n || sum > target {
		return
	}
	flags[curr] = true
	subsetSumUtil(arr, n, flags, sum+arr[curr], curr+1, target)
	flags[curr] = false
	subsetSumUtil(arr, n, flags, sum, curr+1, target)
}

func printSubset(flags []bool, arr []int, size int) {
	for i := 0; i < size; i++ {
		if flags[i] {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
	fmt.Print("fdsfds")
}

func main() {
	arr := []int{15, 22, 14, 26, 32, 9, 16, 8}
	target := 53
	n := len(arr)
	SubsetSum(arr, n, target)
}

/*
15 22 16 
15 14 16 8 
22 14 9 8 
*/