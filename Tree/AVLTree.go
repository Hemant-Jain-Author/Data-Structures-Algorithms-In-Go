package main

import "fmt"

type AVLTree struct {
	root *Node
}

type Node struct {
	data, height int
	left, right  *Node
}

func NewNode(d int, l, r *Node) *Node {
	return &Node{
		data:   d,
		left:   l,
		right:  r,
		height: 0,
	}
}

func (avlTree *AVLTree) height(n *Node) int {
	if n == nil {
		return -1
	}
	return n.height
}

func (avlTree *AVLTree) getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return avlTree.height(node.left) - avlTree.height(node.right)
}

func (avlTree *AVLTree) Insert(data int) {
	avlTree.root = avlTree.insertUtil(avlTree.root, data)
}

func (avlTree *AVLTree) insertUtil(node *Node, data int) *Node {
	if node == nil {
		return NewNode(data, nil, nil)
	}
	if node.data > data {
		node.left = avlTree.insertUtil(node.left, data)
	} else if node.data < data {
		node.right = avlTree.insertUtil(node.right, data)
	} else { // Duplicate data not allowed
		return node
	}
	node.height = avlTree.max(avlTree.height(node.left), avlTree.height(node.right)) + 1
	balance := avlTree.getBalance(node)
	if balance > 1 {
		if data < node.left.data { // Left Left Case
			return avlTree.rightRotate(node)
		}
		if data > node.left.data { // Left Right Case
			return avlTree.leftRightRotate(node)
		}
	}
	if balance < -1 {
		if data > node.right.data { // Right Right Case
			return avlTree.leftRotate(node)
		}
		if data < node.right.data { // Right Left Case
			return avlTree.rightLeftRotate(node)
		}
	}
	return node
}

// Function to right rotate subtree rooted with x
func (avlTree *AVLTree) rightRotate(x *Node) *Node {
	y := x.left
	T := y.right

	// Rotation
	y.right = x
	x.left = T

	// Update heights
	x.height = avlTree.max(avlTree.height(x.left), avlTree.height(x.right)) + 1
	y.height = avlTree.max(avlTree.height(y.left), avlTree.height(y.right)) + 1

	// Return new root
	return y
}

// Function to left rotate subtree rooted with x
func (avlTree *AVLTree) leftRotate(x *Node) *Node {
	y := x.right
	T := y.left

	// Rotation
	y.left = x
	x.right = T

	// Update heights
	x.height = avlTree.max(avlTree.height(x.left), avlTree.height(x.right)) + 1
	y.height = avlTree.max(avlTree.height(y.left), avlTree.height(y.right)) + 1

	// Return new root
	return y
}

// Function to right then left rotate subtree rooted with x
func (avlTree *AVLTree) rightLeftRotate(x *Node) *Node {
	x.right = avlTree.rightRotate(x.right)
	return avlTree.leftRotate(x)
}

// Function to left then right rotate subtree rooted with x
func (avlTree *AVLTree) leftRightRotate(x *Node) *Node {
	x.left = avlTree.leftRotate(x.left)
	return avlTree.rightRotate(x)
}

func (avlTree *AVLTree) Delete(data int) {
	avlTree.root = avlTree.deleteUtil(avlTree.root, data)
}

func (avlTree *AVLTree) deleteUtil(node *Node, data int) *Node {
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
			minNode := avlTree.FindMin(node.right)
			node.data = minNode.data
			node.right = avlTree.deleteUtil(node.right, minNode.data)
		}
	} else {
		if node.data > data {
			node.left = avlTree.deleteUtil(node.left, data)
		} else {
			node.right = avlTree.deleteUtil(node.right, data)
		}
	}
	node.height = avlTree.max(avlTree.height(node.left), avlTree.height(node.right)) + 1
	balance := avlTree.getBalance(node)

	if balance > 1 {
		if data >= node.left.data { // Left Left Case
			return avlTree.rightRotate(node)
		}
		if data < node.left.data { // Left Right Case
			return avlTree.leftRightRotate(node)
		}
	}
	if balance < -1 {
		if data <= node.right.data { // Right Right Case
			return avlTree.leftRotate(node)
		}
		if data > node.right.data { // Right Left Case
			return avlTree.rightLeftRotate(node)
		}
	}
	return node
}

func (avlTree *AVLTree) FindMin(curr *Node) *Node {
	node := curr
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (avlTree *AVLTree) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (avlTree *AVLTree) PrintTree() {
	avlTree.printTreeUtil(avlTree.root, "", false)
	fmt.Println()
}

func (avlTree *AVLTree) printTreeUtil(node *Node, indent string, isLeft bool) {
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
	avlTree.printTreeUtil(node.left, indent, true)
	avlTree.printTreeUtil(node.right, indent, false)
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
