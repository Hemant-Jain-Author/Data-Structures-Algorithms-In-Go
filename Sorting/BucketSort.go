package main

import (
	"fmt"
	"math"
	"sort"
)

func BucketSort(arr []int, maxValue int, numBucket int) {
	length := len(arr)
	if length == 0 {
		return
	}
	bucket := make([][]int, numBucket)
	div := int(math.Ceil(float64(maxValue / numBucket)))
	
	for i := 0; i < length; i++ {
		if arr[i] < 0 || arr[i] > maxValue {
			fmt.Println("Value out of range.")
			return
		}
		bucketIndex := arr[i] / div
		if bucketIndex >= numBucket {
			bucketIndex = numBucket - 1
		}
		bucket[bucketIndex] = append(bucket[bucketIndex], arr[i])
	}
	for i := 0; i < numBucket; i++ {
		sort.Slice(bucket[i], func(a, b int) bool {
			return bucket[i][a] < bucket[i][b]
		})
	}
	index := 0
	var count int
	for i := 0; i < numBucket; i++ {
		temp := bucket[i]
		count = len(temp)
		for j := 0; j < count; j++ {
			arr[index] = temp[j]
			index++
		}
	}
}

// Testing code.
func main() {
	array := []int{1, 34, 7, 99, 5, 23, 45, 88, 77, 19, 91, 100}
	maxValue := 100
	numBucket := 5
	BucketSort(array, maxValue, numBucket)
	fmt.Print(array)
}

// [1 5 7 19 23 34 45 77 88 91 99 100]
