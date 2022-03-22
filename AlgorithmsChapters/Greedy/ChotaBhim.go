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
	cmp := func(a, b interface{}) bool { 
		return a.(int) > b.(int) 
	}
	pq := NewHeap(cmp)
	for i := 0; i < size; i++ {
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
