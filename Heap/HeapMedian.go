package main

import (
	"fmt"
)

type MedianHeap struct {
	minHeap *Heap
	maxHeap *Heap
}

func NewMedianHeap() *MedianHeap {
	min := NewHeap(nil, true)
	max := NewHeap(nil, false)

	return &MedianHeap{
		minHeap: min,
		maxHeap: max,
	}
}

func (h *MedianHeap) insert(value int) {
	empty := h.maxHeap.Empty()

	if empty {
		h.maxHeap.Add(value)
	} else {
		top := h.maxHeap.Peek()
		if top >= value {
			h.maxHeap.Add(value)
		} else {
			h.minHeap.Add(value)
		}
	}

	// size balancing
	if h.maxHeap.Size() > h.minHeap.Size()+1 {
		value := h.maxHeap.Remove()
		h.minHeap.Add(value)
	}

	if h.minHeap.Size() > h.maxHeap.Size()+1 {
		value := h.minHeap.Remove()
		h.maxHeap.Add(value)
	}
}

func (h *MedianHeap) getMedian() int {
	if h.maxHeap.Size() == 0 && h.minHeap.Size() == 0 {
		fmt.Println("HeapEmptyException")
		return 0
	}

	if h.maxHeap.Size() == h.minHeap.Size() {
		val1 := h.maxHeap.Peek()
		val2 := h.minHeap.Peek()
		return (val1 + val2) / 2
	} else if h.maxHeap.Size() > h.minHeap.Size() {
		val1 := h.maxHeap.Peek()
		return val1
	} else {
		val2 := h.minHeap.Peek()
		return val2
	}
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 1, 9, 2, 8, 3, 7, 4, 6, 5, 10, 10}
	hp := NewMedianHeap()

	for i := 0; i < 20; i++ {
		hp.insert(arr[i])
		fmt.Println("Median after insertion of ", arr[i], " is ", hp.getMedian())
	}
}
