package main

import (
	"fmt"
	"math"
)

type SegmentTree struct {
	segArr []int
	size   int
}

func NewSegmentTree(input []int) (self *SegmentTree) {
	self = &SegmentTree{}
	self.size = len(input)
	// Height of segment tree.
	x := (math.Ceil(math.Log(float64(self.size)) / math.Log(2)))
	//Maximum size of segment tree
	maxSize := int(2*math.Pow(2.0, x) - 1)
	// Allocate memory for segment tree
	self.segArr = make([]int, maxSize)
	self.constructST(input, 0, self.size-1, 0)
	return
}

func (self *SegmentTree) constructST(input []int, start int, end int, index int) int {
	// Store it in current node of the segment tree and return
	if start == end {
		self.segArr[index] = input[start]
		return input[start]
	}
	// If there are more than one elements,
	// then traverse left and right subtrees
	// and store the sum of values in current node.
	mid := (start + end) / 2
	self.segArr[index] = self.constructST(input, start, mid, index*2+1) + self.constructST(input, mid+1, end, index*2+2)
	return self.segArr[index]
}

func (self *SegmentTree) GetSum(start int, end int) int {
	// Check for error conditions.
	if start > end || start < 0 || end > self.size-1 {
		fmt.Println("Invalid Input.")
		return -1
	}
	return self.getSumUtil(0, self.size-1, start, end, 0)
}

func (self *SegmentTree) getSumUtil(segStart int, segEnd int, queryStart int, queryEnd int, index int) int {
	if queryStart <= segStart && segEnd <= queryEnd {
		return self.segArr[index]
	}
	if segEnd < queryStart || queryEnd < segStart {
		return 0
	}
	// Segment tree is partly overlaps with the query range.
	mid := (segStart + segEnd) / 2
	return self.getSumUtil(segStart, mid, queryStart, queryEnd, 2*index+1) + self.getSumUtil(mid+1, segEnd, queryStart, queryEnd, 2*index+2)
}

func (self *SegmentTree) Set(arr []int, ind int, val int) {
	// Check for error conditions.
	if ind < 0 || ind > self.size-1 {
		fmt.Println("Invalid Input.")
		return
	}
	arr[ind] = val
	// Set new value in segment tree
	self.setUtil(0, self.size-1, ind, val, 0)
}

// Always diff will be returned.
func (self *SegmentTree) setUtil(segStart int, segEnd int, ind int, val int, index int) int {
	// set index lies outside the range of current segment.
	// So diff to its parent node will be zero.
	if ind < segStart || ind > segEnd {
		return 0
	}
	// If the input index is in range of this node, then set the
	// value of the node and its children
	if segStart == segEnd {
		if segStart == ind { // Index that needs to be set.
			diff := val - self.segArr[index]
			self.segArr[index] = val
			return diff
		} else {
			return 0
		}
	}
	mid := (segStart + segEnd) / 2
	diff := self.setUtil(segStart, mid, ind, val, 2*index+1) +
		self.setUtil(mid+1, segEnd, ind, val, 2*index+2)
	// Current node value is set with diff.
	self.segArr[index] = self.segArr[index] + diff
	// Value of diff is propagated to the parent node.
	return diff
}

func main() {
	arr := []int{1, 2, 4, 8, 16, 32, 64}
	tree := NewSegmentTree(arr)
	fmt.Println("Sum of values in the range(0, 3): ", tree.GetSum(1, 3))
	fmt.Println("Sum of values of all the elements: ", tree.GetSum(0, len(arr)-1))
	tree.Set(arr, 1, 10)
	fmt.Println("Sum of values in the range(0, 3): ", tree.GetSum(1, 3))
	fmt.Println("Sum of values of all the elements: ", tree.GetSum(0, len(arr)-1))
}

/*
Sum of values in the range(0, 3):  14
Sum of values of all the elements:  127
Sum of values in the range(0, 3):  22
Sum of values of all the elements:  135
*/
