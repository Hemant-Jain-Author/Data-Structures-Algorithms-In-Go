package main

import (
	"fmt"
)

type TreeNode struct {
	value       int
	left, right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Add(value int) {
	t.root = Add(t.root, value)
}

func Add(n *TreeNode, value int) *TreeNode {
	if n == nil {
		// Equivalent to return &Tree{value: value}.
		n = new(TreeNode)
		n.value = value
		return n
	}
	if value < n.value {
		n.left = Add(n.left, value)
	} else {
		n.right = Add(n.right, value)
	}
	return n
}

func (t *Tree) InOrder() {
	InOrder(t.root)
	fmt.Println()
}

func InOrder(n *TreeNode) {
	if n == nil {
		return
	}
	InOrder(n.left)
	fmt.Print(n.value, " ")
	InOrder(n.right)
}

func main() {
	t := new(Tree)
	t.Add(2)
	t.Add(1)
	t.Add(3)
	t.Add(4)
	t.InOrder()
}