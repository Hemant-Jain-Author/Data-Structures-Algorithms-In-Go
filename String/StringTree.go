package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	value  string
	count  int
	lChild *TreeNode
	rChild *TreeNode
}

type StringTree struct {
	root *TreeNode
}

func (t *StringTree) print() {
	t.printUtil(t.root)
}

func (t *StringTree) printUtil(curr *TreeNode) {
	if curr != nil {
		fmt.Println(" value is ::", curr.value)
		fmt.Println(" count is :: ", curr.count)
		t.printUtil(curr.lChild)
		t.printUtil(curr.rChild)
	}
}

func (t *StringTree) Insert(value string) {
	t.root = t.insertUtil(value, t.root)
}

func (t *StringTree) insertUtil(value string, curr *TreeNode) *TreeNode {
	var compare int
	if curr == nil {
		curr = new(TreeNode)
		curr.value = value
	} else {
		compare = strings.Compare(curr.value, value)
		if compare == 0 {
			curr.count++
		} else if compare == 1 {
			curr.lChild = t.insertUtil(value, curr.lChild)
		} else {
			curr.rChild = t.insertUtil(value, curr.rChild)
		}
	}
	return curr
}

func (t *StringTree) freeTree() {
	t.root = nil
}

func (t *StringTree) Find(value string) bool {
	return t.findUtil(t.root, value)
}

func (t *StringTree) findUtil(curr *TreeNode, value string) bool {
	var compare int
	if curr == nil {
		return false
	}
	compare = strings.Compare(curr.value, value)
	if compare == 0 {
		return true
	}

	if compare == 1 {
		return t.findUtil(curr.lChild, value)
	}
	return t.findUtil(curr.rChild, value)
}

func (t *StringTree) frequency(value string) int {
	return t.frequencyUtil(t.root, value)
}

func (t *StringTree) frequencyUtil(curr *TreeNode, value string) int {
	var compare int
	if curr == nil {
		return 0
	}
	compare = strings.Compare(curr.value, value)

	if compare == 0 {
		return curr.count
	}
	if compare > 0 {
		return t.frequencyUtil(curr.lChild, value)
	}
	return t.frequencyUtil(curr.rChild, value)
}

func main() {
	t := new(StringTree)
	t.Insert("banana")
	t.Insert("apple")
	t.Insert("mango")
	fmt.Println("Apple Found :", t.Find("apple"))
	fmt.Println("Grapes Found :", t.Find("grapes"))
	fmt.Println("Banana Found :", t.Find("banana"))
}

/*
Apple Found : true
Grapes Found : false
Banana Found : true
*/
