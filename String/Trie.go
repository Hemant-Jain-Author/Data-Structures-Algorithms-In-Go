package main

import (
	"fmt"
	"strings"
)

const numChars = 26

type TrieNode struct {
	isLastChar bool
	child      [numChars]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(s string) {
	if s == "" {
		return
	}

	str := strings.ToLower(s)
	t.root = t.insertUtil(t.root, str, 0)
}

func (t *Trie) insertUtil(curr *TrieNode, str string, index int) *TrieNode {
	if curr == nil {
		curr = &TrieNode{}
	}
	if index == len(str) {
		curr.isLastChar = true
	} else {
		charIndex := str[index] - 'a'
		curr.child[charIndex] = t.insertUtil(curr.child[charIndex], str, index+1)
	}
	return curr
}

func (t *Trie) Remove(s string) {
	if s == "" {
		return
	}

	str := strings.ToLower(s)
	t.removeUtil(t.root, str, 0)
}

func (t *Trie) removeUtil(curr *TrieNode, str string, index int) {
	if curr == nil {
		return
	}

	if index == len(str) {
		curr.isLastChar = false
		return
	}
	charIndex := str[index] - 'a'
	t.removeUtil(curr.child[charIndex], str, index+1)
}

func (t *Trie) Find(s string) bool {
	if s == "" {
		return false
	}
	str := strings.ToLower(s)
	return t.findUtil(t.root, str, 0)
}

func (t *Trie) findUtil(curr *TrieNode, str string, index int) bool {
	if curr == nil {
		return false
	}
	if index == len(str) {
		return curr.isLastChar
	}
	charIndex := str[index] - 'a'
	return t.findUtil(curr.child[charIndex], str, index+1)
}

// Testing code.
func main() {
	t := &Trie{}
	t.Insert("banana")
	t.Insert("apple")
	t.Insert("mango")
	fmt.Println("Apple Found:", t.Find("apple"))
	fmt.Println("Grapes Found:", t.Find("grapes"))
	fmt.Println("Banana Found:", t.Find("banana"))
}
