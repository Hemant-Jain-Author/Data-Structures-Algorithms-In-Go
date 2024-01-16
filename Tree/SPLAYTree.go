package main

import "fmt"

type SPLAYTree struct {
	root *Node
}

func NewSPLAYTree() (self *SPLAYTree) {
	self = &SPLAYTree{}
	self.root = nil
	return
}

type Node struct {
	data                int
	left, right, parent *Node
}

func NewNode(d int, l, r *Node) (nd *Node) {
	nd = &Node{}
	nd.data = d
	nd.left = l
	nd.right = r
	nd.parent = nil
	return
}

func (self *SPLAYTree) PrintInOrder() {
	self.printInOrder(self.root)
	fmt.Println()
}

func (self *SPLAYTree) printInOrder(node *Node) {
	if node != nil {
		self.printInOrder(node.left)
		fmt.Print(node.data, " ")
		self.printInOrder(node.right)
	}
}

func (self *SPLAYTree) PrintTree() {
	self.printTree(self.root, "", false)
	fmt.Println()
}

func (self *SPLAYTree) printTree(node *Node, indent string, isLeft bool) {
	if node == nil {
		return
	}
	if isLeft {
		fmt.Print(indent, "L:")
		indent += "|  "
	} else {
		fmt.Print(indent, "R:")
		indent += "   "
	}
	fmt.Println(node.data)
	self.printTree(node.left, indent, true)
	self.printTree(node.right, indent, false)
}

// Function to right rotate subtree rooted with x
func (self *SPLAYTree) rightRotate(x *Node) *Node {
	y := x.left
	T := y.right

	// Rotation
	y.parent = x.parent
	y.right = x
	x.parent = y
	x.left = T
	if T != nil {
		T.parent = x
	}
	if y.parent != nil && y.parent.left == x {
		y.parent.left = y
	} else if y.parent != nil && y.parent.right == x {
		y.parent.right = y
	}
	// Return new root
	return y
}

// Function to left rotate subtree rooted with x
func (self *SPLAYTree) leftRotate(x *Node) *Node {
	y := x.right
	T := y.left

	// Rotation
	y.parent = x.parent
	y.left = x
	x.parent = y
	x.right = T
	if T != nil {
		T.parent = x
	}
	if y.parent != nil && y.parent.left == x {
		y.parent.left = y
	} else if y.parent != nil && y.parent.right == x {
		y.parent.right = y
	}
	// Return new root
	return y
}

func (self *SPLAYTree) parent(node *Node) *Node {
	if node == nil || node.parent == nil {
		return nil
	}
	return node.parent
}

func (self *SPLAYTree) splay(node *Node) {
	var parent *Node
	var grand *Node
	for node != self.root {
		parent = self.parent(node)
		grand = self.parent(parent)
		if parent == nil { // rotations had created new root, always last condition.
			self.root = node
		} else if grand == nil { // single rotation case.
			if parent.left == node {
				node = self.rightRotate(parent)
			} else {
				node = self.leftRotate(parent)
			}
		} else if grand.left == parent && parent.left == node { // Zig Zig case.
			self.rightRotate(grand)
			node = self.rightRotate(parent)
		} else if grand.right == parent && parent.right == node { // Zag Zag case.
			self.leftRotate(grand)
			node = self.leftRotate(parent)
		} else if grand.left == parent && parent.right == node { //Zig Zag case.
			self.leftRotate(parent)
			node = self.rightRotate(grand)
		} else if grand.right == parent && parent.left == node { // Zag Zig case.
			self.rightRotate(parent)
			node = self.leftRotate(grand)
		}
	}
}

func (self *SPLAYTree) Find(data int) bool {
	curr := self.root
	for curr != nil {
		if curr.data == data {
			self.splay(curr)
			return true
		} else if curr.data > data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (self *SPLAYTree) Insert(data int) {
	newNode := NewNode(data, nil, nil)
	if self.root == nil {
		self.root = newNode
		return
	}
	node := self.root
	var parent *Node = nil
	for node != nil {
		parent = node
		if node.data > data {
			node = node.left
		} else if node.data < data {
			node = node.right
		} else {
			self.splay(node) // duplicate insertion not allowed but splaying for it.
			return
		}
	}
	newNode.parent = parent
	if parent.data > data {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	self.splay(newNode)
}

func (self *SPLAYTree) FindMinNode(curr *Node) *Node {
	node := curr
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (self *SPLAYTree) Delete(data int) {
	node := self.root
	var parent *Node = nil
	var next *Node = nil
	for node != nil {
		if node.data == data {
			parent = node.parent
			if node.left == nil && node.right == nil {
				next = nil
			} else if node.left == nil {
				next = node.right
			} else if node.right == nil {
				next = node.left
			}
			if node.left == nil || node.right == nil {
				if node == self.root {
					self.root = next
					return
				}
				if parent.left == node {
					parent.left = next
				} else {
					parent.right = next
				}
				if next != nil {
					next.parent = parent
				}
				break
			}
			minNode := self.FindMinNode(node.right)
			data = minNode.data
			node.data = data
			node = node.right
		} else if node.data > data {
			parent = node
			node = node.left
		} else {
			parent = node
			node = node.right
		}
	}
	self.splay(parent) // Splaying for the parent of the node deleted.
}

// Testing code.
func main() {
	tree := NewSPLAYTree()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	tree.PrintTree()
	fmt.Println("Value 2 found:", tree.Find(2))
	tree.Delete(2)
	tree.Delete(5)
	tree.PrintTree()
}

/*
R:3
   L:2
   |  L:1
   R:6
      L:4
      |  R:5

Value 2 found: true
R:4
   L:3
   |  L:1
   R:6
*/
