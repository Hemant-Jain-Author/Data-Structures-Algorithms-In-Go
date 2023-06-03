package main

import (
	"fmt"
)

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	SelectionSort(data)
	fmt.Println(data)

	data2 := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	SelectionSort2(data2)
	fmt.Println(data2)
}

/*
[1 2 3 4 5 6 7 8 9]
[1 2 3 4 5 6 7 8 9]

*/

func SelectionSort(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		max := 0
		for j := 1; j < size-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
	}
}

func SelectionSort2(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
