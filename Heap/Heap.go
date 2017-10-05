package main

import "fmt"

func main() {
	a := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
	HeapSort(a)
	fmt.Println(a)
	hp := NewHeap(nil, true)
	n := len(a)
	for i := 0; i < n; i++ {
		hp.Add(a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println("pop value :: ", hp.Remove())
	}
	*/
	/*n := len(a)
	hp := NewHeap(a, true)

	for i := 0; i < n; i++ {
		hp.Add(i)
	}
	for i := 0; i < n; i++ {
		fmt.Println("pop value :: ", hp.Remove())
	}
	aa := []int{1, 9, 6, 7, 8, 1, 2, 4, 5, 3}
	HeapSort(aa)
	fmt.Println("value after heap sort::")
	for i := 0; i < n; i++ {
		fmt.Print(" ", aa[i])
	}
	bb := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(IsMinHeap(bb))
	cc := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(IsMaxHeap(cc))
	*/

}

type Heap struct {
	size  int
	arr   []int
	isMin bool
}

func NewHeap(arrInput []int, isMin bool) *Heap {
	size := len(arrInput)
	arr := []int{1}
	arr = append(arr, arrInput...)
	h := &Heap{size: size, arr: arr, isMin: isMin}
	for i := (h.size / 2); i > 0; i-- {
		h.proclateDown(i)
	}
	return h
}

func NewHeap2(isMin bool) *Heap {
	arr := []int{1}
	h := &Heap{size: 0, arr: arr, isMin: isMin}
	return h
}

func (h *Heap) comp(i, j int) bool { // always i < j in use
	if h.isMin == true {
		return h.arr[i] > h.arr[j] // swaps for min heap
	}
	return h.arr[i] < h.arr[j] // swap for max heap.
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) proclateDown(parent int) {
	lChild := 2 * parent
	rChild := lChild + 1
	small := -1
	if lChild <= h.size {
		small = lChild
	}
	if rChild <= h.size && h.comp(lChild, rChild) {
		small = rChild
	}
	if small != -1 && h.comp(parent, small) {
		h.swap(parent, small)
		h.proclateDown(small)
	}
}

func (h *Heap) proclateUp(child int) {
	parent := child / 2
	if parent == 0 {
		return
	}
	if h.comp(parent, child) {
		h.swap(child, parent)
		h.proclateUp(parent)
	}
}

func (h *Heap) Add(value int) {
	h.size++
	h.arr = append(h.arr, value)
	h.proclateUp(h.size)
}

func (h *Heap) Remove() int {
	if h.Empty() {
		fmt.Println("HeapEmptyException.")
		return 0
	}
	value := h.arr[1]
	h.arr[1] = h.arr[h.size]
	h.size--
	h.proclateDown(1)
	h.arr = h.arr[0 : h.size+1]
	return value
}

func (h *Heap) Print() {
	fmt.Print("Printing Heap size :", h.size, " :: ")
	for i := 1; i <= h.size; i++ {
		fmt.Print(" ", h.arr[i])
	}
	fmt.Println()
}

func IsMinHeap(arr []int) bool {
	size := len(arr)
	for i := 0; i <= (size-2)/2; i++ {
		if 2*i+1 < size {
			if arr[i] > arr[2*i+1] {
				return false
			}
		}
		if 2*i+2 < size {
			if arr[i] > arr[2*i+2] {
				return false
			}
		}
	}
	return true
}

func IsMaxHeap(arr []int) bool {
	size := len(arr)
	for i := 0; i <= (size-2)/2; i++ {
		if 2*i+1 < size {
			if arr[i] < arr[2*i+1] {
				return false
			}
		}
		if 2*i+2 < size {
			if arr[i] < arr[2*i+2] {
				return false
			}
		}
	}
	return true
}

func (h *Heap) Empty() bool {
	return (h.size == 0)
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() int {
	if h.Empty() {
		fmt.Println("Heap empty exception.")
		return 0
	}
	return h.arr[1]
}

func HeapSort(arrInput []int) {
	hp := NewHeap(arrInput, true)
	n := len(arrInput)
	for i := 0; i < n; i++ {
		arrInput[i] = hp.Remove()
	}
}
