package main

import (
	"fmt"
	"math"
)

type AVLTree struct {
	root *Node
}

type Node struct {
	data, height int
	left, right  *Node
}

func NewNode(data int, left, right *Node) *Node {
	return &Node{
		data:   data,
		left:   left,
		right:  right,
		height: 0,
	}
}

func (t *AVLTree) height(n *Node) int {
	if n == nil {
		return -1
	}
	return n.height
}

func (t *AVLTree) getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return t.height(node.left) - t.height(node.right)
}

func (t *AVLTree) Insert(data int) {
	t.root = t.insertUtil(t.root, data)
}

func (t *AVLTree) insertUtil(node *Node, data int) *Node {
	if node == nil {
		return NewNode(data, nil, nil)
	}
	if node.data > data {
		node.left = t.insertUtil(node.left, data)
	} else if node.data < data {
		node.right = t.insertUtil(node.right, data)
	} else { // Duplicate data not allowed
		return node
	}
	node.height = int(math.Max(float64(t.height(node.left)), float64(t.height(node.right)))) + 1
	balance := t.getBalance(node)
	if balance > 1 {
		if data < node.left.data { // Left Left Case
			return t.rightRotate(node)
		}
		if data > node.left.data { // Left Right Case
			return t.leftRightRotate(node)
		}
	}
	if balance < -1 {
		if data > node.right.data { // Right Right Case
			return t.leftRotate(node)
		}
		if data < node.right.data { // Right Left Case
			return t.rightLeftRotate(node)
		}
	}
	return node
}

// Function to right rotate subtree rooted with x
func (t *AVLTree) rightRotate(x *Node) *Node {
	y := x.left
	T := y.right

	// Rotation
	y.right = x
	x.left = T

	// Update heights
	x.height = int(math.Max(float64(t.height(x.left)), float64(t.height(x.right)))) + 1
	y.height = int(math.Max(float64(t.height(y.left)), float64(t.height(y.right)))) + 1

	// Return new root
	return y
}

// Function to left rotate subtree rooted with x
func (t *AVLTree) leftRotate(x *Node) *Node {
	y := x.right
	T := y.left

	// Rotation
	y.left = x
	x.right = T

	// Update heights
	x.height = int(math.Max(float64(t.height(x.left)), float64(t.height(x.right)))) + 1
	y.height = int(math.Max(float64(t.height(y.left)), float64(t.height(y.right)))) + 1

	// Return new root
	return y
}

// Function to right then left rotate subtree rooted with x
func (t *AVLTree) rightLeftRotate(x *Node) *Node {
	x.right = t.rightRotate(x.right)
	return t.leftRotate(x)
}

// Function to left then right rotate subtree rooted with x
func (t *AVLTree) leftRightRotate(x *Node) *Node {
	x.left = t.leftRotate(x.left)
	return t.rightRotate(x)
}

func (t *AVLTree) Delete(data int) {
	t.root = t.deleteUtil(t.root, data)
}

func (t *AVLTree) deleteUtil(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if node.data == data {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil {
			return node.right // no need to modify height
		} else if node.right == nil {
			return node.left // no need to modify height
		} else {
			minNode := t.FindMin(node.right)
			node.data = minNode.data
			node.right = t.deleteUtil(node.right, minNode.data)
		}
	} else {
		if node.data > data {
			node.left = t.deleteUtil(node.left, data)
		} else {
			node.right = t.deleteUtil(node.right, data)
		}
	}
	node.height = int(math.Max(float64(t.height(node.left)), float64(t.height(node.right)))) + 1
	balance := t.getBalance(node)
	if balance > 1 {
		if data >= node.left.data { // Left Left Case
			return t.rightRotate(node)
		}
		if data < node.left.data { // Left Right Case
			return t.leftRightRotate(node)
		}
	}
	if balance < -1 {
		if data <= node.right.data { // Right Right Case
			return t.leftRotate(node)
		}
		if data > node.right.data { // Right Left Case
			return t.rightLeftRotate(node)
		}
	}
	return node
}

func (t *AVLTree) FindMin(curr *Node) *Node {
	node := curr
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *AVLTree) PrintTree() {
	t.printTreeUtil(t.root, "", false)
	fmt.Println()
}

func (t *AVLTree) printTreeUtil(node *Node, indent string, isLeft bool) {
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
	fmt.Println(node.data, "(", node.height, ")")
	t.printTreeUtil(node.left, indent, true)
	t.printTreeUtil(node.right, indent, false)
}

func main() {
	t := AVLTree{}
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)
	t.Insert(8)
	t.PrintTree()
	t.Delete(5)
	t.PrintTree()
	t.Delete(1)
	t.PrintTree()
	t.Delete(2)
	t.PrintTree()
}

/*
R:4 ( 3 )
   L:2 ( 1 )
   |  L:1 ( 0 )
   |  R:3 ( 0 )
   R:6 ( 2 )
      L:5 ( 0 )
      R:7 ( 1 )
         R:8 ( 0 )

R:4 ( 2 )
   L:2 ( 1 )
   |  L:1 ( 0 )
   |  R:3 ( 0 )
   R:7 ( 1 )
      L:6 ( 0 )
      R:8 ( 0 )

R:4 ( 2 )
   L:2 ( 1 )
   |  R:3 ( 0 )
   R:7 ( 1 )
      L:6 ( 0 )
      R:8 ( 0 )

R:4 ( 2 )
   L:3 ( 0 )
   R:7 ( 1 )
      L:6 ( 0 )
      R:8 ( 0 )
*/
