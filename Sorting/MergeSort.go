package main

import (
	"fmt"
)

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	MergeSort(data)
	fmt.Println(data)
}

func MergeSort(arr []int) {
	size := len(arr)
	tempArray := make([]int, size)
	mergeSrt(arr, tempArray, 0, size-1)
}

func merge(arr []int, tempArray []int, lowerIndex int, middleIndex int, upperIndex int) {
	lowerStart := lowerIndex
	lowerStop := middleIndex
	upperStart := middleIndex + 1
	upperStop := upperIndex
	count := lowerIndex
	for lowerStart <= lowerStop && upperStart <= upperStop {
		if arr[lowerStart] < arr[upperStart] {
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

func mergeSrt(arr []int, tempArray []int, lowerIndex int, upperIndex int) {
	if lowerIndex >= upperIndex {
		return
	}
	middleIndex := (lowerIndex + upperIndex) / 2
	mergeSrt(arr, tempArray, lowerIndex, middleIndex)
	mergeSrt(arr, tempArray, middleIndex+1, upperIndex)
	merge(arr, tempArray, lowerIndex, middleIndex, upperIndex)
}