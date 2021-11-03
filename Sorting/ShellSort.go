package main

import "fmt"

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func ShellSort(arr []int) {
	n := len(arr)
	// Gap starts with n/2 and half in each iteration.
	for gap := n / 2; gap > 0; gap /= 2 {
		// Do a gapped insertion sort.
		for i := gap; i < n; i += 1 {
			curr := arr[i]
			// Shift elements of already sorted list
			// to find right position for curr value.
			var j int
			for j = i; j >= gap && more(arr[j-gap], curr); j -= gap {
				arr[j] = arr[j-gap]
			}
			// Put current value in its correct location
			arr[j] = curr
		}
	}
}

func main() {
	array := []int{36, 32, 11, 6, 19, 31, 17, 3}
	ShellSort(array)
	fmt.Print(array)
}
