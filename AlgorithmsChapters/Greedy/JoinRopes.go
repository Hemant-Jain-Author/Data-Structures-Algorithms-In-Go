package main

import (
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
		return a.(int) > b.(int) 
	}
	hp := CreateHeap(cmp)
	
	for i := 0; i < size; i++ {
		hp.Add(ropes[i])
	}
	total := 0
	value := 0
	for hp.Size() > 1 {
		value = hp.Remove().(int)
		value += hp.Remove().(int)
		hp.Add(value)
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
	size  int
	arr   []interface{}
	comp func(x interface{}, y interface{}) bool
}

func CreateHeap(comp func(x interface{}, y interface{}) bool, args ...[]interface{}) *Heap {
	var arr []interface{}
	size := 0
	if len(args) > 0 {
		arrInput := args[0]
		arr = append(arr, arrInput...)
		size = len(arrInput)
	}

	h := &Heap{comp: comp, arr : arr, size : size}
	for i := (size / 2); i >= 0; i-- {
		h.percolateDown(i)
	}

	return h
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) percolateDown(parent int) {
	lChild := 2 * parent + 1
	rChild := lChild + 1
	child := -1
	if lChild < h.size {
		child = lChild
	}
	if rChild < h.size && h.comp(h.arr[lChild], h.arr[rChild]) {
		child = rChild
	}
	if child != -1 && h.comp(h.arr[parent], h.arr[child]) {
		h.swap(parent, child)
		h.percolateDown(child)
	}
}

func (h *Heap) percolateUp(child int) {
	parent := (child - 1) / 2
	if parent >= 0 && h.comp(h.arr[parent], h.arr[child]) {
		h.swap(child, parent)
		h.percolateUp(parent)
	}
}

func (h *Heap) Add(value interface{}) {
	h.arr = append(h.arr, value)
	h.size++
	h.percolateUp(h.size-1)
}

func (h *Heap) Remove() interface{} {
	if h.IsEmpty() {
		fmt.Println("HeapEmptyException.")
		return 0
	}
	value := h.arr[0]
	h.arr[0] = h.arr[h.size - 1]
	h.size--
	h.percolateDown(0)
	h.arr = h.arr[0 : h.size]
	return value
}


func (h *Heap) Delete( value interface{}) bool {
    for i := 0; i < h.size; i++ {
        if (h.arr[i] == value) {
            h.arr[i] = h.arr[h.size - 1]
            h.size -= 1
            h.percolateUp(i)
            h.percolateDown(i)
            return true
        }
    }
    return false
}


func (h *Heap) IsEmpty() bool {
	return (h.size == 0)
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() interface{} {
	if h.IsEmpty() {
		fmt.Println("Heap empty exception.")
		return 0
	}
	return h.arr[0]
}

func (h *Heap) Print() {
	fmt.Println("Heap size :", h.size)
	fmt.Print("Heap Array :")
	for i := 0; i < h.size; i++ {
		fmt.Print(" ", h.arr[i])
	}
	fmt.Println()
}