package main

import (
	"fmt"
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

/*Binarysearch algorithm */
func Binarysearch(data []int, value int) bool {
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

/*BinarySearchRecursive algorithm */
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

func main1(){
	first := []int{1, 3, 5, 7, 9, 7, 25, 21, 30}
	second := []int{2, 4, 6, 8, 10, 12, 14, 18, 21, 23, 24}
	fmt.Println(linearSearchUnsorted(first, 8))
	fmt.Println(linearSearchUnsorted(second, 8))
	fmt.Println(linearSearchSorted(first, 8))
	fmt.Println(linearSearchSorted(second, 8))
	fmt.Println(Binarysearch(second, 18))
	fmt.Println(Binarysearch(second, 16))
	fmt.Println(BinarySearchRecursive(first, 8));
	fmt.Println(BinarySearchRecursive(first, 25));

}

/*FirstRepeated function*/
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

func main2() {
	first := []int { 34, 56, 77, 1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 34, 20, 30 };
	fmt.Println(FirstRepeated(first));
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

func main3() {
	first := []int{1, 3, 5, 7, 9, 7, 25, 21, 30}
	second := []int{2, 4, 6, 8, 10, 12, 14, 18, 21, 23, 24}
	printRepeating(second)
	printRepeating(first)
	printRepeating2(second)
	printRepeating2(first)
	printRepeating3(second)
	printRepeating3(first)
	printRepeating4(second, 100)
	printRepeating4(first, 100)
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

func main4() {
	first := []int { 1, 3, 5, 3, 9, 1, 30 }
	ret := removeDuplicates(first);
	for i := 0; i < ret; i++ {
		fmt.Print(first[i] , " ")
	}
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

func findMissingNumber3(arr []int, size int, upperRange int) int {
	st := make(Set) 
	i := 0
	for i < size {
		st.Add(arr[i])
		i += 1
	}
	i = 1
	for i <= upperRange {
		if (st.Find(i) == false) {
			return i
		}
		i += 1
	}
	fmt.Println("NoNumberMissing")
	return -1
}

func main5() {
	first := []int {1, 3, 5, 4, 6, 8, 7 }
	fmt.Println(findMissingNumber(first))
	fmt.Println(findMissingNumber2(first, 8))
	fmt.Println(findMissingNumber3(first, len(first), 8))
}

func MissingValues(arr []int, size int) {
	sort.Ints(arr)
	value := arr[0]
	i := 0
	for i < size {
		if (value == arr[i]) {
			value += 1
			i += 1
		} else {
			fmt.Println(value)
			value += 1
		}
	}
}

func MissingValues2(arr []int, size int) {
	ht := make(Set)
	minVal := 999999
	maxVal := -999999
	i := 0
	for i = 0; i < size; i++ {
		ht.Add(arr[i])
		if (minVal > arr[i]) {
			minVal = arr[i]
		}
		if (maxVal < arr[i]) {
			maxVal = arr[i]
		}
	}
	for i = minVal; i < maxVal + 1; i++ {
		if (ht.Find(i) == false) {
			fmt.Println(i)
		}
	}
}

func main6() {
	arr := []int { 1, 9, 2, 8, 3, 7, 4, 6 }
	size := len(arr)
	MissingValues(arr, size)
	MissingValues2(arr, size)
}

func OddCount(arr []int, size int) {
	ctr := make(map[int]int)
	count := 0

	for i := 0; i < size; i++ {
		val, ok := ctr[arr[i]]
		if ok {
			ctr[arr[i]] = val + 1;
		} else {
			ctr[arr[i]] = 1
		}
	}
	for i := 0; i < size; i++ {
		val, ok := ctr[arr[i]]
		if (ok && (val % 2 == 1)) {
			fmt.Println(arr[i])
			count++
			delete(ctr, arr[i])
		}
	}
	fmt.Println("Odd count is :: " , count)
}

func OddCount2(arr []int, size int) {
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
	setBit = xorSum &^(xorSum - 1)

	/*
		* Dividing elements in two group: Elements having setBit bit as 1. Elements
		* having setBit bit as 0. Even elements cancelled themselves if group and we
		* get our numbers.
		*/
	for i := 0; i < size; i++ {
		if ((arr[i] & setBit) != 0) {
			first ^= arr[i]
		} else {
			second ^= arr[i]
		}
	}
	fmt.Println(first + second)
}


func SumDistinct(arr []int, size int) {
	sum := 0
	sort.Ints(arr)
	for i := 0; i < (size - 1); i++ {
		if (arr[i] != arr[i + 1]) {
			sum += arr[i]
		}
	}
	sum += arr[size - 1]
	fmt.Println(sum)
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

func main7() {
	first := []int{ 1, 5, -10, 3, 2, -6, 8, 9, 6 }
	minAbsSumPair2(first)
	minAbsSumPair(first)

}

/*FindPair function*/
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

/*FindPair2 function*/
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

/*FindPair3 function*/
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

func main8() {
	first := []int {1, 5, 4, 3, 2, 7, 8, 9, 6 }
	fmt.Println(FindPair(first, 8))
	fmt.Println(FindPair2(first, 8))
	fmt.Println(FindPair3(first, 8))

}

func FindDifference(arr []int, size int, value int) bool {
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (abs(arr[i] - arr[j]) == value) {
				fmt.Println("The pair is:: " , arr[i] , " & " , arr[j])
				return true
			}
		}
	}
	return false
}

func FindDifference2(arr []int, size int, value int) bool {
	first := 0
	second := 0
	var diff int
	sort.Ints(arr)
	for (first < size && second < size) {
		diff = abs(arr[first] - arr[second])
		if (diff == value) {
			fmt.Println("The pair is::" , arr[first] , " & " , arr[second])
			return true
		} else if (diff > value) {
			first += 1
		} else {
			second += 1
		}
	}
	return false
}

func findMinDiff(arr []int, size int) int {
	sort.Ints(arr)
	diff := 9999999

	for i := 0; i < (size - 1); i++ {
		if ((arr[i + 1] - arr[i]) < diff) {
			diff = arr[i + 1] - arr[i]
		}
	}
	return diff
}

func MinDiffPair(arr1 []int, size1 int, arr2 []int, size2 int) int {
	minDiff := 9999999
	first := 0
	second := 0
	out1 := 0 
	out2 := 0
	diff := 0
	sort.Ints(arr1);
	sort.Ints(arr2);
	for (first < size1 && second < size2) {
		diff = abs(arr1[first] - arr2[second]);
		if (minDiff > diff) {
			minDiff = diff
			out1 = arr1[first]
			out2 = arr2[second]
		}
		if (arr1[first] < arr2[second]) {
			first += 1
		} else {
			second += 1
		}
	}
	fmt.Println("The pair is :: " , out1 + out2)
	fmt.Println("Minimum difference is :: " , minDiff)
	return minDiff
}

func main9() {
	first := []int {1, 5, 4, 3, 2, 7, 8, 9, 6 }
	fmt.Println(FindDifference(first, len(first), 6))
	fmt.Println(FindDifference2(first, len(first), 6))
	fmt.Println(findMinDiff(first, len(first)))
	fmt.Println(MinDiffPair(first, len(first), first, len(first)))
}

func ClosestPair(arr []int, size int, value int) {
	diff := 999999
	first := -1
	second := -1
	curr := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			curr = abs(value - (arr[i] + arr[j]))
			if (curr < diff) {
				diff = curr
				first = arr[i]
				second = arr[j]
			}
		}
	}
	fmt.Println("closest pair is ::" , first , second);
}

func ClosestPair2(arr []int, size int, value int) {
	first := 0
	second := 0
	start := 0
	stop := size - 1
	diff := 0
	curr := 0
	sort.Ints(arr)
	diff = 9999999
	
	for (start < stop) {
		curr = (value - (arr[start] + arr[stop]))
		if (abs(curr) < diff) {
			diff = abs(curr)
			first = arr[start]
			second = arr[stop]
		}
		if (curr == 0) {
			break
		} else if (curr > 0) {
			start += 1
		} else {
			stop -= 1
		}
	}

	fmt.Println("closest pair is :: " , first , second)
}

func main10() {
	first := []int {1, 5, 4, 3, 2, 7, 8, 9, 6 }
	ClosestPair(first, len(first), 6)
	ClosestPair2(first, len(first), 6)
}

func SumPairRestArray(arr []int, size int) bool {
	var total, low, high, curr, value int
	sort.Ints(arr)
	total = 0
	for i := 0; i < size; i++ {
		total += arr[i]
	}
	value = total / 2
	low = 0
	high = size - 1
	for (low < high) {
		curr = arr[low] + arr[high]
		if (curr == value) {
			fmt.Println("Pair is :: " , arr[low] , arr[high])
			return true
		} else if (curr < value) {
			low += 1
		} else {
			high -= 1
		}
	}
	return false
}

func ZeroSumTriplets(arr []int, size int) {
	for i := 0; i < (size - 2); i++ {
		for j := i + 1; j < (size - 1); j++ {
			for k := j + 1; k < size; k++ {
				if (arr[i] + arr[j] + arr[k] == 0) {
					fmt.Println("Triplet :: " , arr[i] , arr[j] , arr[k])
				}
			}
		}
	}
}

func ZeroSumTriplets2(arr []int, size int) {
	var start, stop int
	sort.Ints(arr)
	for i := 0; i < (size - 2); i++ {
		start = i + 1
		stop = size - 1

		for (start < stop) {
			if (arr[i] + arr[start] + arr[stop] == 0) {
				fmt.Println("Triplet :: " , arr[i] , arr[start] , arr[stop])
				start += 1
				stop -= 1
			} else if (arr[i] + arr[start] + arr[stop] > 0) {
				stop -= 1
			} else {
				start += 1;
			}
		}
	}
}

func findTriplet(arr []int, size int, value int) {
	for i := 0; i < (size - 2); i++ {
		for j := i + 1; j < (size - 1); j++ {
			for k := j + 1; k < size; k++ {
				if ((arr[i] + arr[j] + arr[k]) == value) {
					fmt.Println("Triplet :: ", arr[i], arr[j], arr[k]);
				}
			}
		}
	}
}

func findTriplet2(arr []int, size int, value int) {
	var start, stop int
	sort.Ints(arr)
	for i := 0; i < size - 2; i++ {
		start = i + 1
		stop = size - 1
		for (start < stop) {
			if (arr[i] + arr[start] + arr[stop] == value) {
				fmt.Println("Triplet ::" , arr[i] , arr[start] , arr[stop])
				start += 1
				stop -= 1
			} else if (arr[i] + arr[start] + arr[stop] > value) {
				stop -= 1
			} else {
				start += 1
			}
		}
	}
}

func ABCTriplet(arr []int, size int) {
	var start, stop int
	sort.Ints(arr)
	for i := 0; i < (size - 2); i++ {
		start = i + 1
		stop = size - 1
		for (start < stop) {
			if (arr[i] == arr[start] + arr[stop]) {
				fmt.Println("Triplet ::" , arr[i] , arr[start] , arr[stop]);
				start += 1
				stop -= 1
			} else if (arr[i] > arr[start] + arr[stop]) {
				stop -= 1
			} else {
				start += 1
			}
		}
	}
}

func SmallerThenTripletCount(arr []int, size int, value int) {
	var start, stop int
	count := 0
	sort.Ints(arr)

	for i := 0; i < (size - 2); i++ {
		start = i + 1
		stop = size - 1
		for (start < stop) {
			if (arr[i] + arr[start] + arr[stop] >= value) {
				stop -= 1
			} else {
				count += stop - start
				start += 1
			}
		}
	}
	fmt.Println(count);
}

func APTriplets(arr []int, size int) {
	var i, j, k int
	for i = 1; i < size - 1; i++ {
		j = i - 1
		k = i + 1
		for (j >= 0 && k < size) {
			if (arr[j] + arr[k] == 2 * arr[i]) {
				fmt.Println("Triplet ::" , arr[j] , arr[i] , arr[k])
				k += 1
				j -= 1
			} else if (arr[j] + arr[k] < 2 * arr[i]) {
				k += 1
			} else {
				j -= 1
			}
		}
	}
}

func GPTriplets(arr []int, size int) {
	var i, j, k int
	for i = 1; i < size - 1; i++ {
		j = i - 1
		k = i + 1
		for (j >= 0 && k < size) {
			if (arr[j] * arr[k] == arr[i] * arr[i]) {
				fmt.Println("Triplet is :: " , arr[j] , arr[i] , arr[k]);
				k += 1
				j -= 1
			} else if (arr[j] + arr[k] < 2 * arr[i]) {
				k += 1
			} else {
				j -= 1
			}
		}
	}
}

func numberOfTriangles(arr []int, size int) int {
	var i, j, k int 
	count := 0
	for i = 0; i < (size - 2); i++ {
		for j = i + 1; j < (size - 1); j++ {
			for k = j + 1; k < size; k++ {
				if (arr[i] + arr[j] > arr[k]) {
					count += 1
				}
			}
		}
	}
	return count
}

func numberOfTriangles2(arr []int, size int) int {
	var i, j, k int 
	count := 0
	sort.Ints(arr)

	for i = 0; i < (size - 2); i++ {
		k = i + 2;
		for j = i + 1; j < (size - 1); j++ {
			/*
				* if sum of arr[i] & arr[j] is greater arr[k] then sum of arr[i] & arr[j + 1]
				* is also greater than arr[k] this improvement make algo O(n2)
				*/
			for (k < size && arr[i] + arr[j] > arr[k]) {
				k += 1
			}

			count += k - j - 1
		}
	}
	return count
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

func main11() {
	first := []int {1, 30, 5, 13, 9, 31, 5 };
	fmt.Println(getMax(first))
	fmt.Println(getMax2(first))
	fmt.Println(getMax3(first, 50));
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

func main12() {
	first := []int {1, 5, 5, 13, 5, 31, 5 }
	fmt.Println(getMajority(first))
	fmt.Println(getMajority2(first))
	fmt.Println(getMajority3(first))
}

func getMedian(arr []int, size int) int {
	sort.Ints(arr)
	return arr[size / 2]
}



/*SearchBitonicArrayMax function*/
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


/*SearchBitonicArray function*/
func SearchBitonicArray(data []int, key int) int {
	size := len(data)
	maxIndex, _ := FindMaxBitonicArray(data)
	k := BinarySearch(data, 0, maxIndex, key, true)
	if k != -1 {
		return k
	}
	return BinarySearch(data, maxIndex+1, size-1, key, false)
}

/*FindMaxBitonicArray function*/
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


func main13() {
	first := []int {1, 5, 10, 13, 20, 30, 8, 7, 6 }
	fmt.Println(SearchBitonicArrayMax(first))
	fmt.Println(FindMaxBitonicArray(first))
	fmt.Println(SearchBitonicArray(first, 6))
	fmt.Println(SearchBitonicArray(first, 7))
}

/*BinarySearch algorithm. */
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

/* FindKeyCount2 function.*/
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


func isMajority(arr []int, size int) bool {
	majority := arr[size / 2];
	i := findFirstIndex(arr, 0, size - 1, majority)
	/*
		* we are using majority element form array so we will get some valid index
		* always.
		*/
	if (((i + size / 2) <= (size - 1)) && arr[i + size / 2] == majority) {
		return true
	} else {
		return false
	}
}

func main14() {
	first := []int {1, 5, 10, 13, 20, 30, 8, 7, 6 }
	fmt.Println(findKeyCount(first, 6))
	fmt.Println(findKeyCount2(first, 6))
}

func maxProfit(stocks []int) int {
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
	return maxProfit
}

func main15() {
	first := []int {10, 150, 6, 67, 61, 16, 86, 6, 67, 78, 150, 3, 28, 143 }
	fmt.Println(maxProfit(first))
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

func main16() {
	first := []int {1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30 }
	second := []int { 1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30 }
	fmt.Println(findMedian(first, second));
}

/*BinarySearch01 Function */
func BinarySearch01(data []int) int {
	size := len(data)
	if size == 1 && data[0] == 1 {
		return -1
	}
	return binarySearch01Util(data, 0, size-1)
}

func binarySearch01Util(data []int, start int, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if 1 == data[mid] && 0 == data[mid-1] {
		return mid
	}
	if 0 == data[mid] {
		return binarySearch01Util(data, mid+1, end)
	}
	return binarySearch01Util(data, start, mid-1)
}

func main17() {
	first := []int {0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1 }
	fmt.Println(BinarySearch01(first))
}

func RotationMaxUtil( arr []int, start int, end int) int {
	if (end <= start) {
		return arr[start]
	}
	mid := (start + end) / 2
	if (arr[mid] > arr[mid + 1]) {
		return arr[mid]
	}

	if (arr[start] <= arr[mid]) {// increasing part.
		return RotationMaxUtil(arr, mid + 1, end);
	} else {
		return RotationMaxUtil(arr, start, mid - 1);
	}
}

func RotationMax(arr []int, size int) int {
	return RotationMaxUtil(arr, 0, size - 1)
}

func FindRotationMaxUtil(arr []int, start int, end int) int {
	/* single element case. */
	if (end <= start) {
		return start
	}

	mid := (start + end) / 2
	if (arr[mid] > arr[mid + 1]) {
		return mid
	}

	if (arr[start] <= arr[mid]) { // increasing part.
		return FindRotationMaxUtil(arr, mid + 1, end)
	} else {
		return FindRotationMaxUtil(arr, start, mid - 1)
	}
}

func FindRotationMax(arr []int, size int) int {
	return FindRotationMaxUtil(arr, 0, size - 1)
}

func CountRotation(arr []int, size int) int {
	maxIndex := FindRotationMaxUtil(arr, 0, size - 1)
	return (maxIndex + 1) % size
}

/* Function */
func binarySearchRotateArrayUtil(data []int, start int, end int, key int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] {
		return mid
	}
	if data[mid] > data[start] {
		if data[start] <= key && key < data[mid] {
			return binarySearchRotateArrayUtil(data, start, mid-1, key)
		}
		return binarySearchRotateArrayUtil(data, mid+1, end, key)
	}
	if data[mid] < key && key <= data[end] {
		return binarySearchRotateArrayUtil(data, mid+1, end, key)
	}
	return binarySearchRotateArrayUtil(data, start, mid-1, key)
}

/*BinarySearchRotateArray Function*/
func BinarySearchRotateArray(data []int, key int) int {
	size := len(data)
	return binarySearchRotateArrayUtil(data, 0, size-1, key)
}

func main18() {
	first := []int {34, 56, 77, 1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30 }
	fmt.Println(BinarySearchRotateArray(first, 20))
	fmt.Println(CountRotation(first, len(first)))
	fmt.Println(first[FindRotationMax(first, len(first))])
}

func minAbsDiffAdjCircular(arr []int, size int) int {
	diff := 9999999
	if (size < 2) {
		return -1
	}

	for i := 0; i < size; i++ {
		diff = min(diff, abs(arr[i] - arr[(i + 1) % size]))
	}

	return diff
}

func main19() {
	arr := []int { 5, 29, 18, 51, 11 }
	fmt.Println(minAbsDiffAdjCircular(arr, len(arr)))
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

/*
func getMedian(data []int) int {
	size := len(data)
	sort.Ints(data)
	return data[size/2]
}
*/
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

func main20() {
	str := "aaaabbbb"
	transformArrayAB1(str)
	fmt.Println(str)
}

/*CheckPermutation function */
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

/*Search Function */
func Search(data []int, value int) bool {
	for _, key := range data {
		if key == value {
			return true
		}
	}
	return false
}

/*CheckPermutation2 function*/
func CheckPermutation2(data1 []int, data2 []int) bool {
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

// func main21() {
// 	str1 := "aaaabbbb"
// 	str2 := "bbaaaabb"

// 	fmt.Println(CheckPermutation(str1, str2));
// }

/*FindElementIn2DArray function*/
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
	if (size <= 1) {
		return true
	}

	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < size; i++ {
		if (arr[i] - arr[i - 1] != diff) {
			return false
		}
	}
	return true
}
/*
func isAP2(arr []int, size int) bool {
	first := 9999999
	second := 9999999
	hs := make(Set)
	for i := 0; i < size; i++ {
		if (arr[i] < first) {
			second = first
			first = arr[i]
		} else if (arr[i] < second) {
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
		value = first + i * diff
		_, ok := hs[arr[i]] 
		if !ok {
			return false
		}
	}
	return true
}
*/
func isAP3(arr []int, size int) bool {
	first := 9999999
	second := 9999999
	count := make([]int, size)
	index := -1
	for i := 0; i < size; i++ {
		if (arr[i] < first) {
			second = first
			first = arr[i]
		} else if (arr[i] < second) {
			second = arr[i]
		}
	}
	diff := second - first

	for i := 0; i < size; i++ {
		index = (arr[i] - first) / diff
	}
	if (index > size - 1 || count[index] != 0) {
		return false
	}
	count[index] = 1

	for i := 0; i < size; i++ {
		if (count[i] != 1) {
			return false
		}
	}
	return true
}

func findBalancedPoint(arr []int, size int) int {
	first := 0
	second := 0
	for i := 1; i < size; i++ {
		second += arr[i]
	}

	for i := 0; i < size; i++ {
		if (first == second) {
			fmt.Println(i)
			return i
		}
		if (i < size - 1) {
			first += arr[i]
		}
		second -= arr[i + 1]
	}
	return -1
}

// Testing code
func main22() {
	arr := []int { -7, 1, 5, 2, -4, 3, 0 }
	fmt.Println(findBalancedPoint(arr, len(arr)))

}


func findFloor(arr []int, size int, value int) int {
	start := 0
	stop := size - 1
	var mid int
	for (start <= stop) {
		mid = (start + stop) / 2
		/*
			* search value is equal to arr[mid] value.. search value is greater than mid
			* index value and less than mid+1 index value. value is greater than
			* arr[size-1] then floor is arr[size-1]
			*/
		if (arr[mid] == value || (arr[mid] < value && (mid == size - 1 || arr[mid + 1] > value))) {
			return mid
		} else if (arr[mid] < value) {
			start = mid + 1
		} else {
			stop = mid - 1
		}
	}
	return -1
}

func findCeil(arr []int, size int, value int) int {
	start := 0
	stop := size - 1
	var mid int

	for (start <= stop) {
		mid = (start + stop) / 2
		/*
			* search value is equal to arr[mid] value.. search value is less than mid index
			* value and greater than mid-1 index value. value is less than arr[0] then ceil
			* is arr[0]
			*/
		if (arr[mid] == value || (arr[mid] > value && (mid == 0 || arr[mid - 1] < value))) {
			return mid
		} else if (arr[mid] < value) {
			start = mid + 1
		} else {
			stop = mid - 1
		}
	}
	return -1
}

func ClosestNumber(arr []int, size int, num int) int {
	start := 0
	stop := size - 1
	output := -1
	minDist := 9999
	var mid int

	for (start <= stop) {
		mid = (start + stop) / 2
		if (minDist > abs(arr[mid] - num)) {
			minDist = abs(arr[mid] - num)
			output = arr[mid]
		}
		if (arr[mid] == num) {
			break
		} else if (arr[mid] > num) {
			stop = mid - 1
		} else {
			start = mid + 1
		}
	}
	return output
}

func DuplicateKDistance(arr []int, size int, k int) bool {
	hm := make(map[int]int)

	for i := 0; i < size; i++ {
		val, ok := hm[arr[i]]

		if (ok && i - val <= k) {
			fmt.Println("Value:" , arr[i] , " Index: " , hm[arr[i]] , " & " , i);
			return true
		} else {
			hm[arr[i]] = i
		}
	}
	return false
}

// Testing code
	
func main23() {
	arr := []int { 1, 2, 3, 1, 4, 5 }
	DuplicateKDistance(arr, len(arr), 3)
}


func frequencyCounts(arr []int, size int) {
	var index int
	for i := 0; i < size; i++ {
		for (arr[i] > 0) {
			index = arr[i] - 1
			if (arr[index] > 0) {
				arr[i] = arr[index]
				arr[index] = -1
			} else {
				arr[index] -= 1
				arr[i] = 0
			}
		}
	}
	for i := 0; i < size; i++ {
		fmt.Println((i + 1) , abs(arr[i]));
	}
}

func KLargestElements(arrIn []int, size int, k int) int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = arrIn[i]
	}

	sort.Ints(arr)
	for i := 0; i < size; i++ {
		if (arrIn[i] >= arr[size - k]) {
			fmt.Println(arrIn[i])
			return arrIn[i]
		}
	}
	return -1
}

func QuickSelectUtil(arr []int, lower int, upper int, k int) {
	if (upper <= lower) {
		return;
	}

	pivot := arr[lower]
	start := lower
	stop := upper

	for (lower < upper) {
		for (arr[lower] <= pivot) {
			lower++
		}
		for (arr[upper] > pivot) {
			upper--
		}
		if (lower < upper) {
			arr[upper], arr[lower] = arr[lower], arr[upper] 
		}
	}

	arr[upper], arr[start] = arr[start], arr[upper] // upper is the pivot position
	if (k < upper) {
		QuickSelectUtil(arr, start, upper - 1, k) // pivot -1 is the upper for left sub array.
	}
	if (k > upper) {
		QuickSelectUtil(arr, upper + 1, stop, k) // pivot + 1 is the lower for right sub array.
	}
}

func KLargestElements2(arrIn []int, size int, k int) int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = arrIn[i]
	}

	QuickSelectUtil(arr, 0, size - 1, size - k)
	for i := 0; i < size; i++ {
		if (arrIn[i] >= arr[size - k]) {
			fmt.Println(arrIn[i])
			return arrIn[i]
		}
	}
	return -1
}

/* linear search method */
func FixPoint(arr []int, size int) int {
	for i := 0; i < size; i++ {
		if (arr[i] == i) {
			return i
		}
	} // fix point not found so return invalid index
	return -1
}

/* Binary search method */
func FixPoint2(arr []int, size int) int {
	low := 0
	high := size - 1
	var mid int
	for (low <= high) {
		mid = (low + high) / 2
		if (arr[mid] == mid) {
			return mid
		} else if (arr[mid] < mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	/* fix point not found so return invalid index */
	return -1
}

func subArraySums(arr []int, size int, value int) int {
	first := 0
	second := 0
	sum := arr[first]
	for (second < size && first < size) {
		if (sum == value) {
			fmt.Println(first + second)
		}

		if (sum < value) {
			second += 1
			if (second < size) {
				sum += arr[second]
			}
		} else {
			sum -= arr[first]
			first += 1
		}
	}
	return sum
}

func MaxConSub(arr []int, size int) int {
	currMax := 0
	maximum := 0
	for i := 0; i < size; i++ {
		currMax = max(arr[i], currMax + arr[i])
		if (currMax < 0) {
			currMax = 0
		}
		if (maximum < currMax) {
			maximum = currMax
		}
	}
	fmt.Println(maximum)
	return maximum
}

func MaxConSubArr( A []int, sizeA int, B []int, sizeB int) int {
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
			currMax = max(A[i], currMax + A[i])
		}
	}
	if (currMax < 0) {
		currMax = 0
	}
	if (maximum < currMax) {
		maximum = currMax
	}
	fmt.Println(maximum)
	return maximum
}

func MaxConSubArr2(A []int, sizeA int, B []int, sizeB int) int {
	sort.Ints(B)
	currMax := 0
	maximum := 0

	for i := 0; i < sizeA; i++ {
		if (Binarysearch(B, A[i])) {
			currMax = 0
		} else {
			currMax = max(A[i], currMax + A[i]);
			if (currMax < 0) {
				currMax = 0
			}
			if (maximum < currMax) {
				maximum = currMax
			}
		}
	}
	fmt.Println(maximum)
	return maximum
}

func RainWater(arr []int, size int) int {
	leftHigh := make([]int, size)
	rightHigh := make([]int, size)

	max := arr[0]
	leftHigh[0] = arr[0];
	for i := 1; i < size; i++ {
		if (max < arr[i]) {
			max = arr[i]
		}
		leftHigh[i] = max
	}
	max = arr[size - 1]
	rightHigh[size - 1] = arr[size - 1]
	for i := (size - 2); i >= 0; i-- {
		if (max < arr[i]) {
			max = arr[i]
		}
		rightHigh[i] = max
	}

	water := 0
	for i := 0; i < size; i++ {
		water += min(leftHigh[i], rightHigh[i]) - arr[i]
	}
	fmt.Println("Water : " , water)
	return water
}

func RainWater2(arr []int, size int) int {
	water := 0
	leftMax := 0
	rightMax := 0
	left := 0
	right := size - 1

	for (left <= right) {
		if (arr[left] < arr[right]) {
			if (arr[left] > leftMax) {
				leftMax = arr[left]
			} else {
				water += leftMax - arr[left]
			}
			left += 1
		} else {
			if (arr[right] > rightMax) {
				rightMax = arr[right]
			} else {
				water += rightMax - arr[right]
			}
			right -= 1
		}
	}
	fmt.Println("Water : " , water)
	return water
}

func main24() {
	first := []int {1, 5, 6, 6, 6, 6, 6, 6, 7, 8, 10, 13, 20, 30 }
	seperateEvenAndOdd(first)
	for _, val := range first {
		fmt.Print(val , " ")
	}
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
	main11()
	main12()
	main13()
	main14()
	main15()
	main16()
	main17()
	main18()
	main19()
	main20()
	//main21()
	main22()
	main23()
	main24()
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