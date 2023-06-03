package main

import (
	"fmt"
	"math"
)

type rmqST struct {
	segArr []int
	n      int
}

func NewRmqST(input []int) (self *rmqST) {
	self = &rmqST{}
	// Height of segment tree.
	self.n = len(input)
	// Maximum size of segment tree
	x := (math.Ceil(math.Log(float64(self.n)) / math.Log(2)))
	// Allocate memory for segment tree
	maxSize := int(2*math.Pow(2.0, x) - 1)
	self.segArr = make([]int, maxSize)
	self.constructST(input, 0, self.n-1, 0)
	return
}

func (self *rmqST) constructST(input []int, start int, end int, index int) int {
	// Store it in the current node of the segment tree and return
	if start == end {
		self.segArr[index] = input[start]
		return input[start]
	}
	// If there are more than one element,
	// then traverse left and right subtrees
	// and store the minimum of values in the current node.
	mid := (start + end) / 2
	self.segArr[index] = self.min(self.constructST(input, start, mid, index*2+1), self.constructST(input, mid+1, end, index*2+2))
	return self.segArr[index]
}

func (self *rmqST) Update(ind int, val int) {
	// Check for error conditions.
	if ind < 0 || ind > self.n-1 {
		fmt.Println("Invalid Input.")
		return
	}
	// Update the values in the segment tree
	self.updateUtil(0, self.n-1, ind, val, 0)
}

// Always min inside the valid range will be returned.
func (self *rmqST) updateUtil(segStart int, segEnd int, ind int, val int, index int) int {
	// Update index lies outside the range of the current segment.
	// So the minimum will not change.
	if ind < segStart || ind > segEnd {
		return self.segArr[index]
	}
	// If the input index is in the range of this node, then update the
	// value of the node and its children
	if segStart == segEnd {
		if segStart == ind { // Index value needs to be updated.
			self.segArr[index] = val
			return val
		} else {
			return self.segArr[index] // index value is not changed.
		}
	}
	mid := (segStart + segEnd) / 2
	// Current node value is updated with min.
	self.segArr[index] = self.min(self.updateUtil(segStart, mid, ind, val, 2*index+1), self.updateUtil(mid+1, segEnd, ind, val, 2*index+2))
	// Value of diff is propagated to the parent node.
	return self.segArr[index]
}

func (self *rmqST) GetMin(start int, end int) int {
	// Check for error conditions.
	if start > end || start < 0 || end > self.n-1 {
		fmt.Println("Invalid Input.")
		return math.MaxInt32
	}
	return self.getMinUtil(0, self.n-1, start, end, 0)
}

func (self *rmqST) getMinUtil(segStart int, segEnd int, queryStart int, queryEnd int, index int) int {
	if queryStart <= segStart && segEnd <= queryEnd {
		return self.segArr[index]
	}
	if segEnd < queryStart || queryEnd < segStart {
		return math.MaxInt32
	}
	// Segment tree partly overlaps with the query range.
	mid := (segStart + segEnd) / 2
	return self.min(self.getMinUtil(segStart, mid, queryStart, queryEnd, 2*index+1), self.getMinUtil(mid+1, segEnd, queryStart, queryEnd, 2*index+2))
}

func (self *rmqST) min(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

func main() {
	arr := []int{2, 3, 1, 7, 12, 5}
	tree := NewRmqST(arr)
	fmt.Println("Min value in the range(1, 5):", tree.GetMin(1, 5))
	fmt.Println("Min value of all the elements:", tree.GetMin(0, len(arr)-1))
	tree.Update(2, -1)
	fmt.Println("Min value in the range(1, 5):", tree.GetMin(1, 5))
	fmt.Println("Min value of all the elements:", tree.GetMin(0, len(arr)-1))
	tree.Update(5, -2)
	fmt.Println("Min value in the range(0, 4):", tree.GetMin(0, 4))
	fmt.Println("Min value of all the elements:", tree.GetMin(0, len(arr)-1))
}

/*
Min value in the range(1, 5): 1
Min value of all the elements: 1
Min value in the range(1, 5): -1
Min value of all the elements: -1
Min value in the range(0, 4): -1
Min value of all the elements: -2
*/
