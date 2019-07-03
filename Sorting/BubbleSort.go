package main

import (
	"fmt"
)

func less(value1 int, value2 int) bool {
	return value1 < value2
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

//BubbleSort sorting method.
func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				/* Swapping */
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

//BubbleSort2 sorting method.
func BubbleSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := 1
	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data, more)
	fmt.Println(data)

	data3 := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data3, less)
	fmt.Println(data3)

	data2 := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort2(data2, less)
	fmt.Println(data2)
}