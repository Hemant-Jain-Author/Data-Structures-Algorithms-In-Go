package main

import "fmt"

func MakePairs(nuts []int, bolts []int) {
	makePairsUtil(nuts, bolts, 0, len(nuts)-1)
	fmt.Print("Matched nuts and bolts are : ", nuts, " & ", bolts)
}

func makePairsUtil(nuts []int, bolts []int, low int, high int) {
	if low < high {
		pivot := partition(nuts, low, high, bolts[low])
		partition(bolts, low, high, nuts[pivot])
		makePairsUtil(nuts, bolts, low, pivot-1)
		makePairsUtil(nuts, bolts, pivot+1, high)
	}
}

func partition(arr []int, low int, high int, pivot int) (int) {
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			swap(arr, i, j)
			i++
		} else if arr[j] == pivot {
			swap(arr, high, j)
			j--
		}
	}
	swap(arr, i, high)
	return i
}

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func main() {
	nuts := []int{1, 2, 6, 5, 4, 3}
	bolts := []int{6, 4, 5, 1, 3, 2}
	MakePairs(nuts, bolts)
}

/*
Matched nuts and bolts are : [1 2 3 4 5 6] & [1 2 3 4 5 6]
*/