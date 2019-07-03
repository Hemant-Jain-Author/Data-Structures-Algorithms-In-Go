package main

import "fmt"

func QuickSelect(arr []int, key int) int {
	size := len(arr)
	quickSelectUtil(arr, 0, size-1, key)
	return arr[key-1]
}

func quickSelectUtil(arr []int, lower int, upper int, key int) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower++
		}
		for arr[upper] > pivot && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)                   // upper is the pivot position
	quickSelectUtil(arr, start, upper-1, key) // pivot -1 is the upper for left sub array.
	quickSelectUtil(arr, upper+1, stop, key)  // pivot + 1 is the lower for right sub array.
}

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Println("Fifth value is :", QuickSelect(data, 5))
}
