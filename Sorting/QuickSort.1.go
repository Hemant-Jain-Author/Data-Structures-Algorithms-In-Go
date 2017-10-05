package main

import (
	"fmt"
)

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	QuickSort(data, more)
	fmt.Println(data)
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func less(value1 int, value2 int) bool {
	return value1 < value2
}

func QuickSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1, comp)
}

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func quickSortUtil(arr []int, lower int, upper int, comp func(int, int) bool) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for comp(arr[lower], pivot) == false && lower < upper {
			lower++
		}
		for comp(arr[upper], pivot) && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)                  // upper is the pivot position
	quickSortUtil(arr, start, upper-1, comp) // pivot -1 is the upper for left sub array.
	quickSortUtil(arr, upper+1, stop, comp)  // pivot + 1 is the lower for right sub array.
}
