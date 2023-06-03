package main

import "fmt"

// CountSort performs counting sort on the given array within the specified range.
func CountSort(arr []int, lowerRange int, upperRange int) {
	var i int
	var j int
	size := len(arr)
	rangeVal := upperRange - lowerRange
	count := make([]int, rangeVal)
	for i = 0; i < size; i++ {
		count[arr[i]-lowerRange]++
	}
	j = 0
	for i = 0; i < rangeVal; i++ {
		for ; count[i] > 0; count[i]-- {
			arr[j] = i + lowerRange
			j++
		}
	}
}

func main() {
	array := []int{23, 24, 22, 21, 26, 25, 27, 28, 21, 21}
	CountSort(array, 20, 30)
	fmt.Print(array)
}

// [21 21 21 22 23 24 25 26 27 28]
