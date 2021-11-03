package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func JoinRopes(ropes []int, size int) int {
	sort.Slice(ropes, func(i, j int) bool {
		return ropes[i] > ropes[j]
	}) // Decreasing order.

	total := 0
	value := 0
	var index int
	length := size
	for length >= 2 {
		value = ropes[length-1] + ropes[length-2]
		total += value
		index = length - 2
		for index > 0 && ropes[index-1] < value {
			ropes[index] = ropes[index-1]
			index -= 1
		}
		ropes[index] = value
		length--
	}
	fmt.Println("Total:", total)
	return total
}
func JoinRopes2(ropes []int, size int) int {
	pq := &PriorityQueue{}
	i := 0
	for i = 0; i < size; i++ {
		heap.Push(pq, ropes[i])
	}
	total := 0
	value := 0
	for pq.Len() > 1 {
		value = heap.Pop(pq).(int)
		value += heap.Pop(pq).(int)
		heap.Push(pq, value)
		total += value
	}
	fmt.Println("Total:", total)
	return total
}
func main() {
	ropes := []int{4, 3, 2, 6}
	JoinRopes(ropes, len(ropes))
	rope2 := []int{4, 3, 2, 6}
	JoinRopes2(rope2, len(rope2))
}

/*
Total: 29
Total: 29
*/

type PriorityQueue []int

func (pq *PriorityQueue) Print() {
	fmt.Println(pq)
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a] < pq[b]
}

func (pq PriorityQueue) Empty() bool {
	return (pq.Len() == 0)
}

func (pq PriorityQueue) Peek() interface{} {
	return pq[0]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	value := old[n-1]
	*pq = old[0 : pq.Len()-1]
	return value
}
