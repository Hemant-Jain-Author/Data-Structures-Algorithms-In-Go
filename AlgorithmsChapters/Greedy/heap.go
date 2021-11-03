package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (pq MinHeap) Len() int {
	return len(pq)
}

func (pq MinHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	value := old[n-1]
	*pq = old[0 : pq.Len()-1]
	return value
}

func (pq *MinHeap) Print() {
	fmt.Println(pq)
}

func (pq *MinHeap) Empty() bool {
	return (pq.Len() == 0)
}

func (pq *MinHeap) Peek() interface{} {
	return (*pq)[0]
}

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

func main() {
	pq := &MinHeap{}
	ar := []int{2, 4, 3, 1, 7, 8, 4, 1}
	for i := 0; i < len(ar); i++ {
		heap.Push(pq, ar[i])
	}
	for pq.Len() > 0 {
		fmt.Print(pq.Peek(), heap.Pop(pq).(int), " ")
	}
	fmt.Println()

	pq2 := &MaxHeap{}
	for i := 0; i < len(ar); i++ {
		heap.Push(pq2, ar[i])
	}
	for pq2.Len() > 0 {
		fmt.Print(heap.Pop(pq2).(int), " ")
	}
}

/*
func (pq *MinHeap) percolateDown(parent int) {
	lChild := 2*parent + 1
	rChild := lChild + 1
	child := -1

	if lChild < pq.Len() {
		child = lChild
	}

	if rChild < pq.Len() && pq.Less(lChild, rChild) {
		child = rChild
	}

	if child != -1 && pq.Less(parent, child) {
		pq.Swap(parent, child)
		pq.percolateDown(child)
	}
}
func (pq *MinHeap) percolateUp(position int) {
	parent := (position - 1) / 2
	if parent < 0 {
		return
	}
	if pq.Less(parent, position) {
		pq.Swap(parent, position)
		pq.percolateUp(parent)
	}
}
*/
