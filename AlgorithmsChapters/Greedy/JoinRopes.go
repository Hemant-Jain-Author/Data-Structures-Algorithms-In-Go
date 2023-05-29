package main

import (
	"fmt"
	"sort"
)

// JoinRopes joins the ropes and returns the total cost
func JoinRopes(ropes []int, size int) int {
	// Sort ropes in decreasing order
	sort.Slice(ropes, func(i, j int) bool {
		return ropes[i] > ropes[j]
	})

	total := 0
	value := 0
	var index int
	length := size

	// Join ropes until there are at least 2 ropes remaining
	for length >= 2 {
		value = ropes[length-1] + ropes[length-2] // Combine the last two ropes
		total += value                            // Add the combined value to the total cost
		index = length - 2

		// Shift ropes to make room for the combined rope
		for index > 0 && ropes[index-1] < value {
			ropes[index] = ropes[index-1]
			index -= 1
		}

		ropes[index] = value // Place the combined rope in the appropriate position
		length--
	}

	fmt.Println("Total:", total)
	return total
}

// JoinRopes2 joins the ropes using a heap and returns the total cost
func JoinRopes2(ropes []int, size int) int {
	// Define the comparison function for the heap
	cmp := func(a, b interface{}) bool {
		return a.(int) > b.(int)
	}

	// Create a heap and add the ropes to it
	hp := CreateHeap(cmp)
	for i := 0; i < size; i++ {
		hp.Add(ropes[i])
	}

	total := 0
	value := 0

	// Combine ropes until there is only one rope remaining in the heap
	for hp.Size() > 1 {
		value = hp.Remove().(int) // Remove the smallest two ropes from the heap
		value += hp.Remove().(int)
		hp.Add(value) // Add the combined rope back to the heap
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

// Heap struct represents a binary heap
type Heap struct {
	size int
	arr  []interface{}
	comp func(x interface{}, y interface{}) bool
}

// CreateHeap creates a new heap with the given comparison function
func CreateHeap(comp func(x interface{}, y interface{}) bool, args ...[]interface{}) *Heap {
	var arr []interface{}
	size := 0
	if len(args) > 0 {
		arrInput := args[0]
		arr = append(arr, arrInput...)
		size = len(arrInput)
	}

	h := &Heap{comp: comp, arr: arr, size: size}

	// Build the heap from the input array
	for i := (size / 2); i >= 0; i-- {
		h.percolateDown(i)
	}

	return h
}

// swap swaps the elements at the given indices in the heap's array
func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// percolateDown moves the element at the given index down the heap to its correct position
func (h *Heap) percolateDown(parent int) {
	lChild := 2*parent + 1
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

// percolateUp moves the element at the given index up the heap to its correct position
func (h *Heap) percolateUp(child int) {
	parent := (child - 1) / 2

	if parent >= 0 && h.comp(h.arr[parent], h.arr[child]) {
		h.swap(child, parent)
		h.percolateUp(parent)
	}
}

// Add adds a new element to the heap
func (h *Heap) Add(value interface{}) {
	h.arr = append(h.arr, value)
	h.size++
	h.percolateUp(h.size - 1)
}

// Remove removes and returns the smallest element from the heap
func (h *Heap) Remove() interface{} {
	if h.IsEmpty() {
		fmt.Println("HeapEmptyException.")
		return 0
	}

	value := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.percolateDown(0)
	h.arr = h.arr[0:h.size]
	return value
}

// Delete deletes the given value from the heap
func (h *Heap) Delete(value interface{}) bool {
	for i := 0; i < h.size; i++ {
		if h.arr[i] == value {
			h.arr[i] = h.arr[h.size-1]
			h.size -= 1
			h.percolateUp(i)
			h.percolateDown(i)
			return true
		}
	}
	return false
}

// IsEmpty checks if the heap is empty
func (h *Heap) IsEmpty() bool {
	return (h.size == 0)
}

// Size returns the size of the heap
func (h *Heap) Size() int {
	return h.size
}

// Peek returns the smallest element in the heap without removing it
func (h *Heap) Peek() interface{} {
	if h.IsEmpty() {
		fmt.Println("Heap empty exception.")
		return 0
	}
	return h.arr[0]
}

// Print prints the heap's size and array
func (h *Heap) Print() {
	fmt.Println("Heap size:", h.size)
	fmt.Print("Heap Array:")
	for i := 0; i < h.size; i++ {
		fmt.Print(" ", h.arr[i])
	}
	fmt.Println()
}
