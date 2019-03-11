package main

import (
	"fmt"
	"sort"
	"math"
)

func main1() {
	a := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
	pq := NewPQueue(false)
	for _, val := range a {
		pq.Add(val, val)
	}

	for _, _ = range a {
		value := pq.Remove()
		fmt.Print(value, " ")
	}

	fmt.Println()
	pq2 := NewPQueue(false)
	for _, val := range a {
		pq2.Add(val, val)
	}
	for _, _ = range a {
		value := pq2.Remove()
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func KthSmallest(arr[] int, size int, k int) int {
	sort.Ints(arr)
	return arr[k - 1]
}

func KthSmallest2(arr[] int, size int, k int) int {
	i := 0
	value := 0
	pq := NewPQueue(false)
	for i = 0; i < size; i++ {
		pq.Add(arr[i], arr[i])
	}

	for i < size && i < k {
		value = pq.Remove().(int)
		i += 1
	}
	return value
}

func isMinHeap(arr[] int, size int) bool {
	var lchild, rchild int
	// last element index size - 1
	for parent := 0; parent < (size / 2 + 1); parent++ {
		lchild = parent * 2 + 1
		rchild = parent * 2 + 2
		// heap property check.
		if (((lchild < size) && (arr[parent] > arr[lchild])) ||
		((rchild < size) && (arr[parent] > arr[rchild]))) {
			return false
		}
	}
	return true
}

func isMaxHeap(arr[] int, size int) bool {
	var lchild, rchild int
	// last element index size - 1
	for parent := 0; parent < (size / 2 + 1); parent++ {
		lchild = parent * 2 + 1
		rchild = lchild + 1
		// heap property check.
		if (((lchild < size) && (arr[parent] < arr[lchild])) ||
		((rchild < size) && (arr[parent] < arr[rchild]))) {
			return false
		}
	}
	return true
}

func main2() {
	arr :=  []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest :: " , KthSmallest(arr, len(arr), 3))
	arr2 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest :: " , KthSmallest2(arr2, len(arr2), 3))
	arr3 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("isMaxHeap :: " , isMaxHeap(arr3, len(arr3)))
	arr4 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	sort.Ints(arr4)
	fmt.Println("isMinHeap :: " , isMinHeap(arr4, len(arr4)))
}

func KSmallestProduct(arr[] int, size int, k int) int {
	sort.Ints(arr)
	product := 1
	for i := 0; i < k; i++ {
		product *= arr[i]
	}
	return product
}

func swap(arr[] int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func QuickSelectUtil(arr[] int, lower int, upper int, k int) {
	if (upper <= lower) {
		return
	}

	pivot := arr[lower]
	start := lower
	stop := upper

	for (lower < upper) {
		for (lower < upper && arr[lower] <= pivot) {
			lower++
		}
		for (lower <= upper && arr[upper] > pivot) {
			upper--
		}
		if (lower < upper) {
			swap(arr, upper, lower)
		}
	}

	swap(arr, upper, start) // upper is the pivot position
	if (k < upper) {
		QuickSelectUtil(arr, start, upper - 1, k) // pivot -1 is the upper for left sub array.
	}
	if (k > upper) {
		QuickSelectUtil(arr, upper + 1, stop, k) // pivot + 1 is the lower for right sub array.
	}
}
func KSmallestProduct2(arr[] int, size int, k int) int {
	pq := NewPQueue(false)
	i := 0
	product := 1
	for i = 0; i < size; i++ {
		pq.Add(arr[i], arr[i])
	}

	for (i < size && i < k) {
		product *= pq.Remove().(int)
		i += 1
	}
	return product
}

func KSmallestProduct3(arr[] int, size int, k int) int {
	QuickSelectUtil(arr, 0, size - 1, k)
	product := 1
	for i := 0; i < k; i++ {
		product *= arr[i]
	}
	return product
}

func main3() {
	arr := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest product:: " , KSmallestProduct(arr, 8, 3))
	arr2 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest product:: " , KSmallestProduct2(arr2, 8, 3))
	arr3 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest product:: " , KSmallestProduct3(arr3, 8, 3))
}

func PrintLargerHalf(arr[] int, size int) {
	sort.Ints(arr) // , size, 1)
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
}

func PrintLargerHalf2(arr[] int, size int) {
	pq := NewPQueue(false)
	for i := 0; i < size; i++ {
		pq.Add(arr[i], arr[i])
	}

	for i := 0; i < size / 2; i++ {
		pq.Remove()
	}
	fmt.Println(pq)
}

func PrintLargerHalf3(arr[] int, size int) {
	QuickSelectUtil(arr, 0, size - 1, size / 2)
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
}

func main4() {
	arr := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	PrintLargerHalf(arr, 8)
	arr2 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	PrintLargerHalf2(arr2, 8)
	arr3 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	PrintLargerHalf3(arr3, 8)
}

func sortK(arr[] int, size int, k int) {
	pq := NewPQueue(true)
	i := 0
	for i = 0; i < k; i++ {
		pq.Add(arr[i], arr[i])
	}

	output := make([]int, size)
	index := 0

	for i = k; i < size; i++ {
		output[index] = pq.Remove().(int)
		index++
		pq.Add(arr[i], arr[i])
	}

	for (pq.Size() > 0) {
		output[index] = pq.Remove().(int)
		index++
	}

	for i = k; i < size; i++ {
		arr[i] = output[i]
	}
	fmt.Println(output)
}

// Testing Code
func main5() {
	k := 3
	arr := []int { 1, 5, 4, 10, 50, 9 }
	size := len(arr)
	sortK(arr, size, k)
}

func ChotaBhim(cups []int, size int) int {
	time := 60
	sort.Ints(cups)	
	total := 0
	var index, temp int
	for (time > 0) 	{
		total += cups[0]
		cups[0] = int(math.Ceil(float64(cups[0]) / 2.0))
		index = 0
		temp = cups[0]
		for (index < size - 1 && temp < cups[index + 1]) {
			cups[index] = cups[index + 1]
			index += 1
		}
		cups[index] = temp
		time -= 1
	}
	fmt.Println("Total " , total)
	return total;
}

func ChotaBhim3(cups []int, size int) int {
	time := 60;
	pq := NewPQueue(false)
	i := 0;
	for i = 0; i < size; i++ {
		pq.Add(cups[i], cups[i])
	}

	total := 0;
	var value int
	for (time > 0) {
		value = pq.Remove().(int)
		total += value;
		value = int(math.Ceil(float64(value) / 2.0))
		pq.Add(value, value)
		time -= 1;
	}
	fmt.Println("Total : " , total)
	return total;
}

func JoinRopes(ropes []int , size int) int {

	sort.Slice(ropes, func(i, j int) bool {
	    return ropes[i] > ropes[j]
	})


	fmt.Println(ropes)
	total := 0
	value := 0
	var index int
	length := size

	for (length >= 2) {
		value = ropes[length - 1] + ropes[length - 2];
		fmt.Print("[", value, "]")
		total += value
		index = length - 2

		for (index > 0 && ropes[index - 1] < value) {
			ropes[index] = ropes[index - 1]
			index -= 1
		}
		ropes[index] = value
		length--
	}
	fmt.Println("Total : " , total)
	return total
}

func JoinRopes2(ropes []int, size int) int {
	pq := NewPQueue(true)
	i := 0
	for i = 0; i < size; i++ {
		pq.Add(ropes[i], ropes[i])
	}

	total := 0
	value := 0
	for (pq.Size() > 1) {
		value = pq.Remove().(int)
		value += pq.Remove().(int)
		pq.Add(value, value)
		total += value
	}
	fmt.Println("Total : " , total)
	return total
}

func main6() {
	cups := []int { 2, 1, 7, 4, 2 }
	ChotaBhim(cups, len(cups))
	cups3 := []int { 2, 1, 7, 4, 2 }
	ChotaBhim3(cups3, len(cups))

	ropes := []int { 2, 1, 7, 4, 2 }
	JoinRopes(ropes, len(ropes))
	rope2 := []int { 2, 1, 7, 4, 2 }
	JoinRopes2(rope2, len(rope2))
}
/*
func kthLargestStream(k int) int {
	pq := NewPQueue(false)
	size := 0
	data := 0
	for (true) {
		fmt.Println("Enter data: ")
		data = fmt.readline()

		if (size < k - 1) {
			pq.Add(data, data)
		} else {
			if (size == k - 1) {
				pq.Add(data, data)
			} else if (pq.peek() < data) {
				pq.Add(data, data)
				pq.Remove().(int)
			}
			fmt.Println("Kth larges element is :: " + pq.peek())
		}
		size += 1;
	}
}

func Main888() {
	kthLargestStream(3)
}
*/

func main(){
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
}

//*********************
type item struct {
	value    interface{}
	priority int
}

type PQueue struct {
	items []*item
	Count int
	isMin bool
}

func newItem(value interface{}, priority int) *item {
	return &item{
		value:    value,
		priority: priority,
	}
}

func NewPQueue(isMin bool) *PQueue {
	items := make([]*item, 1)
	items[0] = nil // Heap queue first element should always be nil

	return &PQueue{
		items: items,
		Count: 0,
		isMin: isMin,
	}
}

func (pq *PQueue) comp(i, j int) bool { // always i < j in use
	if pq.isMin == true {
		return pq.items[i].priority > pq.items[j].priority // swaps for min heap
	}
	return pq.items[i].priority < pq.items[j].priority // swap for max heap.
}

func (pq *PQueue) proclateDown(position int) {
	lChild := 2 * position
	rChild := lChild + 1
	small := -1

	if lChild <= pq.Count {
		small = lChild
	}

	if rChild <= pq.Count && pq.comp(lChild, rChild) {
		small = rChild
	}

	if small != -1 && pq.comp(position, small) {
		pq.items[position], pq.items[small] = pq.items[small], pq.items[position]
		pq.proclateDown(small)
	}
}

func (pq *PQueue) proclateUp(position int) {
	parent := position / 2
	if parent == 0 {
		return
	}

	if pq.comp(parent, position) {
		pq.items[position], pq.items[parent] = pq.items[parent], pq.items[position]
		pq.proclateUp(parent)
	}
}

func (pq *PQueue) Add(value interface{}, priority int) {
	item := newItem(value, priority)
	pq.items = append(pq.items, item)
	pq.Count++
	pq.proclateUp(pq.Count)
}

func (pq *PQueue) Print() {
	n := pq.Count
	for i := 1; i <= n; i++ {
		fmt.Print(*(pq.items[i]), " ")
	}
	fmt.Println()
}

func (pq *PQueue) Remove() (interface{}) {
	if pq.Empty() {
		fmt.Println("Heap Empty Error.")
		return 0
	}

	value := pq.items[1].value
	pq.items[1] = pq.items[pq.Count]
	pq.items = pq.items[0:pq.Count]
	pq.Count--
	pq.proclateDown(1)
	return value
}

func (pq *PQueue) Empty() bool {
	return (pq.Count == 0)
}

func (pq *PQueue) Size() int {
	return pq.Count
}

func (pq *PQueue) Peek() (interface{}) {
	if pq.Empty() {
		fmt.Println("Heap empty exception.")
		return 0
	}
	return pq.items[1].value
}

//*********************