package main

import (
	"container/heap"
	"fmt"
)

func OptimalMergePattern(lists []int, size int) int {
	cmp := func(a, b interface{}) bool { 
		return a.(int) < b.(int) 
	}
	pq := NewHeap(cmp)
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

type Heap struct {
	heap []interface{}
	comp func(x interface{}, y interface{}) bool
}

func NewHeap(comp func(x interface{}, y interface{}) bool) *Heap {
	hp := new(Heap)
	hp.comp = comp
	return hp
}

func (hp Heap) Len() int {
	return len(hp.heap)
}

func (hp Heap) Less(i, j int) bool {
	return hp.comp(hp.heap[i], hp.heap[j])
}

func (hp Heap) Swap(i, j int) {
	hp.heap[i], hp.heap[j] = hp.heap[j], hp.heap[i]
}

func (hp *Heap) Push(x interface{}) {
	hp.heap = append(hp.heap, x)
}

func (hp *Heap) Pop() interface{} {
	n := len(hp.heap)
	value := hp.heap[n-1]
	hp.heap = hp.heap[0 : n-1]
	return value
}

func (hp Heap) Print() {
	fmt.Println(hp.heap)
}

func (hp Heap) Empty() bool {
	return len(hp.heap) == 0
}

func (hp Heap) Peek() interface{} {
	return hp.heap[0]
}