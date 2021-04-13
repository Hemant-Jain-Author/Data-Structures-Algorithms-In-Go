package main

import (
	"fmt"
	"sort"
	"math"
)

type Heap struct {
	size  int
	arr   []int
	isMin bool
}

func CreateHeap(isMin bool, args ...[]int) *Heap {
	arr := []int{1}
	size := 0
	if len(args) > 0 {
		arrInput := args[0]
		arr = append(arr, arrInput...)
		size = len(arrInput)
	}
	h := &Heap{size: size, arr: arr, isMin: isMin}
	for i := (size / 2); i > 0; i-- {
		h.proclateDown(i)
	}

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
	if h.IsEmpty() {
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

func (h *Heap) IsEmpty() bool {
	return (h.size == 0)
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() int {
	if h.IsEmpty() {
		fmt.Println("Heap empty exception.")
		return 0
	}
	return h.arr[1]
}

func (h *Heap) Print() {
	fmt.Print("Printing Heap size :", h.size, " :: ")
	for i := 1; i <= h.size; i++ {
		fmt.Print(" ", h.arr[i])
	}
	fmt.Println()
}

func main1() {
    hp := CreateHeap(true)
	hp.Print();
    hp.Add(1)
    hp.Add(9)
    hp.Add(6)
    hp.Add(7)
    hp.Print()
    for  hp.IsEmpty() == false {
        fmt.Println(hp.Remove())
    }
}
/*
Printing Heap size :0 :: 
Printing Heap size :4 ::  1 7 6 9
1
6
7
9
*/

func main2() {
    a := []int{1, 0, 2, 4, 5, 3}
    hp := CreateHeap(true, a)// Min Heap
    hp.Print()
    for hp.IsEmpty() == false {
        fmt.Println(hp.Remove())
    }
}
/*
Printing Heap size :6 ::  0 1 2 4 5 3
0
1
2
3
4
5
*/

func main3() {
    a := []int{1, 0, 2, 4, 5, 3}
    hp := CreateHeap(false, a)// Max Heap
    hp.Print()
    for hp.IsEmpty() == false {
        fmt.Println(hp.Remove())
    }
}
/*
Printing Heap size :6 ::  5 4 3 1 0 2
5
4
3
2
1
0
*/

func HeapSort(arrInput []int) {
	hp := CreateHeap(true, arrInput)
	n := len(arrInput)
	for i := 0; i < n; i++ {
		arrInput[i] = hp.Remove()
	}
}

//Testing Code 
func main4() {
	a := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
	HeapSort(a)
	fmt.Println("value after heap sort::", a)
}

/*
value after heap sort:: [-1 1 2 3 4 5 6 7 8 9]
*/

func IsMinHeap(arr[] int) bool {
	var lchild, rchild int
	size := len(arr)
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

func IsMaxHeap(arr[] int) bool {
	var lchild, rchild int
	size := len(arr)
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

//Testing Code 
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
	min := CreateHeap(true)
	max := CreateHeap(false)

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

//Testing Code 
func main6() {
	arr := []int{1, 9, 2, 8, 3, 7}
	hp := NewMedianHeap()

	for i := 0; i < 6; i++ {
		hp.insert(arr[i])
		fmt.Println("Median after insertion of ", arr[i], " is ", hp.getMedian())
	}
}
/*
Median after insertion of  1  is  1
Median after insertion of  9  is  5
Median after insertion of  2  is  2
Median after insertion of  8  is  5
Median after insertion of  3  is  3
Median after insertion of  7  is  5
*/
func KthSmallest(arr[] int, size int, k int) int {
	sort.Ints(arr)
	return arr[k - 1]
}

func KthSmallest2(arr[] int, size int, k int) int {
	i := 0
	value := 0
	hp := CreateHeap(true)
	for i = 0; i < size; i++ {
		hp.Add(arr[i])
	}
	i = 0
	for i < size && i < k {
		value = hp.Remove()
		i += 1
	}
	return value
}

//Testing Code 
func main7() {
	arr :=  []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest :: " , KthSmallest(arr, len(arr), 3))
	arr2 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest :: " , KthSmallest2(arr2, len(arr2), 3))
}
/*
Kth Smallest ::  5
Kth Smallest ::  5
*/

func KSmallestProduct(arr[] int, size int, k int) int {
	sort.Ints(arr)
	product := 1
	for i := 0; i < k; i++ {
		product *= arr[i]
	}
	return product
}

func KSmallestProduct2(arr[] int, size int, k int) int {
	hp := CreateHeap(true)
	i := 0
	product := 1
	for i = 0; i < size; i++ {
		hp.Add(arr[i])
	}
	i = 0
	for (i < size && i < k) {
		product *= hp.Remove()
		i += 1
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

func KSmallestProduct3(arr[] int, size int, k int) int {
	QuickSelectUtil(arr, 0, size - 1, k)
	product := 1
	for i := 0; i < k; i++ {
		product *= arr[i]
	}
	return product
}

//Testing Code 
func main8() {
	arr := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest product:: " , KSmallestProduct(arr, 8, 3))
	arr2 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest product:: " , KSmallestProduct2(arr2, 8, 3))
	arr3 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	fmt.Println("Kth Smallest product:: " , KSmallestProduct3(arr3, 8, 3))
}
/*
Kth Smallest product::  10
Kth Smallest product::  10
Kth Smallest product::  10
*/

func PrintLargerHalf(arr[] int, size int) {
	sort.Ints(arr) // , size, 1)
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func PrintLargerHalf2(arr[] int, size int) {
	hp := CreateHeap(false)
	for i := 0; i < size; i++ {
		hp.Add(arr[i])
	}

	for i := 0; i < size / 2; i++ {
		fmt.Print(hp.Remove(), " ")
	}
	fmt.Println()
}

func PrintLargerHalf3(arr[] int, size int) {
	QuickSelectUtil(arr, 0, size - 1, size / 2)
	for i := size / 2; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

//Testing Code 
func main9() {
	arr := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	PrintLargerHalf(arr, 8)
	arr2 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	PrintLargerHalf2(arr2, 8)
	arr3 := []int { 8, 7, 6, 5, 7, 5, 2, 1 }
	PrintLargerHalf3(arr3, 8)
}

/*
6 7 7 8 
8 7 7 6 
6 7 7 8
*/

func sortK(arr[] int, size int, k int) {
	hp := CreateHeap(true)
	i := 0
	for i = 0; i < k; i++ {
		hp.Add(arr[i])
	}

	output := make([]int, size)
	index := 0

	for i = k; i < size; i++ {
		output[index] = hp.Remove()
		index++
		hp.Add(arr[i])
	}

	for (hp.Size() > 0) {
		output[index] = hp.Remove()
		index++
	}

	for i = k; i < size; i++ {
		arr[i] = output[i]
	}
	fmt.Println(output)
}

// Testing Code
func main10() {
	k := 3
	arr := []int { 1, 5, 4, 10, 50, 9 }
	size := len(arr)
	sortK(arr, size, k)
}

/*
[1 4 5 9 10 50]
*/

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
	fmt.Println("Total :" , total)
	return total;
}

func ChotaBhim3(cups []int, size int) int {
	time := 60;
	hp := CreateHeap(false)
	i := 0;
	for i = 0; i < size; i++ {
		hp.Add(cups[i])
	}

	total := 0;
	var value int
	for (time > 0) {
		value = hp.Remove()
		total += value;
		value = int(math.Ceil(float64(value) / 2.0))
		hp.Add(value)
		time -= 1;
	}
	fmt.Println("Total :" , total)
	return total;
}

//Testing Code 
func main11() {
	cups := []int { 2, 1, 7, 4, 2 }
	ChotaBhim(cups, len(cups))
	cups3 := []int { 2, 1, 7, 4, 2 }
	ChotaBhim3(cups3, len(cups))
}

/*
Total : 76
Total : 76
*/

func JoinRopes(ropes []int , size int) int {
	sort.Slice(ropes, func(i, j int) bool {
	    return ropes[i] > ropes[j]
	})

	total := 0
	value := 0
	var index int
	length := size

	for (length >= 2) {
		value = ropes[length - 1] + ropes[length - 2];
		total += value
		index = length - 2

		for (index > 0 && ropes[index - 1] < value) {
			ropes[index] = ropes[index - 1]
			index -= 1
		}
		ropes[index] = value
		length--
	}
	fmt.Println("Total :" , total)
	return total
}

func JoinRopes2(ropes []int, size int) int {
	hp := CreateHeap(true)
	i := 0
	for i = 0; i < size; i++ {
		hp.Add(ropes[i])
	}

	total := 0
	value := 0
	for (hp.Size() > 1) {
		value = hp.Remove()
		value += hp.Remove()
		hp.Add(value)
		total += value
	}
	fmt.Println("Total :" , total)
	return total
}

//Testing Code 
func main12() {
	ropes := []int { 2, 1, 7, 4, 2 }
	JoinRopes(ropes, len(ropes))
	rope2 := []int { 2, 1, 7, 4, 2 }
	JoinRopes2(rope2, len(rope2))
}

/*
Total : 33
Total : 33
*/

func kthLargestStream(k int) {
	hp := CreateHeap(true)
	size := 0
	data := 0

	for (true) {
		fmt.Println("Enter data: ")
		fmt.Scanf("%d", &data)

		if (size < k - 1) {
			hp.Add(data)
		} else {
			if (size == k - 1) {
				hp.Add(data)
			} else if (hp.Peek() < data) {
				hp.Add(data)
				hp.Remove()
			}
			fmt.Println("Kth larges element is :: ", hp , hp.Peek())
		}
		size += 1;
	}
}

func main13() {
	kthLargestStream(3)
}

func main(){
	//main1()
	//main2()
	//main3()
	//main4()
	//main5()
	//main6()
	//main7()
	//main8()
	//main9()
	//main10()
	//main11()
	main12()
	//main13()
	
}