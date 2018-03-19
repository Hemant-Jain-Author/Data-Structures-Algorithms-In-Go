package main

import "fmt"

func IncrementPassByValue(x int) {
	x++
}

func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}

func main1() {
	i := 10
	fmt.Println("Value of i before increment is :  ", i)
	IncrementPassByValue(i)
	fmt.Println("Value of i after increment is :  ", i)

	fmt.Println("Value of i before increment is :  ", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is :  ", i)
}

type coord struct {
	x int
	y int
}

func ArrayExample() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	PrintSlice(arr[:])
}

func SliceExample() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	PrintSlice(s)
}

func PrintSlice(data []int) {
	count := len(data)
	fmt.Print("Values stored are : ")
	for i := 0; i < count; i++ {
		fmt.Print(" ", data[i])
	}
	fmt.Println()
}

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func Permutation(data []int, i int, length int) {
	if length == i {
		PrintSlice(data)
		return
	}

	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func main2() {
	var data [5]int
	for i := 0; i < 5; i++ {
		data[i] = i
	}
	Permutation(data[:], 0, 5)
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

func function2() {
	fmt.Println("fun2 line 1")
}

func function1() {
	fmt.Println("fun1 line 1")
	function2()
	fmt.Println("fun1 line 2")
}

func main3() {
	fmt.Println("main line 1")
	function1()
	fmt.Println("main line 2")
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

func main4() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	fmt.Println("Max sub array sum :", MaxSubArraySum(data))
}

func VariableExample() {
	var1 := 100
	var2 := 200
	var3 := var1 + var2
	fmt.Println("Adding ", var1, " and ", var2, " will give ", var3)
}

func VectorExample() {
	var data [10]int
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	PrintSlice(data[:])
}

func twoDArrayExample() {
	var data [4][2]int
	count := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			count++
			data[i][j] = count
		}
	}
	print2DArray(data, 4, 2)
}

func print2DArray(data [4][2]int, R int, C int) {
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Println(" ", data[i][j])
		}
	}
}

func SumArray(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}

func main5() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("sum of all the values in array:", SumArray(data))
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

func main() {
	point := &coord{10, 10}
	fmt.Println("X axis coord value is  ", point.x)
	fmt.Println("Y axis coord value is  ", point.y)
}

func (s student) GetFirstName() string {
	return s.firstName
}

func (s student) GetLastName() string {
	return s.lastName
}

func (s student) GetRollNo() int {
	return s.rollNo
}

type student struct {
	rollNo    int
	firstName string
	lastName  string
}

func main77() {
	ptrStud := &student{1, "hemant", "jain"}
	fmt.Println("Roll No:   Student Name:  ", ptrStud.GetFirstName(), ptrStud.GetLastName(), ptrStud.GetRollNo())
}

func Sum(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func main77() {
	// calling a function to calculate sum
	result := Sum(10, 20)
	fmt.Println("Sum is : ", result)
}

func main7() {
	// local variable definition
	x := 10
	y := 20
	// calling a function to find sum
	result := Sum(x, y)
	fmt.Println("Sum is : ", result)
}

func Factorial(i int) int {
	// Termination Condition
	if i <= 1 {
		return 1
	}
	// Body, Recursive Expansion
	return i * Factorial(i-1)
}

func main(){
	fmt.Println("factorial 5 is :: ", Factorial(5))
}

func BinarySearchRecursive(data []int, low int, high int, value int) int {
	mid := low + (high-low)/2 // To afunc the overflow
	if data[mid] == value {
		return mid
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}
