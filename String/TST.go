package main

import "fmt"

type TSTNode struct {
	data               byte
	isLastChar         bool
	left, equal, right *TSTNode
}

type TST struct {
	root *TSTNode
}

func (t *TST) Insert(word string) {
	t.root = t.insertUtil(t.root, word, 0)
}

func (t *TST) insertUtil(curr *TSTNode, word string, wordIndex int) *TSTNode {
	if curr == nil {
		curr = new(TSTNode)
		curr.data = word[wordIndex]
	}
	if word[wordIndex] < curr.data {
		curr.left = t.insertUtil(curr.left, word, wordIndex)
	} else if word[wordIndex] > curr.data {
		curr.right = t.insertUtil(curr.right, word, wordIndex)
	} else {
		if wordIndex < len(word)-1 {
			curr.equal = t.insertUtil(curr.equal, word, wordIndex+1)
		} else {
			curr.isLastChar = true
		}
	}
	return curr
}

func (t *TST) findUtil(curr *TSTNode, word string, wordIndex int) bool {
	if curr == nil {
		return false
	}
	if word[wordIndex] < curr.data {
		return t.findUtil(curr.left, word, wordIndex)
	} else if word[wordIndex] > curr.data {
		return t.findUtil(curr.right, word, wordIndex)
	} else {
		if wordIndex == len(word)-1 {
			return curr.isLastChar
		}
		return t.findUtil(curr.equal, word, wordIndex+1)
	}
}

func (t *TST) Find(word string) bool {
	ret := t.findUtil(t.root, word, 0)
	fmt.Print(word, " :: ")
	if ret {
		fmt.Println(" Found ")
	} else {
		fmt.Println("Not Found ")
	}
	return ret
}

func main() {
	tt := new(TST)
	tt.Insert("banana")
	tt.Insert("apple")
	tt.Insert("mango")
	fmt.Println("Search results for apple, banana, grapes and mango :")
	tt.Find("apple")
	tt.Find("banana")
	tt.Find("mango")
	tt.Find("grapes")
}
