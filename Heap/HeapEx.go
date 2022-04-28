package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func Demo(args []string) {
	pq := &MinHeap{}
	arr := []int{1, 2, 10, 8, 7, 3, 4, 6, 5, 9}
	for _, i := range arr {
		heap.Push(pq, i)
	}
	fmt.Println("Printing Priority Queue Heap : ", pq)
	fmt.Print("Dequeue elements of Priority Queue ::")
	for pq.Len() != 0 {
		fmt.Print(" ", heap.Pop(pq))
	}
}
func IsMaxHeap(arr []int, size int) bool {
	var lchild int
	var rchild int
	for parent := 0; parent < size/2+1; parent++ {
		lchild = parent*2 + 1
		rchild = lchild + 1
		if lchild < size && arr[parent] < arr[lchild] || rchild < size && arr[parent] < arr[rchild] {
			return false
		}
	}
	return true
}
func IsMinHeap(arr []int, size int) bool {
	var lchild int
	var rchild int
	for parent := 0; parent < size/2+1; parent++ {
		lchild = parent*2 + 1
		rchild = parent*2 + 2
		if lchild < size && arr[parent] > arr[lchild] || rchild < size && arr[parent] > arr[rchild] {
			return false
		}
	}
	return true
}
func KSmallestProduct(arr []int, size int, k int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	product := 1
	for i := 0; i < k; i++ {
		product *= arr[i]
	}
	return product
}
func KSmallestProduct2(arr []int, size int, k int) int {
	pq := &MinHeap{}
	i := 0
	product := 1
	for i = 0; i < size; i++ {
		heap.Push(pq, arr[i])
	}
	i = 0
	for i < size && i < k {
		product *= heap.Pop(pq).(int)
		i += 1
	}
	return product
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
	pq := &MaxHeap{}
	for i := 0; i < size; i++ {
		if i < k {
			heap.Push(pq, arr[i])
		} else {
			if pq.Peek().(int) > arr[i] {
				heap.Push(pq, arr[i])
				heap.Pop(pq)
			}
		}
	}
	product := 1
	for i := 0; i < k; i++ {
		product *= heap.Pop(pq).(int)
	}
	return product
}
func KthLargest(arr []int, size int, k int) int {
	value := 0
	pq := &MaxHeap{}
	for i := 0; i < size; i++ {
		heap.Push(pq, arr[i])
	}
	for i := 0; i < k; i++ {
		value = heap.Pop(pq).(int)
	}
	return value
}
func KthSmallest(arr []int, size int, k int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[k-1]
}
func KthSmallest2(arr []int, size int, k int) int {
	pq := &MinHeap{}
	for i := 0; i < size; i++ {
		heap.Push(pq, arr[i])
	}
	for i := 0; i < k-1; i++ {
		heap.Pop(pq)
	}
	return pq.Peek().(int)
}
func KthSmallest3(arr []int, size int, k int) int {
	pq := &MaxHeap{}
	for i := 0; i < size; i++ {
		if i < k {
			heap.Push(pq, arr[i])
		} else {
			if pq.Peek().(int) > arr[i] {
				heap.Push(pq, arr[i])
				heap.Pop(pq)
			}
		}
	}
	return pq.Peek().(int)
}

func Main0() {
	arr := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest :: ", KthSmallest(arr, len(arr), 3))
	arr2 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest :: ", KthSmallest2(arr2, len(arr2), 3))
}
func Main1() {
	arr3 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("isMaxHeap :: ", IsMaxHeap(arr3, len(arr3)))
	arr4 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("isMinHeap :: ", IsMinHeap(arr4, len(arr4)))
}
func Main2() {
	arr := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product:: ", KSmallestProduct(arr, 8, 3))
	arr2 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product:: ", KSmallestProduct2(arr2, 8, 3))
	arr3 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product:: ", KSmallestProduct3(arr3, 8, 3))
	arr4 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	fmt.Println("Kth Smallest product:: ", KSmallestProduct4(arr4, 8, 3))
}
func Main3() {
	arr := []int{8, 7, 6, 5, 7, 5, 2, 1}
	PrintLargerHalf(arr, 8)
	arr2 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	PrintLargerHalf2(arr2, 8)
	arr3 := []int{8, 7, 6, 5, 7, 5, 2, 1}
	PrintLargerHalf3(arr3, 8)
}

func SortK(arr []int, size int, k int) {
	pq := &MinHeap{}
	i := 0
	for ; i < k; i++ {
		heap.Push(pq, arr[i])
	}
	output := make([]int, size)
	index := 0
	for i = k; i < size; i++ {
		output[index] = heap.Pop(pq).(int)
		index += 1
		heap.Push(pq, arr[i])
	}
	for pq.Len() > 0 {
		output[index] = heap.Pop(pq).(int)
		index += 1
	}
	for i = 0; i < size; i++ {
		arr[i] = output[i]
	}
}

func Main4() {
	k := 3
	arr := []int{1, 5, 4, 10, 50, 9}
	size := len(arr)
	SortK(arr, size, k)
	for i := 0; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
}
func PrintLargerHalf(arr []int, size int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
func PrintLargerHalf2(arr []int, size int) {
	pq := &MinHeap{}
	for i := 0; i < size; i++ {
		heap.Push(pq, arr[i])
	}
	for i := 0; i < size/2; i++ {
		heap.Pop(pq)
	}
	fmt.Println(pq)
}
func PrintLargerHalf3(arr []int, size int) {
	QuickSelectUtil(arr, 0, size-1, size/2)
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
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
			Swap(arr, upper, lower)
		}
	}
	Swap(arr, upper, start)
	if k < upper {
		QuickSelectUtil(arr, start, upper-1, k)
	}
	if k > upper {
		QuickSelectUtil(arr, upper+1, stop, k)
	}
}

func Swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	Main0()
	Main1()
	Main2()
	Main3()
	Main4()
}

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

func (pq MinHeap) Print() {
	fmt.Println(pq)
}

func (pq MinHeap) Empty() bool {
	return (pq.Len() == 0)
}

func (pq MinHeap) Peek() interface{} {
	return pq[0]
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

func (pq MaxHeap) Print() {
	fmt.Println(pq)
}

func (pq MaxHeap) Empty() bool {
	return (pq.Len() == 0)
}

func (pq MaxHeap) Peek() interface{} {
	return pq[0]
}