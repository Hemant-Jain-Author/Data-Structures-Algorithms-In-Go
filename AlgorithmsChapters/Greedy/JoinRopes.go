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
	cmp := func(a, b interface{}) bool { 
		return a.(int) < b.(int) 
	}
	pq := NewHeap(cmp)
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