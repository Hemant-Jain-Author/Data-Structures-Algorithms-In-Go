package main

import (
	"fmt"
)

func less(value1, value2 int) bool {
	return value1 < value2
}

func greater(value1, value2 int) bool {
	return value1 > value2
}

// BubbleSort sorts the array using the specified comparison function.
func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				// Swapping
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

// BubbleSort2 sorts the array using the specified comparison function.
// It includes an optimization to stop early if no swaps are made in a pass.
func BubbleSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := true
	for i := 0; i < size-1 && swapped; i++ {
		swapped = false
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = true
			}
		}
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data, greater)
	fmt.Println(data)

	data3 := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data3, less)
	fmt.Println(data3)

	data2 := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort2(data2, less)
	fmt.Println(data2)
}

/*
[1 2 3 4 5 6 7 8 9]
[9 8 7 6 5 4 3 2 1]
[9 8 7 6 5 4 3 2 1]
*/
