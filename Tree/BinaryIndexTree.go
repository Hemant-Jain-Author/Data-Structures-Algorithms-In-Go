package main

import "fmt"

type BinaryIndexTree struct {
	BIT  []int
	size int
}

func NewBinaryIndexTree(arr []int) *BinaryIndexTree {
	t := &BinaryIndexTree{
		size: len(arr),
		BIT:  make([]int, len(arr)+1),
	}

	// Populating bit.
	for i := 0; i < t.size; i++ {
		t.Update(i, arr[i])
	}
	return t
}

func (t *BinaryIndexTree) Update(index int, val int) {
	// Index in bit is 1 more than the input array.
	index++

	// Traverse to ancestors of nodes.
	for index <= t.size {
		// Add val to the current node of the Binary Index Tree.
		t.BIT[index] += val
		// Next element that needs to store val.
		index += index & -index
	}
}

func (t *BinaryIndexTree) Set(arr []int, index int, val int) {
	diff := val - arr[index]
	arr[index] = val

	// Difference is propagated.
	t.Update(index, diff)
}

// Range sum in the range start to end.
func (t *BinaryIndexTree) RangeSum(start int, end int) int {
	// Check for error conditions.
	if start > end || start < 0 || end > t.size-1 {
		fmt.Println("Invalid input.")
		return -1
	}
	return t.PrefixSum(end) - t.PrefixSum(start-1)
}

// Prefix sum in the range 0 to index.
func (t *BinaryIndexTree) PrefixSum(index int) int {
	sum := 0
	index++

	// Traverse ancestors of Binary Index Tree nodes.
	for index > 0 {
		// Add the current element to the sum.
		sum += t.BIT[index]
		// Calculate the parent index.
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
