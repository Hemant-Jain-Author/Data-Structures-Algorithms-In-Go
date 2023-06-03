package main

import (
	"fmt"
	"sort"
)

type Heap struct {
	size int
	arr  []interface{}
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

	h := &Heap{comp: comp, arr: arr, size: size}
	for i := (size / 2); i >= 0; i-- {
		h.percolateDown(i)
	}

	return h
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

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
	h.percolateUp(h.size - 1)
}

func (h *Heap) Remove() interface{} {
	if h.IsEmpty() {
		fmt.Println("HeapEmptyException.")
		return nil
	}
	value := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.percolateDown(0)
	h.arr = h.arr[0:h.size]
	return value
}

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

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() interface{} {
	if h.IsEmpty() {
		fmt.Println("Heap empty exception.")
		return nil
	}
	return h.arr[0]
}

func (h *Heap) Print() {
	fmt.Println("Heap size:", h.size)
	fmt.Print("Heap Array:")
	for i := 0; i < h.size; i++ {
		fmt.Print(" ", h.arr[i])
	}
	fmt.Println()
}

func minComp(i, j interface{}) bool { // always i < j in use
	return i.(int) > j.(int) // swaps for min heap
}

func maxComp(i, j interface{}) bool { // always i < j in use
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
	for !hp.IsEmpty() {
		fmt.Print(hp.Remove(), " ")
	}
}

/*
Heap size : 7
Heap Array : 1 3 2 7 6 5 4
1 2 3 4 5 6 7
*/

//Testing Code
func main2() {
	a := []interface{}{1, 2, 10, 8, 7, 3, 4, 6, 5, 9}
	hp := CreateHeap(minComp, a) // Min Heap
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

//Testing Code
func main3() {
	a := []interface{}{1, 2, 10, 8, 7, 3, 4, 6, 5, 9}
	hp := CreateHeap(maxComp, a) // Max Heap
	hp.Print()
	for !hp.IsEmpty() {
		fmt.Print(hp.Remove(), " ")
	}
}

/*
Heap size : 10
Heap Array : 10 9 4 8 7 3 1 6 5 2
10 9 8 7 6 5 4 3 2 1
*/

func HeapSort(arrInput []interface{}, ismin bool) {
	var hp *Heap
	if ismin {
		hp = CreateHeap(minComp, arrInput)
	} else {
		hp = CreateHeap(maxComp, arrInput)
	}

	n := len(arrInput)
	for i := 0; i < n; i++ {
		arrInput[i] = hp.Remove().(int)
	}
}

//Testing Code
func main4() {
	a := []interface{}{1, 9, 6, 7, 8, 0, 2, 4, 5, 3}
	HeapSort(a, true)
	fmt.Println(a)

	a1 := []interface{}{1, 9, 6, 7, 8, 0, 2, 4, 5, 3}
	HeapSort(a1, false)
	fmt.Println(a1)
}

/*
[0 1 2 3 4 5 6 7 8 9]
[9 8 7 6 5 4 3 2 1 0]
*/

func IsMinHeap(arr []int) bool {
	var lchild, rchild int
	size := len(arr)
	for parent := 0; parent < (size/2 + 1); parent++ {
		lchild = parent*2 + 1
		rchild = parent*2 + 2
		// heap property check.
		if ((lchild < size) && (arr[parent] > arr[lchild])) ||
			((rchild < size) && (arr[parent] > arr[rchild])) {
			return false
		}
	}
	return true
}

func IsMaxHeap(arr []int) bool {
	var lchild, rchild int
	size := len(arr)
	for parent := 0; parent < (size/2 + 1); parent++ {
		lchild = parent*2 + 1
		rchild = lchild + 1
		// heap property check.
		if ((lchild < size) && (arr[parent] < arr[lchild])) ||
			((rchild < size) && (arr[parent] < arr[rchild])) {
			return false
		}
	}
	return true
}

// Testing Code
func main5() {
	bb := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(IsMinHeap(bb))
	cc := []int{8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(IsMaxHeap(cc))
}

/*
true
true
*/

type MedianHeap struct {
	minHeap *Heap
	maxHeap *Heap
}

func NewMedianHeap() *MedianHeap {
	min := CreateHeap(minComp)
	max := CreateHeap(maxComp)

	return &MedianHeap{
		minHeap: min,
		maxHeap: max,
	}
}

func (h *MedianHeap) insert(value int) {
	empty := h.maxHeap.IsEmpty()

	if empty {
		h.maxHeap.Add(value)
	} else {
		top := h.maxHeap.Peek().(int)
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
		val1 := h.maxHeap.Peek().(int)
		val2 := h.minHeap.Peek().(int)
		return (val1 + val2) / 2
	} else if h.maxHeap.Size() > h.minHeap.Size() {
		val1 := h.maxHeap.Peek().(int)
		return val1
	} else {
		val2 := h.minHeap.Peek().(int)
		return val2
	}
}

// Testing Code
func main6() {
	arr := []int{1, 9, 2, 8, 3, 7}
	hp := NewMedianHeap()

	for i := 0; i < 6; i++ {
		hp.insert(arr[i])
		fmt.Println("Median after insertion of", arr[i], "is", hp.getMedian())
	}
}

/*
Median after insertion of 1 is 1
Median after insertion of 9 is 5
Median after insertion of 2 is 2
Median after insertion of 8 is 5
Median after insertion of 3 is 3
Median after insertion of 7 is 5
*/

func KthSmallest(arr []int, size int, k int) int {
	sort.Ints(arr)
	return arr[k-1]
}

func KthSmallest2(arr []int, size int, k int) int {
	i := 0
	value := 0
	hp := CreateHeap(minComp)
	for i = 0; i < size; i++ {
		hp.Add(arr[i])
	}
	i = 0
	for i < size && i < k {
		value = hp.Remove().(int)
		i += 1
	}
	return value
}

func KthSmallest3(arr []int, size int, k int) int {
	hp := CreateHeap(maxComp)
	for i := 0; i < size; i++ {
		if i < k {
			hp.Add(arr[i])
		} else {
			if hp.Peek().(int) > arr[i] {
				hp.Add(arr[i])
				hp.Remove()
			}
		}
	}
	return hp.Peek().(int)
}

//Testing Code
func main7() {
	arr := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest ::", KthSmallest(arr, len(arr), 3))
	arr2 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest ::", KthSmallest2(arr2, len(arr2), 3))
	arr3 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest ::", KthSmallest3(arr3, len(arr3), 3))
}

/*
Kth Smallest :: 5
Kth Smallest :: 5
*/

func KSmallestProduct(arr []int, size int, k int) int {
	sort.Ints(arr)
	product := 1
	for i := 0; i < k; i++ {
		product *= arr[i]
	}
	return product
}

func KSmallestProduct2(arr []int, size int, k int) int {
	hp := CreateHeap(minComp)
	i := 0
	product := 1
	for i = 0; i < size; i++ {
		hp.Add(arr[i])
	}
	i = 0
	for i < size && i < k {
		product *= hp.Remove().(int)
		i += 1
	}
	return product
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func QuickSelectUtil(arr []int, lower int, upper int, k int) {
	if upper <= lower {
		return
	}

	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for lower < upper && arr[lower] <= pivot {
			lower++
		}
		for lower <= upper && arr[upper] > pivot {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}

	swap(arr, upper, start) // upper is the pivot position
	if k < upper {
		QuickSelectUtil(arr, start, upper-1, k) // pivot -1 is the upper for left sub array.
	}
	if k > upper {
		QuickSelectUtil(arr, upper+1, stop, k) // pivot + 1 is the lower for right sub array.
	}
}

func KSmallestProduct3(arr []int, size int, k int) int {
	QuickSelectUtil(arr, 0, size-1, k)
	product := 1
	for i := 0; i < k; i++ {
		product *= arr[i]
	}
	return product
}
func KSmallestProduct4(arr []int, size int, k int) int {
	hp := CreateHeap(maxComp)
	for i := 0; i < size; i++ {
		if i < k {
			hp.Add(arr[i])
		} else {
			if hp.Peek().(int) > arr[i] {
				hp.Add(arr[i])
				hp.Remove()
			}
		}
	}
	product := 1
	for i := 0; i < k; i++ {
		product *= hp.Remove().(int)
	}
	return product
}

//Testing Code
func main8() {
	arr := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product ::", KSmallestProduct(arr, 8, 3))
	arr2 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product ::", KSmallestProduct2(arr2, 8, 3))
	arr3 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product ::", KSmallestProduct3(arr3, 8, 3))
	arr4 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product ::", KSmallestProduct4(arr4, 8, 3))
}

/*
Kth Smallest product :: 10
Kth Smallest product :: 10
Kth Smallest product :: 10
Kth Smallest product :: 10
*/

func KthLargest(arr []int, size int, k int) int {
	value := 0
	hp := CreateHeap(maxComp)
	for i := 0; i < size; i++ {
		hp.Add(arr[i])
	}
	for i := 0; i < k; i++ {
		value = hp.Remove().(int)
	}
	return value
}

func PrintLargerHalf(arr []int, size int) {
	sort.Ints(arr) // , size, 1)
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func PrintLargerHalf2(arr []int, size int) {
	hp := CreateHeap(maxComp)
	for i := 0; i < size; i++ {
		hp.Add(arr[i])
	}

	for i := 0; i < size/2; i++ {
		fmt.Print(hp.Remove(), " ")
	}
	fmt.Println()
}

func PrintLargerHalf3(arr []int, size int) {
	QuickSelectUtil(arr, 0, size-1, size/2)
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

//Testing Code
func main9() {
	arr := []int{8, 7, 6, 5, 7, 5, 2, 1}
	PrintLargerHalf(arr, 8)
	arr2 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	PrintLargerHalf2(arr2, 8)
	arr3 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	PrintLargerHalf3(arr3, 8)
}

/*
6 7 7 8
8 7 7 6
6 7 7 8
*/

func sortK(arr []int, size int, k int) {
	hp := CreateHeap(minComp)
	for i := 0; i < k && i < size; i++ {
		hp.Add(arr[i])
	}

	index := 0
	for i := k; i < size; i++ {
		arr[index] = hp.Remove().(int)
		index++
		hp.Add(arr[i])
	}

	for hp.Size() > 0 {
		arr[index] = hp.Remove().(int)
		index++
	}
}

// Testing Code
func main10() {
	k := 3
	arr := []int{1, 5, 4, 10, 50, 9}
	sortK(arr, len(arr), k)
	fmt.Println(arr)
}

/*
[1 4 5 9 10 50]
*/

func kthLargestStream(k int) {
	hp := CreateHeap(minComp)
	size := 0
	data := 0

	for true {
		fmt.Println("Enter data: ")
		fmt.Scanf("%d", &data)

		if size < k-1 {
			hp.Add(data)
		} else {
			if size == k-1 {
				hp.Add(data)
			} else if hp.Peek().(int) < data {
				hp.Add(data)
				hp.Remove()
			}
			fmt.Println("Kth larges element is ::", hp, hp.Peek())
		}
		size += 1
	}
}

//Testing Code
func main11() {
	kthLargestStream(3)
}

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
	main10()
	// main11()
}
