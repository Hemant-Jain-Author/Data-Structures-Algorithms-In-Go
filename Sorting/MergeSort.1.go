package main

import (
	"fmt"
)

//will return true when the elements are out of order.
func more(value1 int, value2 int) bool {
	return value1 > value2
}

func less(value1 int, value2 int) bool {
	return value1 < value2
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	MergeSort(data, more)
	fmt.Println(data)
}

func MergeSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	tempArray := make([]int, size)
	mergeSrt(arr, tempArray, 0, size-1, comp)
}

func merge(arr []int, tempArray []int, lowerIndex int, middleIndex int, upperIndex int, comp func(int, int) bool) {
	lowerStart := lowerIndex
	lowerStop := middleIndex
	upperStart := middleIndex + 1
	upperStop := upperIndex
	count := lowerIndex
	for lowerStart <= lowerStop && upperStart <= upperStop {
		if comp(arr[lowerStart], arr[upperStart]) == false { // comp return true when out of order.
			tempArray[count] = arr[lowerStart]
			lowerStart++
		} else {
			tempArray[count] = arr[upperStart]
			upperStart++
		}
		count++
	}
	for lowerStart <= lowerStop {
		tempArray[count] = arr[lowerStart]
		count++
		lowerStart++
	}
	for upperStart <= upperStop {
		tempArray[count] = arr[upperStart]
		count++
		upperStart++
	}
	for i := lowerIndex; i <= upperIndex; i++ {
		arr[i] = tempArray[i]
	}
}

func mergeSrt(arr []int, tempArray []int, lowerIndex int, upperIndex int, comp func(int, int) bool) {
	if lowerIndex >= upperIndex {
		return
	}
	middleIndex := (lowerIndex + upperIndex) / 2
	mergeSrt(arr, tempArray, lowerIndex, middleIndex, comp)
	mergeSrt(arr, tempArray, middleIndex+1, upperIndex, comp)
	merge(arr, tempArray, lowerIndex, middleIndex, upperIndex, comp)
}
