package main

import (
	"container/heap"
	"fmt"
)

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

type Person struct {
	Age int
}

func main() {
	cmp := func(a, b interface{}) bool { 
		return a.(int) < b.(int) 
	}
	pq := NewHeap(cmp)
	arr := []int{1, 2, 10, 8, 7, 3, 4, 6, 5, 9}
	for _, i := range arr {
		heap.Push(pq, i)
	}
	fmt.Println("Printing Priority Queue Heap : ", pq)
	fmt.Print("Dequeue elements of Priority Queue ::")
	for pq.Len() != 0 {
		fmt.Print(" ", heap.Pop(pq))
	}
	p2 := NewHeap(func(a, b interface{}) bool { 
		return a.(Person).Age > b.(Person).Age 
	})
	p2Data := []Person{{23}, {42}, {70}, {30}}
	for _, d := range p2Data {
		heap.Push(p2, d)
	}
	for p2.Len() > 0 {
		fmt.Println(heap.Pop(p2))
	}
}
