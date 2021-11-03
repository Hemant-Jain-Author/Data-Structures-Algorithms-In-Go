package main

import (
	"container/heap"
	"fmt"
)

func OptimalMergePattern(lists []int, size int) int {
	pq := &PriorityQueue{}
	i := 0
	for i = 0; i < size; i++ {
		heap.Push(pq, lists[i])
	}
	total := 0
	value := 0
	for pq.Len() > 1 {
		value = heap.Pop(pq).(int)
		value += heap.Pop(pq).(int)
		heap.Push(pq, value)
		total += value
	}
	return total
}

func main() {
	lists := []int{4, 3, 2, 6}
	fmt.Println("Total:", OptimalMergePattern(lists, len(lists)))
}

/*
Total: 29
*/

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Empty() bool {
	return (pq.Len() == 0)
}

func (pq PriorityQueue) Peek() interface{} {
	return pq[0]
}

func (pq *PriorityQueue) Print() {
	fmt.Println(pq)
}

func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a] < pq[b]
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
