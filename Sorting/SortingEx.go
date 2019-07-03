package main

import (
	"fmt"
	"sort"
	"math"
)

func PrintArr(arr []int, count int) {
	fmt.Print("[")
	for i := 0; i < count; i++ {
		fmt.Print(" " , arr[i])
	}
	fmt.Print(" ]\n")
}

func swap(arr []int, x int, y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}

func Partition01(arr []int, size int) int {
	left := 0
	right := size - 1
	count := 0
	for left < right {
		for arr[left] == 0 {
			left += 1
		}

		for arr[right] == 1 {
			right -= 1
		}

		if (left < right) {
			swap(arr, left, right)
			count += 1
		}
	}
	return count
}

func Partition012(arr []int, size int) {
	left := 0
	right := size - 1
	i := 0
	for i <= right {
		if (arr[i] == 0) {
			swap(arr, i, left)
			i += 1
			left += 1
		} else if (arr[i] == 2) {
			swap(arr, i, right)
			right -= 1
		} else {
			i += 1
		}
	}
}

// Testing code
func main1() {
	arr := []int{ 0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1 }
	Partition01(arr, len(arr))
	fmt.Println(arr)
	arr2 := []int{ 0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1 }
	Partition012(arr2, len(arr))
	fmt.Println(arr2)
}

func RangePartition(arr []int, size int, lower int, higher int) {
	start := 0
	end := size - 1
	i := 0
	for i <= end {
		if (arr[i] < lower) {
			swap(arr, i, start)
			i += 1
			start += 1
		} else if (arr[i] > higher) {
			swap(arr, i, end)
			end -= 1
		} else {
			i += 1
		}
	}
}

// Testing code
func main2() {
	arr  := []int{ 1, 21, 2, 20, 3, 19, 4, 18, 5, 17, 6, 16, 7, 15, 8, 14, 9, 13, 10, 12, 11 }
	RangePartition(arr, len(arr), 9, 12)
	fmt.Println(arr)
}

func minSwaps(arr []int, size int, val int) int {
	swapCount := 0
	first := 0
	second := size - 1
	var temp int
	for first < second {
		if (arr[first] < val) {
			first += 1
		} else if (arr[second] >= val) {
			second -= 1
		} else {
			temp = arr[first]
			arr[first] = arr[second]
			arr[second] = temp
			swapCount += 1
		}
	}
	return swapCount;
}

func main3(){
	arr := []int {2,7,5,6,1,3,4,9,10,8}
	val := 5
	fmt.Println(minSwaps(arr, 10, val))
}

func seperateEvenAndOdd(data []int, size int) {
	left := 0 
	right := size - 1
	for left < right {
		if (data[left] % 2 == 0) {
			left++
		} else if (data[right] % 2 == 1) {
			right--
		} else {
			swap(data, left, right)
			left++
			right--
		}
	}
}

func main4(){
	arr := []int{2,7,5,6,1,3,4,9,10,8}
	seperateEvenAndOdd(arr, 10)
	fmt.Println(arr)
}

func absMore(value1 int, value2 int, reference int)  bool{
	return (math.Abs(float64(value1 - reference)) > math.Abs(float64(value2 - reference)))
}

func absBubbleSort(arr []int, size int, reference int) {
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < (size - i - 1); j++ {
			if (absMore(arr[j], arr[j + 1], reference)) {
				swap(arr, j, j + 1)
			}
		}
	}
}

// Testing code
func main5() {
	array := []int{ 9, 1, 8, 2, 7, 3, 6, 4, 5 }
	reference := 5
	absBubbleSort(array, len(array), reference)
	fmt.Println(array)
}

func EqMore(value1 int, value2 int, A int) bool {
	value1 = A * value1 * value1
	value2 = A * value2 * value2
	return value1 > value2
}

func SortByOrder(arr []int, size int, arr2 []int, size2 int) {	
	ht := make(map[int]int)
	for i := 0; i < size; i++ {
		value, ok := ht[arr[i]]
		if ok {
			ht[arr[i]] = value + 1
		} else {
			ht[arr[i]] = 1
		}
	}

	for j := 0; j < size2; j++ {
		value, ok := ht[arr2[j]]
		if ok {
			for k := 0; k < value; k++ {
				fmt.Print(arr2[j], " ")
			}
			delete(ht,arr2[j])
		}
	}

	for i := 0; i < size; i++ {
		value, ok := ht[arr[i]]
		if ok {
			for k := 0; k < value; k++ {
				fmt.Print(arr[i], " ")
			}
			delete(ht, arr[i])
		}
	}
}

// Testing code
func main6() {
	arr := []int { 2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8 }
	arr2 := []int{ 2, 1, 8, 3 }
	SortByOrder(arr, len(arr), arr2, len(arr2))
}

func ArrayReduction(arr []int, size int) {
	sort.Ints(arr)
	count := 1
	reduction := arr[0]

	for i := 0; i < size; i++ {
		if (arr[i] - reduction > 0) {
			reduction = arr[i]
			count += 1
		}
	}
	fmt.Println("Total number of reductions ", count)
}

// Testing code
func main7() {
	arr := []int { 5, 1, 1, 1, 2, 3, 5 }
	ArrayReduction(arr, len(arr))
}

func merge(arr1 []int, size1 int, arr2 []int, size2 int) {
	index := 0
	for index < size1 {
		if (arr1[index] <= arr2[0]) {
			index += 1
		} else {
			// always first element of arr2 is compared.
			arr1[index], arr2[0] =  arr2[0], arr1[index]
			index += 1
			// After swap arr2 may be unsorted.
			// Insertion of the element in proper sorted position.
			for i := 0; i < (size2 - 1); i++ {
				if (arr2[i] < arr2[i + 1]) {
					break
				}
				arr2[i], arr2[i + 1] = arr2[i + 1], arr2[i]
			}
		}
	}
}

// Testing code.
func main8() {
	arr1 := []int{ 1, 5, 9, 10, 15, 20 }
	arr2 := []int{ 2, 3, 8, 13 }
	merge(arr1, len(arr1), arr2, len(arr2))
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func checkReverse(arr []int, size int) bool {
	start := -1
	stop := -1
	for i := 0; i < (size - 1); i++ {
		if (arr[i] > arr[i + 1]) {
			start = i
			break
		}
	}

	if (start == -1) {
		return true
	}

	for i := start; i < (size - 1); i++ {
		if (arr[i] < arr[i + 1]) {
			stop = i
			break
		}
	}

	if (stop == -1) {
		return true
	}

	// increasing property
	// after reversal the sub array should fit in the array.
	if (arr[start - 1] > arr[stop] || arr[stop + 1] < arr[start]) {
		return false
	}

	for i := stop + 1; i < size - 1; i++ {
		if (arr[i] > arr[i + 1]) {
			return false
		}
	}
	return true
}

func main9() {
	arr := []int {1, 3, 8, 5, 4, 3, 10, 11, 12, 18, 28}
	fmt.Println("checkReverse : ", checkReverse(arr, len(arr)))
}

func min(X int, Y int) int {
	if (X < Y) {
		return X
	}
	return Y
}

func UnionIntersectionSorted(arr1 []int, size1 int, arr2 []int, size2 int) {
	first := 0
	second := 0
	unionArr := make([]int, size1 + size2)
	interArr := make([]int, min(size1, size2))
	uIndex := 0
	iIndex := 0

	for first < size1 && second < size2 {
		if (arr1[first] == arr2[second]) {
			unionArr[uIndex] = arr1[first]
			uIndex++
			interArr[iIndex] = arr1[first]
			iIndex++
			first += 1
			second += 1
		} else if (arr1[first] < arr2[second]) {
			unionArr[uIndex] = arr1[first]
			uIndex++
			first += 1
		} else {
			unionArr[uIndex] = arr2[second]
			uIndex++
			second += 1
		}
	}

	for first < size1 {
		unionArr[uIndex] = arr1[first]
		uIndex++
		first += 1
	}

	for second < size2 {
		unionArr[uIndex] = arr2[second]
		uIndex++
		second += 1
	}
	PrintArr(unionArr, uIndex)
	PrintArr(interArr, iIndex)
}

func UnionIntersectionUnsorted(arr1 []int, size1 int, arr2 []int, size2 int) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	UnionIntersectionSorted(arr1, size1, arr2, size2)
}

func main10() {
	arr1 := []int { 1, 11, 2, 3, 14, 5, 6, 8, 9 }
	arr2 := []int { 2, 4, 5, 12, 7, 8, 13, 10 }
	UnionIntersectionUnsorted(arr1, len(arr1), arr2, len(arr2))
}

func main(){
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
}