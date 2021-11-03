package main

import (
	"fmt"
	"math"
)

type rangeMaxST struct {
	segArr []int
	n      int
}

func NewrangeMaxST(input []int) (self *rangeMaxST) {
	self = &rangeMaxST{}
	self.n = len(input)
	// Height of segment tree.
	x := (math.Ceil(math.Log(float64(self.n)) / math.Log(2)))
	//Maximum size of segment tree
	max_size := int(2*math.Pow(2.0, x) - 1)
	// Allocate memory for segment tree
	self.segArr = make([]int, max_size)
	self.constructST(input, 0, self.n-1, 0)
	return
}

func (self *rangeMaxST) constructST(input []int, start int, end int, index int) int {
	// Store it in current node of the segment tree and return
	if start == end {
		self.segArr[index] = input[start]
		return input[start]
	}
	// If there are more than one elements,
	// then traverse left and right subtrees
	// and store the minimum of values in current node.
	mid := (start + end) / 2
	self.segArr[index] = self.max(self.constructST(input, start, mid, index*2+1), self.constructST(input, mid+1, end, index*2+2))
	return self.segArr[index]
}

func (self *rangeMaxST) max(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func (self *rangeMaxST) GetMax(start int, end int) int {
	// Check for error conditions.
	if start > end || start < 0 || end > self.n-1 {
		fmt.Println("Invalid Input.")
		return math.MinInt32
	}
	return self.getMaxUtil(0, self.n-1, start, end, 0)
}

func (self *rangeMaxST) getMaxUtil(segStart int, segEnd int, queryStart int, queryEnd int, index int) int {
	if queryStart <= segStart && segEnd <= queryEnd {
		return self.segArr[index]
	}
	if segEnd < queryStart || queryEnd < segStart {
		return math.MinInt32
	}
	// Segment tree is partly overlaps with the query range.
	mid := (segStart + segEnd) / 2
	return self.max(self.getMaxUtil(segStart, mid, queryStart, queryEnd, 2*index+1), self.getMaxUtil(mid+1, segEnd, queryStart, queryEnd, 2*index+2))
}

func (self *rangeMaxST) Update(ind int, val int) {
	// Check for error conditions.
	if ind < 0 || ind > self.n-1 {
		fmt.Println("Invalid Input.")
		return
	}
	// Update the values in segment tree
	self.updateUtil(0, self.n-1, ind, val, 0)
}

// Always min inside valid range will be returned.
func (self *rangeMaxST) updateUtil(segStart int, segEnd int, ind int, val int, index int) int {
	// Update index lies outside the range of current segment.
	// So minimum will not change.
	if ind < segStart || ind > segEnd {
		return self.segArr[index]
	}
	// If the input index is in range of this node, then update the
	// value of the node and its children
	if segStart == segEnd {
		if segStart == ind { // Index value need to be updated.
			self.segArr[index] = val
			return val
		} else {
			return self.segArr[index] // index value is not changed.
		}
	}
	mid := (segStart + segEnd) / 2
	// Current node value is updated with min. 
    self.segArr[index] = self.max(self.updateUtil(segStart, mid, ind, val, 2*index+1), 
					self.updateUtil(mid+1, segEnd, ind, val, 2*index+2))
	// Value of diff is propagated to the parent node.
    return self.segArr[index]
}

func main() {
	arr := []int{1, 8, 2, 7, 3, 6, 4, 5}
	tree := NewrangeMaxST(arr)
	fmt.Println("Max value in the range(1, 5): ", tree.GetMax(1, 5))
	fmt.Println("Max value in the range(2, 7): ", tree.GetMax(2, 7))
	fmt.Println("Max value of all the elements: ", tree.GetMax(0, len(arr)-1))
	tree.Update(2, 9)
	fmt.Println("Max value in the range(1, 5): ", tree.GetMax(1, 5))
	fmt.Println("Max value of all the elements: ", tree.GetMax(0, len(arr)-1))
}

/*
Max value in the range(1, 5):  8
Max value in the range(2, 7):  7
Max value of all the elements:  8
Max value in the range(1, 5):  9
Max value of all the elements:  9
*/