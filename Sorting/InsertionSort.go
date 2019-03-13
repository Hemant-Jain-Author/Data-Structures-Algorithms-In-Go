package main

import (
	"fmt"
)

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	InsertionSort(data, more)
	fmt.Println(data)
}

func less(value1 int, value2 int) bool {
	return value1 < value2
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
