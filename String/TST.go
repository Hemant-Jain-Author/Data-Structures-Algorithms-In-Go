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

func (t *TST) insertUtil(currentNode *TSTNode, word string, currentIndex int) *TSTNode {
	if currentNode == nil {
		currentNode = &TSTNode{
			data: word[currentIndex],
		}
	}

	if word[currentIndex] < currentNode.data {
		currentNode.left = t.insertUtil(currentNode.left, word, currentIndex)
		return currentNode
	}

	if word[currentIndex] > currentNode.data {
		currentNode.right = t.insertUtil(currentNode.right, word, currentIndex)
		return currentNode
	}

	if currentIndex < len(word)-1 {
		currentNode.equal = t.insertUtil(currentNode.equal, word, currentIndex+1)
	} else {
		currentNode.isLastChar = true
	}

	return currentNode
}

func (t *TST) findUtil(currentNode *TSTNode, searchWord string, currentIndex int) bool {
	if currentNode == nil {
		return false
	}

	if searchWord[currentIndex] < currentNode.data {
		return t.findUtil(currentNode.left, searchWord, currentIndex)
	}

	if searchWord[currentIndex] > currentNode.data {
		return t.findUtil(currentNode.right, searchWord, currentIndex)
	}

	if currentIndex == len(searchWord)-1 {
		return currentNode.isLastChar
	}

	return t.findUtil(currentNode.equal, searchWord, currentIndex+1)
}

func (t *TST) Find(searchWord string) bool {
	return t.findUtil(t.root, searchWord, 0)
}

// Testing code.
func main() {
	t := new(TST)
	t.Insert("banana")
	t.Insert("apple")
	t.Insert("mango")
	fmt.Println("Apple Found:", t.Find("apple"))
	fmt.Println("Banana Found:", t.Find("banana"))
	fmt.Println("Grapes Found:", t.Find("grapes"))
}

/*
Apple Found: true
Banana Found: true
Grapes Found: false
*/
