package main

import "fmt"

type BinaryIndexTree struct {
	BIT  []int
	size int
}

func NewBinaryIndexTree(arr []int) (self *BinaryIndexTree) {
	self = &BinaryIndexTree{}
	self.size = len(arr)
	self.BIT = make([]int, self.size+1)

	// Populating bit.
	for i := 0; i < self.size; i++ {
		self.Update(i, arr[i])
	}
	return
}

func (self *BinaryIndexTree) Update(index int, val int) {
	// Index in bit is 1 more than the input array.
	index = index + 1

	// Traverse to ancestors of nodes.
	for index <= self.size {
		// Add val to current node of Binary Index Tree.
		self.BIT[index] += val
		// Next element which need to store val.
		index += index & -index
	}
}

func (self *BinaryIndexTree) Set(arr []int, index int, val int) {
	diff := val - arr[index]
	arr[index] = val

	// Difference is propagated.
	self.Update(index, diff)
}

// Range sum in the range start to end.
func (self *BinaryIndexTree) RangeSum(start int, end int) int {
	// Check for error conditions.
	if start > end || start < 0 || end > self.size-1 {
		fmt.Println("Invalid Input.")
		return -1
	}
	return self.PrefixSum(end) - self.PrefixSum(start-1)
}

// Prefix sum in the range 0 to index.
func (self *BinaryIndexTree) PrefixSum(index int) int {
	sum := 0
	index = index + 1

	// Traverse ancestors of Binary Index Tree nodes.
	for index > 0 {
		// Add current element to sum.
		sum += self.BIT[index]
		// Parent index calculation.
		index -= index & -index
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	tree := NewBinaryIndexTree(arr)
	fmt.Println("Sum of elements in range(0, 5): ", tree.PrefixSum(5))
	fmt.Println("Sum of elements in range(2, 5): ", tree.RangeSum(2, 5))
	tree.Set(arr, 3, 10)
	fmt.Println("Sum of elements in range(0, 5): ", tree.PrefixSum(5))
}
