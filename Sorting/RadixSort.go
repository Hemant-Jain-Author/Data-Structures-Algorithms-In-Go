package main

import "fmt"

func RadixSort(arr []int) {
	n := len(arr)
	m := getMax(arr, n)
	// Counting sort for every digit.
	// The dividend passed is used to calculate current working digit.
	for div := 1; m/div > 0; div *= 10 {
		countSort(arr, n, div)
	}
}

func getMax(arr []int, n int) int {
	max := arr[0]
	for i := 1; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func countSort(arr []int, n int, dividend int) {
	temp := make([]int, n)
	for i := range arr {
		temp[i] = arr[i]
	}

	// Store count of occurrences in count array.
	// (number / dividend) % 10 is used to find the working digit.
	count := make([]int, 10)
	for i := 0; i < n; i++ {
		count[(temp[i]/dividend)%10]++
	}

	// Change count[i] so that count[i] contains
	// number of elements till index i in output.
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Copy content to input arr.
	for i := n - 1; i >= 0; i -= 1 {
		arr[count[(temp[i]/dividend)%10]-1] = temp[i]
		count[(temp[i]/dividend)%10] -= 1
	}
}

func main() {
	array := []int{100, 49, 65, 91, 702, 29, 4, 55}
	RadixSort(array)
	fmt.Println(array)
}

/*
[4 29 49 55 65 91 100 702]
*/
