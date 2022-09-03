package main

import (
	"fmt"
	"math"
)

func Permutation(arr []int, i int, length int) {
	if length == i {
		fmt.Println(arr)
		return
	}
	for j := i; j < length; j++ {
		arr[i], arr[j] = arr[j], arr[i]
		Permutation(arr, i+1, length)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return
}

func Permutation2(arr []int, i int, length int) {
	if length == i {
		if isValid(arr, length) {
			fmt.Println(arr)
		}
		return
	}
	for j := i; j < length; j++ {
		arr[i], arr[j] = arr[j], arr[i]
		Permutation2(arr, i+1, length)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return
}

func isValid(arr []int, n int) bool {
	for j := 1; j < n; j++ {
		if math.Abs(float64(arr[j]-arr[j-1])) < 2 {
			return false
		}
	}
	return true
}

func Permutation3(arr []int, i int, length int) {
	if length == i {
		fmt.Println(arr)
		return
	}
	for j := i; j < length; j++ {
		arr[i], arr[j] = arr[j], arr[i]
		if isValid2(arr, i) {
			Permutation3(arr, i+1, length)
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	return
}

func isValid2(arr []int, i int) bool {
	if i < 1 || math.Abs(float64(arr[i]-arr[i-1])) >= 2 {
		return true
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4}
	Permutation(arr, 0, 4)
	fmt.Println()

	Permutation2(arr, 0, 4)
	fmt.Println()
	
	Permutation3(arr, 0, 4)
}

/*
[1 2 3 4]
[1 2 4 3]
....
[4 1 3 2]
[4 1 2 3]

[2 4 1 3]
[3 1 4 2]

[2 4 1 3]
[3 1 4 2]
*/
