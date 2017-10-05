package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		first := []int{1, 3, 5, 7, 9, 7, 25, 21, 30}
		second := []int{2, 4, 6, 8, 10, 12, 14, 8, 21, 23, 24}
		fmt.Println(linearSearchUnsorted(first, 8))
		fmt.Println(linearSearchUnsorted(second, 8))
		fmt.Println(Binarysearch(second, 18))
		fmt.Println(Binarysearch(second, 16))

		printRepeating(second)
		printRepeating(first)
		printRepeating2(second)
		printRepeating2(first)
		printRepeating3(second)
		printRepeating3(first)
		printRepeating4(second, 100)
		printRepeating4(first, 100)

		fmt.Println(getMax(first))
		fmt.Println(getMax(second))
		fmt.Println(getMax2(first))
		fmt.Println(getMax2(second))
		fmt.Println(getMax3(first,100))
		fmt.Println(getMax3(second,100))
	*/
	/*
		first := []int{1, 3, 5, 7, 9, 7, 25, 21, 30}
		second := []int{1, 3, 3, 7, 3, 7, 2, 3, 3}
		fmt.Println(getMajority(first))
		fmt.Println(getMajority(second))
		fmt.Println(getMajority2(first))
		fmt.Println(getMajority2(second))
		fmt.Println(getMajority3(first))
		fmt.Println(getMajority3(second))
	*/
	/*
		first := []int{1, 3, 5, 7, 2, 4, 6, 10, 9}
		second := []int{1, 3, 2, 5, 4, 7, 6, 8, 9}
		fmt.Println(findMissingNumber(first))
		fmt.Println(findMissingNumber(second))
		fmt.Println(findMissingNumber2(first,10))
		fmt.Println(findMissingNumber2(second,9))
		fmt.Println(FindPair(first, 7))
		fmt.Println(FindPair(second, 7))
		fmt.Println(FindPair2(first, 7))
		fmt.Println(FindPair2(second, 7))
		fmt.Println(FindPair3(first, 7))
		fmt.Println(FindPair3(second, 7))
	*/
	/*
		first := []int{1, 3, 5, -7, 2, -4, 6, -10, -9}
		minabsSumPair(first)
		minabsSumPair2(first)
	*/
	/*
		first := []int{1, 3, 5, 7, 4, 1, -1, -9}
		fmt.Println(SearchBitonicArrayMax(first))
		fmt.Println(FindMaxBitonicArray(first))
		fmt.Println(SearchBitonicArray(first, 6))
	*/
	/*
		first := []int{1, 3, 5, 7, 2, 4, 6, 8, 9}
		second := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println(checkPermutation(first, second))
		fmt.Println(checkPermutation2(first, second))
	*/
	/*
		first := []int{1, 3, 5, 7, 3, 4, 3, 8, 9}
		second := []int{1, 2, 3, 3, 5, 5, 5, 6, 6}

		fmt.Println(findKeyCount(first, 3))
		fmt.Println(findKeyCount(second, 3))
		fmt.Println(findKeyCount(second, 5))
		fmt.Println(findKeyCount(second, 6))
	*/
	/*
		first := []int{1, 3, 5, 7, 3, 4, 3, 8, 9}
		maxProfit(first)
		seperateEvenAndOdd(first)
		fmt.Println(first)
		fmt.Println(getMedian(first))
		fmt.Println(first)
		second := []int{1, 2, 3, 3, 5, 5, 5, 6, 6}
		sort.Ints(second)
		fmt.Println(findMedian(first, second))
	*/
	/*
		first := []int{0, 0, 1, 1, 1, 1, 1, 1, 1}
		fmt.Println(BinarySearch01(first))

		second := []int{5, 5, 6, 6, 1, 2, 3, 3, 4}
		fmt.Println(BinarySearchRotateArray(second, 4))
		fmt.Println(BinarySearchRotateArray(second, 2))
		fmt.Println(BinarySearchRotateArray(second, 5))
		fmt.Println(BinarySearchRotateArray(second, 7))

		first := []int{1, 3, 5, 7,6,6, 3, 4, 3, 8, 9}
		fmt.Println(FirstRepeated(first))
	*/

}

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

func Binarysearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0

	for low <= high {
		mid = low + (high-low)/2
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

func BinarySearchRecursive(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}

func printRepeating(data []int) {
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Print(" ", data[i])
			}
		}
	}
	fmt.Println()
}

func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data) // Sort(data,size)
	fmt.Print("Repeating elements are : ")

	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print(" ", data[i])
		}
	}
	fmt.Println()
}

type Set map[interface{}]bool

func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}

func printRepeating3(data []int) {
	s := make(Set)
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		if s.Find(data[i]) {
			fmt.Print(" ", data[i])
		} else {
			s.Add(data[i])
		}
	}
	fmt.Println()
}

func printRepeating4(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)
	for i := 0; i < size; i++ {
		count[i] = 0
	}
	fmt.Println("Repeating elements are : ")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Println(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}

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

func findMissingNumber(data []int) (int, bool) {
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
func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}
func findMissingNumber2(data []int, dataRange int) int {
	size := len(data)
	xorSum := 0
	// get the XOR of all the numbers from 1 to dataRange
	for i := 1; i <= dataRange; i++ {
		xorSum ^= i
	}
	// loop through the array and get the XOR of elements
	for i := 0; i < size; i++ {
		xorSum ^= data[i]
	}
	return xorSum
}

func FindPair(data []int, value int) bool {
	size := len(data)
	ret := false
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (data[i] + data[j]) == value {
				fmt.Println("The pair is : ", data[i], ",", data[j])
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
	for first < second {
		curr := data[first] + data[second]
		if curr == value {
			fmt.Println("The pair is ", data[first], ",", data[second])
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
	for i := 0; i < size; i++ {
		if s.Find(value - data[i]) {
			fmt.Println(i, "The pair is : ", data[i], " , ", (value - data[i]))
			ret = true
		} else {
			s.Add(data[i])
		}
	}
	return ret
}

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
	fmt.Println(" The two elements with minimum sum are : ", data[minFirst], " , ", data[minSecond])
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
			minSum = abs(sum) /// just corrected......hemant
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
	fmt.Println(" The two elements with minimum sum are : ", data[minFirst], " , ", data[minSecond])
}

func SearchBitonicArrayMax(data []int) (int, bool) {
	size := len(data)
	start := 0
	end := size - 1
	mid := (start + end) / 2
	maximaFound := 0
	if size < 3 {
		fmt.Println("InvalidInput")
		return 0, false
	}
	for start <= end {
		mid := (start + end) / 2

		if data[mid-1] < data[mid] && data[mid+1] < data[mid] { //maxima
			maximaFound = 1
			break
		} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] { // increasing
			start = mid + 1
		} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] { // decreasing
			end = mid - 1
		} else {
			break
		}
	}
	if maximaFound == 0 {
		fmt.Println("NoMaximaFound")
		return 0, false
	}
	return mid, true
}

func SearchBitonicArray(data []int, key int) int {
	size := len(data)
	maxIndex, _ := FindMaxBitonicArray(data)
	k := BinarySearch(data, 0, maxIndex, key, true)
	if k != -1 {
		return k
	} else {
		return BinarySearch(data, maxIndex+1, size-1, key, false)
	}
}

func FindMaxBitonicArray(data []int) (int, bool) {
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
			start = mid + 1
		} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] /* decreasing */ {
			end = mid - 1
		} else {
			break
		}
	}
	fmt.Println("error")
	return -1, false
}

func BinarySearch(data []int, start int, end int, key int, isInc bool) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] {
		return mid
	}
	if isInc != false && key < data[mid] || isInc == false && key > data[mid] {
		return BinarySearch(data, start, mid-1, key, isInc)
	}
	return BinarySearch(data, mid+1, end, key, isInc)
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

func maxProfit(stocks []int) {
	size := len(stocks)
	buy := 0
	sell := 0
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
	fmt.Println("Purchase day is- ", buy, " at price ", stocks[buy])
	fmt.Println("Sell day is- ", sell, " at price ", stocks[sell])
	fmt.Println("Max Profit :: ", maxProfit)
}

func getMedian(data []int) int {
	size := len(data)
	sort.Ints(data)
	return data[size/2]
}

func findMedian(dataFirst []int, dataSecond []int) int {
	sizeFirst := len(dataFirst)
	sizeSecond := len(dataSecond)
	medianIndex := ((sizeFirst + sizeSecond) + (sizeFirst+sizeSecond)%2) / 2 // cealing function.
	i := 0
	j := 0
	count := 0
	for count < medianIndex-1 {
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

func BinarySearchRotateArrayUtil(data []int, start int, end int, key int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] {
		return mid
	}
	if data[mid] > data[start] {
		if data[start] <= key && key < data[mid] {
			return BinarySearchRotateArrayUtil(data, start, mid-1, key)
		}
		return BinarySearchRotateArrayUtil(data, mid+1, end, key)
	} else {
		if data[mid] < key && key <= data[end] {
			return BinarySearchRotateArrayUtil(data, mid+1, end, key)
		}
		return BinarySearchRotateArrayUtil(data, start, mid-1, key)
	}
}

func BinarySearchRotateArray(data []int, key int) int {
	size := len(data)
	return BinarySearchRotateArrayUtil(data, 0, size-1, key)
}

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

func transformArrayAB1(str string) string {
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

func CheckPermutation(data1 []int, data2 []int) bool {
	size1 := len(data1)
	size2 := len(data2)

	if size1 != size2 {
		return false
	}

	sort.Ints(data1)
	sort.Ints(data2)

	for i := 0; i < size1; i++ {
		if data1[i] != data2[i] {
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

func checkPermutation2(data1 []int, data2 []int) bool {
	size1 := len(data1)
	size2 := len(data2)
	h := make(map[int]int)

	if size1 != size2 {
		return false
	}

	for i := 0; i < size1; i++ {
		h[data1[i]]++
		h[data2[i]]--
	}

	for i := 0; i < size1; i++ {
		if h[data1[i]] != 0 {
			return false
		}
	}
	return true
}

func removeDuplicates(data []int) int {
	j := 0
	size := len(data)
	if size == 0 {
		return 0
	}

	sort.Ints(data)
	for i := 1; i < size; i++ {
		if data[i] != data[j] {
			j++
			data[j] = data[i]
		}
	}
	return j + 1
}

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
