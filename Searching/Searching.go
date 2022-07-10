package main

import (
	"fmt"
	"math"
	"sort"
)

func linearSearchUnsorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func linearSearchSorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		} else if value < data[i] {
			return false
		}
	}
	return false
}

// Binary Search Algorithm - Iterative Way
func BinarySearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func BinarySearchRecursive(data []int, value int) bool {
	size := len(data)
	return BinarySearchRecursiveUtil(data, 0, size-1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := (low + high) / 2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursiveUtil(data, mid+1, high, value)
	} else {
		return BinarySearchRecursiveUtil(data, low, mid-1, value)
	}
}

func FibonacciSearch(arr []int, value int) bool {
	size := len(arr)
	/* Initialize fibonacci numbers */
	fibNMn2 := 0
	fibNMn1 := 1
	fibN := fibNMn2 + fibNMn1
	for fibN < size {
		fibNMn2 = fibNMn1
		fibNMn1 = fibN
		fibN = fibNMn2 + fibNMn1
	}
	low := 0
	for fibN > 1 { // fibonacci series start with 0, 1, 1, 2
		i := min(low+fibNMn2, size-1)
		if arr[i] == value {
			return true
		} else if arr[i] < value {
			fibN = fibNMn1
			fibNMn1 = fibNMn2
			fibNMn2 = fibN - fibNMn1
			low = i
		} else { // for feb2 <= 1, these will be invalid.
			fibN = fibNMn2
			fibNMn1 = fibNMn1 - fibNMn2
			fibNMn2 = fibN - fibNMn1
		}
	}
	// above loop does not check when fibNMn2 = 0
	if arr[low+fibNMn2] == value {
		return true
	}
	return false
}

func main1() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 18, 21, 23, 24}
	fmt.Println(linearSearchUnsorted(arr, 18))
	fmt.Println(linearSearchUnsorted(arr, 11))
	fmt.Println(linearSearchSorted(arr, 18))
	fmt.Println(linearSearchSorted(arr, 11))
	fmt.Println(BinarySearch(arr, 18))
	fmt.Println(BinarySearch(arr, 11))
	fmt.Println(BinarySearchRecursive(arr, 18))
	fmt.Println(BinarySearchRecursive(arr, 11))
	fmt.Println(FibonacciSearch(arr, 18))
	fmt.Println(FibonacciSearch(arr, 11))
}

/*

true
false
true
false
true
false
true
false

*/

func FirstRepeated(data []int) int {
	size := len(data)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				return data[i]
			}
		}
	}
	return 0
}

func FirstRepeated2(arr []int) int {
	size := len(arr)
	hm := make(map[int]int)
	for i := 0; i < size; i++ {
		_, ok := hm[arr[i]]
		if ok {
			hm[arr[i]] = 2
		} else {
			hm[arr[i]] = 1
		}
	}
	for i := 0; i < size; i++ {
		if hm[arr[i]] == 2 {
			return arr[i]
		}
	}
	return 0
}

func main2() {
	first := []int{1, 3, 5, 3, 9, 1, 30}
	fmt.Println("FirstRepeated :", FirstRepeated(first))
	fmt.Println("FirstRepeated :", FirstRepeated2(first))
}

/*
FirstRepeated : 1
FirstRepeated : 1
*/

func printRepeating(data []int) {
	size := len(data)
	fmt.Print("Repeating elements: ")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Print(data[i], " ")
			}
		}
	}
	fmt.Println()
}

func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data)
	fmt.Print("Repeating elements: ")

	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func printRepeating3(data []int) {
	s := make(Set)
	size := len(data)
	fmt.Print("Repeating elements: ")
	for i := 0; i < size; i++ {
		if s.Find(data[i]) {
			fmt.Print(data[i], " ")
		} else {
			s.Add(data[i])
		}
	}
	fmt.Println()
}

func printRepeating4(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)
	fmt.Print("Repeating elements: ")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Print(data[i], " ")
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}

func main3() {
	first := []int{1, 3, 5, 7, 9, 7, 25, 21, 30}
	printRepeating(first)
	printRepeating2(first)
	printRepeating3(first)
	printRepeating4(first, 100)
}

/*
Repeating elements: 7
Repeating elements: 7
Repeating elements: 7
Repeating elements: 7
*/

func RemoveDuplicates(data []int) []int {
	size := len(data)
	sort.Ints(data)
	j := 0
	for i := 1; i < size; i++ {
		if data[i] != data[j] {
			j++
			data[j] = data[i]
		}
	}
	return data[:j+1]
}

func RemoveDuplicates2(arr []int) []int {
	size := len(arr)
	hm := make(Set)
	j := 0
	for i := 0; i < size; i++ {
		if !hm.Find(arr[i]) {
			arr[j] = arr[i]
			j++
			hm.Add(arr[i])
		}
	}
	return arr[:j]
}

func main4() {
	first := []int{1, 3, 5, 3, 9, 1, 30}
	fmt.Println(RemoveDuplicates(first))
	first = []int{1, 3, 5, 3, 9, 1, 30}
	fmt.Println(RemoveDuplicates2(first))
}

func FindMissingNumber(data []int) (int, bool) {
	var found int
	size := len(data)
	for i := 1; i <= size; i++ {
		found = 0
		for j := 0; j < size; j++ {
			if data[j] == i {
				found = 1
				break
			}
		}
		if found == 0 {
			return i, true
		}
	}
	fmt.Println("NoNumberMissing")
	return 0, false
}

func FindMissingNumber2(arr []int) (int, bool) {
	size := len(arr)
	sort.Ints(arr)
	for i := 0; i < size; i++ {
		if arr[i] != i+1 {
			return i + 1, true
		}
	}
	return size, false
}

func FindMissingNumber3(arr []int) (int, bool) {
	size := len(arr)
	hm := make(map[int]int)
	for i := 0; i < size; i++ {
		hm[arr[i]] = 1
	}
	for i := 1; i <= size; i++ {
		_, ok := hm[i]
		if !ok {
			return i, true
		}
	}
	return math.MaxInt32, false
}
func FindMissingNumber4(arr []int) (int, bool) {
	size := len(arr)
	count := make([]int, size+1)
	for i := range count {
		count[i] = -1
	}
	for i := 0; i < size; i++ {
		count[arr[i]-1] = 1
	}
	for i := 0; i <= size; i++ {
		if count[i] == -1 {
			return i + 1, true
		}
	}
	return math.MaxInt32, false
}
func FindMissingNumber5(arr []int) (int, bool) {
	size := len(arr)
	sum := 0
	for i := 1; i < size+2; i++ {
		sum += i
	}
	for i := 0; i < size; i++ {
		sum -= arr[i]
	}

	if sum == 0 {
		return -1, false
	} else {
		return sum, true
	}
}

func FindMissingNumber6(arr []int) (int, bool) {
	size := len(arr)
	for i := 0; i < size; i++ {
		// len(arr)+1 value should be ignored.
		if arr[i] != size+1 && arr[i] != size*3+1 {
			// 1 should not become (len(arr)+1) so multiplied by 2
			arr[(arr[i]-1)%(size)] += (size * 2)
		}
	}
	for i := 0; i < size; i++ {
		if arr[i] <= (size * 2) {
			return i + 1, true
		}
	}
	return math.MaxInt32, false
}

func FindMissingNumber7(arr []int, dataRange int) (int, bool) {
	size := len(arr)
	xorSum := 0
	// get the XOR of all the numbers from 1 to dataRange
	for i := 1; i <= dataRange; i++ {
		xorSum ^= i
	}
	// loop through the array and get the XOR of elements
	for i := 0; i < size; i++ {
		xorSum ^= arr[i]
	}
	if xorSum == 0 {
		return -1, false
	} else {
		return xorSum, true
	}
}

func FindMissingNumber8(arr []int, upperRange int) (int, bool) {
	size := len(arr)
	st := make(Set)
	i := 0
	for i < size {
		st.Add(arr[i])
		i += 1
	}
	i = 1
	for i <= upperRange {
		if st.Find(i) == false {
			return i, true
		}
		i += 1
	}
	fmt.Println("NoNumberMissing")
	return -1, false
}

func main5() {
	first := []int{1, 3, 5, 4, 6, 8, 2}
	fmt.Println(FindMissingNumber(first))
	fmt.Println(FindMissingNumber2(first))
	fmt.Println(FindMissingNumber3(first))
	fmt.Println(FindMissingNumber4(first))
	fmt.Println(FindMissingNumber5(first))
	first = []int{1, 3, 5, 4, 6, 8, 2}
	fmt.Println(FindMissingNumber6(first))
	first = []int{1, 3, 5, 4, 6, 8, 2}
	fmt.Println(FindMissingNumber7(first, 8))
	fmt.Println(FindMissingNumber8(first, 8))
}

/*
2 true
2 true
2 true
2 true
2 true
2 true
2 true
2 true
*/

func MissingValues(arr []int) {
	size := len(arr)
	max := arr[0]
	min := arr[0]
	for i := 1; i < size; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	var found bool
	for i := min + 1; i < max; i++ {
		found = false
		for j := 0; j < size; j++ {
			if arr[j] == i {
				found = true
				break
			}
		}
		if !found {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func MissingValues2(arr []int) {
	size := len(arr)
	sort.Ints(arr)
	value := arr[0]
	i := 0
	for i < size {
		if value == arr[i] {
			value += 1
			i += 1
		} else {
			fmt.Print(value, " ")
			value += 1
		}
	}
	fmt.Println()
}

func MissingValues3(arr []int) {
	size := len(arr)
	ht := make(Set)
	minVal := math.MaxInt32
	maxVal := math.MinInt32
	i := 0
	for i = 0; i < size; i++ {
		ht.Add(arr[i])
		if minVal > arr[i] {
			minVal = arr[i]
		}
		if maxVal < arr[i] {
			maxVal = arr[i]
		}
	}
	for i = minVal; i < maxVal+1; i++ {
		if ht.Find(i) == false {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func main6() {
	arr := []int{11, 14, 13, 17, 21, 18, 19, 23, 24}
	MissingValues(arr)
	MissingValues2(arr)
	MissingValues3(arr)
}

/*
12 15 16 20 22 
12 15 16 20 22 
12 15 16 20 22 
*/

func OddCount(arr []int) {
	size := len(arr)
	xorSum := 0
	for i := 0; i < size; i++ {
		xorSum ^= arr[i]
	}
	fmt.Println("Odd values: ", xorSum)
}

func OddCount2(arr []int) {
	size := len(arr)
	ctr := make(map[int]int)
	count := 0

	for i := 0; i < size; i++ {
		val, ok := ctr[arr[i]]
		if ok {
			ctr[arr[i]] = val + 1
		} else {
			ctr[arr[i]] = 1
		}
	}
	for i := 0; i < size; i++ {
		val, ok := ctr[arr[i]]
		if ok && (val%2 == 1) {
			count++
			delete(ctr, arr[i])
		}
	}
	fmt.Println("Odd count is :: ", count)
}

func OddCountElements(arr []int) {
	size := len(arr)
	xorSum := 0
	first := 0
	second := 0
	var setBit int
	/*
	* xor of all elements in arr[] even occurrence will cancel each other. sum will
	* contain sum of two odd elements.
	 */
	for i := 0; i < size; i++ {
		xorSum = xorSum ^ arr[i]
	}

	/* Rightmost set bit. */
	setBit = xorSum &^ (xorSum - 1)

	/*
	* Dividing elements in two group: Elements having setBit bit as 1. Elements
	* having setBit bit as 0. Even elements cancelled themselves if group and we
	* get our numbers.
	 */
	for i := 0; i < size; i++ {
		if (arr[i] & setBit) != 0 {
			first ^= arr[i]
		} else {
			second ^= arr[i]
		}
	}
	fmt.Println("Odd count Elements are :: ", first, "&", second)
}

func main7() {
	arr := []int{10, 25, 30, 10, 15, 25, 15}
	OddCount(arr)
	OddCount2(arr)
	arr = []int{10, 25, 30, 10, 15, 25, 15, 40}
	OddCountElements(arr)
}

/*
Odd values:  30
Odd count is ::  1
Odd count Elements are ::  30 & 40
*/

func SumDistinct(arr []int) {
	sum := 0
	size := len(arr)
	sort.Ints(arr)
	for i := 0; i < (size - 1); i++ {
		if arr[i] != arr[i+1] {
			sum += arr[i]
		}
	}
	sum += arr[size-1]
	fmt.Println(sum)
}

func main8() {
	arr := []int{1, 9, 2, 4, 3, 5, 4, 5}
	SumDistinct(arr)
}

/*
24
*/

func minAbsSumPair(data []int) {
	var sum int
	size := len(data)
	// Array should have at least two elements
	if size < 2 {
		fmt.Println("InvalidInput")
	}
	// Initialization of values
	minFirst := 0
	minSecond := 1
	minSum := abs(data[0] + data[1])
	for l := 0; l < size-1; l++ {
		for r := l + 1; r < size; r++ {
			sum = abs(data[l] + data[r])
			if sum < minSum {
				minSum = sum
				minFirst = l
				minSecond = r
			}
		}
	}
	fmt.Println("The two elements with minimum sum are : ", data[minFirst], " & ", data[minSecond])
}

func minAbsSumPair2(data []int) {
	var sum int
	size := len(data)
	// Array should have at least two elements
	if size < 2 {
		fmt.Println("InvalidInput")
	}
	sort.Ints(data) // Sort(data, size)

	// Initialization of values
	minFirst := 0
	minSecond := size - 1
	minSum := abs(data[minFirst] + data[minSecond])
	for l, r := 0, (size - 1); l < r; {
		sum = (data[l] + data[r])
		if abs(sum) < minSum {
			minSum = abs(sum)
			minFirst = l
			minSecond = r
		}

		if sum < 0 {
			l++
		} else if sum > 0 {
			r--
		} else {
			break
		}
	}
	fmt.Println("The two elements with minimum sum are : ", data[minFirst], " & ", data[minSecond])
}

func main9() {
	first := []int{1, 5, -10, 3, 2, -6, 8, 9, 6}
	minAbsSumPair2(first)
	minAbsSumPair(first)
}

/*
The two elements with minimum sum are :  -6  &  6
The two elements with minimum sum are :  -6  &  6
*/

func FindPair(data []int, value int) bool {
	size := len(data)
	ret := false
	fmt.Print("Pairs with sum ", value, " are : ")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (data[i] + data[j]) == value {
				fmt.Print("(", data[i], ", ", data[j], ") ")
				ret = true
			}
		}
	}
	return ret
}

func FindPair2(data []int, value int) bool {
	size := len(data)
	first := 0
	second := size - 1
	ret := false
	sort.Ints(data) // Sort(data, size)
	fmt.Print("Pairs with sum ", value, " are : ")
	for first < second {
		curr := data[first] + data[second]
		if curr == value {
			fmt.Print("(", data[first], ", ", data[second], ") ")
			ret = true
		}
		if curr < value {
			first++
		} else {
			second--
		}
	}
	return ret
}

func FindPair3(data []int, value int) bool {
	s := make(Set)
	size := len(data)
	ret := false
	fmt.Print("Pairs with sum ", value, " are : ")
	for i := 0; i < size; i++ {
		if s.Find(value - data[i]) {
			fmt.Print("(", data[i], ", ", (value - data[i]), ") ")
			ret = true
		} else {
			s.Add(data[i])
		}
	}
	return ret
}

func FindPair4(data []int, rangeVal int, value int) bool {
	size := len(data)
	count := make([]int, rangeVal+1)
	for i := 0; i < size; i++ {
		if count[value-data[i]] > 0 {
			fmt.Println("The pair is : ", data[i], ", ", value-data[i])
			return true
		}
		count[data[i]] += 1
	}
	return false
}

func main10() {
	first := []int{1, 5, 4, 3, 2, 7, 8, 9, 6}
	FindPair(first, 8)
	FindPair2(first, 8)
	FindPair3(first, 8)
	FindPair4(first, 9, 8)
}

/*
Pairs with sum 8 are : (1, 7) (5, 3) (2, 6)
Pairs with sum 8 are : (1, 7) (2, 6) (3, 5)
Pairs with sum 8 are : (5, 3) (6, 2) (7, 1)
The pair is :  5 ,  3
*/

func FindPairTwoLists(arr1 []int, size1 int, arr2 []int, size2 int, value int) bool {
	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			if arr1[i]+arr2[j] == value {
				fmt.Println("The pair is :", arr1[i], ",", arr2[j])
				return true
			}
		}
	}
	return false
}
func FindPairTwoLists2(arr1 []int, size1 int, arr2 []int, size2 int, value int) bool {
	sort.Ints(arr2)
	for i := 0; i < size1; i++ {
		if BinarySearch(arr2, value-arr1[i]) {
			fmt.Println("The pair is :", arr1[i], ",", value-arr1[i])
		}
		return true
	}
	return false
}
func FindPairTwoLists3(arr1 []int, size1 int, arr2 []int, size2 int, value int) bool {
	first := 0
	second := size2 - 1
	curr := 0
	sort.Ints(arr1)
	sort.Ints(arr2)
	for first < size1 && second >= 0 {
		curr = arr1[first] + arr2[second]
		if curr == value {
			fmt.Println("The pair is :", arr1[first], ",", arr2[second])
			return true
		} else if curr < value {
			first++
		} else {
			second--
		}
	}
	return false
}
func FindPairTwoLists4(arr1 []int, size1 int, arr2 []int, size2 int, value int) bool {
	hs := make(Set)
	for i := 0; i < size2; i++ {
		hs.Add(arr2[i])
	}
	for i := 0; i < size1; i++ {
		if hs.Find(value - arr1[i]) {
			fmt.Println("The pair is :", arr1[i], ",", value-arr1[i])
			return true
		}
	}
	return false
}
func FindPairTwoLists5(arr1 []int, size1 int, arr2 []int, size2 int, rangeVal int, value int) bool {
	count := make([]int, rangeVal+1)
	for i := 0; i < size2; i++ {
		count[arr2[i]] = 1
	}
	for i := 0; i < size1; i++ {
		if count[value-arr1[i]] != 0 {
			fmt.Println("The pair is :", arr1[i], ",", value-arr1[i])
			return true
		}
	}
	return false
}

func Main10A() {
	first := []int{1, 5, 4, 3, 2, 7, 8, 9, 6}
	second := []int{1, 5, 4, 3, 2, 7, 8, 9, 6}
	fmt.Println(FindPairTwoLists(first, len(first), second, len(second), 8))
	fmt.Println(FindPairTwoLists2(first, len(first), second, len(second), 8))
	fmt.Println(FindPairTwoLists3(first, len(first), second, len(second), 8))
	fmt.Println(FindPairTwoLists4(first, len(first), second, len(second), 8))
	fmt.Println(FindPairTwoLists5(first, len(first), second, len(second), 9, 8))
}

/*
The pair is : 1 , 7
true
The pair is : 1 , 7
true
The pair is : 1 , 7
true
The pair is : 1 , 7
true
The pair is : 1 , 7
true
*/

func FindDifference(arr []int, value int) bool {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if abs(arr[i]-arr[j]) == value {
				fmt.Println("The pair is::", arr[i], "&", arr[j])
				return true
			}
		}
	}
	return false
}

func FindDifference2(arr []int, value int) bool {
	first := 0
	second := 0
	size := len(arr)
	var diff int
	sort.Ints(arr)
	for first < size && second < size {
		diff = abs(arr[first] - arr[second])
		if diff == value {
			fmt.Println("The pair is::", arr[first], "&", arr[second])
			return true
		} else if diff > value {
			first += 1
		} else {
			second += 1
		}
	}
	return false
}

func main11() {
	arr := []int{1, 5, 4, 3, 2, 7, 8, 9, 6}
	fmt.Println(FindDifference(arr, 6))
	fmt.Println(FindDifference2(arr, 6))
}

/*
The pair is:: 1 & 7
true
The pair is:: 1 & 7
true
*/

func FindMinDiff(arr []int) int {
	size := len(arr)
	diff := math.MaxInt32
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			value := abs(arr[i] - arr[j])
			if diff > value {
				diff = value
			}
		}
	}
	return diff
}

func FindMinDiff2(arr []int) int {
	sort.Ints(arr)
	diff := math.MaxInt32
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		if (arr[i+1] - arr[i]) < diff {
			diff = arr[i+1] - arr[i]
		}
	}
	return diff
}

func MinDiffPair(arr1 []int, arr2 []int) int {
	diff := math.MaxInt32
	first, second := 0, 0
	size1 := len(arr1)
	size2 := len(arr2)
	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			value := abs(arr1[i] - arr2[j])
			if diff > value {
				diff = value
				first = arr1[i]
				second = arr2[j]
			}
		}
	}
	fmt.Println("The pair is::", first, "&", second)
	fmt.Println("Minimum difference is :: ", diff)
	return diff
}

func MinDiffPair2(arr1 []int, arr2 []int) int {
	minDiff := math.MaxInt32
	first, second, diff := 0, 0, 0
	out1, out2 := 0, 0
	size1 := len(arr1)
	size2 := len(arr2)
	sort.Ints(arr1)
	sort.Ints(arr2)
	for first < size1 && second < size2 {
		diff = abs(arr1[first] - arr2[second])
		if minDiff > diff {
			minDiff = diff
			out1 = arr1[first]
			out2 = arr2[second]
		}
		if arr1[first] < arr2[second] {
			first += 1
		} else {
			second += 1
		}
	}
	fmt.Println("The pair is ::", out1, "&", out2)
	fmt.Println("Minimum difference is :: ", minDiff)
	return minDiff
}

func main12() {
	first := []int{1, 3, 2, 7, 8, 9}
	second := []int{5, 10, 15}
	fmt.Println("MinDiff", FindMinDiff(first))
	fmt.Println("MinDiff", FindMinDiff2(first))
	MinDiffPair(first, second)
	MinDiffPair2(first, second)
}

/*
MinDiff 1
MinDiff 1
The pair is:: 9 & 10
Minimum difference is ::  1
The pair is :: 9 & 10
Minimum difference is ::  1
*/

func ClosestPair(arr []int, value int) {
	diff := math.MaxInt32
	first := -1
	second := -1
	curr := 0
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			curr = abs(value - (arr[i] + arr[j]))
			if curr < diff {
				diff = curr
				first = arr[i]
				second = arr[j]
			}
		}
	}
	fmt.Println("closest pair is ::", first, "&", second)
}

func ClosestPair2(arr []int, value int) {
	first := 0
	second := 0
	start := 0
	size := len(arr)

	stop := size - 1
	diff := 0
	curr := 0
	sort.Ints(arr)
	diff = math.MaxInt32

	for start < stop {
		curr = (value - (arr[start] + arr[stop]))
		if abs(curr) < diff {
			diff = abs(curr)
			first = arr[start]
			second = arr[stop]
		}
		if curr == 0 {
			break
		} else if curr > 0 {
			start += 1
		} else {
			stop -= 1
		}
	}

	fmt.Println("closest pair is ::", first, "&", second)
}

func main13() {
	first := []int{1, 5, 4, 3, 2, 7, 8, 9, 6}
	ClosestPair(first, 6)
	ClosestPair2(first, 6)
}

/*
closest pair is :: 1 & 5
closest pair is :: 1 & 5
*/

func SumPairRestArray(arr []int) bool {
	var total, low, high, curr, value int
	size := len(arr)
	sort.Ints(arr)
	total = 0
	for i := 0; i < size; i++ {
		total += arr[i]
	}
	value = total / 2
	low = 0
	high = size - 1
	for low < high {
		curr = arr[low] + arr[high]
		if curr == value {
			fmt.Println("Pair is ::", arr[low], "&", arr[high])
			return true
		} else if curr < value {
			low += 1
		} else {
			high -= 1
		}
	}
	return false
}

func main14() {
	first := []int{1, 2, 4, 3, 7, 3}
	SumPairRestArray(first)
}

/*
Pair is :: 3 & 7
*/

func ZeroSumTriplets(arr []int) {
	size := len(arr)
	fmt.Print("Zero Sum Triplets are :: ")
	for i := 0; i < (size - 2); i++ {
		for j := i + 1; j < (size - 1); j++ {
			for k := j + 1; k < size; k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					fmt.Print("(", arr[i], arr[j], arr[k], ") ")
				}
			}
		}
	}
}

func ZeroSumTriplets2(arr []int) {
	size := len(arr)
	var start, stop int
	sort.Ints(arr)
	fmt.Print("Zero Sum Triplets are :: ")
	for i := 0; i < (size - 2); i++ {
		start = i + 1
		stop = size - 1

		for start < stop {
			if arr[i]+arr[start]+arr[stop] == 0 {
				fmt.Print("(", arr[i], arr[start], arr[stop], ")")
				start += 1
				stop -= 1
			} else if arr[i]+arr[start]+arr[stop] > 0 {
				stop -= 1
			} else {
				start += 1
			}
		}
	}
}

func main15() {
	first := []int{1, 2, -4, 3, 7, -3}
	ZeroSumTriplets(first)
	ZeroSumTriplets2(first)
}

/*
Zero Sum Triplets are :: (1 2 -3) (1 -4 3) (-4 7 -3)
Zero Sum Triplets are :: (-4 -3 7)(-4 1 3)(-3 1 2)
*/

func findTriplet(arr []int, value int) {
	size := len(arr)
	fmt.Print("Triplet with sum ", value, " are :: ")
	for i := 0; i < (size - 2); i++ {
		for j := i + 1; j < (size - 1); j++ {
			for k := j + 1; k < size; k++ {
				if (arr[i] + arr[j] + arr[k]) == value {
					fmt.Print("(", arr[i], arr[j], arr[k], ")")
				}
			}
		}
	}
}

func findTriplet2(arr []int, value int) {
	size := len(arr)
	var start, stop int
	sort.Ints(arr)
	fmt.Print("Triplet with sum ", value, " are :: ")
	for i := 0; i < size-2; i++ {
		start = i + 1
		stop = size - 1
		for start < stop {
			if arr[i]+arr[start]+arr[stop] == value {
				fmt.Print("(", arr[i], arr[start], arr[stop], ")")
				start += 1
				stop -= 1
			} else if arr[i]+arr[start]+arr[stop] > value {
				stop -= 1
			} else {
				start += 1
			}
		}
	}
}

func main16() {
	first := []int{1, 2, -4, 3, 7, -3}
	findTriplet(first, 6)
	findTriplet2(first, 6)
}

/*
Triplet with sum 6 are :: (1 2 3)(2 7 -3)(-4 3 7)
Triplet with sum 6 are :: (-4 3 7)(-3 2 7)(1 2 3)
*/

func AbcTriplet(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			for k := 0; k < size; k++ {
				if k != i && k != j && arr[i]+arr[j] == arr[k] {
					fmt.Println("AbcTriplet ::", arr[i], arr[j], arr[k])
					return
				}
			}
		}
	}
}

func AbcTriplet2(arr []int) {
	var start, stop int
	size := len(arr)
	sort.Ints(arr)
	for i := 0; i < (size - 1); i++ {
		start = 0
		stop = size - 1
		for start < stop {
			if arr[i] == arr[start]+arr[stop] {
				fmt.Println("AbcTriplet ::", arr[i], arr[start], arr[stop])
				return
			} else if arr[i] < arr[start]+arr[stop] {
				stop -= 1
			} else {
				start += 1
			}
		}
	}
}

func main17() {
	first := []int{1, 2, -4, 3, 8, -3}
	AbcTriplet(first)
	AbcTriplet2(first)
}

/*
AbcTriplet::  1   2   3
AbcTriplet::  1   -4   -3
AbcTriplet::  -3   -4   1
*/

func SmallerThenTripletCount(arr []int, value int) {
	count := 0
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if arr[i]+arr[j]+arr[k] < value {
					count += 1
				}
			}
		}
	}
	fmt.Println("SmallerThenTripletCount::", count)
}

func SmallerThenTripletCount2(arr []int, value int) {
	var start, stop int
	count := 0
	size := len(arr)
	sort.Ints(arr)

	for i := 0; i < (size - 2); i++ {
		start = i + 1
		stop = size - 1
		for start < stop {
			if arr[i]+arr[start]+arr[stop] >= value {
				stop -= 1
			} else {
				count += stop - start
				start += 1
			}
		}
	}
	fmt.Println("SmallerThenTripletCount::", count)
}

func main18() {
	first := []int{1, 2, -4, 3, 7, -3}
	SmallerThenTripletCount(first, 6)
	SmallerThenTripletCount2(first, 6)
}

/*
SmallerThenTripletCount:: 13
SmallerThenTripletCount:: 13
*/

func APTriplets(arr []int) {
	size := len(arr)
	var i, j, k int
	for i = 1; i < size-1; i++ {
		j = i - 1
		k = i + 1
		for j >= 0 && k < size {
			if arr[j]+arr[k] == 2*arr[i] {
				fmt.Println("Triplet ::", arr[j], arr[i], arr[k])
				k += 1
				j -= 1
			} else if arr[j]+arr[k] < 2*arr[i] {
				k += 1
			} else {
				j -= 1
			}
		}
	}
}

func GPTriplets(arr []int) {
	size := len(arr)
	var i, j, k int
	for i = 1; i < size-1; i++ {
		j = i - 1
		k = i + 1
		for j >= 0 && k < size {
			if arr[j]*arr[k] == arr[i]*arr[i] {
				fmt.Println("Triplet is :: ", arr[j], arr[i], arr[k])
				k += 1
				j -= 1
			} else if arr[j]+arr[k] < 2*arr[i] {
				k += 1
			} else {
				j -= 1
			}
		}
	}
}

func main19() {
	first := []int{1, 2, 3, 4, 9, 17, 23}
	APTriplets(first)
	GPTriplets(first)
}

/*
Triplet :: 1 2 3
Triplet :: 2 3 4
Triplet :: 1 9 17

Triplet is ::  1 3 9
*/

func NumberOfTriangles(arr []int, size int) int {
	var i, j, k int
	count := 0
	for i = 0; i < (size - 2); i++ {
		for j = i + 1; j < (size - 1); j++ {
			for k = j + 1; k < size; k++ {
				if (arr[i]+arr[j] > arr[k]) && (arr[k]+arr[j] > arr[i]) && (arr[i]+arr[k] > arr[j]) {
					count += 1
				}
			}
		}
	}
	return count
}

func NumberOfTriangles2(arr []int, size int) int {
	var i, j, k int
	count := 0
	sort.Ints(arr)

	for i = 0; i < (size - 2); i++ {
		k = i + 2
		for j = i + 1; j < (size - 1); j++ {
			/*
			* if sum of arr[i] & arr[j] is greater arr[k] then sum of arr[i] & arr[j + 1]
			* is also greater than arr[k] this improvement make algo O(n2)
			 */
			for k < size && arr[i]+arr[j] > arr[k] {
				k += 1
			}

			count += k - j - 1
		}
	}
	return count
}

func main20() {
	first := []int{1, 2, 3, 4, 5}
	fmt.Println(NumberOfTriangles(first, len(first)))
	fmt.Println(NumberOfTriangles2(first, len(first)))
}

/*
3
3
*/

func getMax(data []int) int {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1
	for i := 0; i < size; i++ {
		count = 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	return max
}

func getMax2(data []int) int {
	size := len(data)
	max := data[0]
	maxCount := 1
	curr := data[0]
	currCount := 1
	sort.Ints(data) // Sort(data,size)
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}

func getMax3(data []int, dataRange int) int {
	max := data[0]
	maxCount := 1
	size := len(data)
	count := make([]int, dataRange)
	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}

func main21() {
	first := []int{1, 30, 5, 13, 9, 31, 5}
	fmt.Println(getMax(first))
	fmt.Println(getMax2(first))
	fmt.Println(getMax3(first, 50))
}

/*
5
5
5
*/

func getMajority(data []int) (int, bool) {
	size := len(data)
	max := 0
	count := 0
	maxCount := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	if maxCount > size/2 {
		return max, true
	}
	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

func getMajority2(data []int) (int, bool) {
	size := len(data)
	majIndex := size / 2
	sort.Ints(data) // Sort(data,size)
	candidate := data[majIndex]
	count := 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

func getMajority3(data []int) (int, bool) {
	majIndex := 0
	count := 1
	size := len(data)

	for i := 1; i < size; i++ {
		if data[majIndex] == data[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	candidate := data[majIndex]
	count = 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

func IsMajority(arr []int) bool {
	size := len(arr)
	count := 0
	mid := arr[size/2]
	for i := 0; i < size; i++ {
		if arr[i] == mid {
			count += 1
		}
	}
	if count > size/2 {
		return true
	}
	return false
}

func IsMajority2(arr []int) bool {
	size := len(arr)
	majority := arr[size/2]
	i := findFirstIndex(arr, 0, size-1, majority)
	/*
	* we are using majority element form array so we will get some valid index
	* always.
	 */
	if ((i + size/2) <= (size - 1)) && arr[i+size/2] == majority {
		return true
	} else {
		return false
	}
}

func main22() {
	first := []int{1, 5, 5, 13, 5, 31, 5}
	fmt.Println(getMajority(first))
	fmt.Println(getMajority2(first))
	fmt.Println(getMajority3(first))
	fmt.Println(IsMajority(first))
	fmt.Println(IsMajority2(first))

}

/*
5 true
5 true
5 true
true
true
*/

func FindBitonicArrayMax(arr []int) int {
	size := len(arr)
	for i := 0; i < size-2; i++ {
		if arr[i] > arr[i+1] {
			return arr[i]
		}
	}
	fmt.Println("error not a bitonic array")
	return 0
}

func FindBitonicArrayMax2(data []int) (int, bool) {
	size := len(data)
	start := 0
	end := size - 1

	if size < 3 {
		fmt.Println("InvalidInput")
		return 0, false
	}

	for start <= end {
		mid := (start + end) / 2

		if data[mid-1] < data[mid] && data[mid+1] < data[mid] { //maxima
			return data[mid], true
		} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] { // increasing
			start = mid + 1
		} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] { // decreasing
			end = mid - 1
		} else {
			break
		}
	}
	fmt.Println("NoMaximaFound")
	return 0, false
}

func SearchBitonicArray(data []int, key int) bool {
	size := len(data)
	maxIndex, _ := FindBitonicArrayMaxIndex(data)
	k := BinarySearch2(data, 0, maxIndex, key, true)
	if k != -1 {
		return true
	}
	k = BinarySearch2(data, maxIndex+1, size-1, key, false)
	if k != -1 {
		return true
	} else {
		return false
	}
}

func FindBitonicArrayMaxIndex(data []int) (int, bool) {
	size := len(data)
	start := 0
	end := size - 1
	mid := 0
	if size < 3 {
		fmt.Println("InvalidInput")
		return -1, false
	}
	for start <= end {
		mid = (start + end) / 2
		if data[mid-1] < data[mid] && data[mid+1] < data[mid] /* maxima */ {
			return mid, true
		} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] /* increasing */ {
			start = mid
		} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] /* decreasing */ {
			end = mid
		} else {
			break
		}
	}
	fmt.Println("error")
	return -1, false
}

func main23() {
	first := []int{1, 5, 10, 13, 20, 30, 8, 6, 5}
	fmt.Println(FindBitonicArrayMax(first))
	fmt.Println(FindBitonicArrayMaxIndex(first))
	fmt.Println(SearchBitonicArray(first, 7))
	fmt.Println(SearchBitonicArray(first, 8))
}

/*
30
5 true
false
true
*/

func BinarySearch2(data []int, start int, end int, key int, isInc bool) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] {
		return mid
	}
	if isInc != false && key < data[mid] || isInc == false && key > data[mid] {
		return BinarySearch2(data, start, mid-1, key, isInc)
	}
	return BinarySearch2(data, mid+1, end, key, isInc)
}

func findKeyCount(data []int, key int) int {
	count := 0
	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == key {
			count++
		}
	}
	return count
}

func findKeyCount2(data []int, key int) int {
	size := len(data)
	firstIndex := findFirstIndex(data, 0, size-1, key)
	lastIndex := findLastIndex(data, 0, size-1, key)
	return (lastIndex - firstIndex + 1)
}

func findFirstIndex(data []int, start int, end int, key int) int {
	if end < start {
		return -1
	}

	mid := (start + end) / 2
	if key == data[mid] && (mid == start || data[mid-1] != key) {
		return mid
	}

	if key <= data[mid] {
		return findFirstIndex(data, start, mid-1, key)
	}
	return findFirstIndex(data, mid+1, end, key)
}

func findLastIndex(data []int, start int, end int, key int) int {
	if end < start {
		return -1
	}

	mid := (start + end) / 2
	if key == data[mid] && (mid == end || data[mid+1] != key) {
		return mid
	}

	if key < data[mid] {
		return findLastIndex(data, start, mid-1, key)
	}
	return findLastIndex(data, mid+1, end, key)
}

func main24() {
	first := []int{1, 5, 6, 6, 6, 6, 7, 10, 13, 20, 30}
	fmt.Println(findKeyCount(first, 6))
	fmt.Println(findKeyCount2(first, 6))
}

/*
4
4
*/

func MaxProfit(stocks []int) int {
	size := len(stocks)
	maxProfit := 0
	buy, sell := 0, 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if maxProfit < stocks[j]-stocks[i] {
				maxProfit = stocks[j] - stocks[i]
				buy = i
				sell = j
			}
		}
	}
	fmt.Println("Purchase day is ", buy, " at price ", stocks[buy])
	fmt.Println("Sell day is ", sell, " at price ", stocks[sell])
	return maxProfit
}

func MaxProfit2(stocks []int) int {
	size := len(stocks)
	buy, sell := 0, 0
	curMin := 0
	currProfit := 0
	maxProfit := 0

	for i := 0; i < size; i++ {
		if stocks[i] < stocks[curMin] {
			curMin = i
		}

		currProfit = stocks[i] - stocks[curMin]
		if currProfit > maxProfit {
			buy = curMin
			sell = i
			maxProfit = currProfit
		}
	}
	fmt.Println("Purchase day is ", buy, " at price ", stocks[buy])
	fmt.Println("Sell day is ", sell, " at price ", stocks[sell])
	fmt.Println("Max Profit :: ", maxProfit)
	return maxProfit
}

func main25() {
	first := []int{10, 150, 6, 67, 61, 16, 86, 6, 67, 78, 150, 3, 28, 143}
	MaxProfit(first)
	MaxProfit2(first)
}

/*
Purchase day is  2  at price  6
Sell day is  10  at price  150
Purchase day is  2  at price  6
Sell day is  10  at price  150
Max Profit ::  144
*/

func findMedian(dataFirst []int, dataSecond []int) int {
	sizeFirst := len(dataFirst)
	sizeSecond := len(dataSecond)
	medianIndex := (sizeFirst + sizeSecond) / 2
	i := 0
	j := 0
	count := 0
	for count < medianIndex {
		if i < sizeFirst-1 && dataFirst[i] < dataSecond[j] {
			i++
		} else {
			j++
		}
		count++
	}
	if dataFirst[i] < dataSecond[j] {
		return dataFirst[i]
	}
	return dataSecond[j]
}

func main26() {
	first := []int{1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30}
	second := []int{1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30}
	fmt.Println(findMedian(first, second))
}

/*
6
*/

func Search01(arr []int) int {
	size := len(arr)
	for i := 0; i < size; i++ {
		if arr[i] == 1 {
			return i
		}
	}
	return -1
}

func BinarySearch01(data []int) int {
	size := len(data)
	if size == 1 && data[0] == 1 {
		return -1
	}
	return BinarySearch01Util(data, 0, size-1)
}

func BinarySearch01Util(data []int, start int, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if 1 == data[mid] && 0 == data[mid-1] {
		return mid
	}
	if 0 == data[mid] {
		return BinarySearch01Util(data, mid+1, end)
	}
	return BinarySearch01Util(data, start, mid-1)
}

func main27() {
	first := []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	fmt.Println(Search01(first))
	fmt.Println(BinarySearch01(first))
}

/*
8
8
*/

func FindRotationMaxUtil(arr []int, start int, end int) int {
	if end <= start {
		return start
	}
	mid := (start + end) / 2
	if arr[mid] > arr[mid+1] {
		return mid
	}

	if arr[start] <= arr[mid] { // increasing part.
		return FindRotationMaxUtil(arr, mid+1, end)
	} else {
		return FindRotationMaxUtil(arr, start, mid-1)
	}
}

func RotationMax(arr []int) int {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		if arr[i] > arr[i+1] {
			return arr[i]
		}
	}
	return -1
}

func RotationMax2(arr []int) int {
	size := len(arr)
	index := FindRotationMaxUtil(arr, 0, size-1)
	return arr[index]
}

func FindRotationMax(arr []int) int {
	size := len(arr)
	return FindRotationMaxUtil(arr, 0, size-1)
}

func CountRotation(arr []int) int {
	size := len(arr)
	maxIndex := FindRotationMaxUtil(arr, 0, size-1)
	return (maxIndex + 1) % size
}

func SearchRotateArray(arr []int, key int) int {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		if arr[i] == key {
			return i
		}
	}
	return -1
}

func BinarySearchRotateArrayUtil(data []int, start int, end int, key int) bool {
	if end < start {
		return false
	}
	mid := (start + end) / 2
	if key == data[mid] {
		return true
	}
	if data[mid] > data[start] {
		if data[start] <= key && key < data[mid] {
			return BinarySearchRotateArrayUtil(data, start, mid-1, key)
		}
		return BinarySearchRotateArrayUtil(data, mid+1, end, key)
	}
	if data[mid] < key && key <= data[end] {
		return BinarySearchRotateArrayUtil(data, mid+1, end, key)
	}
	return BinarySearchRotateArrayUtil(data, start, mid-1, key)
}

func BinarySearchRotateArray(data []int, key int) bool {
	size := len(data)
	return BinarySearchRotateArrayUtil(data, 0, size-1, key)
}

func main28() {
	first := []int{8, 9, 10, 11, 3, 5, 7}
	fmt.Println(SearchRotateArray(first, 5))
	fmt.Println(BinarySearchRotateArray(first, 7))
	fmt.Println(BinarySearchRotateArray(first, 6))
	fmt.Println("RotationMax is ::", RotationMax(first))
	fmt.Println("RotationMax is ::", RotationMax2(first))
	fmt.Println("FindRotationMax is ::", FindRotationMax(first))
	fmt.Println("CountRotation is ::", CountRotation(first))
}

/*
5
true
false
RotationMax is :: 11
RotationMax is :: 11
FindRotationMax is :: 3
CountRotation is :: 4
*/

func minAbsDiffAdjCircular(arr []int, size int) int {
	diff := math.MaxInt32
	if size < 2 {
		return -1
	}

	for i := 0; i < size; i++ {
		diff = min(diff, abs(arr[i]-arr[(i+1)%size]))
	}
	return diff
}

func main29() {
	arr := []int{5, 29, 22, 51, 11}
	fmt.Println(minAbsDiffAdjCircular(arr, len(arr)))
}

/*
6
*/

func seperateEvenAndOdd(data []int) {
	size := len(data)
	left := 0
	right := size - 1
	for left < right {
		if data[left]%2 == 0 {
			left++
		} else if data[right]%2 == 1 {
			right--
		} else {
			data[left], data[right] = data[right], data[left] // swap
			left++
			right--
		}
	}
}

func main30() {
	first := []int{1, 0, 5, 7, 9, 11, 12, 8, 5, 3, 1}
	seperateEvenAndOdd(first)
	fmt.Println(first)
}

/*
[8 0 12 7 9 11 5 1 5 3 1]
*/

func GetMedian(data []int) int {
	size := len(data)
	sort.Ints(data)
	return data[size/2]
}

func GetMedian2(arr []int) int {
	size := len(arr)
	QuickSelectUtil(arr, 0, size-1, size/2)
	return arr[size/2]
}

func main30A() {
	first := []int{1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30}
	fmt.Println(GetMedian(first))
	fmt.Println(GetMedian2(first))
}

/*
6
6
*/

func transformArrayAB(str string) string {
	data := []rune(str)
	size := len(data)
	N := size / 2
	for i := 1; i < N; i++ {
		for j := 0; j < i; j++ {
			data[N-i+2*j], data[N-i+2*j+1] = data[N-i+2*j+1], data[N-i+2*j]
		}
	}
	return string(data)
}

func main31() {
	str := "aaaabbbb"
	str = transformArrayAB(str)
	fmt.Println(str)
}

/*
abababab
*/

func CheckPermutation(data1 string, data2 string) bool {
	size1 := len(data1)
	size2 := len(data2)

	if size1 != size2 {
		return false
	}

	arr1 := []rune(data1)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	arr2 := []rune(data2)
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	for i := 0; i < size1; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func Search(data []int, value int) bool {
	for _, key := range data {
		if key == value {
			return true
		}
	}
	return false
}

func CheckPermutation2(data1 string, data2 string) bool {
	size1 := len(data1)
	size2 := len(data2)
	h := make(map[int]int)

	if size1 != size2 {
		return false
	}

	for i := 0; i < size1; i++ {
		h[int(data1[i])]++
		h[int(data2[i])]--
	}

	for i := 0; i < size1; i++ {
		if h[int(data1[i])] != 0 {
			return false
		}
	}
	return true
}

func CheckPermutation3(data1 string, data2 string) bool {
	size1 := len(data1)
	size2 := len(data2)

	if size1 != size2 {
		return false
	}
	count := make([]int, 256)
	for i := 0; i < size1; i++ {
		count[data1[i]]++
		count[data2[i]]--
	}
	for i := 0; i < size1; i++ {
		if count[i] != 0 {
			fmt.Println("Not Permutation")
			return false
		}
	}
	return true
}

func main32() {
	str1 := "aaaabbbb"
	str2 := "bbaaaabb"

	fmt.Println(CheckPermutation(str1, str2))
	fmt.Println(CheckPermutation2(str1, str2))
	fmt.Println(CheckPermutation3(str1, str2))
}

/*
true
true
true
*/

func FindElementIn2DArray(data [][]int, r int, c int, value int) bool {
	row := 0
	column := c - 1
	for row < r && column >= 0 {
		if data[row][column] == value {
			return true
		} else if data[row][column] > value {
			column--
		} else {
			row++
		}
	}
	return false
}

func isAP(arr []int, size int) bool {
	if size <= 1 {
		return true
	}

	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < size; i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

func isAP2(arr []int, size int) bool {
	first := math.MaxInt32
	second := math.MaxInt32
	hs := make(Set)

	for i := 0; i < size; i++ {
		if arr[i] < first {
			second = first
			first = arr[i]
		} else if arr[i] < second {
			second = arr[i]
		}
	}

	diff := second - first

	for i := 0; i < size; i++ {
		_, ok := hs[arr[i]]
		if ok {
			return false
		}
		hs.Add(arr[i])
	}
	for i := 0; i < size; i++ {
		value := first + i*diff
		_, ok := hs[value]
		if !ok {
			return false
		}
	}
	return true
}

func isAP3(arr []int, size int) bool {
	first := math.MaxInt32
	second := math.MaxInt32
	count := make([]int, size)
	index := -1
	for i := 0; i < size; i++ {
		if arr[i] < first {
			second = first
			first = arr[i]
		} else if arr[i] < second {
			second = arr[i]
		}
	}
	diff := second - first

	for i := 0; i < size; i++ {
		index = (arr[i] - first) / diff
		if index > size-1 || count[index] != 0 {
			return false
		}
		count[index] = 1
	}

	for i := 0; i < size; i++ {
		if count[i] != 1 {
			return false
		}
	}
	return true
}

func main33() {
	first := []int{3, 6, 9, 12, 15}
	size := len(first)
	fmt.Println("isAP :", isAP(first, size))
	fmt.Println("isAP :", isAP2(first, size))
	fmt.Println("isAP :", isAP3(first, size))
}

/*
isAP : true
isAP : true
isAP : true
*/

func findBalancedPoint(arr []int, size int) int {
	first := 0
	second := 0
	for i := 1; i < size; i++ {
		second += arr[i]
	}

	for i := 0; i < size; i++ {
		if first == second {
			return i
		}
		if i < size-1 {
			first += arr[i]
		}
		second -= arr[i+1]
	}
	return -1
}

// Testing code
func main34() {
	arr := []int{-7, 1, 5, 2, -4, 3, 0}
	fmt.Println("BalancedPoint : ", findBalancedPoint(arr, len(arr)))

}

/*
BalancedPoint :  3
*/

func findFloor(arr []int, size int, value int) (int, bool) {
	start := 0
	stop := size - 1
	var mid int
	for start <= stop {
		mid = (start + stop) / 2
		/*
		* search value is equal to arr[mid] value.. search value is greater than mid
		* index value and less than mid+1 index value. value is greater than
		* arr[size-1] then floor is arr[size-1]
		 */
		if arr[mid] == value || (arr[mid] < value && (mid == size-1 || arr[mid+1] > value)) {
			return arr[mid], true
		} else if arr[mid] < value {
			start = mid + 1
		} else {
			stop = mid - 1
		}
	}
	return -1, false
}

func findCeil(arr []int, size int, value int) (int, bool) {
	start := 0
	stop := size - 1
	var mid int

	for start <= stop {
		mid = (start + stop) / 2
		/*
		* search value is equal to arr[mid] value.. search value is less than mid index
		* value and greater than mid-1 index value. value is less than arr[0] then ceil
		* is arr[0]
		 */
		if arr[mid] == value || (arr[mid] > value && (mid == 0 || arr[mid-1] < value)) {
			return arr[mid], true
		} else if arr[mid] < value {
			start = mid + 1
		} else {
			stop = mid - 1
		}
	}
	return -1, false
}

func ClosestNumber(arr []int, size int, num int) int {
	start := 0
	stop := size - 1
	output := -1
	minDist := math.MaxInt32
	var mid int

	for start <= stop {
		mid = (start + stop) / 2
		if minDist > abs(arr[mid]-num) {
			minDist = abs(arr[mid] - num)
			output = arr[mid]
		}
		if arr[mid] == num {
			break
		} else if arr[mid] > num {
			stop = mid - 1
		} else {
			start = mid + 1
		}
	}
	return output
}

func main35() {
	arr := []int{-7, 1, 2, 3, 6, 8, 10}
	fmt.Println(findFloor(arr, len(arr), 4))
	fmt.Println(findCeil(arr, len(arr), 4))
	fmt.Println("ClosestNumber : ", ClosestNumber(arr, len(arr), 4))
}

/*
3 true
6 true
ClosestNumber :  3
*/

func DuplicateKDistance(arr []int, size int, k int) bool {
	hm := make(map[int]int)

	for i := 0; i < size; i++ {
		val, ok := hm[arr[i]]
		if ok && i-val <= k {
			fmt.Println("Value:", arr[i], " Index: ", hm[arr[i]], " & ", i)
			return true
		} else {
			hm[arr[i]] = i
		}
	}
	return false
}

// Testing code
func main36() {
	arr := []int{1, 2, 3, 1, 4, 5}
	DuplicateKDistance(arr, len(arr), 3)
}

/*
Value: 1  Index:  0  &  3
*/

func FrequencyCounts(arr []int) {
	size := len(arr)
	hm := make(map[int]int)
	for i := 0; i < size; i++ {
		val, ok := hm[arr[i]]
		if ok {
			hm[arr[i]] = val + 1
		} else {
			hm[arr[i]] = 1
		}
	}
	for key, val := range hm {
		fmt.Print("(", key, " : ", val, ") ")
	}
	fmt.Println()
}

func FrequencyCounts2(arr []int) {
	size := len(arr)
	sort.Ints(arr)
	count := 1
	for i := 1; i < size; i++ {
		if (arr[i] == arr[i - 1]) {
            count++
		} else {
            fmt.Print("(", arr[i - 1], " : ", count, ") ")
            count = 1
        }
	}
    fmt.Print("(", arr[size - 1], " : ", count, ")")
}

func FrequencyCounts3(arr []int) {
	size := len(arr)
	aux := make([]int, size+1)
	for i := 0; i < size; i++ {
		aux[arr[i]] += 1
	}
	for i := 0; i < size+1; i++ {
		if aux[i] > 0 {
			fmt.Print("(", i, " : ", aux[i], ") ")
		}
	}
	fmt.Println()
}

func FrequencyCounts4(arr []int) {
	size := len(arr)
	var index int
	for i := 0; i < size; i++ {
		for arr[i] > 0 {
			index = arr[i] - 1
			if arr[index] > 0 {
				arr[i] = arr[index]
				arr[index] = -1
			} else {
				arr[index] -= 1
				arr[i] = 0
			}
		}
	}
	for i := 0; i < size; i++ {
		if arr[i] != 0 {
			fmt.Print("(", i+1, " : ", abs(arr[i]), ") ")
		}
	}
	fmt.Println()
}

func main37() {
	arr := []int{1, 2, 2, 2, 1}
	FrequencyCounts(arr)
	FrequencyCounts2(arr)
	arr = []int{1, 2, 2, 2, 1}
	FrequencyCounts3(arr)
	FrequencyCounts4(arr)
}

/*
(1 : 2) (2 : 3)
(1 : 2) (2 : 3)
(1 : 2) (2 : 3)
(1 : 2) (2 : 3)
*/

func KLargestElements(arrIn []int, size int, k int) {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = arrIn[i]
	}
	sort.Ints(arr)
	fmt.Print("KLargestElements are ::")
	for i := 0; i < size; i++ {
		if arrIn[i] >= arr[size-k] {
			fmt.Print(arrIn[i], " ")
		}
	}
}

func QuickSelectUtil(arr []int, lower int, upper int, k int) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for arr[lower] <= pivot {
			lower++
		}
		for arr[upper] > pivot {
			upper--
		}
		if lower < upper {
			arr[upper], arr[lower] = arr[lower], arr[upper]
		}
	}

	arr[upper], arr[start] = arr[start], arr[upper] // upper is the pivot position
	if k < upper {
		QuickSelectUtil(arr, start, upper-1, k) // pivot -1 is the upper for left sub array.
	}
	if k > upper {
		QuickSelectUtil(arr, upper+1, stop, k) // pivot + 1 is the lower for right sub array.
	}
}

func KLargestElements2(arrIn []int, size int, k int) {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = arrIn[i]
	}

	QuickSelectUtil(arr, 0, size-1, size-k)

	fmt.Print("KLargestElements are ::")
	for i := 0; i < size; i++ {
		if arrIn[i] >= arr[size-k] {
			fmt.Print(arrIn[i], " ")
		}
	}
}

func main38() {
	arr := []int{1, 3, 4, 2, 2, 1, 5, 9, 3}
	KLargestElements(arr, len(arr), 7)
	KLargestElements2(arr, len(arr), 7)
}

/*
KLargestElements are ::3 4 2 2 5 9 3
KLargestElements are ::3 4 2 2 5 9 3
*/


func FixPoint(arr []int, size int) int {
	for i := 0; i < size; i++ {
		if arr[i] == i {
			return i
		}
	} // fix point not found so return invalid index
	return -1
}

func FixPoint2(arr []int, size int) int {
	low := 0
	high := size - 1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if arr[mid] == mid {
			return mid
		} else if arr[mid] < mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	/* fix point not found so return invalid index */
	return -1
}

func main39() {
	arr := []int{-1, 0, 2, 3, 6, 7, 9, 10, 18}
	fmt.Println(FixPoint(arr, len(arr)))
	fmt.Println(FixPoint2(arr, len(arr)))
}

/*
2
2
*/

func subArraySums(arr []int, size int, value int) {
	start := 0
	end := 0
	sum := 0

	for (start < size && end < size) {
        if (sum < value) {
            sum += arr[end]
            end += 1
        } else {
            sum -= arr[start]
            start += 1;
        }

        if (sum == value) {
            fmt.Print("(", start, " to ", (end - 1), ") ")
        }
    }
}

func main40() {
	arr := []int{15, 5, 5, 20, 10, 5, 5, 20, 10, 10}
	subArraySums(arr, len(arr), 20)
}

/*
(0 to 1) (3 to 3) (4 to 6) (7 to 7) (8 to 9) 
*/

func MaxConSub(arr []int, size int) int {
	currMax := 0
	maximum := 0
	for i := 0; i < size; i++ {
		currMax = max(arr[i], currMax+arr[i])
		if currMax < 0 {
			currMax = 0
		}
		if maximum < currMax {
			maximum = currMax
		}
	}
	return maximum
}

func main41() {
	arr := []int{1, -2, 3, 4, -4, 6, -4, 8, 2}
	fmt.Println(MaxConSub(arr, len(arr)))
}

/*
15
*/

func MaxConSubArr(A []int, sizeA int, B []int, sizeB int) int {
	currMax := 0
	maximum := 0
	hs := make(Set)

	for i := 0; i < sizeB; i++ {
		hs.Add(B[i])
	}

	for i := 0; i < sizeA; i++ {
		_, ok := hs[A[i]]
		if ok {
			currMax = 0
		} else {
			currMax = max(A[i], currMax+A[i])
		}
		if currMax < 0 {
			currMax = 0
		}
		if maximum < currMax {
			maximum = currMax
		}
	}
	return maximum
}

func MaxConSubArr2(A []int, sizeA int, B []int, sizeB int) int {
	sort.Ints(B)
	currMax := 0
	maximum := 0

	for i := 0; i < sizeA; i++ {
		if BinarySearch(B, A[i]) {
			currMax = 0
		} else {
			currMax = max(A[i], currMax+A[i])
			if currMax < 0 {
				currMax = 0
			}
			if maximum < currMax {
				maximum = currMax
			}
		}
	}
	return maximum
}

func main42() {
	arr := []int{1, 2, 3, 4, 4, 6, 4, 8, 2}
	arr2 := []int{2, 4, 8, 18, 10}

	fmt.Println(MaxConSubArr(arr, len(arr), arr2, len(arr2)))
	fmt.Println(MaxConSubArr2(arr, len(arr), arr2, len(arr2)))
}

/*
6
6
*/

func RainWater(arr []int, size int) int {
	leftHigh := make([]int, size)
	rightHigh := make([]int, size)

	max := arr[0]
	leftHigh[0] = arr[0]
	for i := 1; i < size; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		leftHigh[i] = max
	}
	max = arr[size-1]
	rightHigh[size-1] = arr[size-1]
	for i := (size - 2); i >= 0; i-- {
		if max < arr[i] {
			max = arr[i]
		}
		rightHigh[i] = max
	}

	water := 0
	for i := 0; i < size; i++ {
		water += min(leftHigh[i], rightHigh[i]) - arr[i]
	}
	fmt.Println("Water :", water)
	return water
}

func RainWater2(arr []int, size int) int {
	water := 0
	leftMax := 0
	rightMax := 0
	left := 0
	right := size - 1

	for left <= right {
		if arr[left] < arr[right] {
			if arr[left] > leftMax {
				leftMax = arr[left]
			} else {
				water += leftMax - arr[left]
			}
			left += 1
		} else {
			if arr[right] > rightMax {
				rightMax = arr[right]
			} else {
				water += rightMax - arr[right]
			}
			right -= 1
		}
	}
	fmt.Println("Water :", water)
	return water
}

func main43() {
	arr := []int{4, 0, 1, 5}
	RainWater(arr, len(arr))
	RainWater2(arr, len(arr))
}

/*
Water : 7
Water : 7
*/

func SeparateEvenAndOdd(arr []int, size int) {
	left := 0
	right := size - 1
	for left < right {
		if arr[left]%2 == 0 {
			left++
		} else if arr[right]%2 == 1 {
			right--
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}

func Main50() {
	first := []int{1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30}
	SeparateEvenAndOdd(first, len(first))
	for _, val := range first {
		fmt.Print(val, " ")
	}
}

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5()
	// main6()
	// main7()
	// main8()
	// main9()
	// main10()
	// Main10A()
	// main11()
	// main12()
	// main13()
	// main14()
	// main15()
	// main16()
	// main17()
	// main18()
	// main19()
	// main20()
	// main21()
	// main22()
	// main23()
	// main24()
	// main25()
	// main26()
	// main27()
	// main28()
	// main29()
	// main30()
	// main30A()
	// main31()
	// main32()
	// main33()
	// main34()
	// main35()
	// main36()
	// main37()
	// main38()
	// main39()
	main40()
	// main41()
	// main42()
	// main43()
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// *********************
/*Set type is defined using map */
type Set map[interface{}]bool

/*Add function of Set type */
func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}

/*Remove function of Set type */
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

/*Find function of Set type */
func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}

//*********************
