package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func ChotaBhim(cups []int) int {
	size := len(cups)
	time := 60

	sort.Slice(cups, func(i, j int) bool {
		return cups[i] > cups[j]
	}) // sort decreasing

	total := 0
	var index int
	var temp int
	for time > 0 {
		total += cups[0]
		cups[0] = int(math.Ceil(float64(cups[0]) / 2.0))
		index = 0
		temp = cups[0]
		for index < size-1 && temp < cups[index+1] {
			cups[index] = cups[index+1]
			index += 1
		}
		cups[index] = temp
		time -= 1
	}
	fmt.Println("Total:", total)
	return total
}

func ChotaBhim2(cups []int) int {
	size := len(cups)
	time := 60
	pq := &MaxHeap{}

	i := 0
	for i = 0; i < size; i++ {
		heap.Push(pq, cups[i])
	}
	total := 0
	var value int
	for time > 0 {
		value = heap.Pop(pq).(int)
		total += value
		value = int(math.Ceil(float64(value) / 2.0))
		heap.Push(pq, value)
		time -= 1
	}
	fmt.Println("Total:", total)
	return total
}

func main() {
	cups := []int{2, 1, 7, 4, 2}
	ChotaBhim(cups)
	cups2 := []int{2, 1, 7, 4, 2}
	ChotaBhim2(cups2)
}

/*
Total: 76
Total: 76
*/

type MaxHeap []int

func (pq MaxHeap) Len() int {
	return len(pq)
}

func (pq MaxHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MaxHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	value := old[n-1]
	*pq = old[0 : pq.Len()-1]
	return value
}

func (pq *MaxHeap) Print() {
	fmt.Println(pq)
}

func (pq *MaxHeap) Empty() bool {
	return (pq.Len() == 0)
}

func (pq *MaxHeap) Peek() interface{} {
	return (*pq)[0]
}