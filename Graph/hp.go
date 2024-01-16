package main

import (
	"fmt"
)


type Heap struct {
	size  int
	arr   []interface{}
	comp func(x interface{}, y interface{}) bool
}

func CreateHeap(comp func(x interface{}, y interface{}) bool) *Heap {
	var arr []interface{}
	h := &Heap{comp: comp, arr : arr, size : 0}
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

func minComp (i, j interface{}) bool { // always i < j in use
	return i.(int) > j.(int) // swaps for min heap
}

func maxComp (i, j interface{}) bool { // always i < j in use
	return i.(int) < j.(int) // swap for max heap.
}

//Testing Code 
func main1() {
    hp := CreateHeap(minComp)
    hp.Add(1)
    hp.Add(6)
    hp.Add(5)
    hp.Add(7)
    hp.Add(3)
    hp.Add(4)
    hp.Add(2)
    hp.Print()
    for  !hp.IsEmpty() {
        fmt.Print(hp.Remove(), " ")
    }
}

/*
Heap size : 7
Heap Array : 1 3 2 7 6 5 4
1 2 3 4 5 6 7
*/

//Testing Code 
func main() {
    hp := CreateHeap(maxComp)
    hp.Add(1)
    hp.Add(6)
    hp.Add(5)
    hp.Add(7)
    hp.Add(3)
    hp.Add(4)
    hp.Add(2)
    hp.Print()
    for  !hp.IsEmpty() {
        fmt.Print(hp.Remove(), " ")
    }
}

/* Testing Code 
func main2() {
    a := []int{1, 2, 10, 8, 7, 3, 4, 6, 5, 9}
    hp := CreateHeap(true, a)// Min Heap
    hp.Print()
    for !hp.IsEmpty() {
        fmt.Print(hp.Remove(), " ")
    }
}
/*
Heap size : 10
Heap Array : 1 2 3 5 7 10 4 6 8 9
1 2 3 4 5 6 7 8 9 10
*/