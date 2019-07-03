package main

import "fmt"

func BucketSort(data []int, lowerRange int, upperRange int) {
	rng := upperRange - lowerRange
	size := len(data)
	count := make([]int, rng)

	for i := 0; i < size; i++ {
		count[data[i]-lowerRange]++
	}

	j := 0
	for i := 0; i < rng; i++ {
		for ; count[i] > 0; count[i]-- {
			data[j] = i + lowerRange
			j++
		}
	}
}

func main() {
	data := []int{23, 24, 22, 21, 26, 25, 27, 28, 21, 21}
	BucketSort(data, 20, 30)
	fmt.Println(data)
}