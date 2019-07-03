package main

import (
	"fmt"
	"sort"
)

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func SumArray(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}

func main1() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sum of all the values in array:", SumArray(data))
}

func function2() {
	fmt.Println("fun2 line 1")
}

func function1() {
	fmt.Println("fun1 line 1")
	function2()
	fmt.Println("fun1 line 2")
}

func main2() {
	fmt.Println("main line 1")
	function1()
	fmt.Println("main line 2")
}

func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func BinarySearch(data []int, value int) bool {
	size := len(data)
	var mid int
	low := 0
	high := size - 1
	for low <= high {
		mid = low + (high-low)/2 // To afunc the overflow
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

func main3() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("SequentialSearch:", SequentialSearch(arr, 7))
	fmt.Println("SequentialSearch:", SequentialSearch(arr, 8))
	fmt.Println("BinarySearch:", BinarySearch(arr, 7))
	fmt.Println("BinarySearch:", BinarySearch(arr, 8))
}

func RotateArray(data []int, k int) {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
}

func ReverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func ReverseArray2(data []int) {
	i := 0
	j := len(data) - 1
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func main4() {
	arr := []int{1, 2, 3, 4, 5, 6}
    fmt.Println("Input array :", arr)
	RotateArray(arr, 2)
    fmt.Println("Rotated array :", arr)
}

func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0

	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func main5() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	fmt.Println("Max sub array sum :", MaxSubArraySum(data))
}

func WaveArray(arr []int) {
	size := len(arr)
	/* Odd elements are lesser then even elements. */
	for i := 1; i < size; i += 2 {
		if (i-1) >= 0 && arr[i] > arr[i-1] {
			swap(arr, i, i-1)
		}
		if (i+1) < size && arr[i] > arr[i+1] {
			swap(arr, i, i+1)
		}
	}
}

func WaveArray2(arr []int) {
	size := len(arr)
	sort.Ints(arr)
	for i := 0; i < size-1; i += 2 {
		swap(arr, i, i+1)
	}
}

/* Testing code */
func main6() {
	arr := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	WaveArray(arr)
    fmt.Println(arr)

	arr2 := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	WaveArray2(arr2)
    fmt.Println(arr2)
}


func indexArray(arr []int, size int) {
	for i := 0; i < size; i++ {
		curr := i
		value := -1

		/* swaps to move elements in proper position. */
		for arr[curr] != -1 && arr[curr] != curr {
			temp := arr[curr]
			arr[curr] = value
			value = temp
			curr = temp
		}

		/* check if some swaps happened. */
		if value != -1 {
			arr[curr] = value
		}
	}
}

func indexArray2(arr []int, size int) {
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != -1 && arr[i] != i {
			/* swap arr[i] and arr[arr[i]] */
			temp = arr[i]
			arr[i] = arr[temp]
			arr[temp] = temp
		}
	}
}

/* Testing code */
func main7() {
	arr := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	size := len(arr)
	indexArray2(arr, size)
    fmt.Println(arr)
	arr2 := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	size2 := len(arr2)
	indexArray(arr2, size2)
    fmt.Println(arr2)
}

func Sort1toN(arr []int, size int) {
	curr, value, next := 0, 0, 0
	for i := 0; i < size; i++ {
		curr = i
		value = -1
		/* swaps to move elements in proper position. */
		for curr >= 0 && curr < size && arr[curr] != curr+1 {
			next = arr[curr]
			arr[curr] = value
			value = next
			curr = next - 1
		}
	}
}

func Sort1toN2(arr []int, size int) {
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 1 {
			temp = arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
		}
	}
}

func main8() {
	arr := []int{8, 5, 6, 1, 9, 3, 2, 7, 4, 10}
	size := len(arr)
	Sort1toN2(arr, size)
    fmt.Println(arr)

	arr2 := []int{8, 5, 6, 1, 9, 3, 2, 7, 4, 10}
	size2 := len(arr2)
	Sort1toN(arr2, size2)
    fmt.Println(arr2)
}


func SmallestPositiveMissingNumber(arr []int, size int) int {
	found := 0
	for i := 1; i < size+1; i++ {
		found = 0
		for j := 0; j < size; j++ {
			if arr[j] == i {
				found = 1
				break
			}
		}
		if found == 0 {
			return i
		}
	}
	return -1
}

func SmallestPositiveMissingNumber2(arr []int, size int) int {
	hs := make(map[int]int)
	for i := 0; i < size; i++ {
		hs[arr[i]] = 1
	}
	for i := 1; i < size+1; i++ {
		_, ok := hs[i]
		if ok == false {
			return i
		}
	}
	return -1
}

func SmallestPositiveMissingNumber3(arr []int, size int) int {
	aux := make([]int, size)

	for index, _ := range aux {
		aux[index] = -1
	}

	for i := 0; i < size; i++ {
		if arr[i] > 0 && arr[i] <= size {
			aux[arr[i]-1] = arr[i]
		}
	}
	for i := 0; i < size; i++ {
		if aux[i] != i+1 {
			return i + 1
		}
	}
	return -1
}

func SmallestPositiveMissingNumber4(arr []int, size int) int {
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 0 && arr[i] <= size {
			temp = arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
		}
	}
	for i := 0; i < size; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	return -1
}

func main9() {
	arr := []int{8, 5, 6, 1, 9, 11, 2, 7, 4, 10}
	size := len(arr)

	fmt.Println("Missing Number :", SmallestPositiveMissingNumber(arr, size))
	fmt.Println("Missing Number :", SmallestPositiveMissingNumber2(arr, size))
	fmt.Println("Missing Number :", SmallestPositiveMissingNumber3(arr, size))
	fmt.Println("Missing Number :", SmallestPositiveMissingNumber4(arr, size))
}

func MaxMinArr(arr []int, size int) {
	aux := make([]int, size)
	copy(aux, arr)
	start := 0
	stop := size - 1
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			arr[i] = aux[stop]
			stop -= 1
		} else {
			arr[i] = aux[start]
			start += 1
		}
	}
}

func ReverseArr(arr []int, start int, stop int) {
	for start < stop {
		swap(arr, start, stop)
		start += 1
		stop -= 1
	}
}

func MaxMinArr2(arr []int, size int) {
	for i := 0; i < (size - 1); i++ {
		ReverseArr(arr, i, size-1)
	}
}

/* Testing code */
func main10() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	size := len(arr)
	MaxMinArr(arr, size)
    fmt.Println(arr)

	arr2 := []int{1, 2, 3, 4, 5, 6, 7}
	size2 := len(arr2)
	MaxMinArr2(arr2, size2)
    fmt.Println(arr2)
}

func maxCircularSum(arr []int, size int) int {
	sumAll := 0
	currVal := 0

	for i := 0; i < size; i++ {
		sumAll += arr[i]
		currVal += (i * arr[i])
	}
	maxVal := currVal
	for i := 1; i < size; i++ {
		currVal = (currVal + sumAll) - (size * arr[size-i])
		if currVal > maxVal {
			maxVal = currVal
		}
	}
	return maxVal
}

/* Testing code */
func main11() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("MaxCirculrSum: ", maxCircularSum(arr, len(arr)))
}

func ArrayIndexMaxDiff(arr []int, size int) int {
	maxDiff := -1
	j := 0
	for i := 0; i < size; i++ {
		j = size - 1
		for j > i {
			if arr[j] > arr[i] {
				if maxDiff < (j - i) {
					maxDiff = (j - i)
				}
				break
			}
			j -= 1
		}
	}
	return maxDiff
}

func ArrayIndexMaxDiff2(arr []int, size int) int {
	leftMin := make([]int, size)
	rightMax := make([]int, size)
	leftMin[0] = arr[0]
	i, j := 0, 0
	var maxDiff = 0
	for i = 1; i < size; i++ {
		if leftMin[i-1] < arr[i] {
			leftMin[i] = leftMin[i-1]
		} else {
			leftMin[i] = arr[i]
		}
	}
	rightMax[size-1] = arr[size-1]
	for i = size - 2; i >= 0; i-- {
		if rightMax[i+1] > arr[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = arr[i]
		}
	}
	i = 0
	j = 0
	maxDiff = -1
	for j < size && i < size {
		if leftMin[i] < rightMax[j] {
			if maxDiff < j-i {
				maxDiff = j - i
			}
			j = j + 1
		} else {
			i = i + 1
		}
	}
	return maxDiff
}

func main12() {
	arr := []int{33, 9, 10, 3, 2, 60, 30, 33, 1}
	fmt.Println("ArrayIndexMaxDiff : ", ArrayIndexMaxDiff(arr, len(arr)))
	fmt.Println("ArrayIndexMaxDiff : ", ArrayIndexMaxDiff2(arr, len(arr)))
}

func maxPathSum(arr1 []int, size1 int, arr2 []int, size2 int) int {
	i := 0
	j := 0
	result := 0
	sum1 := 0
	sum2 := 0

	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			sum1 += arr1[i]
			i += 1
		} else if arr1[i] > arr2[j] {
			sum2 += arr2[j]
			j += 1
		} else {
			if sum1 > sum2 {
				result += sum1
			} else {
				result += sum2
			}
			result = result + arr1[i]
			sum1 = 0
			sum2 = 0
			i += 1
			j += 1
		}
	}
	for i < size1 {
		sum1 += arr1[i]
		i += 1
	}

	for j < size2 {
		sum2 += arr2[j]
		j += 1
	}

	if sum1 > sum2 {
		result += sum1
	} else {
		result += sum2
	}

	return result
}

/* Testing code */
func main13() {
	arr1 := []int{12, 13, 18, 20, 22, 26, 70}
	arr2 := []int{11, 15, 18, 19, 20, 26, 30, 31}
	fmt.Println("Max Path Sum :: ", maxPathSum(arr1, len(arr1), arr2, len(arr2)))
}

func Factorial(i int) int {
	// Termination Condition
	if i <= 1 {
		return 1
	}
	// Body, Recursive Expansion
	return i * Factorial(i-1)
}

func main14() {
	fmt.Println("factorial 5 is :: ", Factorial(5))
}

func printInt1(number int) {
	digit := (number%10 + '0')
	number = number / 10
	if number != 0 {
		printInt1(number)
	}
	fmt.Printf("%c", digit)
}

func printInt(number int) {
	conversion := "0123456789ABCDEF"
	base := 16
	digit := number % base
	number = number / base
	
	if number != 0 {
		printInt(number)
	}
	
	fmt.Print(string(conversion[digit]))
}
func main15() {
	printInt(95)
}

func towerOfHanoi(num int, src byte, dst byte, temp byte) {
	if num < 1 {
		return
	}

	towerOfHanoi(num-1, src, temp, dst)
	fmt.Printf("Move %d disk from peg %c to peg %c \n", num, src, dst)
	towerOfHanoi(num-1, temp, dst, src)
}

func main16() {
	num := 3
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	towerOfHanoi(num, 'A', 'C', 'B')
}

func GCD(m int, n int) int {
	if m < n {
		return GCD(n, m)
	}
	if m%n == 0 {
		return n
	}
	return GCD(n, m%n)
}

func main17() {
fmt.Println("GCD is :", GCD(24, 18))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main18() {
	fmt.Println("10th number in fibonacci series :", fibonacci(10))
}

func Permutation(data []int, i int, length int) {
	if length == i {
		fmt.Println(data)
		return
	}

	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func main19() {
	var data [3]int
	for i := 0; i < 3; i++ {
		data[i] = i
	}
	Permutation(data[:], 0, 3)
}



func BinarySearchRecursive(data []int, value int) bool {
	size := len(data)
	return BinarySearchRecursiveUtil(data, 0, size - 1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := (low + high)/2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursiveUtil(data, mid+1, high, value)
	} else {
		return BinarySearchRecursiveUtil(data, low, mid-1, value)
	}
}

func main20() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("BinarySearchRecursive:", BinarySearchRecursive(arr, 7))
	fmt.Println("BinarySearchRecursive:", BinarySearchRecursive(arr, 8))
}
BinarySearchRecursive: true
BinarySearchRecursive: false
func main() {
//	main1()
//	main2()
//	main3()
//	main4()
//	main5()
//	main6()
//	main7()
//	main8()
//	main9()
//	main10()
//  main11()
//  main12()
//  main13()
//  main14()
//  main15()
//  main16()
// main17()
// main18()
// main19()
 main20()
}
