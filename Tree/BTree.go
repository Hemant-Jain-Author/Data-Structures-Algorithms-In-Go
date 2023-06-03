package main

import (
	"fmt"
)

type BTree struct {
	root *Node // Pointer to root node
	max  int   // Maximum degree
	min  int   // Minimum degree
}

// Constructor
func NewBTree(dg int) *BTree {
	return &BTree{
		root: nil,
		max:  dg,     // Max number of children.
		min:  dg / 2, // Minimum number of children.
	}
}

type Node struct {
	n    int     // Current number of keys
	keys []int   // An array of keys
	arr  []*Node // An array of child pointers
	leaf bool    // Is true when node is leaf. Otherwise false
}

// Constructor
func NewNode(leaf bool, max int) *Node {
	return &Node{
		n:    0,
		keys: make([]int, max),
		arr:  make([]*Node, max+1),
		leaf: leaf,
	}
}

func (self *BTree) PrintTree() {
	self.printTree(self.root, "")
	fmt.Println()
}

func (self *BTree) printTree(node *Node, indent string) {
	if node == nil {
		return
	}
	for i := 0; i < node.n; i++ {
		self.printTree(node.arr[i], indent+"    ")
		fmt.Println(indent, "key[", i, "]:", node.keys[i])
	}
	self.printTree(node.arr[node.n], indent+"    ")
}

func (self *BTree) Search(key int) *Node {
	if self.root == nil {
		return nil
	}
	return self.searchUtil(self.root, key)
}

func (self *BTree) searchUtil(node *Node, key int) *Node {
	i := 0
	for i < node.n && node.keys[i] < key {
		i++
	}

	// If the found key is equal to key, return this node
	if node.keys[i] == key {
		return node
	}
	// If the key is not found and this is a leaf node
	if node.leaf == true {
		return nil
	}
	// Search in the appropriate child
	return self.searchUtil(node.arr[i], key)
}

func (self *BTree) Insert(key int) {
	// If tree is empty
	if self.root == nil {
		// Allocate memory for root
		self.root = NewNode(true, self.max)
		self.root.keys[0] = key // Insert key
		self.root.n = 1         // Update number of keys in root
		return
	}

	if self.root.leaf {
		// Finds the location where new key can be inserted.
		// By moving all keys greater than key to one place forward.
		i := self.root.n - 1
		for i >= 0 && self.root.keys[i] > key {
			self.root.keys[i+1] = self.root.keys[i]
			i--
		}
		// Insert the new key at found location
		self.root.keys[i+1] = key
		self.root.n++
	} else {
		i := 0
		for i < self.root.n && self.root.keys[i] < key {
			i++
		}
		self.insertUtil(self.root, self.root.arr[i], i, key)
	}

	if self.root.n == self.max {
		// If root contains more then allowed nodes, then tree grows in height.
		// Allocate memory for new root
		rt := NewNode(false, self.max)
		rt.arr[0] = self.root
		self.split(rt, self.root, 0)
		self.root = rt
	}
}

// Insert a new key in this node
// Arguments are parent, child, index of child and key.
func (self *BTree) insertUtil(parent *Node, child *Node, index int, key int) {
	if child.leaf == true {
		// Finds the location where new key will be inserted
		// by moving all keys greater than key to one place forward.
		i := child.n - 1
		for i >= 0 && child.keys[i] > key {
			child.keys[i+1] = child.keys[i]
			i--
		}
		// Insert the new key at found location
		child.keys[i+1] = key
		child.n++
	} else {
		// insert the node to the proper child.
		i := 0
		for i < child.n && child.keys[i] < key {
			i++
		}
		self.insertUtil(child, child.arr[i], i, key)
	}
	if child.n == self.max { // More nodes than allowed.
		// divide the child into two and then add the median to the parent.
		self.split(parent, child, index)
	}
}

func (self *BTree) split(parent *Node, child *Node, index int) {
	// Getting index of median.
	median := self.max / 2
	// Reduce the number of keys in child
	child.n = median

	node := NewNode(child.leaf, self.max)
	// Copy the second half keys of child to node
	j := 0
	for median+1+j < self.max {
		node.keys[j] = child.keys[median+1+j]
		j++
	}
	node.n = j

	// Copy the second half children of child to node
	j = 0
	for child.leaf == false && median+j <= self.max-1 {
		node.arr[j] = child.arr[median+1+j]
		j++
	}
	// parent is going to have a new child,
	// create space of new child
	for j = parent.n; j >= index+1; j-- {
		parent.arr[j+1] = parent.arr[j]
	}

	// Link the new child to the parent node
	parent.arr[index+1] = node

	// A key of child will move to the parent node.
	// Find the location of new key by moving
	// all greater keys one space forward.
	for j = parent.n - 1; j >= index; j-- {
		parent.keys[j+1] = parent.keys[j]
	}
	// Copy the middle key of child to the parent
	parent.keys[index] = child.keys[median]

	// Increment count of keys in this parent
	parent.n++
}

func (self *BTree) Remove(key int) {
	self.removeUtil(self.root, key)
	if self.root.n == 0 {
		// Shrinking : If root is pointing to empty node.
		// If that node is a leaf node then root will become null.
		// Else root will point to first child of node.
		if self.root.leaf {
			self.root = nil
		} else {
			self.root = self.root.arr[0]
		}
	}
}

func (self *BTree) removeUtil(node *Node, key int) {
	index := self.findKey(node, key)

	if node.leaf {
		if index < node.n && node.keys[index] == key {
			self.removeFromLeaf(node, index) // Leaf node key found.
		} else {
			fmt.Println("The key", key, "was not found.")
			return
		}
	} else {
		if index < node.n && node.keys[index] == key {
			self.removeFromNonLeaf(node, index) // Internal node key found.
		} else {
			self.removeUtil(node.arr[index], key)
		}

		// All the property violation in index subtree only.
		// which include remove from leaf case too.
		if node.arr[index].n < self.min {
			self.fixBTree(node, index)
		}
	}
}

// Returns the index of first key which is greater than or equal to key.
func (self *BTree) findKey(node *Node, key int) int {
	index := 0
	for index < node.n && node.keys[index] < key {
		index++
	}
	return index
}

// Remove the index key from leaf node.
func (self *BTree) removeFromLeaf(node *Node, index int) {
	// Move all the keys after the index position one step left.
	for i := index + 1; i < node.n; i++ {
		node.keys[i-1] = node.keys[i]
	}
	// Reduce the key count.
	node.n--
	return
}

// Remove the index key from a non-leaf node.
func (self *BTree) removeFromNonLeaf(node *Node, index int) {
	key := node.keys[index]

	if node.arr[index].n > self.min {
		// If the child that precedes key has at least min keys,
		// Find the predecessor 'pred' of key in the subtree rooted at index.
		// Replace key by pred and recursively delete pred in arr[index]
		pred := self.getPred(node, index)
		node.keys[index] = pred
		self.removeUtil(node.arr[index], pred)
	} else if node.arr[index+1].n > self.min {
		// If the child that succeeds key has at least min keys,
		// Find the successor 'succ' of key in the subtree rooted at index+1.
		// Replace key by succ and recursively delete succ in arr[index+1].
		succ := self.getSucc(node, index)
		node.keys[index] = succ
		self.removeUtil(node.arr[index+1], succ)
	} else {
		// If both left and right subtree has min number of keys.
		// Then merge left, right child along with parent key.
		// Then call remove on the merged child.
		self.merge(node, index)
		self.removeUtil(node.arr[index], key)
	}
	return
}

// To get predecessor of keys[index]
func (self *BTree) getPred(node *Node, index int) int {
	// Keep moving to the right most node of left subtree until we reach a leaf.
	cur := node.arr[index]
	for !cur.leaf {
		cur = cur.arr[cur.n]
	}
	// Return the last key of the leaf
	return cur.keys[cur.n-1]
}

// To get successor of keys[index]
func (self *BTree) getSucc(node *Node, index int) int {
	// Keep moving to the left most node of right subtree until we reach a leaf
	cur := node.arr[index+1]
	for !cur.leaf {
		cur = cur.arr[0]
	}
	// Return the first key of the leaf
	return cur.keys[0]
}

// Make sure that the node have at least min number of keys
func (self *BTree) fixBTree(node *Node, index int) {
	if index != 0 && node.arr[index-1].n > self.min {
		// If the left sibling has more than min keys.
		self.borrowFromLeft(node, index)
	} else if index != node.n && node.arr[index+1].n > self.min {
		// If the right sibling has more than min keys.
		self.borrowFromRight(node, index)
	} else {
		// If both siblings has less than min keys.
		// When right sibling exist always merge with the right sibling.
		// When right sibling does not exist then merge with left sibling.
		if index != node.n {
			self.merge(node, index)
		} else {
			self.merge(node, index-1)
		}
	}
}

// Move a key from parent to right and left to parent.
func (self *BTree) borrowFromLeft(node *Node, index int) {
	child := node.arr[index]
	sibling := node.arr[index-1]

	// Moving all key in child one step forward.
	for i := child.n - 1; i >= 0; i-- {
		child.keys[i+1] = child.keys[i]
	}
	// Move all its child pointers one step forward.
	for i := child.n; !child.leaf && i >= 0; i-- {
		child.arr[i+1] = child.arr[i]
	}
	// Setting child's first key equal to of the current node.
	child.keys[0] = node.keys[index-1]
	// Moving sibling's last child as child's first child.
	if !child.leaf {
		child.arr[0] = sibling.arr[sibling.n]
	}
	// Moving the key from the sibling to the parent
	node.keys[index-1] = sibling.keys[sibling.n-1]
	// Increase child key count and decrease sibling key count.
	child.n++
	sibling.n--
	return
}

// Move a key from parent to left and right to parent.
func (self *BTree) borrowFromRight(node *Node, index int) {
	child := node.arr[index]
	sibling := node.arr[index+1]

	// node key is inserted as the last key in child.
	child.keys[child.n] = node.keys[index]
	// Sibling's first child is inserted as the last child of child.
	if !child.leaf {
		child.arr[child.n+1] = sibling.arr[0]
	}
	//First key from sibling is inserted into node.
	node.keys[index] = sibling.keys[0]
	// Moving all keys in sibling one step left
	for i := 1; i < sibling.n; i++ {
		sibling.keys[i-1] = sibling.keys[i]
	}
	// Moving the child pointers one step behind
	for i := 1; !sibling.leaf && i <= sibling.n; i++ {
		sibling.arr[i-1] = sibling.arr[i]
	}
	// Increase child key count and decrease sibling key count.

	child.n++
	sibling.n--
	return
}

// Merge node's children at index and index+1.
func (self *BTree) merge(node *Node, index int) {
	left := node.arr[index]
	right := node.arr[index+1]
	start := left.n
	// Adding a key from node to the left child.
	left.keys[start] = node.keys[index]
	// Copying the keys from right to left.
	for i := 0; i < right.n; i++ {
		left.keys[start+1+i] = right.keys[i]
	}
	// Copying the child pointers from right to left.
	for i := 0; !left.leaf && i <= right.n; i++ {
		left.arr[start+1+i] = right.arr[i]
	}
	// Moving all keys after  index in the current node one step forward.
	for i := index + 1; i < node.n; i++ {
		node.keys[i-1] = node.keys[i]
	}
	// Moving the child pointers after (index+1) in the current node one step forward.
	for i := index + 2; i <= node.n; i++ {
		node.arr[i-1] = node.arr[i]
	}
	// Updating the key count of child and the current node
	left.n += right.n + 1
	node.n--
	return
}

func main() {
	t := NewBTree(3)
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)
	t.Insert(8)
	t.Insert(9)
	t.Insert(10)
	t.PrintTree()
	fmt.Println("6 : ", (t.Search(6) != nil))
	fmt.Println("11 : ", (t.Search(11) != nil))
	t.Remove(6)
	t.Remove(3)
	t.Remove(7)
	t.PrintTree()
}

/*
         key[ 0 ]: 1
     key[ 0 ]: 2
         key[ 0 ]: 3
 key[ 0 ]: 4
         key[ 0 ]: 5
     key[ 0 ]: 6
         key[ 0 ]: 7
     key[ 1 ]: 8
         key[ 0 ]: 9
         key[ 1 ]: 10

6 :  true
11 :  false
     key[ 0 ]: 1
     key[ 1 ]: 2
 key[ 0 ]: 4
     key[ 0 ]: 5
 key[ 1 ]: 8
     key[ 0 ]: 9
     key[ 1 ]: 10

*/
